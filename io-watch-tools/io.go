package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Process struct {
	pid     string
	read    float64
	write   float64
	command string
}

func main() {
	output := flag.String("o", "", "output file")
	pid := flag.String("p", "", "pid")
	tresString := flag.String("t", "", "treshold")
	tres := 0.0
	flag.Parse()
	if *pid == "" {
		log.Fatal("pid is empty")
	}
	if *output == "" {
		log.Fatal("output file is empty")
	}
	if *tresString != "" {
		tres, _ = strconv.ParseFloat(*tresString, 64)
	}
	f, err := os.Create(*output)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.WriteString("Read(kB/s)\t\tWrite(kB/s)\t\tTIME\t\t\t\tPID\t\tCOMMAND\n")

	for {
		process := GetProcessInfoUseTop(*pid)
		if process == nil || process.read <= tres {
			continue
		}

		printProcessInfo(process)

		f.WriteString(strconv.FormatFloat(process.read, 'f', 2, 64) + "\t\t")
		f.WriteString(strconv.FormatFloat(process.write, 'f', 2, 64) + "\t\t")
		f.WriteString(time.Now().Format("2006-01-02 15:04:05") + "\t\t")
		f.WriteString(process.pid + "\t\t")
		f.WriteString(process.command + "\n")

		time.Sleep(1000 * time.Millisecond)
	}

}

func printProcessInfo(process *Process) {
	log.Printf("pid: %s, read(kB/s): %f, write(kB/s): %f, command: %s\n", process.pid, process.read, process.write, process.command)
}

func GetProcessInfoUseTop(pid string) *Process {
	cmd := exec.Command("pidstat", "-dl")
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + errb.String())
	}
	for {
		line, err := outb.ReadString('\n')
		if err != nil {
			break
		}
		tokens := strings.Split(line, " ")
		ft := make([]string, 0)
		for _, t := range tokens {
			t = strings.TrimSpace(t)
			if t != "" && t != "\t" {
				ft = append(ft, t)
			}
		}

		if len(ft) < 4 || ft[3] != pid {
			continue
		}

		read, err := strconv.ParseFloat(ft[4], 64)
		if err != nil {
			log.Fatal(err)
		}

		write, err := strconv.ParseFloat(ft[5], 64)
		if err != nil {
			log.Fatal(err)
		}

		if read == 0 && write == 0 {
			continue
		}
		return &Process{pid, read, write, ft[8]}
	}
	return nil
}
