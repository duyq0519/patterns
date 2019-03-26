package prototype

import "testing"

func TestPrototype(t *testing.T) {
	mc := NewMessageCache()

	note := mc.GetMessage(NoteBookType)
	note.Print()
	page := mc.GetMessage(PageInfoType)
	page.Print()
}
