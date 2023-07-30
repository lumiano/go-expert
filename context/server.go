package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r = r.WithContext(r.Context())

	log.Printf("Request received. Context: %+v\n", r.Context())

	defer log.Println("Handler ended.")

	select {

	case <-r.Context().Done():
		log.Println("Handler cancelled, Timeout reached.", r.Context().Err())

		http.Error(w, http.StatusText(http.StatusRequestTimeout), http.StatusRequestTimeout)

		return

	case <-time.After(5 * time.Second):

		response := make(map[string]string)

		response["message"] = "Request processed."
		response["status"] = strconv.Itoa(http.StatusNoContent)

		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			return
		}

		w.WriteHeader(http.StatusNoContent)

		log.Println("Response sent.")

	}

}
