package main

import "fmt"

var a, b int

func main() {
	fmt.Println("Введите размеры поля")
	_, _ = fmt.Scanf("%d ", &a)
	_, _ = fmt.Scanf("%d ", &b)

	field := make([][]int, a)

	for i := range field {
		field[i] = make([]int, b)
		for j := range field[i] {
			field[i][j] = 0
		}
	}

	kol := calculate(a, b)

	arrange(a, b, kol, &field)

	score(&field)

	println(&field)
}

func println(field *[][]int)  {
	b:=len((*field)[0])
	fmt.Print("_")
	for i:=0;i<b;i++{
		fmt.Print("_")
	}
	fmt.Print("_")
	for i:=range *field {
		fmt.Println()
		fmt.Print("|")
		for j:=range (*field)[i]{
			if (*field)[i][j]==0{
				fmt.Print(" ")
			} else if (*field)[i][j]==-1{
				fmt.Print("*")
			} else {
				fmt.Print((*field)[i][j])
			}
		}
		fmt.Print("|")
	}
	fmt.Println()
	fmt.Print("-")
	for i:=0;i<b;i++{
		fmt.Print("-")
	}
	fmt.Print("-")
}
