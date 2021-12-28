package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Task 5.
Имеется матрица MxN, состоящая из блоков,
каждый блок (ячейка) имеет какой-либо цвет
Найти наибольшую группу смежных блоков одного цвета

Пример входных данных
M=5 N=5
a a a a a
c b b c a
a a a c a
a b b c a
a a a a a

Программа выведет:
Группа смежных блоков № 1 длина = 17 цвет = a
▮▮▮▮▮
▯▯▯▯▮
▮▮▮▯▮
▮▯▯▯▮
▮▮▮▮▮

Группа смежных блоков № 2 длина = 1 цвет = c
▯▯▯▯▯
▮▯▯▯▯
▯▯▯▯▯
▯▯▯▯▯
▯▯▯▯▯

Группа смежных блоков № 3 длина = 2 цвет = b
▯▯▯▯▯
▯▮▮▯▯
▯▯▯▯▯
▯▯▯▯▯
▯▯▯▯▯

Группа смежных блоков № 4 длина = 3 цвет = c
▯▯▯▯▯
▯▯▯▮▯
▯▯▯▮▯
▯▯▯▮▯
▯▯▯▯▯

Группа смежных блоков № 5 длина = 2 цвет = b
▯▯▯▯▯
▯▯▯▯▯
▯▯▯▯▯
▯▮▮▯▯
▯▯▯▯▯
Итого: наибольшая группа смежных блоков состоит из 17 блоков цвета a
*/

type Block struct {
	color    string
	visited  bool
	group    int
	position Position
}

type Position struct {
	x, y int
}

func Task5() {
	var M, N int

	fmt.Println("Введите M:")
	fmt.Scanf("%d\n", &M)

	fmt.Println("Введите N:")
	fmt.Scanf("%d\n", &N)

	// Ввод матрицы
	fmt.Println("Введите матрицу (цвета, через пробел):")
	matrix := make([][]Block, M)
	for i := 0; i < M; i++ {
		matrix[i] = make([]Block, N)

		in := bufio.NewReader(os.Stdin)
		row, err := in.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения")
			return
		}

		// доп. очистка строки от возврата каретки и новой строки, а также ведущих и конечных пробелов
		row = strings.Trim(row, "\r\n ")
		colors := strings.Split(row, " ")
		if len(colors) != N {
			fmt.Printf("Количество блоков в '%s' не совпадает с %d", row, N)
			return
		}

		for j, color := range colors {
			block := Block{color: color, position: Position{i, j}}
			matrix[i][j] = block
		}
	}

	// Алгоритм: начиная с левого верхнего блока осуществляем поиск смежных, помечаем их, до тех пор, пока не пройдемся по всем
	var head Position
	var groupID int
	var maxBlocksInGroup int
	var colorOfMaxBlocksInGroup string
	for {
		// поиск позиции первого непомеченного блока. В blocs будут храниться смежные блоки
		var blocks []Block
		groupID++
		for i := 0; i < M; i++ {
			if len(blocks) == 0 {
				for j := 0; j < N; j++ {
					if !matrix[i][j].visited {
						matrix[i][j].visited = true
						matrix[i][j].group = groupID
						head = matrix[i][j].position
						blocks = append(blocks, matrix[i][j])
						break
					}
				}
			} else {
				break
			}
		}

		// выход из цикла, когда все блоки помечены
		if len(blocks) == 0 {
			break
		}

		for {
			// проверка соседнего блока справа: не помечен и совпадают цвета
			if head.y < N-1 && matrix[head.x][head.y].color == matrix[head.x][head.y+1].color && !matrix[head.x][head.y+1].visited {
				matrix[head.x][head.y+1].visited = true
				blocks = append(blocks, matrix[head.x][head.y+1])
			}
			// проверка соседнего блока снизу: не помечен и совпадают цвета
			if head.x < M-1 && matrix[head.x][head.y].color == matrix[head.x+1][head.y].color && !matrix[head.x+1][head.y].visited {
				matrix[head.x+1][head.y].visited = true
				blocks = append(blocks, matrix[head.x+1][head.y])
			}
			// проверка соседнего блока слева: не помечен и совпадают цвета
			if head.y > 0 && matrix[head.x][head.y].color == matrix[head.x][head.y-1].color && !matrix[head.x][head.y-1].visited {
				matrix[head.x][head.y-1].visited = true
				blocks = append(blocks, matrix[head.x][head.y-1])
			}
			// проверка соседнего блока сверху: не помечен и совпадают цвета
			if head.x > 0 && matrix[head.x][head.y].color == matrix[head.x-1][head.y].color && !matrix[head.x-1][head.y].visited {
				matrix[head.x-1][head.y].visited = true
				blocks = append(blocks, matrix[head.x-1][head.y])
			}

			unknownGroupBlockFound := true
			for i := 0; i < len(blocks); i++ {
				if blocks[i].group == 0 {
					blocks[i].group = groupID
					matrix[blocks[i].position.x][blocks[i].position.y].group = groupID
					head = blocks[i].position
					unknownGroupBlockFound = false
					break
				}
			}

			// если все блоки в списке смежных блоков проверены - выход их этого цикла
			if unknownGroupBlockFound {
				break
			}
		}

		// данные для ответа
		if len(blocks) > maxBlocksInGroup {
			maxBlocksInGroup = len(blocks)
			colorOfMaxBlocksInGroup = blocks[0].color
		}

		// вывод графического ответа
		fmt.Println()
		fmt.Println("Группа смежных блоков №", groupID, "длина =", len(blocks), "цвет =", blocks[0].color)
		for i := 0; i < M; i++ {
			for j := 0; j < N; j++ {
				findedBlock := false
				for k := 0; k < len(blocks); k++ {
					if blocks[k].position.x == i && blocks[k].position.y == j {
						findedBlock = true
						break
					}
				}

				if findedBlock {
					fmt.Print("▮")
				} else {
					fmt.Print("▯")
				}
			}
			fmt.Println()
		}
	}
	fmt.Printf("Итого: наибольшая группа смежных блоков состоит из %d блоков цвета %s\r\n", maxBlocksInGroup, colorOfMaxBlocksInGroup)
}
