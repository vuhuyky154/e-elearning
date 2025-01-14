package main

import (
	"app/cmd/quizz-service/initialize"
	_ "app/internal/connection"
)

func main() {
	initialize.Run()
}
