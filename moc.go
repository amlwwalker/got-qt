package main

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type ApplicationUI_ITF interface {
	std_core.QObject_ITF
	ApplicationUI_PTR() *ApplicationUI
}

func (ptr *ApplicationUI) ApplicationUI_PTR() *ApplicationUI {
	return ptr
}

func (ptr *ApplicationUI) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *ApplicationUI) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromApplicationUI(ptr ApplicationUI_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ApplicationUI_PTR().Pointer()
	}
	return nil
}

func NewApplicationUIFromPointer(ptr unsafe.Pointer) *ApplicationUI {
	var n *ApplicationUI
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(ApplicationUI)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *ApplicationUI:
			n = deduced

		case *std_core.QObject:
			n = &ApplicationUI{QObject: *deduced}

		default:
			n = new(ApplicationUI)
			n.SetPointer(ptr)
		}
	}
	return n
}

//export callbackApplicationUI_Constructor
func callbackApplicationUI_Constructor(ptr unsafe.Pointer) {
	gPtr := NewApplicationUIFromPointer(ptr)
	qt.Register(ptr, gPtr)
}

//export callbackApplicationUI_SwapThemePalette
func callbackApplicationUI_SwapThemePalette(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "swapThemePalette"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := make([]string, 0)
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *ApplicationUI) ConnectSwapThemePalette(f func() []string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "swapThemePalette"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "swapThemePalette", func() []string {
				signal.(func() []string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "swapThemePalette", f)
		}
	}
}

func (ptr *ApplicationUI) DisconnectSwapThemePalette() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "swapThemePalette")
	}
}

func (ptr *ApplicationUI) SwapThemePalette() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.ApplicationUI_SwapThemePalette(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackApplicationUI_DefaultThemePalette
func callbackApplicationUI_DefaultThemePalette(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "defaultThemePalette"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := make([]string, 0)
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *ApplicationUI) ConnectDefaultThemePalette(f func() []string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "defaultThemePalette"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "defaultThemePalette", func() []string {
				signal.(func() []string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "defaultThemePalette", f)
		}
	}
}

func (ptr *ApplicationUI) DisconnectDefaultThemePalette() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "defaultThemePalette")
	}
}

func (ptr *ApplicationUI) DefaultThemePalette() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.ApplicationUI_DefaultThemePalette(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackApplicationUI_PrimaryPalette
func callbackApplicationUI_PrimaryPalette(ptr unsafe.Pointer, paletteIndex C.int) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "primaryPalette"); signal != nil {
		tempVal := signal.(func(int) []string)(int(int32(paletteIndex)))
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := make([]string, 0)
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *ApplicationUI) ConnectPrimaryPalette(f func(paletteIndex int) []string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "primaryPalette"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "primaryPalette", func(paletteIndex int) []string {
				signal.(func(int) []string)(paletteIndex)
				return f(paletteIndex)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "primaryPalette", f)
		}
	}
}

func (ptr *ApplicationUI) DisconnectPrimaryPalette() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "primaryPalette")
	}
}

func (ptr *ApplicationUI) PrimaryPalette(paletteIndex int) []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.ApplicationUI_PrimaryPalette(ptr.Pointer(), C.int(int32(paletteIndex)))), "|")
	}
	return make([]string, 0)
}

//export callbackApplicationUI_AccentPalette
func callbackApplicationUI_AccentPalette(ptr unsafe.Pointer, paletteIndex C.int) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "accentPalette"); signal != nil {
		tempVal := signal.(func(int) []string)(int(int32(paletteIndex)))
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := make([]string, 0)
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *ApplicationUI) ConnectAccentPalette(f func(paletteIndex int) []string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "accentPalette"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "accentPalette", func(paletteIndex int) []string {
				signal.(func(int) []string)(paletteIndex)
				return f(paletteIndex)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "accentPalette", f)
		}
	}
}

func (ptr *ApplicationUI) DisconnectAccentPalette() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "accentPalette")
	}
}

func (ptr *ApplicationUI) AccentPalette(paletteIndex int) []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.ApplicationUI_AccentPalette(ptr.Pointer(), C.int(int32(paletteIndex)))), "|")
	}
	return make([]string, 0)
}

//export callbackApplicationUI_DefaultPrimaryPalette
func callbackApplicationUI_DefaultPrimaryPalette(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "defaultPrimaryPalette"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := make([]string, 0)
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *ApplicationUI) ConnectDefaultPrimaryPalette(f func() []string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "defaultPrimaryPalette"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "defaultPrimaryPalette", func() []string {
				signal.(func() []string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "defaultPrimaryPalette", f)
		}
	}
}

func (ptr *ApplicationUI) DisconnectDefaultPrimaryPalette() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "defaultPrimaryPalette")
	}
}

func (ptr *ApplicationUI) DefaultPrimaryPalette() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.ApplicationUI_DefaultPrimaryPalette(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackApplicationUI_DefaultAccentPalette
func callbackApplicationUI_DefaultAccentPalette(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "defaultAccentPalette"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := make([]string, 0)
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *ApplicationUI) ConnectDefaultAccentPalette(f func() []string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "defaultAccentPalette"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "defaultAccentPalette", func() []string {
				signal.(func() []string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "defaultAccentPalette", f)
		}
	}
}

func (ptr *ApplicationUI) DisconnectDefaultAccentPalette() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "defaultAccentPalette")
	}
}

func (ptr *ApplicationUI) DefaultAccentPalette() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.ApplicationUI_DefaultAccentPalette(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func ApplicationUI_QRegisterMetaType() int {
	return int(int32(C.ApplicationUI_ApplicationUI_QRegisterMetaType()))
}

func (ptr *ApplicationUI) QRegisterMetaType() int {
	return int(int32(C.ApplicationUI_ApplicationUI_QRegisterMetaType()))
}

func ApplicationUI_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ApplicationUI_ApplicationUI_QRegisterMetaType2(typeNameC)))
}

func (ptr *ApplicationUI) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ApplicationUI_ApplicationUI_QRegisterMetaType2(typeNameC)))
}

func ApplicationUI_QmlRegisterType() int {
	return int(int32(C.ApplicationUI_ApplicationUI_QmlRegisterType()))
}

func (ptr *ApplicationUI) QmlRegisterType() int {
	return int(int32(C.ApplicationUI_ApplicationUI_QmlRegisterType()))
}

func ApplicationUI_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.ApplicationUI_ApplicationUI_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ApplicationUI) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.ApplicationUI_ApplicationUI_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ApplicationUI) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.ApplicationUI___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *ApplicationUI) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.ApplicationUI___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *ApplicationUI) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.ApplicationUI___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *ApplicationUI) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.ApplicationUI___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApplicationUI) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApplicationUI___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApplicationUI) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.ApplicationUI___findChildren_newList2(ptr.Pointer()))
}

func (ptr *ApplicationUI) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.ApplicationUI___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApplicationUI) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApplicationUI___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApplicationUI) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.ApplicationUI___findChildren_newList3(ptr.Pointer()))
}

func (ptr *ApplicationUI) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.ApplicationUI___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApplicationUI) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApplicationUI___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApplicationUI) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.ApplicationUI___findChildren_newList(ptr.Pointer()))
}

func (ptr *ApplicationUI) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.ApplicationUI___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ApplicationUI) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ApplicationUI___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ApplicationUI) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.ApplicationUI___children_newList(ptr.Pointer()))
}

func NewApplicationUI(parent std_core.QObject_ITF) *ApplicationUI {
	var tmpValue = NewApplicationUIFromPointer(C.ApplicationUI_NewApplicationUI(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackApplicationUI_DestroyApplicationUI
func callbackApplicationUI_DestroyApplicationUI(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~ApplicationUI"); signal != nil {
		signal.(func())()
	} else {
		NewApplicationUIFromPointer(ptr).DestroyApplicationUIDefault()
	}
}

func (ptr *ApplicationUI) ConnectDestroyApplicationUI(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~ApplicationUI"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~ApplicationUI", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~ApplicationUI", f)
		}
	}
}

func (ptr *ApplicationUI) DisconnectDestroyApplicationUI() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~ApplicationUI")
	}
}

func (ptr *ApplicationUI) DestroyApplicationUI() {
	if ptr.Pointer() != nil {
		C.ApplicationUI_DestroyApplicationUI(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *ApplicationUI) DestroyApplicationUIDefault() {
	if ptr.Pointer() != nil {
		C.ApplicationUI_DestroyApplicationUIDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackApplicationUI_Event
func callbackApplicationUI_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewApplicationUIFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *ApplicationUI) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.ApplicationUI_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackApplicationUI_EventFilter
func callbackApplicationUI_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewApplicationUIFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *ApplicationUI) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.ApplicationUI_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackApplicationUI_ChildEvent
func callbackApplicationUI_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewApplicationUIFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *ApplicationUI) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApplicationUI_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackApplicationUI_ConnectNotify
func callbackApplicationUI_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewApplicationUIFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ApplicationUI) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ApplicationUI_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackApplicationUI_CustomEvent
func callbackApplicationUI_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewApplicationUIFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *ApplicationUI) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApplicationUI_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackApplicationUI_DeleteLater
func callbackApplicationUI_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewApplicationUIFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *ApplicationUI) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.ApplicationUI_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackApplicationUI_Destroyed
func callbackApplicationUI_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackApplicationUI_DisconnectNotify
func callbackApplicationUI_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewApplicationUIFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ApplicationUI) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ApplicationUI_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackApplicationUI_ObjectNameChanged
func callbackApplicationUI_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackApplicationUI_TimerEvent
func callbackApplicationUI_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewApplicationUIFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *ApplicationUI) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ApplicationUI_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
