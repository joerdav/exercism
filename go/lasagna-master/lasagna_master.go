package lasagna

const (
	sauce   = "sauce"
	noodles = "noodles"
)

func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	return len(layers) * timePerLayer
}

func Quantities(layers []string) (noodlesCount int, sauceQuantity float64) {
	for i := range layers {
		switch layers[i] {
		case noodles:
			noodlesCount += 50
		case sauce:
			sauceQuantity += .2
		}
	}
	return noodlesCount, sauceQuantity
}

func AddSecretIngredient(friendsLayers, myLayers []string) {
	myLayers[len(myLayers)-1] = friendsLayers[len(friendsLayers)-1]
}

func ScaleRecipe(quantities []float64, targetPortions int) []float64 {
	result := make([]float64, len(quantities))
	for i := range quantities {
		result[i] = quantities[i] / float64(2) * float64(targetPortions)
	}
	return result
}
