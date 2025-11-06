package main

import (
	"fmt"
	"httpfromtcp/internal/request"
	"httpfromtcp/internal/response"
	"httpfromtcp/internal/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const port = ":42069"

func respond400() []byte {
	return []byte(`
	<html>
			<head>
			<title>400 Bad Request</title>
			</head>
			<body>
			<h1> Bad Request</h1>
			<p>Your Request Kinda Sucks</p>
			</body>
			</html>`)
}

func respond500() []byte {
	return []byte(`
	<html>
			<head>
			<title>400 Bad Request</title>
			</head>
			<body>
			<h1>Bad Request</h1>
			<p>Your Request Kinda Sucks</p>
			</body>
			</html>`)
}

func respond200() []byte {
	return []byte(`
	<html>
			<head>
			<title>200 OK</title>
			</head>
			<body>
			<h1>Success!</h1>
			<p>Your Request Is Absolute Banger</p>
			</body>
			</html>`)
}

func main() {
	s, err := server.Serve(port, func(w *response.Writer, req *request.Request) {
		h := response.GetDefaultHeaders(0)
		body := respond200()
		status := response.StatusOK

		if req.RequestLine.RequestTarget == "/yourproblem" {
			w.WriteStatusLine(response.StatusBadRequest)
			body = respond400()
			status = response.StatusBadRequest
		} else if req.RequestLine.RequestTarget == "/myproblem" {
			w.WriteStatusLine(response.StatusBadRequest)
			body = respond500()
			status = response.StatusInternalServerError
		}
		h.Replace("Content-Length", fmt.Sprintf("%d", len(body)))
		h.Replace("Content-Type", "text/html")
		w.WriteStatusLine(status)
		w.WriteHeaders(*h)

		w.WriteBody(body)
	})
	if err != nil {
		log.Fatalf("Error Starting Server: %v", err)
	}

	defer s.Close()
	log.Println("Server started on port", port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Server gracefully stopped")
}
