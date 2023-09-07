package handlers

import (
	"context"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

const URL = "https://mastersimageandtext.blob.core.windows.net/"
const CONTAINER_NAME = "peace-and-conflicts-resolution"

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
func Upload(path string) string {

	data, err := os.ReadFile(path)
	handleError(err)
	println(string(data))

	ctx := context.Background()

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	client, err := azblob.NewClient(URL, credential, nil)
	handleError(err)

	println(client.URL())

	_, err = client.UploadBuffer(ctx, CONTAINER_NAME, "blobName", data, &azblob.UploadBufferOptions{})
	handleError(err)

	//	client.ServiceClient().GetSASURL(CONTAINER_NAME, "blobName")

	return URL + CONTAINER_NAME + "/" + "blobName"
}
