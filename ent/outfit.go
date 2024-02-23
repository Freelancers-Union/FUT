// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/mikestefanello/pagoda/ent/outfit"
)

// Outfit is the model entity for the Outfit schema.
type Outfit struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Tag holds the value of the "tag" field.
	Tag string `json:"tag,omitempty"`
	// Faction holds the value of the "faction" field.
	Faction uint8 `json:"faction,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OutfitQuery when eager-loading is set.
	Edges        OutfitEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OutfitEdges holds the relations/edges for other nodes in the graph.
type OutfitEdges struct {
	// Characters holds the value of the characters edge.
	Characters []*Character `json:"characters,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CharactersOrErr returns the Characters value or an error if the edge
// was not loaded in eager-loading.
func (e OutfitEdges) CharactersOrErr() ([]*Character, error) {
	if e.loadedTypes[0] {
		return e.Characters, nil
	}
	return nil, &NotLoadedError{edge: "characters"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Outfit) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case outfit.FieldID, outfit.FieldFaction:
			values[i] = new(sql.NullInt64)
		case outfit.FieldName, outfit.FieldTag:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Outfit fields.
func (o *Outfit) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case outfit.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = uint64(value.Int64)
		case outfit.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				o.Name = value.String
			}
		case outfit.FieldTag:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tag", values[i])
			} else if value.Valid {
				o.Tag = value.String
			}
		case outfit.FieldFaction:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field faction", values[i])
			} else if value.Valid {
				o.Faction = uint8(value.Int64)
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Outfit.
// This includes values selected through modifiers, order, etc.
func (o *Outfit) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// QueryCharacters queries the "characters" edge of the Outfit entity.
func (o *Outfit) QueryCharacters() *CharacterQuery {
	return NewOutfitClient(o.config).QueryCharacters(o)
}

// Update returns a builder for updating this Outfit.
// Note that you need to call Outfit.Unwrap() before calling this method if this Outfit
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Outfit) Update() *OutfitUpdateOne {
	return NewOutfitClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Outfit entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Outfit) Unwrap() *Outfit {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Outfit is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Outfit) String() string {
	var builder strings.Builder
	builder.WriteString("Outfit(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("name=")
	builder.WriteString(o.Name)
	builder.WriteString(", ")
	builder.WriteString("tag=")
	builder.WriteString(o.Tag)
	builder.WriteString(", ")
	builder.WriteString("faction=")
	builder.WriteString(fmt.Sprintf("%v", o.Faction))
	builder.WriteByte(')')
	return builder.String()
}

// Outfits is a parsable slice of Outfit.
type Outfits []*Outfit