Generate protobuffers: `protoc -I .\CRUD-blog\proto --go_out=.\CRUD-blog\proto .\CRUD-blog\proto\blog.proto`

Generate protobuffers + gRPC Server: `protoc -I .\CRUD-blog\proto --go_out=plugins=grpc:.\CRUD-blog\proto .\CRUD-blog\proto\blog.proto`
