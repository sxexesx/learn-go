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