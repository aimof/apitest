package main

import (
	"os"
	"time"

	"github.com/aimof/apitest"
	"github.com/aimof/apitest/comparer/jsoncomparer"
	"github.com/aimof/apitest/kicker"
	"github.com/aimof/apitest/logger"
	"github.com/aimof/apitest/parser/yamlparser"
)

func main() {
	logger.LEVEL = logger.ALL
	if len(os.Args) < 2 {
		logger.Fatal("Not enough args")
		os.Exit(1)
	}
	path := os.Args[1]

	scenarios, err := yamlparser.NewYamlParser().Parse(path)
	if err != nil {
		logger.Fatal(err.Error())
		os.Exit(1)
	}

	go func() {
		time.Sleep(120 * time.Second)
		os.Exit(8)
	}()
	match, err := do(scenarios)
	if err != nil {
		logger.Fatal(err.Error())
		os.Exit(1)
	}
	if !match {
		logger.Error("Not match")
		os.Exit(2)
	}
	os.Exit(0)
}

func do(scenarios []apitest.Scenario) (bool, error) {
	matchAll := true
	for _, s := range scenarios {
		match, err := apitest.Do(s, kicker.NewKicker(), jsoncomparer.NewJSONComparer())
		if err != nil {
			return false, err
		}
		if !match {
			matchAll = false
		}
	}
	return matchAll, nil
}
