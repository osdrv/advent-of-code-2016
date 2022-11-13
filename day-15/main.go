package main

/*

	p1 = (t+5)%5
	p2 = (t+3)%2

	p1 = (15+1+t)%17
	p2 = (2+2+t)%3
	p3 = (4+3+t)%19
	p4 = (2+4+t)%13
	p5 = (2+5+t)%7
	p6 = (0+6+t)%5

*/

func solve(input [][2]int) int {
	t := 0
Tick:
	for {
		prev := (input[0][1] + 1 + t) % input[0][0]
		for ix, i := range input {
			if v := (i[1] + ix + 1 + t) % i[0]; v != prev {
				t++
				continue Tick
			}
		}
		return t
	}
}

func main() {
	//input := [][2]int{
	//	{5, 4},
	//	{2, 1},
	//}

	input := [][2]int{
		{17, 15},
		{3, 2},
		{19, 4},
		{13, 2},
		{7, 2},
		{5, 0},
		{11, 0},
	}

	res := solve(input)
	printf("solved: %d", res)
}
