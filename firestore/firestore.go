package main

import (
	"bkirk/utilities"
	"context"
	firebase "firebase.google.com/go"
	"github.com/davecgh/go-spew/spew"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("/Users/billykirk/Downloads/wkirk01-testdata-8b5908301c93.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	utilities.Check(err)
	client, err := app.Firestore(ctx)
	utilities.Check(err)
	defer client.Close()

	iter := client.Collection("posts").Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		utilities.Check(err)
		spew.Dump(doc.Data())
	}
}
