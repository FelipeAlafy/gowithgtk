package urltools

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Download(url, path string) (*os.File, string, error) {
	file, filePath := createFile(path, url)
	err := putFile(file, httpClient(), url)
	if err != nil {
		log.Fatal("downloader - Download: ", err)
		return file, filePath, err
	}
	return file, filePath, err
}

func putFile(filePath *os.File, client *http.Client, url string) error {
	res, err := client.Get(url)
	defer res.Body.Close()
	defer filePath.Close()
	if err != nil {
		log.Fatal("downloader - putFile 1: ", err)
		return err
	}
	_, err = io.Copy(filePath, res.Body)
	if err != nil {
		log.Fatal("downloader - putFile 2: ", err)
		return err
	}
	return nil
}

func httpClient() *http.Client {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	log.Println("Client return was sucess")
	return &client
}

func fileName(urlPath string) string {
	url, err := url.Parse(urlPath)
	if err != nil {
		panic(err)
	}
	path := url.Path
	segments := strings.Split(path, "/")
	return segments[len(segments)-1]
}

func createFile(path, urlPath string) (*os.File, string) {
	log.Println(path)
	t := strings.SplitAfter(path, "file://")
	log.Println(t)

	nPath := t[1]
	nPath += "/" + fileName(urlPath)
	file, err := os.Create(nPath)
	if err != nil {
		log.Fatal("downloader - createFile: ", err, "\nPath: ", nPath)
		panic(err)
	}
	return file, nPath
}
