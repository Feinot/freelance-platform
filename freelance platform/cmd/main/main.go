package main

import (
	"321/internal/user"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	user.Init()
	http.HandleFunc("/user", user.WorkSpace)
	http.HandleFunc("/user/order_info", user.OrderWorkSpace)
	fmt.Println("Server started on localhost:8000")
	//Starting the server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
