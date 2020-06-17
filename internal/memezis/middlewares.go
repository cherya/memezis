package memezis

import (
	"context"
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

type Client struct {
	Name string
}

type authMiddleware struct {
	clients map[string]*Client
}

func newAuthMiddleware(clients map[string]*Client) func(next http.Handler) http.Handler {
	amw := &authMiddleware{
		clients: clients,
	}
	return amw.Middleware
}

func (amw *authMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if client, found := amw.clients[token]; found {
			// add client to context
			ctx := context.WithValue(r.Context(), "client", client)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			// Write an error and stop the handler chain
			http.Error(w, "access denied", http.StatusForbidden)
		}
	})
}
