package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"github.com/MatsuoTakuro/learning-gqlgen/graph/model"
)

type Resolver struct {
	todos  []*model.Todo
	lastID int
}
