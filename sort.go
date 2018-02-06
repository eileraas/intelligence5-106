package main

import "fmt"

func swap(list []int, a, b int){
	tem := list[b]
	list[b]=list[a]
	list[a]=tem
}


func sort(list []int){

	sorted := true
	for sorted {
		sorted=false
		for i := 0; i < len(list) - 1; i++{
			if list[i+1] < list[i]{
				swap(list, i, i+1)
				sorted=true
			}
		}
	}
}

func main(){

		liste := []int{5, 2, 7, 3, 87, 2, 3, 4, 0, 1}
	fmt.Println("usortert:", liste)
	sort(liste)
	fmt.Println("vingardium leviosa!", liste)
}
