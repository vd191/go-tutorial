package main

import "fmt"

// Go has pointers. A pointer holds the memory address of a value.
func main() {

	// Declare a variable x = 10
	x := 10
	// Call changeValue() function for changing value of x to 0
	changeValue(x)
	// Print out x variable
	fmt.Println(x) // x = 10
	// -> x = 10 because the value of variable x is referenced to a different
	// x variable of changeValue() function. that why that function didn't affect to
	// the real x variable.

	// --------------------------------------------------------------------
	y := 20
	changeValueWithPointer(&y)
	fmt.Println(y)

	// --------------------------------------------------------------------
	z := new(int)
	changeValueWithNewKeyword(z)
	fmt.Println(*z)

	// --------------------------------------------------------------------
	a := 1
	// & to get the memory adrress of a.
	b := &a
	fmt.Println(b) // 0xc00006e018

	// * to get the value of pointer variable (it is b)
	fmt.Println(*b) // 1

}

func changeValue(x int) {
	x = 0
}

func changeValueWithPointer(y *int) {
	*y = 0
}

func changeValueWithNewKeyword(z *int) {
	*z = 0
}

/*
	CONCLUSION:
	khi truyền 1 biến vào 1 hàm thì trương trình sẽ tạo 1 biến reference của biến ấy
	chứ không trực tiếp sử dụng biến. Bởi vậy, khi thực hiện thay đổi giá trị thì biến
	không thay đổi. Đó là lý do cần sử dụng Pointer.
	* : lấy giá trị của biến
	& : lấy địa chỉ memory của biến
*/
