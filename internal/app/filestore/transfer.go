package filestore

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
)

func (s *store) MakeObjPermanent(objKey string) error {
	sess := s.getSession()
	s3client := s3.New(sess)

	resp, err := s3client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.s3Bucket),
		Key:    aws.String(s.getTempObjKey(objKey)),
	})
	if err != nil {
		return errors.Wrapf(err, "MakeObjPermanent: can't get object (bucket=%s, key=%s)", s.s3Bucket, objKey)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "MakeObjPermanent: can't read resp body (bucket=%s, key=%s)", s.s3Bucket, objKey)
	}

	_, err = s3client.CopyObject(&s3.CopyObjectInput{
		Bucket:             aws.String(s.s3Bucket),
		CopySource:         aws.String(s.s3Bucket + "/" + s.getTempObjKey(objKey)),
		Key:                aws.String(objKey),
		ContentDisposition: aws.String("permanent"),
		ContentType:        aws.String(http.DetectContentType(body)),
		MetadataDirective:  aws.String(s3.MetadataDirectiveReplace),
	})
	if err != nil {
		return errors.Wrapf(err, "MakeObjPermanent: can't copy object (bucket=%s, key=%s)", s.s3Bucket, objKey)
	}

	go func() {
		_, err = s3client.DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String(s.s3Bucket),
			Key:    aws.String(s.getTempObjKey(objKey)),
		})
		if err != nil {
			log.Print(errors.Wrapf(err, "MakeObjPermanent: can't delete object (bucket=%s, key=%s)", s.s3Bucket, objKey))
		}
	}()

	return nil
}
