package main

import (
	"Inheritance/Inheri"
	"Inheritance/Polymorphism"
)

func main() {
	author1 := Inheri.Author{
		"Naveen",
		"Ramana than",
		"Golang Enthusiast",
	}
	blogPost1 := Inheri.BlogPost{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	blogPost2 := Inheri.BlogPost{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	blogPost3 := Inheri.BlogPost{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}
	w := Inheri.Website{
		BlogPosts: []Inheri.BlogPost{blogPost1, blogPost2, blogPost3},
	}
	w.Contents()
	project1 := Polymorphism.FixedBilling{
		ProjectName: "Project 1",
		BidedAmount: 5000,
	}
	project2 := Polymorphism.FixedBilling{
		ProjectName: "Project 2",
		BidedAmount: 10000,
	}
	project3 := Polymorphism.TimeAndMaterial{
		ProjectName: "Project 3",
		NoofHours:   160, HourlyRate: 25,
	}
	incomeStreams := []Polymorphism.Income{project1, project2, project3}
	Polymorphism.CalculateNetIncome(incomeStreams)
}
