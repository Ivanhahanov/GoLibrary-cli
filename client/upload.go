package client

import (
	"bytes"
	"fmt"
	"github.com/Ivanhahanov/GoLibrary-cli/models"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func UploadData(filePath string, descriptionPath string) (err error) {

	description := ParseDescription(descriptionPath)
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("file", filePath)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	// open file handle
	fh, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}
	bodyWriter.WriteField("title", description.Title)
	bodyWriter.WriteField("author", description.Author)
	bodyWriter.WriteField("publisher", description.Publisher)
	bodyWriter.WriteField("description", description.BookDescription)
	bodyWriter.Close()

	req, err := http.NewRequest("PUT", url+"/books/", bodyBuf)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", bodyWriter.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}

func ParseDescription(filePath string) (description *models.Description) {

	filename, _ := filepath.Abs(filePath)
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &description)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Value: %#v\n", description)

	return
}