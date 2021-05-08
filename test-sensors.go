package main

import (
  "./sensors"
  "fmt"
)

func main () {
  i := 0
  temperature1, humidity, err := sensors.GetTempAndHumid()
  for err != nil && i < 10 {
    temperature1, humidity, err = sensors.GetTempAndHumid()
    i++
  }
  fmt.Printf("Temperature:\t%f C\n", temperature1)
  fmt.Printf("Humidity:\t%f %%rh\n", humidity)

  i = 0
  light, err := sensors.GetLight()
  for err != nil && i < 10 {
    light, err = sensors.GetLight()
    i++
  }
  fmt.Printf("Light:\t%f Lux\n", light)

  i = 0
  temperature2, pressure, err := sensors.GetTempAndPressure()
  for err != nil && i < 10 {
    temperature2, pressure, err = sensors.GetTempAndPressure()
    i++
  }
  fmt.Printf("Temperature:\t%f C\n", temperature2)
  fmt.Printf("Pressure:\t%f %%rh\n", pressure)

  co2, _ := sensors.GetCo2()
  fmt.Printf("Co2:\t%f ppm\n", co2)
}
