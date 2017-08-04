package store

import (
	"fmt"
	"io"
	"time"

	minio "github.com/minio/minio-go"
)

// MinioStore for save scores
type MinioStore struct {
	bucketName string
	client     *minio.Client
}

func (store *MinioStore) initMonio() (*minio.Client, error) {
	endpoint := "s3.meimor.com"
	accessKeyID := "Gf0328FY89GHF32J4H9F"
	secretAccessKey := "8fh6Lf56ieGf03jfSkjf28r3FY/89Go4aF3EWpf9"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		return nil, err
	}

	bucketName := store.bucketName
	location := "us-east-1"

	exists, err := minioClient.BucketExists(bucketName)
	if err == nil && !exists {
		err = minioClient.MakeBucket(bucketName, location)
		if err != nil {
			return nil, err
		}
	}

	return minioClient, nil
}

// Save save
func (store *MinioStore) Save(name string, reader io.Reader, contentType string) (string, error) {

	now := time.Now()

	objectName := fmt.Sprintf("%s/%s", now.Format("2006/01/02"), name)

	_, err := store.client.PutObject(store.bucketName, objectName, reader, contentType)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("http://s3.meimor.com/%s/%s", store.bucketName, objectName), nil
}

func newMinioStore() (*MinioStore, error) {
	store := &MinioStore{bucketName: "scores"}

	client, err := store.initMonio()
	if err != nil {
		return nil, err
	}

	store.client = client

	return store, nil
}
