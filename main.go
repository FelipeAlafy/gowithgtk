package main

import (
	"log"

	"github.com/felipealafy/gowithgtk/urltools"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window, see: ", err)
	}
	win.SetTitle("Image downloader")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	//Container
	vbox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	if err != nil {
		log.Fatal("Unable to create vbox: ", err)
	}

	//label
	lbl, err := gtk.LabelNew("Put your url:")
	if err != nil {
		log.Fatal("Unable to create label", err)
	}
	vbox.Add(lbl)

	//Entry
	entry, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Unable to create entry: ", err)
	}
	vbox.Add(entry)

	//Save Folder Chooser
	chooser, err := gtk.FileChooserNativeDialogNew("Select your folder", win, gtk.FILE_CHOOSER_ACTION_SELECT_FOLDER, "Select", "Cancel")

	//button
	btn, _ := gtk.ButtonNewWithLabel("Download")
	btn.SetSensitive(false)
	if err != nil {
		log.Fatal("Unable to create label: ", err)
	}

	entry.Event(gdk.EVENT_KEY_PRESS)

	btn.Connect("clicked", func() {
		log.Println("Download Started")
		lbl.SetText("Download Started")
		url, _ := entry.GetText()
		re := chooser.Run()
		var path string
		log.Println("Chooser return: ", re)
		if re < 0 {
			path = chooser.GetURI()
		}
		_, filePath, err := urltools.Download(url, path)
		if err != nil {
			log.Panic("Download error: ", err)
		} else {
			lbl.SetText("Download Finished!")
			log.Println("Download saved in ", filePath)
		}
	})
	vbox.Add(btn)

	//Size
	win.Add(vbox)
	vbox.ShowAll()
	win.SetDefaultSize(300, 100)
	win.ShowAll()
	gtk.Main()
}
