package main

import (
	"log"

	"go.uber.org/zap"
)

func main() {
	/*logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("Hello, World!")*/

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	logger.Info("Hello, world!")
}
