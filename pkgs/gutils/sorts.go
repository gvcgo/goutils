package gutils

type IComparable interface {
	Less(IComparable) bool
}

func partition(itemList []IComparable, left int, right int) int {
	target := itemList[right]
	mid := left
	j := left
	for j < right {
		if itemList[j].Less(target) {
			itemList[mid], itemList[j] = itemList[j], itemList[mid]
			mid++
		}
		j++
	}

	//all in [l,mid) < target && all in (mid,j] > target
	itemList[mid], itemList[right] = itemList[right], itemList[mid]
	return mid
}

func QuickSort(itemList []IComparable, left, right int) {
	if left < right {
		m := partition(itemList, left, right)
		QuickSort(itemList, left, m-1)
		QuickSort(itemList, m+1, right)
	}
}
