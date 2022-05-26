// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/api-manager/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/api-manager/pkg/db/ent/serviceapi"
)

// ServiceAPIUpdate is the builder for updating ServiceAPI entities.
type ServiceAPIUpdate struct {
	config
	hooks    []Hook
	mutation *ServiceAPIMutation
}

// Where appends a list predicates to the ServiceAPIUpdate builder.
func (sau *ServiceAPIUpdate) Where(ps ...predicate.ServiceAPI) *ServiceAPIUpdate {
	sau.mutation.Where(ps...)
	return sau
}

// SetProtocol sets the "protocol" field.
func (sau *ServiceAPIUpdate) SetProtocol(s string) *ServiceAPIUpdate {
	sau.mutation.SetProtocol(s)
	return sau
}

// SetServiceName sets the "service_name" field.
func (sau *ServiceAPIUpdate) SetServiceName(s string) *ServiceAPIUpdate {
	sau.mutation.SetServiceName(s)
	return sau
}

// SetDomains sets the "domains" field.
func (sau *ServiceAPIUpdate) SetDomains(s []string) *ServiceAPIUpdate {
	sau.mutation.SetDomains(s)
	return sau
}

// SetMethod sets the "method" field.
func (sau *ServiceAPIUpdate) SetMethod(s string) *ServiceAPIUpdate {
	sau.mutation.SetMethod(s)
	return sau
}

// SetMethodName sets the "method_name" field.
func (sau *ServiceAPIUpdate) SetMethodName(s string) *ServiceAPIUpdate {
	sau.mutation.SetMethodName(s)
	return sau
}

// SetPath sets the "path" field.
func (sau *ServiceAPIUpdate) SetPath(s string) *ServiceAPIUpdate {
	sau.mutation.SetPath(s)
	return sau
}

// SetExported sets the "exported" field.
func (sau *ServiceAPIUpdate) SetExported(b bool) *ServiceAPIUpdate {
	sau.mutation.SetExported(b)
	return sau
}

// SetPathPrefix sets the "path_prefix" field.
func (sau *ServiceAPIUpdate) SetPathPrefix(s string) *ServiceAPIUpdate {
	sau.mutation.SetPathPrefix(s)
	return sau
}

// SetCreateAt sets the "create_at" field.
func (sau *ServiceAPIUpdate) SetCreateAt(u uint32) *ServiceAPIUpdate {
	sau.mutation.ResetCreateAt()
	sau.mutation.SetCreateAt(u)
	return sau
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (sau *ServiceAPIUpdate) SetNillableCreateAt(u *uint32) *ServiceAPIUpdate {
	if u != nil {
		sau.SetCreateAt(*u)
	}
	return sau
}

// AddCreateAt adds u to the "create_at" field.
func (sau *ServiceAPIUpdate) AddCreateAt(u int32) *ServiceAPIUpdate {
	sau.mutation.AddCreateAt(u)
	return sau
}

// SetUpdateAt sets the "update_at" field.
func (sau *ServiceAPIUpdate) SetUpdateAt(u uint32) *ServiceAPIUpdate {
	sau.mutation.ResetUpdateAt()
	sau.mutation.SetUpdateAt(u)
	return sau
}

// AddUpdateAt adds u to the "update_at" field.
func (sau *ServiceAPIUpdate) AddUpdateAt(u int32) *ServiceAPIUpdate {
	sau.mutation.AddUpdateAt(u)
	return sau
}

// SetDeleteAt sets the "delete_at" field.
func (sau *ServiceAPIUpdate) SetDeleteAt(u uint32) *ServiceAPIUpdate {
	sau.mutation.ResetDeleteAt()
	sau.mutation.SetDeleteAt(u)
	return sau
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (sau *ServiceAPIUpdate) SetNillableDeleteAt(u *uint32) *ServiceAPIUpdate {
	if u != nil {
		sau.SetDeleteAt(*u)
	}
	return sau
}

// AddDeleteAt adds u to the "delete_at" field.
func (sau *ServiceAPIUpdate) AddDeleteAt(u int32) *ServiceAPIUpdate {
	sau.mutation.AddDeleteAt(u)
	return sau
}

// Mutation returns the ServiceAPIMutation object of the builder.
func (sau *ServiceAPIUpdate) Mutation() *ServiceAPIMutation {
	return sau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sau *ServiceAPIUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	sau.defaults()
	if len(sau.hooks) == 0 {
		affected, err = sau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServiceAPIMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sau.mutation = mutation
			affected, err = sau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sau.hooks) - 1; i >= 0; i-- {
			if sau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sau *ServiceAPIUpdate) SaveX(ctx context.Context) int {
	affected, err := sau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sau *ServiceAPIUpdate) Exec(ctx context.Context) error {
	_, err := sau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sau *ServiceAPIUpdate) ExecX(ctx context.Context) {
	if err := sau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sau *ServiceAPIUpdate) defaults() {
	if _, ok := sau.mutation.UpdateAt(); !ok {
		v := serviceapi.UpdateDefaultUpdateAt()
		sau.mutation.SetUpdateAt(v)
	}
}

func (sau *ServiceAPIUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   serviceapi.Table,
			Columns: serviceapi.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: serviceapi.FieldID,
			},
		},
	}
	if ps := sau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sau.mutation.Protocol(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serviceapi.FieldProtocol,
		})
	}
	if value, ok := sau.mutation.ServiceName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serviceapi.FieldServiceName,
		})
	}
	if value, ok := sau.mutation.Domains(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: serviceapi.FieldDomains,
		})
	}
	if value, ok := sau.mutation.Method(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serviceapi.FieldMethod,
		})
	}
	if value, ok := sau.mutation.MethodName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serviceapi.FieldMethodName,
		})
	}
	if value, ok := sau.mutation.Path(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serviceapi.FieldPath,
		})
	}
	if value, ok := sau.mutation.Exported(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: serviceapi.FieldExported,
		})
	}
	if value, ok := sau.mutation.PathPrefix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serviceapi.FieldPathPrefix,
		})
	}
	if value, ok := sau.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: serviceapi.FieldCreateAt,
		})
	}
	if value, ok := sau.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: serviceapi.FieldCreateAt,
		})
	}
	if value, ok := sau.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: serviceapi.FieldUpdateAt,
		})
	}
	if value, ok := sau.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: serviceapi.FieldUpdateAt,
		})
	}
	if value, ok := sau.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: serviceapi.FieldDeleteAt,
		})
	}
	if value, ok := sau.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: serviceapi.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{serviceapi.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ServiceAPIUpdateOne is the builder for updating a single ServiceAPI entity.
type ServiceAPIUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ServiceAPIMutation
}

// SetProtocol sets the "protocol" field.
func (sauo *ServiceAPIUpdateOne) SetProtocol(s string) *ServiceAPIUpdateOne {
	sauo.mutation.SetProtocol(s)
	return sauo
}

// SetServiceName sets the "service_name" field.
func (sauo *ServiceAPIUpdateOne) SetServiceName(s string) *ServiceAPIUpdateOne {
	sauo.mutation.SetServiceName(s)
	return sauo
}

// SetDomains sets the "domains" field.
func (sauo *ServiceAPIUpdateOne) SetDomains(s []string) *ServiceAPIUpdateOne {
	sauo.mutation.SetDomains(s)
	return sauo
}

// SetMethod sets the "method" field.
func (sauo *ServiceAPIUpdateOne) SetMethod(s string) *ServiceAPIUpdateOne {
	sauo.mutation.SetMethod(s)
	return sauo
}

// SetMethodName sets the "method_name" field.
func (sauo *ServiceAPIUpdateOne) SetMethodName(s string) *ServiceAPIUpdateOne {
	sauo.mutation.SetMethodName(s)
	return sauo
}

// SetPath sets the "path" field.
func (sauo *ServiceAPIUpdateOne) SetPath(s string) *ServiceAPIUpdateOne {
	sauo.mutation.SetPath(s)
	return sauo
}

// SetExported sets the "exported" field.
func (sauo *ServiceAPIUpdateOne) SetExported(b bool) *ServiceAPIUpdateOne {
	sauo.mutation.SetExported(b)
	return sauo
}

// SetPathPrefix sets the "path_prefix" field.
func (sauo *ServiceAPIUpdateOne) SetPathPrefix(s string) *ServiceAPIUpdateOne {
	sauo.mutation.SetPathPrefix(s)
	return sauo
}

// SetCreateAt sets the "create_at" field.
func (sauo *ServiceAPIUpdateOne) SetCreateAt(u uint32) *ServiceAPIUpdateOne {
	sauo.mutation.ResetCreateAt()
	sauo.mutation.SetCreateAt(u)
	return sauo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (sauo *ServiceAPIUpdateOne) SetNillableCreateAt(u *uint32) *ServiceAPIUpdateOne {
	if u != nil {
		sauo.SetCreateAt(*u)
	}
	return sauo
}

// AddCreateAt adds u to the "create_at" field.
func (sauo *ServiceAPIUpdateOne) AddCreateAt(u int32) *ServiceAPIUpdateOne {
	sauo.mutation.AddCreateAt(u)
	return sauo
}

// SetUpdateAt sets the "update_at" field.
func (sauo *ServiceAPIUpdateOne) SetUpdateAt(u uint32) *ServiceAPIUpdateOne {
	sauo.mutation.ResetUpdateAt()
	sauo.mutation.SetUpdateAt(u)
	return sauo
}

// AddUpdateAt adds u to the "update_at" field.
func (sauo *ServiceAPIUpdateOne) AddUpdateAt(u int32) *ServiceAPIUpdateOne {
	sauo.mutation.AddUpdateAt(u)
	return sauo
}

// SetDeleteAt sets the "delete_at" field.
func (sauo *ServiceAPIUpdateOne) SetDeleteAt(u uint32) *ServiceAPIUpdateOne {
	sauo.mutation.ResetDeleteAt()
	sauo.mutation.SetDeleteAt(u)
	return sauo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (sauo *ServiceAPIUpdateOne) SetNillableDeleteAt(u *uint32) *ServiceAPIUpdateOne {
	if u != nil {
		sauo.SetDeleteAt(*u)
	}
	return sauo
}

// AddDeleteAt adds u to the "delete_at" field.
func (sauo *ServiceAPIUpdateOne) AddDeleteAt(u int32) *ServiceAPIUpdateOne {
	sauo.mutation.AddDeleteAt(u)
	return sauo
}

// Mutation returns the ServiceAPIMutation object of the builder.
func (sauo *ServiceAPIUpdateOne) Mutation() *ServiceAPIMutation {
	return sauo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sauo *ServiceAPIUpdateOne) Select(field string, fields ...string) *ServiceAPIUpdateOne {
	sauo.fields = append([]string{field}, fields...)
	return sauo
}

// Save executes the query and returns the updated ServiceAPI entity.
func (sauo *ServiceAPIUpdateOne) Save(ctx context.Context) (*ServiceAPI, error) {
	var (
		err  error
		node *ServiceAPI
	)
	sauo.defaults()
	if len(sauo.hooks) == 0 {
		node, err = sauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServiceAPIMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sauo.mutation = mutation
			node, err = sauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sauo.hooks) - 1; i >= 0; i-- {
			if sauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sauo *ServiceAPIUpdateOne) SaveX(ctx context.Context) *ServiceAPI {
	node, err := sauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sauo *ServiceAPIUpdateOne) Exec(ctx context.Context) error {
	_, err := sauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sauo *ServiceAPIUpdateOne) ExecX(ctx context.Context) {
	if err := sauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sauo *ServiceAPIUpdateOne) defaults() {
	if _, ok := sauo.mutation.UpdateAt(); !ok {
		v := serviceapi.UpdateDefaultUpdateAt()
		sauo.mutation.SetUpdateAt(v)
	}
}

func (sauo *ServiceAPIUpdateOne) sqlSave(ctx context.Context) (_node *ServiceAPI, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   serviceapi.Table,
			Columns: serviceapi.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: serviceapi.FieldID,
			},
		},
	}
	id, ok := sauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ServiceAPI.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, serviceapi.FieldID)
		for _, f := range fields {
			if !serviceapi.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != serviceapi.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sauo.mutation.Protocol(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serviceapi.FieldProtocol,
		})
	}
	if value, ok := sauo.mutation.ServiceName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serviceapi.FieldServiceName,
		})
	}
	if value, ok := sauo.mutation.Domains(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: serviceapi.FieldDomains,
		})
	}
	if value, ok := sauo.mutation.Method(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serviceapi.FieldMethod,
		})
	}
	if value, ok := sauo.mutation.MethodName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serviceapi.FieldMethodName,
		})
	}
	if value, ok := sauo.mutation.Path(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serviceapi.FieldPath,
		})
	}
	if value, ok := sauo.mutation.Exported(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: serviceapi.FieldExported,
		})
	}
	if value, ok := sauo.mutation.PathPrefix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serviceapi.FieldPathPrefix,
		})
	}
	if value, ok := sauo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: serviceapi.FieldCreateAt,
		})
	}
	if value, ok := sauo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: serviceapi.FieldCreateAt,
		})
	}
	if value, ok := sauo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: serviceapi.FieldUpdateAt,
		})
	}
	if value, ok := sauo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: serviceapi.FieldUpdateAt,
		})
	}
	if value, ok := sauo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: serviceapi.FieldDeleteAt,
		})
	}
	if value, ok := sauo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: serviceapi.FieldDeleteAt,
		})
	}
	_node = &ServiceAPI{config: sauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{serviceapi.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
