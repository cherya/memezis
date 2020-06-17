package store

import (
	"github.com/lib/pq"
	"time"
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
	Media             []Media
	Votes             VotesCount
}

type Media struct {
	ID       int    `db:"id"`
	PostID   int64  `db:"post_id"`
	Key      string `db:"url"`
	Type     string `db:"type"`
	SourceID string `db:"source_id"`
	SHA1     string `db:"sha1"`
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