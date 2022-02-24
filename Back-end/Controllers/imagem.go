package Controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//FUNCTION FOR TESTS
func ImagemPost(w http.ResponseWriter, r *http.Request) {
	//parse input from a form data  multipart/form-data
	//colocar entre 10mb a 20 mb o limite
	r.ParseMultipartForm(10 << 20)
	//sacar o file
	file, handler, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		fmt.Println(err)
	}
	defer file.Close()
	// Get substring after a string.
	pos := strings.LastIndex(handler.Filename, ".")
	adjustedPos := pos + len(".")
	extension := handler.Filename[adjustedPos:len(handler.Filename)]
	println(extension)
	fmt.Printf("UPLOADING FILE: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	//escrever temporariamente o file no nosso server
	if extension == "png" || extension == "jpeg" || extension == "jpg" || extension == "PNG" || extension == "JPEG" || extension == "JPG" {
		os.Remove("temp/upload1052450194.jpg")
		done := fmt.Sprintf("upload*.%s", extension)
		tempFile, err := ioutil.TempFile("temp", done)
		println(tempFile.Name())
		if err != nil {
			fmt.Println(err)
			return
		}
		defer tempFile.Close()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		println(fileBytes)
		tempFile.Write(fileBytes)

		//retornar com sucesso ou nao
		fmt.Fprintf(w, "Sucess\n")
	} else {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Formats error")

	}
}
