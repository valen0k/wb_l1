package ex10

func Fluctuations(nums []float64) map[int][]float64 {
	set := make(map[int][]float64)
	var buf int

	for _, val := range nums {
		//определяем группу
		buf = int(val/10) * 10
		//добавляем в определенную группу
		set[buf] = append(set[buf], val)
	}

	return set
}
