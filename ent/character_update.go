// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/character"
	"github.com/mikestefanello/pagoda/ent/outfit"
	"github.com/mikestefanello/pagoda/ent/predicate"
	"github.com/mikestefanello/pagoda/ent/ribbon"
)

// CharacterUpdate is the builder for updating Character entities.
type CharacterUpdate struct {
	config
	hooks    []Hook
	mutation *CharacterMutation
}

// Where appends a list predicates to the CharacterUpdate builder.
func (cu *CharacterUpdate) Where(ps ...predicate.Character) *CharacterUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CharacterUpdate) SetName(s string) *CharacterUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *CharacterUpdate) SetNillableName(s *string) *CharacterUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// SetFactionID sets the "faction_id" field.
func (cu *CharacterUpdate) SetFactionID(u uint8) *CharacterUpdate {
	cu.mutation.ResetFactionID()
	cu.mutation.SetFactionID(u)
	return cu
}

// SetNillableFactionID sets the "faction_id" field if the given value is not nil.
func (cu *CharacterUpdate) SetNillableFactionID(u *uint8) *CharacterUpdate {
	if u != nil {
		cu.SetFactionID(*u)
	}
	return cu
}

// AddFactionID adds u to the "faction_id" field.
func (cu *CharacterUpdate) AddFactionID(u int8) *CharacterUpdate {
	cu.mutation.AddFactionID(u)
	return cu
}

// SetDateCreated sets the "date_created" field.
func (cu *CharacterUpdate) SetDateCreated(t time.Time) *CharacterUpdate {
	cu.mutation.SetDateCreated(t)
	return cu
}

// SetNillableDateCreated sets the "date_created" field if the given value is not nil.
func (cu *CharacterUpdate) SetNillableDateCreated(t *time.Time) *CharacterUpdate {
	if t != nil {
		cu.SetDateCreated(*t)
	}
	return cu
}

// SetOutfitID sets the "outfit" edge to the Outfit entity by ID.
func (cu *CharacterUpdate) SetOutfitID(id uint64) *CharacterUpdate {
	cu.mutation.SetOutfitID(id)
	return cu
}

// SetNillableOutfitID sets the "outfit" edge to the Outfit entity by ID if the given value is not nil.
func (cu *CharacterUpdate) SetNillableOutfitID(id *uint64) *CharacterUpdate {
	if id != nil {
		cu = cu.SetOutfitID(*id)
	}
	return cu
}

// SetOutfit sets the "outfit" edge to the Outfit entity.
func (cu *CharacterUpdate) SetOutfit(o *Outfit) *CharacterUpdate {
	return cu.SetOutfitID(o.ID)
}

// SetRibbonsID sets the "ribbons" edge to the Ribbon entity by ID.
func (cu *CharacterUpdate) SetRibbonsID(id int) *CharacterUpdate {
	cu.mutation.SetRibbonsID(id)
	return cu
}

// SetNillableRibbonsID sets the "ribbons" edge to the Ribbon entity by ID if the given value is not nil.
func (cu *CharacterUpdate) SetNillableRibbonsID(id *int) *CharacterUpdate {
	if id != nil {
		cu = cu.SetRibbonsID(*id)
	}
	return cu
}

// SetRibbons sets the "ribbons" edge to the Ribbon entity.
func (cu *CharacterUpdate) SetRibbons(r *Ribbon) *CharacterUpdate {
	return cu.SetRibbonsID(r.ID)
}

// Mutation returns the CharacterMutation object of the builder.
func (cu *CharacterUpdate) Mutation() *CharacterMutation {
	return cu.mutation
}

// ClearOutfit clears the "outfit" edge to the Outfit entity.
func (cu *CharacterUpdate) ClearOutfit() *CharacterUpdate {
	cu.mutation.ClearOutfit()
	return cu
}

// ClearRibbons clears the "ribbons" edge to the Ribbon entity.
func (cu *CharacterUpdate) ClearRibbons() *CharacterUpdate {
	cu.mutation.ClearRibbons()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CharacterUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CharacterUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CharacterUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CharacterUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CharacterUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(character.Table, character.Columns, sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(character.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.FactionID(); ok {
		_spec.SetField(character.FieldFactionID, field.TypeUint8, value)
	}
	if value, ok := cu.mutation.AddedFactionID(); ok {
		_spec.AddField(character.FieldFactionID, field.TypeUint8, value)
	}
	if value, ok := cu.mutation.DateCreated(); ok {
		_spec.SetField(character.FieldDateCreated, field.TypeTime, value)
	}
	if cu.mutation.OutfitCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   character.OutfitTable,
			Columns: []string{character.OutfitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(outfit.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.OutfitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   character.OutfitTable,
			Columns: []string{character.OutfitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(outfit.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.RibbonsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   character.RibbonsTable,
			Columns: []string{character.RibbonsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ribbon.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RibbonsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   character.RibbonsTable,
			Columns: []string{character.RibbonsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ribbon.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{character.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CharacterUpdateOne is the builder for updating a single Character entity.
type CharacterUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CharacterMutation
}

// SetName sets the "name" field.
func (cuo *CharacterUpdateOne) SetName(s string) *CharacterUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *CharacterUpdateOne) SetNillableName(s *string) *CharacterUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// SetFactionID sets the "faction_id" field.
func (cuo *CharacterUpdateOne) SetFactionID(u uint8) *CharacterUpdateOne {
	cuo.mutation.ResetFactionID()
	cuo.mutation.SetFactionID(u)
	return cuo
}

// SetNillableFactionID sets the "faction_id" field if the given value is not nil.
func (cuo *CharacterUpdateOne) SetNillableFactionID(u *uint8) *CharacterUpdateOne {
	if u != nil {
		cuo.SetFactionID(*u)
	}
	return cuo
}

// AddFactionID adds u to the "faction_id" field.
func (cuo *CharacterUpdateOne) AddFactionID(u int8) *CharacterUpdateOne {
	cuo.mutation.AddFactionID(u)
	return cuo
}

// SetDateCreated sets the "date_created" field.
func (cuo *CharacterUpdateOne) SetDateCreated(t time.Time) *CharacterUpdateOne {
	cuo.mutation.SetDateCreated(t)
	return cuo
}

// SetNillableDateCreated sets the "date_created" field if the given value is not nil.
func (cuo *CharacterUpdateOne) SetNillableDateCreated(t *time.Time) *CharacterUpdateOne {
	if t != nil {
		cuo.SetDateCreated(*t)
	}
	return cuo
}

// SetOutfitID sets the "outfit" edge to the Outfit entity by ID.
func (cuo *CharacterUpdateOne) SetOutfitID(id uint64) *CharacterUpdateOne {
	cuo.mutation.SetOutfitID(id)
	return cuo
}

// SetNillableOutfitID sets the "outfit" edge to the Outfit entity by ID if the given value is not nil.
func (cuo *CharacterUpdateOne) SetNillableOutfitID(id *uint64) *CharacterUpdateOne {
	if id != nil {
		cuo = cuo.SetOutfitID(*id)
	}
	return cuo
}

// SetOutfit sets the "outfit" edge to the Outfit entity.
func (cuo *CharacterUpdateOne) SetOutfit(o *Outfit) *CharacterUpdateOne {
	return cuo.SetOutfitID(o.ID)
}

// SetRibbonsID sets the "ribbons" edge to the Ribbon entity by ID.
func (cuo *CharacterUpdateOne) SetRibbonsID(id int) *CharacterUpdateOne {
	cuo.mutation.SetRibbonsID(id)
	return cuo
}

// SetNillableRibbonsID sets the "ribbons" edge to the Ribbon entity by ID if the given value is not nil.
func (cuo *CharacterUpdateOne) SetNillableRibbonsID(id *int) *CharacterUpdateOne {
	if id != nil {
		cuo = cuo.SetRibbonsID(*id)
	}
	return cuo
}

// SetRibbons sets the "ribbons" edge to the Ribbon entity.
func (cuo *CharacterUpdateOne) SetRibbons(r *Ribbon) *CharacterUpdateOne {
	return cuo.SetRibbonsID(r.ID)
}

// Mutation returns the CharacterMutation object of the builder.
func (cuo *CharacterUpdateOne) Mutation() *CharacterMutation {
	return cuo.mutation
}

// ClearOutfit clears the "outfit" edge to the Outfit entity.
func (cuo *CharacterUpdateOne) ClearOutfit() *CharacterUpdateOne {
	cuo.mutation.ClearOutfit()
	return cuo
}

// ClearRibbons clears the "ribbons" edge to the Ribbon entity.
func (cuo *CharacterUpdateOne) ClearRibbons() *CharacterUpdateOne {
	cuo.mutation.ClearRibbons()
	return cuo
}

// Where appends a list predicates to the CharacterUpdate builder.
func (cuo *CharacterUpdateOne) Where(ps ...predicate.Character) *CharacterUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CharacterUpdateOne) Select(field string, fields ...string) *CharacterUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Character entity.
func (cuo *CharacterUpdateOne) Save(ctx context.Context) (*Character, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CharacterUpdateOne) SaveX(ctx context.Context) *Character {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CharacterUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CharacterUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CharacterUpdateOne) sqlSave(ctx context.Context) (_node *Character, err error) {
	_spec := sqlgraph.NewUpdateSpec(character.Table, character.Columns, sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Character.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, character.FieldID)
		for _, f := range fields {
			if !character.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != character.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(character.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.FactionID(); ok {
		_spec.SetField(character.FieldFactionID, field.TypeUint8, value)
	}
	if value, ok := cuo.mutation.AddedFactionID(); ok {
		_spec.AddField(character.FieldFactionID, field.TypeUint8, value)
	}
	if value, ok := cuo.mutation.DateCreated(); ok {
		_spec.SetField(character.FieldDateCreated, field.TypeTime, value)
	}
	if cuo.mutation.OutfitCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   character.OutfitTable,
			Columns: []string{character.OutfitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(outfit.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.OutfitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   character.OutfitTable,
			Columns: []string{character.OutfitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(outfit.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.RibbonsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   character.RibbonsTable,
			Columns: []string{character.RibbonsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ribbon.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RibbonsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   character.RibbonsTable,
			Columns: []string{character.RibbonsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ribbon.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Character{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{character.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
