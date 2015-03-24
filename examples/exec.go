package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	words := strings.Fields("-a -l -h")
	out, err := exec.Command("ls", words...).Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	out, err = exec.Command("aplay", "-D", "plughw:1,0", "out.wav").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
