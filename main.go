package main

import (
	"context"
	"flag"

	"cloud.google.com/go/storage"

	"google.golang.org/api/option"

	"github.com/k0kubun/pp"
)

func main() {
	var err error
	var givenProjectId = flag.String("project_id", "", "project_id")
	var givenBucketName = flag.String("bucket_name", "", "bucket_name")
	var givenServiceAccountFile = flag.String("service_account_credential_file", "", "service_account_credential_file")

	flag.Parse()
	pp.Println("service_account_credential_file:" + *givenServiceAccountFile)

	if givenProjectId == nil || *givenProjectId == "" {
		pp.Println("no project_id!")
		return
	}
	projectID := *givenProjectId

	if givenBucketName == nil || *givenBucketName == "" {
		pp.Println("no bucket_name!")
		return
	}
	bucketName := *givenBucketName

	if givenServiceAccountFile == nil || *givenServiceAccountFile == "" {
		pp.Println("no serviceAccountFile!")
		return
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithServiceAccountFile(*givenServiceAccountFile))
	if err != nil {
		pp.Printf("Failed to create client: %v\n", err)
	}

	bucketAttrs := &storage.BucketAttrs{
		Location:     "asia-east1",
		StorageClass: "REGIONAL",
	}
	err = client.Bucket(bucketName).Create(ctx, projectID, bucketAttrs)
	if err != nil {
		pp.Printf("Failed to create bucket: %v\n", err)
	}

	pp.Printf("Bucket %v created.\n", bucketName)
}
