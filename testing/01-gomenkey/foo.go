package foo

import "fmt"

func logicFunc(a, b int) (int, error) {
	res := insertData(a, b)
	fmt.Println("insertData result:", res)
	sum, err := netWorkFunc(a, b)
	if err != nil {
		return 0, err
	}

	return sum, nil
}

func netWorkFunc(a, b int) (int, error) {
	if a < 0 && b < 0 {
		errmsg := "a<0 && b<0" //gomonkey有bug，函数一定要有栈分配变量，不然mock不住
		return 0, fmt.Errorf("%v", errmsg)
	}

	return a + b, nil
}

func insertData(a, b int) int {
	fmt.Printf("number a %d and b %d are inserted into database\n", a, b)
	return a * b
}
