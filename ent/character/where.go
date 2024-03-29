// Code generated by ent, DO NOT EDIT.

package character

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Character {
	return predicate.Character(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Character {
	return predicate.Character(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Character {
	return predicate.Character(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Character {
	return predicate.Character(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Character {
	return predicate.Character(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Character {
	return predicate.Character(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Character {
	return predicate.Character(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Character {
	return predicate.Character(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Character {
	return predicate.Character(sql.FieldLTE(FieldID, id))
}

// CharacterID applies equality check predicate on the "character_id" field. It's identical to CharacterIDEQ.
func CharacterID(v uint64) predicate.Character {
	return predicate.Character(sql.FieldEQ(FieldCharacterID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Character {
	return predicate.Character(sql.FieldEQ(FieldName, v))
}

// FactionID applies equality check predicate on the "faction_id" field. It's identical to FactionIDEQ.
func FactionID(v uint8) predicate.Character {
	return predicate.Character(sql.FieldEQ(FieldFactionID, v))
}

// DateCreated applies equality check predicate on the "date_created" field. It's identical to DateCreatedEQ.
func DateCreated(v time.Time) predicate.Character {
	return predicate.Character(sql.FieldEQ(FieldDateCreated, v))
}

// CharacterIDEQ applies the EQ predicate on the "character_id" field.
func CharacterIDEQ(v uint64) predicate.Character {
	return predicate.Character(sql.FieldEQ(FieldCharacterID, v))
}

// CharacterIDNEQ applies the NEQ predicate on the "character_id" field.
func CharacterIDNEQ(v uint64) predicate.Character {
	return predicate.Character(sql.FieldNEQ(FieldCharacterID, v))
}

// CharacterIDIn applies the In predicate on the "character_id" field.
func CharacterIDIn(vs ...uint64) predicate.Character {
	return predicate.Character(sql.FieldIn(FieldCharacterID, vs...))
}

// CharacterIDNotIn applies the NotIn predicate on the "character_id" field.
func CharacterIDNotIn(vs ...uint64) predicate.Character {
	return predicate.Character(sql.FieldNotIn(FieldCharacterID, vs...))
}

// CharacterIDGT applies the GT predicate on the "character_id" field.
func CharacterIDGT(v uint64) predicate.Character {
	return predicate.Character(sql.FieldGT(FieldCharacterID, v))
}

// CharacterIDGTE applies the GTE predicate on the "character_id" field.
func CharacterIDGTE(v uint64) predicate.Character {
	return predicate.Character(sql.FieldGTE(FieldCharacterID, v))
}

// CharacterIDLT applies the LT predicate on the "character_id" field.
func CharacterIDLT(v uint64) predicate.Character {
	return predicate.Character(sql.FieldLT(FieldCharacterID, v))
}

// CharacterIDLTE applies the LTE predicate on the "character_id" field.
func CharacterIDLTE(v uint64) predicate.Character {
	return predicate.Character(sql.FieldLTE(FieldCharacterID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Character {
	return predicate.Character(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Character {
	return predicate.Character(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Character {
	return predicate.Character(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Character {
	return predicate.Character(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Character {
	return predicate.Character(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Character {
	return predicate.Character(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Character {
	return predicate.Character(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Character {
	return predicate.Character(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Character {
	return predicate.Character(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Character {
	return predicate.Character(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Character {
	return predicate.Character(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Character {
	return predicate.Character(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Character {
	return predicate.Character(sql.FieldContainsFold(FieldName, v))
}

// FactionIDEQ applies the EQ predicate on the "faction_id" field.
func FactionIDEQ(v uint8) predicate.Character {
	return predicate.Character(sql.FieldEQ(FieldFactionID, v))
}

// FactionIDNEQ applies the NEQ predicate on the "faction_id" field.
func FactionIDNEQ(v uint8) predicate.Character {
	return predicate.Character(sql.FieldNEQ(FieldFactionID, v))
}

// FactionIDIn applies the In predicate on the "faction_id" field.
func FactionIDIn(vs ...uint8) predicate.Character {
	return predicate.Character(sql.FieldIn(FieldFactionID, vs...))
}

// FactionIDNotIn applies the NotIn predicate on the "faction_id" field.
func FactionIDNotIn(vs ...uint8) predicate.Character {
	return predicate.Character(sql.FieldNotIn(FieldFactionID, vs...))
}

// FactionIDGT applies the GT predicate on the "faction_id" field.
func FactionIDGT(v uint8) predicate.Character {
	return predicate.Character(sql.FieldGT(FieldFactionID, v))
}

// FactionIDGTE applies the GTE predicate on the "faction_id" field.
func FactionIDGTE(v uint8) predicate.Character {
	return predicate.Character(sql.FieldGTE(FieldFactionID, v))
}

// FactionIDLT applies the LT predicate on the "faction_id" field.
func FactionIDLT(v uint8) predicate.Character {
	return predicate.Character(sql.FieldLT(FieldFactionID, v))
}

// FactionIDLTE applies the LTE predicate on the "faction_id" field.
func FactionIDLTE(v uint8) predicate.Character {
	return predicate.Character(sql.FieldLTE(FieldFactionID, v))
}

// DateCreatedEQ applies the EQ predicate on the "date_created" field.
func DateCreatedEQ(v time.Time) predicate.Character {
	return predicate.Character(sql.FieldEQ(FieldDateCreated, v))
}

// DateCreatedNEQ applies the NEQ predicate on the "date_created" field.
func DateCreatedNEQ(v time.Time) predicate.Character {
	return predicate.Character(sql.FieldNEQ(FieldDateCreated, v))
}

// DateCreatedIn applies the In predicate on the "date_created" field.
func DateCreatedIn(vs ...time.Time) predicate.Character {
	return predicate.Character(sql.FieldIn(FieldDateCreated, vs...))
}

// DateCreatedNotIn applies the NotIn predicate on the "date_created" field.
func DateCreatedNotIn(vs ...time.Time) predicate.Character {
	return predicate.Character(sql.FieldNotIn(FieldDateCreated, vs...))
}

// DateCreatedGT applies the GT predicate on the "date_created" field.
func DateCreatedGT(v time.Time) predicate.Character {
	return predicate.Character(sql.FieldGT(FieldDateCreated, v))
}

// DateCreatedGTE applies the GTE predicate on the "date_created" field.
func DateCreatedGTE(v time.Time) predicate.Character {
	return predicate.Character(sql.FieldGTE(FieldDateCreated, v))
}

// DateCreatedLT applies the LT predicate on the "date_created" field.
func DateCreatedLT(v time.Time) predicate.Character {
	return predicate.Character(sql.FieldLT(FieldDateCreated, v))
}

// DateCreatedLTE applies the LTE predicate on the "date_created" field.
func DateCreatedLTE(v time.Time) predicate.Character {
	return predicate.Character(sql.FieldLTE(FieldDateCreated, v))
}

// HasOutfit applies the HasEdge predicate on the "outfit" edge.
func HasOutfit() predicate.Character {
	return predicate.Character(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OutfitTable, OutfitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOutfitWith applies the HasEdge predicate on the "outfit" edge with a given conditions (other predicates).
func HasOutfitWith(preds ...predicate.Outfit) predicate.Character {
	return predicate.Character(func(s *sql.Selector) {
		step := newOutfitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRibbons applies the HasEdge predicate on the "ribbons" edge.
func HasRibbons() predicate.Character {
	return predicate.Character(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, RibbonsTable, RibbonsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRibbonsWith applies the HasEdge predicate on the "ribbons" edge with a given conditions (other predicates).
func HasRibbonsWith(preds ...predicate.Ribbon) predicate.Character {
	return predicate.Character(func(s *sql.Selector) {
		step := newRibbonsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Character) predicate.Character {
	return predicate.Character(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Character) predicate.Character {
	return predicate.Character(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Character) predicate.Character {
	return predicate.Character(sql.NotPredicates(p))
}
