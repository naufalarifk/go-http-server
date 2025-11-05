package response

import (
	"fmt"
	"httpfromtcp/internal/headers"
	"io"
)

type StatusCode int

const (
	StatusOK                  StatusCode = 200
	StatusBadRequest          StatusCode = 400
	StatusInternalServerError StatusCode = 500
)

func GetDefaultHeaders(contentLen int) *headers.Headers {
	h := headers.NewHeaders()
	h.Set("Content-Length", fmt.Sprintf("%d", contentLen))
	h.Set("Connection", "close")
	h.Set("Content-Type", "text-plain")
	return h
}

func WriteHeaders(w io.Writer, h *headers.Headers) error {
	b := []byte{}
	h.Foreach(func(n, v string) {
		b = fmt.Appendf(b, "%s: %s\r\n", n, v)
	})
	b = fmt.Append(b, "\r\n")
	_, err := w.Write(b)
	return err
}

func WriteStatusLine(w io.Writer, statusCode StatusCode) error {
	statusLine := []byte{}
	switch statusCode {
	case StatusOK:
		statusLine = []byte("HTTP/1.1 200 OK")
	case StatusBadRequest:
		statusLine = []byte("HTTP/1.1 400 OK")
	case StatusInternalServerError:
		statusLine = []byte("HTTP/1.1 500 OK")
	default:
		return fmt.Errorf("Unrecognized Error Code")
	}
	_, err := w.Write(statusLine)
	return err
}
