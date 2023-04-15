package main

import (
	"context"

	"github.com/ServiceWeaver/weaver"
)

type Reverser interface {
	Reverse(context.Context, string) (string, error)
}

type reverser struct {
	weaver.Implements[Reverser]
	cache Cache
}

func (r *reverser) Init(context.Context) error {
	var err error
	r.cache, err = weaver.Get[Cache](r)
	return err
}

func (r *reverser) Reverse(ctx context.Context, s string) (string, error) {
	if reversed, err := r.cache.Get(ctx, s); err == nil {
		return reversed, nil
	}
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	reversed := string(runes)
	r.cache.Set(ctx, s, reversed)
	return reversed, nil
}
