package main

import (
	"bytes"
	"encoding/base64"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"image"
	"image/jpeg"
)

type Response struct {
	Message string `json:"message"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//mw := imagick.NewMagickWand()
	//defer mw.Destroy()
	//dw := imagick.NewDrawingWand()
	//pw := imagick.NewPixelWand()

	//pw.SetColor("white")
	//dw.Rectangle(0, 0, 100, 100)

	img := image.NewRGBA(image.Rect(0, 0, 100, 50))
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
	//imagick.Initialize()
	//defer imagick.Terminate()
	lambda.Start(Handler)
}
