## Error группа

Полная версия кода [тут](note_b3/ex_1/main.go)

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