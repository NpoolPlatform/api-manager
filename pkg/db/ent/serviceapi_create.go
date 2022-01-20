// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/api-manager/pkg/db/ent/serviceapi"
	"github.com/google/uuid"
)

// ServiceAPICreate is the builder for creating a ServiceAPI entity.
type ServiceAPICreate struct {
	config
	mutation *ServiceAPIMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetDomain sets the "domain" field.
func (sac *ServiceAPICreate) SetDomain(s string) *ServiceAPICreate {
	sac.mutation.SetDomain(s)
	return sac
}

// SetMethod sets the "method" field.
func (sac *ServiceAPICreate) SetMethod(s string) *ServiceAPICreate {
	sac.mutation.SetMethod(s)
	return sac
}

// SetPath sets the "path" field.
func (sac *ServiceAPICreate) SetPath(s string) *ServiceAPICreate {
	sac.mutation.SetPath(s)
	return sac
}

// SetExported sets the "exported" field.
func (sac *ServiceAPICreate) SetExported(b bool) *ServiceAPICreate {
	sac.mutation.SetExported(b)
	return sac
}

// SetPathPrefix sets the "path_prefix" field.
func (sac *ServiceAPICreate) SetPathPrefix(s string) *ServiceAPICreate {
	sac.mutation.SetPathPrefix(s)
	return sac
}

// SetCreateAt sets the "create_at" field.
func (sac *ServiceAPICreate) SetCreateAt(u uint32) *ServiceAPICreate {
	sac.mutation.SetCreateAt(u)
	return sac
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (sac *ServiceAPICreate) SetNillableCreateAt(u *uint32) *ServiceAPICreate {
	if u != nil {
		sac.SetCreateAt(*u)
	}
	return sac
}

// SetUpdateAt sets the "update_at" field.
func (sac *ServiceAPICreate) SetUpdateAt(u uint32) *ServiceAPICreate {
	sac.mutation.SetUpdateAt(u)
	return sac
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (sac *ServiceAPICreate) SetNillableUpdateAt(u *uint32) *ServiceAPICreate {
	if u != nil {
		sac.SetUpdateAt(*u)
	}
	return sac
}

// SetDeleteAt sets the "delete_at" field.
func (sac *ServiceAPICreate) SetDeleteAt(u uint32) *ServiceAPICreate {
	sac.mutation.SetDeleteAt(u)
	return sac
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (sac *ServiceAPICreate) SetNillableDeleteAt(u *uint32) *ServiceAPICreate {
	if u != nil {
		sac.SetDeleteAt(*u)
	}
	return sac
}

// SetID sets the "id" field.
func (sac *ServiceAPICreate) SetID(u uuid.UUID) *ServiceAPICreate {
	sac.mutation.SetID(u)
	return sac
}

// Mutation returns the ServiceAPIMutation object of the builder.
func (sac *ServiceAPICreate) Mutation() *ServiceAPIMutation {
	return sac.mutation
}

// Save creates the ServiceAPI in the database.
func (sac *ServiceAPICreate) Save(ctx context.Context) (*ServiceAPI, error) {
	var (
		err  error
		node *ServiceAPI
	)
	sac.defaults()
	if len(sac.hooks) == 0 {
		if err = sac.check(); err != nil {
			return nil, err
		}
		node, err = sac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServiceAPIMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sac.check(); err != nil {
				return nil, err
			}
			sac.mutation = mutation
			if node, err = sac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sac.hooks) - 1; i >= 0; i-- {
			if sac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sac *ServiceAPICreate) SaveX(ctx context.Context) *ServiceAPI {
	v, err := sac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sac *ServiceAPICreate) Exec(ctx context.Context) error {
	_, err := sac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sac *ServiceAPICreate) ExecX(ctx context.Context) {
	if err := sac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sac *ServiceAPICreate) defaults() {
	if _, ok := sac.mutation.CreateAt(); !ok {
		v := serviceapi.DefaultCreateAt()
		sac.mutation.SetCreateAt(v)
	}
	if _, ok := sac.mutation.UpdateAt(); !ok {
		v := serviceapi.DefaultUpdateAt()
		sac.mutation.SetUpdateAt(v)
	}
	if _, ok := sac.mutation.DeleteAt(); !ok {
		v := serviceapi.DefaultDeleteAt()
		sac.mutation.SetDeleteAt(v)
	}
	if _, ok := sac.mutation.ID(); !ok {
		v := serviceapi.DefaultID()
		sac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sac *ServiceAPICreate) check() error {
	if _, ok := sac.mutation.Domain(); !ok {
		return &ValidationError{Name: "domain", err: errors.New(`ent: missing required field "domain"`)}
	}
	if _, ok := sac.mutation.Method(); !ok {
		return &ValidationError{Name: "method", err: errors.New(`ent: missing required field "method"`)}
	}
	if _, ok := sac.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "path"`)}
	}
	if _, ok := sac.mutation.Exported(); !ok {
		return &ValidationError{Name: "exported", err: errors.New(`ent: missing required field "exported"`)}
	}
	if _, ok := sac.mutation.PathPrefix(); !ok {
		return &ValidationError{Name: "path_prefix", err: errors.New(`ent: missing required field "path_prefix"`)}
	}
	if _, ok := sac.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := sac.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := sac.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (sac *ServiceAPICreate) sqlSave(ctx context.Context) (*ServiceAPI, error) {
	_node, _spec := sac.createSpec()
	if err := sqlgraph.CreateNode(ctx, sac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (sac *ServiceAPICreate) createSpec() (*ServiceAPI, *sqlgraph.CreateSpec) {
	var (
		_node = &ServiceAPI{config: sac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: serviceapi.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: serviceapi.FieldID,
			},
		}
	)
	_spec.OnConflict = sac.conflict
	if id, ok := sac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sac.mutation.Domain(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serviceapi.FieldDomain,
		})
		_node.Domain = value
	}
	if value, ok := sac.mutation.Method(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serviceapi.FieldMethod,
		})
		_node.Method = value
	}
	if value, ok := sac.mutation.Path(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serviceapi.FieldPath,
		})
		_node.Path = value
	}
	if value, ok := sac.mutation.Exported(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: serviceapi.FieldExported,
		})
		_node.Exported = value
	}
	if value, ok := sac.mutation.PathPrefix(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serviceapi.FieldPathPrefix,
		})
		_node.PathPrefix = value
	}
	if value, ok := sac.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: serviceapi.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := sac.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: serviceapi.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := sac.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: serviceapi.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ServiceAPI.Create().
//		SetDomain(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ServiceAPIUpsert) {
//			SetDomain(v+v).
//		}).
//		Exec(ctx)
//
func (sac *ServiceAPICreate) OnConflict(opts ...sql.ConflictOption) *ServiceAPIUpsertOne {
	sac.conflict = opts
	return &ServiceAPIUpsertOne{
		create: sac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ServiceAPI.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (sac *ServiceAPICreate) OnConflictColumns(columns ...string) *ServiceAPIUpsertOne {
	sac.conflict = append(sac.conflict, sql.ConflictColumns(columns...))
	return &ServiceAPIUpsertOne{
		create: sac,
	}
}

type (
	// ServiceAPIUpsertOne is the builder for "upsert"-ing
	//  one ServiceAPI node.
	ServiceAPIUpsertOne struct {
		create *ServiceAPICreate
	}

	// ServiceAPIUpsert is the "OnConflict" setter.
	ServiceAPIUpsert struct {
		*sql.UpdateSet
	}
)

// SetDomain sets the "domain" field.
func (u *ServiceAPIUpsert) SetDomain(v string) *ServiceAPIUpsert {
	u.Set(serviceapi.FieldDomain, v)
	return u
}

// UpdateDomain sets the "domain" field to the value that was provided on create.
func (u *ServiceAPIUpsert) UpdateDomain() *ServiceAPIUpsert {
	u.SetExcluded(serviceapi.FieldDomain)
	return u
}

// SetMethod sets the "method" field.
func (u *ServiceAPIUpsert) SetMethod(v string) *ServiceAPIUpsert {
	u.Set(serviceapi.FieldMethod, v)
	return u
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *ServiceAPIUpsert) UpdateMethod() *ServiceAPIUpsert {
	u.SetExcluded(serviceapi.FieldMethod)
	return u
}

// SetPath sets the "path" field.
func (u *ServiceAPIUpsert) SetPath(v string) *ServiceAPIUpsert {
	u.Set(serviceapi.FieldPath, v)
	return u
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *ServiceAPIUpsert) UpdatePath() *ServiceAPIUpsert {
	u.SetExcluded(serviceapi.FieldPath)
	return u
}

// SetExported sets the "exported" field.
func (u *ServiceAPIUpsert) SetExported(v bool) *ServiceAPIUpsert {
	u.Set(serviceapi.FieldExported, v)
	return u
}

// UpdateExported sets the "exported" field to the value that was provided on create.
func (u *ServiceAPIUpsert) UpdateExported() *ServiceAPIUpsert {
	u.SetExcluded(serviceapi.FieldExported)
	return u
}

// SetPathPrefix sets the "path_prefix" field.
func (u *ServiceAPIUpsert) SetPathPrefix(v string) *ServiceAPIUpsert {
	u.Set(serviceapi.FieldPathPrefix, v)
	return u
}

// UpdatePathPrefix sets the "path_prefix" field to the value that was provided on create.
func (u *ServiceAPIUpsert) UpdatePathPrefix() *ServiceAPIUpsert {
	u.SetExcluded(serviceapi.FieldPathPrefix)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *ServiceAPIUpsert) SetCreateAt(v uint32) *ServiceAPIUpsert {
	u.Set(serviceapi.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *ServiceAPIUpsert) UpdateCreateAt() *ServiceAPIUpsert {
	u.SetExcluded(serviceapi.FieldCreateAt)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *ServiceAPIUpsert) SetUpdateAt(v uint32) *ServiceAPIUpsert {
	u.Set(serviceapi.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *ServiceAPIUpsert) UpdateUpdateAt() *ServiceAPIUpsert {
	u.SetExcluded(serviceapi.FieldUpdateAt)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *ServiceAPIUpsert) SetDeleteAt(v uint32) *ServiceAPIUpsert {
	u.Set(serviceapi.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *ServiceAPIUpsert) UpdateDeleteAt() *ServiceAPIUpsert {
	u.SetExcluded(serviceapi.FieldDeleteAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ServiceAPI.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(serviceapi.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ServiceAPIUpsertOne) UpdateNewValues() *ServiceAPIUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(serviceapi.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.ServiceAPI.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *ServiceAPIUpsertOne) Ignore() *ServiceAPIUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ServiceAPIUpsertOne) DoNothing() *ServiceAPIUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ServiceAPICreate.OnConflict
// documentation for more info.
func (u *ServiceAPIUpsertOne) Update(set func(*ServiceAPIUpsert)) *ServiceAPIUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ServiceAPIUpsert{UpdateSet: update})
	}))
	return u
}

// SetDomain sets the "domain" field.
func (u *ServiceAPIUpsertOne) SetDomain(v string) *ServiceAPIUpsertOne {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.SetDomain(v)
	})
}

// UpdateDomain sets the "domain" field to the value that was provided on create.
func (u *ServiceAPIUpsertOne) UpdateDomain() *ServiceAPIUpsertOne {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.UpdateDomain()
	})
}

// SetMethod sets the "method" field.
func (u *ServiceAPIUpsertOne) SetMethod(v string) *ServiceAPIUpsertOne {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.SetMethod(v)
	})
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *ServiceAPIUpsertOne) UpdateMethod() *ServiceAPIUpsertOne {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.UpdateMethod()
	})
}

// SetPath sets the "path" field.
func (u *ServiceAPIUpsertOne) SetPath(v string) *ServiceAPIUpsertOne {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.SetPath(v)
	})
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *ServiceAPIUpsertOne) UpdatePath() *ServiceAPIUpsertOne {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.UpdatePath()
	})
}

// SetExported sets the "exported" field.
func (u *ServiceAPIUpsertOne) SetExported(v bool) *ServiceAPIUpsertOne {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.SetExported(v)
	})
}

// UpdateExported sets the "exported" field to the value that was provided on create.
func (u *ServiceAPIUpsertOne) UpdateExported() *ServiceAPIUpsertOne {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.UpdateExported()
	})
}

// SetPathPrefix sets the "path_prefix" field.
func (u *ServiceAPIUpsertOne) SetPathPrefix(v string) *ServiceAPIUpsertOne {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.SetPathPrefix(v)
	})
}

// UpdatePathPrefix sets the "path_prefix" field to the value that was provided on create.
func (u *ServiceAPIUpsertOne) UpdatePathPrefix() *ServiceAPIUpsertOne {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.UpdatePathPrefix()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *ServiceAPIUpsertOne) SetCreateAt(v uint32) *ServiceAPIUpsertOne {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *ServiceAPIUpsertOne) UpdateCreateAt() *ServiceAPIUpsertOne {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *ServiceAPIUpsertOne) SetUpdateAt(v uint32) *ServiceAPIUpsertOne {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *ServiceAPIUpsertOne) UpdateUpdateAt() *ServiceAPIUpsertOne {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *ServiceAPIUpsertOne) SetDeleteAt(v uint32) *ServiceAPIUpsertOne {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *ServiceAPIUpsertOne) UpdateDeleteAt() *ServiceAPIUpsertOne {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *ServiceAPIUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ServiceAPICreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ServiceAPIUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ServiceAPIUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ServiceAPIUpsertOne.ID is not supported by MySQL driver. Use ServiceAPIUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ServiceAPIUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ServiceAPICreateBulk is the builder for creating many ServiceAPI entities in bulk.
type ServiceAPICreateBulk struct {
	config
	builders []*ServiceAPICreate
	conflict []sql.ConflictOption
}

// Save creates the ServiceAPI entities in the database.
func (sacb *ServiceAPICreateBulk) Save(ctx context.Context) ([]*ServiceAPI, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sacb.builders))
	nodes := make([]*ServiceAPI, len(sacb.builders))
	mutators := make([]Mutator, len(sacb.builders))
	for i := range sacb.builders {
		func(i int, root context.Context) {
			builder := sacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ServiceAPIMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = sacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, sacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sacb *ServiceAPICreateBulk) SaveX(ctx context.Context) []*ServiceAPI {
	v, err := sacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sacb *ServiceAPICreateBulk) Exec(ctx context.Context) error {
	_, err := sacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sacb *ServiceAPICreateBulk) ExecX(ctx context.Context) {
	if err := sacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ServiceAPI.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ServiceAPIUpsert) {
//			SetDomain(v+v).
//		}).
//		Exec(ctx)
//
func (sacb *ServiceAPICreateBulk) OnConflict(opts ...sql.ConflictOption) *ServiceAPIUpsertBulk {
	sacb.conflict = opts
	return &ServiceAPIUpsertBulk{
		create: sacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ServiceAPI.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (sacb *ServiceAPICreateBulk) OnConflictColumns(columns ...string) *ServiceAPIUpsertBulk {
	sacb.conflict = append(sacb.conflict, sql.ConflictColumns(columns...))
	return &ServiceAPIUpsertBulk{
		create: sacb,
	}
}

// ServiceAPIUpsertBulk is the builder for "upsert"-ing
// a bulk of ServiceAPI nodes.
type ServiceAPIUpsertBulk struct {
	create *ServiceAPICreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ServiceAPI.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(serviceapi.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ServiceAPIUpsertBulk) UpdateNewValues() *ServiceAPIUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(serviceapi.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ServiceAPI.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *ServiceAPIUpsertBulk) Ignore() *ServiceAPIUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ServiceAPIUpsertBulk) DoNothing() *ServiceAPIUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ServiceAPICreateBulk.OnConflict
// documentation for more info.
func (u *ServiceAPIUpsertBulk) Update(set func(*ServiceAPIUpsert)) *ServiceAPIUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ServiceAPIUpsert{UpdateSet: update})
	}))
	return u
}

// SetDomain sets the "domain" field.
func (u *ServiceAPIUpsertBulk) SetDomain(v string) *ServiceAPIUpsertBulk {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.SetDomain(v)
	})
}

// UpdateDomain sets the "domain" field to the value that was provided on create.
func (u *ServiceAPIUpsertBulk) UpdateDomain() *ServiceAPIUpsertBulk {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.UpdateDomain()
	})
}

// SetMethod sets the "method" field.
func (u *ServiceAPIUpsertBulk) SetMethod(v string) *ServiceAPIUpsertBulk {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.SetMethod(v)
	})
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *ServiceAPIUpsertBulk) UpdateMethod() *ServiceAPIUpsertBulk {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.UpdateMethod()
	})
}

// SetPath sets the "path" field.
func (u *ServiceAPIUpsertBulk) SetPath(v string) *ServiceAPIUpsertBulk {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.SetPath(v)
	})
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *ServiceAPIUpsertBulk) UpdatePath() *ServiceAPIUpsertBulk {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.UpdatePath()
	})
}

// SetExported sets the "exported" field.
func (u *ServiceAPIUpsertBulk) SetExported(v bool) *ServiceAPIUpsertBulk {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.SetExported(v)
	})
}

// UpdateExported sets the "exported" field to the value that was provided on create.
func (u *ServiceAPIUpsertBulk) UpdateExported() *ServiceAPIUpsertBulk {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.UpdateExported()
	})
}

// SetPathPrefix sets the "path_prefix" field.
func (u *ServiceAPIUpsertBulk) SetPathPrefix(v string) *ServiceAPIUpsertBulk {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.SetPathPrefix(v)
	})
}

// UpdatePathPrefix sets the "path_prefix" field to the value that was provided on create.
func (u *ServiceAPIUpsertBulk) UpdatePathPrefix() *ServiceAPIUpsertBulk {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.UpdatePathPrefix()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *ServiceAPIUpsertBulk) SetCreateAt(v uint32) *ServiceAPIUpsertBulk {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *ServiceAPIUpsertBulk) UpdateCreateAt() *ServiceAPIUpsertBulk {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *ServiceAPIUpsertBulk) SetUpdateAt(v uint32) *ServiceAPIUpsertBulk {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *ServiceAPIUpsertBulk) UpdateUpdateAt() *ServiceAPIUpsertBulk {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *ServiceAPIUpsertBulk) SetDeleteAt(v uint32) *ServiceAPIUpsertBulk {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *ServiceAPIUpsertBulk) UpdateDeleteAt() *ServiceAPIUpsertBulk {
	return u.Update(func(s *ServiceAPIUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *ServiceAPIUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ServiceAPICreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ServiceAPICreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ServiceAPIUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
