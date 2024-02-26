package controllers

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-story-service/internal/app"
	"grpc-story-service/protobuffs/auth-service"
	"log"
	"net/http"
)

type HttpRoutes struct {
	app        app.App
	authClient auth.AuthServiceClient
}

func NewHttpRoutes(a app.App) HttpRoutes {
	grpcClient, err := grpc.Dial("http://157.230.46.45/auth-service", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}
	return HttpRoutes{
		app:        a,
		authClient: auth.NewAuthServiceClient(grpcClient),
	}
}

func (routes HttpRoutes) middlewares(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing middleware")

		res, err := routes.authClient.VerifyToken(context.Background(), &auth.VerifyTokenRequest{Token: r.Header.Get("Authorization")})

		if err == nil {
			r = r.WithContext(context.WithValue(r.Context(), "userId", res.JwtContent.Id))
			r = r.WithContext(context.WithValue(r.Context(), "mail", res.JwtContent.Mail))
		}

		//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
		// log.Print("Executing middlewareTwo again")
	})
}

func (routes HttpRoutes) Routes() http.Handler {
	// Declare a new router
	mux := http.NewServeMux()

	// Use middleware for all routes
	mux.Handle("/story", routes.middlewares(http.HandlerFunc(routes.StoryRoute)))
	mux.Handle("/recommend", routes.middlewares(http.HandlerFunc(routes.GetRecommendStory)))
	mux.Handle("/comment", routes.middlewares(http.HandlerFunc(routes.CommentRoute)))
	mux.Handle("/subComment", routes.middlewares(http.HandlerFunc(routes.SubCommentRoute)))

	return mux
}
