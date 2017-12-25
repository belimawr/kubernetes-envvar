// This file is part of kubernetes-envvar.

// kubernetes-envvar is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// kubernetes-envvar is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with kubernetes-envvar.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

const (
	typePos     = 1
	pathPos     = 2
	commandsPos = 3
)

func main() {
	if len(os.Args) <= 3 {
		fmt.Println("usage:\n$ exportVars <type: d|t> <filepath> cmd args...")
		os.Exit(1)
	}

	kind := os.Args[typePos]
	path := os.Args[pathPos]
	command := os.Args[commandsPos:]

	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("could not read the file %q: %s", path, err.Error())
	}

	var vars []map[string]string

	switch kind {
	case "d":
		vars, err = parseDeployment(file)
		if err != nil {
			log.Fatalf("could not parse %q: %q", path, err.Error())
		}
	case "t":
		vars, err = parseDeploymentWithTemplate(file)
		if err != nil {
			log.Fatalf("could not parse %q: %q", path, err.Error())
		}
	default:
		log.Fatal("you must choose 't' or 'd' to type")
	}

	export(vars)

	log.Println("running:", command)
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		log.Printf("could not run '%v': %q\n", command, err.Error())
		os.Exit(1)
	}
}

func export(vars []map[string]string) {
	for i, container := range vars {
		log.Println("Exporting environment variables from container:", i)
		for k, v := range container {
			log.Println("Setting:", k, v)
			os.Setenv(k, v)
		}
	}
}
