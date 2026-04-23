## 1. Сделать из непредсказуемой функции предсказуемую // реализовать функцию обертку

```golang
func main() {
	predicableTimeWork()
}

func unpredictableFunc() int {
    n := rand.Intn(100)
	time.Sleep(time.Duration(n) * time.Second)
    return n
}

func predicableFunc() {
    ...
}

```
<details>
    <summary>Ответ</summary>

```golang
func predicableFunc() {
	c := make(chan int)

	go func() {
		c <- unpredictableFunc()
		close(c)
	}()

	select {
	case v := <-c:
		fmt.Println("Func done with ", v, " seconds")
	case <-time.After(pause):
		fmt.Println("Deadline is exceeded")
	}
}
```

</details>

Решение [тут](ex_1/main.go)


## 2. Сделать из непредсказуемой функции предсказуемую через Context

<details>
    <summary>Ответ</summary>

```golang
func predicableFunc() (int, error) {
	c := make(chan int)
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	go func() {
		c <- unpredictableFunc()
		close(c)
	}()

	select {
	case v := <-c:
		return v, nil
	case <-ctx.Done():
		return 0, errors.New("Deadline is exceeded")
	}
}
```

</details>

Решение [тут](ex_2/main.go)



## 3. Необходимо написать функцию `reader(doubler(writer()))` при условии, что doubler отрабатывает с задержкой 500 мс

<details>
    <summary>Ответ</summary>

```golang
func main() {
	// writer()
	// doubler
	// reader
	reader(doubler(writer()))
}

func writer() <-chan int {
	c := make(chan int)

	go func() {
		for i := range 10 {
			c <- i
		}
		close(c)
	}()

	return c
}

func doubler(c <-chan int) <-chan int {
	result := make(chan int)

	go func() {
		for v := range c {
			time.Sleep(500 * time.Millisecond)
			result <- v * 2
		}
		close(result)
	}()
	return result
}

func reader(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("done")
}
```

</details>


Решение [тут](ex_4/main.go)


## 4. Необходимо реализовать паттерн Write-Behind Cache + Batching. Каждую секунду значение записывается в in-memory хранилище и в очередь. Когда количество значений в очереди достигает 10, тогда in-memory хранилище чистится, а значения из очереди записываются в основное хранилище storage. 

<details>
    <summary>Ответ</summary>

```golang
func main() {
	storage := make([]int, 0)
	buffer := make([]int, 0)
	c := make(chan int, 10)
	rw := sync.RWMutex{}

	go func() {
		var counter int

		for {
			time.Sleep(1 * time.Second)
			counter++
			if counter > 10 {
				counter = 1
			}

			rw.Lock()
			buffer = append(buffer, counter)
			rw.Unlock()

			c <- counter
			fmt.Println("writing buffer... ", counter)
		}
	}()

	go func() {
		for {
			if len(c) == 10 {
				fmt.Println("flushing buffer...")

				tmp := make([]int, 0, 10)
				for i := 0; i < 10; i++ {
					tmp = append(tmp, <-c)
				}

				rw.Lock()
				storage = append(storage, tmp...)
				buffer = buffer[:0]
				rw.Unlock()
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	for {
		time.Sleep(15 * time.Second)

		rw.RLock()
		fmt.Println("read storage after 15 sec")
		fmt.Println("Buffer: ", buffer)
		fmt.Println("Storage: ", storage)
		println("-------------------------")
		rw.RUnlock()
	}
}
```

</details>

Решение [тут](ex_8/main.go)

## 5a. Rate Limiter. Количество коннектов

<details>
    <summary>Ответ</summary>

```golang
func main() {

	requests := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		requests = append(requests, fmt.Sprintf("%d", i*i))
	}

	ch := make(chan string)
	go func() {
		for _, v := range requests {
			ch <- v
		}
		close(ch)
	}()

	// ограничение коннектов
	maxConnects := 10

	wg := sync.WaitGroup{}
	wg.Add(maxConnects)

	for range maxConnects {
		go func() {
			defer wg.Done()

			for req := range ch {
				sendRequest(req)
			}
		}()
	}
	wg.Wait()
}
```

</details>


## 5b. Rate Limiter. Количество коннектов

<details>
    <summary>Ответ</summary>

```golang
func main() {
	// ctx := context.Background()
	// c := client{}

	println("start")

	requests := make([]string, 0, 4)
	for i := 0; i < 4; i++ {
		requests = append(requests, fmt.Sprintf("%d", i*i))
	}
	fmt.Printf("%+v", requests)

	// ограничение коннектов
	maxGoroutines := 2
	tokens := make(chan struct{}, maxGoroutines)

	// заполняем хранилище токенов
	go func() {
		for range maxGoroutines {
			tokens <- struct{}{}
		}
	}()

	for _, req := range requests {
		<-tokens
		go func() {
			defer func() {
				tokens <- struct{}{}
			}()

			// sendRequest(req)
			fmt.Println(req)
		}()
	}

	// необходимо для синхронизации
	for range maxGoroutines {
		<-tokens
	}
}
```

</details>


## 5c. Rate Limiter. Количество rps