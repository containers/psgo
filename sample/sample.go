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
		descriptors []string
		pidsList    []string
		data        [][]string
		err         error
	)

	pid := flag.String("pid", "", "join mount namespace of the process ID")
	pids := flag.String("pids", "", "comma separated list of process IDs to retrieve")
	format := flag.String("format", "", "ps(1) AIX format comma-separated string")
	list := flag.Bool("list", false, "list all supported descriptors")

	flag.Parse()

	if *list {
		fmt.Println(strings.Join(psgo.ListDescriptors(), ", "))
		return
	}

	if *format != "" {
		descriptors = strings.Split(*format, ",")
	}

	if *pid != "" && *pids != "" {
		logrus.Error("you can't pass both --pid and --pids options")
		os.Exit(-1)
	}

	if *pids != "" {
		pidsList = strings.Split(*pids, ",")
	}

	if *pid != "" {
		data, err = psgo.JoinNamespaceAndProcessInfo(*pid, descriptors)
		if err != nil {
			logrus.Panic(err)
		}
	} else if len(pidsList) > 0 {
		data, err = psgo.ProcessInfoByPids(pidsList, descriptors)
		if err != nil {
			logrus.Panic(err)
		}
	} else {
		data, err = psgo.ProcessInfo(descriptors)
		if err != nil {
			logrus.Panic(err)
		}
	}

	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
	for _, d := range data {
		fmt.Fprintln(tw, strings.Join(d, "\t"))
	}
	tw.Flush()
}
