package main

import (
	"github.com/Abdulrahman-02/Go-bank/internal/api"
)

func main(){
	server := api.NewAPIserver(":3000")
	server.Run()
}