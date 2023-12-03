package applicator

import (
	"payservice/internal/user/config"
	"payservice/internal/user/controller/grpc"
	"payservice/internal/user/controller/http"
	"payservice/internal/user/database"
	"payservice/internal/user/repository"
)

func Run(cfg *config.Config) {
	mainDB := database.Connect(cfg.Database.Main)
	defer mainDB.Close()

	replicaDB := database.Connect(cfg.Database.Replica)
	defer replicaDB.Close()

	repo := repository.NewUserRepository(mainDB, replicaDB)

	grpcService := grpc.NewService(repo)
	grpcServer := grpc.NewServer(cfg.GrpcServer.Port, grpcService)
	err := grpcServer.Start()
	if err != nil {
		panic(err)
	}

	defer grpcServer.Close()

	http.Start(repo)
}
