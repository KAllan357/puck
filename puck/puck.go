package main

import (
	// "bytes"
	"encoding/json"
	"fmt"
	"github.com/KAllan357/puck"
	"io/ioutil"
	"net/http"
	// "os/exec"
	// "strings"
)

func main() {
	http.HandleFunc("/resource", resourceHandler)
	http.ListenAndServe(":8080", nil)
}

func resourceHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var resources puck.Resources
	json.Unmarshal(body, &resources)
	fmt.Printf("%#v", resources.Resources)

	for _, v := range resources.Resources {
		fmt.Println(v.ToChefResource())
	}

	fmt.Fprint(w, string(body))
}

// func runChefApply(input string) {
// 	cmd := exec.Command("chef-apply", "-s")
// 	cmd.Stdin = strings.NewReader(input)
// 	var out bytes.Buffer
// 	cmd.Stdout = &out
// 	err := cmd.Run()
// 	if err != nil {
// 		fmt.Println("got an error")
// 		fmt.Println(err)
// 	}
// 	fmt.Println(out.String())
// }
