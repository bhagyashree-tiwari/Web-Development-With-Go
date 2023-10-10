package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// homeHandler is a function that handles incoming HTTP requests.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Respond to the request with an HTML message.
	fmt.Fprint(w, "<h1>Welcome to my site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:bhagyashreetiwari12@gmail.com\">bhagyashreetiwari12@gmail.com</a> </p>")

}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ page</h1> 
	<p>You can write faqs here</p>
	<ul>
		<li>Is ther any questions?</li>
		<li> please keep it to yourself</li>
	</ul>`)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
