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


Решение [тут](ex_1/main.go)
