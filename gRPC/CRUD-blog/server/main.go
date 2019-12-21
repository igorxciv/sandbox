package main

import (
	"context"
	blogpb "crud-proto/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type BlogServiceServer struct{}

func (s *BlogServiceServer) CreateBlog(ctx context.Context, req *blogpb.CreateBlogReq) (*blogpb.CreateBlogRes, error) {
	return nil, nil
}

func (s *BlogServiceServer) ReadBlog(ctx context.Context, req *blogpb.ReadBlogReq) (*blogpb.ReadBlogRes, error) {
	return nil, nil
}

func (s *BlogServiceServer) UpdateBlog(ctx context.Context, req *blogpb.UpdateBlogReq) (*blogpb.UpdateBlogRes, error) {
	return nil, nil
}

func (s *BlogServiceServer) DeleteBlog(ctx context.Context, req *blogpb.DeleteBlogReq) (*blogpb.DeleteBlogRes, error) {
	return nil, nil
}

func (s *BlogServiceServer) ListBlog(ctx context.Context, req *blogpb.ListBlogReq) (*blogpb.ListBlogRes, error) {
	return nil, nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port :50051...")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Unable to listen on port :50051: %v", err)
	}

	var opts []grpc.ServerOption

	s := grpc.NewServer(opts...)
	srv := &BlogServiceServer{}

	blogpb.RegisterBlogServiceServer(s, srv)
}
