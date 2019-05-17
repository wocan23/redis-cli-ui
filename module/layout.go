package module

import "github.com/gotk3/gotk3/gtk"
import "../common"

func Layout() *gtk.Box{

	box,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)

	// ip
	ipbox,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)

	ipLabel,_ := gtk.LabelNew("ip")
	ipLabel.SetSizeRequest(common.DEFAULT_COM_WIDTH,common.DEFAULT_COM_HEIGHT)

	ipEntry,_ := gtk.EntryNew()
	ipEntry.SetWidthChars(20)
	ipEntry.SetPlaceholderText("127.0.0.1")

	ipbox.Add(ipLabel)
	ipbox.Add(ipEntry)
	// port
	portbox,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)

	portLabel,_ := gtk.LabelNew("port")
	portLabel.SetSizeRequest(common.DEFAULT_COM_WIDTH,common.DEFAULT_COM_HEIGHT)

	portEntry,_ := gtk.EntryNew()
	portEntry.SetWidthChars(20)
	portEntry.SetText("6379")

	portbox.Add(portLabel)
	portbox.Add(portEntry)

	// show
	showbox,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)
	showView,_ := gtk.TextViewNew()
	showView.SetSizeRequest(400,300)

	showbox.Add(showView)

	box.Add(ipbox)
	box.Add(portbox)
	box.Add(showbox)


	return box
}


