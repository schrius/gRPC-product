package main

import (
	"context"
	"log"
	pb "productinfo/client/ecommerce"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection with the server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}

	// close connection
	defer conn.Close()

	// Pass the connection and create a stub
	c := pb.NewProductInfoClient(conn)

	name := "iPhone X"
	description := `Meet iPhone x. All-new dual-camera system with Ultra Wide and Night mode.`

	price := float32(1000.0)
	// Create a Context to pass with the remote call
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// call addProduct method
	r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description, Price: price})
	if err != nil {
		log.Fatal("Cloud not add product: %v", err)
	}
	log.Printf("Product ID: %s added successfully", r.Value)

	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})

	if err != nil {
		log.Fatalf("Cloud not get product: %v", err)
	}
	log.Printf("Product: ", product.String())
}
