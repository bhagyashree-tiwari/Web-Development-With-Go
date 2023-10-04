package main

import (
	"fmt"
	"net/http"
)

// homeHandler is a function that handles incoming HTTP requests.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Respond to the request with an HTML message.
	fmt.Fprint(w, "<h1>Welcome to my site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:bhagyashreetiwari12@gmail.com\">bhagyashreetiwari12@gmail.com</a>. </p>")

}
func main() {
	// Register the homeHandler function to handle requests at the root ("/") path.
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting the server on :3000....")

	// Start the HTTP server on port 3000 and handle incoming requests using the default router (nil).
	http.ListenAndServe(":3000", nil)
}
