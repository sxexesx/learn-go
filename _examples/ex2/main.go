package main

import (
	"context"
	"errors"
	"time"
)

func main() {
}

type KVAgent struct {
	client Client
}

type Client interface {
	Get(ctx context.Context, key string) (result string, err error)
}

type ResultWithError struct {
	result string
	err    error
}

func (k *KVAgent) GetHedged(
	ctx context.Context,
	key string,
	hedgeDelay time.Duration,
	maxDelay time.Duration,
) (result string, err error) {

	ctx, cancel := context.WithTimeout(ctx, maxDelay)
	defer cancel()

	// master + 2 replicas
	resultCh := make(chan ResultWithError, 3)

	// отправляем запрос на мастер
	go func() {
		res, err := k.client.Get(ctx, key)
		resultCh <- ResultWithError{
			result: res,
			err:    err,
		}
	}()

	select {
	case v := <-resultCh:
		return v.result, v.err
	case <-time.After(hedgeDelay):
	}

	for i := 0; i < 2; i++ {
		go func() {
			res, err := k.client.Get(ctx, key)
			resultCh <- ResultWithError{result: res, err: err}
		}()
	}

	select {
	case v := <-resultCh:
		return v.result, v.err
	case <-ctx.Done():
	}
	return "", errors.New("timeout")
}
