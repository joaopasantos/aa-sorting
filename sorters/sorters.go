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

func Selection() {

}

func Insertion() {

}

func Quick() {

}
