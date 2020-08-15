package store

import (
	"time"

	"github.com/lib/pq"
)

type Post struct {
	ID                int64         `db:"id"`
	SubmittedBy       string        `db:"submitted_by"`
	Source            string        `db:"source"`
	Text              string        `db:"text"`
	CreatedAt         time.Time     `db:"created_at"`
	Tags              pq.Int64Array `db:"tags"`
	OriginalCreatedAt time.Time     `db:"original_created_at"`
	HasMedia          bool          `db:"has_media"`
	SourceURL         string        `db:"source_url"`
	Publish           []Publish
	Media             []Media
	Votes             VotesCount
}

type Media struct {
	ID       int    `db:"id"`
	PostID   int64  `db:"post_id"`
	Key      string `db:"url"`
	Type     string `db:"type"`
	SourceID string `db:"source_id"`
	Phash    string `db:"phash"`
}

type Publish struct {
	ID          int       `db:"id"`
	PostID      int64     `db:"post_id"`
	PublishedAt time.Time `db:"published_at"`
	PublishedTo string    `db:"published_to"`
	URL         string    `db:"url"`
	Status      string    `db:"status"`
}

type VotesCount struct {
	ID     int `db:"id"`
	PostID int `db:"post_id"`
	Up     int `db:"up"`
	Down   int `db:"down"`
}

type Tag struct {
	ID   int    `db:"id"`
	Text string `db:"text"`
}

type NewVote struct{}

type PublishStatus string

const (
	PublishStatusUnknown   PublishStatus = "unknown"
	PublishStatusPublished PublishStatus = "published"
	PublishStatusEnqueued  PublishStatus = "enqueued"
	PublishStatusDeclined  PublishStatus = "declined"
)
