module github.com/NpoolPlatform/api-manager

go 1.16

require (
	entgo.io/ent v0.10.1
	github.com/BurntSushi/toml v0.4.1 // indirect
	github.com/NpoolPlatform/go-service-framework v0.0.0-20220120091626-4e8035637592
	github.com/NpoolPlatform/libent-cruder v0.0.0-20220526050249-956b54fac9f1
	github.com/NpoolPlatform/message v0.0.0-20220530011435-281c2b6aedd4
	github.com/go-resty/resty/v2 v2.7.0
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.8.0
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/miekg/dns v1.1.43 // indirect
	github.com/prometheus/common v0.30.0 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.1-0.20210427113832-6241f9ab9942
	github.com/ugorji/go v1.2.6 // indirect
	github.com/urfave/cli/v2 v2.3.0
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/grpc v1.46.2
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.2.0
	google.golang.org/protobuf v1.28.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.41.0
