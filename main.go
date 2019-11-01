/*
   This is a simple code example for connecting, uploading, downloading and listing files
   from an AWS S3 Bucket.
   Author: Antonio Sanchez antonio@asanchez.dev
*/

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var sess = connectAWS()

func connectAWS() *session.Session {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(AWS_S3_REGION)})
	if err != nil {
		panic(err)
	}
	return sess
}

const (
	AWS_S3_REGION = ""
	AWS_S3_BUCKET = ""
)

func main() {

	http.HandleFunc("/upload/", handlerUpload) // Upload
	http.HandleFunc("/get/", handlerDownload)  // Get the file
	http.HandleFunc("/list/", handlerList)     // List all files
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func showError(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, message)
}
