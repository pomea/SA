package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Menutype holds the schema definition for the Menutype entity.
type Menutype struct {
	ent.Schema
}

// Fields of the Menutype.
func (Menutype) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").Unique(),
	}
}

// Edges of the Menutype.
func (Menutype) Edges() []ent.Edge {
	return []ent.Edge{
        edge.To("menus", Menu.Type).StorageKey(edge.Column("type_id")),
	}
}