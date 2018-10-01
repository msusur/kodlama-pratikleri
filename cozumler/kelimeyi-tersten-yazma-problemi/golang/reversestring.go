package main

func recursive(arr []rune, val string) string {

	lastIndex := (len(arr) - 1) - len(([]rune)(val))

	if lastIndex < 0 {
		return val
	}

	val = val + string(arr[lastIndex])

	return recursive(arr, val)
}

func reverse(val string) string {

	arr := []rune(val)

	return recursive(arr, "")
}

func main() {
	print(reverse("Hello World!!!"))
}
