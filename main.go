package main

import "fmt"

type BikeCategory struct {
	Name          string
	Description   string
	Subcategories []BikeCategory
}

func main() {
	// Define the bike taxonomy
	bikeTaxonomy := []BikeCategory{
		{
			Name:        "Mountain Bikes",
			Description: "Off-road bikes for tackling rough terrains.",
			Subcategories: []BikeCategory{
				{
					Name:        "Hardtail",
					Description: "Mountain bikes with front suspension only.",
					Subcategories: []BikeCategory{
						{
							Name:        "Cross-Country",
							Description: "Hardtail bikes optimized for speed and endurance.",
						},
						{
							Name:        "Trail",
							Description: "Hardtail bikes suitable for a variety of trails.",
						},
					},
				},
				{
					Name:        "Full Suspension",
					Description: "Mountain bikes with both front and rear suspension.",
					Subcategories: []BikeCategory{
						{
							Name:        "All-Mountain",
							Description: "Full suspension bikes designed for varied terrain.",
						},
						{
							Name:        "Downhill",
							Description: "Full suspension bikes for fast descents.",
						},
					},
				},
			},
		},
		{
			Name:        "Road Bikes",
			Description: "Bikes designed for speed on paved roads.",
			Subcategories: []BikeCategory{
				{
					Name:        "Race",
					Description: "Lightweight road bikes optimized for racing.",
				},
				{
					Name:        "Endurance",
					Description: "Road bikes designed for comfort on long rides.",
				},
			},
		},
	}

	// Print the bike taxonomy
	printBikeCategories(bikeTaxonomy, 0)
}

func printBikeCategories(categories []BikeCategory, indent int) {
	for _, category := range categories {
		fmt.Printf("%s%s: %s\n", getIndent(indent), category.Name, category.Description)
		printBikeCategories(category.Subcategories, indent+1)
	}
}

func getIndent(indent int) string {
	return fmt.Sprintf("%"+fmt.Sprintf("%ds", indent*4), "")
}
