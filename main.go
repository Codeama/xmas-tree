package main

import ("fmt"
"github.com/fatih/color"
"time"
)


func main() {
	topSpace := 22
	topStars := 1
	green := color.New(color.Bold, color.FgGreen)
	white := color.New(color.Bold, color.FgWhite)
	fmt.Println()
	fmt.Println()

	for i := 0; i < 5; i++ {
		for j := 0; j < topSpace; j++ {
		fmt.Print(" ")
		}
		topSpace -= 2
		for k := 0; k < topStars; k++ {
			if k == 0 || k == (topStars - 1){
				white.Print("0")
			}else{
			green.Print("*")}
		}
		fmt.Println()
		topStars += 4
	}
	time.Sleep(1 * time.Second)

	midSpace := 18
	midStars := 9
	for i := 0; i < 5; i++ {
		for j := 0; j < midSpace; j++ {
			fmt.Print(" ")
		}
		midSpace -= 3
		for k := 0; k < midStars; k++ {
			if k == 0 || k == (midStars - 1) {
				white.Print("0")
			}else{
			green.Print("*")
		}
		}
		fmt.Println()
		midStars += 6
	}
	time.Sleep(2 * time.Second)

	bottomSpace := 14
	bottomStar := 17
	for i := 0; i < 5; i++ {
		for j := 0; j < bottomSpace; j++ {
			fmt.Print(" ")
		}
		bottomSpace -= 3
		for k := 0; k < bottomStar; k++ {
			if k == 0 || k == (bottomStar - 1){
			white.Print("0")
		}else{
			green.Print("*")
		}
		}
		fmt.Println()
		bottomStar += 6
	}

	trunkSpace := 20
	trunkStar := 6
	red := color.New(color.Bold, color.FgRed)
	for l := 0; l < 3; l++ {
		for j := 0; j < trunkSpace; j++ {
			fmt.Print(" ")
		}

		for k := 0; k < trunkStar; k++ {
			red.Print("*")
	}
		fmt.Println()
	}
	fmt.Println()
	red.Println(".................MERRY CHRISTMAS................")
	fmt.Println()
}

