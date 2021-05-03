package main

func score(field *[][]int) {
	for i := range *field{
		for j := range (*field)[i]{
			(*field)[i][j]=check(i,j,*field)
		}
	}
}

func check(i, j int, field [][]int) int {
	if field[i][j] == -1 {
		return -1
	}
	a,b := len(field), len(field[0])

	var mines int
	if i!=0 && j!=0 && field[i-1][j-1]==-1{
		mines++
	}
	if j!=0 && field[i][j-1]==-1{
		mines++
	}
	if i+1!=a && j!=0 && field[i+1][j-1]==-1{
		mines++
	}
	if i+1!=a && field[i+1][j]==-1{
		mines++
	}
	if i+1!=a && j+1!=b && field[i+1][j+1]==-1{
		mines++
	}
	if j+1!=b && field[i][j+1]==-1{
		mines++
	}
	if i!=0 && j+1!=b && field[i-1][j+1]==-1{
		mines++
	}
	if i!=0 && field[i-1][j]==-1{
		mines++
	}
	return mines
}