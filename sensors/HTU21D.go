package sensors
// Temperature and Humidity sensor

import (
	"log"
	"os/exec"
  "strings"
  "strconv"
)

func GetTempAndHumid() (temperature float64, humidity float64, err error) {
	out, err := exec.Command("/home/pi/anavi-examples/sensors/HTU21D/c/HTU21D").Output()
	if err != nil {
    log.Printf("Failed to get sensor value from HTU21D:  %s\n", err)
    return
	}

  s := string(out[:len(out)])
  lines := strings.Split(s, "\n")
  lineTemp := lines[1]
  lineHumid := lines[2]

  tempStr := strings.Split(lineTemp, ": ")[1]
  temperature, err = strconv.ParseFloat(tempStr[0:len(tempStr)-1], 32)
  humidStr := strings.Split(lineHumid, ": ")[1]
  humidity, err = strconv.ParseFloat(humidStr[0:len(humidStr)-3], 32)

  return
}

