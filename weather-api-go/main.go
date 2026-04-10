package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type WeatherResult struct {
	City        string `json:"city"`
	Description string `json:"description"`
}

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
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": "Failed to read response",
			})
			return
		}
		var weatherData map[string]interface{}
		if err := json.Unmarshal(body, &weatherData); err != nil {
			ctx.JSON(500, gin.H{
				"error": "Failed to parse body",
			})
		}

		result := WeatherResult{
			City:        fmt.Sprintf("%v", weatherData["resolvedAddress"]),
			Description: fmt.Sprintf("%v", weatherData["description"]),
		}

		rdb.Set(ctx, "test", "hello", time.Hour)
		val, _ := rdb.Get(ctx, "test").Result()

		cached, err := rdb.get(ctx, cacheKey).Result()

		ctx.JSON(200, result)
	})

	r.Run(":8080")
}
