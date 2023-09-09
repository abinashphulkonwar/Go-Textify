package handlers

import (
	"os"

	errorhandlers "github.com/abinashphulkonwar/Textify/error-handler"
)

func OutputTextFile(text []TextComponent, output_file_path string) {
	var data []byte
	path := output_file_path + ".txt"
	data, err := os.ReadFile(path)

	if os.IsNotExist(err) {
		println("createing a new file")
	} else {
		errorhandlers.HandleError(err)
	}

	string_text := ""
	string_text = "\n\nnew page\n\n"
	for _, val := range text {
		string_text = string_text + "\n" + val.Text
	}
	data = append(data, []byte(string_text)...)

	err = os.WriteFile(path, data, 0644)
	errorhandlers.HandleError(err)
}
