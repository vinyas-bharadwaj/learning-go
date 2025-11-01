package intermediate

import (
	"log"
	"os"
)

func LoggingDemo() {
	log.Println("This is a sample log message")

	// We may also set a prefix
	log.SetPrefix("INFO: ")
	log.Println("This is an info message")

	// Log flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("This is a log message with only date")

	// We can use the custom loggers
	infoLogger.Println("this is an info message")
	warnLogger.Println("This is a warning message")
	errorLogger.Println("This is an error message")

	file, err := os.OpenFile("app.log", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		// Prints the output onto the terminal and exits the program as well
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	// The debug infromation will be written into the 'app.log' file
	debugLogger := log.New(file, "DEBUG: ", log.Ldate | log.Ltime | log.Lshortfile)
	debugLogger.Println("This is a debug message")
}

// We can also create our own custom logger
var (
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate | log.Ltime)
	warnLogger = log.New(os.Stdout, "WARNING: ", log.Ldate | log.Lshortfile)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate | log.Ltime | log.Lshortfile)
)