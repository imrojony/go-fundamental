package main

import "fmt" 
func main() {


	if firstRank , secondRank := 500,400; firstRank < secondRank {
			fmt.Println("second course is batter than first course")
			if firstRank > 100 {
				fmt.Println("first block greter 100")
			}


		}else if firstRank > secondRank {
			if secondRank > 100 {
				fmt.Println("second block greter 100")

			}
			fmt.Println("first course is doing batter tahn second course")
	}else {
		fmt.Println("first course is epual to second course")
	}
	
}
