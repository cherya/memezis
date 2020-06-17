package filestore

import (
	"fmt"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gofrs/uuid"
)

type store struct {
	baseURL    string
	session    *session.Session
	s3Endpoint string
	s3Key      string
	s3Secret   string
	s3Region   string
	s3Bucket   string
}

func NewStore(baseURL, s3Endpoint, s3Key, s3Secret, s3Region, s3Bucket string) *store {
	return &store{
		s3Endpoint: s3Endpoint,
		baseURL:    baseURL,
		s3Key:      s3Key,
		s3Secret:   s3Secret,
		s3Region:   s3Region,
		s3Bucket:   s3Bucket,
	}
}

func (s *store) GetObjAbsoluteURL(objKey string) string {
	return fmt.Sprintf("%s/%s", s.baseURL, objKey)
}

const (
	tempPrefix = "temp/"
)

func normalizeFilename(fn string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9.]+")
	fn = reg.ReplaceAllString(fn, "")
	if len(fn) > 15 {
		fn = fn[len(fn)-15:]
	}
	u4, _ := uuid.NewV4()
	fn = fmt.Sprintf("%s_%s", u4.String()[:6], fn)
	return fn
}

func (s *store) getSession() *session.Session {
	if s.session != nil {
		return s.session
	}
	s.session = session.Must(session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(s.s3Key, s.s3Secret, ""),
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String(s.s3Endpoint),
		Region:           aws.String(s.s3Region),
	}))
	return s.session
}

func (s *store) getTempObjKey(key string) string {
	return tempPrefix + key
}
