package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/distribution/distribution/v3/version"
	"github.com/opencontainers/go-digest"

	_ "crypto/sha256"
	_ "crypto/sha512"
)

var (
	algorithm   = digest.Canonical
	showVersion bool
)

type job struct {
	name   string
	reader io.Reader
}

func init() {
	flag.Var(&algorithm, "a", "select the digest algorithm (shorthand)")
	flag.Var(&algorithm, "algorithm", "select the digest algorithm")
	flag.BoolVar(&showVersion, "version", false, "show the version and exit")

	log.SetFlags(0)
	log.SetPrefix(os.Args[0] + ": ")
}

func usage() { _ = "STUB: not implemented"; return }

func unsupported() { _ = "STUB: not implemented"; return }

func main() {
	var jobs []job

	flag.Usage = usage
	flag.Parse()
	if showVersion {
		version.PrintVersion()
		return
	}

	var fail bool
	if flag.NArg() > 0 {
		for _, path := range flag.Args() {
			fp, err := os.Open(path)
			if err != nil {
				log.Printf("%s: %v", path, err)
				fail = true
				continue
			}
			defer fp.Close()

			jobs = append(jobs, job{name: path, reader: fp})
		}
	} else {

		jobs = append(jobs, job{name: "-", reader: os.Stdin})
	}

	digestFn := algorithm.FromReader

	if !algorithm.Available() {
		unsupported()
	}

	for _, job := range jobs {
		dgst, err := digestFn(job.reader)
		if err != nil {
			log.Printf("%s: %v", job.name, err)
			fail = true
			continue
		}

		fmt.Printf("%v\t%s\n", dgst, job.name)
	}

	if fail {
		os.Exit(1)
	}
}
