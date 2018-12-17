/*
@Time : 2018/12/14 上午9:26
@Author : xiaoxuez

排序算法
*/

package main

import (
	"fmt"
)

/**
当排序长度小于50时，优于选择插入排序和选择排序
*/

/**
插入排序，保证数组有序，插入数据，时间复杂度为O(n²)
*/
func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		var j = i
		//从后往前，依次后挪，若for循环退出，退出的j就是适合插入的位置
		for j := i; j > 0 && arr[j-1] > temp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
}

/**
选择排序
*/
func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
	}
}

/**
当排序长度较大时，优于选择快排和堆排序
*/

/**
快排，每次将选择第0个作为比较对象
*/
func quickSort(arr []int, start, end int) {
	if end <= start || end == 0 || start == len(arr)-1 {
		return
	}

	flag := arr[start]
	head, tail := start, end
	for i := start + 1; i <= tail; {
		//如果大于的情况，与tail的换，换完i的位置不变，因为tail位置换上来的也是还没比较过的数
		if arr[i] < flag {
			arr[i], arr[head] = arr[head], arr[i]
			head++
			i++
		} else {
			arr[i], arr[tail] = arr[tail], arr[i]
			tail--
		}
	}
	//一轮下来，head == tail == flag应该在的位置
	arr[head] = flag
	quickSort(arr, start, head-1)
	quickSort(arr, head+1, end)
}

func main() {
	arr := []int{12, 35, 4, 23, 13, 20, 4}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
