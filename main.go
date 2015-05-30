package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "nowis"
	app.Author = "Naoto Kaneko"
	app.Email = "naoty.k@gmail.com"
	app.Version = "0.1.0"
	app.Usage = "Ask when is now"
	app.Action = nowis
	app.Run(os.Args)
}

func nowis(context *cli.Context) {
	if len(context.Args()) < 1 {
		cli.ShowAppHelp(context)
		os.Exit(1)
	}

	weekday, err := parseWeekday(context.Args()[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	now := time.Now()
	if now.Weekday() == weekday {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func parseWeekday(token string) (time.Weekday, error) {
	normalized := strings.ToLower(token)
	switch normalized {
	case "sunday":
		return time.Sunday, nil
	case "monday":
		return time.Monday, nil
	case "tuesday":
		return time.Tuesday, nil
	case "wednesday":
		return time.Wednesday, nil
	case "thursday":
		return time.Thursday, nil
	case "friday":
		return time.Friday, nil
	case "saturday":
		return time.Saturday, nil
	}

	return time.Sunday, errors.New(fmt.Sprintf("Cannot parse as weekday: %s", token))
}
