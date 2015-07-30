package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/getmillipede/millipede-go"
)

func handler(w http.ResponseWriter, r *http.Request) {
	millipede := millipede.New(20)
	for k, v := range r.URL.Query() {
		switch k {
		// uint64 fields
		case "size":
			size, err := strconv.ParseUint(v[0], 10, 32)
			if err != nil {
				fmt.Fprintf(w, "Error: %v", err)
				return
			}
			millipede.Size = size
		case "curve":
			curve, err := strconv.ParseUint(v[0], 10, 32)
			if err != nil {
				fmt.Fprintf(w, "Error: %v", err)
				return
			}
			millipede.Curve = curve
		case "width":
			width, err := strconv.ParseUint(v[0], 10, 32)
			if err != nil {
				fmt.Fprintf(w, "Error: %v", err)
				return
			}
			millipede.Width = width
		case "steps":
			steps, err := strconv.ParseUint(v[0], 10, 32)
			if err != nil {
				fmt.Fprintf(w, "Error: %v", err)
				return
			}
			millipede.Steps = steps

			// string fields
		case "skin":
			millipede.Skin = v[0]

			// boolean fields
		case "reverse":
			reverse, err := strconv.ParseBool(v[0])
			if err != nil {
				fmt.Fprintf(w, "error: %v", err)
			}
			millipede.Reverse = reverse
		case "zalgo":
			zalgo, err := strconv.ParseBool(v[0])
			if err != nil {
				fmt.Fprintf(w, "error: %v", err)
			}
			millipede.Zalgo = zalgo
		case "opposite":
			opposite, err := strconv.ParseBool(v[0])
			if err != nil {
				fmt.Fprintf(w, "error: %v", err)
			}
			millipede.Opposite = opposite
		}
	}
	fmt.Fprint(w, millipede)
}

func main() {
	http.HandleFunc("/", handler)
	log.Printf("About to listen on 4242. Go to http://127.0.0.1:4242/")
	err := http.ListenAndServe(":4242", nil)
	if err != nil {
		log.Fatal(err)
	}
}
