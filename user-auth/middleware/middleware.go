package middleware

import (
	"fmt"
	"net/http"
	"user-auth/models"
)

type RequireUser struct {
	*models.UserService
}

func (mw *RequireUser) Applyfn(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Redirect(w, r, "/home", http.StatusFound)
			return
		}
		user, err := mw.ByRememberToken(cookie.Value)

		if err != nil {
			http.Redirect(w, r, "/home", http.StatusFound)
			return
		}
		fmt.Println("User Found", user)
		next(w, r)

	})

}
