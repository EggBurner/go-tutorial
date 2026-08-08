package main

import (
	"errors"
	"fmt"
	"strings"
)

type gasEngine struct{
	kpl uint8
	liters uint8
	owner
}

type electricMotor struct{
	kmpkw uint8
	battery uint16
	owner
}

type car interface{
	kmLeft() uint16
}
func (e gasEngine) kmLeft() uint16{
	return uint16(e.kpl)*uint16(e.liters)
}
func(e electricMotor) kmLeft() uint16{
	return uint16(e.kmpkw)*uint16(e.battery)
}

type owner struct{
	name string
	age uint8
}

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

	var strArr = []rune{'a','b','d','u','l','l','a','h'}
	catStr := "";

	var strBuilder strings.Builder;

	for _,v := range strArr{
		strBuilder.WriteRune(v)
	}

	catStr = strBuilder.String();
	fmt.Println(catStr)
	fmt.Println()

	var myOwner owner = owner{"ali", 19}
	var myEngine gasEngine = gasEngine{13, 35, myOwner}

	var myEV electricMotor = electricMotor{7, 90, myOwner}

	fmt.Printf("Kilometers left for engine car %v \n", myEngine.kmLeft())
	fmt.Printf("Kilometers left for EV car %v \n", myEV.kmLeft())


	canIMakeIt(myEngine, uint16(500))
	canIMakeIt(myEV, uint16(500))
	
}

func canIMakeIt(e car, miles uint16){
	if miles<= e.kmLeft() {
		fmt.Println("Can make it")
	} else{
		fmt.Println("No cant make it")
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