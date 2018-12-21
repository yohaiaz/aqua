package main

import (
	"flag"
	"fmt"
	"myrepo/aqua/sender"
	"myrepo/aqua/services"
	"os"
)

const help = `			Command manual

NAME
	aqua file directory traversal 

SYNOPSIS
	aqua [-hv] [arg]

DESCRIPTION
	aqua traversal is a tool designed for traversing a directory 
	and printing useful stats info on directory

EXAMPLES
	aqua -h		// print help
	aqua -v		// print version
	aqua .		// print stat info for content starting on current directory
	aqua ~/folder1/		// print stat info for content starting on ~/folder1/ directory
	aqua -u http://localhost:12345/files .		// crawl recursively current dir and post files info to specified endpoint
	aqua -u http://localhost:12345/stats		// prints statistics`

func main() {

	v := flag.Bool("v", false, "print aqua version")
	h := flag.Bool("h", false, "for help usage")
	u := flag.String("u", "http://localhost:12345/files", "endpoint listening to post request")

	flag.Parse()

	if flag.NFlag()+len(flag.Args()) == 0 {
		flag.Usage()
		return
	}

	if *v {
		fmt.Println("aqua version 1.0.0")
		return
	}

	if *h {
		fmt.Println(help)
		return
	}

	dir := "."
	if len(flag.Args()) > 0 {
		dir = flag.Args()[0]
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Println(fmt.Errorf("path not valid: %s", err))
		return
	}

	filesChan := Traverse(dir)

	c := services.NewCrawler(filesChan, sender.NewHttp(*u))

	c.Crawl()
}
