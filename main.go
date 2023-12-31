package main

import (
	"fmt"

	"github.com/cli/cli/v2/pkg/cmd/factory"
	"github.com/patrickblackjr/gh-stats/cmd/gh-stats/root"
)

func main() {
	cmdFactory := factory.New("0.1.0")

	cmd := root.NewRootCmd(cmdFactory)
	err := cmd.Execute()

	if err != nil {
		fmt.Println(err)
	}
}
