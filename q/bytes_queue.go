package q

import "container/list"

// NewBytesQueue creates a non blocking bytearray queue channel.
func NewBytesQueue() (chan<- []byte, <-chan []byte) {
	send := make(chan []byte, 1)
	receive := make(chan []byte, 1)
	go manageBytesQueue(send, receive)
	return send, receive
}

func manageBytesQueue(send <-chan []byte, receive chan<- []byte) {
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
			case receive <- front.Value.([]byte):
				queue.Remove(front)
			case value, ok := <-send:
				if ok {
					queue.PushBack(value)
				}
			}
		}
	}
}
