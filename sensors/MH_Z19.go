package sensors
// Temperature and Humidity sensor

import (
	"log"
	"os/exec"
  "encoding/json"
)

func GetCo2() (co2 float64, err error) {
	out, err := exec.Command("/usr/bin/python", "-m", "mh_z19").Output()
	if err != nil {
    log.Printf("Failed to get sensor value from MH_Z19:  %s\n", err)
    return
	}

  s := string(out[:len(out)])
  var data map[string]interface{}
  err = json.Unmarshal([]byte(s), &data)

	if err == nil {
    co2 = data["co2"].(float64)
  }

  return
}
