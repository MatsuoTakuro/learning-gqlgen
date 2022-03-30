package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/MatsuoTakuro/learning-gqlgen/graph/generated"
	"github.com/MatsuoTakuro/learning-gqlgen/graph/model"
)

func (r *myQueryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	time.Sleep(220 * time.Millisecond)

	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}
	if intId == 666 {
		panic("critical failure")
	}

	for _, todo := range r.todos {
		if todo.ID == intId {
			fmt.Printf("Todo : %v\n", *todo)
			fmt.Printf("Owner : %v\n\n", *todo.Owner())
			return todo, nil
		}
	}
	return nil, errors.New("not found")
}

func (r *myQueryResolver) LastTodo(ctx context.Context) (*model.Todo, error) {
	if len(r.todos) == 0 {
		return nil, errors.New("not found")
	}
	fmt.Printf("LastTodo : %v\n", *r.todos[len(r.todos)-1])
	fmt.Printf("Owner : %v\n\n", *r.todos[len(r.todos)-1].Owner())
	return r.todos[len(r.todos)-1], nil
}

func (r *myQueryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	for i, v := range r.todos {
		fmt.Printf("Todos #%v : %v\n", i+1, *v)
		fmt.Printf("Owner #%v : %v\n\n", i+1, *v.Owner())
	}
	return r.todos, nil
}

// MyQuery returns generated.MyQueryResolver implementation.
func (r *Resolver) MyQuery() generated.MyQueryResolver { return &myQueryResolver{r} }

type myQueryResolver struct{ *Resolver }
