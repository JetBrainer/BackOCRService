package model

import (
	"log"
	"os/exec"
)

// Runs python script that makes document in vertical position
func CheckVertical(){
	if err := exec.Command("cvrotate.py").Run();err != nil{
		log.Fatal("Error running script!", err)
	}
}