package handlers

import (
	"context"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/service"
	errorhandlers "github.com/abinashphulkonwar/Textify/error-handler"
)

const ACCOUNT_NAME = "mastersimageandtext"
const URL = "https://" + ACCOUNT_NAME + ".blob.core.windows.net/"
const CONTAINER_NAME = "peace-and-conflicts-resolution"

func Upload(path string) string {

	data, err := os.ReadFile(path)
	errorhandlers.HandleError(err)

	ctx := context.Background()

	credential, err := azidentity.NewDefaultAzureCredential(nil)

	errorhandlers.HandleError(err)

	client, err := azblob.NewClient(URL, credential, nil)
	errorhandlers.HandleError(err)
	println(client.URL())

	_, err = client.UploadBuffer(ctx, CONTAINER_NAME, "blobName", data, &azblob.UploadBufferOptions{})
	errorhandlers.HandleError(err)
	now := time.Now()
	exp := now.Add(time.Minute * 10)

	info := service.KeyInfo{
		Start:  to.Ptr(now.UTC().Format(sas.TimeFormat)),
		Expiry: to.Ptr(exp.UTC().Format(sas.TimeFormat)),
	}

	serviceClient := client.ServiceClient()

	udc, err := serviceClient.GetUserDelegationCredential(context.TODO(), info, nil)
	errorhandlers.HandleError(err)
	sasQueryParams, err := sas.BlobSignatureValues{
		Protocol:      sas.ProtocolHTTPS,
		StartTime:     time.Now().UTC().Add(time.Second * -10),
		ExpiryTime:    time.Now().UTC().Add(15 * time.Minute),
		Permissions:   to.Ptr(sas.ContainerPermissions{Read: true}).String(),
		ContainerName: CONTAINER_NAME,
	}.SignWithUserDelegation(udc)

	errorhandlers.HandleError(err)
	return URL + CONTAINER_NAME + "/" + "blobName" + "?" + sasQueryParams.Encode()
}
