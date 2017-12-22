package philipshue

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Start() {
	philip := NewPhilipsHue("settings")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Command line reader started")
	for true {
		text, _ := reader.ReadString('\n')
		if !strings.Contains(text, "-") {
			fmt.Println("Incorrect command")
			fmt.Println("Try: {{lampname}}-{{state}}")
		} else {
			var result []string
			result = strings.Split(text, "-")
			if strings.Contains(result[1], "on") {
				philip.turnLightOn(result[0])
			} else if strings.Contains(result[1], "off") {
				philip.turnLightOff(result[0])
			} else if strings.Contains(result[1], "#") {
				philip.setLightColor(result[0], result[1])
			}
		}
	}
}
