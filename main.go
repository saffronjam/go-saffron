package main

import (
	"go-saffron/pkg/app"
	"go-saffron/pkg/sys"
	"runtime"
)

func init() { runtime.LockOSThread() }

func main() {
	saffronApp, err := app.NewApp(&app.Config{
		WindowProps: &sys.WindowProps{
			Width:  1600,
			Height: 900,
			Title:  "go-saffron Example",
		},
	})
	if err != nil {
		panic(err)
	}

	err = saffronApp.Run()
	if err != nil {
		panic(err)
	}
}
