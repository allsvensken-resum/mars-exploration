package main

func main() {
	moveMap := make(map[string][2]int)
	moveMap["L"] = [2]int{0, 1}
	moveMap["R"] = [2]int{0, -1}
	moveMap["F"] = [2]int{1, 0}

}
