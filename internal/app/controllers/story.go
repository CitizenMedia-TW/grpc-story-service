package controllers

import (
	"encoding/json"
	"grpc-story-service/internal/database"
	"grpc-story-service/protobuffs/story-service"
	"net/http"
)

func (routes HttpRoutes) StoryRoute(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		routes.CreateStory(writer, request)
		return
	case "GET":
		routes.GetOneStory(writer, request)
		return
	case "DELETE":
		routes.DeleteStory(writer, request)
		return
	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (routes HttpRoutes) CreateStory(writer http.ResponseWriter, request *http.Request) {
	userId := request.Context().Value("userId")

	if userId == nil {
		http.Error(writer, "UnAuthorized", http.StatusUnauthorized)
		return
	}

	in := &story.CreateStoryRequest{}
	err := json.NewDecoder(request.Body).Decode(in)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	in.AuthorId = userId.(string)

	res, err := routes.app.CreateStory(request.Context(), in)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(writer).Encode(res)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func (routes HttpRoutes) GetOneStory(writer http.ResponseWriter, request *http.Request) {
	storyId := request.URL.Query().Get("storyId")
	res, err := routes.app.GetOneStory(request.Context(), storyId)
	if err != nil {
		if err == database.ErrNotFound {
			http.Error(writer, "Story not found", http.StatusNotFound)
			return
		}
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(writer).Encode(res)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func (routes HttpRoutes) DeleteStory(writer http.ResponseWriter, request *http.Request) {
	userId := request.Context().Value("userId")
	if userId == nil {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}

	in := &story.DeleteStoryRequest{}
	err := json.NewDecoder(request.Body).Decode(in)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	in.DeleterId = userId.(string)

	res, err := routes.app.DeleteStory(request.Context(), in)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(writer).Encode(res)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
