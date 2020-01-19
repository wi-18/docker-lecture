package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	NoPortGif string = "https://media1.tenor.com/images/2b6138c8abd50d00965e784d948a88df/tenor.gif"
	PortGif   string = "https://media1.tenor.com/images/1fa0bf4c58df5e5f5142af2098943832/tenor.gif"
)
const HelloWorldHtml string = `
<!doctype html>
<html lang="en">
<head>
	<meta charset="UTF-8">            
	<title>Hello World</title>
</head>
<body style="background: black; width: 100%; height: 100%">
	<img src="{{.Gif}}" alt="Well done" style="display: block;margin-left: auto;margin-right: auto;width: 50%;">
</body>
</html>
`

func main() {
	tmpl, err := template.New("index").Parse(HelloWorldHtml)
	if err != nil {
		log.Fatal(err)
		return
	}

	port := 80
	data := WebTemplateData{Gif: NoPortGif}

	if _, ok := os.LookupEnv("KUBERNETES_SERVICE_HOST"); ok {
		fmt.Println("Oh look, we are running inside a kubernetes cluster! Well done!")
	}

	if envPort, ok := os.LookupEnv("PORT"); ok {
		if val, err := strconv.Atoi(envPort); err != nil {
			fmt.Println(fmt.Sprintf("provided environment variable PORT is not an integer (%s)", envPort))
			os.Exit(1)
		} else {
			port = val
			data.Gif = PortGif
		}
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		_ = tmpl.Execute(writer, data)
	})

	fmt.Println(fmt.Sprintf("listetning on port %d", port))
	_ = http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}

type WebTemplateData struct {
	Gif string
}
