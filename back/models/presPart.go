package models

// PresPart : model for a presPart
type PresPart struct {
	ID    int    `json:"id"`
	Label string `json:"label"`
}

// PresParts : an array of presParts
type PresParts []PresPart

// NewPresPart : add a presPart
func NewPresPart(p *PresPart) {

}
