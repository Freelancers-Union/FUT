// Code generated by ent, DO NOT EDIT.

package outfit

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the outfit type in the database.
	Label = "outfit"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldTag holds the string denoting the tag field in the database.
	FieldTag = "tag"
	// FieldFaction holds the string denoting the faction field in the database.
	FieldFaction = "faction"
	// EdgeCharacters holds the string denoting the characters edge name in mutations.
	EdgeCharacters = "characters"
	// Table holds the table name of the outfit in the database.
	Table = "outfits"
	// CharactersTable is the table that holds the characters relation/edge.
	CharactersTable = "characters"
	// CharactersInverseTable is the table name for the Character entity.
	// It exists in this package in order to avoid circular dependency with the "character" package.
	CharactersInverseTable = "characters"
	// CharactersColumn is the table column denoting the characters relation/edge.
	CharactersColumn = "outfit_characters"
)

// Columns holds all SQL columns for outfit fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldTag,
	FieldFaction,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Outfit queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByTag orders the results by the tag field.
func ByTag(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTag, opts...).ToFunc()
}

// ByFaction orders the results by the faction field.
func ByFaction(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFaction, opts...).ToFunc()
}

// ByCharactersCount orders the results by characters count.
func ByCharactersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCharactersStep(), opts...)
	}
}

// ByCharacters orders the results by characters terms.
func ByCharacters(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCharactersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCharactersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CharactersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CharactersTable, CharactersColumn),
	)
}