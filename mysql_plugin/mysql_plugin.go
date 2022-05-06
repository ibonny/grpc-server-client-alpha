package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	pb "grpc-server-client-alpha/protocol"

	empty "github.com/golang/protobuf/ptypes/empty"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"google.golang.org/grpc"
)

var files []*pb.FileRecord
var db *gorm.DB

type FileRecord struct {
	Uuid      string
	Bucket    string
	Prefix    string
	Filename  string
	Chunksize int32
	Filesize  int64
	Chunklist string
	Writetime int64
}

func (fr *FileRecord) putChunkList(cl []string) {
	data, err := json.Marshal(cl)

	if err != nil {
		log.Fatalf("Error marshalling: %v\n", err)
	}

	fr.Chunklist = string(data)
}

func (fr *FileRecord) getChunkList() []string {
	var cl []string

	err := json.Unmarshal([]byte(fr.Chunklist), &cl)

	if err != nil {
		log.Fatalf("Error unmarshalling: %v\n", err)
	}

	return cl
}

type Server struct {
	pb.UnimplementedDataInterfaceServer
}

func (s *Server) GetByFilename(ctx context.Context, in *pb.FileRequest) (*pb.FileResponse, error) {
	retVal := pb.FileResponse{Success: false, Message: "File Not Found."}

	for _, file := range files {
		if file.Filename == in.Fr.Filename {
			retVal.Success = true

			retVal.Fr = append(retVal.Fr, file)

			return &retVal, nil
		}
	}

	return &retVal, nil
}

func (s *Server) GetAllFiles(ctx context.Context, in *empty.Empty) (*pb.FileResponse, error) {
	result := &pb.FileResponse{Success: true}

	var fileRecs []FileRecord

	errResult := db.Find(&fileRecs)

	fmt.Println(errResult.RowsAffected)
	fmt.Println(fileRecs)

	for _, file := range fileRecs {
		fileRec := &pb.FileRecord{
			Uuid:      file.Uuid,
			Bucket:    file.Bucket,
			Prefix:    file.Prefix,
			Filename:  file.Filename,
			Chunksize: file.Chunksize,
			Chunklist: file.getChunkList(),
			Filesize:  file.Filesize,
			Writetime: file.Writetime,
		}

		result.Fr = append(result.Fr, fileRec)
	}

	return result, nil
}

func (s *Server) CreateFile(ctx context.Context, fr *pb.FileRequest) (*pb.FileResponse, error) {
	for _, file := range files {
		if file.Filename == fr.Fr.Filename {
			return &pb.FileResponse{
				Success: false,
				Message: "Error: File already exists.",
			}, nil
		}
	}

	outRec := &FileRecord{
		Uuid:      fr.Fr.Uuid,
		Bucket:    fr.Fr.Bucket,
		Prefix:    fr.Fr.Prefix,
		Filename:  fr.Fr.Filename,
		Chunksize: fr.Fr.Chunksize,
		Filesize:  fr.Fr.Filesize,
		Writetime: time.Now().UnixMilli(),
	}

	outRec.putChunkList(fr.Fr.Chunklist)

	result := db.Create(outRec)

	fmt.Println(result.RowsAffected)

	return &pb.FileResponse{
		Success: true,
		Message: "File created successfully.",
	}, nil
}

func openDatabaseConnection() {
	var err error

	db, err = gorm.Open(mysql.Open("root:password@tcp(127.0.0.1:3306)/storage"))

	if err != nil {
		log.Fatalf("Unable to open MySQL connection: %v\n", err.Error())
	}

	db.AutoMigrate(&FileRecord{})
}

func main() {
	port := 50051

	var lis net.Listener
	var err error

	if lis, err = net.Listen("tcp", fmt.Sprintf("localhost:%d", port)); err != nil {
		log.Fatalf("Error listening on port: %v", err)
	}

	openDatabaseConnection()

	s := grpc.NewServer()

	pb.RegisterDataInterfaceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	log.Printf("Server listening at %v", lis.Addr())
}
