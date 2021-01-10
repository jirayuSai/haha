package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Personal holds the schema definition for the Personal entity.
type Personal struct {
	ent.Schema
}

// Fields of the Personal.
func (Personal) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name"),
		field.String("Email"),
		field.String("Password").
			Unique(),
	}
}

// Edges of the Personal.
func (Personal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("customers", Customer.Type).StorageKey(edge.Column("Personal_ID")),
	}
}
