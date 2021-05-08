package sensors
// Temperature and Pressure sensor

import (
	"log"
	"os/exec"
  "strings"
  "strconv"
)


func GetTempAndPressure() (temperature float64, pressure float64, err error) {
	out, err := exec.Command("/home/pi/anavi-examples/sensors/BMP180/c/BMP180").Output()
	if err != nil {
    log.Printf("Failed to get sensor value from BMP180:  %s\n", err)
    return
	}

  s := string(out[:len(out)])
  lines := strings.Split(s, "\n")
  lineTemp := lines[1]
  linePress := lines[2]

  tempStr := strings.Split(lineTemp, ": ")[1]
  temperature, err = strconv.ParseFloat(tempStr[0:len(tempStr)-2], 32)
  pressStr := strings.Split(linePress, ": ")[1]
  pressure, err = strconv.ParseFloat(pressStr[0:len(pressStr)-4], 32)

  return
}

