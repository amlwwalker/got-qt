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

extern void callbackQmlBridge_Constructor(void* p0);

extern void callbackQmlBridge_UpdateLoader(void* p0, struct Moc_PackedString p1);

extern void callbackQmlBridge_UpdateSettings(void* p0, struct Moc_PackedString p1, struct Moc_PackedString p2, struct Moc_PackedString p3, struct Moc_PackedString p4, struct Moc_PackedString p5, struct Moc_PackedString p6, char p7);

extern struct Moc_PackedString callbackQmlBridge_SendToGo(void* p0, struct Moc_PackedString p1);

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
