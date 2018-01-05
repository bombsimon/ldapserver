package ldapserver

import (
	"fmt"

	"github.com/bombsimon/ldapserver/message"
)

type Message struct {
	*message.LDAPMessage
	Client *client
	Done   chan bool
}

func (m *Message) String() string {
	return fmt.Sprintf("MessageId=%d, %s", m.MessageID(), m.ProtocolOpName())
}

// Abandon close the Done channel, to notify handler's user function to stop any
// running process
func (m *Message) Abandon() {
	m.Done <- true
}

func (m *Message) GetAbandonRequest() message.AbandonRequest {
	return m.ProtocolOp().(message.AbandonRequest)
}

func (m *Message) GetSearchRequest() message.SearchRequest {
	return m.ProtocolOp().(message.SearchRequest)
}

func (m *Message) GetBindRequest() message.BindRequest {
	return m.ProtocolOp().(message.BindRequest)
}

func (m *Message) GetAddRequest() message.AddRequest {
	return m.ProtocolOp().(message.AddRequest)
}

func (m *Message) GetDeleteRequest() message.DelRequest {
	return m.ProtocolOp().(message.DelRequest)
}

func (m *Message) GetModifyRequest() message.ModifyRequest {
	return m.ProtocolOp().(message.ModifyRequest)
}

func (m *Message) GetCompareRequest() message.CompareRequest {
	return m.ProtocolOp().(message.CompareRequest)
}

func (m *Message) GetExtendedRequest() message.ExtendedRequest {
	return m.ProtocolOp().(message.ExtendedRequest)
}
