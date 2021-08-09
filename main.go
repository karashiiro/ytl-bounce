package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type failure struct {
	Error string `json:"error"`
}

var videoIdRegex = regexp.MustCompile(`https?:\/\/www\.youtube\.com\/watch\?v=(?P<videoId>[^#\&\?%"<>]*)`)

func jsonFailure(message string) events.APIGatewayProxyResponse {
	res := failure{Error: message}
	body, _ := json.Marshal(res)
	return events.APIGatewayProxyResponse{
		StatusCode: 400,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	channelId, ok := request.QueryStringParameters["c"]
	if !ok {
		fail := jsonFailure("Failed to get channel ID from request")
		return &fail, nil
	}

	livePath := fmt.Sprintf("https://www.youtube.com/channel/%s/live", channelId)

	res, err := http.Get(livePath)
	if err != nil {
		fail := jsonFailure(err.Error())
		return &fail, nil
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fail := jsonFailure(err.Error())
		return &fail, nil
	}

	matches := videoIdRegex.FindStringSubmatch(string(body))
	if len(matches) == 0 {
		fail := jsonFailure("No live stream found")
		return &fail, nil
	}

	videoId := matches[1]
	videoPath := fmt.Sprintf("https://www.youtube.com/live_chat?is_popout=1&v=%s", videoId)

	return &events.APIGatewayProxyResponse{
		StatusCode: 302,
		Headers: map[string]string{
			"Location": videoPath,
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
