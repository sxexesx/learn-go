## 1. Написать свой краулер и ограничить его таким образом, чтобы сайт парсился 1 раз в секунду.

По сути задача сводится к rate limiter-у

Ответ [тут](/_tasks/ex_301/main.go)


## 2. Используя error group конкурентно выполнить N-запросов  

Ответ [тут](/_tasks/ex_302/main.go)

<details>
    <summary>Ответ</summary>

```golang
func process(ctx context.Context, users []User) (map[string]int64, error) {
	names := make(map[string]int64, 0)
	mu := sync.Mutex{}

	egroup, ectx := errgroup.WithContext(ctx)
    // количество горутин
    egroup.SetLimit(100)

	for _, u := range users {
		egroup.Go(
			func() error {
				name, err := fetch(ctx, u)
				if err != nil {
					return err
				}

				mu.Lock()
				defer mu.Unlock()
				names[name] = names[name] + 1

				return nil
			})
	}

	if err := egroup.Wait(); err != nil {
		return nil, err
	}

	return names, nil
}
```
</details>


## 3. Реализовать WorkerPool с дополнительными условиями

```golang
type WorkerPool struct {
}

func NewWorkerPool(numberOfWorkers int) *WorkerPool {
}

// Submit - добавить таску в воркер пул и возвращает управление (неблокирующая операция).
// Обеспечить соблюдение очереди при запуске задач.
func (wp *WorkerPool) Submit(task func()) error {
}

// SubmitWait - добавить таску в воркер пул и дождаться окончания ее выполнения.
// Если был вызван метод Stop, SubmitWait выходит с ошибкой ErrStopped для задач
// которые ждут выполнения. Задачи которые выполняются должны доработать.
func (wp *WorkerPool) SubmitWait(task func()) error {
}

// Stop - остановить воркер пул, дождаться выполнения только тех тасок, которые выполняются сейчас
func (wp *WorkerPool) Stop() error {
}

// StopWait - остановить воркер пул, дождаться выполнения всех тасок, даже тех, что не начали выполняться, но лежат в очереди
func (wp *WorkerPool) StopWait() error {
}
```