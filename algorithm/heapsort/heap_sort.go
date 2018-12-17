package main

import "fmt"

/**
堆排序的时间复杂度为O(NlogN)
最大堆排序
特性: 1. 堆为完全二叉树或近似完全二叉树  2. 堆中，任意父节点大于子节点。可推导出的结论是根节点是最大
实现: 1. 堆构建，每一个父节点都需要大于其子节点，故构建可从最后一个父节点开始，保证父节点大于子节点，再依次往前向上构建，这样就能构建每一个分支节点(即父节点)
      2. 完全排序，构建好的堆中，只能满足部分有序，即任意父节点大于子节点，要排序的话，需要完全有序，结合堆的根节点最大的特性，多次构建堆，将最大值依次取出即可
*/
func buildHeap(array []int) {
	father := (len(array))/2 - 1
	//构建的父节点个数为0到最后一个父节点，倒序构建
	for i := father; i >= 0; i-- {
		//根据父节点和子节点的值，调整保证父节点大于子节点
		fatherMax(array, i, len(array))
	}
}

func fatherMax(array []int, father int, length int) {
	//左子树： 2*father+1， 右子树： 2*father+2，有father就一定有左子树 但不一定有右子树， length为可达长度
	lastFath := length/2 - 1
	if array[father] < array[2*father+1] {
		array[father], array[2*father+1] = array[2*father+1], array[father]
		//如果子节点是父亲，则还需要调整子节点作为父亲的时候，父节点大于子节点
		if 2*father+1 <= lastFath {
			fatherMax(array, 2*father+1, length)
		}
	}
	if 2*father+2 < length && array[father] < array[2*father+2] {
		array[father], array[2*father+2] = array[2*father+2], array[father]
		if 2*father+2 <= lastFath {
			fatherMax(array, 2*father+2, length)
		}
	}

}

func heapSort(array []int) {
	buildHeap(array) //排序 构建一次，即满足父都大于子
	//每次将[0]调出来，换到末尾，剩下的再进行fatherMax即可
	array[0], array[len(array)-1] = array[len(array)-1], array[0]
	for i := 1; i < len(array)-1; i++ {
		fatherMax(array, 0, len(array)-i)
		array[0], array[len(array)-1-i] = array[len(array)-1-i], array[0]
	}
}

func main() {
	a := []int{3, 6, 5, 10, 1, 2, 8}
	heapSort(a)
	fmt.Println(a)

}
