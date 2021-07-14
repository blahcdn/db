// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/blahcdn/db/ent/predicate"
	"github.com/blahcdn/db/ent/user"
	"github.com/blahcdn/db/ent/zone"
)

// ZoneUpdate is the builder for updating Zone entities.
type ZoneUpdate struct {
	config
	hooks    []Hook
	mutation *ZoneMutation
}

// Where adds a new predicate for the ZoneUpdate builder.
func (zu *ZoneUpdate) Where(ps ...predicate.Zone) *ZoneUpdate {
	zu.mutation.predicates = append(zu.mutation.predicates, ps...)
	return zu
}

// SetDomain sets the "domain" field.
func (zu *ZoneUpdate) SetDomain(s string) *ZoneUpdate {
	zu.mutation.SetDomain(s)
	return zu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (zu *ZoneUpdate) SetOwnerID(id int) *ZoneUpdate {
	zu.mutation.SetOwnerID(id)
	return zu
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (zu *ZoneUpdate) SetNillableOwnerID(id *int) *ZoneUpdate {
	if id != nil {
		zu = zu.SetOwnerID(*id)
	}
	return zu
}

// SetOwner sets the "owner" edge to the User entity.
func (zu *ZoneUpdate) SetOwner(u *User) *ZoneUpdate {
	return zu.SetOwnerID(u.ID)
}

// Mutation returns the ZoneMutation object of the builder.
func (zu *ZoneUpdate) Mutation() *ZoneMutation {
	return zu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (zu *ZoneUpdate) ClearOwner() *ZoneUpdate {
	zu.mutation.ClearOwner()
	return zu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (zu *ZoneUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(zu.hooks) == 0 {
		affected, err = zu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ZoneMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			zu.mutation = mutation
			affected, err = zu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(zu.hooks) - 1; i >= 0; i-- {
			mut = zu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, zu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (zu *ZoneUpdate) SaveX(ctx context.Context) int {
	affected, err := zu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (zu *ZoneUpdate) Exec(ctx context.Context) error {
	_, err := zu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (zu *ZoneUpdate) ExecX(ctx context.Context) {
	if err := zu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (zu *ZoneUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   zone.Table,
			Columns: zone.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: zone.FieldID,
			},
		},
	}
	if ps := zu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := zu.mutation.Domain(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: zone.FieldDomain,
		})
	}
	if zu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   zone.OwnerTable,
			Columns: []string{zone.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := zu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   zone.OwnerTable,
			Columns: []string{zone.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, zu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{zone.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ZoneUpdateOne is the builder for updating a single Zone entity.
type ZoneUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ZoneMutation
}

// SetDomain sets the "domain" field.
func (zuo *ZoneUpdateOne) SetDomain(s string) *ZoneUpdateOne {
	zuo.mutation.SetDomain(s)
	return zuo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (zuo *ZoneUpdateOne) SetOwnerID(id int) *ZoneUpdateOne {
	zuo.mutation.SetOwnerID(id)
	return zuo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (zuo *ZoneUpdateOne) SetNillableOwnerID(id *int) *ZoneUpdateOne {
	if id != nil {
		zuo = zuo.SetOwnerID(*id)
	}
	return zuo
}

// SetOwner sets the "owner" edge to the User entity.
func (zuo *ZoneUpdateOne) SetOwner(u *User) *ZoneUpdateOne {
	return zuo.SetOwnerID(u.ID)
}

// Mutation returns the ZoneMutation object of the builder.
func (zuo *ZoneUpdateOne) Mutation() *ZoneMutation {
	return zuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (zuo *ZoneUpdateOne) ClearOwner() *ZoneUpdateOne {
	zuo.mutation.ClearOwner()
	return zuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (zuo *ZoneUpdateOne) Select(field string, fields ...string) *ZoneUpdateOne {
	zuo.fields = append([]string{field}, fields...)
	return zuo
}

// Save executes the query and returns the updated Zone entity.
func (zuo *ZoneUpdateOne) Save(ctx context.Context) (*Zone, error) {
	var (
		err  error
		node *Zone
	)
	if len(zuo.hooks) == 0 {
		node, err = zuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ZoneMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			zuo.mutation = mutation
			node, err = zuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(zuo.hooks) - 1; i >= 0; i-- {
			mut = zuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, zuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (zuo *ZoneUpdateOne) SaveX(ctx context.Context) *Zone {
	node, err := zuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (zuo *ZoneUpdateOne) Exec(ctx context.Context) error {
	_, err := zuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (zuo *ZoneUpdateOne) ExecX(ctx context.Context) {
	if err := zuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (zuo *ZoneUpdateOne) sqlSave(ctx context.Context) (_node *Zone, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   zone.Table,
			Columns: zone.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: zone.FieldID,
			},
		},
	}
	id, ok := zuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Zone.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := zuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, zone.FieldID)
		for _, f := range fields {
			if !zone.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != zone.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := zuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := zuo.mutation.Domain(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: zone.FieldDomain,
		})
	}
	if zuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   zone.OwnerTable,
			Columns: []string{zone.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := zuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   zone.OwnerTable,
			Columns: []string{zone.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Zone{config: zuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, zuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{zone.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
