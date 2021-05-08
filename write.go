package main

import (
  "./sensors"
  "log"
  "time"
  "context"
  firebase "firebase.google.com/go"
  "google.golang.org/api/option"
)

type RoomData struct {
  temperature float64
  pressure float64
  humidity float64
  light float64
  co2 float64
}

func recordData (roomData *RoomData) {
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
    "humidity":  roomData.humidity,
    "light":  roomData.light,
    "pressure": roomData.pressure,
    "temperature": roomData.temperature,
    "co2": roomData.co2,
  })
  if err != nil {
    log.Fatalf("Failed adding alovelace: %v", err)
  }
}

func main () {
  log.Println("Start capturing my room confitions")
  i := 0

  temperature1, humidity, err := sensors.GetTempAndHumid()
  for err != nil && i < 10 {
    temperature1, humidity, err = sensors.GetTempAndHumid()
    i++
  }
  if (err == nil) {
    log.Printf("Temperature:  %f C\n", temperature1)
    log.Printf("Humidity:     %f %%rh\n", humidity)
  }

  i = 0
  light, err := sensors.GetLight()
  for err != nil && i < 10 {
    light, err = sensors.GetLight()
    i++
  }
  if (err == nil) {
    log.Printf("Light:        %f Lux\n", light)
  }

  i = 0
  temperature2, pressure, err := sensors.GetTempAndPressure()
  for err != nil && i < 10 {
    temperature2, pressure, err = sensors.GetTempAndPressure()
    i++
  }
  if (err == nil) {
    log.Printf("Temperature:  %f C\n", temperature2)
    log.Printf("Pressure:     %f %%rh\n", pressure)
  }

  co2, err := sensors.GetCo2()
  if (err == nil) {
    log.Printf("Co2:\t%f ppm\n", co2)
  }

  data := RoomData{
    temperature: temperature1,
    pressure: pressure,
    humidity: humidity,
    light: light,
    co2: co2,
  }

  recordData(&data)
  log.Println("Finished recording data!")
}
