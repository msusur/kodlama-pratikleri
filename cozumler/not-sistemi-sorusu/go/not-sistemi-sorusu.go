package main

var scores = []int{90, 80, 70, 65, 60, 50, 30, 0}
var codes = []string{"AA", "BA", "BB", "CB", "CC", "DD", "FD", "FF"}

func scoreToResultCode(score int) (result string) {
	for idx, _score := range scores {
		if score >= _score {
			result = codes[idx]
			break
		}
	}
	
	return 
}

func test(scores []int, codes []string) bool {
	for idx, _score := range scores {
		if scoreToResultCode(_score) != codes[idx] { panic("ERROR") }
	}

	return true
}

func main() {
	codesToTest := []int{30, 90, 66}
	scoresToTest := []string{"FD", "AA", "CB"}
	test(codesToTest, scoresToTest)
}
