package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

//ByAge implements sort.Interface for []Person based on the field Age

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []Person

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].Name < bn[j].Name }

func main() {
	//Primitive
	xi := []int{4, 5, 1, 2, 3}
	xs := []string{"zama", "tom", "batman", "harry", "charlie"}
	fmt.Println("Primitives")
	fmt.Println("Before sorting...")
	fmt.Println(xi)
	fmt.Println(xs)

	fmt.Println("After sorting...")
	sort.Ints(xi)
	sort.Strings(xs)
	fmt.Println(xi)
	fmt.Println(xs)

	fmt.Println("------------")
	//Custom Objects
	fmt.Println("Custom Objects")
	people := []Person{
		{Name: "Zama", Age: 31},
		{Name: "Alice", Age: 43},
		{Name: "Charlie", Age: 40},
		{Name: "Bob", Age: 22},
	}
	fmt.Println(people)
	fmt.Println("---Sort by Age---")
	sort.Sort(ByAge(people))
	fmt.Println(people)
	fmt.Println("---Sort by Name---")
	sort.Sort(ByName(people))
	fmt.Println(people)
}
