// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Jecosine/alioth-kratos/app/data_service/ent/judgerecord"
	"github.com/Jecosine/alioth-kratos/app/data_service/ent/user"
)

// JudgeRecordCreate is the builder for creating a JudgeRecord entity.
type JudgeRecordCreate struct {
	config
	mutation *JudgeRecordMutation
	hooks    []Hook
}

// SetJudgeTime sets the "judge_time" field.
func (jrc *JudgeRecordCreate) SetJudgeTime(t time.Time) *JudgeRecordCreate {
	jrc.mutation.SetJudgeTime(t)
	return jrc
}

// SetFinishedTime sets the "finished_time" field.
func (jrc *JudgeRecordCreate) SetFinishedTime(t time.Time) *JudgeRecordCreate {
	jrc.mutation.SetFinishedTime(t)
	return jrc
}

// SetTimeCost sets the "time_cost" field.
func (jrc *JudgeRecordCreate) SetTimeCost(i int64) *JudgeRecordCreate {
	jrc.mutation.SetTimeCost(i)
	return jrc
}

// SetMemoryCost sets the "memory_cost" field.
func (jrc *JudgeRecordCreate) SetMemoryCost(i int64) *JudgeRecordCreate {
	jrc.mutation.SetMemoryCost(i)
	return jrc
}

// SetStatus sets the "status" field.
func (jrc *JudgeRecordCreate) SetStatus(i int64) *JudgeRecordCreate {
	jrc.mutation.SetStatus(i)
	return jrc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (jrc *JudgeRecordCreate) SetNillableStatus(i *int64) *JudgeRecordCreate {
	if i != nil {
		jrc.SetStatus(*i)
	}
	return jrc
}

// SetID sets the "id" field.
func (jrc *JudgeRecordCreate) SetID(i int64) *JudgeRecordCreate {
	jrc.mutation.SetID(i)
	return jrc
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (jrc *JudgeRecordCreate) AddUserIDs(ids ...int64) *JudgeRecordCreate {
	jrc.mutation.AddUserIDs(ids...)
	return jrc
}

// AddUser adds the "user" edges to the User entity.
func (jrc *JudgeRecordCreate) AddUser(u ...*User) *JudgeRecordCreate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return jrc.AddUserIDs(ids...)
}

// Mutation returns the JudgeRecordMutation object of the builder.
func (jrc *JudgeRecordCreate) Mutation() *JudgeRecordMutation {
	return jrc.mutation
}

// Save creates the JudgeRecord in the database.
func (jrc *JudgeRecordCreate) Save(ctx context.Context) (*JudgeRecord, error) {
	var (
		err  error
		node *JudgeRecord
	)
	jrc.defaults()
	if len(jrc.hooks) == 0 {
		if err = jrc.check(); err != nil {
			return nil, err
		}
		node, err = jrc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JudgeRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = jrc.check(); err != nil {
				return nil, err
			}
			jrc.mutation = mutation
			if node, err = jrc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(jrc.hooks) - 1; i >= 0; i-- {
			if jrc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = jrc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, jrc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (jrc *JudgeRecordCreate) SaveX(ctx context.Context) *JudgeRecord {
	v, err := jrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jrc *JudgeRecordCreate) Exec(ctx context.Context) error {
	_, err := jrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jrc *JudgeRecordCreate) ExecX(ctx context.Context) {
	if err := jrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jrc *JudgeRecordCreate) defaults() {
	if _, ok := jrc.mutation.Status(); !ok {
		v := judgerecord.DefaultStatus
		jrc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jrc *JudgeRecordCreate) check() error {
	if _, ok := jrc.mutation.JudgeTime(); !ok {
		return &ValidationError{Name: "judge_time", err: errors.New(`ent: missing required field "JudgeRecord.judge_time"`)}
	}
	if _, ok := jrc.mutation.FinishedTime(); !ok {
		return &ValidationError{Name: "finished_time", err: errors.New(`ent: missing required field "JudgeRecord.finished_time"`)}
	}
	if _, ok := jrc.mutation.TimeCost(); !ok {
		return &ValidationError{Name: "time_cost", err: errors.New(`ent: missing required field "JudgeRecord.time_cost"`)}
	}
	if _, ok := jrc.mutation.MemoryCost(); !ok {
		return &ValidationError{Name: "memory_cost", err: errors.New(`ent: missing required field "JudgeRecord.memory_cost"`)}
	}
	if _, ok := jrc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "JudgeRecord.status"`)}
	}
	return nil
}

func (jrc *JudgeRecordCreate) sqlSave(ctx context.Context) (*JudgeRecord, error) {
	_node, _spec := jrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, jrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (jrc *JudgeRecordCreate) createSpec() (*JudgeRecord, *sqlgraph.CreateSpec) {
	var (
		_node = &JudgeRecord{config: jrc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: judgerecord.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: judgerecord.FieldID,
			},
		}
	)
	if id, ok := jrc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := jrc.mutation.JudgeTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: judgerecord.FieldJudgeTime,
		})
		_node.JudgeTime = value
	}
	if value, ok := jrc.mutation.FinishedTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: judgerecord.FieldFinishedTime,
		})
		_node.FinishedTime = value
	}
	if value, ok := jrc.mutation.TimeCost(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: judgerecord.FieldTimeCost,
		})
		_node.TimeCost = value
	}
	if value, ok := jrc.mutation.MemoryCost(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: judgerecord.FieldMemoryCost,
		})
		_node.MemoryCost = value
	}
	if value, ok := jrc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: judgerecord.FieldStatus,
		})
		_node.Status = value
	}
	if nodes := jrc.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// JudgeRecordCreateBulk is the builder for creating many JudgeRecord entities in bulk.
type JudgeRecordCreateBulk struct {
	config
	builders []*JudgeRecordCreate
}

// Save creates the JudgeRecord entities in the database.
func (jrcb *JudgeRecordCreateBulk) Save(ctx context.Context) ([]*JudgeRecord, error) {
	specs := make([]*sqlgraph.CreateSpec, len(jrcb.builders))
	nodes := make([]*JudgeRecord, len(jrcb.builders))
	mutators := make([]Mutator, len(jrcb.builders))
	for i := range jrcb.builders {
		func(i int, root context.Context) {
			builder := jrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*JudgeRecordMutation)
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
					_, err = mutators[i+1].Mutate(root, jrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, jrcb.driver, spec); err != nil {
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
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, jrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (jrcb *JudgeRecordCreateBulk) SaveX(ctx context.Context) []*JudgeRecord {
	v, err := jrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jrcb *JudgeRecordCreateBulk) Exec(ctx context.Context) error {
	_, err := jrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jrcb *JudgeRecordCreateBulk) ExecX(ctx context.Context) {
	if err := jrcb.Exec(ctx); err != nil {
		panic(err)
	}
}