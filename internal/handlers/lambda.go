package handlers

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gabrielmq/zapgpt/internal/chatgpt"
)

func Handle(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	result, err := parseBase64RequestData(request.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	text, err := chatgpt.GetText(result)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       text,
	}, nil
}

func parseBase64RequestData(r string) (string, error) {
	dataBytes, err := base64.StdEncoding.DecodeString(r)
	if err != nil {
		return "", err
	}

	data, err := url.ParseQuery(string(dataBytes))
	if err != nil {
		return "", err
	}

	if data.Has("Body") {
		return data.Get("Body"), nil
	}

	return "", errors.New("body not found")
}
