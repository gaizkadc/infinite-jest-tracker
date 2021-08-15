package utils

import (
	"fmt"
)

func CreateCaption(progress int) string {
	adornments := [4]string{"", "jodido ", "puto ", "jodido puto "}

	n := progress % len(adornments)

	caption := fmt.Sprintf("Ya me he leído el %d%% de %s#InfiniteJest.\n", progress, adornments[n])
	caption += createProgressBar(progress)

	return caption
}

func createProgressBar(progress int) string {
	completeChar := "▓"
	incompleteChar := "░"

	numCompleteChars := progress / 5
	numIncompleteChars := 20 - numCompleteChars

	progressBar := ""
	for i := 0; i < numCompleteChars; i++ {
		progressBar += completeChar
	}
	for j := 0; j < numIncompleteChars; j++ {
		progressBar += incompleteChar
	}

	return progressBar
}