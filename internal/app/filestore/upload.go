package filestore

import (
	"bytes"
	"github.com/pkg/errors"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (s *store) Upload(file *bytes.Buffer, filename string) (string, error) {
	filename = normalizeFilename(filename)
	err := s.upload(file, filename, s.s3Bucket, "attachment", s3.BucketCannedACLPublicRead, nil)
	if err != nil {
		err = errors.Wrap(err, "Upload: failed to upload file ")
		return "", err
	}

	return filename, nil
}

func (s *store) UploadTemp(file *bytes.Buffer, filename string) (string, error) {
	filename = normalizeFilename(filename)
	err := s.upload(file, s.getTempObjKey(filename), s.s3Bucket, "temp_upload", "", nil)
	if err != nil {
		return "", errors.Wrap(err, "UploadTemp: failed to upload temp file ")
	}

	return filename, nil
}

func (s *store) upload(file *bytes.Buffer, filename, bucketName, disposition, acl string, expires *time.Time) error {
	sess := s.getSession()
	r := bytes.NewReader(file.Bytes())
	s3client := s3.New(sess)
	req := &s3.PutObjectInput{
		Bucket:             aws.String(bucketName),
		Key:                aws.String(filename),
		ACL:                aws.String(acl),
		Body:               r,
		ContentLength:      aws.Int64(r.Size()),
		ContentType:        aws.String(http.DetectContentType(file.Bytes())),
		ContentDisposition: aws.String(disposition),
	}
	if expires != nil {
		req.Expires = expires
	}
	_, err := s3client.PutObject(req)
	if err != nil {
		return errors.Wrap(err, "upload: can't put object")
	}

	return err
}
