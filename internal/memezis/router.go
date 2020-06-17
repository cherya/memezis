package memezis

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (m *memezis) createRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/queue/{queue_name}/info", asJson(m.GetQueueInfo)).Methods(http.MethodGet)
	r.HandleFunc("/post/add", asJson(m.AddPost)).Methods(http.MethodPost)
	r.HandleFunc("/post/random", asJson(m.GetRandomPost)).Methods(http.MethodGet)
	r.HandleFunc("/post/{post_id:[0-9]+}", asJson(m.GetPostByID)).Methods(http.MethodGet)
	r.HandleFunc("/post/{post_id:[0-9]+}/upvote", asJson(m.UpVote)).Methods(http.MethodPost)
	r.HandleFunc("/post/{post_id:[0-9]+}/downvote", asJson(m.DownVote)).Methods(http.MethodPost)
	r.HandleFunc("/post/{post_id:[0-9]+}/publish", asJson(m.PublishPost)).Methods(http.MethodPost)
	r.HandleFunc("/upload", asJson(m.UploadMedia)).Methods(http.MethodPost)

	return r
}
