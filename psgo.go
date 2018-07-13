package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/containers/psgo/ps"
	"github.com/sirupsen/logrus"
)

func main() {
	var (
		data []string
		err  error
	)

	pid := flag.String("pid", "", "join mount namespace of the process ID")
	format := flag.String("format", "", "ps (1) AIX format comma-separated string")

	flag.Parse()

	if *pid != "" {
		data, err = ps.JoinNamespaceAndProcessInfo(*pid, *format)
		if err != nil {
			logrus.Panic(err)
		}
	} else {
		data, err = ps.ProcessInfo(*format)
		if err != nil {
			logrus.Panic(err)
		}
	}

	tw := tabwriter.NewWriter(os.Stdout, 20, 1, 3, ' ', 0)
	for _, d := range data {
		fmt.Fprintln(tw, d)
	}
	tw.Flush()
}
