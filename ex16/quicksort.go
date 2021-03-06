package ex16

/*
Алгоритм состоит из трёх шагов:
	1. Выбрать элемент из массива. Назовём его опорным.
	2. Разбиение: перераспределение элементов в массиве таким образом,
		что элементы, меньшие опорного, помещаются перед ним, а большие или равные - после.
	3. Рекурсивно применить первые два шага к двум подмассивам слева и справа от опорного элемента.
		Рекурсия не применяется к массиву, в котором только один элемент или отсутствуют элементы.
*/
func Quicksort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	//есть два варианта поиска опорного элемента
	//pivot := PartitionLomuto(arr)
	pivot := PartitionHoare(arr)

	Quicksort(arr[:pivot])
	Quicksort(arr[pivot+1:])
}

/*
В данном примере опорным выбирается последний элемент.
Алгоритм хранит индекс в переменной left.
Каждый раз, когда находится элемент,
меньше или равный опорному,
индекс увеличивается,
и элемент вставляется перед опорным.
*/
func PartitionLomuto(arr []int) int {
	left, right := 0, len(arr)-1
	pivot := arr[right]
	for i := 0; i < right; i++ {
		if arr[i] <= pivot {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	return left
}

/*
Данная схема использует два индекса (один в начале массива, другой в конце),
которые приближаются друг к другу, пока не найдётся пара элементов,
где один больше опорного и расположен перед ним, а второй меньше и расположен после.
Эти элементы меняются местами. Обмен происходит до тех пор, пока индексы не пересекутся.
*/
func PartitionHoare(arr []int) int {
	left, right := 0, len(arr)-1
	pivot := arr[right/2]

	for {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left >= right {
			return right
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
}
