package apt

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

var (
	aptGetPath string
)

func init() {
	aptGetP, err := exec.LookPath("apt-get")
	if err != nil {
		log.Fatalln("apt-get doesn't exists")
	}
	aptGetPath = aptGetP
}

/*
example use:
	packages := []string{"curl"}
	apt.Install(packages)
*/
func Install(packages []string) {
	args := []string{aptGetPath, "install"}
	args = append(args, packages...)

	cmd := exec.Command("sudo", args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalln(err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalln(err)
	}

	cmd.Stdin = os.Stdin

	err = cmd.Start()
	if err != nil {
		log.Fatalln(err)
	}

	read(stdout)
	read(stderr)
}

func read(reader io.Reader) {
	buf := bufio.NewReader(reader)
	for {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(line))
	}
}
