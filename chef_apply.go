package puck

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type ChefApplyOutput struct {
	Output []string
}

func RunChefApply(input string) []byte {
	cmd := exec.Command("chef-apply", "-s")
	cmd.Stdin = strings.NewReader(input)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("got an error")
		fmt.Println(err)
	}

	chefApplyOutput := ChefApplyOutput{Output: strings.Split(out.String(), "\n")}
	b, err := json.Marshal(chefApplyOutput)
	return b
}
