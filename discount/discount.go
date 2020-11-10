package discount

type Discount struct {
	Description string `json:"description"`
	Percentage  int    `json:"percentage"`
}

var BonusDiscount = &Discount{
	Description: "Bonuskorting",
	Percentage:  20,
}

var OutOfDateDiscount = &Discount{
	Description: "Overdatumkorting",
	Percentage:  20,
}
