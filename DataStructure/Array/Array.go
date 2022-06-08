package Array

type Array struct {
	elemNum int
	arraySize int
	arr []int
}

func newArray(arraySize int) *Array {
	arr := Array{}
	arr.elemNum = 0
	arr.arraySize = arraySize
	arr.arr = make([]int, arr.arraySize)

	for i := 0; i < arraySize; i++ {
		arr.arr[i] = 0
	}

	return &arr
}

func (arr Array) At(idx int) int {
	if idx > arr.arraySize-1 || idx < 0 {
		panic("Array Out of Index!")
	}

	return arr.arr[idx]
}

func (arr *Array) Add(idx, data int) {
	if arr.elemNum == arr.arraySize {
		panic("Array is full!")
	}

	if idx < -1 || idx > arr.arraySize-1 {
		panic("Array Out of Index!")
	}

	if arr.arr[idx] == 0 {
		arr.arr[idx] = data
		arr.elemNum++
	}else {
		for i := arr.elemNum; i > idx; i-- {
			arr.arr[idx] = arr.arr[idx-1]
		}
		arr.elemNum++
		arr.arr[idx] = data
	}
}

func (arr *Array) Set(idx, data int) {
	if idx > arr.arraySize-1 || idx < 0{
		panic("Array Out of Index!")
	}

	arr.arr[idx] = data
}

func (arr *Array) Remove(idx int) {
	if arr.elemNum == 0 {
		panic("Array is empty!")
	}

	if idx < 0 || arr.arraySize-1 < idx {
		panic("Array Out of Index!")
	}

	for i := idx; i < arr.elemNum-1; i++ {
		arr.arr[i] = arr.arr[i+1]
	}

	arr.arr[arr.elemNum-1] = 0
	arr.elemNum--
}

