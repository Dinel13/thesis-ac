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
	cpu     float64
	mem     float64
	command string
}

func main() {
	output := flag.String("o", "", "output file")
	pid := flag.String("p", "", "pid")
	tresString := flag.String("t", "", "treshold")
	tres := 1.0
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

	f.WriteString("CPU\t\tMEM\t\tTIME\t\t\t\tPID\t\tCOMMAND\n")

	for {
		process := GetProcessInfoUseTop(*pid)
		if process == nil || process.cpu <= tres {
			continue
		}

		go printProcessInfo(process)

		f.WriteString(strconv.FormatFloat(process.cpu, 'f', 2, 64) + "\t\t")
		f.WriteString(strconv.FormatFloat(process.mem, 'f', 2, 64) + "\t\t")
		f.WriteString(time.Now().Format("2006-01-02 15:04:05") + "\t\t")
		f.WriteString(process.pid + "\t\t")
		f.WriteString(process.command + "\n")

		// time.Sleep(500 * time.Millisecond)
	}

}

func GetProcessInfoUsePs(pid string) *Process {
	cmd := exec.Command("ps", "-o", "%cpu,%mem,command", "-p", pid)
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
		// if header , skip it
		if tokens[0] == "%CPU" {
			continue
		}
		ft := make([]string, 0)
		for _, t := range tokens {
			if t != "" && t != "\t" {
				ft = append(ft, t)
			}
		}
		cpu, err := strconv.ParseFloat(ft[0], 64)
		if err != nil {
			log.Fatal("error parsing cpu", err)
		}
		mem, err := strconv.ParseFloat(ft[1], 64)
		if err != nil {
			log.Fatal("error parsing mem", err)
		}
		return &Process{pid, cpu, mem, ft[2]}
	}
	return nil
}

func printProcessInfo(process *Process) {
	log.Printf("pid: %s, cpu: %f, mem: %f, command: %s\n", process.pid, process.cpu, process.mem, process.command)
}

func GetProcessInfoUseTop(pid string) *Process {
	// top -b -n 2 -d 0.2 -p 18368 | tail -1 | awk '{print $9,$10,$12}'
	cmd := exec.Command("top", "-b", "-n", "2", "-d", "0.1", "-p", pid)
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
			if t != "" && t != "\t" {
				ft = append(ft, t)
			}
		}
		if pid != ft[0] {
			continue
		}
		cpu, err := strconv.ParseFloat(ft[8], 64)
		if err != nil {
			log.Fatal(err)
		}
		if cpu == 0 {
			continue
		}
		mem, err := strconv.ParseFloat(ft[9], 64)
		if err != nil {
			log.Fatal(err)
		}
		return &Process{pid, cpu, mem, ft[11]}
	}
	return nil
}
