package memezis

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func asJson(handler func(ctx context.Context, req *http.Request) (interface{}, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := handler(context.Background(), r)
		w.Header().Add("Content-Type", "application/json")
		if err != nil {
			httpErr, ok := err.(clientError)
			if !ok {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(httpErr.Code())
			json.NewEncoder(w).Encode(&errorResponse{
				Message: err.Error(),
			})
			log.Println("http handler error: ", err)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&resp)
	}
}