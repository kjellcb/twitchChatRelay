package irc

import (
	"github.com/spf13/viper"
)

func ReadConfigFile() error {
	viper.SetConfigName("twitch_settings")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	Bot.Cred.User = viper.GetString("Twitch.User")
	Bot.Cred.Token = viper.GetString("Twitch.Token")
	Bot.Relay = viper.GetString("Twitch.Relay")
	Bot.Chan = viper.GetStringSlice("Twitch.Channels")
	return nil
}
