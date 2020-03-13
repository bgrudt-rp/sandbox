package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Credtype holds the schema definition for provider credential types
type ProviderType struct {
	ent.Schema
}

// Mixin common fields - created and updated time and user
func (ProviderType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		UserMixin{},
	}
}

// Fields of the provider type
func (ProviderType) Fields() []ent.Field {
	return []ent.Field{
		field.String("provider_type_name").Unique(), // Radiologist, Cardiologist, Technologist, etc.
	}
}

// Edges of the provider type
func (ProviderType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("provider", Provider.Type),
	}
}
