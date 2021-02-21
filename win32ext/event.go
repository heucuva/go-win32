// +build windows

package win32ext

import (
	"github.com/heucuva/go-win32"
	"golang.org/x/sys/windows"
)

// EventToChannel turns an event handle into a channel
//  returns: channel, cancelFunc
func EventToChannel(event windows.Handle) (<-chan struct{}, func()) {
	ch := make(chan struct{}, 1)
	done := make(chan struct{}, 1)
	go func() {
		defer close(ch)
		for {
			select {
			case <-done:
				return
			default:
			}
			if err := win32.WaitForSingleObjectInfinite(event); err != nil {
				panic(err)
			}
			ch <- struct{}{}
		}
	}()
	return ch, func() {
		done <- struct{}{}
		close(done)
	}
}
