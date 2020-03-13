package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Credtype holds the schema definition for provider credential types
type Credtype struct {
	ent.Schema
}

// Mixin common fields - created and updated time and user
func (Credtype) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		UserMixin{},
	}
}

// Fields of the credential type
func (Credtype) Fields() []ent.Field {
	return []ent.Field{
		field.String("credential_type_name").Unique(), // State License, Payor Credentialing, MC/MD, Workstation Login, etc.
	}
}

// Edges of the credential type
func (Credtype) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("provider_credential", ProviderCredential.Type),
	}
}
