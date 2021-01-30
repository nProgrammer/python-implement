package python_implement

import (
	"os"
	"os/exec"
)

type Python struct{}

func (p Python) Execute(file string) {
	cmd := exec.Command("python3", file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func (p Python) ExecuteCode(code string) {
	cmd := exec.Command("python3", "-c", code)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
