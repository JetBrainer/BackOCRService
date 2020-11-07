package model

import (
	"log"
	"os/exec"
)

func CheckVertical(){
	cmd := exec.Command("python","cvrotate.py")
	out, err := cmd.Output()
	if err != nil{
		log.Println(err)
	}
	log.Println(string(out))
	if err := exec.Command("cvrotate.py").Run();err != nil{
		log.Fatal("Error running script!", err)
	}
}