package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

var jmeter *string
var mode *string

func main() {
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

	// run jmeter test or convert jtl to html
	if *mode == "test" {
		runJmxTest2Jtl(sourceDir, destDir)
	} else if *mode == "convert" {
		// read files
		files, err := ioutil.ReadDir(sourceDir)
		if err != nil {
			log.Fatal(err)
		}
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
func runJmxTest2Jtl(file string, destDir string) {
	fileSave := file[:len(file)-4] // remove .jmx mislanya login

	// loop 10 times
	for i := 0; i < 10; i++ {
		fmt.Println("iterasi ke-: ", i+1)
		cmd := exec.Command(*jmeter, "-n", "-t", file, "-l", fmt.Sprintf("%s/%s%d.jtl", destDir, fileSave, i+1))
		// show the result
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("iterasi ke-: ", i+1, " selesai")
		fmt.Println("sleep for 3 seconds")
		time.Sleep(time.Second * 2)
		fmt.Println("-------------------------------------------------------")
		fmt.Println("")
	}

}
