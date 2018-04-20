package main

import (
	"bytes"
	"encoding/base64"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"gopkg.in/gographics/imagick.v1/imagick"
	"image"
	"image/color"
	"image/jpeg"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func CreateTestImage() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))
	col := color.RGBA{255, 0, 0, 255}
	for x1 := 30; x1 <= 50; x1++ {
		for y1 := 30; y1 <= 50; y1++ {
			img.Set(x1, y1, col)
		}
	}
	return img
}

func SaveImageAsJpeg(img *image.RGBA) *bytes.Buffer {
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)
	return buf
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//buf := SaveImageAsJpeg(CreateTestImage())

	response, _ := http.Get("https://img.melonjump.com/oH4H9poeMVrGtQtXqnFXyxypLzUBYaBU.jpg")
	b, _ := ioutil.ReadAll(response.Body)

	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()

	if err := mw.ReadImageBlob(b); err != nil {
		panic(err)
	}

	//mw.SepiaToneImage(20)
	mw.NegateImage(false)

	headers := make(map[string]string)
	headers["Content-Type"] = "image/jpeg"

	return events.APIGatewayProxyResponse{
		//Body: base64.StdEncoding.EncodeToString(buf.Bytes()),
		Body:            base64.StdEncoding.EncodeToString(mw.GetImageBlob()),
		StatusCode:      200,
		IsBase64Encoded: true,
		Headers:         headers,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
