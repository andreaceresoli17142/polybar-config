package main

import (
	"fmt"
	"math"
	"os/exec"
	"strconv"
	"strings"
)

func main() {

	out, err := exec.Command("speedtest-cli", "--csv").Output()

	if err != nil {
		fmt.Print("↓ - Mb/s  ↑ - Mb/s")
		return
	}

	splitted := strings.Split(string(out), ",")

	down, _ := strconv.ParseFloat(splitted[6], 64)

	down = math.Round(down/10000) / 100

	up, _ := strconv.ParseFloat(splitted[7], 64)

	up = math.Round(up/10000) / 100

	/*up := 12.11

	down := 8.11*/

	fmt.Printf("↓ %v Mb/s  ↑ %v Mb/s", up, down)
}
