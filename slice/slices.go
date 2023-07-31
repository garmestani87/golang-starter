package main

import "fmt"

func main() {

	//var m_slice = []int{1, 2, 3}
	//fmt.Printf("len : %d , cap : %d ", len(m_slice), cap(m_slice))
	//
	//var m_arr = [5]int{1, 2, 3, 4, 5}
	//var m_slice = m_arr[:]
	//var m_slice2 = m_arr[2:]
	//var m_slice3 = m_arr[2:3]
	//
	//fmt.Printf("values : %v \n", m_slice)
	//fmt.Printf("values : %v \n", m_slice2)
	//fmt.Printf("values : %v \n", m_slice3)
	//
	//var m_slice4 = make([]int, 5, 8)
	//fmt.Printf("values : %d ", m_slice4)
	//fmt.Printf("len : %d ,  cap : %d ", len(m_slice4), cap(m_slice4))
	//
	//m_slice4 = append(append(append(append(append(m_slice4, 2), 3), 5), 6), 9)
	//
	//fmt.Printf("values %v \n", m_slice4)
	//fmt.Printf("len : %d , cap : %d ", len(m_slice4), cap(m_slice4))
	//
	var m_arr = [2]int{1, 2}
	var m_slice5 = m_arr[:]
	m_arr[1] = 10
	fmt.Printf("values : %v \n", m_slice5)
	fmt.Printf("values : %v \n", m_arr)

	m_slice5 = append(m_slice5, 3)
	fmt.Printf("values : %v \n", m_slice5)
	fmt.Printf("values : %v \n", m_arr)

	m_arr[1] = 15
	fmt.Printf("values : %v \n", m_slice5)
	fmt.Printf("values : %v \n", m_arr)

}
