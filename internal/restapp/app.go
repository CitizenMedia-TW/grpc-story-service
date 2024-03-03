package restapp

import (
	"context"
	"grpc-story-service/internal/helper"
	"grpc-story-service/protobuffs/auth-service"
	"log"
	"net/http"
)

type RestApp struct {
	helper helper.Helper
}

func New(authClient auth.AuthServiceClient) RestApp {
	h := helper.New(authClient)

	return RestApp{
		helper: h,
	}
}

func (s RestApp) middlewares(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing middleware")

		res, err := s.helper.AuthClient.VerifyToken(context.Background(), &auth.VerifyTokenRequest{Token: r.Header.Get("Authorization")})

		if err == nil && res.Message != "Failed" {
			println(res.Message)
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

func (s RestApp) Routes() http.Handler {
	// Declare a new router
	mux := http.NewServeMux()

	// Use middleware for all routes
	mux.Handle("/story", s.middlewares(http.HandlerFunc(s.StoryRoute)))
	mux.Handle("/recommend", s.middlewares(http.HandlerFunc(s.GetRecommendStory)))
	mux.Handle("/comment", s.middlewares(http.HandlerFunc(s.CommentRoute)))
	mux.Handle("/subComment", s.middlewares(http.HandlerFunc(s.SubCommentRoute)))

	return mux
}
