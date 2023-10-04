package main

import (
	"fmt"
	"net/http"
)

// handlerfunc is a function that handles incoming HTTP requests.
func handlerfunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Respond to the request with an HTML message.
	fmt.Fprint(w, "<h1>Welcome to my site!</h1>")
}

func main() {
	// Register the handlerfunc function to handle requests at the root ("/") path.
	http.HandleFunc("/", handlerfunc)

	fmt.Println("Starting the server on :3000....")

	// Start the HTTP server on port 3000 and handle incoming requests using the default router (nil).
	http.ListenAndServe(":3000", nil)
}
