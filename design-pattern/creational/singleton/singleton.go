package singleton

import (
	"fmt"
	"sync"
)

type Logger struct {
}

func (l *Logger) Print(message string) {
	fmt.Println("Log: ", message)
}

var (
	loggerInstance *Logger
	once           sync.Once
)

func InitLogger() *Logger {
	once.Do(func() {
		fmt.Println("Initializing Logger...")
		loggerInstance = &Logger{}
	})

	return loggerInstance
}

func GetLogger() *Logger {
	once.Do(func() {
		fmt.Println("Initializing Logger...")
		loggerInstance = &Logger{}
	})
	return loggerInstance
}

func main() {
	l1 := GetLogger()
	l1.Print("This is the first log.")

	l2 := GetLogger()
	l2.Print("This is the second log.")

	fmt.Println("Are l1 and l2 the same?", l1 == l2)
}
