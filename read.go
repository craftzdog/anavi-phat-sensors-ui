package main

import (
  "log"
  "fmt"
  "context"
  firebase "firebase.google.com/go"
  firestore "cloud.google.com/go/firestore"
  "google.golang.org/api/iterator"
  "google.golang.org/api/option"
)

func main () {
  ctx := context.Background()
  opt := option.WithCredentialsFile("./service-account-key.json")
  app, err := firebase.NewApp(ctx, nil, opt)
  client, err := app.Firestore(ctx)
  if err != nil {
    log.Fatal(err)
  }

  iter := client.Collection("conditions").OrderBy("createdAt", firestore.Desc).Limit(100).Documents(ctx)
  for {
    doc, err := iter.Next()
    if err == iterator.Done {
      break
    }
    if err != nil {
      log.Fatalf("Failed to iterate: %v", err)
    }
    fmt.Println(doc.Data())
  }
}
