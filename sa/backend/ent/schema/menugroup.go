package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Menugroup holds the schema definition for the Menugroup entity.
type Menugroup struct {
	ent.Schema
}

// Fields of the Menugroup.
func (Menugroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("group").Unique(),
	}
}

// Edges of the Menugroup.
func (Menugroup) Edges() []ent.Edge {
	return []ent.Edge{
        edge.To("menus", Menu.Type).StorageKey(edge.Column("group_id")),
	}
}