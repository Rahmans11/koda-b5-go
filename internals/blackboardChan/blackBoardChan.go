package blackboardchan

import (
	"fmt"
	"sync"
)

type Message struct {
	Sender  string
	Content string
}

func NewMessage(Sender string, Content string) Message {
	return Message{
		Sender:  Sender,
		Content: Content,
	}
}

func Blackboard(wg *sync.WaitGroup, mc chan Message) {
	defer wg.Done()

	for msg := range mc {
		fmt.Printf("\npesan dari %s: %s\n", msg.Sender, msg.Content)
	}

}

func MsgSender(mc chan Message, msg Message) {
	var wg sync.WaitGroup
	wg.Go(func() {
		func() {
			mc <- msg
		}()
	})
	wg.Wait()
}
