# Микропаттерны

## Генератор

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

## WorkerPool

```golang
var max = 10

func (c client) Proc(ctx context.Context, ch chan int) {
	
	for range max {
		go func() {
			for v := range ch {
				c.DoSth(ctx, v)
			}
		}()

	}
}
```

## Fanin
Слияние нескольких каналов в один

[Пример](note_b2/ex_1/main.go)

```golang
func fanIn(chans ...<-chan int) <-chan int {
	result := make(chan int)
	wg := &sync.WaitGroup{}
	
	go func() {
		for _, ch := range chans {
			wg.Add(1)

			go func() {
				defer wg.Done()
				for val := range ch {
					result <- val
				}
			}()
		}

		wg.Wait()
		close(result)
	}()

	return results
}
```

## Сбор данных из нескольких источников

[Пример](note_b2/ex_3/main.go)

```golang
func main() {
	wg := new(sync.WaitGroup)
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			ch <- i
		}()
	}

	go func() {
		wg.Wait()

		close(ch)
	}()

	<-ch
}
```

## Раздача данных для обработки нескольким воркерам 

[Пример](note_b2/ex_2/main.go)

```golang
func main() {
	ch := make(chan int)

	go func() {
		for i := range 1000 {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for v := range ch {
			fmt.Println("v = ", v, "worker1")
		}
	}()

	for v := range ch {
		fmt.Println("v = ", v, "worker2")
	}
}
```

## Буфферизация потока выполнения

[Пример](note_b2/ex_4/main.go)

```golang

func main() {
	ch := make(chan int, 5)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for {
			v, ok := <-ch
			if !ok {
				return
			}
			time.Sleep(time.Duration(v) * time.Second)
		}
	}()
}
```