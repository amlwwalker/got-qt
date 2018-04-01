package main

import (
	"os"
	"path/filepath"

	"github.com/therecipe/qt/core"
	// "github.com/therecipe/qt/quick"
	// "github.com/therecipe/qt/widgets"

	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	// "github.com/therecipe/qt/qml"
	// "github.com/therecipe/qt/gui"
)

type QmlBridge struct {
    core.QObject

    _ func(p string)        `signal:"updateLoader"`
    _ func(data string) string `slot:"sendToGo"`
}

func (q *QmlBridge)initQQuickView(path string) *qml.QQmlApplicationEngine {

	var view = qml.NewQQmlApplicationEngine(nil)

	var watcher = core.NewQFileSystemWatcher2([]string{filepath.Dir(path)}, nil)

	var reload = func(p string) {
		println("changed:", p)
		// view.ClearComponentCache()
		q.UpdateLoader(p)
	}

	// watcher.ConnectFileChanged(reload)
	watcher.ConnectDirectoryChanged(reload)

	return view
}

func main() {

	var path = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "amlwwalker", "qt-recipe", "hotreload", "qml", "loader.qml")

	app := gui.NewQGuiApplication(len(os.Args), os.Args)
	app.SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	// widgets.NewQApplication(len(os.Args), os.Args)
	// gui.NewQGuiApplication(len(os.Args), os.Args)
	os.Setenv("QT_QUICK_CONTROLS_STYLE", "material")
	// gui.NewQGuiApplication(len(os.Args), os.Args)
	// widgets.NewQApplication(len(os.Args), os.Args)


	var qmlBridge = NewQmlBridge(nil)
	var view = qmlBridge.initQQuickView(path)
	// view.SetSource(core.NewQUrl3("qrc:/" + path, 0))
	// view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	// view.Show()

	view.RootContext().SetContextProperty("QmlBridge", qmlBridge)
	view.Load(core.NewQUrl3(path, 0))
	// widgets.QApplication_Exec()
	gui.QGuiApplication_Exec()



}

// func main() {
// 	var path = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "amlwwalker", "qt-recipe", "hotreload", "qml", "main.qml")

// 	gui.NewQGuiApplication(len(os.Args), os.Args)
// 	var view = qml.NewQQmlApplicationEngine(nil)
//     view.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

// 	gui.QGuiApplication_Exec()
// }
