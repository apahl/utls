package utls

import "log"

// QuitOnError quits the program when an error occurs.
// Helper func for small scripts
func QuitOnError(e error) {
	if e != nil {
		// log.Fatalf("PERR: %+v", e)
		log.Panic("PERR: ", e)
	}
}

// QuitOnErrorWithMsg quits the program with a custom message,
// when an error occurrs.
func QuitOnErrorWithMsg(e error, msg string) {
	if e != nil {
		log.Fatal(msg)
	}
}

// PanicOnError quits the program when an error occurs.
// Calls log.Panic() which provides a stack trace.
// Helper func for small scripts
func PanicOnError(e error) {
	if e != nil {
		log.Panic("PANIC: ", e)
	}
}

// PanicOnErrorWithMsg quits the program with a custom message,
// when an error occurrs.
// Calls log.Panic() which provides a stack trace.
func PanicOnErrorWithMsg(e error, msg string) {
	if e != nil {
		log.Panic(msg)
	}
}

// LogError logs the error to stdout
func LogError(e error) {
	if e != nil {
		log.Println(e)
	}
}
