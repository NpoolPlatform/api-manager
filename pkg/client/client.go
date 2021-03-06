package client

import (
	"context"
	"fmt"
	"reflect"
	"time"
	"unsafe"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	config "github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	apimgr "github.com/NpoolPlatform/message/npool/apimgr"

	constant "github.com/NpoolPlatform/api-manager/pkg/message/const"

	"google.golang.org/grpc"
)

func reliableRegister(apis *apimgr.ServiceApis) {
	for {
		conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
		if err != nil {
			logger.Sugar().Errorf("fail get api manager connection: %v", err)
			time.Sleep(time.Minute)
			continue
		}

		cli := apimgr.NewApiManagerClient(conn)

		ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)

		_, err = cli.Register(ctx, &apimgr.RegisterRequest{
			Info: apis,
		})
		if err == nil {
			cancel()
			conn.Close()
			return
		}

		logger.Sugar().Errorf("fail register apis: %v", err)
		time.Sleep(time.Minute)

		cancel()
		conn.Close()
	}
}

func MuxApis(mux *runtime.ServeMux) *apimgr.ServiceApis {
	apis := &apimgr.ServiceApis{
		ServiceName: config.GetStringValueWithNameSpace("", config.KeyHostname),
		Protocol:    "http",
	}

	valueOfMux := reflect.ValueOf(mux).Elem()
	handlers := valueOfMux.FieldByName("handlers")
	methIter := handlers.MapRange()
	for methIter.Next() {
		for i := 0; i < methIter.Value().Len(); i++ {
			pat := methIter.Value().Index(i).FieldByName("pat")
			tmp := reflect.NewAt(pat.Type(), unsafe.Pointer(pat.UnsafeAddr())).Elem()
			str := tmp.MethodByName("String").Call(nil)[0].String()
			apis.Paths = append(apis.Paths, &apimgr.Path{
				Method:     methIter.Key().String(),
				Path:       str,
				MethodName: "NONE",
			})
		}
	}

	return apis
}

func Register(mux *runtime.ServeMux) {
	apis := MuxApis(mux)
	go reliableRegister(apis)
}

func GrpcApis(server grpc.ServiceRegistrar) *apimgr.ServiceApis {
	srvInfo := server.(*grpc.Server).GetServiceInfo()
	apis := &apimgr.ServiceApis{
		ServiceName: config.GetStringValueWithNameSpace("", config.KeyHostname),
		Protocol:    "grpc",
	}

	for key, info := range srvInfo {
		for _, method := range info.Methods {
			apis.Paths = append(apis.Paths, &apimgr.Path{
				Method:     "NONE",
				Path:       fmt.Sprintf("%v/%v", key, method.Name),
				MethodName: method.Name,
			})
		}
	}

	return apis
}

func RegisterGRPC(server grpc.ServiceRegistrar) {
	apis := GrpcApis(server)
	go reliableRegister(apis)
}

func do(ctx context.Context, fn func(_ctx context.Context, cli apimgr.ApiManagerClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get stock connection: %v", err)
	}
	defer conn.Close()

	cli := apimgr.NewApiManagerClient(conn)

	return fn(_ctx, cli)
}

func GetServiceMethodAPI(ctx context.Context, serviceName, methodName string) (*apimgr.ServicePath, error) {
	info, err := do(ctx, func(_ctx context.Context, cli apimgr.ApiManagerClient) (cruder.Any, error) {
		resp, err := cli.GetServiceMethodAPI(ctx, &apimgr.GetServiceMethodAPIRequest{
			ServiceName: serviceName,
			MethodName:  methodName,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get api: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get api: %v", err)
	}
	return info.(*apimgr.ServicePath), nil
}
