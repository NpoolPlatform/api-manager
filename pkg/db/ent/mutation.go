// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"github.com/NpoolPlatform/api-manager/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/api-manager/pkg/db/ent/serviceapi"
	"github.com/google/uuid"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeServiceAPI = "ServiceAPI"
)

// ServiceAPIMutation represents an operation that mutates the ServiceAPI nodes in the graph.
type ServiceAPIMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	domain        *string
	method        *string
	_path         *string
	exported      *bool
	path_prefix   *string
	create_at     *uint32
	addcreate_at  *uint32
	update_at     *uint32
	addupdate_at  *uint32
	delete_at     *uint32
	adddelete_at  *uint32
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*ServiceAPI, error)
	predicates    []predicate.ServiceAPI
}

var _ ent.Mutation = (*ServiceAPIMutation)(nil)

// serviceapiOption allows management of the mutation configuration using functional options.
type serviceapiOption func(*ServiceAPIMutation)

// newServiceAPIMutation creates new mutation for the ServiceAPI entity.
func newServiceAPIMutation(c config, op Op, opts ...serviceapiOption) *ServiceAPIMutation {
	m := &ServiceAPIMutation{
		config:        c,
		op:            op,
		typ:           TypeServiceAPI,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withServiceAPIID sets the ID field of the mutation.
func withServiceAPIID(id uuid.UUID) serviceapiOption {
	return func(m *ServiceAPIMutation) {
		var (
			err   error
			once  sync.Once
			value *ServiceAPI
		)
		m.oldValue = func(ctx context.Context) (*ServiceAPI, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().ServiceAPI.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withServiceAPI sets the old ServiceAPI of the mutation.
func withServiceAPI(node *ServiceAPI) serviceapiOption {
	return func(m *ServiceAPIMutation) {
		m.oldValue = func(context.Context) (*ServiceAPI, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ServiceAPIMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ServiceAPIMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of ServiceAPI entities.
func (m *ServiceAPIMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ServiceAPIMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetDomain sets the "domain" field.
func (m *ServiceAPIMutation) SetDomain(s string) {
	m.domain = &s
}

// Domain returns the value of the "domain" field in the mutation.
func (m *ServiceAPIMutation) Domain() (r string, exists bool) {
	v := m.domain
	if v == nil {
		return
	}
	return *v, true
}

// OldDomain returns the old "domain" field's value of the ServiceAPI entity.
// If the ServiceAPI object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ServiceAPIMutation) OldDomain(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDomain is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDomain requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDomain: %w", err)
	}
	return oldValue.Domain, nil
}

// ResetDomain resets all changes to the "domain" field.
func (m *ServiceAPIMutation) ResetDomain() {
	m.domain = nil
}

// SetMethod sets the "method" field.
func (m *ServiceAPIMutation) SetMethod(s string) {
	m.method = &s
}

// Method returns the value of the "method" field in the mutation.
func (m *ServiceAPIMutation) Method() (r string, exists bool) {
	v := m.method
	if v == nil {
		return
	}
	return *v, true
}

// OldMethod returns the old "method" field's value of the ServiceAPI entity.
// If the ServiceAPI object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ServiceAPIMutation) OldMethod(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldMethod is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldMethod requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMethod: %w", err)
	}
	return oldValue.Method, nil
}

// ResetMethod resets all changes to the "method" field.
func (m *ServiceAPIMutation) ResetMethod() {
	m.method = nil
}

// SetPath sets the "path" field.
func (m *ServiceAPIMutation) SetPath(s string) {
	m._path = &s
}

// Path returns the value of the "path" field in the mutation.
func (m *ServiceAPIMutation) Path() (r string, exists bool) {
	v := m._path
	if v == nil {
		return
	}
	return *v, true
}

// OldPath returns the old "path" field's value of the ServiceAPI entity.
// If the ServiceAPI object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ServiceAPIMutation) OldPath(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPath is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPath requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPath: %w", err)
	}
	return oldValue.Path, nil
}

// ResetPath resets all changes to the "path" field.
func (m *ServiceAPIMutation) ResetPath() {
	m._path = nil
}

// SetExported sets the "exported" field.
func (m *ServiceAPIMutation) SetExported(b bool) {
	m.exported = &b
}

// Exported returns the value of the "exported" field in the mutation.
func (m *ServiceAPIMutation) Exported() (r bool, exists bool) {
	v := m.exported
	if v == nil {
		return
	}
	return *v, true
}

// OldExported returns the old "exported" field's value of the ServiceAPI entity.
// If the ServiceAPI object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ServiceAPIMutation) OldExported(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldExported is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldExported requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldExported: %w", err)
	}
	return oldValue.Exported, nil
}

// ResetExported resets all changes to the "exported" field.
func (m *ServiceAPIMutation) ResetExported() {
	m.exported = nil
}

// SetPathPrefix sets the "path_prefix" field.
func (m *ServiceAPIMutation) SetPathPrefix(s string) {
	m.path_prefix = &s
}

// PathPrefix returns the value of the "path_prefix" field in the mutation.
func (m *ServiceAPIMutation) PathPrefix() (r string, exists bool) {
	v := m.path_prefix
	if v == nil {
		return
	}
	return *v, true
}

// OldPathPrefix returns the old "path_prefix" field's value of the ServiceAPI entity.
// If the ServiceAPI object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ServiceAPIMutation) OldPathPrefix(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPathPrefix is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPathPrefix requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPathPrefix: %w", err)
	}
	return oldValue.PathPrefix, nil
}

// ResetPathPrefix resets all changes to the "path_prefix" field.
func (m *ServiceAPIMutation) ResetPathPrefix() {
	m.path_prefix = nil
}

// SetCreateAt sets the "create_at" field.
func (m *ServiceAPIMutation) SetCreateAt(u uint32) {
	m.create_at = &u
	m.addcreate_at = nil
}

// CreateAt returns the value of the "create_at" field in the mutation.
func (m *ServiceAPIMutation) CreateAt() (r uint32, exists bool) {
	v := m.create_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateAt returns the old "create_at" field's value of the ServiceAPI entity.
// If the ServiceAPI object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ServiceAPIMutation) OldCreateAt(ctx context.Context) (v uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreateAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreateAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateAt: %w", err)
	}
	return oldValue.CreateAt, nil
}

// AddCreateAt adds u to the "create_at" field.
func (m *ServiceAPIMutation) AddCreateAt(u uint32) {
	if m.addcreate_at != nil {
		*m.addcreate_at += u
	} else {
		m.addcreate_at = &u
	}
}

// AddedCreateAt returns the value that was added to the "create_at" field in this mutation.
func (m *ServiceAPIMutation) AddedCreateAt() (r uint32, exists bool) {
	v := m.addcreate_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetCreateAt resets all changes to the "create_at" field.
func (m *ServiceAPIMutation) ResetCreateAt() {
	m.create_at = nil
	m.addcreate_at = nil
}

// SetUpdateAt sets the "update_at" field.
func (m *ServiceAPIMutation) SetUpdateAt(u uint32) {
	m.update_at = &u
	m.addupdate_at = nil
}

// UpdateAt returns the value of the "update_at" field in the mutation.
func (m *ServiceAPIMutation) UpdateAt() (r uint32, exists bool) {
	v := m.update_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdateAt returns the old "update_at" field's value of the ServiceAPI entity.
// If the ServiceAPI object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ServiceAPIMutation) OldUpdateAt(ctx context.Context) (v uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdateAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdateAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdateAt: %w", err)
	}
	return oldValue.UpdateAt, nil
}

// AddUpdateAt adds u to the "update_at" field.
func (m *ServiceAPIMutation) AddUpdateAt(u uint32) {
	if m.addupdate_at != nil {
		*m.addupdate_at += u
	} else {
		m.addupdate_at = &u
	}
}

// AddedUpdateAt returns the value that was added to the "update_at" field in this mutation.
func (m *ServiceAPIMutation) AddedUpdateAt() (r uint32, exists bool) {
	v := m.addupdate_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetUpdateAt resets all changes to the "update_at" field.
func (m *ServiceAPIMutation) ResetUpdateAt() {
	m.update_at = nil
	m.addupdate_at = nil
}

// SetDeleteAt sets the "delete_at" field.
func (m *ServiceAPIMutation) SetDeleteAt(u uint32) {
	m.delete_at = &u
	m.adddelete_at = nil
}

// DeleteAt returns the value of the "delete_at" field in the mutation.
func (m *ServiceAPIMutation) DeleteAt() (r uint32, exists bool) {
	v := m.delete_at
	if v == nil {
		return
	}
	return *v, true
}

// OldDeleteAt returns the old "delete_at" field's value of the ServiceAPI entity.
// If the ServiceAPI object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ServiceAPIMutation) OldDeleteAt(ctx context.Context) (v uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDeleteAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDeleteAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeleteAt: %w", err)
	}
	return oldValue.DeleteAt, nil
}

// AddDeleteAt adds u to the "delete_at" field.
func (m *ServiceAPIMutation) AddDeleteAt(u uint32) {
	if m.adddelete_at != nil {
		*m.adddelete_at += u
	} else {
		m.adddelete_at = &u
	}
}

// AddedDeleteAt returns the value that was added to the "delete_at" field in this mutation.
func (m *ServiceAPIMutation) AddedDeleteAt() (r uint32, exists bool) {
	v := m.adddelete_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetDeleteAt resets all changes to the "delete_at" field.
func (m *ServiceAPIMutation) ResetDeleteAt() {
	m.delete_at = nil
	m.adddelete_at = nil
}

// Where appends a list predicates to the ServiceAPIMutation builder.
func (m *ServiceAPIMutation) Where(ps ...predicate.ServiceAPI) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *ServiceAPIMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (ServiceAPI).
func (m *ServiceAPIMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ServiceAPIMutation) Fields() []string {
	fields := make([]string, 0, 8)
	if m.domain != nil {
		fields = append(fields, serviceapi.FieldDomain)
	}
	if m.method != nil {
		fields = append(fields, serviceapi.FieldMethod)
	}
	if m._path != nil {
		fields = append(fields, serviceapi.FieldPath)
	}
	if m.exported != nil {
		fields = append(fields, serviceapi.FieldExported)
	}
	if m.path_prefix != nil {
		fields = append(fields, serviceapi.FieldPathPrefix)
	}
	if m.create_at != nil {
		fields = append(fields, serviceapi.FieldCreateAt)
	}
	if m.update_at != nil {
		fields = append(fields, serviceapi.FieldUpdateAt)
	}
	if m.delete_at != nil {
		fields = append(fields, serviceapi.FieldDeleteAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ServiceAPIMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case serviceapi.FieldDomain:
		return m.Domain()
	case serviceapi.FieldMethod:
		return m.Method()
	case serviceapi.FieldPath:
		return m.Path()
	case serviceapi.FieldExported:
		return m.Exported()
	case serviceapi.FieldPathPrefix:
		return m.PathPrefix()
	case serviceapi.FieldCreateAt:
		return m.CreateAt()
	case serviceapi.FieldUpdateAt:
		return m.UpdateAt()
	case serviceapi.FieldDeleteAt:
		return m.DeleteAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ServiceAPIMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case serviceapi.FieldDomain:
		return m.OldDomain(ctx)
	case serviceapi.FieldMethod:
		return m.OldMethod(ctx)
	case serviceapi.FieldPath:
		return m.OldPath(ctx)
	case serviceapi.FieldExported:
		return m.OldExported(ctx)
	case serviceapi.FieldPathPrefix:
		return m.OldPathPrefix(ctx)
	case serviceapi.FieldCreateAt:
		return m.OldCreateAt(ctx)
	case serviceapi.FieldUpdateAt:
		return m.OldUpdateAt(ctx)
	case serviceapi.FieldDeleteAt:
		return m.OldDeleteAt(ctx)
	}
	return nil, fmt.Errorf("unknown ServiceAPI field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ServiceAPIMutation) SetField(name string, value ent.Value) error {
	switch name {
	case serviceapi.FieldDomain:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDomain(v)
		return nil
	case serviceapi.FieldMethod:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMethod(v)
		return nil
	case serviceapi.FieldPath:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPath(v)
		return nil
	case serviceapi.FieldExported:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetExported(v)
		return nil
	case serviceapi.FieldPathPrefix:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPathPrefix(v)
		return nil
	case serviceapi.FieldCreateAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateAt(v)
		return nil
	case serviceapi.FieldUpdateAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdateAt(v)
		return nil
	case serviceapi.FieldDeleteAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeleteAt(v)
		return nil
	}
	return fmt.Errorf("unknown ServiceAPI field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ServiceAPIMutation) AddedFields() []string {
	var fields []string
	if m.addcreate_at != nil {
		fields = append(fields, serviceapi.FieldCreateAt)
	}
	if m.addupdate_at != nil {
		fields = append(fields, serviceapi.FieldUpdateAt)
	}
	if m.adddelete_at != nil {
		fields = append(fields, serviceapi.FieldDeleteAt)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ServiceAPIMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case serviceapi.FieldCreateAt:
		return m.AddedCreateAt()
	case serviceapi.FieldUpdateAt:
		return m.AddedUpdateAt()
	case serviceapi.FieldDeleteAt:
		return m.AddedDeleteAt()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ServiceAPIMutation) AddField(name string, value ent.Value) error {
	switch name {
	case serviceapi.FieldCreateAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddCreateAt(v)
		return nil
	case serviceapi.FieldUpdateAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddUpdateAt(v)
		return nil
	case serviceapi.FieldDeleteAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddDeleteAt(v)
		return nil
	}
	return fmt.Errorf("unknown ServiceAPI numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ServiceAPIMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ServiceAPIMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ServiceAPIMutation) ClearField(name string) error {
	return fmt.Errorf("unknown ServiceAPI nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ServiceAPIMutation) ResetField(name string) error {
	switch name {
	case serviceapi.FieldDomain:
		m.ResetDomain()
		return nil
	case serviceapi.FieldMethod:
		m.ResetMethod()
		return nil
	case serviceapi.FieldPath:
		m.ResetPath()
		return nil
	case serviceapi.FieldExported:
		m.ResetExported()
		return nil
	case serviceapi.FieldPathPrefix:
		m.ResetPathPrefix()
		return nil
	case serviceapi.FieldCreateAt:
		m.ResetCreateAt()
		return nil
	case serviceapi.FieldUpdateAt:
		m.ResetUpdateAt()
		return nil
	case serviceapi.FieldDeleteAt:
		m.ResetDeleteAt()
		return nil
	}
	return fmt.Errorf("unknown ServiceAPI field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ServiceAPIMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ServiceAPIMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ServiceAPIMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ServiceAPIMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ServiceAPIMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ServiceAPIMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ServiceAPIMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown ServiceAPI unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ServiceAPIMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown ServiceAPI edge %s", name)
}
