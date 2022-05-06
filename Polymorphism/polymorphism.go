package Polymorphism

import "fmt"

type Income interface {
	Calculate() int
	Source() string
}
type FixedBilling struct {
	ProjectName string
	BidedAmount int
}
type TimeAndMaterial struct {
	ProjectName string
	NoofHours   int
	HourlyRate  int
}

func (fb FixedBilling) Calculate() int {
	return fb.BidedAmount
}
func (fb FixedBilling) Source() string {
	return fb.ProjectName
}
func (tm TimeAndMaterial) Calculate() int {
	return tm.NoofHours * tm.HourlyRate
}
func (tm TimeAndMaterial) Source() string {
	return tm.ProjectName
}
func CalculateNetIncome(ic []Income) {
	NetIncome := 0
	for _, v := range ic {
		fmt.Printf("Income From %s = $%d\n", v.Source(), v.Calculate())
		NetIncome = NetIncome + v.Calculate()
	}
	fmt.Printf("Net income of organization = $%d\n", NetIncome)
}
