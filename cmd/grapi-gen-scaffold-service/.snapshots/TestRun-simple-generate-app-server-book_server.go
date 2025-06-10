package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/x-izumin/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "testapp/api"
)

// BookServiceServer is a composite interface of api_pb.BookServiceServer and grapiserver.Server.
type BookServiceServer interface {
	api_pb.BookServiceServer
	grapiserver.Server
}

// NewBookServiceServer creates a new BookServiceServer instance.
func NewBookServiceServer() BookServiceServer {
	return &bookServiceServerImpl{}
}

type bookServiceServerImpl struct {
}

func (s *bookServiceServerImpl) ListBooks(ctx context.Context, req *api_pb.ListBooksRequest) (*api_pb.ListBooksResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *bookServiceServerImpl) GetBook(ctx context.Context, req *api_pb.GetBookRequest) (*api_pb.Book, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *bookServiceServerImpl) CreateBook(ctx context.Context, req *api_pb.CreateBookRequest) (*api_pb.Book, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *bookServiceServerImpl) UpdateBook(ctx context.Context, req *api_pb.UpdateBookRequest) (*api_pb.Book, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *bookServiceServerImpl) DeleteBook(ctx context.Context, req *api_pb.DeleteBookRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

