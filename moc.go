package main

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
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

type QmlBridge_ITF interface {
	std_core.QObject_ITF
	QmlBridge_PTR() *QmlBridge
}

func (ptr *QmlBridge) QmlBridge_PTR() *QmlBridge {
	return ptr
}

func (ptr *QmlBridge) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlBridge) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQmlBridge(ptr QmlBridge_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlBridge_PTR().Pointer()
	}
	return nil
}

func NewQmlBridgeFromPointer(ptr unsafe.Pointer) *QmlBridge {
	var n *QmlBridge
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(QmlBridge)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *QmlBridge:
			n = deduced

		case *std_core.QObject:
			n = &QmlBridge{QObject: *deduced}

		default:
			n = new(QmlBridge)
			n.SetPointer(ptr)
		}
	}
	return n
}

//export callbackQmlBridge_Constructor
func callbackQmlBridge_Constructor(ptr unsafe.Pointer) {
	gPtr := NewQmlBridgeFromPointer(ptr)
	qt.Register(ptr, gPtr)
}

//export callbackQmlBridge_UpdateLoader
func callbackQmlBridge_UpdateLoader(ptr unsafe.Pointer, p C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "updateLoader"); signal != nil {
		signal.(func(string))(cGoUnpackString(p))
	}

}

func (ptr *QmlBridge) ConnectUpdateLoader(f func(p string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "updateLoader") {
			C.QmlBridge_ConnectUpdateLoader(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "updateLoader"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "updateLoader", func(p string) {
				signal.(func(string))(p)
				f(p)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateLoader", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectUpdateLoader() {
	if ptr.Pointer() != nil {
		C.QmlBridge_DisconnectUpdateLoader(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "updateLoader")
	}
}

func (ptr *QmlBridge) UpdateLoader(p string) {
	if ptr.Pointer() != nil {
		var pC *C.char
		if p != "" {
			pC = C.CString(p)
			defer C.free(unsafe.Pointer(pC))
		}
		C.QmlBridge_UpdateLoader(ptr.Pointer(), C.struct_Moc_PackedString{data: pC, len: C.longlong(len(p))})
	}
}

//export callbackQmlBridge_UpdateSettings
func callbackQmlBridge_UpdateSettings(ptr unsafe.Pointer, author C.struct_Moc_PackedString, mode C.struct_Moc_PackedString, date C.struct_Moc_PackedString, host C.struct_Moc_PackedString, version C.struct_Moc_PackedString, port C.struct_Moc_PackedString, hotload C.char) {
	if signal := qt.GetSignal(ptr, "updateSettings"); signal != nil {
		signal.(func(string, string, string, string, string, string, bool))(cGoUnpackString(author), cGoUnpackString(mode), cGoUnpackString(date), cGoUnpackString(host), cGoUnpackString(version), cGoUnpackString(port), int8(hotload) != 0)
	}

}

func (ptr *QmlBridge) ConnectUpdateSettings(f func(author string, mode string, date string, host string, version string, port string, hotload bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "updateSettings") {
			C.QmlBridge_ConnectUpdateSettings(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "updateSettings"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "updateSettings", func(author string, mode string, date string, host string, version string, port string, hotload bool) {
				signal.(func(string, string, string, string, string, string, bool))(author, mode, date, host, version, port, hotload)
				f(author, mode, date, host, version, port, hotload)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateSettings", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectUpdateSettings() {
	if ptr.Pointer() != nil {
		C.QmlBridge_DisconnectUpdateSettings(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "updateSettings")
	}
}

func (ptr *QmlBridge) UpdateSettings(author string, mode string, date string, host string, version string, port string, hotload bool) {
	if ptr.Pointer() != nil {
		var authorC *C.char
		if author != "" {
			authorC = C.CString(author)
			defer C.free(unsafe.Pointer(authorC))
		}
		var modeC *C.char
		if mode != "" {
			modeC = C.CString(mode)
			defer C.free(unsafe.Pointer(modeC))
		}
		var dateC *C.char
		if date != "" {
			dateC = C.CString(date)
			defer C.free(unsafe.Pointer(dateC))
		}
		var hostC *C.char
		if host != "" {
			hostC = C.CString(host)
			defer C.free(unsafe.Pointer(hostC))
		}
		var versionC *C.char
		if version != "" {
			versionC = C.CString(version)
			defer C.free(unsafe.Pointer(versionC))
		}
		var portC *C.char
		if port != "" {
			portC = C.CString(port)
			defer C.free(unsafe.Pointer(portC))
		}
		C.QmlBridge_UpdateSettings(ptr.Pointer(), C.struct_Moc_PackedString{data: authorC, len: C.longlong(len(author))}, C.struct_Moc_PackedString{data: modeC, len: C.longlong(len(mode))}, C.struct_Moc_PackedString{data: dateC, len: C.longlong(len(date))}, C.struct_Moc_PackedString{data: hostC, len: C.longlong(len(host))}, C.struct_Moc_PackedString{data: versionC, len: C.longlong(len(version))}, C.struct_Moc_PackedString{data: portC, len: C.longlong(len(port))}, C.char(int8(qt.GoBoolToInt(hotload))))
	}
}

//export callbackQmlBridge_SendTime
func callbackQmlBridge_SendTime(ptr unsafe.Pointer, data C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "sendTime"); signal != nil {
		signal.(func(string))(cGoUnpackString(data))
	}

}

func (ptr *QmlBridge) ConnectSendTime(f func(data string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sendTime") {
			C.QmlBridge_ConnectSendTime(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sendTime"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "sendTime", func(data string) {
				signal.(func(string))(data)
				f(data)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sendTime", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectSendTime() {
	if ptr.Pointer() != nil {
		C.QmlBridge_DisconnectSendTime(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sendTime")
	}
}

func (ptr *QmlBridge) SendTime(data string) {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		C.QmlBridge_SendTime(ptr.Pointer(), C.struct_Moc_PackedString{data: dataC, len: C.longlong(len(data))})
	}
}

//export callbackQmlBridge_Calculator
func callbackQmlBridge_Calculator(ptr unsafe.Pointer, number1 C.struct_Moc_PackedString, number2 C.struct_Moc_PackedString) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "calculator"); signal != nil {
		tempVal := signal.(func(string, string) string)(cGoUnpackString(number1), cGoUnpackString(number2))
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QmlBridge) ConnectCalculator(f func(number1 string, number2 string) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "calculator"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "calculator", func(number1 string, number2 string) string {
				signal.(func(string, string) string)(number1, number2)
				return f(number1, number2)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "calculator", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectCalculator() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "calculator")
	}
}

func (ptr *QmlBridge) Calculator(number1 string, number2 string) string {
	if ptr.Pointer() != nil {
		var number1C *C.char
		if number1 != "" {
			number1C = C.CString(number1)
			defer C.free(unsafe.Pointer(number1C))
		}
		var number2C *C.char
		if number2 != "" {
			number2C = C.CString(number2)
			defer C.free(unsafe.Pointer(number2C))
		}
		return cGoUnpackString(C.QmlBridge_Calculator(ptr.Pointer(), C.struct_Moc_PackedString{data: number1C, len: C.longlong(len(number1))}, C.struct_Moc_PackedString{data: number2C, len: C.longlong(len(number2))}))
	}
	return ""
}

func QmlBridge_QRegisterMetaType() int {
	return int(int32(C.QmlBridge_QmlBridge_QRegisterMetaType()))
}

func (ptr *QmlBridge) QRegisterMetaType() int {
	return int(int32(C.QmlBridge_QmlBridge_QRegisterMetaType()))
}

func QmlBridge_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QmlBridge_QmlBridge_QRegisterMetaType2(typeNameC)))
}

func (ptr *QmlBridge) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QmlBridge_QmlBridge_QRegisterMetaType2(typeNameC)))
}

func QmlBridge_QmlRegisterType() int {
	return int(int32(C.QmlBridge_QmlBridge_QmlRegisterType()))
}

func (ptr *QmlBridge) QmlRegisterType() int {
	return int(int32(C.QmlBridge_QmlBridge_QmlRegisterType()))
}

func QmlBridge_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.QmlBridge_QmlBridge_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QmlBridge) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.QmlBridge_QmlBridge_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QmlBridge) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.QmlBridge___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *QmlBridge) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QmlBridge___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QmlBridge) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QmlBridge___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QmlBridge___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QmlBridge) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QmlBridge___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QmlBridge___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QmlBridge) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QmlBridge___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QmlBridge___findChildren_newList(ptr.Pointer()))
}

func (ptr *QmlBridge) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QmlBridge___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QmlBridge___children_newList(ptr.Pointer()))
}

func NewQmlBridge(parent std_core.QObject_ITF) *QmlBridge {
	var tmpValue = NewQmlBridgeFromPointer(C.QmlBridge_NewQmlBridge(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQmlBridge_DestroyQmlBridge
func callbackQmlBridge_DestroyQmlBridge(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QmlBridge"); signal != nil {
		signal.(func())()
	} else {
		NewQmlBridgeFromPointer(ptr).DestroyQmlBridgeDefault()
	}
}

func (ptr *QmlBridge) ConnectDestroyQmlBridge(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QmlBridge"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QmlBridge", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QmlBridge", f)
		}
	}
}

func (ptr *QmlBridge) DisconnectDestroyQmlBridge() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QmlBridge")
	}
}

func (ptr *QmlBridge) DestroyQmlBridge() {
	if ptr.Pointer() != nil {
		C.QmlBridge_DestroyQmlBridge(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QmlBridge) DestroyQmlBridgeDefault() {
	if ptr.Pointer() != nil {
		C.QmlBridge_DestroyQmlBridgeDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQmlBridge_Event
func callbackQmlBridge_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQmlBridgeFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *QmlBridge) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QmlBridge_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQmlBridge_EventFilter
func callbackQmlBridge_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQmlBridgeFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *QmlBridge) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QmlBridge_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQmlBridge_ChildEvent
func callbackQmlBridge_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QmlBridge) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackQmlBridge_ConnectNotify
func callbackQmlBridge_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlBridgeFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlBridge) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQmlBridge_CustomEvent
func callbackQmlBridge_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *QmlBridge) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackQmlBridge_DeleteLater
func callbackQmlBridge_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQmlBridgeFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QmlBridge) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QmlBridge_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQmlBridge_Destroyed
func callbackQmlBridge_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackQmlBridge_DisconnectNotify
func callbackQmlBridge_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlBridgeFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlBridge) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQmlBridge_ObjectNameChanged
func callbackQmlBridge_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQmlBridge_TimerEvent
func callbackQmlBridge_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QmlBridge) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridge_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
