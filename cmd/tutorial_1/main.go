package main
import (
	"fmt" 
	"errors"
)

func main(){

	name := "Sarah"

	printMe(name)
	// printMe(age)

	var result, remainder, err = divideForMe(uint32(12), uint32(0))

	if(err != nil){
		fmt.Println(err.Error())
	} else if(remainder == 0){
		fmt.Printf("The result is %v", result)
	} else {
		fmt.Printf("The result is %v and the remainder is %v", result, remainder)
	}

	var arr [3]int32 = [3]int32{1,2,3}

	var arr2  = [...]int32{1,2,3,4,}

	arr3 := [...]int32{1,2,3,4,5}

	fmt.Println(len(arr), len(arr2), len(arr3))

	var intSlice = []int32{1,2,3,4,5}

	fmt.Printf("Capacity is %v and length is %v \n", cap(intSlice), len(intSlice))

	intSlice2 := []int32{6,7,8}

	intSlice = append(intSlice, intSlice2...)

	fmt.Printf("Capacity is %v and length is %v \n", cap(intSlice), len(intSlice))

	fmt.Println(arr3)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 0, 8);


	fmt.Printf("Capacity is %v and length is %v \n", cap(intSlice3), len(intSlice3))

	var myMap = map[string]uint8{"Adam": 18, "Eve": 19};
	fmt.Println(myMap)

	fmt.Println(myMap["Eve"])

	for name,age:= range myMap{
		fmt.Println(name,age)
	}

	for index,value := range arr3{
		fmt.Printf("%v %v\n",index, value)
	}
	fmt.Println()
	for index,value := range intSlice{
		fmt.Printf("%v %v\n",index, value)
	}
	fmt.Println()

	myString := "résumé"

	fmt.Println(len(myString))
	fmt.Println(myString)

	for i, v := range myString{
		fmt.Printf("%v %v \n", i,v)
	}

	var myString2 = []rune("résumé")

	fmt.Println(myString2)
	fmt.Println(len(myString2))

		for i, v := range myString2{
		fmt.Printf("%v %v \n", i,v)
	}

}

func printMe(printValue string){
	fmt.Printf(printValue + "\n")
}

func divideForMe(numerator uint32, denominator uint32) (uint32, uint32, error) {

	var err error;
	if(denominator == 0){
		err = errors.New("Cannot divide by zero. Sorry :(")
		return 0,0,err
	}
	var remainder uint32 = numerator%denominator;
	var result uint32 = numerator /denominator

	return result, remainder, err;
}