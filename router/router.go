package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kurmidev/webox/handler"
)

var mySigningKey = []byte(os.Getenv("JWT_SECRET"))

func Routes(h *handler.Handlers) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	//r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/user", func(r chi.Router) {
			r.Post("/login", h.Login)
			r.Post("/send-login-otp", h.SendOtp)
			r.Post("/login-otp", h.LoginOtp)
		})

		r.Get("/bouque/list", h.BouqueList)
		r.Get("/bouque/{id}", h.Bouque)

		r.Group(func(r chi.Router) {
			r.Route("/subscriber", func(r chi.Router) {
				r.Get("/profile/{id}", h.GetProfile)
				// r.Get("/me/", h.MeDetails)
				r.Get("/smc-details/{number}/{type}", h.SmcDetails)
				r.Get("/transaction/{id}", h.ToTransactionResponse)
			})
			//r.Use(AuthMiddleware) // Use the authentication middleware

		})

	})
	return r
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return mySigningKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
