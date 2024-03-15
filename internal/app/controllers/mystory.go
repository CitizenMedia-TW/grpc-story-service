package controllers

import (
	"encoding/json"
	"net/http"
)

func (routes HttpRoutes) MyStoryRoute(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		routes.GetMyStory(writer, request)
		return
	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (routes HttpRoutes) GetMyStory(writer http.ResponseWriter, request *http.Request) {
	userId := request.Context().Value(UserIdContextKey{})
	if userId == nil {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}
	stories, err := routes.app.GetMyStoryIds(request.Context(), userId.(string))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(writer).Encode(&StoryIds{StoryIds: stories})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

type StoryIds struct {
	StoryIds []string `json:"storyIds"`
}
