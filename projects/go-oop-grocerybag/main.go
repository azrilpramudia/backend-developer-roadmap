package main

import (
	"github.com/azrilpramudia/go-oop-grocerybag/grocerybag"
	"github.com/azrilpramudia/go-oop-grocerybag/labels"
)

func main() {
	label1 := labels.NewLabel("jonh Smith", "BC-1001")
	bag := grocerybag.NewBag([]grocerybag.Item{
		{
			Name: "Apple",
			Cost: 0.5,
		},
		{
			Name: "Banana",
			Cost: 0.25,
		},
		{
			Name: "Orange",
			Cost: 0.75,
		},
		{
			Name: "Watermelon",
			Cost: 1.00,
		},
	}, label1)

	bag.Add("Milk", 2.99)
	bag.PrintList()

	label2 := labels.NewLabel("Jane Smith", "BC-1002")
	bag2 := grocerybag.NewBag([]grocerybag.Item{
		{
			Name: "Apple",
			Cost: 0.5,
		},
	}, label2)
	bag2.PrintList()
}