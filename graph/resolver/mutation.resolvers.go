package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/MatsuoTakuro/learning-gqlgen/graph/generated"
	"github.com/MatsuoTakuro/learning-gqlgen/graph/model"
	"github.com/mitchellh/mapstructure"
)

func (r *myMutationResolver) CreateTodo(ctx context.Context, todo model.TodoInput) (*model.Todo, error) {
	newID := r.id()

	newTodo := &model.Todo{
		ID:   newID,
		Text: todo.Text,
	}
	newTodo.SetOwner(you)

	if todo.Done != nil {
		newTodo.Done = *todo.Done
	}

	r.todos = append(r.todos, newTodo)

	fmt.Printf("NewTodo : %v\n", *newTodo)
	fmt.Printf("Owner : %v\n\n", *newTodo.Owner())
	return newTodo, nil
}

func (r *myMutationResolver) UpdateTodo(ctx context.Context, id string, changes map[string]interface{}) (*model.Todo, error) {
	var updatedTodo *model.Todo

	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}
	for i := 0; i < len(r.todos); i++ {
		if r.todos[i].ID == intId {
			updatedTodo = r.todos[i]
			break
		}
	}

	if updatedTodo == nil {
		return nil, nil
	}

	err = mapstructure.Decode(changes, updatedTodo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("UpdatedTodo : %v\n", updatedTodo)
	fmt.Printf("Owner : %v\n\n", updatedTodo.Owner())
	return updatedTodo, nil
}

func (r *myMutationResolver) id() int {
	r.lastID++
	return r.lastID
}

// MyMutation returns generated.MyMutationResolver implementation.
func (r *Resolver) MyMutation() generated.MyMutationResolver { return &myMutationResolver{r} }

type myMutationResolver struct{ *Resolver }
