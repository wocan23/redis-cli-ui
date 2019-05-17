package main

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
	"os"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gdk"
	"./module"
	"./common"
)

func main(){
	mainFunc()
}

func mainFunc()  {
	// Create Gtk Application, change appID to your application domain name reversed.
	const appID = "org.gtk.es"
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	// Check to make sure no errors when creating Gtk Application
	if err != nil {
		log.Fatal("Could not create application.", err)
	}
	application.Connect("activate", func() { onActivate(application) })
	os.Exit(application.Run(os.Args))
}

func onActivate(application *gtk.Application) {
	// Create ApplicationWindow
	appWindow, err := gtk.ApplicationWindowNew(application)
	if err != nil {
		log.Fatal("Could not create application window.", err)
	}
	// Set ApplicationWindow Properties
	appWindow.SetTitle("Redis Cli")

	appWindow.SetKeepAbove(true)
	layout := module.Layout()
	appWindow.Add(layout)

	appWindow.SetPosition(gtk.WIN_POS_CENTER)
	appWindow.SetTypeHint(gdk.WINDOW_TYPE_HINT_DIALOG)
	appWindow.SetDefaultSize(common.DEFAULT_WINDOW_WIDTH,common.DEFAULT_WINDOW_HEIGHT)
	appWindow.ShowAll()
}





