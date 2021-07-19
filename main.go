package main

import (
	. "simpl.com/loggers"

	services "simpl.com/services"
)
func main() {

	Logger.Info("App Started")

	simplePaylaterService := services.NewSimplePaylaterService()
	simplePaylaterService.StartServing()

}