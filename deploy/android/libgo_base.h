/* Created by "go tool cgo" - DO NOT EDIT. */

/* package github.com/amlwwalker/qt-hotreloader */

/* Start of preamble from import "C" comments.  */



#line 3 "/media/sf_GOPATH0/src/github.com/amlwwalker/qt-hotreloader/moc.go"
#include <stdint.h>
#include <stdlib.h>
#include <string.h>
#include "moc.h"

#line 1 "cgo-generated-wrapper"


/* End of preamble from import "C" comments.  */


/* Start of boilerplate cgo prologue.  */
#line 1 "cgo-gcc-export-header-prolog"

#ifndef GO_CGO_PROLOGUE_H
#define GO_CGO_PROLOGUE_H

typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt32 GoInt;
typedef GoUint32 GoUint;
typedef __SIZE_TYPE__ GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;
typedef float _Complex GoComplex64;
typedef double _Complex GoComplex128;

/*
  static assertion to make sure the file is being used on architecture
  at least with matching size of GoInt.
*/
typedef char _check_for_32_bit_pointer_matching_GoInt[sizeof(void*)==32/8 ? 1:-1];

typedef struct { const char *p; GoInt n; } GoString;
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;

#endif

/* End of boilerplate cgo prologue.  */

#ifdef __cplusplus
extern "C" {
#endif


extern void go_main_wrapper();

extern void callbackPerson_Constructor(void* p0);

extern struct Moc_PackedString callbackPerson_FirstName(void* p0);

extern void callbackPerson_SetFirstName(void* p0, struct Moc_PackedString p1);

extern void callbackPerson_FirstNameChanged(void* p0, struct Moc_PackedString p1);

extern struct Moc_PackedString callbackPerson_LastName(void* p0);

extern void callbackPerson_SetLastName(void* p0, struct Moc_PackedString p1);

extern void callbackPerson_LastNameChanged(void* p0, struct Moc_PackedString p1);

extern struct Moc_PackedString callbackPerson_Email(void* p0);

extern void callbackPerson_SetEmail(void* p0, struct Moc_PackedString p1);

extern void callbackPerson_EmailChanged(void* p0, struct Moc_PackedString p1);

extern void callbackPerson_DestroyPerson(void* p0);

extern char callbackPerson_Event(void* p0, void* p1);

extern char callbackPerson_EventFilter(void* p0, void* p1, void* p2);

extern void callbackPerson_ChildEvent(void* p0, void* p1);

extern void callbackPerson_ConnectNotify(void* p0, void* p1);

extern void callbackPerson_CustomEvent(void* p0, void* p1);

extern void callbackPerson_DeleteLater(void* p0);

extern void callbackPerson_Destroyed(void* p0, void* p1);

extern void callbackPerson_DisconnectNotify(void* p0, void* p1);

extern void callbackPerson_ObjectNameChanged(void* p0, struct Moc_PackedString p1);

extern void callbackPerson_TimerEvent(void* p0, void* p1);

extern void callbackPersonModel_Constructor(void* p0);

extern void callbackPersonModel_AddPerson(void* p0, void* p1);

extern void callbackPersonModel_EditPerson(void* p0, int p1, struct Moc_PackedString p2, struct Moc_PackedString p3, struct Moc_PackedString p4);

extern void callbackPersonModel_RemovePerson(void* p0, int p1);

extern void* callbackPersonModel_Roles(void* p0);

extern void callbackPersonModel_SetRoles(void* p0, struct Moc_PackedList p1);

extern void callbackPersonModel_RolesChanged(void* p0, struct Moc_PackedList p1);

extern void* callbackPersonModel_People(void* p0);

extern void callbackPersonModel_SetPeople(void* p0, struct Moc_PackedList p1);

extern void callbackPersonModel_PeopleChanged(void* p0, struct Moc_PackedList p1);

extern char callbackPersonModel_DropMimeData(void* p0, void* p1, long long int p2, int p3, int p4, void* p5);

extern void* callbackPersonModel_Index(void* p0, int p1, int p2, void* p3);

extern void* callbackPersonModel_Sibling(void* p0, int p1, int p2, void* p3);

extern long long int callbackPersonModel_Flags(void* p0, void* p1);

extern char callbackPersonModel_InsertColumns(void* p0, int p1, int p2, void* p3);

extern char callbackPersonModel_InsertRows(void* p0, int p1, int p2, void* p3);

extern char callbackPersonModel_MoveColumns(void* p0, void* p1, int p2, int p3, void* p4, int p5);

extern char callbackPersonModel_MoveRows(void* p0, void* p1, int p2, int p3, void* p4, int p5);

extern char callbackPersonModel_RemoveColumns(void* p0, int p1, int p2, void* p3);

extern char callbackPersonModel_RemoveRows(void* p0, int p1, int p2, void* p3);

extern char callbackPersonModel_SetData(void* p0, void* p1, void* p2, int p3);

extern char callbackPersonModel_SetHeaderData(void* p0, int p1, long long int p2, void* p3, int p4);

extern char callbackPersonModel_SetItemData(void* p0, void* p1, struct Moc_PackedList p2);

extern char callbackPersonModel_Submit(void* p0);

extern void callbackPersonModel_ColumnsAboutToBeInserted(void* p0, void* p1, int p2, int p3);

extern void callbackPersonModel_ColumnsAboutToBeMoved(void* p0, void* p1, int p2, int p3, void* p4, int p5);

extern void callbackPersonModel_ColumnsAboutToBeRemoved(void* p0, void* p1, int p2, int p3);

extern void callbackPersonModel_ColumnsInserted(void* p0, void* p1, int p2, int p3);

extern void callbackPersonModel_ColumnsMoved(void* p0, void* p1, int p2, int p3, void* p4, int p5);

extern void callbackPersonModel_ColumnsRemoved(void* p0, void* p1, int p2, int p3);

extern void callbackPersonModel_DataChanged(void* p0, void* p1, void* p2, struct Moc_PackedList p3);

extern void callbackPersonModel_FetchMore(void* p0, void* p1);

extern void callbackPersonModel_HeaderDataChanged(void* p0, long long int p1, int p2, int p3);

extern void callbackPersonModel_LayoutAboutToBeChanged(void* p0, struct Moc_PackedList p1, long long int p2);

extern void callbackPersonModel_LayoutChanged(void* p0, struct Moc_PackedList p1, long long int p2);

extern void callbackPersonModel_ModelAboutToBeReset(void* p0);

extern void callbackPersonModel_ModelReset(void* p0);

extern void callbackPersonModel_ResetInternalData(void* p0);

extern void callbackPersonModel_Revert(void* p0);

extern void callbackPersonModel_RowsAboutToBeInserted(void* p0, void* p1, int p2, int p3);

extern void callbackPersonModel_RowsAboutToBeMoved(void* p0, void* p1, int p2, int p3, void* p4, int p5);

extern void callbackPersonModel_RowsAboutToBeRemoved(void* p0, void* p1, int p2, int p3);

extern void callbackPersonModel_RowsInserted(void* p0, void* p1, int p2, int p3);

extern void callbackPersonModel_RowsMoved(void* p0, void* p1, int p2, int p3, void* p4, int p5);

extern void callbackPersonModel_RowsRemoved(void* p0, void* p1, int p2, int p3);

extern void callbackPersonModel_Sort(void* p0, int p1, long long int p2);

extern void* callbackPersonModel_RoleNames(void* p0);

extern void* callbackPersonModel_ItemData(void* p0, void* p1);

extern void* callbackPersonModel_MimeData(void* p0, struct Moc_PackedList p1);

extern void* callbackPersonModel_Buddy(void* p0, void* p1);

extern void* callbackPersonModel_Parent(void* p0, void* p1);

extern void* callbackPersonModel_Match(void* p0, void* p1, int p2, void* p3, int p4, long long int p5);

extern void* callbackPersonModel_Span(void* p0, void* p1);

extern struct Moc_PackedString callbackPersonModel_MimeTypes(void* p0);

extern void* callbackPersonModel_Data(void* p0, void* p1, int p2);

extern void* callbackPersonModel_HeaderData(void* p0, int p1, long long int p2, int p3);

extern long long int callbackPersonModel_SupportedDragActions(void* p0);

extern long long int callbackPersonModel_SupportedDropActions(void* p0);

extern char callbackPersonModel_CanDropMimeData(void* p0, void* p1, long long int p2, int p3, int p4, void* p5);

extern char callbackPersonModel_CanFetchMore(void* p0, void* p1);

extern char callbackPersonModel_HasChildren(void* p0, void* p1);

extern int callbackPersonModel_ColumnCount(void* p0, void* p1);

extern int callbackPersonModel_RowCount(void* p0, void* p1);

extern char callbackPersonModel_Event(void* p0, void* p1);

extern char callbackPersonModel_EventFilter(void* p0, void* p1, void* p2);

extern void callbackPersonModel_ChildEvent(void* p0, void* p1);

extern void callbackPersonModel_ConnectNotify(void* p0, void* p1);

extern void callbackPersonModel_CustomEvent(void* p0, void* p1);

extern void callbackPersonModel_DeleteLater(void* p0);

extern void callbackPersonModel_Destroyed(void* p0, void* p1);

extern void callbackPersonModel_DisconnectNotify(void* p0, void* p1);

extern void callbackPersonModel_ObjectNameChanged(void* p0, struct Moc_PackedString p1);

extern void callbackPersonModel_TimerEvent(void* p0, void* p1);

extern void callbackQmlBridge_Constructor(void* p0);

extern void callbackQmlBridge_UpdateLoader(void* p0, struct Moc_PackedString p1);

extern void callbackQmlBridge_UpdateSettings(void* p0, struct Moc_PackedString p1, struct Moc_PackedString p2, struct Moc_PackedString p3, struct Moc_PackedString p4, struct Moc_PackedString p5, struct Moc_PackedString p6, char p7);

extern void callbackQmlBridge_SendTime(void* p0, struct Moc_PackedString p1);

extern struct Moc_PackedString callbackQmlBridge_Calculator(void* p0, struct Moc_PackedString p1, struct Moc_PackedString p2);

extern void callbackQmlBridge_DestroyQmlBridge(void* p0);

extern char callbackQmlBridge_Event(void* p0, void* p1);

extern char callbackQmlBridge_EventFilter(void* p0, void* p1, void* p2);

extern void callbackQmlBridge_ChildEvent(void* p0, void* p1);

extern void callbackQmlBridge_ConnectNotify(void* p0, void* p1);

extern void callbackQmlBridge_CustomEvent(void* p0, void* p1);

extern void callbackQmlBridge_DeleteLater(void* p0);

extern void callbackQmlBridge_Destroyed(void* p0, void* p1);

extern void callbackQmlBridge_DisconnectNotify(void* p0, void* p1);

extern void callbackQmlBridge_ObjectNameChanged(void* p0, struct Moc_PackedString p1);

extern void callbackQmlBridge_TimerEvent(void* p0, void* p1);

#ifdef __cplusplus
}
#endif
