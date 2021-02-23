package pkg

import (
	"encoding/json"
	"fmt"
)

var (
	BuildDescription string = ""
	BuildSHA         string = ""
	BuildVersion     string = ""
	BuildTime        string = ""
)

// Metadata is to store the application info
type Metadata struct {
	Version       string `json:"version"`
	Description   string `json:"description"`
	LastCommitSha string `json:"lastcommitsha"`
}

func GetMetadata() Metadata {
	return Metadata{
		Version:       BuildVersion,
		Description:   BuildDescription,
		LastCommitSha: BuildSHA,
	}
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
