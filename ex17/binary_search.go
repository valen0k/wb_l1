package main

import (
	"errors"
)

//На вход дается отсортированный массив и искомое значение
func BinarySearch(s []int, target int) (int, error) {
	//поиск с помощью двух указателей
	startPoint, endPoint := 0, len(s)-1
	var mid int

	//поиск идет до тех пор, пока начальный указатель меньше конечного
	for startPoint < endPoint {
		//поиск осуществляется с постоянным сравниванием среднего значения массива с искомым значением
		mid = (startPoint + endPoint) / 2
		if s[mid] == target {
			return mid, nil
			//	если среднее значение меньше, то начальный указатель смещается на средний+1 и цикл продолжается
		} else if s[mid] < target {
			startPoint = mid + 1
			//	иначе, конечный указатель становится средним
		} else {
			endPoint = mid
		}
	}
	return 0, errors.New("not found")
}
