package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("aplay", "-D", "plughw:1,0", "out.wav").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
