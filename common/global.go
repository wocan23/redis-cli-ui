package common

import (
	"github.com/gotk3/gotk3/gtk"
	"sync"
)

const (
	DEFAULT_WINDOW_WIDTH = 500
	DEFAULT_WINDOW_HEIGHT = 400


	DEFAULT_COM_WIDTH = 50
	DEFAULT_COM_HEIGHT = 30
)


var(
	AppWindow *gtk.ApplicationWindow
	ComponentPool = make(map[string]interface{})
	EventPool = make(map[string]func())

	LastCmdStrs = make([]string,0)
	LastCmdIndex = 0

	LastCmdStrLock = new(sync.RWMutex)
)
