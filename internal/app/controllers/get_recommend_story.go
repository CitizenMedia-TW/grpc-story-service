package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (routes HttpRoutes) GetRecommendStory(writer http.ResponseWriter, request *http.Request) {
	skip, err := strconv.ParseInt(request.URL.Query().Get("skip"), 10, 64)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	count, err := strconv.ParseInt(request.URL.Query().Get("count"), 10, 64)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := routes.app.GetRecommended(request.Context(), skip, count)
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
