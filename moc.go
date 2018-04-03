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

type PersonModel_ITF interface {
	std_core.QAbstractListModel_ITF
	PersonModel_PTR() *PersonModel
}

func (ptr *PersonModel) PersonModel_PTR() *PersonModel {
	return ptr
}

func (ptr *PersonModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModel_PTR().Pointer()
	}
	return nil
}

func (ptr *PersonModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractListModel_PTR().SetPointer(p)
	}
}

func PointerFromPersonModel(ptr PersonModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.PersonModel_PTR().Pointer()
	}
	return nil
}

func NewPersonModelFromPointer(ptr unsafe.Pointer) *PersonModel {
	var n *PersonModel
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(PersonModel)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *PersonModel:
			n = deduced

		case *std_core.QAbstractListModel:
			n = &PersonModel{QAbstractListModel: *deduced}

		default:
			n = new(PersonModel)
			n.SetPointer(ptr)
		}
	}
	return n
}

//export callbackPersonModel_Constructor
func callbackPersonModel_Constructor(ptr unsafe.Pointer) {
	gPtr := NewPersonModelFromPointer(ptr)
	qt.Register(ptr, gPtr)
	gPtr.init()
}

//export callbackPersonModel_AddPerson
func callbackPersonModel_AddPerson(ptr unsafe.Pointer, v0 unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addPerson"); signal != nil {
		signal.(func(*Person))(NewPersonFromPointer(v0))
	}

}

func (ptr *PersonModel) ConnectAddPerson(f func(v0 *Person)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addPerson"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "addPerson", func(v0 *Person) {
				signal.(func(*Person))(v0)
				f(v0)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addPerson", f)
		}
	}
}

func (ptr *PersonModel) DisconnectAddPerson() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addPerson")
	}
}

func (ptr *PersonModel) AddPerson(v0 Person_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel_AddPerson(ptr.Pointer(), PointerFromPerson(v0))
	}
}

//export callbackPersonModel_EditPerson
func callbackPersonModel_EditPerson(ptr unsafe.Pointer, row C.int, firstName C.struct_Moc_PackedString, lastName C.struct_Moc_PackedString, email C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "editPerson"); signal != nil {
		signal.(func(int, string, string, string))(int(int32(row)), cGoUnpackString(firstName), cGoUnpackString(lastName), cGoUnpackString(email))
	}

}

func (ptr *PersonModel) ConnectEditPerson(f func(row int, firstName string, lastName string, email string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "editPerson"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "editPerson", func(row int, firstName string, lastName string, email string) {
				signal.(func(int, string, string, string))(row, firstName, lastName, email)
				f(row, firstName, lastName, email)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "editPerson", f)
		}
	}
}

func (ptr *PersonModel) DisconnectEditPerson() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "editPerson")
	}
}

func (ptr *PersonModel) EditPerson(row int, firstName string, lastName string, email string) {
	if ptr.Pointer() != nil {
		var firstNameC *C.char
		if firstName != "" {
			firstNameC = C.CString(firstName)
			defer C.free(unsafe.Pointer(firstNameC))
		}
		var lastNameC *C.char
		if lastName != "" {
			lastNameC = C.CString(lastName)
			defer C.free(unsafe.Pointer(lastNameC))
		}
		var emailC *C.char
		if email != "" {
			emailC = C.CString(email)
			defer C.free(unsafe.Pointer(emailC))
		}
		C.PersonModel_EditPerson(ptr.Pointer(), C.int(int32(row)), C.struct_Moc_PackedString{data: firstNameC, len: C.longlong(len(firstName))}, C.struct_Moc_PackedString{data: lastNameC, len: C.longlong(len(lastName))}, C.struct_Moc_PackedString{data: emailC, len: C.longlong(len(email))})
	}
}

//export callbackPersonModel_RemovePerson
func callbackPersonModel_RemovePerson(ptr unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "removePerson"); signal != nil {
		signal.(func(int))(int(int32(row)))
	}

}

func (ptr *PersonModel) ConnectRemovePerson(f func(row int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "removePerson"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "removePerson", func(row int) {
				signal.(func(int))(row)
				f(row)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "removePerson", f)
		}
	}
}

func (ptr *PersonModel) DisconnectRemovePerson() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "removePerson")
	}
}

func (ptr *PersonModel) RemovePerson(row int) {
	if ptr.Pointer() != nil {
		C.PersonModel_RemovePerson(ptr.Pointer(), C.int(int32(row)))
	}
}

//export callbackPersonModel_Roles
func callbackPersonModel_Roles(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roles"); signal != nil {
		return func() unsafe.Pointer {
			var tmpList = NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__roles_newList())
			for k, v := range signal.(func() map[int]*std_core.QByteArray)() {
				tmpList.__roles_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		var tmpList = NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__roles_newList())
		for k, v := range NewPersonModelFromPointer(ptr).RolesDefault() {
			tmpList.__roles_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *PersonModel) ConnectRoles(f func() map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "roles"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "roles", func() map[int]*std_core.QByteArray {
				signal.(func() map[int]*std_core.QByteArray)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "roles", f)
		}
	}
}

func (ptr *PersonModel) DisconnectRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "roles")
	}
}

func (ptr *PersonModel) Roles() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			var out = make(map[int]*std_core.QByteArray, int(l.len))
			for _, i := range NewPersonModelFromPointer(l.data).__roles_keyList() {
				out[i] = NewPersonModelFromPointer(l.data).__roles_atList(i)
			}
			return out
		}(C.PersonModel_Roles(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

func (ptr *PersonModel) RolesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			var out = make(map[int]*std_core.QByteArray, int(l.len))
			for _, i := range NewPersonModelFromPointer(l.data).__roles_keyList() {
				out[i] = NewPersonModelFromPointer(l.data).__roles_atList(i)
			}
			return out
		}(C.PersonModel_RolesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackPersonModel_SetRoles
func callbackPersonModel_SetRoles(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setRoles"); signal != nil {
		signal.(func(map[int]*std_core.QByteArray))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			var out = make(map[int]*std_core.QByteArray, int(l.len))
			for _, i := range NewPersonModelFromPointer(l.data).__setRoles_keyList() {
				out[i] = NewPersonModelFromPointer(l.data).__setRoles_roles_atList(i)
			}
			return out
		}(roles))
	} else {
		NewPersonModelFromPointer(ptr).SetRolesDefault(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			var out = make(map[int]*std_core.QByteArray, int(l.len))
			for _, i := range NewPersonModelFromPointer(l.data).__setRoles_keyList() {
				out[i] = NewPersonModelFromPointer(l.data).__setRoles_roles_atList(i)
			}
			return out
		}(roles))
	}
}

func (ptr *PersonModel) ConnectSetRoles(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRoles"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setRoles", func(roles map[int]*std_core.QByteArray) {
				signal.(func(map[int]*std_core.QByteArray))(roles)
				f(roles)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRoles", f)
		}
	}
}

func (ptr *PersonModel) DisconnectSetRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRoles")
	}
}

func (ptr *PersonModel) SetRoles(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.PersonModel_SetRoles(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *PersonModel) SetRolesDefault(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.PersonModel_SetRolesDefault(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackPersonModel_RolesChanged
func callbackPersonModel_RolesChanged(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "rolesChanged"); signal != nil {
		signal.(func(map[int]*std_core.QByteArray))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			var out = make(map[int]*std_core.QByteArray, int(l.len))
			for _, i := range NewPersonModelFromPointer(l.data).__rolesChanged_keyList() {
				out[i] = NewPersonModelFromPointer(l.data).__rolesChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

func (ptr *PersonModel) ConnectRolesChanged(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rolesChanged") {
			C.PersonModel_ConnectRolesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "rolesChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "rolesChanged", func(roles map[int]*std_core.QByteArray) {
				signal.(func(map[int]*std_core.QByteArray))(roles)
				f(roles)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rolesChanged", f)
		}
	}
}

func (ptr *PersonModel) DisconnectRolesChanged() {
	if ptr.Pointer() != nil {
		C.PersonModel_DisconnectRolesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rolesChanged")
	}
}

func (ptr *PersonModel) RolesChanged(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.PersonModel_RolesChanged(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__rolesChanged_roles_newList())
			for k, v := range roles {
				tmpList.__rolesChanged_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackPersonModel_People
func callbackPersonModel_People(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "people"); signal != nil {
		return func() unsafe.Pointer {
			var tmpList = NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__people_newList())
			for _, v := range signal.(func() []*Person)() {
				tmpList.__people_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		var tmpList = NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__people_newList())
		for _, v := range NewPersonModelFromPointer(ptr).PeopleDefault() {
			tmpList.__people_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *PersonModel) ConnectPeople(f func() []*Person) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "people"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "people", func() []*Person {
				signal.(func() []*Person)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "people", f)
		}
	}
}

func (ptr *PersonModel) DisconnectPeople() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "people")
	}
}

func (ptr *PersonModel) People() []*Person {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*Person {
			var out = make([]*Person, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewPersonModelFromPointer(l.data).__people_atList(i)
			}
			return out
		}(C.PersonModel_People(ptr.Pointer()))
	}
	return make([]*Person, 0)
}

func (ptr *PersonModel) PeopleDefault() []*Person {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*Person {
			var out = make([]*Person, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewPersonModelFromPointer(l.data).__people_atList(i)
			}
			return out
		}(C.PersonModel_PeopleDefault(ptr.Pointer()))
	}
	return make([]*Person, 0)
}

//export callbackPersonModel_SetPeople
func callbackPersonModel_SetPeople(ptr unsafe.Pointer, people C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setPeople"); signal != nil {
		signal.(func([]*Person))(func(l C.struct_Moc_PackedList) []*Person {
			var out = make([]*Person, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewPersonModelFromPointer(l.data).__setPeople_people_atList(i)
			}
			return out
		}(people))
	} else {
		NewPersonModelFromPointer(ptr).SetPeopleDefault(func(l C.struct_Moc_PackedList) []*Person {
			var out = make([]*Person, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewPersonModelFromPointer(l.data).__setPeople_people_atList(i)
			}
			return out
		}(people))
	}
}

func (ptr *PersonModel) ConnectSetPeople(f func(people []*Person)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setPeople"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setPeople", func(people []*Person) {
				signal.(func([]*Person))(people)
				f(people)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setPeople", f)
		}
	}
}

func (ptr *PersonModel) DisconnectSetPeople() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setPeople")
	}
}

func (ptr *PersonModel) SetPeople(people []*Person) {
	if ptr.Pointer() != nil {
		C.PersonModel_SetPeople(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__setPeople_people_newList())
			for _, v := range people {
				tmpList.__setPeople_people_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *PersonModel) SetPeopleDefault(people []*Person) {
	if ptr.Pointer() != nil {
		C.PersonModel_SetPeopleDefault(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__setPeople_people_newList())
			for _, v := range people {
				tmpList.__setPeople_people_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackPersonModel_PeopleChanged
func callbackPersonModel_PeopleChanged(ptr unsafe.Pointer, people C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "peopleChanged"); signal != nil {
		signal.(func([]*Person))(func(l C.struct_Moc_PackedList) []*Person {
			var out = make([]*Person, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewPersonModelFromPointer(l.data).__peopleChanged_people_atList(i)
			}
			return out
		}(people))
	}

}

func (ptr *PersonModel) ConnectPeopleChanged(f func(people []*Person)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "peopleChanged") {
			C.PersonModel_ConnectPeopleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "peopleChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "peopleChanged", func(people []*Person) {
				signal.(func([]*Person))(people)
				f(people)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "peopleChanged", f)
		}
	}
}

func (ptr *PersonModel) DisconnectPeopleChanged() {
	if ptr.Pointer() != nil {
		C.PersonModel_DisconnectPeopleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "peopleChanged")
	}
}

func (ptr *PersonModel) PeopleChanged(people []*Person) {
	if ptr.Pointer() != nil {
		C.PersonModel_PeopleChanged(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__peopleChanged_people_newList())
			for _, v := range people {
				tmpList.__peopleChanged_people_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func PersonModel_QRegisterMetaType() int {
	return int(int32(C.PersonModel_PersonModel_QRegisterMetaType()))
}

func (ptr *PersonModel) QRegisterMetaType() int {
	return int(int32(C.PersonModel_PersonModel_QRegisterMetaType()))
}

func PersonModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.PersonModel_PersonModel_QRegisterMetaType2(typeNameC)))
}

func (ptr *PersonModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.PersonModel_PersonModel_QRegisterMetaType2(typeNameC)))
}

func PersonModel_QmlRegisterType() int {
	return int(int32(C.PersonModel_PersonModel_QmlRegisterType()))
}

func (ptr *PersonModel) QmlRegisterType() int {
	return int(int32(C.PersonModel_PersonModel_QmlRegisterType()))
}

func PersonModel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.PersonModel_PersonModel_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *PersonModel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.PersonModel_PersonModel_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *PersonModel) ____setItemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.PersonModel_____setItemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *PersonModel) ____setItemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.PersonModel_____setItemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *PersonModel) ____setItemData_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel_____setItemData_keyList_newList(ptr.Pointer()))
}

func (ptr *PersonModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.PersonModel_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *PersonModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.PersonModel_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *PersonModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel_____roleNames_keyList_newList(ptr.Pointer()))
}

func (ptr *PersonModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.PersonModel_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *PersonModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.PersonModel_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *PersonModel) ____itemData_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel_____itemData_keyList_newList(ptr.Pointer()))
}

func (ptr *PersonModel) __setItemData_roles_atList(i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQVariantFromPointer(C.PersonModel___setItemData_roles_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *PersonModel) __setItemData_roles_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel___setItemData_roles_newList(ptr.Pointer()))
}

func (ptr *PersonModel) __setItemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			var out = make([]int, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewPersonModelFromPointer(l.data).____setItemData_keyList_atList(i)
			}
			return out
		}(C.PersonModel___setItemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *PersonModel) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.PersonModel___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *PersonModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel___changePersistentIndexList_from_newList(ptr.Pointer()))
}

func (ptr *PersonModel) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.PersonModel___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *PersonModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel___changePersistentIndexList_to_newList(ptr.Pointer()))
}

func (ptr *PersonModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.PersonModel___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *PersonModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.PersonModel___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *PersonModel) __dataChanged_roles_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel___dataChanged_roles_newList(ptr.Pointer()))
}

func (ptr *PersonModel) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQPersistentModelIndexFromPointer(C.PersonModel___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *PersonModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel___layoutAboutToBeChanged_parents_newList(ptr.Pointer()))
}

func (ptr *PersonModel) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQPersistentModelIndexFromPointer(C.PersonModel___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *PersonModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel___layoutChanged_parents_newList(ptr.Pointer()))
}

func (ptr *PersonModel) __roleNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.PersonModel___roleNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *PersonModel) __roleNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel___roleNames_newList(ptr.Pointer()))
}

func (ptr *PersonModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			var out = make([]int, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewPersonModelFromPointer(l.data).____roleNames_keyList_atList(i)
			}
			return out
		}(C.PersonModel___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *PersonModel) __itemData_atList(i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQVariantFromPointer(C.PersonModel___itemData_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *PersonModel) __itemData_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel___itemData_newList(ptr.Pointer()))
}

func (ptr *PersonModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			var out = make([]int, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewPersonModelFromPointer(l.data).____itemData_keyList_atList(i)
			}
			return out
		}(C.PersonModel___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *PersonModel) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.PersonModel___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *PersonModel) __mimeData_indexes_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel___mimeData_indexes_newList(ptr.Pointer()))
}

func (ptr *PersonModel) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.PersonModel___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *PersonModel) __match_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel___match_newList(ptr.Pointer()))
}

func (ptr *PersonModel) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.PersonModel___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *PersonModel) __persistentIndexList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel___persistentIndexList_newList(ptr.Pointer()))
}

func (ptr *PersonModel) ____doSetRoleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.PersonModel_____doSetRoleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *PersonModel) ____doSetRoleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.PersonModel_____doSetRoleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *PersonModel) ____doSetRoleNames_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel_____doSetRoleNames_keyList_newList(ptr.Pointer()))
}

func (ptr *PersonModel) ____setRoleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.PersonModel_____setRoleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *PersonModel) ____setRoleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.PersonModel_____setRoleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *PersonModel) ____setRoleNames_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel_____setRoleNames_keyList_newList(ptr.Pointer()))
}

func (ptr *PersonModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.PersonModel___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *PersonModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *PersonModel) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.PersonModel___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *PersonModel) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel___findChildren_newList2(ptr.Pointer()))
}

func (ptr *PersonModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.PersonModel___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *PersonModel) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel___findChildren_newList3(ptr.Pointer()))
}

func (ptr *PersonModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.PersonModel___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *PersonModel) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel___findChildren_newList(ptr.Pointer()))
}

func (ptr *PersonModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.PersonModel___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *PersonModel) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel___children_newList(ptr.Pointer()))
}

func (ptr *PersonModel) __roles_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.PersonModel___roles_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel___roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *PersonModel) __roles_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel___roles_newList(ptr.Pointer()))
}

func (ptr *PersonModel) __roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			var out = make([]int, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewPersonModelFromPointer(l.data).____roles_keyList_atList(i)
			}
			return out
		}(C.PersonModel___roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *PersonModel) __setRoles_roles_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.PersonModel___setRoles_roles_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __setRoles_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel___setRoles_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *PersonModel) __setRoles_roles_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel___setRoles_roles_newList(ptr.Pointer()))
}

func (ptr *PersonModel) __setRoles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			var out = make([]int, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewPersonModelFromPointer(l.data).____setRoles_keyList_atList(i)
			}
			return out
		}(C.PersonModel___setRoles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *PersonModel) __rolesChanged_roles_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.PersonModel___rolesChanged_roles_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __rolesChanged_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel___rolesChanged_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *PersonModel) __rolesChanged_roles_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel___rolesChanged_roles_newList(ptr.Pointer()))
}

func (ptr *PersonModel) __rolesChanged_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			var out = make([]int, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewPersonModelFromPointer(l.data).____rolesChanged_keyList_atList(i)
			}
			return out
		}(C.PersonModel___rolesChanged_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *PersonModel) __people_atList(i int) *Person {
	if ptr.Pointer() != nil {
		var tmpValue = NewPersonFromPointer(C.PersonModel___people_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __people_setList(i Person_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel___people_setList(ptr.Pointer(), PointerFromPerson(i))
	}
}

func (ptr *PersonModel) __people_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel___people_newList(ptr.Pointer()))
}

func (ptr *PersonModel) __setPeople_people_atList(i int) *Person {
	if ptr.Pointer() != nil {
		var tmpValue = NewPersonFromPointer(C.PersonModel___setPeople_people_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __setPeople_people_setList(i Person_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel___setPeople_people_setList(ptr.Pointer(), PointerFromPerson(i))
	}
}

func (ptr *PersonModel) __setPeople_people_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel___setPeople_people_newList(ptr.Pointer()))
}

func (ptr *PersonModel) __peopleChanged_people_atList(i int) *Person {
	if ptr.Pointer() != nil {
		var tmpValue = NewPersonFromPointer(C.PersonModel___peopleChanged_people_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *PersonModel) __peopleChanged_people_setList(i Person_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel___peopleChanged_people_setList(ptr.Pointer(), PointerFromPerson(i))
	}
}

func (ptr *PersonModel) __peopleChanged_people_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel___peopleChanged_people_newList(ptr.Pointer()))
}

func (ptr *PersonModel) ____roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.PersonModel_____roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *PersonModel) ____roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.PersonModel_____roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *PersonModel) ____roles_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel_____roles_keyList_newList(ptr.Pointer()))
}

func (ptr *PersonModel) ____setRoles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.PersonModel_____setRoles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *PersonModel) ____setRoles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.PersonModel_____setRoles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *PersonModel) ____setRoles_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel_____setRoles_keyList_newList(ptr.Pointer()))
}

func (ptr *PersonModel) ____rolesChanged_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.PersonModel_____rolesChanged_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *PersonModel) ____rolesChanged_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.PersonModel_____rolesChanged_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *PersonModel) ____rolesChanged_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.PersonModel_____rolesChanged_keyList_newList(ptr.Pointer()))
}

func NewPersonModel(parent std_core.QObject_ITF) *PersonModel {
	var tmpValue = NewPersonModelFromPointer(C.PersonModel_NewPersonModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *PersonModel) DestroyPersonModel() {
	if ptr.Pointer() != nil {
		C.PersonModel_DestroyPersonModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackPersonModel_DropMimeData
func callbackPersonModel_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *PersonModel) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.PersonModel_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackPersonModel_Index
func callbackPersonModel_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewPersonModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *PersonModel) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.PersonModel_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackPersonModel_Sibling
func callbackPersonModel_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewPersonModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *PersonModel) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.PersonModel_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackPersonModel_Flags
func callbackPersonModel_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong(signal.(func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewPersonModelFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *PersonModel) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.PersonModel_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackPersonModel_InsertColumns
func callbackPersonModel_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *PersonModel) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.PersonModel_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackPersonModel_InsertRows
func callbackPersonModel_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *PersonModel) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.PersonModel_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackPersonModel_MoveColumns
func callbackPersonModel_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *PersonModel) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return C.PersonModel_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

//export callbackPersonModel_MoveRows
func callbackPersonModel_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *PersonModel) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return C.PersonModel_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

//export callbackPersonModel_RemoveColumns
func callbackPersonModel_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *PersonModel) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.PersonModel_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackPersonModel_RemoveRows
func callbackPersonModel_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *PersonModel) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.PersonModel_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackPersonModel_SetData
func callbackPersonModel_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *PersonModel) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return C.PersonModel_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

//export callbackPersonModel_SetHeaderData
func callbackPersonModel_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *PersonModel) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return C.PersonModel_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

//export callbackPersonModel_SetItemData
func callbackPersonModel_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, map[int]*std_core.QVariant) bool)(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			var out = make(map[int]*std_core.QVariant, int(l.len))
			for _, i := range NewPersonModelFromPointer(l.data).__setItemData_keyList() {
				out[i] = NewPersonModelFromPointer(l.data).__setItemData_roles_atList(i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).SetItemDataDefault(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
		var out = make(map[int]*std_core.QVariant, int(l.len))
		for _, i := range NewPersonModelFromPointer(l.data).__setItemData_keyList() {
			out[i] = NewPersonModelFromPointer(l.data).__setItemData_roles_atList(i)
		}
		return out
	}(roles)))))
}

func (ptr *PersonModel) SetItemDataDefault(index std_core.QModelIndex_ITF, roles map[int]*std_core.QVariant) bool {
	if ptr.Pointer() != nil {
		return C.PersonModel_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			var tmpList = NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}()) != 0
	}
	return false
}

//export callbackPersonModel_Submit
func callbackPersonModel_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *PersonModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return C.PersonModel_SubmitDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackPersonModel_ColumnsAboutToBeInserted
func callbackPersonModel_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackPersonModel_ColumnsAboutToBeMoved
func callbackPersonModel_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackPersonModel_ColumnsAboutToBeRemoved
func callbackPersonModel_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackPersonModel_ColumnsInserted
func callbackPersonModel_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackPersonModel_ColumnsMoved
func callbackPersonModel_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackPersonModel_ColumnsRemoved
func callbackPersonModel_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackPersonModel_DataChanged
func callbackPersonModel_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		signal.(func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			var out = make([]int, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewPersonModelFromPointer(l.data).__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackPersonModel_FetchMore
func callbackPersonModel_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		signal.(func(*std_core.QModelIndex))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewPersonModelFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *PersonModel) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackPersonModel_HeaderDataChanged
func callbackPersonModel_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		signal.(func(std_core.Qt__Orientation, int, int))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackPersonModel_LayoutAboutToBeChanged
func callbackPersonModel_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		signal.(func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			var out = make([]*std_core.QPersistentModelIndex, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewPersonModelFromPointer(l.data).__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackPersonModel_LayoutChanged
func callbackPersonModel_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		signal.(func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			var out = make([]*std_core.QPersistentModelIndex, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewPersonModelFromPointer(l.data).__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackPersonModel_ModelAboutToBeReset
func callbackPersonModel_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackPersonModel_ModelReset
func callbackPersonModel_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackPersonModel_ResetInternalData
func callbackPersonModel_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		signal.(func())()
	} else {
		NewPersonModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *PersonModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.PersonModel_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackPersonModel_Revert
func callbackPersonModel_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		signal.(func())()
	} else {
		NewPersonModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *PersonModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.PersonModel_RevertDefault(ptr.Pointer())
	}
}

//export callbackPersonModel_RowsAboutToBeInserted
func callbackPersonModel_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackPersonModel_RowsAboutToBeMoved
func callbackPersonModel_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackPersonModel_RowsAboutToBeRemoved
func callbackPersonModel_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackPersonModel_RowsInserted
func callbackPersonModel_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackPersonModel_RowsMoved
func callbackPersonModel_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackPersonModel_RowsRemoved
func callbackPersonModel_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackPersonModel_Sort
func callbackPersonModel_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		signal.(func(int, std_core.Qt__SortOrder))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewPersonModelFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *PersonModel) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.PersonModel_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackPersonModel_RoleNames
func callbackPersonModel_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			var tmpList = NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__roleNames_newList())
			for k, v := range signal.(func() map[int]*std_core.QByteArray)() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		var tmpList = NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__roleNames_newList())
		for k, v := range NewPersonModelFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *PersonModel) RoleNamesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			var out = make(map[int]*std_core.QByteArray, int(l.len))
			for _, i := range NewPersonModelFromPointer(l.data).__roleNames_keyList() {
				out[i] = NewPersonModelFromPointer(l.data).__roleNames_atList(i)
			}
			return out
		}(C.PersonModel_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackPersonModel_ItemData
func callbackPersonModel_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			var tmpList = NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__itemData_newList())
			for k, v := range signal.(func(*std_core.QModelIndex) map[int]*std_core.QVariant)(std_core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		var tmpList = NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__itemData_newList())
		for k, v := range NewPersonModelFromPointer(ptr).ItemDataDefault(std_core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *PersonModel) ItemDataDefault(index std_core.QModelIndex_ITF) map[int]*std_core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			var out = make(map[int]*std_core.QVariant, int(l.len))
			for _, i := range NewPersonModelFromPointer(l.data).__itemData_keyList() {
				out[i] = NewPersonModelFromPointer(l.data).__itemData_atList(i)
			}
			return out
		}(C.PersonModel_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackPersonModel_MimeData
func callbackPersonModel_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData(signal.(func([]*std_core.QModelIndex) *std_core.QMimeData)(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			var out = make([]*std_core.QModelIndex, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewPersonModelFromPointer(l.data).__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return std_core.PointerFromQMimeData(NewPersonModelFromPointer(ptr).MimeDataDefault(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
		var out = make([]*std_core.QModelIndex, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewPersonModelFromPointer(l.data).__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *PersonModel) MimeDataDefault(indexes []*std_core.QModelIndex) *std_core.QMimeData {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQMimeDataFromPointer(C.PersonModel_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__mimeData_indexes_newList())
			for _, v := range indexes {
				tmpList.__mimeData_indexes_setList(v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackPersonModel_Buddy
func callbackPersonModel_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewPersonModelFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *PersonModel) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.PersonModel_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackPersonModel_Parent
func callbackPersonModel_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewPersonModelFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *PersonModel) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQModelIndexFromPointer(C.PersonModel_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackPersonModel_Match
func callbackPersonModel_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			var tmpList = NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__match_newList())
			for _, v := range signal.(func(*std_core.QModelIndex, int, *std_core.QVariant, int, std_core.Qt__MatchFlag) []*std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		var tmpList = NewPersonModelFromPointer(NewPersonModelFromPointer(nil).__match_newList())
		for _, v := range NewPersonModelFromPointer(ptr).MatchDefault(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *PersonModel) MatchDefault(start std_core.QModelIndex_ITF, role int, value std_core.QVariant_ITF, hits int, flags std_core.Qt__MatchFlag) []*std_core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			var out = make([]*std_core.QModelIndex, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewPersonModelFromPointer(l.data).__match_atList(i)
			}
			return out
		}(C.PersonModel_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackPersonModel_Span
func callbackPersonModel_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize(signal.(func(*std_core.QModelIndex) *std_core.QSize)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewPersonModelFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *PersonModel) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQSizeFromPointer(C.PersonModel_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackPersonModel_MimeTypes
func callbackPersonModel_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := NewPersonModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *PersonModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.PersonModel_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackPersonModel_Data
func callbackPersonModel_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(*std_core.QModelIndex, int) *std_core.QVariant)(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewPersonModelFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *PersonModel) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQVariantFromPointer(C.PersonModel_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackPersonModel_HeaderData
func callbackPersonModel_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewPersonModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *PersonModel) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQVariantFromPointer(C.PersonModel_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackPersonModel_SupportedDragActions
func callbackPersonModel_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewPersonModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *PersonModel) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.PersonModel_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackPersonModel_SupportedDropActions
func callbackPersonModel_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewPersonModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *PersonModel) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.PersonModel_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackPersonModel_CanDropMimeData
func callbackPersonModel_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *PersonModel) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.PersonModel_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackPersonModel_CanFetchMore
func callbackPersonModel_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *PersonModel) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.PersonModel_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackPersonModel_HasChildren
func callbackPersonModel_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *PersonModel) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return C.PersonModel_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackPersonModel_ColumnCount
func callbackPersonModel_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewPersonModelFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *PersonModel) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.PersonModel_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackPersonModel_RowCount
func callbackPersonModel_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewPersonModelFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *PersonModel) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.PersonModel_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackPersonModel_Event
func callbackPersonModel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *PersonModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.PersonModel_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackPersonModel_EventFilter
func callbackPersonModel_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *PersonModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.PersonModel_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackPersonModel_ChildEvent
func callbackPersonModel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewPersonModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *PersonModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackPersonModel_ConnectNotify
func callbackPersonModel_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewPersonModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *PersonModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackPersonModel_CustomEvent
func callbackPersonModel_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewPersonModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *PersonModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackPersonModel_DeleteLater
func callbackPersonModel_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewPersonModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *PersonModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.PersonModel_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackPersonModel_Destroyed
func callbackPersonModel_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackPersonModel_DisconnectNotify
func callbackPersonModel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewPersonModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *PersonModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackPersonModel_ObjectNameChanged
func callbackPersonModel_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackPersonModel_TimerEvent
func callbackPersonModel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewPersonModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *PersonModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PersonModel_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
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

type Person_ITF interface {
	std_core.QObject_ITF
	Person_PTR() *Person
}

func (ptr *Person) Person_PTR() *Person {
	return ptr
}

func (ptr *Person) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *Person) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromPerson(ptr Person_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Person_PTR().Pointer()
	}
	return nil
}

func NewPersonFromPointer(ptr unsafe.Pointer) *Person {
	var n *Person
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(Person)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *Person:
			n = deduced

		case *std_core.QObject:
			n = &Person{QObject: *deduced}

		default:
			n = new(Person)
			n.SetPointer(ptr)
		}
	}
	return n
}

//export callbackPerson_Constructor
func callbackPerson_Constructor(ptr unsafe.Pointer) {
	gPtr := NewPersonFromPointer(ptr)
	qt.Register(ptr, gPtr)
}

//export callbackPerson_FirstName
func callbackPerson_FirstName(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "firstName"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewPersonFromPointer(ptr).FirstNameDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *Person) ConnectFirstName(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "firstName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "firstName", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "firstName", f)
		}
	}
}

func (ptr *Person) DisconnectFirstName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "firstName")
	}
}

func (ptr *Person) FirstName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Person_FirstName(ptr.Pointer()))
	}
	return ""
}

func (ptr *Person) FirstNameDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Person_FirstNameDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackPerson_SetFirstName
func callbackPerson_SetFirstName(ptr unsafe.Pointer, firstName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setFirstName"); signal != nil {
		signal.(func(string))(cGoUnpackString(firstName))
	} else {
		NewPersonFromPointer(ptr).SetFirstNameDefault(cGoUnpackString(firstName))
	}
}

func (ptr *Person) ConnectSetFirstName(f func(firstName string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFirstName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setFirstName", func(firstName string) {
				signal.(func(string))(firstName)
				f(firstName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFirstName", f)
		}
	}
}

func (ptr *Person) DisconnectSetFirstName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFirstName")
	}
}

func (ptr *Person) SetFirstName(firstName string) {
	if ptr.Pointer() != nil {
		var firstNameC *C.char
		if firstName != "" {
			firstNameC = C.CString(firstName)
			defer C.free(unsafe.Pointer(firstNameC))
		}
		C.Person_SetFirstName(ptr.Pointer(), C.struct_Moc_PackedString{data: firstNameC, len: C.longlong(len(firstName))})
	}
}

func (ptr *Person) SetFirstNameDefault(firstName string) {
	if ptr.Pointer() != nil {
		var firstNameC *C.char
		if firstName != "" {
			firstNameC = C.CString(firstName)
			defer C.free(unsafe.Pointer(firstNameC))
		}
		C.Person_SetFirstNameDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: firstNameC, len: C.longlong(len(firstName))})
	}
}

//export callbackPerson_FirstNameChanged
func callbackPerson_FirstNameChanged(ptr unsafe.Pointer, firstName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "firstNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(firstName))
	}

}

func (ptr *Person) ConnectFirstNameChanged(f func(firstName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "firstNameChanged") {
			C.Person_ConnectFirstNameChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "firstNameChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "firstNameChanged", func(firstName string) {
				signal.(func(string))(firstName)
				f(firstName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "firstNameChanged", f)
		}
	}
}

func (ptr *Person) DisconnectFirstNameChanged() {
	if ptr.Pointer() != nil {
		C.Person_DisconnectFirstNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "firstNameChanged")
	}
}

func (ptr *Person) FirstNameChanged(firstName string) {
	if ptr.Pointer() != nil {
		var firstNameC *C.char
		if firstName != "" {
			firstNameC = C.CString(firstName)
			defer C.free(unsafe.Pointer(firstNameC))
		}
		C.Person_FirstNameChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: firstNameC, len: C.longlong(len(firstName))})
	}
}

//export callbackPerson_LastName
func callbackPerson_LastName(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "lastName"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewPersonFromPointer(ptr).LastNameDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *Person) ConnectLastName(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "lastName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "lastName", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lastName", f)
		}
	}
}

func (ptr *Person) DisconnectLastName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "lastName")
	}
}

func (ptr *Person) LastName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Person_LastName(ptr.Pointer()))
	}
	return ""
}

func (ptr *Person) LastNameDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Person_LastNameDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackPerson_SetLastName
func callbackPerson_SetLastName(ptr unsafe.Pointer, lastName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setLastName"); signal != nil {
		signal.(func(string))(cGoUnpackString(lastName))
	} else {
		NewPersonFromPointer(ptr).SetLastNameDefault(cGoUnpackString(lastName))
	}
}

func (ptr *Person) ConnectSetLastName(f func(lastName string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setLastName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setLastName", func(lastName string) {
				signal.(func(string))(lastName)
				f(lastName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setLastName", f)
		}
	}
}

func (ptr *Person) DisconnectSetLastName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setLastName")
	}
}

func (ptr *Person) SetLastName(lastName string) {
	if ptr.Pointer() != nil {
		var lastNameC *C.char
		if lastName != "" {
			lastNameC = C.CString(lastName)
			defer C.free(unsafe.Pointer(lastNameC))
		}
		C.Person_SetLastName(ptr.Pointer(), C.struct_Moc_PackedString{data: lastNameC, len: C.longlong(len(lastName))})
	}
}

func (ptr *Person) SetLastNameDefault(lastName string) {
	if ptr.Pointer() != nil {
		var lastNameC *C.char
		if lastName != "" {
			lastNameC = C.CString(lastName)
			defer C.free(unsafe.Pointer(lastNameC))
		}
		C.Person_SetLastNameDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: lastNameC, len: C.longlong(len(lastName))})
	}
}

//export callbackPerson_LastNameChanged
func callbackPerson_LastNameChanged(ptr unsafe.Pointer, lastName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "lastNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(lastName))
	}

}

func (ptr *Person) ConnectLastNameChanged(f func(lastName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "lastNameChanged") {
			C.Person_ConnectLastNameChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "lastNameChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "lastNameChanged", func(lastName string) {
				signal.(func(string))(lastName)
				f(lastName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lastNameChanged", f)
		}
	}
}

func (ptr *Person) DisconnectLastNameChanged() {
	if ptr.Pointer() != nil {
		C.Person_DisconnectLastNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "lastNameChanged")
	}
}

func (ptr *Person) LastNameChanged(lastName string) {
	if ptr.Pointer() != nil {
		var lastNameC *C.char
		if lastName != "" {
			lastNameC = C.CString(lastName)
			defer C.free(unsafe.Pointer(lastNameC))
		}
		C.Person_LastNameChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: lastNameC, len: C.longlong(len(lastName))})
	}
}

//export callbackPerson_Email
func callbackPerson_Email(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "email"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewPersonFromPointer(ptr).EmailDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *Person) ConnectEmail(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "email"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "email", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "email", f)
		}
	}
}

func (ptr *Person) DisconnectEmail() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "email")
	}
}

func (ptr *Person) Email() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Person_Email(ptr.Pointer()))
	}
	return ""
}

func (ptr *Person) EmailDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Person_EmailDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackPerson_SetEmail
func callbackPerson_SetEmail(ptr unsafe.Pointer, email C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setEmail"); signal != nil {
		signal.(func(string))(cGoUnpackString(email))
	} else {
		NewPersonFromPointer(ptr).SetEmailDefault(cGoUnpackString(email))
	}
}

func (ptr *Person) ConnectSetEmail(f func(email string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setEmail"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setEmail", func(email string) {
				signal.(func(string))(email)
				f(email)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setEmail", f)
		}
	}
}

func (ptr *Person) DisconnectSetEmail() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setEmail")
	}
}

func (ptr *Person) SetEmail(email string) {
	if ptr.Pointer() != nil {
		var emailC *C.char
		if email != "" {
			emailC = C.CString(email)
			defer C.free(unsafe.Pointer(emailC))
		}
		C.Person_SetEmail(ptr.Pointer(), C.struct_Moc_PackedString{data: emailC, len: C.longlong(len(email))})
	}
}

func (ptr *Person) SetEmailDefault(email string) {
	if ptr.Pointer() != nil {
		var emailC *C.char
		if email != "" {
			emailC = C.CString(email)
			defer C.free(unsafe.Pointer(emailC))
		}
		C.Person_SetEmailDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: emailC, len: C.longlong(len(email))})
	}
}

//export callbackPerson_EmailChanged
func callbackPerson_EmailChanged(ptr unsafe.Pointer, email C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "emailChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(email))
	}

}

func (ptr *Person) ConnectEmailChanged(f func(email string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "emailChanged") {
			C.Person_ConnectEmailChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "emailChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "emailChanged", func(email string) {
				signal.(func(string))(email)
				f(email)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "emailChanged", f)
		}
	}
}

func (ptr *Person) DisconnectEmailChanged() {
	if ptr.Pointer() != nil {
		C.Person_DisconnectEmailChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "emailChanged")
	}
}

func (ptr *Person) EmailChanged(email string) {
	if ptr.Pointer() != nil {
		var emailC *C.char
		if email != "" {
			emailC = C.CString(email)
			defer C.free(unsafe.Pointer(emailC))
		}
		C.Person_EmailChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: emailC, len: C.longlong(len(email))})
	}
}

func Person_QRegisterMetaType() int {
	return int(int32(C.Person_Person_QRegisterMetaType()))
}

func (ptr *Person) QRegisterMetaType() int {
	return int(int32(C.Person_Person_QRegisterMetaType()))
}

func Person_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Person_Person_QRegisterMetaType2(typeNameC)))
}

func (ptr *Person) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Person_Person_QRegisterMetaType2(typeNameC)))
}

func Person_QmlRegisterType() int {
	return int(int32(C.Person_Person_QmlRegisterType()))
}

func (ptr *Person) QmlRegisterType() int {
	return int(int32(C.Person_Person_QmlRegisterType()))
}

func Person_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.Person_Person_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Person) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.Person_Person_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Person) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.Person___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *Person) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.Person___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *Person) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.Person___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *Person) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.Person___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Person) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Person___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Person) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.Person___findChildren_newList2(ptr.Pointer()))
}

func (ptr *Person) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.Person___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Person) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Person___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Person) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.Person___findChildren_newList3(ptr.Pointer()))
}

func (ptr *Person) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.Person___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Person) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Person___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Person) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.Person___findChildren_newList(ptr.Pointer()))
}

func (ptr *Person) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.Person___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Person) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Person___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Person) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.Person___children_newList(ptr.Pointer()))
}

func NewPerson(parent std_core.QObject_ITF) *Person {
	var tmpValue = NewPersonFromPointer(C.Person_NewPerson(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackPerson_DestroyPerson
func callbackPerson_DestroyPerson(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~Person"); signal != nil {
		signal.(func())()
	} else {
		NewPersonFromPointer(ptr).DestroyPersonDefault()
	}
}

func (ptr *Person) ConnectDestroyPerson(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~Person"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~Person", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~Person", f)
		}
	}
}

func (ptr *Person) DisconnectDestroyPerson() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~Person")
	}
}

func (ptr *Person) DestroyPerson() {
	if ptr.Pointer() != nil {
		C.Person_DestroyPerson(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *Person) DestroyPersonDefault() {
	if ptr.Pointer() != nil {
		C.Person_DestroyPersonDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackPerson_Event
func callbackPerson_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *Person) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.Person_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackPerson_EventFilter
func callbackPerson_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPersonFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *Person) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.Person_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackPerson_ChildEvent
func callbackPerson_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewPersonFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *Person) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Person_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackPerson_ConnectNotify
func callbackPerson_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewPersonFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Person) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Person_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackPerson_CustomEvent
func callbackPerson_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewPersonFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Person) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Person_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackPerson_DeleteLater
func callbackPerson_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewPersonFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *Person) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.Person_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackPerson_Destroyed
func callbackPerson_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackPerson_DisconnectNotify
func callbackPerson_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewPersonFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Person) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Person_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackPerson_ObjectNameChanged
func callbackPerson_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackPerson_TimerEvent
func callbackPerson_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewPersonFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *Person) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Person_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
