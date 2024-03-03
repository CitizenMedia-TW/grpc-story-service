package helper

import (
	"grpc-story-service/internal/database"
	"grpc-story-service/protobuffs/auth-service"
)

type Helper struct {
	database   database.Database
	AuthClient auth.AuthServiceClient
}

func New(authClient auth.AuthServiceClient) Helper {
	db := database.New()

	// grpcClient, err := grpc.Dial("157.230.46.45:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	panic(err)
	// }

	return Helper{
		database:   db,
		AuthClient: authClient,
	}
}
