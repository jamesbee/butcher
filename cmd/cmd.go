package cmd

import (
	"butcher/lib"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"time"
)

var (
	Client *lib.TelnetClient

	globalSleep         int
	globalSleepDuration time.Duration

	rootCmd = &cobra.Command{
		Use:               "Butcher",
		Short:             "A set of Dubbo utils provided by Pudge meant to be helpful.",
		PersistentPreRun:  PreRun,
		PersistentPostRun: PostRun,
	}
)

func InitCmd() {
	initInvoke()

	rootCmd.PersistentFlags().StringP("host", "H", "0.0.0.0", "Dubbo connection host")
	rootCmd.PersistentFlags().IntP("port", "P", 20880, "Dubbo connection port")
	rootCmd.PersistentFlags().IntVarP(&globalSleep, "sleep", "S", 300, "Global command wait duration")
	rootCmd.AddCommand(lsCmd, invokeCmd)

	initConfig()

	globalSleepDuration = time.Duration(globalSleep) * time.Millisecond
}

func PreRun(cmd *cobra.Command, args []string) {
	var err error
	Client, err = lib.Dial(viper.GetString(DubboHost), viper.GetInt(DubboPort))
	if err != nil {
		panic(err)
	}
	Client.Cmd("pwd")
	waitGlobal()
}

func PostRun(cmd *cobra.Command, args []string) {
	Client.Cmd("exit")
	waitGlobal()

	err := Client.Close()
	if err != nil {
		panic(err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func waitGlobal() {
	time.Sleep(globalSleepDuration)
}
