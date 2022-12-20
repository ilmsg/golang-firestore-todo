package model

type Todo struct {
	ID    string `json:"id" firestore:"id"`
	Title string `json:"title" firestore:"title"`
}
