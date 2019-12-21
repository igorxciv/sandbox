package main

import (
	"context"
	blogpb "crud-proto/proto"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
)

type BlogServiceServer struct{}

type BlogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

func (s *BlogServiceServer) CreateBlog(ctx context.Context, req *blogpb.CreateBlogReq) (*blogpb.CreateBlogRes, error) {
	blog := req.GetBlog()
	data := BlogItem{
		AuthorID: blog.AuthorId,
		Content:  blog.Content,
		Title:    blog.Title,
	}
	result, err := blogdb.InsertOne(mongoCtx, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v", err))
	}

	oid := result.InsertedID.(primitive.ObjectID)
	blog.Id = oid.Hex()
	return &blogpb.CreateBlogRes{
		Blog: blog,
	}, nil
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

	blogdb = db.Database("mydb").Collection("blog")

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server successfully started on port :50051")

	c := make(chan os.Signal)

	signal.Notify(c, os.Interrupt)

	<-c

	fmt.Println("\nStopping the server...")
	s.Stop()
	if err := listener.Close(); err != nil {
		log.Fatalf("Failed to close TCP connection: %v", err)
	}
	fmt.Println("Closing MongoDB connection...")

	if err := db.Disconnect(mongoCtx); err != nil {
		log.Fatalf("Failed to close MongoDB connection: %v", err)
	}
	fmt.Println("Done.")
}
