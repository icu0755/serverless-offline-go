package main

import (
	"bytes"
	"encoding/base64"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"image"
	"image/color"
	"image/jpeg"
)

type Response struct {
	Message string `json:"message"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))
	col := color.RGBA{255, 0, 0, 255}
	for x1 := 30; x1 <= 50; x1++ {
		for y1 := 30; y1 <= 50; y1++ {
			img.Set(x1, y1, col)
		}
	}

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)
	headers := make(map[string]string)
	headers["Content-Type"] = "image/jpeg"

	return events.APIGatewayProxyResponse{
		Body:            base64.StdEncoding.EncodeToString(buf.Bytes()),
		StatusCode:      200,
		IsBase64Encoded: true,
		Headers:         headers,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
