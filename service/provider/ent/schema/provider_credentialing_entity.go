package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Credtype holds the schema definition for provider credential types
type Credent struct {
	ent.Schema
}

// Mixin common fields - created and updated time and user
func (Credent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		UserMixin{},
	}
}

// Fields of the credential type
func (Credent) Fields() []ent.Field {
	return []ent.Field{
		field.String("credentialing_entity_name"),
		field.String("credentialing_entity_location_id").Optional(),
	}
}

// Edges of the credential type
func (Credent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("provider_credential", ProviderCredential.Type),
	}
}
