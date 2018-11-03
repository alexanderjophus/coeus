package thing

import (
	"io"
	"log"
)

// Exec is entry point
func Exec(w io.Writer, r RawResponse) error {
	for _, v := range r.Highlights.GameCenter.Items {
		_, err := w.Write([]byte(v.Description + ".\n"))
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}
