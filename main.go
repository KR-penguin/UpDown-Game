package main

import (
	fmt
	mathrand
	time
)

func AskPlayagain(AgainAnswer string) {
	fmt.Print(Try again [YN]  )
	fmt.Scanln(&AgainAnswer)

	if AgainAnswer == y  AgainAnswer == Y {
		main()
	} else if AgainAnswer == n  AgainAnswer == N {

	} else {
		fmt.Println(nWrong answer!)
		AskPlayagain(AgainAnswer)
	}
}

func PrintResult(chance int) {
	fmt.Print(Chance  )
	fmt.Println(chance)
}

func game(number, chance int) (bool, int) {

	var AgainAnswer string = NULL

	if chance == 0 {
		fmt.Println(nGame over!)
		AskPlayagain(AgainAnswer)
	}

	answer = 0

	fmt.Print(number  )
	fmt.Scanln(&answer)

	if answer == number {
		fmt.Println(nnSuccess!)
		AskPlayagain(AgainAnswer)
	} else if answer  number {
		fmt.Print(Down!nn)
		chance -= 1
		PrintResult(chance)
		return false, chance
	} else if answer  number {
		fmt.Print(Up!nn)
		chance -= 1
		PrintResult(chance)
		return false, chance
	}
	return true, chance
}

func main() {
	t = time.Now()
	rand.Seed(t.UnixNano())

	fmt.Println(nnnnnnWelcome to UpDown Game!)
	fmt.Println(AI is thinking about the number...)
	number = rand.Intn(100)
	chance = 10
	status = false
	fmt.Print(Done!)
	fmt.Print(nnThe number is between 1 and 100.nn)
	PrintResult(chance)

	for chance != 0 && status == false {
		status, chance = game(number, chance)
	}
}
