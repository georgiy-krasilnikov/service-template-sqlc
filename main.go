package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"

	"grpc-service-template-sqlc/db"
	"grpc-service-template-sqlc/pb"
	"grpc-service-template-sqlc/services/users"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := ReadConfig()
	if err != nil {
		log.WithError(err).Fatal("failed to read config")
		os.Exit(1)
	}

	dbc, err := pgxpool.Connect(context.Background(), config.DBConnString)
	if err != nil {
		log.WithError(err).Error("failed connect to database")
		os.Exit(1)
	}

	query := db.New(dbc)
	if err := db.Migrate(config.DBConnString); err != nil {
		log.WithError(err).Error("failed to run migrations")
		os.Exit(1)
	}

	usersServiceHandlers := users.New(query)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Port))
	if err != nil {
		log.WithError(err).Fatal("failed to start listen")
		os.Exit(1)
	}

	server := grpc.NewServer()
	reflection.Register(server)

	pb.RegisterUsersServer(server, usersServiceHandlers)

	if err := server.Serve(lis); err != nil {
		log.WithError(err).Fatal("failed to start serve")
		os.Exit(1)
	}
}
