package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"os"

	errorhandlers "github.com/abinashphulkonwar/Textify/error-handler"
)

type TextComponent struct {
	Text string
	Box  string
}

type JsonSchema struct {
	List []TextComponent `json:"list"`
}

func GenerateHash(key string) string {
	hash := md5.Sum([]byte(key))
	hashString := hex.EncodeToString(hash[:])
	return hashString
}

func OutputJson(textList []TextComponent, path string) {
	fileName := GenerateHash(path) + ".json"

	println(fileName)
	var data []byte
	data, err := os.ReadFile(fileName)

	if os.IsNotExist(err) {
		err = os.WriteFile(fileName, data, 0644)
		errorhandlers.HandleError(err)
	} else {
		errorhandlers.HandleError(err)
	}

	println(data)

	fileDataJson := JsonSchema{}
	if len(data) > 0 {
		err = json.Unmarshal(data, &fileDataJson)
		errorhandlers.HandleError(err)
	}

	fileDataJson.List = append(fileDataJson.List, textList...)
	println("fileDataJson ")
	fileDataByts, err := json.Marshal(fileDataJson)
	errorhandlers.HandleError(err)
	println("fileDataByts ")

	err = os.WriteFile(fileName, fileDataByts, 0644)
	errorhandlers.HandleError(err)

}
