package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var jmeter *string
var mode *string

func run() {
	// get flags
	jmeter = flag.String("j", "/home/din/Downloads/apache-jmeter-5.4.1/bin/jmeter", "jmeter home")
	mode = flag.String("m", "test", "test or convert")
	flag.Parse()
	fmt.Println("jmeter: ", *jmeter)
	fmt.Println("mode: ", *mode)

	// get arguments
	sourceDir := os.Args[3]
	destDir := os.Args[4]
	fmt.Println("sourceDir: ", sourceDir)
	fmt.Println("destDir: ", destDir)

	// read files
	files, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		log.Fatal(err)
	}

	// run jmeter test or convert jtl to html
	if *mode == "test" {
		runJmxTest2Jtl(files, sourceDir, destDir)
	} else if *mode == "convert" {
		convertJtl2html(files, sourceDir, destDir)
	} else {
		log.Fatal("invalid mode")
	}
}

// convertjtl2html convert jtl to html
// it need files, jtl file path, html file path
func convertJtl2html(files []fs.FileInfo, sourceDir string, destDir string) {
	for _, file := range files {
		log.Println("file: ", file.Name())
		if file.IsDir() {
			newFiles, err := ioutil.ReadDir(fmt.Sprintf("%s/%s", sourceDir, file.Name()))
			if err != nil {
				log.Fatal(err)
			}
			convertJtl2html(newFiles, fmt.Sprintf("%s/%s", sourceDir, file.Name()), fmt.Sprintf("%s/%s", destDir, file.Name()))
			continue
		}
		cmd := exec.Command(*jmeter, "-g", fmt.Sprintf("%s/%s", sourceDir, file.Name()), "-o", fmt.Sprintf("%s/%s", destDir, file.Name()))
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}

// runJMeterTest2Jtl run jmeter test and save result to jtl
// it need files, jmx file path, jtl file path
func runJmxTest2Jtl(files []fs.FileInfo, sourceDir string, destDir string) {
	for _, file := range files {
		log.Println("file: ", file.Name())
		if file.IsDir() {
			newFiles, err := ioutil.ReadDir(fmt.Sprintf("%s/%s", sourceDir, file.Name()))
			if err != nil {
				log.Fatal(err)
			}
			runJmxTest2Jtl(newFiles, fmt.Sprintf("%s/%s", sourceDir, file.Name()), fmt.Sprintf("%s/%s", destDir, file.Name()))
			continue
		}

		fileJmx := fmt.Sprintf("%s/%s", sourceDir, file.Name())
		fileReport := file.Name()[len(file.Name())-4:] // remove .jmx
		fileJtl := fmt.Sprintf("%s/%s.jtl", destDir, fileReport)

		cmd := exec.Command(*jmeter, "-n", "-t", fileJmx, "-l", fileJtl)
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}
