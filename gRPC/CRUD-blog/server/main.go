package main

import (
	"context"
	blogpb "crud-proto/proto"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

var db *mongo.Client
var blogdb *mongo.Collection
var mongoCtx context.Context

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

	fmt.Println("Connecting to MongoDB")
	mongoCtx = context.Background()

	db, err := mongo.Connect(mongoCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping(mongoCtx, nil)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	} else {
		log.Println("Connected to MongoDB!")
	}
}
