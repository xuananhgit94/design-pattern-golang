package main

import (
	"designpattern/builderdemo/builder"
	"log"
)

func main() {
	normalBuilder := builder.GetBuilder("normal")
	iglooBuilder := builder.GetBuilder("igloo")

	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()
	log.Println(normalHouse)
	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()
	log.Println(iglooHouse)
}
