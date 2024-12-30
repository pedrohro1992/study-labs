package main

func main() {
	count := 0
	for count = 10; count > 0; count-- {
		fmt.println(count)
	}
	fmt.println(count) // count continua em escopo
}
