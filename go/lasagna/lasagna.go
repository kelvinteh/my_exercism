package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime = 40

// TODO: define the 'RemainingOvenTime()' function
func RemainingOvenTime(timeInOven int) int {
	return OvenTime - timeInOven
}

// TODO: define the 'PreparationTime()' function
func PreparationTime(numLayers int) int {
	return numLayers * 2
}

// TODO: define the 'ElapsedTime()' function
func ElapsedTime(layers int, time int) int {
	return PreparationTime(layers) + time
}
