# Задачи

## 1. Что выведется на экран? 
```golang
func main() {
	c := make(chan int)

	go func() {
		fmt.Println(3)
		time.Sleep(5 * time.Second)

		c <- 100

		fmt.Println(4)
		fmt.Println("goroutine done")
	}()

	fmt.Println(1)
	<-c
	fmt.Println(2)

	fmt.Println("main done")
}
```

<details>
    <summary>Ответ</summary>

	1
	3
	4 (опционально)
	goroutine done (опционально)
	2
	main done
	
</details>


## 2. Что выведет код и как его исправить?
```golang
func main() {
    for i := 0; i < 100; i++ {
        go func() {
            fmt.Println(i)
        }()
    }
}
```
<details>
    <summary>Ответ</summary>
	Выведется НЕ от 0 до 100, может вывестиcь всё, что угодно.  

Решение 1: не упорядоченный порядок вывода
```golang
func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}
```

Решение 2: упорядоченный вывод:
```golang
func main() {
    arr := make([]int, 100)
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			arr[i] = i
		}()
	}
	wg.Wait()

	for _, v := range arr {
		fmt.Println(v)
	}
}
```

Решение 3: не оптимальное (через каналы)
```golang
func main() {
    c := make(chan int)

	for i := 0; i < 100; i++ {
		go func() {
			c <- i
		}()
	}

	for range 100 {
		v := <-c
		fmt.Println(v)
	}
}
```
</details>

## 3. Что выведется на экран?
```golang
func worker() <-chan int {
    ch := make(chan int)

    go func() {
        time.Sleep(time.Second)
        close(ch)
    }()
}

func main() {
    start := time.Now()
    _, _ = worker(), worker() (*)

    fmt.Println(time.Since(start))
}
```
<details>
    <summary>Ответ</summary>
	Если в строчке нет операции чтения <-worker(), то код выполнится практически моментально. Если операции присутствуют, тогда операции будут выполнены **последовательно**
</details>


## 4. Что будет выведено на экран?
```golang
func main() {
    c := make(chan int)

    go func() {
	    for i := 0; i < 100; i++ {
			c <- i
    	}
    }()

	for v := range c {
		fmt.Println(v)
	}
}
```
<details>
    <summary>Ответ</summary>
Deadlock, потому что в какой-то момент времени мы повиснем на чтении из канала.  
func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c) // <--
	}()

	for v := range c {
		fmt.Println(v)
	}
}	
</details>



## 5. Что будет выведено на экран? Сколько горутин запуститься одновременно?

[Пример](ex_6/main.go)

```golang
func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(1 * time.Millisecond)
}

func main() {
	// количество логических процессоров на потоке
	runtime.GOMAXPROCS(1)
	MAX_TASKS := 10_000

	wg := &sync.WaitGroup{}
	wg.Add(MAX_TASKS)

	start := time.Now()
	for range MAX_TASKS {
		go worker(wg)
	}

	wg.Wait()
	fmt.Println(time.Since(start))
}
```
<details>
    <summary>Ответ</summary>
	Планировщик не будет ждать миллисекунду и блокировать горутины, а начнет выполнять все горутины и сразу отправит их в слип. Когда они пробудятся, то он  все их завершит. Но так как есть переключение горутин и забирание их из очереди занимает какое-то время, то не 1 мс, а побольше - 10 мс.<br>
	По факту приложение однопоточное, т.к. runtime.GOMAXPROCS(1), но увеличение логических процессоров не даст значительного прироста, т.к. планировщик эффективно считает ресурсы.
</details>


## 6. Что будет выведено на экран?

```golang
func spawnMessages(n int) chan string {
	ch := make(chan string, 1)

	for i := 0; i < n; i++ {
		ch <- fmt.Sprintf("msg %d", i+1)
	}

	return ch
}

func main() {
	n := 10

	for msg := range spawnMessages(n) {
		fmt.Println("received:", msg)
	}
}
```

<details>
    <summary>Ответ</summary>
	Ответ: deadlock. Запишется одно значение в канал, т.к. есть буфер, но в функции `spawnMessages` чтения не происходит, поэтому функция залочит сама себя.  
</details>

Оба решения [тут](ex_7/main.go)


## 7. Что будет выведено на экран?
```golang
func main() {
	ch := make(chan int, 1)

	for i := 0; i < 5; i++ {
		select {
		case val := <-ch:
			fmt.Println(val)
		case ch <- i:
		}
	}
}
```
<details>
    <summary>Ответ</summary>
	0, 2.<br>
	// 0 -> добавление 0 в канал<br>  
	// 1 -> чтение из канала<br>  
	// 2 -> добавление 2 в канал<br>  
	// 3 -> чтение из канала<br>
	// 4 -> добавление в канал<br>    
</details>
