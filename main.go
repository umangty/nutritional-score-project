package main

import "fmt"

func main() {
	ns := GetNutritionalScore(NutritionalData{
		Energy: EnergyFromKcal(100),
		// negative
		Sugars:              SugarGram(10),
		SaturatedFattyAcids: SaturatedFattyAcids(2),
		Sodium:              SodiumMilligram(500),
		// positve
		Fruits:  FruitsPercent(60),
		Fibre:   FibreGram(4),
		Protein: ProetinGram(2),
	}, Food)
	fmt.Printf("Nutritional score: %d\n", ns.Value)
	fmt.Printf("Nutri score: %s\n", ns.GetNutriScore())
}
