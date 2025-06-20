package main
import (
	"log"
	"videochat/internal/server"
)


func main(){
	if err := server.Run(); err !=nil{
		log.Fatalln(err.Error())
	}
}