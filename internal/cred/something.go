package cred

import (
	_ "embed"
	"encoding/json"
)

//go:embed data/something.json
var credSomethingJson []byte

type Cred struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Something() (Cred, error) {
	var res Cred

	err := json.Unmarshal(credSomethingJson, &res)
	if err != nil {
		return Cred{}, err
	}

	return res, nil
}
