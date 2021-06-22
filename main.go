package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Reywaltz/backend_xsolla/internal/models"
)

func main() {
	http.HandleFunc("/", TestHandler)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	var item models.Item

	if err := item.Bind(r); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println(item)

	return
}
