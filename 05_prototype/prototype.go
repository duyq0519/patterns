package prototype

import "fmt"

const (
	NoteBookType = "notebook"
	PageInfoType = "pageinfo"
)

type Message interface {
	Print()
}

type NoteBook struct {
	Msg string
}

func (n *NoteBook) Print() {
	fmt.Printf("msg is %s\n", n.Msg)
}

type Pages struct {
	Info string
}

func (p *Pages) Print() {
	fmt.Printf("info is %s\n", p.Info)
}

type MessageCache struct {
	msgs map[string]Message
}

func NewMessageCache() *MessageCache {
	ms := make(map[string]Message)
	note := &NoteBook{Msg: "hello moto"}
	ms[NoteBookType] = note
	page := &Pages{Info: "hello apple"}
	ms[PageInfoType] = page
	return &MessageCache{
		msgs: ms,
	}
}

func (m *MessageCache) GetMessage(t string) Message {
	return m.msgs[t]
}
