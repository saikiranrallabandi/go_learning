package main
import (
	"net/http"
	"text/template"
	"log"
)

func Home(writer http.ResponseWriter, reader *http.Request)  {
	var template_html *template.Template
	template_html = template.Must(template.ParseFiles("main.html"))
	template_html.Execute(writer, nil)

}
func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8080", nil)
}