package utls

import "log"

// QuitOnError quits the program when an error occurs.
// Helper func for small scripts
func QuitOnError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// QuitOnErrorWithMsg quits the program with a custom message,
// when an error occurrs.
func QuitOnErrorWithMsg(e error, msg string) {
	if e != nil {
		log.Fatal(msg)
	}
}

// LogError logs the error to stdout
func LogError(e error) {
	if e != nil {
		log.Println(e)
	}
}
