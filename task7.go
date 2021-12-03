package main

import (
	"fmt"
	"time"
)

// Task 7. Напишите собственную функцию Sleep, используя time.After
func Sleep(n int) {
	select {
	case <-time.After(time.Second * time.Duration(n)):
	}
}

func main() {
	// Some code
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			// select в бесконечном цикле ждет сообщения из канала. Следующая итерация цикла не выполняется,
			// до тех пор, пока в одном из каналов не появится сообщение
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
				// ИЛИ выполняется case 3 с таймаутом 1с: time After создаёт канал, по которому посылаем метки времени с заданным интервалом.
				//case <-time.After(time.Second):
				//	fmt.Println("timeout")
				// дефолтное значение (в бесконечном цикле) будет выполняться очень быстро: Выполняемые по умолчанию команды исполняются
				// сразу же, если все каналы заняты.
				//default:
				//	fmt.Println("nothing ready")
			}
		}
	}()

	fmt.Println("Таймер 5с запущен")
	// реализация собственного метода Sleep
	Sleep(5)
	fmt.Println("Таймер: 5с прошло")
}
