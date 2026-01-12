package chance
import "math/rand"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	n := rand.Intn(20) + 1
    return n
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	n := rand.Float64()*12.0
    return n
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	anims := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}

    rand.Shuffle(len(anims), func(i, j int){
        anims[i], anims[j] = anims[j], anims[i]
    })

    return anims
}
