package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"user-auth/models"
	"user-auth/rand"
)

type UsersController struct {
	us *models.UserService
}

func NewUserController(us *models.UserService) *UsersController {
	return &UsersController{us: us}
}

type jsonData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
	panic("hello")
}

func (uc *UsersController) Create(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var u jsonData
	json.Unmarshal(b, &u)

	user := models.User{
		Email:    u.Email,
		Password: u.Password,
	}

	err = uc.us.Create(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = uc.signIn(w, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Successfully created your account")
}

func (uc *UsersController) Login(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var userDetails models.User

	json.Unmarshal(b, &userDetails)

	user, err := uc.us.Authenticate(userDetails.Email, userDetails.Password)
	_ = user
	if err != nil {
		switch err {
		case models.RecordNotFound:
			http.Error(w, "Invalid email address", http.StatusInternalServerError)
		case models.ErrInvalidPassword:
			http.Error(w, "Invalid Password", http.StatusInternalServerError)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	err = uc.signIn(w, &userDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "logged in")

}

func (uc *UsersController) signIn(w http.ResponseWriter, user *models.User) error {

	token, err := rand.RememberToken()
	if err != nil {
		return err
	}
	user.Remember = token

	err = uc.us.Update(user)

	if err != nil {
		return err
	}

	cookie := http.Cookie{
		Name:  "token",
		Value: user.Remember,
	}

	http.SetCookie(w, &cookie)
	return nil

}

func (uc *UsersController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page")
}

func (uc *UsersController) Hello(w http.ResponseWriter, r *http.Request) {

	cookie, _ := r.Cookie("token")
	user, _ := uc.us.ByRememberToken(cookie.Value)
	fmt.Fprintln(w, "Hello", user.Email)

}
