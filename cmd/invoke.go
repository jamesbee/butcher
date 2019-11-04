package cmd

import (
	"bufio"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

const (
	DubboHost    = "dubbo.host"
	DubboPort    = "dubbo.port"
	ButcherSleep = "butcher.sleep"
	InvokeSleep  = "butcher.invoke.sleep"
	InvokeFile   = "butcher.invoke.file"
)

var invokeCmd = &cobra.Command{
	Use:   "invoke",
	Short: "Invoke given command on dubbo instance.",
	Run: func(cmd *cobra.Command, args []string) {
		fileName := viper.GetString(InvokeFile)
		if fileName == "" {
			// 单次执行
			if len(args) == 0 {
				return
			}
			Client.Cmd(fill(args[0]))
		} else {
			file, err := os.Open(fileName)
			if err != nil {
				panic(err)
			}
			defer func() { check(file.Close()) }()

			sleepDuration := globalSleepDuration
			sleep := viper.GetInt(InvokeSleep)
			if sleep > 0 {
				sleepDuration = time.Duration(sleep) * time.Millisecond
			}

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				Client.Cmd(fill(scanner.Text()))
				time.Sleep(sleepDuration)
			}

			if err := scanner.Err(); err != nil {
				panic(err)
			}
		}
	},
}

func initInvoke() {
	invokeCmd.PersistentFlags().StringP("file", "F", "", "Invoke from file")
	invokeCmd.PersistentFlags().IntP("sleep", "S", 0, "Sleep after each command execution")
}

func fill(cmd string) string {
	if !strings.HasPrefix(cmd, "invoke") {
		cmd = "invoke " + cmd
	}
	return cmd
}
