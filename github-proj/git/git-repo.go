package git

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github-proj/model"
	"github-proj/utils"
	"io/ioutil"
	"log"
	"net/http"
)

const url = "https://api.github.com/user/repos"

func getAuthToken(accessToken string) string {
	return fmt.Sprintf("token %s", accessToken)
}

func CreateRepo(accessToken string, repo model.CreateRepoRequest) (*model.CreateRepoResponse, *utils.GithubErrorResponse) {

	headers := http.Header{}
	headers.Set("Authorization", getAuthToken(accessToken))

	bytesData, _ := json.Marshal(repo)

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bytesData))

	if err != nil {
		log.Println("Error in creating repo")
		return nil, &utils.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	request.Header = headers
	client := http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return nil, &utils.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	bytesData, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, &utils.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		var err utils.GithubErrorResponse
		json.Unmarshal(bytesData, &err)
		err.StatusCode = resp.StatusCode
		return nil, &err
	}

	var result model.CreateRepoResponse

	err = json.Unmarshal(bytesData, &result)

	if err != nil {
		log.Println("Error while unmarshalling")
		return nil, &utils.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return &result, nil

}
