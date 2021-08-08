package client

import (
	"encoding/json"
	"fmt"
	"github.com/Ivanhahanov/GoLibrary-cli/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

func GetBooks() {
	client := http.Client{Timeout: time.Second * 2}
	req, err := http.NewRequest(http.MethodGet, "http://localhost/books", nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// unmarshal response to Books struct
	var books models.Books
	jsonErr := json.Unmarshal(body, &books)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	// choose output len default: 10
	outputLen := 10
	if len(books.Books) < outputLen {
		outputLen = len(books.Books)
	}

	// print data
	w := tabwriter.NewWriter(os.Stdout, 5, 0, 2, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "TITLE\tAUTHOR\tTAGS\t")
	for i := 0; i < outputLen; i++ {
		book := books.Books[i]
		fmt.Fprintf(w, "%v\t%v\t%v\t\n", book.Title, book.Author, strings.Join(book.Tags, ", "))
	}
	w.Flush()
}
