package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/ilmsg/eloop-api/pkg/model"
)

var collName = "eloop-todo"

type TodoRepo struct {
	db *firestore.Client
}

func NewTodoRepo(db *firestore.Client) *TodoRepo {
	return &TodoRepo{db}
}

func (t *TodoRepo) GetTodos() ([]*model.Todo, error) {
	context := context.Background()
	docs, err := t.db.Collection(collName).Documents(context).GetAll()
	if err != nil {
		return nil, err
	}

	var todos []*model.Todo
	for _, doc := range docs {
		var todo *model.Todo
		err := doc.DataTo(&todo)
		if err != nil {
			return nil, err
		}
		todo.ID = doc.Ref.ID
		todos = append(todos, todo)
	}

	return todos, nil
}

func (t *TodoRepo) GetTodo(id string) (*model.Todo, error) {
	context := context.Background()
	doc, err := t.db.Collection(collName).Doc(id).Get(context)
	if err != nil {
		return nil, err
	}

	var todo *model.Todo
	err = doc.DataTo(&todo)
	if err != nil {
		return nil, err
	}

	todo.ID = doc.Ref.ID
	return todo, nil
}

func (t *TodoRepo) UpdateTodo(id string, todo *model.Todo) error {
	context := context.Background()

	updateTodo := map[string]interface{}{
		"title": todo.Title,
	}

	_, err := t.db.Collection(collName).Doc(id).Set(context, updateTodo, firestore.MergeAll)
	if err != nil {
		return err
	}

	return nil
}

func (t *TodoRepo) CreateTodo(todo *model.Todo) (string, error) {
	context := context.Background()
	doc, _, err := t.db.Collection(collName).Add(context, todo)
	if err != nil {
		return "", err
	}
	return doc.ID, nil
}

func (t *TodoRepo) DeleteTodo(id string) error {
	context := context.Background()
	_, err := t.db.Collection(collName).Doc(id).Delete(context)
	if err != nil {
		return err
	}
	return nil
}
