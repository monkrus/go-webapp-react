package mutations

import (
	fields "app/mutations/fields"
	"github.com/graphql-go/graphql"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createNotTodo": fields.CreateNotTodo,
	},
})
