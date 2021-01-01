package service

import (
	"errors"
	"github-proj/git"
	"github-proj/model"
	"github-proj/utils"
	"net/http"
	"os"
)

type repoServiceInterface interface {
	CreateRepo(request model.CreateRepoRequest) (*model.CreateRepoResponse, *utils.GithubErrorResponse)
}

var RepoService repoServiceInterface = &repoService{}

type repoService struct{}

var invalidName = errors.New("invalid name")

func (r *repoService) CreateRepo(request model.CreateRepoRequest) (*model.CreateRepoResponse, *utils.GithubErrorResponse) {

	if request.Name == "" {
		return nil, &utils.GithubErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    invalidName.Error(),
		}
	}
	resp, appErr := git.CreateRepo(os.Getenv("AccessToken"), request)

	if appErr != nil {
		return nil, appErr
	}

	return resp, nil

}
