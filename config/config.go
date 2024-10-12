package config

import "github.com/spf13/viper"

type Config struct {
	TWITCH_CLIENT_ID           string `mapstructure:"TWITCH_CLIENT_ID"`
	TWITCH_CLIENT_SECRET       string `mapstructure:"TWITCH_CLIENT_SECRET"`
	DISCORD_BOT_TOKEN          string `mapstructure:"DISCORD_BOT_TOKEN"`
	DISCORD_BOT_CHANNEL_ID     string `mapstructure:"DISCORD_BOT_CHANNEL_ID"`
	TWITCH_CHANNELS_TO_MONITOR string `mapstructure:"TWITCH_CHANNELS_TO_MONITOR"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
