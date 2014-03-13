package main

import (
	"github.com/KAllan357/puck"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/resources", resourcesHandler)
	http.ListenAndServe(":8080", nil)
}

func resourcesHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	resources := puck.ParseResourcesJSON(body)

	resourceStrings := puck.CompileChefResources(resources.Resources)
	chefApplyOutput := puck.RunChefApply(resourceStrings)

	w.Header().Set("Content-Type", "application/json")
	w.Write(chefApplyOutput)
}
