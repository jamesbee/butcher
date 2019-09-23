package cmd

import (
	"bufio"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

var invokeCmd = &cobra.Command{
	Use:     "invoke",
	Short:   "Invoke given command on dubbo instance.",
	PreRun:  PreRun,
	PostRun: PostRun,
	Run: func(cmd *cobra.Command, args []string) {
		fileName := viper.GetString("invokeFile")
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
			defer file.Close()

			sleep := viper.GetInt("invokeSleep")
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				Client.Cmd(fill(scanner.Text()))
				time.Sleep(time.Duration(sleep) * time.Millisecond)
			}

			if err := scanner.Err(); err != nil {
				panic(err)
			}
		}
	},
}

func fill(cmd string) string {
	if !strings.HasPrefix(cmd, "invoke") {
		cmd = "invoke " + cmd
	}
	return cmd
}