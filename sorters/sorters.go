package sorters

func Bubble(array []int) {
	for i := len(array) - 1; i >= 1; i-- {
		for j := 0; j < i; j++ {
			if array[j] > array[j+1] {
				temp := array[j+1]
				array[j+1] = array[j]
				array[j] = temp
			}
		}
	}
}

func Selection(array []int) {
	for i := 0; i < len(array)-1; i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[min] > array[j] {
				min = j
			}
		}
		if i != min {
			temp := array[i]
			array[i] = array[min]
			array[min] = temp
		}
	}
}

func Insertion(array []int) {
	for j := 1; j < len(array); j++ {
		temp := array[j]
		i := j - 1
		for i >= 0 && temp < array[i] {
			array[i+1] = array[i]
			i--
		}
		array[i+1] = temp
	}
}

func Quick() {

}
