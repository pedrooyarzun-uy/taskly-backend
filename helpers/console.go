package helpers

import (
	"os"
	"os/exec"
)

func ConsoleCleaner() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}
