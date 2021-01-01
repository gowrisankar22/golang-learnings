package controller

import (
	"encoding/json"
	"fmt"
	"github-proj/model"
	"github-proj/service"
	"io/ioutil"
	"net/http"
)

func CreateRepo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var repo model.CreateRepoRequest
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(b, &repo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, appErr := service.RepoService.CreateRepo(repo)

	if appErr != nil {
		data, _ := json.Marshal(appErr)
		w.WriteHeader(appErr.StatusCode)
		w.Write(data)
		return
	}

	fmt.Fprintln(w, "Repo Created")
	data, _ := json.Marshal(resp)
	w.Write(data)

}
