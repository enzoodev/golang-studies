package math

import (
	"fmt"
	"strconv"
)

func CalculateMedia(values ...float64) float64 {
	total := 0.0

	for _, value := range values {
		total += value
	}

	media := total / float64(len(values))
	roundedMedia, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", media), 64)

	return roundedMedia
}
