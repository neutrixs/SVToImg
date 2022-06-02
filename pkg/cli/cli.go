package cli

import (
	"flag"
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/neutrixs/SVToImg/pkg/api"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func Run() int {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println(help())
		return 0
	}

	panoid := flag.String("p", "", "panoid")
	URL := flag.String("u", "", "URL")
	output := flag.String("o", "", "Path to output JPEG file")
	quality := flag.Int("q", 80, "Output JPEG Quality")

	flag.Parse()

	// if they're both specified or they're both not specified
	if (*panoid != "") == (*URL != "") {
		fmt.Println("You must specify exactly 1 source!")
		return 1
	}

	if *output == "" {
		fmt.Println("You must specify output path!")
		return 1
	}

	if *panoid == "" {
		newPanoid, err := api.ShortlinkToPanoid(*URL)

		if err != nil {
			log.Fatal(err)
		}

		*panoid = newPanoid
	}

	img, err := api.GetImages(*panoid)
	
	if err != nil {
		log.Fatal(err)
	}

	wd, _ := os.Getwd()
	splittedPath := strings.Split(*output, "/"); splittedPath = splittedPath[:len(splittedPath) - 1]
	
	pathDirOnly := strings.Join(splittedPath, "/")

	err = os.MkdirAll(filepath.Join(wd, pathDirOnly), os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(*output)

	if err != nil {
		log.Fatal(err)
	}

	jpeg.Encode(file, img, &jpeg.Options{Quality: *quality})

	return 0
}