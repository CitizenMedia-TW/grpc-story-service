package restapp

import (
	"encoding/json"
	"grpc-story-service/internal/database"
	"grpc-story-service/protobuffs/story-service"
	"net/http"
)

func (s RestApp) StoryRoute(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		s.CreateStory(writer, request)
		return
	case "GET":
		s.GetOneStory(writer, request)
		return
	case "DELETE":
		s.DeleteStory(writer, request)
		return
	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (s RestApp) CreateStory(writer http.ResponseWriter, request *http.Request) {
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

	res, err := s.helper.CreateStory(request.Context(), in)

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

func (s RestApp) GetOneStory(writer http.ResponseWriter, request *http.Request) {
	in := &story.GetOneStoryRequest{}
	err := json.NewDecoder(request.Body).Decode(in)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := s.helper.GetOneStory(request.Context(), in)
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

func (s RestApp) DeleteStory(writer http.ResponseWriter, request *http.Request) {
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

	res, err := s.helper.DeleteStory(request.Context(), in)

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
