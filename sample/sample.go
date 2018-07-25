package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/containers/psgo"
	"github.com/sirupsen/logrus"
)

func main() {
	var (
		data []string
		err  error
	)

	pid := flag.String("pid", "", "join mount namespace of the process ID")
	format := flag.String("format", "", "ps (1) AIX format comma-separated string")
	list := flag.Bool("list", false, "list all supported descriptors")

	flag.Parse()

	if *list {
		fmt.Println(strings.Join(psgo.ListDescriptors(), ", "))
		return
	}

	if *pid != "" {
		data, err = psgo.JoinNamespaceAndProcessInfo(*pid, *format)
		if err != nil {
			logrus.Panic(err)
		}
	} else {
		data, err = psgo.ProcessInfo(*format)
		if err != nil {
			logrus.Panic(err)
		}
	}

	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
	for _, d := range data {
		fmt.Fprintln(tw, d)
	}
	tw.Flush()
}
