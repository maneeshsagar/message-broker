package core

type MessageBroker struct {
	Msg      []string
	Offset   int
	Interval int
	MsgLimit int
}

func NewMessageBroker(interval, msgLimit int) *MessageBroker {
	return &MessageBroker{
		MsgLimit: msgLimit,
		Interval: interval,
	}
}

func (m *MessageBroker) AddMsg(msgs []string) {
	m.Msg = append(m.Msg, msgs...)
}

func (m *MessageBroker) Emit() []string {
	if len(m.Msg) < m.MsgLimit {
		messageToSend := m.Msg[:len(m.Msg)-1]
		m.Msg = make([]string, 0)
		return messageToSend
	}
	
	messageToSend := m.Msg[:m.MsgLimit]
	m.Msg = m.Msg[m.MsgLimit:]
	return messageToSend
}
