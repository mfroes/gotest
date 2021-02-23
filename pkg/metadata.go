package pkg

import (
	"encoding/json"
	"fmt"
)

// Metadata is to store the application info
type Metadata struct {
	Version       string `json:"version"`
	Description   string `json:"description"`
	LastCommitSha string `json:"lastcommitsha"`
}

// ToJSON is just a example of fucntion to unit test
func (meta *Metadata) ToJSON() string {
	e, err := json.Marshal(meta)
	if err != nil {
		fmt.Println(err)
		return "Err"
	}
	return string(e)
}

// "myapplication": [
// 	{
// 		"version": "1.0",
// 		"description": "pre-interview technical test",
// 		"lastcommitsha": "abc57858585"
// 	}
// ]
