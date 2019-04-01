package main

import (
	"log"
	"os"

	"github.com/aimof/apitest"
	compselector "github.com/aimof/apitest/comparer"
	"github.com/aimof/apitest/kicker"

	"github.com/aimof/apitest/mapper/yamlmapper"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Not enough args")
	}
	inputpath := os.Args[1]
	tasks, err := yamlmapper.NewYamlMapper().Tasks(inputpath)
	if err != nil {
		log.Fatalln(err)
	}
	results, err := apitest.DoTasks(tasks, kicker.NewKicker(), make(compselector.Selector, 0))
	if err != nil {
		log.Fatalln(err)
	}
	for _, r := range results {
		if !r.Match {
			n, b := r.Got.Got()
			log.Fatalf("%s is fail\ngot: %d %s", r.Name, n, string(b))
		}
	}
}
