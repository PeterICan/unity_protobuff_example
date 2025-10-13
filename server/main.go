package main

import (
    "fmt"
    "github.com/magiclvzs/antnet"
    "github.com/magiclvzs/antnet/parser"
    "proto_buffer_example/server/handlers" // Import the new package
)

func main() {
    s, err := antnet.NewServer("tcp://:6666",
        antnet.WithParser(parser.NewPBParser(4096, true)),
    )
    if err != nil {
        fmt.Printf("Failed to create server: %v\n", err)
        return
    }

    // Register handlers from the handlers package
    s.Register(1, &handlers.LoginHandler{})
    s.Register(3, &handlers.PositionHandler{}) // Register the new handler

    fmt.Println("Starting antnet server on tcp://:6666")
    s.Run()
}

