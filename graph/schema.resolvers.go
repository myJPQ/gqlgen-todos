package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/myJPQ/gqlgen-todos/graph/generated"
	"github.com/myJPQ/gqlgen-todos/graph/model"
)

func (r *mutationResolver) UpdateMeeting(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	
	if r.todos ==nil{
		r.todos=make(map[int]*model.Todo)
	}	
	todo := &model.Todo{
		ID:      input.ID,
		Meeting: &model.Meeting{Title: input.Title, Description: input.Description, StartTime: input.StartTime},
	}
	r.todos[input.ID] = todo

	return todo, nil
}

func (r *queryResolver) Getmeeting(ctx context.Context, id int) (*model.Todo, error) {
	return r.todos[id], nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
