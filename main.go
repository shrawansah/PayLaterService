package main

import (
	. "simpl.com/loggers"

	"simpl.com/services/paylater"
)
func main() {

	Logger.Info("App Started")

	simplePaylaterService := paylater.NewSimplePaylaterService()
	simplePaylaterService.StartServing()

}