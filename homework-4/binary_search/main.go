package main

import "log"

//
// Задача:
// Реализуйте алгоритм бинарного поиска, посчитайте его сложность: https://ru.wikipedia.org/wiki/Двоичный_поиск
//

func BinarySearchInSortedSliceOfInt(slice []int, search int) (success bool, foundSliceIndex int, iterations int) {
	// считаем сложность, сколько итераций потребуется
	iterationsSpent := 1

	// если слайс пуст, или искомое значение меньше первого элемента слайса, или больше последнего
	if len(slice) == 0 || search < slice[0] || search > slice[len(slice)-1] {
		return false, -1, iterationsSpent // возвращаем: не смогли найти
	}

	fromIdx := 0 // ищем от этого индекса
	toIdx := len(slice) - 1 // до этого индекса
	for {
		//log.Printf("DEBUG LOG: searching %v from [%v]=%v, to [%v]=%v ...\n", search, fromIdx, slice[fromIdx], toIdx, slice[toIdx])

		// главное вовремя остановиться, и не искать то, чего найти не получится
		if toIdx-fromIdx <= 1 && search != slice[fromIdx] && search != slice[toIdx] {
			return false, -1, iterationsSpent
		}

		// проверяем, вдруг уже нашли
		if slice[fromIdx]==search { return true, fromIdx, iterationsSpent }
		if slice[toIdx]==search { return true, toIdx, iterationsSpent }
		var idx int = ( fromIdx + toIdx ) / 2 // берём индекс в середине искомого диапазона
		if slice[idx]==search { return true, idx, iterationsSpent }

		// если не нашли, то сдвигаем поиск левее или правее
		if search < slice[idx] {
			toIdx = idx - 1
		} else {
			fromIdx = idx + 1
		}

		// считаем сложность задачи
		iterationsSpent ++
		// а вдруг алгоритм зациклился
		if iterationsSpent > len(slice) { return false, -1, iterationsSpent }
	}
}

func main() {
	slice := []int{0, 1, 3, 3, 3, 5, 6, 7, 8, 9}
	var maxIterationsSpent int = 0
	var totalIterationsSpent int = 0
	for searchValue:=0; searchValue<10; searchValue++ {
		isFound, foundIndex, iterationsSpent := BinarySearchInSortedSliceOfInt(slice, searchValue)
		log.Printf("slice=%v, searchValue=%v, isFound:%v, foundSliceIndex:%v, iterationsSpent:%v\n", slice, searchValue, isFound, foundIndex, iterationsSpent)

		if iterationsSpent > maxIterationsSpent { maxIterationsSpent = iterationsSpent }
		totalIterationsSpent += iterationsSpent
	}
	log.Println("Максимальное количество проходов цикла на одном слайсе:", maxIterationsSpent)
	log.Println("Суммарное количество проходов цикла на 10 слайсах * 10 элементов:", totalIterationsSpent)
}

