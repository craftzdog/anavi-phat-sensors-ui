package sensors
// Light sensor

import (
	"log"
	"os/exec"
  "strings"
  "strconv"
)

func GetLight() (float64, error) {
	out, err := exec.Command("/home/pi/anavi-examples/sensors/BH1750/c/BH1750").Output()
	if err != nil {
    log.Printf("Failed to get sensor value from BH1750:  %s\n", err)
    return 0, err
	}

  s := string(out[:len(out)])
  lines := strings.Split(s, "\n")
  lineLight := lines[1]

  lightStr := strings.Split(lineLight, ": ")[1]
  light, err := strconv.ParseFloat(lightStr[0:len(lightStr)-4], 32)
  return light, err
}

