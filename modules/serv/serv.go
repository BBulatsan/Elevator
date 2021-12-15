package serv

import (
	"context"
	"fmt"
	"net/http"
)

func Cancel(cancel context.CancelFunc) {
	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Service stoping...")
		if err != nil {
			panic(err)
		}
		cancel()
	})
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic(err)
	}

}
