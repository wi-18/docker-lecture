package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

const HelloWorldHtml string = `
<!doctype html>
<html lang="en">
<head>
	<meta charset="UTF-8">            
	<title>FrontendHello World</title>
</head>
<body style="width: 100%; height: 100%">
	<h1>The backend is running on the host: {{.Host}}</h1>
</body>
</html>
`

func main() {
	port := 80
	tmpl, err := template.New("index").Parse(HelloWorldHtml)
	if err != nil {
		log.Fatal(err)
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		result, err := http.Get("http://orch-backend-svc")
		if err != nil {
			writer.WriteHeader(500)
			return
		}

		defer result.Body.Close()
		body, err := ioutil.ReadAll(result.Body)
		if err != nil {
			writer.WriteHeader(500)
			return
		}

		writer.WriteHeader(200)
		_ = tmpl.Execute(writer, WebTemplateData{Host: string(body)})
	})

	fmt.Println(fmt.Sprintf("listetning on port %d", port))
	_ = http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}

type WebTemplateData struct {
	Host string
}
