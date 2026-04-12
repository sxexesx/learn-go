## Сделать из непредсказуемой функции предсказуемую // реализовать функцию обертку

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

Ответ: [тут](ex_1/main.go)
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

## Сделать из непредсказуемой функции предсказуемую через Context

Ответ:
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