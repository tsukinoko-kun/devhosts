package main

import (
	"fmt"
	"os"
)

func main() {
	switch len(os.Args) {
	case 1:
		// check for config file in the current directory
		cfgPath, err := FindConfig()
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
			return
		}
		fmt.Println("Config file found:", cfgPath)
		cfg, err := LoadConfig(cfgPath)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
			return
		}
		if err := AddHosts(cfg); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
			return
		}
		<-AwaitCancel()
		if err := ResetHosts(); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
			return
		}
	case 2:
		// config file provided in argument
		cfg, err := LoadConfig(os.Args[1])
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
			return
		}
		if err := AddHosts(cfg); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
			return
		}
		<-AwaitCancel()
		if err := ResetHosts(); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
			return
		}
	default:
		fmt.Println("Too many arguments")
		fmt.Printf("Usage: %s <config>\n", os.Args[0])
	}
}
