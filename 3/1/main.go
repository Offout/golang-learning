package main

import (
	"sync"
)

type job func(in, out chan interface{})

func Pipe(funcs ...job) {

	chans := make([]chan interface{}, len(funcs)+1)

	for idx := range chans {
		chans[idx] = make(chan interface{})
	}

	wg := new(sync.WaitGroup)

	for idx, f := range funcs {
		wg.Add(1)
		go func(f job, idx int) {
			f(chans[idx], chans[idx+1])

			wg.Done()
			close(chans[idx+1])
		}(f, idx)
	}
	wg.Wait()
	return
}