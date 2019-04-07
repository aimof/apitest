package main

import (
	"fmt"
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
	path := os.Args[1]

	tasks, err := yamlmapper.NewYamlMapper().Tasks(path)
	if err != nil {
		log.Fatalln(err)
	}

	err = do(tasks)
	if err != nil {
		log.Fatalln(err)
	}
}

func do(tasks apitest.Tasks) error {

	results, err := apitest.DoTasks(tasks, kicker.NewKicker(), make(compselector.Selector, 0))
	if err != nil {
		return err
	}
	for _, r := range results {
		if !r.Match {
			// Log file is not supported now.
			log.Printf("%s : Fail\tPlaese Read Log.", r.Name)
			return fmt.Errorf("Stopped: %s Failed", r.Name)
		}
	}
	return nil
}
