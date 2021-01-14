package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Menu holds the schema definition for the Menu entity.
type Menu struct {
	ent.Schema
}

// Fields of the Menu.
func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.String("menuname").NotEmpty(),
		field.String("ingredient"),
		field.Int("calories").Positive(),
		field.Time("added_time"),
	}
}

// Edges of the Menu.
func (Menu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("menus").Unique(),
		edge.From("type", Menutype.Type).Ref("menus").Unique(),
		edge.From("group", Menugroup.Type).Ref("menus").Unique(),
    }
}
