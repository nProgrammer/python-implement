package python_implement

import (
	"log"
	"os"
	"os/exec"
)

type Python struct{}

func (p Python) Execute(file string) {
	cmd := exec.Command("python3", file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println(cmd.Run())
}
