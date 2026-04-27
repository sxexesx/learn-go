# Микропаттерны

## 0. Генератор

```golang
func producer() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
}
```

## 1. Worker pool

Ограниченное количества горутин обрабатывает поток данных  

[Пример](note_10/ex_1/main.go)

Когда использовать:
- обработка очереди задач
- контроль нагрузки

```golang
func main() {
	// наполенение данными jobs

	for i := 0; i < limit; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := range jobs {
				results <- j * 2
			}
		}()
	}
	// важно дождаться всех горутин и закрыть канал, в который писали
	go func() {
		wg.Wait()
		close(results)
	}()

	// вывод данных
}
```

## 2. Fan in

Объединение результатов в один канал.

[Пример](note_10/ex_2/main.go)

Когда использовать:
- параллельная обработка данных
- map-reduce

```golang
func main() {
	arr := make([]chan int, 0)
	results := make(chan string)
	// наполнение данными results

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for v := range arr[i] {
				results <- fmt.Sprintf("arr [%d]: %d", i, v)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()
	// вывод данных
}
```

### 3. Fan Out

[Пример](note_10/ex_3/main.go)

```golang
func() {
	jobs := make(chan int, 0)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		// наполнение канала
	}()

	go func() {
		for j := range jobs {
			fmt.Println("worker [1]: ", j)
		}
	}()

	for j := range jobs {
		fmt.Println("worker [2]: ", j)
	}
	wg.Wait()
}
```

### 4. Pipeline

Цепочка стадий, где каждая стадия обрабатывает данные и передает их дальше

Когда использовать:
- потоковая обработка данных
- ETL, парсинг

[Пример](note_10/ex_4/main.go)

```golang
func main() {
	c1 := stage1()
	c2 := stage2(c1)
	for v := range c2 {
		fmt.Println(v)
	}
	fmt.Println("done")
}

func stage1() <-chan int {
	c := make(chan int)

	go func() {
		for i := range 10 {
			c <- i
		}
		close(c)
	}()

	return c
}

func stage2(c <-chan int) <-chan int {
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
```

### 5. Context Cancellation

Централизованная отмена работы горутин

Когда использовать:
- HTTP-запросы
- таймауты
- graceful shutdown

[Пример](note_10/ex_5/main.go)

```golang 
func unpredictableFunc(){ ... }

func main() {
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


### 6. Select + timeout

Ожидание нескольких каналов + таймаут

[Пример](note_10/ex_6/main.go)

```golang
func main() {
	jobs := make(chan struct{})

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			jobs <- struct{}{}
		}
		close(jobs)
	}()

	for {
		select {
		case v, ok := <-jobs:
			if !ok {
				fmt.Println("done")
				return
			}
			fmt.Printf("%+v\n", v)

		case <-time.After(1500 * time.Millisecond):
			fmt.Printf("timeout")
			return
		}
	}
}
```

### 7. Semaphore (ограничение параллелизма)

Ограничить количество одновременно выполняемых задач

[Пример](note_10/ex_7/main.go)

```golang
func main() {
	const limit = 3
	const jobs = 20

	sem := make(chan struct{}, limit)
	wg := sync.WaitGroup{}

	for i := 0; i < jobs; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// занимаем слот симафора
			sem <- struct{}{}

			// освобождаем слот
			defer func() {
				<-sem
			}()

			// do process
		}()
	}
	wg.Wait()
}
```

### 8. Pub/Sub

Один источник -> много подписчиков

Когда исползовать:
- события
- уведомления
- реактивные системы

[Пример](note_10/ex_8/main.go)

```golang
type Broker struct {
	mu   sync.RWMutex
	subs []chan string
}

func NewBroker() *Broker {
	return &Broker{
		subs: make([]chan string, 0),
	}
}

func (b *Broker) Subscribe() <-chan string {
	c := make(chan string, 5)

	b.mu.Lock()
	b.subs = append(b.subs, c)
	b.mu.Unlock()

	return c
}

// публикация всем подписчикам
func (b *Broker) Publish(msg string) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, ch := range b.subs {
		ch <- msg
	}
}

func (b *Broker) Close() {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, ch := range b.subs {
		close(ch)
	}
}

func main() {
	broker := NewBroker()

	sub1 := broker.Subscribe()
	sub2 := broker.Subscribe()

	go func() {
		for msg := range sub1 {
			fmt.Println("subscriber 1: ", msg)
		}
	}()

	go func() {
		for msg := range sub2 {
			fmt.Println("subscriber 2: ", msg)
		}
	}()

	broker.Publish("order created")
	broker.Publish("payment receive")

	time.Sleep(time.Second)

	broker.Close()
}
```

### 9. Rate limiting

[Пример](note_10/ex_9/main.go)

Контролировать частоту операций

```golang
func main() {
	jobs := make(chan int)

	go func() {
		defer close(jobs)

		for i := 0; i < 10; i++ {
			jobs <- i
		}
	}()

	ticker := time.Tick(1 * time.Second)

	for j := range jobs {
		<-ticker
		fmt.Println("process job ", j, time.Now().Format("15:04:05.000"))
	}
	fmt.Println("done")
}
```