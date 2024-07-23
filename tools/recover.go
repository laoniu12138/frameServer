package tools

import "context"

func GoWithRecover(ctx context.Context, f func(context.Context)) {
	defer func() {
		if e := recover(); e != nil {

		}
	}()
	go f(ctx)
}
