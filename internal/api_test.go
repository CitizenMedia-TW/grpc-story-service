package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-story-service/protobuffs/auth-service"
	"grpc-story-service/protobuffs/story-service"
	"net/http"
	"testing"
)

func GetAuthToken(t *testing.T) string {
	grpcClient, err := grpc.Dial("157.230.46.45:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)
	token, err := auth.NewAuthServiceClient(grpcClient).GenerateToken(context.TODO(), &auth.GenerateTokenRequest{
		Mail: "110703065@nccu.edu.tw",
		Id:   "65dc9342ec31f6f8e14bbbf6",
		Name: "hahaha",
	})
	assert.NoError(t, err)
	return token.Token
}

func TestCreateAndGetStory(t *testing.T) {
	token := GetAuthToken(t)
	body := story.CreateStoryRequest{
		Tags:    []string{"test1", "test2"},
		Content: "test content",
		Title:   "test title",
	}
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(body)
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "http://localhost:50051/story", b)
	request.Header.Set("Authorization", "Bearer "+token)
	assert.NoError(t, err)

	response, err := http.DefaultClient.Do(request)
	assert.NoError(t, err)

	resBody := story.CreateStoryResponse{}
	println(response.Status + " wwwww")
	err = json.NewDecoder(response.Body).Decode(&resBody)
	assert.NoError(t, err)
	println(resBody.Message)

}
