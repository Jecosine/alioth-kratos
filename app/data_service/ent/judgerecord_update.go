// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Jecosine/alioth-kratos/app/data_service/ent/judgerecord"
	"github.com/Jecosine/alioth-kratos/app/data_service/ent/predicate"
	"github.com/Jecosine/alioth-kratos/app/data_service/ent/user"
)

// JudgeRecordUpdate is the builder for updating JudgeRecord entities.
type JudgeRecordUpdate struct {
	config
	hooks    []Hook
	mutation *JudgeRecordMutation
}

// Where appends a list predicates to the JudgeRecordUpdate builder.
func (jru *JudgeRecordUpdate) Where(ps ...predicate.JudgeRecord) *JudgeRecordUpdate {
	jru.mutation.Where(ps...)
	return jru
}

// SetJudgeTime sets the "judge_time" field.
func (jru *JudgeRecordUpdate) SetJudgeTime(t time.Time) *JudgeRecordUpdate {
	jru.mutation.SetJudgeTime(t)
	return jru
}

// SetFinishedTime sets the "finished_time" field.
func (jru *JudgeRecordUpdate) SetFinishedTime(t time.Time) *JudgeRecordUpdate {
	jru.mutation.SetFinishedTime(t)
	return jru
}

// SetTimeCost sets the "time_cost" field.
func (jru *JudgeRecordUpdate) SetTimeCost(i int64) *JudgeRecordUpdate {
	jru.mutation.ResetTimeCost()
	jru.mutation.SetTimeCost(i)
	return jru
}

// AddTimeCost adds i to the "time_cost" field.
func (jru *JudgeRecordUpdate) AddTimeCost(i int64) *JudgeRecordUpdate {
	jru.mutation.AddTimeCost(i)
	return jru
}

// SetMemoryCost sets the "memory_cost" field.
func (jru *JudgeRecordUpdate) SetMemoryCost(i int64) *JudgeRecordUpdate {
	jru.mutation.ResetMemoryCost()
	jru.mutation.SetMemoryCost(i)
	return jru
}

// AddMemoryCost adds i to the "memory_cost" field.
func (jru *JudgeRecordUpdate) AddMemoryCost(i int64) *JudgeRecordUpdate {
	jru.mutation.AddMemoryCost(i)
	return jru
}

// SetStatus sets the "status" field.
func (jru *JudgeRecordUpdate) SetStatus(i int64) *JudgeRecordUpdate {
	jru.mutation.ResetStatus()
	jru.mutation.SetStatus(i)
	return jru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (jru *JudgeRecordUpdate) SetNillableStatus(i *int64) *JudgeRecordUpdate {
	if i != nil {
		jru.SetStatus(*i)
	}
	return jru
}

// AddStatus adds i to the "status" field.
func (jru *JudgeRecordUpdate) AddStatus(i int64) *JudgeRecordUpdate {
	jru.mutation.AddStatus(i)
	return jru
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (jru *JudgeRecordUpdate) AddUserIDs(ids ...int64) *JudgeRecordUpdate {
	jru.mutation.AddUserIDs(ids...)
	return jru
}

// AddUser adds the "user" edges to the User entity.
func (jru *JudgeRecordUpdate) AddUser(u ...*User) *JudgeRecordUpdate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return jru.AddUserIDs(ids...)
}

// Mutation returns the JudgeRecordMutation object of the builder.
func (jru *JudgeRecordUpdate) Mutation() *JudgeRecordMutation {
	return jru.mutation
}

// ClearUser clears all "user" edges to the User entity.
func (jru *JudgeRecordUpdate) ClearUser() *JudgeRecordUpdate {
	jru.mutation.ClearUser()
	return jru
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (jru *JudgeRecordUpdate) RemoveUserIDs(ids ...int64) *JudgeRecordUpdate {
	jru.mutation.RemoveUserIDs(ids...)
	return jru
}

// RemoveUser removes "user" edges to User entities.
func (jru *JudgeRecordUpdate) RemoveUser(u ...*User) *JudgeRecordUpdate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return jru.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (jru *JudgeRecordUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(jru.hooks) == 0 {
		affected, err = jru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JudgeRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			jru.mutation = mutation
			affected, err = jru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(jru.hooks) - 1; i >= 0; i-- {
			if jru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = jru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, jru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (jru *JudgeRecordUpdate) SaveX(ctx context.Context) int {
	affected, err := jru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (jru *JudgeRecordUpdate) Exec(ctx context.Context) error {
	_, err := jru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jru *JudgeRecordUpdate) ExecX(ctx context.Context) {
	if err := jru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (jru *JudgeRecordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   judgerecord.Table,
			Columns: judgerecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: judgerecord.FieldID,
			},
		},
	}
	if ps := jru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jru.mutation.JudgeTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: judgerecord.FieldJudgeTime,
		})
	}
	if value, ok := jru.mutation.FinishedTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: judgerecord.FieldFinishedTime,
		})
	}
	if value, ok := jru.mutation.TimeCost(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: judgerecord.FieldTimeCost,
		})
	}
	if value, ok := jru.mutation.AddedTimeCost(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: judgerecord.FieldTimeCost,
		})
	}
	if value, ok := jru.mutation.MemoryCost(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: judgerecord.FieldMemoryCost,
		})
	}
	if value, ok := jru.mutation.AddedMemoryCost(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: judgerecord.FieldMemoryCost,
		})
	}
	if value, ok := jru.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: judgerecord.FieldStatus,
		})
	}
	if value, ok := jru.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: judgerecord.FieldStatus,
		})
	}
	if jru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   judgerecord.UserTable,
			Columns: []string{judgerecord.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jru.mutation.RemovedUserIDs(); len(nodes) > 0 && !jru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   judgerecord.UserTable,
			Columns: []string{judgerecord.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   judgerecord.UserTable,
			Columns: []string{judgerecord.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, jru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{judgerecord.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// JudgeRecordUpdateOne is the builder for updating a single JudgeRecord entity.
type JudgeRecordUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *JudgeRecordMutation
}

// SetJudgeTime sets the "judge_time" field.
func (jruo *JudgeRecordUpdateOne) SetJudgeTime(t time.Time) *JudgeRecordUpdateOne {
	jruo.mutation.SetJudgeTime(t)
	return jruo
}

// SetFinishedTime sets the "finished_time" field.
func (jruo *JudgeRecordUpdateOne) SetFinishedTime(t time.Time) *JudgeRecordUpdateOne {
	jruo.mutation.SetFinishedTime(t)
	return jruo
}

// SetTimeCost sets the "time_cost" field.
func (jruo *JudgeRecordUpdateOne) SetTimeCost(i int64) *JudgeRecordUpdateOne {
	jruo.mutation.ResetTimeCost()
	jruo.mutation.SetTimeCost(i)
	return jruo
}

// AddTimeCost adds i to the "time_cost" field.
func (jruo *JudgeRecordUpdateOne) AddTimeCost(i int64) *JudgeRecordUpdateOne {
	jruo.mutation.AddTimeCost(i)
	return jruo
}

// SetMemoryCost sets the "memory_cost" field.
func (jruo *JudgeRecordUpdateOne) SetMemoryCost(i int64) *JudgeRecordUpdateOne {
	jruo.mutation.ResetMemoryCost()
	jruo.mutation.SetMemoryCost(i)
	return jruo
}

// AddMemoryCost adds i to the "memory_cost" field.
func (jruo *JudgeRecordUpdateOne) AddMemoryCost(i int64) *JudgeRecordUpdateOne {
	jruo.mutation.AddMemoryCost(i)
	return jruo
}

// SetStatus sets the "status" field.
func (jruo *JudgeRecordUpdateOne) SetStatus(i int64) *JudgeRecordUpdateOne {
	jruo.mutation.ResetStatus()
	jruo.mutation.SetStatus(i)
	return jruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (jruo *JudgeRecordUpdateOne) SetNillableStatus(i *int64) *JudgeRecordUpdateOne {
	if i != nil {
		jruo.SetStatus(*i)
	}
	return jruo
}

// AddStatus adds i to the "status" field.
func (jruo *JudgeRecordUpdateOne) AddStatus(i int64) *JudgeRecordUpdateOne {
	jruo.mutation.AddStatus(i)
	return jruo
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (jruo *JudgeRecordUpdateOne) AddUserIDs(ids ...int64) *JudgeRecordUpdateOne {
	jruo.mutation.AddUserIDs(ids...)
	return jruo
}

// AddUser adds the "user" edges to the User entity.
func (jruo *JudgeRecordUpdateOne) AddUser(u ...*User) *JudgeRecordUpdateOne {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return jruo.AddUserIDs(ids...)
}

// Mutation returns the JudgeRecordMutation object of the builder.
func (jruo *JudgeRecordUpdateOne) Mutation() *JudgeRecordMutation {
	return jruo.mutation
}

// ClearUser clears all "user" edges to the User entity.
func (jruo *JudgeRecordUpdateOne) ClearUser() *JudgeRecordUpdateOne {
	jruo.mutation.ClearUser()
	return jruo
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (jruo *JudgeRecordUpdateOne) RemoveUserIDs(ids ...int64) *JudgeRecordUpdateOne {
	jruo.mutation.RemoveUserIDs(ids...)
	return jruo
}

// RemoveUser removes "user" edges to User entities.
func (jruo *JudgeRecordUpdateOne) RemoveUser(u ...*User) *JudgeRecordUpdateOne {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return jruo.RemoveUserIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (jruo *JudgeRecordUpdateOne) Select(field string, fields ...string) *JudgeRecordUpdateOne {
	jruo.fields = append([]string{field}, fields...)
	return jruo
}

// Save executes the query and returns the updated JudgeRecord entity.
func (jruo *JudgeRecordUpdateOne) Save(ctx context.Context) (*JudgeRecord, error) {
	var (
		err  error
		node *JudgeRecord
	)
	if len(jruo.hooks) == 0 {
		node, err = jruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JudgeRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			jruo.mutation = mutation
			node, err = jruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(jruo.hooks) - 1; i >= 0; i-- {
			if jruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = jruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, jruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (jruo *JudgeRecordUpdateOne) SaveX(ctx context.Context) *JudgeRecord {
	node, err := jruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (jruo *JudgeRecordUpdateOne) Exec(ctx context.Context) error {
	_, err := jruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jruo *JudgeRecordUpdateOne) ExecX(ctx context.Context) {
	if err := jruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (jruo *JudgeRecordUpdateOne) sqlSave(ctx context.Context) (_node *JudgeRecord, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   judgerecord.Table,
			Columns: judgerecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: judgerecord.FieldID,
			},
		},
	}
	id, ok := jruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "JudgeRecord.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := jruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, judgerecord.FieldID)
		for _, f := range fields {
			if !judgerecord.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != judgerecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := jruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jruo.mutation.JudgeTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: judgerecord.FieldJudgeTime,
		})
	}
	if value, ok := jruo.mutation.FinishedTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: judgerecord.FieldFinishedTime,
		})
	}
	if value, ok := jruo.mutation.TimeCost(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: judgerecord.FieldTimeCost,
		})
	}
	if value, ok := jruo.mutation.AddedTimeCost(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: judgerecord.FieldTimeCost,
		})
	}
	if value, ok := jruo.mutation.MemoryCost(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: judgerecord.FieldMemoryCost,
		})
	}
	if value, ok := jruo.mutation.AddedMemoryCost(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: judgerecord.FieldMemoryCost,
		})
	}
	if value, ok := jruo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: judgerecord.FieldStatus,
		})
	}
	if value, ok := jruo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: judgerecord.FieldStatus,
		})
	}
	if jruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   judgerecord.UserTable,
			Columns: []string{judgerecord.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jruo.mutation.RemovedUserIDs(); len(nodes) > 0 && !jruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   judgerecord.UserTable,
			Columns: []string{judgerecord.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   judgerecord.UserTable,
			Columns: []string{judgerecord.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &JudgeRecord{config: jruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, jruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{judgerecord.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}