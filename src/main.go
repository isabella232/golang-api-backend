package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
	"regexp"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func helloHtml(path string) string {
	no_foo := convertFoo(path)
	return fmt.Sprintf("<h1>Salutations AMPD from Go,</h1>\n<p>You've requested: %s</p>\n", html.EscapeString(no_foo))
}

func convertFoo(path string) string {
	re := regexp.MustCompile(`(?i)foo`)
	path = re.ReplaceAllString(path, "bar")
	return path
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, helloHtml(r.URL.Path))
	})

	log.Debug("Starting hello-ampd, listening on port " + os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
