package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) <= 2 {
		fmt.Println("usage:\n$ exportVars <filepath> cmd args...")
		os.Exit(1)
	}
	path := os.Args[1]
	toRun := os.Args[2:]

	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("could not read the file %q: %s", path, err.Error())
	}

	vars, err := parse(file)
	if err != nil {
		log.Fatalf("could not parse %q: %q", path, err.Error())
	}

	export(vars)

	log.Println("running: ", toRun)
	cmd := exec.Command(toRun[0], toRun[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		log.Printf("could not run '%v': %q\n", toRun, err.Error())
		os.Exit(1)
	}
}

func export(vars []map[string]string) {
	for i, container := range vars {
		log.Println("Exporting environment variables from:", i)
		for k, v := range container {
			log.Println("Setting:", k, v)
			os.Setenv(k, v)
		}
	}
}
