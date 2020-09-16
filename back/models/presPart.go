package models

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Thomgrus/share-your-ci/back/config"
)

// PresPart : model for a presPart
type PresPart struct {
	Code  string `json:"code"`
	Label string `json:"label"`
}

// PresParts : an array of presParts
type PresParts []PresPart

// NewPresPart : add a presPart
func NewPresPart(p *PresPart) {
	json, err := json.Marshal(*p)
	if err != nil {
		fmt.Println(err)
	}
	config.Client().Set(p.Code, json, 0)
	err = config.Client().Set(p.Code, json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}

// FindPresPartByCode : find presPart by code
func FindPresPartByCode(code string) *PresPart {
	var p PresPart

	val, err := config.Client().Get(code).Result()
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal([]byte(val), &p)

	if err != nil {
		log.Println(err)
	}

	return &p
}

func InitPresPartData() {
	var dkrPresPart = PresPart{Code: "DKR", Label: "Docker"}
	NewPresPart(&dkrPresPart)
	fmt.Println(fmt.Sprintf("PresPart %s added", *FindPresPartByCode("DKR")))
}
