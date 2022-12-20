package firebase

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

func NewFirestore() *firestore.Client {
	ctx := context.Background()
	projectID := os.Getenv("PROJECT_ID")
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return client
}
