package account

import "github.com/graphql-go/graphql"

var InsertAccountResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "accountResponse",
		Fields: graphql.Fields{
			"username": &graphql.Field{
				Name: "username",
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Name: "password",
				Type: graphql.String,
			},
		}},
)
