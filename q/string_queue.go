package q

import (
	"container/list"
)

// NewStringQueue creates a non blocking string queue channel.
func NewStringQueue() (chan<- string, <-chan string) {
	send := make(chan string, 1)
	receive := make(chan string, 1)
	go manageStringQueue(send, receive)
	return send, receive
}

func manageStringQueue(send <-chan string, receive chan<- string) {
	queue := list.New()
	defer close(receive)
	for {
		if front := queue.Front(); front == nil {
			if value, ok := <-send; ok {
				queue.PushBack(value)
			} else {
				break
			}
		} else {
			select {
			case receive <- front.Value.(string):
				queue.Remove(front)
			case value, ok := <-send:
				if ok {
					queue.PushBack(value)
				}
			}
		}
	}
}
