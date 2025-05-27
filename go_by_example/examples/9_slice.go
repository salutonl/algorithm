package examples

import "fmt"

func Slices() {
	var s1 []int
	fmt.Println("an empty slice s1 :", s1, s1 == nil, len(s1))

	arr := [...]int{0, 1, 2, 3, 4, 5, 6}
	s2 := arr[1:3]
	fmt.Println("len of s2 is ", len(s2), "cap of s2 is ", cap(s2))

	s3 := make([]int, 3)
	fmt.Println("emp s3:", s3, len(s3))

	s3[0] = 10
	s3[1] = 11
	s3[2] = 12

	fmt.Println("emp s3:", s3, len(s3))

	s3 = append(s3, 22)

	fmt.Println("emp s3:", s3, len(s3))
	s4 := make([]int, len(s3))
	copy(s4, s3)
	fmt.Println("emp s4:", s4, len(s4))

	l := s4[:2]
	fmt.Println(l)
	l = s4[1:2]
	fmt.Println(l)
	l = s4[1:]
	fmt.Println(l)

	twoD := make([][]int, 3)
	for i := 0; i < len(twoD); i++ {
		twoD[i] = make([]int, i+1)
		for j := 0; j < len(twoD[i]); j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)

	twoDigital := [][]int{
		{1},
		{1, 2},
		{1, 2, 3},
		{1, 2, 3, 4},
	}

	fmt.Println(twoDigital)
}
