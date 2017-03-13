package cmd

import (
	"github.com/spf13/cobra"
	"github.com/surajssd/telegrambot/pkg"
	"github.com/spf13/viper"
	"strconv"
	"github.com/Sirupsen/logrus"
)

var RootCmd = &cobra.Command{
	Use:   "telegrambot",
	Short: "telegrambot is used to deploy the bot",
	Run: func(cmd *cobra.Command, args []string) {

		pkg.TOKEN = viper.Get("token").(string)
		pkg.WEBHOOK_URL = viper.Get("webhook_url").(string)
		pkg.NAMES_FILE = viper.Get("names").(string)
		pkg.NOPINGDAYS = viper.Get("nopingdays").(string)

		switch hour := viper.Get("hour").(type) {
		case string:
			var err error
			pkg.HOUR, err = strconv.Atoi(hour)
			if err != nil {
				pkg.HOUR = 12
				logrus.Warningf("Error parsing hour from env HOUR: %v. Using default hour: %d", err, pkg.HOUR)
			}
		case int:
			pkg.HOUR = hour
		}

		switch min := viper.Get("minute").(type) {
		case string:
			var err error
			pkg.MINUTE, err = strconv.Atoi(min)
			if err != nil {
				pkg.MINUTE = 45
				logrus.Warningf("Error parsing minute from env MINUTE: %v. Using default minute: %d", err, pkg.MINUTE)
			}
		case int:
			pkg.MINUTE = min
		}

		// let's do some stuff here
		pkg.StartBot()
	},
}

func init() {
	viper.AutomaticEnv()

	RootCmd.PersistentFlags().StringVar(&pkg.TOKEN, "token", "", "Token of Telegram API server")
	viper.BindPFlag("token", RootCmd.PersistentFlags().Lookup("token"))

	RootCmd.PersistentFlags().StringVar(&pkg.WEBHOOK_URL, "webhook-url", "", "Mention webhook url for telegram server to send push notifications to")
	viper.BindPFlag("webhook_url", RootCmd.PersistentFlags().Lookup("webhook-url"))

	RootCmd.PersistentFlags().StringVar(&pkg.NAMES_FILE, "names-file", "names.yaml", "File with nick names of all users in the group")
	viper.BindPFlag("names", RootCmd.PersistentFlags().Lookup("names-file"))

	RootCmd.PersistentFlags().StringVar(&pkg.NOPINGDAYS, "nopingdays", "Saturday,Sunday", "Don't ping on those days")
	viper.BindPFlag("nopingdays", RootCmd.PersistentFlags().Lookup("nopingdays"))

	RootCmd.PersistentFlags().IntVar(&pkg.HOUR, "hour", 12, "Hour of the day to ping")
	viper.BindPFlag("hour", RootCmd.PersistentFlags().Lookup("hour"))

	RootCmd.PersistentFlags().IntVar(&pkg.MINUTE, "minute", 45, "Minute of the hour to ping")
	viper.BindPFlag("minute", RootCmd.PersistentFlags().Lookup("minute"))
}
