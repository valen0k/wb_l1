package ex12

func NewSet(words []string) []string {
	//для записи и проверки уникальности слов
	buf := make(map[string]struct{})
	//для вывода результата
	set := make([]string, 0)

	for _, word := range words {
		if _, ok := buf[word]; !ok {
			buf[word] = struct{}{}
			set = append(set, word)
		}
	}

	return set
}
