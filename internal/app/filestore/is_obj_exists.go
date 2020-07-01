package filestore

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (s *store) isObjExistsInBucket(objKey string, bucketName string) bool {
	sess := s.getSession()
	s3client := s3.New(sess)

	_, err := s3client.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objKey),
	})

	if err != nil {
		return false
	}

	return true
}

func (s *store) IsObjExists(objKey string) bool {
	return s.isObjExistsInBucket(objKey, s.s3Bucket)
}

func (s *store) IsTempObjExists(objKey string) bool {
	return s.isObjExistsInBucket(tempPrefix+objKey, s.s3Bucket)
}
