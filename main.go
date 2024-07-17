package main

import (
	"log"

	"github.com/rohitkrcodes/go_data_structures/api"
	"github.com/rohitkrcodes/go_data_structures/lists"
)

const (
	serverAddress = "0.0.0.0:8080"
	dbAddress     = "/data/data.txt"
)

func main() {
	// ll := lists.NewDoublyList()

	// ll.Append(1)
	// ll.Append(2)
	// ll.Append(3)
	// ll.Append(4)
	// ll.Append(5)

	// ll.Prepend(0)

	// ll.Remove(3)
	// ll.PrintList()

	db := lists.NewDoublyList(dbAddress)

	server := api.NewServer(db)

	err := server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
