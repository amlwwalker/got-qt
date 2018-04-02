package main

import (
	"os"
	"path/filepath"
	"strings"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
	"fmt"
)

type QmlBridge struct {
    core.QObject

    _ func(p string)        `signal:"updateLoader"`
    _ func(data string) string `slot:"sendToGo"`
}

func main() {

	//you could in essence, pass this to the hotloader,
	//and then actually have the live code anywhere you wanted
	//then have a flag that decides whether to use built on unbuilt qml files.
	var topLevel = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "amlwwalker", "qt-hotreloader", "qml")

	var qmlBridge = NewQmlBridge(nil)

	os.Setenv("QT_QUICK_CONTROLS_STYLE", "material")

	widgets.NewQApplication(len(os.Args), os.Args)
	// widgets.SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	var view = quick.NewQQuickView(nil)
	hotLoader := HotLoader{}

	loader := func(p string) {
		fmt.Println("changed:", p)
		view.SetSource(core.NewQUrl())
		view.Engine().ClearComponentCache()
		view.SetSource(core.NewQUrl3(topLevel + "/loader.qml", 0))
		//this is cool. If its not the loader page,
		//we can just push that file onto the stack
		//so the dev sees it immediately.
		if !strings.Contains(p, "/loader.qml") {
			relativePath := strings.Replace(p, topLevel + "/", "", -1)
			qmlBridge.UpdateLoader(relativePath)
		}
	}
	go hotLoader.startWatcher(loader)
	view.RootContext().SetContextProperty("QmlBridge", qmlBridge)
	view.SetSource(core.NewQUrl3(topLevel + "/loader.qml", 0))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.Show()

	widgets.QApplication_Exec()


}
