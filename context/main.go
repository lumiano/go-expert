package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func BookHotel(ctx context.Context) {

	select {

	case <-time.After(1 * time.Second):
		fmt.Println("Hotel booking done.")
		return

	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled, Timeout reached.", ctx.Err())

		return
	}

}

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*2)

	defer cancel()

	BookHotel(ctx)

	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		return
	}

}
