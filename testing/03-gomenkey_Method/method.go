package method

import "fmt"

type myType struct {
}

func (m *myType) logicFunc(a, b int) (int, error) {
	sum, err := m.NetWorkFunc(a, b)
	if err != nil {
		return 0, err
	}
	return sum, nil
}

func (m *myType) NetWorkFunc(a, b int) (int, error) {
	if a < 0 && b < 0 {
		errmsg := "a<0 && b<0"
		return 0, fmt.Errorf("%v", errmsg)
	}

	return a + b, nil
}
