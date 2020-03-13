package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

// ProviderCredential holds the schema definition for the Provide credential entity.
type ProviderCredential struct {
	ent.Schema
}

// Mixin common fields - created and updated time and user
func (ProviderCredential) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		UserMixin{},
	}
}

// Fields of the credential
func (ProviderCredential) Fields() []ent.Field {
	return []ent.Field{
		field.String("credential_code"),
		field.String("activation_date"),
		field.String("deactivation_date"),
	}
}

// Edges of the credential
func (ProviderCredential) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("provider", Provider.Type).Ref("provider_credentials").Unique(),
		edge.From("credtype", Credtype.Type).Ref("credential_type").Unique(),
		edge.From("credent", Credent.Type).Ref("credentialing_entity").Unique(),
	}
}

// Indexes of the credential
func (ProviderCredential) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("credential_code").Edges("credtype").Edges("credent").Unique(),
		index.Fields("activation_date").Edges("provider").Edges("credtype").Edges("credent").Unique(),
	}
}
