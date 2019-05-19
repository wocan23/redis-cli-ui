package common

const (
	DEFAULT_WINDOW_WIDTH = 500
	DEFAULT_WINDOW_HEIGHT = 400


	DEFAULT_COM_WIDTH = 50
	DEFAULT_COM_HEIGHT = 30
)


var(
	ComponentPool = make(map[string]interface{})
	EventPool = make(map[string]func())
)
