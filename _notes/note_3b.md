## Graceful shutdown 

Пример [тут](note_3/ex_1/main.go)

Graceful shutdown - это способ остановить приложение или систему без потери данных, ошибок и обрыва процессов.

Проблемы:
1. Когда следует завершить работу? Все ли процессы завершены или какой-то застрял?
2. А как мы можем сообщить процессам о необходимости завершения работы?


Решение через каналы: 
```golang
func main() {
	// отслеживаем сигнал прерывания с клавиатуры os.Interrupt
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-c:
			fmt.Println("Break the loop")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("hello in a loop")
		}
	}
}
```

Решение через Context:
```golang
func interruptWithNotifyContext() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Break the loop")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("hello in a loop")
			}
		}
	}()

	wg.Wait()
	fmt.Println("Main done")
}
```