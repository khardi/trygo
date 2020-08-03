package main

import "fmt"


func main() {


	var i int = 10
	fmt.Println("i var value = ", i)
	fmt.Println("p = ", 		  &i)
	fmt.Println("*p = ", 	      *(&i))
	inc(&i)
	fmt.Println("inc(p), *p = ",  i)


	s := []int{1,2,3}
	fmt.Println("slice s = ",              s)
	incslice(&s)
	fmt.Println("incslice(s), slice s = ", s)


	z := map[string]int{"a": 1}
	fmt.Println("z = ", z)
	incmap(&z)
	fmt.Println("incmap z; z = ", z)

}

func incmap(pm *map[string]int) {
	(*pm)["a"]++
	(*pm)["b"] = 43
}

func incslice(ps *[]int) {
	*ps = append(*ps, 77)
}

func inc(p *int) {
	*p++
}
