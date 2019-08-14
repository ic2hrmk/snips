package main

import (
	"flag"
	"log"

	"github.com/mattn/go-gtk/gtk"
	"github.com/mattn/go-webkit/webkit"
)

const (
	defaultURL    = "http://google.com"
	defaultTitle  = "Web Page"
	defaultWidth  = 1024
	defaultHeight = 768
)

var (
	title  string
	url    string
	width  int
	height int
)

func init() {
	log.Println("GTK3+ Based Web Browser")

	flag.StringVar(&url, "url", defaultURL, "URL to open at startup")
	flag.StringVar(&title, "title", defaultTitle, "Window title")
	flag.IntVar(&width, "width", defaultWidth, "Window width")
	flag.IntVar(&height, "height", defaultHeight, "Window height")
	flag.Parse()

	log.Printf("attempting to open: %s", url)
}

func main() {
	gtk.Init(nil)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle(title)
	window.Connect("destroy", gtk.MainQuit)

	vbox := gtk.NewVBox(false, 1)

	swin := gtk.NewScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	swin.SetShadowType(gtk.SHADOW_IN)

	webview := webkit.NewWebView()
	webview.Open(url)
	swin.Add(webview)
	vbox.Add(swin)
	window.Add(vbox)
	window.SetSizeRequest(width, height)
	window.ShowAll()

	gtk.Main()
}
