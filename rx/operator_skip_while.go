package rx

import (
	"context"
)

type operatorSkipWhile struct {
	ctx       context.Context
	predicate Predicate
	skip      bool
}

func (o *operatorSkipWhile) next(item interface{}, dst chan<- interface{}) bool {
	if !o.skip {
		send(dst, item)

		return true
	}

	if !o.predicate(o.ctx, item) {
		o.skip = false
		send(dst, item)

		return true
	}

	return false
}

func (o *operatorSkipWhile) end(dst chan<- interface{}) {}

// SkipWhile discard items until a specified condition becomes false.
func (o *Operable) SkipWhile(predicate Predicate) *Operable {
	o.mu.Lock()
	defer o.mu.Unlock()

	o.operators = append(o.operators, &operatorSkipWhile{
		ctx:       o.ctx,
		predicate: predicate,
		skip:      true,
	})

	return o
}
