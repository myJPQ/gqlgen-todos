package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/myJPQ/gqlgen-todos/graph/generated"
	"github.com/myJPQ/gqlgen-todos/graph/model"
)

func (r *mutationResolver) UpdateMeeting(ctx context.Context, input []*model.NewTodo) ([]*model.Todo, error) {
	if r.todos == nil {
		r.todos = make(map[int]*model.Todo)
	}

	response := make([]*model.Todo, len(input))
	for i := 0; i < len(input); i++ {
		todo := &model.Todo{
			ID:      input[i].ID,
			Meeting: &model.Meeting{Title: input[i].Title, Description: input[i].Description, StartTime: input[i].StartTime},
		}
		r.todos[input[i].ID] = todo
		response[i] = todo
	}

	return response, nil
}

func (r *mutationResolver) DeleteMeeting(ctx context.Context, input []int) ([]*model.Todo, error) {
	for i := 0; i < len(input); i++ {
		if r.todos[input[i]] != nil {

			delete(r.todos, input[i])

		}
	}
	response := make([]*model.Todo, len(r.todos))
	i := 0
	for _, value := range r.todos {
		response[i] = value
		i++
	}

	return response, nil
}

func (r *queryResolver) GetMeeting(ctx context.Context) ([]*model.Todo, error) {
	response := make([]*model.Todo, len(r.todos))
	i := 0
	for _, value := range r.todos {
		response[i] = value
		i++
	}

	return response, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Getmeeting(ctx context.Context, id int) (*model.Todo, error) {
	return r.todos[id], nil
}
