package main

import (
	"fmt"
	"github.com/KAllan357/puck"
	"io/ioutil"
	"net/http"
	// "strings"
)

func main() {
	http.HandleFunc("/resources", resourcesHandler)
	http.ListenAndServe(":8080", nil)
}

func resourcesHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	resources := puck.ParseResourcesJSON(body)

	resourceStrings := puck.CompileChefResources(resources.Resources)
	puck.RunChefApply(resourceStrings)

	fmt.Fprint(w, string(body))
}
