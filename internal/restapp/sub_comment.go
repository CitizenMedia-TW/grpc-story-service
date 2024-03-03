package restapp

import (
	"encoding/json"
	"grpc-story-service/protobuffs/story-service"
	"net/http"
)

func (s RestApp) SubCommentRoute(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		s.CreateSubComment(writer, request)
		return
	case "DELETE":
		s.DeleteSubComment(writer, request)
		return
	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (s RestApp) CreateSubComment(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId")
	if userId == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	in := &story.CreateSubCommentRequest{}
	err := json.NewDecoder(r.Body).Decode(in)
	if err != nil {
		http.Error(w, "Error decoding request", http.StatusUnauthorized)
		return
	}
	in.CommenterId = userId.(string)
	res, err := s.helper.CreateSubComment(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func (s RestApp) DeleteSubComment(writer http.ResponseWriter, request *http.Request) {
	userId := request.Context().Value("userId")
	if userId == nil {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}

	in := &story.DeleteSubCommentRequest{}
	err := json.NewDecoder(request.Body).Decode(in)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	in.DeleterId = userId.(string)

	res, err := s.helper.DeleteSubComment(request.Context(), in)

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
