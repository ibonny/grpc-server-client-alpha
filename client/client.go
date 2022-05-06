package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-server-client-alpha/protocol"
)

var pids = make(map[string]*os.Process)

func startServer(pluginName string) {
	fmt.Printf("Starting plugin %s.\n", pluginName)

	cmd := exec.Command(pluginName)

	cmd.Stdout = os.Stdout

	if err := cmd.Start(); err != nil {
		log.Fatalf("Error starting plugin %s: %v\n", pluginName, err)
	}

	pids[pluginName] = cmd.Process
}

func endServer(pluginName string) {
	if _, found := pids[pluginName]; !found {
		return
	}

	if err := pids[pluginName].Kill(); err != nil {
		log.Fatalf("Error killing plugin %s: %v\n", pluginName, err)
	}
}

func pluginWrapper(pluginName string, callback func()) {
	startServer(pluginName)

	time.Sleep(500 * time.Millisecond)

	callback()

	endServer(pluginName)
}

func main() {
	pluginWrapper("server/server.plugin", func() {
		addr := "localhost:50051"

		var conn *grpc.ClientConn
		var err error

		if conn, err = grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
			log.Fatalf("Did not connect: %v", err)
		}

		defer conn.Close()

		c := pb.NewDataInterfaceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)

		defer cancel()

		var r *pb.ChunkResponse

		if r, err = c.GetChunk(ctx, &pb.ChunkRequest{Id: "12345"}); err != nil {
			log.Fatalf("Could not greet: %v", err)
		}

		log.Printf("Response: %v\n", r)
	})
}
