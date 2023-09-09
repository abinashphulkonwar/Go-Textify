package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.0/computervision"
	"github.com/Azure/go-autorest/autorest"
	errorhandlers "github.com/abinashphulkonwar/Textify/error-handler"
)

var computerVisionContext context.Context

type api_data_schema struct {
	Key string `json:"api_key"`
	Url string `json:"api_url"`
}

func GetApiData() api_data_schema {

	data, err := os.ReadFile("D:/development/web__apps/golang/text-extractor-azure/config.json")
	errorhandlers.HandleError(err)
	api_data := api_data_schema{}

	err = json.Unmarshal(data, &api_data)
	errorhandlers.HandleError(err)

	if api_data.Key == "" {
		errorhandlers.HandleError(errors.New("api key not found"))
	}
	if api_data.Url == "" {
		errorhandlers.HandleError(errors.New("api url not found"))
	}

	return api_data
}

func ExtractText(data io.ReadCloser) []TextComponent {
	api_data := GetApiData()
	computerVisionKey := api_data.Key
	endpointURL := api_data.Url

	computerVisionClient := computervision.New(endpointURL)

	computerVisionClient.Authorizer = autorest.NewCognitiveServicesAuthorizer(computerVisionKey)

	computerVisionContext = context.Background()

	ocrResult, err := computerVisionClient.RecognizePrintedTextInStream(computerVisionContext, true, data, computervision.En)
	errorhandlers.HandleError(err)

	fmt.Printf("Text angle: %.4f\n", *ocrResult.TextAngle)
	textList := []TextComponent{}
	// Get bounding boxes for each line of text and print text.
	for _, region := range *ocrResult.Regions {
		for _, line := range *region.Lines {
			fmt.Printf("\nBounding box: %v\n", *line.BoundingBox)
			s := ""
			for _, word := range *line.Words {
				s += *word.Text + " "
			}
			textList = append(textList, TextComponent{
				Text: s,
				Box:  *line.BoundingBox,
			})
			fmt.Printf("Text: %v", s)
		}
	}
	return textList
}
