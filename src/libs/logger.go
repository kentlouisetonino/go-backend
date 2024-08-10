package libs

import "log"

func APILogger(message string, statusCode int) {

	// Color the status code based on its value.
	statusColor := Green
	if statusCode >= 400 && statusCode <= 500 {
		statusColor = Red
	}

	// Log the request.
	log.Printf("%s%s%s %s%d%s", Blue, message, Reset, statusColor, statusCode, Reset)
}
