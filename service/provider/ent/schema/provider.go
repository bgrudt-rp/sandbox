package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

// Provider holds the schema definition for the Provider entity.
type Provider struct {
	ent.Schema
}

// Mixin common fields - created and updated time and user
func (Provider) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		UserMixin{},
	}
}

// Fields of the Provider.
func (Provider) Fields() []ent.Field {
	return []ent.Field{
		field.String("npi").Unique().Optional(),
		field.String("name_last").Optional(),
		field.String("name_first").Optional(),
		field.String("name_middle").Optional(),
		field.String("name_prefix").Optional(),
		field.String("name_suffix").Optional(),
		field.String("name_degree").Optional(),
		field.Time("birth_date").Optional(),
		field.String("home_addr_location_id").Optional(),
	}
}

// Edges of the provider
func (Provider) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("provider_credential", ProviderCredential.Type),
		edge.From("provider_type", ProviderType.Type).Unique(),
	}
}

// Provider indexes
func (Provider) Indexes() []ent.Index {
	return []ent.Index{
		// NPI is used for lookups
		index.Fields("npi"),
	}
}
