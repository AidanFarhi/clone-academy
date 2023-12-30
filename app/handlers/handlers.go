package handlers

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("yolo")
	http.ServeFile(w, r, "app/web/index.html")
}
