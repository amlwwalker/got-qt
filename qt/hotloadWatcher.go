package main

import (
	"fmt"
	"github.com/radovskyb/watcher"
	"log"
	"path/filepath"
	"time"
)

type HotLoader struct {
	Loader    *func(p string)
	blackList map[string]bool
}

func (h *HotLoader) initBlacklist(extensions ...string) {
	h.blackList = make(map[string]bool)
	for _, v := range extensions {
		h.blackList[v] = true
	}
}
func (h *HotLoader) checkBlacklist(fileName string, dir bool) bool {
	if h.blackList[filepath.Ext(fileName)] == true {
		//ignore this file
		return true
	}
	if dir { //consider changes to a dir itself as ignore
		//ignore raw directory changes
		return true
	}
	return false
}
func (h *HotLoader) handleEvent(event watcher.Event) bool {
	if h.checkBlacklist(event.Name(), event.IsDir()) {
		//its blacklisted
		return true
	}
	fmt.Println("event occured " + event.Name() + ", type: " + event.Op.String() + ", path: " + event.Path)
	return false
}
func (h *HotLoader) startWatcher(loader func(string)) {
	h.initBlacklist(".qmlc", ".cpp", ".h", ".qrc", ".go", ".java")
	w := watcher.New()
	w.IgnoreHiddenFiles(true)
	// SetMaxEvents to 1 to allow at most 1 event's to be received
	// on the Event channel per watching cycle.
	//
	// If SetMaxEvents is not set, the default is to send all events.
	w.SetMaxEvents(1)

	// Only notify rename and move events.
	// w.FilterOps(watcher.Rename, watcher.Move)

	go func() {
		for {
			select {
			case event := <-w.Event:
				if !h.handleEvent(event) {
					fmt.Println("file type OK")
					fmt.Println("loader: ", loader)
					loader(event.Path)
				}
			case err := <-w.Error:
				fmt.Println("error: ", err.Error())
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch this folder for changes.
	if err := w.AddRecursive("."); err != nil {
		log.Fatalln(err)
	}

	// Print a list of all of the files and folders currently
	// being watched and their paths.
	// for path, f := range w.WatchedFiles() {
	// 	if !h.checkBlacklist(f.Name()) {
	// 		fmt.Printf("%s: %s\n", path, f.Name())
	// 	}
	// }

	fmt.Println()

	// Trigger 2 events after watcher started.
	// go func() {
	// 	w.Wait()
	// 	w.TriggerEvent(watcher.Create, nil)
	// 	w.TriggerEvent(watcher.Remove, nil)
	// }()

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
