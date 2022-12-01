package main

func main() {
	var words = []string{"apple", "pie", "apple", "red", "red", "red"}
	getMostRepeated(words)
}

func getMostRepeated(pArray []string) {
	var items = make(map[string]int)
	for _, item := range pArray {
		_, no := items[item]
		if no {
			items[item] = items[item] + 1
		} else {
			items[item] = 1
		}
	}
	var temp int
	var mostRepeatedWord string
	for x, y := range items {
		if y > temp {
			temp = y
			mostRepeatedWord = x
		}
	}
	println(mostRepeatedWord)

}
