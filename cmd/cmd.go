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

	rootCmd = &cobra.Command{
		Use:   "Butcher",
		Short: "Butcher is a dubbo batch invoker util.",
	}
)

func init() {
	invokeCmd.PersistentFlags().StringP("file", "F", "", "Invoke from file")
	invokeCmd.PersistentFlags().StringP("sleep", "S", "500", "Sleep after each command execution")

	rootCmd.PersistentFlags().StringP("host", "H", "0.0.0.0", "Dubbo connection host")
	rootCmd.PersistentFlags().IntP("port", "P", 20880, "Dubbo connection port")
	rootCmd.AddCommand(lsCmd, invokeCmd)

	_ = viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	_ = viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	_ = viper.BindPFlag(InvokeFile, invokeCmd.PersistentFlags().Lookup("file"))
	_ = viper.BindPFlag(InvokeSleep, invokeCmd.PersistentFlags().Lookup("sleep"))

	viper.SetDefault("host", "0.0.0.0")
	viper.SetDefault("port", 20880)
	viper.SetDefault(InvokeSleep, 500)
}

func PreRun(cmd *cobra.Command, args []string) {
	var err error
	Client, err = lib.Dial(viper.GetString("host"), viper.GetInt("port"))
	if err != nil {
		panic(err)
	}
	Client.Cmd("pwd")
	time.Sleep(300 * time.Millisecond)
}

func PostRun(cmd *cobra.Command, args []string) {
	time.Sleep(100 * time.Millisecond)
	Client.Cmd("exit")
	time.Sleep(100 * time.Millisecond)
	err := Client.Close()
	if err != nil {
		panic(err)
	}
	time.Sleep(100 * time.Millisecond)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
