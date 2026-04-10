package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	apiKey := os.Getenv("WEATHER_API_KEY")

	r.GET("/weather/:city", func(ctx *gin.Context) {
		city := ctx.Param("city")

		url := fmt.Sprintf(
			"https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=metric&key=%s&contentType=json",
			city, apiKey,
		)
		resp, err := http.Get(url)

		if err != nil {
			fmt.Printf("Cant get url: %s", err)
			return
		}

		body, _ := io.ReadAll(resp.Body)
		ctx.Data(200, "application/json", body)

	})

	r.Run(":8080")
}
