// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/api-manager/pkg/db/ent/serviceapi"
	"github.com/google/uuid"
)

// ServiceAPI is the model entity for the ServiceAPI schema.
type ServiceAPI struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Domain holds the value of the "domain" field.
	Domain string `json:"domain,omitempty"`
	// Method holds the value of the "method" field.
	Method string `json:"method,omitempty"`
	// Path holds the value of the "path" field.
	Path string `json:"path,omitempty"`
	// Exported holds the value of the "exported" field.
	Exported bool `json:"exported,omitempty"`
	// PathPrefix holds the value of the "path_prefix" field.
	PathPrefix string `json:"path_prefix,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ServiceAPI) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case serviceapi.FieldExported:
			values[i] = new(sql.NullBool)
		case serviceapi.FieldCreateAt, serviceapi.FieldUpdateAt, serviceapi.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case serviceapi.FieldDomain, serviceapi.FieldMethod, serviceapi.FieldPath, serviceapi.FieldPathPrefix:
			values[i] = new(sql.NullString)
		case serviceapi.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ServiceAPI", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ServiceAPI fields.
func (sa *ServiceAPI) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case serviceapi.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				sa.ID = *value
			}
		case serviceapi.FieldDomain:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field domain", values[i])
			} else if value.Valid {
				sa.Domain = value.String
			}
		case serviceapi.FieldMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field method", values[i])
			} else if value.Valid {
				sa.Method = value.String
			}
		case serviceapi.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				sa.Path = value.String
			}
		case serviceapi.FieldExported:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field exported", values[i])
			} else if value.Valid {
				sa.Exported = value.Bool
			}
		case serviceapi.FieldPathPrefix:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path_prefix", values[i])
			} else if value.Valid {
				sa.PathPrefix = value.String
			}
		case serviceapi.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				sa.CreateAt = uint32(value.Int64)
			}
		case serviceapi.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				sa.UpdateAt = uint32(value.Int64)
			}
		case serviceapi.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				sa.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ServiceAPI.
// Note that you need to call ServiceAPI.Unwrap() before calling this method if this ServiceAPI
// was returned from a transaction, and the transaction was committed or rolled back.
func (sa *ServiceAPI) Update() *ServiceAPIUpdateOne {
	return (&ServiceAPIClient{config: sa.config}).UpdateOne(sa)
}

// Unwrap unwraps the ServiceAPI entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sa *ServiceAPI) Unwrap() *ServiceAPI {
	tx, ok := sa.config.driver.(*txDriver)
	if !ok {
		panic("ent: ServiceAPI is not a transactional entity")
	}
	sa.config.driver = tx.drv
	return sa
}

// String implements the fmt.Stringer.
func (sa *ServiceAPI) String() string {
	var builder strings.Builder
	builder.WriteString("ServiceAPI(")
	builder.WriteString(fmt.Sprintf("id=%v", sa.ID))
	builder.WriteString(", domain=")
	builder.WriteString(sa.Domain)
	builder.WriteString(", method=")
	builder.WriteString(sa.Method)
	builder.WriteString(", path=")
	builder.WriteString(sa.Path)
	builder.WriteString(", exported=")
	builder.WriteString(fmt.Sprintf("%v", sa.Exported))
	builder.WriteString(", path_prefix=")
	builder.WriteString(sa.PathPrefix)
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", sa.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", sa.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", sa.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// ServiceAPIs is a parsable slice of ServiceAPI.
type ServiceAPIs []*ServiceAPI

func (sa ServiceAPIs) config(cfg config) {
	for _i := range sa {
		sa[_i].config = cfg
	}
}
