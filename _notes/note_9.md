# Патерны

## fanin

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

	return result
}
```