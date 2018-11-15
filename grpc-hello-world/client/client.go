package main
import (
    "log"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    pb "test.go.dev/grpc-hello-world/proto"
)
func main() {
    creds, err := credentials.NewClientTLSFromFile("/smb/tech/go-v/src/test.go.dev/grpc-hello-world/certs/client.key", "localhost")
    if err != nil {
        log.Println("Failed to create TLS credentials %v", err)
    }
    conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
    defer conn.Close()
    if err != nil {
        log.Println(err)
    }
    c := pb.NewHelloWorldClient(conn)
    context := context.Background()
    body := &pb.HelloWorldRequest{
        Referer : "Grpc",
    }
    r, err := c.SayHelloWorld(context, body)
    if err != nil {
        log.Println(err)
    }
    log.Println(r.Message)
}