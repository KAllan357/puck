package puck

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func RunChefApply(input string) {
	cmd := exec.Command("chef-apply", "-s")
	cmd.Stdin = strings.NewReader(input)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("got an error")
		fmt.Println(err)
	}
	fmt.Println(out.String())
}
