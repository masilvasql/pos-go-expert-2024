package main

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/masilvasql/gprc-pos/internal/database"
	"github.com/masilvasql/gprc-pos/internal/pb"
	"github.com/masilvasql/gprc-pos/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	categoryDb := database.NewCategory(db)
	categryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categryService)

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	fmt.Println("Server running on port 50051")
	grpcServer.Serve(lis)

}
