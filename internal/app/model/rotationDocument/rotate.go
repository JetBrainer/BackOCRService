package rotationDocument

import (
	"log"
	"os/exec"
)

// Runs python script that makes document in vertical position
func CheckVertical(filename string) {
	if err := exec.Command("python", "cvrotate.py "+filename).Run(); err != nil {
		log.Fatal("Error running script!", err)
	}
}
