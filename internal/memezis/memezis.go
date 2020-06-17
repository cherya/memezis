package memezis

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cherya/memezis/internal/store"
	"github.com/cherya/memezis/pkg/queue"
	"github.com/gocraft/work"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type DataStorageManager interface {
	AddPost(ctx context.Context, media []*store.Media, tags []string, createdAt time.Time, source, submittedBy, text string) (*store.Post, error)
	GetPostByID(ctx context.Context, postID int64) (*store.Post, error)
	GetRandomPost(ctx context.Context) (*store.Post, error)
	EnqueuePost(ctx context.Context, postID int64, publishedAt time.Time, to string) error
	UpVote(ctx context.Context, postID int, userID string) (*store.VotesCount, error)
	DownVote(ctx context.Context, postID int, userID string) (*store.VotesCount, error)
	GetTagsByIDs(ctx context.Context, tagsIDs []int64) ([]string, error)
	CheckDuplicate(ctx context.Context, postID int64) ([]int64, error)
}

type FileManager interface {
	Upload(f *bytes.Buffer, filename string) (string, error)
	UploadTemp(f *bytes.Buffer, filename string) (string, error)
	IsObjExists(objKey string) bool
	IsTempObjExists(objKey string) bool
	MakeObjPermanent(objName string) error
	GetObjAbsoluteURL(objKey string) string
}

type QueueManager interface {
	Push(queue string, postID int64) error
	PushWithDelay(queue string, delay time.Duration, postID int64) error
	QueueLength(queue string) (int64, error)
	QueueLastJobTime(queue string) (time.Time, error)
	GetQueueTimeout(queue string) (time.Duration, error)
	ConsumeWithDelay(queue string, handler func(job *work.Job) error)
}

type memezis struct {
	store  DataStorageManager
	qm     QueueManager
	fs     FileManager
	router *mux.Router
}

func NewMemezis(store DataStorageManager, qm QueueManager, fs FileManager, clients map[string]*Client) *memezis {
	m := &memezis{
		store: store,
		qm:    qm,
		fs:    fs,
	}
	m.router = m.createRouter()
	m.router.Use(loggingMiddleware, newAuthMiddleware(clients))

	m.qm.ConsumeWithDelay(queue.EverythingQueue, func(job *work.Job) error {
		fmt.Println("Processing job:", queue.EverythingQueue, job.Name, job.Args)
		return nil
	})

	return m
}

func (m *memezis) Run(port int64) error {
	srv := &http.Server{
		Handler: m.router,
		Addr:    fmt.Sprintf(":%d", port),
	}
	fmt.Println(`
░░░░░▄▀▀▀▄░░░░░░░░░░░░░░░░░
▄███▀░◐░░░▌░░░░░░░░░░░░░░░░
░░░░▌░░░░░▐░░░░░░░░░░░░░░░░
░░░░▐░░░░░▐░░░░░░░░░░░░░░░░
░░░░▌░░░░░▐▄▄░░░░░░░░░░░░░░
░░░░▌░░░░▄▀▒▒▀▀▀▀▄░░░░░░░░░
░░░▐░░░░▐▒▒▒▒▒▒▒▒▀▀▄░░░░░░░
░░░▐░░░░▐▄▒▒▒▒▒▒▒▒▒▒▀▄░░░░░
░░░░▀▄░░░░▀▄▒▒▒▒▒▒▒▒▒▒▀▄░░░
░░░░░░▀▄▄▄▄▄█▄▄▄▄▄▄▄▄▄▄▄▀▄░
░░░░░░░░░░░▌▌░▌▌░░░░░░░░░░░
░░░░░░░░░░░▌▌░▌▌░░░░░░░░░░░
░░░░░░░░░▄▄▌▌▄▌▌░░░░░░░░░░░
	`)
	fmt.Printf("app running on port %d...\n", port)
	log.Fatal(srv.ListenAndServe())

	return nil
}
