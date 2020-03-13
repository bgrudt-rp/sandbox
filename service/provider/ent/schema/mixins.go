package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

type TimeMixin struct{}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

type UserMixin struct{}

func (UserMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("created_by").Immutable(),
		field.String("updated_by"),
	}
}
