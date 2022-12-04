package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

// IndexData ..
type IndexData struct {
	Header    string
	Paragraph string
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	indexData := IndexData{
		Header:    "Заголовок",
		Paragraph: `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Amet volutpat consequat mauris nunc congue nisi. In iaculis nunc sed augue lacus viverra vitae congue. Consequat nisl vel pretium lectus quam. Purus ut faucibus pulvinar elementum. Non diam phasellus vestibulum lorem sed risus ultricies. Ipsum dolor sit amet consectetur adipiscing elit duis. Mi quis hendrerit dolor magna eget. Libero justo laoreet sit amet cursus sit amet dictum. Gravida in fermentum et sollicitudin ac orci phasellus. Commodo sed egestas egestas fringilla phasellus faucibus. Ultrices neque ornare aenean euismod elementum nisi. Nulla aliquet porttitor lacus luctus accumsan tortor posuere ac ut. Donec enim diam vulputate ut pharetra sit amet. Neque laoreet suspendisse interdum consectetur libero id. Aliquam sem fringilla ut morbi. Velit scelerisque in dictum non consectetur a. Leo a diam sollicitudin tempor. Blandit aliquam etiam erat velit scelerisque in.`,
	}
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, indexData)
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", getIndex)
	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
