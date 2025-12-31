package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgMin int) int {
    if avgMin == 0 {
        avgMin = 2
    }

    return len(layers) * avgMin
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    noodle := 0
    sauce := 0
    for _, v := range layers {
        if v == "noodles" {
            noodle += 1
        } else if v == "sauce" {
            sauce += 1
        }
    }

    return noodle * 50, float64(sauce) * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
    if myList[len(myList)-1] == "?" {
        myList[len(myList)-1] = friendsList[len(friendsList)-1]
    } else if friendsList[len(friendsList)-1] == "?" {
        friendsList[len(friendsList)-1] = myList[len(myList)-1]
    }
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(recipe []float64, q int) []float64 {
    scaled := make([]float64, len(recipe))
    copy(scaled, recipe)
    for i := range scaled {
        scaled[i] *= float64(q) / float64(2)
    }

    return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
