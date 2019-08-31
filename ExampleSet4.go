package main

import "fmt"

func main() {
	fmt.Println("Exercise 1")
	var arr[5]int
	j := 0
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Printf("%T", arr)
	fmt.Println()
	for j < len(arr){
		fmt.Println(arr[j])
		j++
	}
	fmt.Println()
	fmt.Println("Exercise 2")
	slc := [] int{1,2,3,4,5,6,7,8,9,10}

	for i,v  := range slc {
		fmt.Println(i,v)
	}
	fmt.Printf("%T", slc)

	fmt.Println()
	fmt.Println("Exercise 3")

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])

	fmt.Println()
	fmt.Println("Exercise 4")

	slcTwo := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	slcTwo = append(slcTwo, 52)
	fmt.Println(slcTwo)
	slcTwo = append(slcTwo, 53, 54, 55)
	fmt.Println(slcTwo)
	y := []int{56, 57, 58, 59, 60}
	slcTwo = append(slcTwo,y...)
	fmt.Println(slcTwo)

	fmt.Println()
	fmt.Println("Exercise 5")
	slcThree := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	slcThree = append(slcThree[:3], slcThree[6:]...)
	fmt.Println(slcThree)

	fmt.Println()
	fmt.Println("Exercise 6")

	m := make([]string, 50, 500)
	fmt.Println(len(m))
	fmt.Println(cap(m))

	for i := 0; i < 50; i++ {
		m[i] = fmt.Sprintf("Position %d", i)
	}

	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))

	m = append(m, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
	)

	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))

	for i := 0; i < len(m); i++ {
		fmt.Println(i, m[i])
	}

	fmt.Println()
	fmt.Println("Exercise 7")

	a :=  []string{"James", "Bond", "Shaken, not stirred"}
	b := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	twoDimentionalSlice := [][]string {a,b}

	fmt.Println(twoDimentionalSlice)

	for i, j := range twoDimentionalSlice{
		fmt.Println("Record: ", i)
		for k, val := range j{
			fmt.Printf("\t index position: %v \t value: \t %v \n", k, val)
		}
	}

	fmt.Println()
	fmt.Println("Exercise 8")

	mTwo := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	for k, v := range mTwo{
		fmt.Println("This is the recorde for" , k)
		for i, v2 := range v{
			fmt.Println("\t", i, v2)
		}
	}

	fmt.Println()
	fmt.Println("Exercise 9")

	mTwo["bayram_beril"] = []string{"Ice Cream", "Kittens", "Dogs"}

	for k,v := range mTwo{
		fmt.Println("This is the record for: ", k)
		for i, v2 := range v{
			fmt.Println("\t",i ,v2)
		}
	}

	fmt.Println()
	fmt.Println("Exercise 10")
	if v, ok := mTwo["bayram_beril"] ; ok{
		fmt.Println("Value: ", v, "\n")
		delete(mTwo, "bayram_beril")
	}

	for k,v := range mTwo{
		fmt.Println("This is the record for: ", k)
		for i, v2 := range v{
			fmt.Println("\t",i ,v2)
		}
	}
}
