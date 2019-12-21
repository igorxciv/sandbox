package main

import (
	"context"
	blogpb "crud-proto/proto"
	"fmt"
	"github.com/jessevdk/go-flags"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"time"
)

type Opts struct {
	Author  string `long:"author" short:"a" description:"Add an author" required:"true"`
	Title   string `long:"title" short:"t" description:"Add a title" required:"true"`
	Content string `long:"content" short:"c" description:"Add a content" required:"true"`
}

var client blogpb.BlogServiceClient
var reqCtx context.Context
var opts Opts

func init() {
	p := flags.NewParser(&opts, flags.Default)
	if _, err := p.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.HelpFlag {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	fmt.Println("Starting Blog Service Client")
	reqCtx, _ = context.WithTimeout(context.Background(), 10*time.Second)

	reqOpts := grpc.WithInsecure()
	conn, err := grpc.Dial(":50051", reqOpts)
	if err != nil {
		log.Fatalf("Unable to establish client connection to localhost:50051: %v", err)
	}
	client = blogpb.NewBlogServiceClient(conn)
}

func main() {
	blog := &blogpb.Blog{
		AuthorId: opts.Author,
		Title:    opts.Title,
		Content:  opts.Content,
	}
	res, err := client.CreateBlog(context.TODO(), &blogpb.CreateBlogReq{
		Blog: blog,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Blog created: %s\n", res.Blog.Id)

	// ADD
	stream, err := client.ListBlog(context.TODO(), &blogpb.ListBlogReq{})
	if err != nil {
		panic(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Println(res.GetBlog())
	}
}
