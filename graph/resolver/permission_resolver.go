package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/MatsuoTakuro/learning-gqlgen/graph/generated"
	"github.com/MatsuoTakuro/learning-gqlgen/graph/model"
)

var (
	you  = &model.User{ID: 1, Name: "You"}
	them = &model.User{ID: 2, Name: "Them"}
)

type ckey string

func getUserId(ctx context.Context) int {
	if id, ok := ctx.Value(ckey("userId")).(int); ok {
		return id
	}
	return you.ID
}

func New() generated.Config {
	t1 := model.Todo{ID: 1, Text: "A todo not to forget", Done: false}
	t1.SetOwner(you)

	t2 := model.Todo{ID: 2, Text: "This is the most important", Done: false}
	t2.SetOwner(you)

	t3 := model.Todo{ID: 3, Text: "Somebody else's todo", Done: false}
	t3.SetOwner(them)

	t4 := model.Todo{ID: 4, Text: "Please do this or else", Done: false}
	t4.SetOwner(you)

	c := generated.Config{
		Resolvers: &Resolver{
			todos: []*model.Todo{
				&t1,
				&t2,
				&t3,
				&t4,
			},
			lastID: t4.ID,
		},
	}
	c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {
		switch role {
		case model.RoleAdmin:
			// No admin for you!
			return nil, nil
		case model.RoleOwner:
			ownable, isOwnable := obj.(model.Ownable)
			if !isOwnable {
				return nil, fmt.Errorf("obj cant be owned")
			}

			if ownable.Owner().ID != getUserId(ctx) {
				return nil, fmt.Errorf("you dont own that")
			}
		}

		return next(ctx)
	}
	c.Directives.User = func(ctx context.Context, obj interface{}, next graphql.Resolver, id string) (interface{}, error) {
		return next(context.WithValue(ctx, ckey("userId"), id))
	}
	return c
}
