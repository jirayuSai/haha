package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("CustomerName"),
		field.String("Address"),
		field.String("PhoneNumber"),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("genders", Gender.Type).
			Ref("customers").
			Unique(),

		edge.From("personals", Personal.Type).
			Ref("customers").
			Unique(),

		edge.From("titles", Title.Type).
			Ref("customers").
			Unique(),
	}
}
