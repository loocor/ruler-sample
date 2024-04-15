package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/gorules/zen-go"
)

func workingDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		panic("Unable to determine working directory")
	}

	return wd
}

func decisionLoader(key string) ([]byte, error) {
	decisionPath := path.Join(workingDirectory(), "graphs", key)
	decisionContent, err := os.ReadFile(decisionPath)
	if err != nil {
		return nil, errors.New("file not found")
	}

	return decisionContent, nil
}

func main() {
	engine := zen.NewEngine(zen.EngineConfig{
		Loader: decisionLoader,
	})

	input := map[string]any{
		"cart": map[string]any{
			"weight": 10,
		},
		"customer": map[string]any{
			"country": "US",
		},
	}

	response, err := engine.Evaluate("graph.json", input)
	if err != nil {
		println(err.Error())
	}

	jsonStr, err := json.MarshalIndent(response.Result, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
	} else {
		fmt.Println(string(jsonStr))
	}
}
