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
	"google.golang.org/protobuf/types/known/emptypb"

	pb "grpc-server-client-alpha/protocol"
)

var pids = make(map[string]*os.Process)

func startServer(pluginName string) {
	// fmt.Printf("Starting plugin %s.\n", pluginName)

	cmd := exec.Command(pluginName)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	if err := cmd.Start(); err != nil {
		log.Fatalf("Error starting plugin %s: %v\n", pluginName, err)
	}

	pids[pluginName] = cmd.Process
}

func endServer(pluginName string) {
	if _, found := pids[pluginName]; !found {
		return
	}

	// fmt.Printf("Attempting to kill process: %v\n", pids[pluginName].Pid)

	if err := pids[pluginName].Kill(); err != nil {
		log.Fatalf("Error killing plugin %s: %v\n", pluginName, err)
	}
}

func pluginWrapper(pluginName string, callback func()) {
	startServer(pluginName)

	// This delay is required in order to give the server some time to start.
	time.Sleep(100 * time.Millisecond)

	callback()

	time.Sleep(100 * time.Millisecond)

	endServer(pluginName)
}

func debugPlugin(commands []string, callback func()) {
	cmd := exec.Command(commands[0], commands[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	if err := cmd.Start(); err != nil {
		log.Fatalf("Error starting plugin %s: %v\n", commands[0], err)
	}

	pids[commands[0]] = cmd.Process

	// This delay is required in order to give the server some time to start.
	time.Sleep(300 * time.Millisecond)

	callback()

	time.Sleep(100 * time.Millisecond)

	endServer(commands[0])
}

func openClientConnection() (pb.DataInterfaceClient, context.Context, context.CancelFunc, error) {
	addr := "localhost:50051"

	var conn *grpc.ClientConn
	var err error

	if conn, err = grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	// defer conn.Close()

	c := pb.NewDataInterfaceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	// defer cancel()

	return c, ctx, cancel, nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("\n   Please use: " + os.Args[0] + " [list|create]\n")

		os.Exit(0)
	}

	// pluginWrapper("mysql_plugin/mysql.plugin", func() {
	pluginWrapper("mysql_plugin/mysql.plugin", func() {
		c, ctx, cancel, err := openClientConnection()

		if os.Args[1] == "create" {
			if len(os.Args) == 2 {
				fmt.Println("\n   Please provide a filename.\n")

				return
			}

			var r *pb.FileResponse

			fr := &pb.FileRequest{
				Fr: &pb.FileRecord{
					Filename:  os.Args[2],
					Filesize:  1024,
					Chunksize: 1024,
					Chunklist: []string{},
				},
			}

			if r, err = c.CreateFile(ctx, fr); err != nil {
				log.Printf("%v\n", r)
				log.Fatalf("Could not Create File: %v", err)
			}

			log.Printf("Response: %v\n", r)
		}

		if os.Args[1] == "list" {
			var r *pb.FileResponse

			if r, err = c.GetAllFiles(ctx, &emptypb.Empty{}); err != nil {
				log.Fatalf("Error: %v\n", err)
			}

			loc, err := time.LoadLocation("US/Eastern")

			if err != nil {
				log.Fatalf("Error loading time location: %v\n", err)
			}

			if len(r.Fr) == 0 {
				fmt.Println("\n   There are no files.\n")
			} else {
				for _, file := range r.Fr {
					fmt.Printf("%s %s %s %s %d %d %v %s\n",
						file.Uuid,
						file.Bucket,
						file.Prefix,
						file.Filename,
						file.Chunksize,
						file.Filesize,
						file.Chunklist,
						time.Unix(0, file.Writetime*int64(time.Millisecond)).In(loc),
					)
				}
			}
		}

		cancel()
	})
}
