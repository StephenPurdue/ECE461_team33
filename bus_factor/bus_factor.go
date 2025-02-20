package bus_factor

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// uses the cloned repository to determine the bus size 
func Get_minimum_bus_size(gitPath string) int {
	// analyze the cloned repository at gitPath
	cmd := exec.Command("python", "bus_factor.py", fmt.Sprintf("\"%s\"", gitPath))
	output, err := cmd.Output()

	if (err != nil) {
		return 0
	}

	// parse bus_size from python output
	outputLines := strings.Split(strings.TrimSpace(string(output)), "\n")
	i, parseError := strconv.Atoi(outputLines[len(outputLines) - 1])

	if (parseError != nil) {
		return 0
	}

	return i
}

// calculates a bus factor (between 0 and 1) from the bus size
func calculate_bus_factor(bus_size int) float32 {
	if (bus_size < 1) {
		return 0.0
	}

	return (float32(bus_size) - 1) / float32(bus_size)
}

// calculates the bus factor by cloning the repo locally then using the truckfactor pyhton library
func Get_bus_factor(githubUrl string) float32 {
	bus_size := Get_minimum_bus_size(githubUrl)
	return calculate_bus_factor(bus_size)
}
