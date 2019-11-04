package cmd

import (
	"github.com/spf13/viper"
	"os"
)

func initConfig() {
	// read from config
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	// create default config file if not exist
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
			f, err := os.Create("config.yml")
			if err != nil {
				panic(err)
			}
			defer func() { check(f.Close()) }()

			resetDefaultConfig()
			check(viper.WriteConfig())
		} else {
			panic(err)
		}
	}

	// read from command line first
	check(viper.BindPFlag(DubboHost, rootCmd.PersistentFlags().Lookup("host")))
	check(viper.BindPFlag(DubboPort, rootCmd.PersistentFlags().Lookup("port")))
	check(viper.BindPFlag(ButcherSleep, rootCmd.PersistentFlags().Lookup("sleep")))
	check(viper.BindPFlag(InvokeFile, invokeCmd.PersistentFlags().Lookup("file")))
	check(viper.BindPFlag(InvokeSleep, invokeCmd.PersistentFlags().Lookup("sleep")))
}

func resetDefaultConfig() {
	viper.SetDefault(DubboHost, "127.0.0.1")
	viper.SetDefault(DubboPort, 20880)
	viper.SetDefault(ButcherSleep, 300)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
