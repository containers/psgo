package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/containers/psgo"
)

func main() {
	var (
		descriptors []string
		pidsList    []string
		data        [][]string
		err         error

		pids         = flag.String("pids", "", "comma separated list of process IDs to retrieve")
		format       = flag.String("format", "", "ps(1) AIX format comma-separated string")
		list         = flag.Bool("list", false, "list all supported descriptors")
		join         = flag.Bool("join", false, "join namespace of provided pids (containers)")
		fillMappings = flag.Bool("fill-mappings", false, "fill the UID and GID mappings with the current user namespace")
	)

	log.SetPrefix("psgo: ")
	log.SetFlags(0)

	flag.Parse()

	if *fillMappings && !*join {
		log.Fatal("-fill-mappings requires -join")
	}

	if *list {
		fmt.Println(strings.Join(psgo.ListDescriptors(), ", "))
		return
	}

	if *format != "" {
		descriptors = strings.Split(*format, ",")
	}

	if *pids != "" {
		pidsList = strings.Split(*pids, ",")
	}

	if len(pidsList) > 0 {
		opts := psgo.JoinNamespaceOpts{FillMappings: *fillMappings}

		if *join {
			data, err = psgo.JoinNamespaceAndProcessInfoByPidsWithOptions(pidsList, descriptors, &opts)
		} else {
			data, err = psgo.ProcessInfoByPids(pidsList, descriptors)
		}
		if err != nil {
			log.Panic(err)
		}
	} else {
		data, err = psgo.ProcessInfo(descriptors)
		if err != nil {
			log.Panic(err)
		}
	}

	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
	for _, d := range data {
		fmt.Fprintln(tw, strings.Join(d, "\t"))
	}
	tw.Flush()
}
