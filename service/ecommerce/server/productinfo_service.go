package main

import (
	"context"
	pb "productinfo/service/ecommerce" // import the generated code

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// server is used to implemented ecommerce/product_info
// abstraction server
type server struct {
	productMap map[string]*pb.Product
}

// AddProduct implements ecommerce.AddProduct
// the method takes Product as a parameter and returns a ProductID
// Context object contains metadata such as the identify of the end user authorization tokens and the requests's deadlin
func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating Product ID", err)
	}

	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in
	return &pb.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

// GetProduct implements ecommerce.GetProduct
// the method takes ProductID as a paramenter and return a Product or error
// Context object contains metadata such as the identify of the end user authorization tokens and the requests's deadlin
func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	value, exists := s.productMap[in.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Product does not exist.", in.Value)
}
