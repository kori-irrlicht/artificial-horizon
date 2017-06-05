package main

import (
	"flag"
	"os"

	astilectron "github.com/asticode/go-astilectron"
	astilog "github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

func main() {

	flag.Parse()

	astilog.SetLogger(astilog.New(astilog.FlagConfig()))

	var err error
	var wd string
	if wd, err = os.Getwd(); err != nil {
		astilog.Fatal(errors.Wrap(err, "finding cwd"))
	}

	var a *astilectron.Astilectron
	if a, err = astilectron.New(astilectron.Options{
		AppName:           "Artificial Horizon",
		BaseDirectoryPath: wd,
	}); err != nil {
		astilog.Fatal(errors.Wrap(err, "creating new astilectron failed"))
	}

	defer a.Close()
	a.HandleSignals()

	if err = a.Start(); err != nil {
		astilog.Fatal(errors.Wrap(err, "starting failed"))
	}

	var w *astilectron.Window
	if w, err = a.NewWindowInDisplay(a.PrimaryDisplay(), "http://127.0.0.1:42424/static", &astilectron.WindowOptions{
		Center:     astilectron.PtrBool(true),
		Height:     astilectron.PtrInt(1920),
		Width:      astilectron.PtrInt(1080),
		Fullscreen: astilectron.PtrBool(true),
		Kiosk:      astilectron.PtrBool(true),
	}); err != nil {
		astilog.Fatal(errors.Wrap(err, "new window failed"))
	}
	if err = w.Create(); err != nil {
		astilog.Fatal(errors.Wrap(err, "creating window failed"))
	}
	//w.OpenDevTools()

	a.Wait()
}
