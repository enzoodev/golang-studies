package math

import "strconv"

func CalculateMedia(values ...float64) float64 {
	total := 0.0

	for _, value := range values {
		total += value
	}

	media := total / float64(len(values))
	roundedMedia, _ := strconv.ParseFloat(strconv.FormatFloat(media, 'f', 2, 64), 64)

	return roundedMedia
}
