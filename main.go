package main

import (
	"butcher/cmd"
	"fmt"
)

var banner = `+---------------------------------------------+
|       ____        __       __               |
|      / __ )__  __/ /______/ /_  ___  _____  |
|     / __  / / / / __/ ___/ __ \/ _ \/ ___/  |
|    / /_/ / /_/ / /_/ /__/ / / /  __/ /      |
|   /_____/\__,_/\__/\___/_/ /_/\___/_/       |
|                                             |
|   VERSION: V1.0.0                           |
|   AUTHOR : zhangchengdong@mallcai.com       |
|                                             |
+---------------------------------------------+
`

func main() {
	fmt.Println(banner)
	cmd.Execute()
}
