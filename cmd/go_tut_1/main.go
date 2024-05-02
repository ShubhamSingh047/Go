package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main(){
	intNum :=12
	/* idelly below line should give length 1 but len give the size of Theta asci char which is 2*/
	fmt.Println(len("Θ"),"Theta length",)
	
	/* to check the length of string we should use RuneCountInString method imported from utf*/
	fmt.Println(utf8.RuneCountInString("Θ"),"Theta length");
	fmt.Printf("Shubham Go %d!",intNum);
	helperFun();
	 name := sayName();
	 fmt.Println(name);
	 numerator:=11;
	 denominator:=0;
	 remainder,quotient,err:=intDevisor(numerator,denominator);
	 if (err!=nil){
		fmt.Printf(err.Error())
	 }else{
	 	fmt.Printf("for our question is remainder in %d and denominator is %d ",remainder,quotient)
	 }
	// test();
	learningArray();
}

func helperFun(){
	const passVal string="shubham"
	printMe("Hello "+passVal)
}

func printMe(printval string){
	fmt.Println(printval);
}

func sayName()string {
	return "Shubham Singh from say name function !"
}

func intDevisor(numerator int,denominator int)(int,int,error){
	var err error;
	if denominator==0{
		err=errors.New("Cannot devide by zero");
		return 0,0,err;
	}
	remainder:=denominator%numerator;
	quotient:=numerator/denominator;

	return remainder,quotient,err;
}

func test() {
	const name, age = "Kim", 22
	fmt.Printf("%s is %d years old.", name, age)
	s:=fmt.Sprintf("\nhello")
	
	// Sprintf retrun a value which need to be store and the printed !
	fmt.Printf(s);
}

func learningArray(){
	var intArr[3]int32;
	intArr[0]=1
	intArr[1]=2
	intArr[2]=3
	/**
	*! can also intitalise array dynamically just like in java !
	**/

	tempArr:=[...]int{1,2,3,4,5}
	fmt.Println(intArr);
	fmt.Println("using range",intArr[1:3])
	fmt.Println("temp arr",tempArr);

	/**
	*? Slice is wrapper build on Array !
	**/

	var tempSlice []int32=[]int32{1,2,3};
	fmt.Println(tempSlice,"slice");
	tempSlice = append(tempSlice, 4);
	fmt.Println(tempSlice,"slice after append");
}