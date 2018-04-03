package main

import (
	"os"
	"path/filepath"
	"strings"
	"fmt"
	"time"
	"encoding/json"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)
type Config struct {
    Author string `"json":"author"`
    Date string `"json":"date"`
    Mode string `"json":"mode"`
    Host string `"json":"host"`
    Version string `"json":"version"`
    Port string `"json":"port"`
    Hotload bool `"json":"hotload"`
}

type QmlBridge struct {
    core.QObject
    Config Config
    //messages to qml
    _ func(p string)        `signal:"updateLoader"`
    _ func(author, mode, date, host, version, port string, hotload bool)        `signal:"updateSettings"`
    _ func(data string) 	`signal:"sendTime"`

    //requests from qml
    _ func(number1, number2 string) string `slot:"calculator"`
}

func LoadConfiguration(file string) (error, Config) {
    var config Config
    configFile, err := os.Open(file)
    defer configFile.Close()
    if err != nil {
        fmt.Println(err.Error())
        return err, config
    }
    jsonParser := json.NewDecoder(configFile)
    jsonParser.Decode(&config)
    return nil, config
}
func addingNumbers(number1, number2 string) string {
	fmt.Println("addingNumbers")
	return number1 + number1 + number2 + number2
}
func main() {
	_, config := LoadConfiguration("config.json")
	fmt.Printf("config: %#v\r\n\r\n", config)
	var qmlBridge = NewQmlBridge(nil)
    qmlBridge.ConnectCalculator(func(number1, number2 string) string {
        return addingNumbers(number1, number2)
    })
	go func() {
		for t := range time.NewTicker(time.Second * 1).C {
			qmlBridge.SendTime(t.Format(time.ANSIC))
		}
	}()
	//configure whether to be in hotloading mode for qml:
	//setting the env var for the compiled qt?
	os.Setenv("QT_QUICK_CONTROLS_STYLE", "material")

	//you could in essence, pass this to the hotloader,
	//and then actually have the live code anywhere you wanted
	//then have a flag that decides whether to use built on unbuilt qml files.
	var topLevel = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "amlwwalker", "qt-hotreloader", "qml")

	widgets.NewQApplication(len(os.Args), os.Args)
	// widgets.SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	var view = quick.NewQQuickView(nil)
	//if we are hotloading, gonna need to inform the front end
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
	view.RootContext().SetContextProperty("QmlBridge", qmlBridge)
	if !config.Hotload {
		fmt.Println("compiling qml into binary...")
		view.SetSource(core.NewQUrl3("qrc:/qml/loader-production.qml", 0))
	} else {
		view.SetSource(core.NewQUrl3(topLevel + "/loader.qml", 0))
		go hotLoader.startWatcher(loader)
	}


	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.Show()
	//this is messy, should be able to pass to frontend immediately
	go func() {
        time.Sleep(5 * time.Second)
        fmt.Println("updating settings with ", config)
        qmlBridge.UpdateSettings(config.Author, config.Mode, config.Date, config.Host, config.Version, config.Port, config.Hotload)
    }()
	widgets.QApplication_Exec()


}
