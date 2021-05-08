package main

import (
  "log"
  "time"
  "context"
  firebase "firebase.google.com/go"
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
  collection := client.Collection("conditions")

  _, _, err = collection.Add(ctx, map[string]interface{}{
    "createdAt": time.Now(),
    "humidity":  1,
    "light":  1815,
    "pressure": 10,
    "temperature": 100,
  })
  if err != nil {
    log.Fatalf("Failed adding alovelace: %v", err)
  }
}
