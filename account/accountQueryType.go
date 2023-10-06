package account

import (
	"context"
	"database/sql"
	"github.com/graphql-go/graphql"
)

var db = &sql.DB{}
var accountRepo = NewAccountRepository(db)
var accountService = NewAccountService(accountRepo)

// create object response
var InsertAccountResponse1 = graphql.NewObject(
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

// create query type
var AccountQueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "query",
		Fields: graphql.Fields{
			"account": &graphql.Field{
				Name: "account",
				Type: InsertAccountResponse1,
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					ctx := context.Background()
					uname := p.Args["username"].(string)
					return accountService.GetByEmail(ctx, uname)
				},
			},
		},
	},
)
