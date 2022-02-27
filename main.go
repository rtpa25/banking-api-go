package main

import (
	"fmt"

	"github.com/rtpa25/banking/app"
	"github.com/rtpa25/banking/logger"
	"github.com/spf13/viper"
)

var ServerURL string

func main() {
	ServerURL = fmt.Sprintf("%s:%s", viper.Get("HOST").(string), viper.Get("PORT").(string))
	logger.Info("Started the application on " + "http://" + ServerURL + " 🚀")
	app.Start(ServerURL)
}
