package config

var Message message

type message string

func (m *message) Set(msg string) {
	*m = message(msg)
	// go func() {
	// 	time.Sleep(time.Second * 1)
	// 	*m = message("")
	// }()
}
