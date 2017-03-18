package main

// The phases of a dish when steps are executed
type Phase int

const (
	EARLY_PREP Phase = iota
	PREP
	COOK
	REST
	SERVE
)

var phases = [...]string {
	"Early Prep",
	"Preparation",
	"Cooking",
	"Food Rest",
	"Plate and Serve",
}

func (phase Phase) String() string {
	return phases[phase]
}

// One ingredient in a recipe
type Ingredient struct {
	Qty float64 `json:"qty"`
	Measure string `json:"measure"`
	Item string `json:"item"`
}

// One cooking step in a recipe
type Step struct {
	Sequence int `json:"sequence"`
	Instruction string `json:"nstruction"`
	Phase Phase `json:"phase"`
	Duration string `json:"duration"`
}

// One piece of equipment used for recipe
type Appliance struct {
	Qty int `json:"qty"`
	Item string `json:"item"`
}

// A recipe
type Recipe struct {
	Name string `json:"name"`
	ID	string `json:"id"`
	YieldQty float64 `json:"yieldQty"`
	YieldLbl string `json:"yieldLbl"`
	Author	string `json:"author"`
	Ingredients []Ingredient `json:"ingredients"`
	Equipment []Appliance `json:"equipment"`
	Steps []Step `json:"steps"`
	Tags []string `json:"tags"`
	Course string `json:"course"`
}
