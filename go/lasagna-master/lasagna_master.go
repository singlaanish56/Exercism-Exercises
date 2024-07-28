package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averagePrepTime int) int{
	if(averagePrepTime==0){
		averagePrepTime =2
	}

	return len(layers) * averagePrepTime
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64){

	noodlesVal :=0 
	sauceVal := 0.0

	for _, k := range layers{
		if(k=="sauce"){
			sauceVal+=float64(0.2)
		}else if(k=="noodles"){
			noodlesVal+= 50
		}
	}

	return noodlesVal, sauceVal
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(ingredients []string, yourIndgredients [] string){

	n := len(ingredients)
	n2 := len(yourIndgredients)
	yourIndgredients = yourIndgredients[:n2-1]
	yourIndgredients = append(yourIndgredients, ingredients[n-1])
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, numberOfPortions int) ([]float64){

	extra:=0.0
	if(numberOfPortions%2!=0){
		extra=0.5
	}

	div := float64(numberOfPortions / 2)

	var quantaties []float64

	for _, k := range amounts{
		q := k*div + k*extra
		quantaties = append(quantaties, q)
	}

	return 	quantaties
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
