package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Terence1105/LineBot/backoffice/cmd"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		cmdString = strings.TrimSuffix(cmdString, "\n")
		arrCommandStr := strings.Fields(cmdString)

		newOsArgs := []string{os.Args[0]}
		newOsArgs = append(newOsArgs, arrCommandStr...)
		os.Args = newOsArgs

		if len(os.Args) == 1 {
			continue
		}

		switch os.Args[1] {
		case "exit":
			os.Exit(1)
		default:
			cmd.Execute()
		}
	}
}
