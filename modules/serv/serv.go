package serv

import (
	"context"
	"fmt"
	"net/http"
)

func Cancel(cancel context.CancelFunc) {
	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Service stoping...")
		cancel()
	})
	http.ListenAndServe(":12345", nil)

}
