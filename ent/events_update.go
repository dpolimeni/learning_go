// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dpolimeni/fiber_app/ent/events"
	"github.com/dpolimeni/fiber_app/ent/predicate"
)

// EventsUpdate is the builder for updating Events entities.
type EventsUpdate struct {
	config
	hooks    []Hook
	mutation *EventsMutation
}

// Where appends a list predicates to the EventsUpdate builder.
func (eu *EventsUpdate) Where(ps ...predicate.Events) *EventsUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetName sets the "name" field.
func (eu *EventsUpdate) SetName(s string) *EventsUpdate {
	eu.mutation.SetName(s)
	return eu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (eu *EventsUpdate) SetNillableName(s *string) *EventsUpdate {
	if s != nil {
		eu.SetName(*s)
	}
	return eu
}

// SetCapacity sets the "capacity" field.
func (eu *EventsUpdate) SetCapacity(i int16) *EventsUpdate {
	eu.mutation.ResetCapacity()
	eu.mutation.SetCapacity(i)
	return eu
}

// SetNillableCapacity sets the "capacity" field if the given value is not nil.
func (eu *EventsUpdate) SetNillableCapacity(i *int16) *EventsUpdate {
	if i != nil {
		eu.SetCapacity(*i)
	}
	return eu
}

// AddCapacity adds i to the "capacity" field.
func (eu *EventsUpdate) AddCapacity(i int16) *EventsUpdate {
	eu.mutation.AddCapacity(i)
	return eu
}

// SetDescription sets the "description" field.
func (eu *EventsUpdate) SetDescription(s string) *EventsUpdate {
	eu.mutation.SetDescription(s)
	return eu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (eu *EventsUpdate) SetNillableDescription(s *string) *EventsUpdate {
	if s != nil {
		eu.SetDescription(*s)
	}
	return eu
}

// Mutation returns the EventsMutation object of the builder.
func (eu *EventsUpdate) Mutation() *EventsMutation {
	return eu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EventsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EventsUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EventsUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EventsUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *EventsUpdate) check() error {
	if v, ok := eu.mutation.Capacity(); ok {
		if err := events.CapacityValidator(v); err != nil {
			return &ValidationError{Name: "capacity", err: fmt.Errorf(`ent: validator failed for field "Events.capacity": %w`, err)}
		}
	}
	return nil
}

func (eu *EventsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(events.Table, events.Columns, sqlgraph.NewFieldSpec(events.FieldID, field.TypeInt32))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Name(); ok {
		_spec.SetField(events.FieldName, field.TypeString, value)
	}
	if value, ok := eu.mutation.Capacity(); ok {
		_spec.SetField(events.FieldCapacity, field.TypeInt16, value)
	}
	if value, ok := eu.mutation.AddedCapacity(); ok {
		_spec.AddField(events.FieldCapacity, field.TypeInt16, value)
	}
	if value, ok := eu.mutation.Description(); ok {
		_spec.SetField(events.FieldDescription, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{events.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EventsUpdateOne is the builder for updating a single Events entity.
type EventsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EventsMutation
}

// SetName sets the "name" field.
func (euo *EventsUpdateOne) SetName(s string) *EventsUpdateOne {
	euo.mutation.SetName(s)
	return euo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (euo *EventsUpdateOne) SetNillableName(s *string) *EventsUpdateOne {
	if s != nil {
		euo.SetName(*s)
	}
	return euo
}

// SetCapacity sets the "capacity" field.
func (euo *EventsUpdateOne) SetCapacity(i int16) *EventsUpdateOne {
	euo.mutation.ResetCapacity()
	euo.mutation.SetCapacity(i)
	return euo
}

// SetNillableCapacity sets the "capacity" field if the given value is not nil.
func (euo *EventsUpdateOne) SetNillableCapacity(i *int16) *EventsUpdateOne {
	if i != nil {
		euo.SetCapacity(*i)
	}
	return euo
}

// AddCapacity adds i to the "capacity" field.
func (euo *EventsUpdateOne) AddCapacity(i int16) *EventsUpdateOne {
	euo.mutation.AddCapacity(i)
	return euo
}

// SetDescription sets the "description" field.
func (euo *EventsUpdateOne) SetDescription(s string) *EventsUpdateOne {
	euo.mutation.SetDescription(s)
	return euo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (euo *EventsUpdateOne) SetNillableDescription(s *string) *EventsUpdateOne {
	if s != nil {
		euo.SetDescription(*s)
	}
	return euo
}

// Mutation returns the EventsMutation object of the builder.
func (euo *EventsUpdateOne) Mutation() *EventsMutation {
	return euo.mutation
}

// Where appends a list predicates to the EventsUpdate builder.
func (euo *EventsUpdateOne) Where(ps ...predicate.Events) *EventsUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EventsUpdateOne) Select(field string, fields ...string) *EventsUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Events entity.
func (euo *EventsUpdateOne) Save(ctx context.Context) (*Events, error) {
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EventsUpdateOne) SaveX(ctx context.Context) *Events {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EventsUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EventsUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *EventsUpdateOne) check() error {
	if v, ok := euo.mutation.Capacity(); ok {
		if err := events.CapacityValidator(v); err != nil {
			return &ValidationError{Name: "capacity", err: fmt.Errorf(`ent: validator failed for field "Events.capacity": %w`, err)}
		}
	}
	return nil
}

func (euo *EventsUpdateOne) sqlSave(ctx context.Context) (_node *Events, err error) {
	if err := euo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(events.Table, events.Columns, sqlgraph.NewFieldSpec(events.FieldID, field.TypeInt32))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Events.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, events.FieldID)
		for _, f := range fields {
			if !events.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != events.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Name(); ok {
		_spec.SetField(events.FieldName, field.TypeString, value)
	}
	if value, ok := euo.mutation.Capacity(); ok {
		_spec.SetField(events.FieldCapacity, field.TypeInt16, value)
	}
	if value, ok := euo.mutation.AddedCapacity(); ok {
		_spec.AddField(events.FieldCapacity, field.TypeInt16, value)
	}
	if value, ok := euo.mutation.Description(); ok {
		_spec.SetField(events.FieldDescription, field.TypeString, value)
	}
	_node = &Events{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{events.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}