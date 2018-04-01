package main

import (
	"os"
	"path/filepath"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"

	// "github.com/therecipe/qt/gui"
	// "github.com/therecipe/qt/qml"
	// "github.com/therecipe/qt/qml"
	// "github.com/therecipe/qt/gui"
)


func initQQuickView(path string) *quick.QQuickView {

	var view = quick.NewQQuickView(nil)

	var watcher = core.NewQFileSystemWatcher2([]string{filepath.Dir(path)}, nil)

	var reload = func(p string) {
		println("changed:", p)
		view.SetSource(core.NewQUrl())
		view.Engine().ClearComponentCache()
		view.SetSource(core.NewQUrl3(path, 0))
	}

	//watcher.ConnectFileChanged(reload)
	watcher.ConnectDirectoryChanged(reload)

	return view
}

func main() {

	var path = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "amlwwalker", "qt-recipe", "hotreload", "qml", "main.qml")

	// widgets.NewQApplication(len(os.Args), os.Args)
	// gui.NewQGuiApplication(len(os.Args), os.Args)
	os.Setenv("QT_QUICK_CONTROLS_STYLE", "material")

	widgets.NewQApplication(len(os.Args), os.Args)

	var view = initQQuickView(path)
	// view.Load(core.NewQUrl3("qrc:/" + path, 0))
	view.SetSource(core.NewQUrl3(path, 0))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.Show()

	widgets.QApplication_Exec()
}



// func main() {
// 	var path = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "amlwwalker", "qt-recipe", "hotreload", "qml", "main.qml")

// 	gui.NewQGuiApplication(len(os.Args), os.Args)
// 	var view = qml.NewQQmlApplicationEngine(nil)
//     view.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

// 	gui.QGuiApplication_Exec()
// }
