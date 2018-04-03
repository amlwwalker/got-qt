

#pragma once

#ifndef GO_MOC_H
#define GO_MOC_H

#include <stdint.h>

#ifdef __cplusplus
class QmlBridge;
void QmlBridge_QmlBridge_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; };
struct Moc_PackedList { void* data; long long len; };
void QmlBridge_ConnectUpdateLoader(void* ptr);
void QmlBridge_DisconnectUpdateLoader(void* ptr);
void QmlBridge_UpdateLoader(void* ptr, struct Moc_PackedString p);
void QmlBridge_ConnectUpdateSettings(void* ptr);
void QmlBridge_DisconnectUpdateSettings(void* ptr);
void QmlBridge_UpdateSettings(void* ptr, struct Moc_PackedString author, struct Moc_PackedString mode, struct Moc_PackedString date, struct Moc_PackedString host, struct Moc_PackedString version, struct Moc_PackedString port, char hotload);
void QmlBridge_ConnectSendTime(void* ptr);
void QmlBridge_DisconnectSendTime(void* ptr);
void QmlBridge_SendTime(void* ptr, struct Moc_PackedString data);
struct Moc_PackedString QmlBridge_Calculator(void* ptr, struct Moc_PackedString number1, struct Moc_PackedString number2);
int QmlBridge_QmlBridge_QRegisterMetaType();
int QmlBridge_QmlBridge_QRegisterMetaType2(char* typeName);
int QmlBridge_QmlBridge_QmlRegisterType();
int QmlBridge_QmlBridge_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* QmlBridge___dynamicPropertyNames_atList(void* ptr, int i);
void QmlBridge___dynamicPropertyNames_setList(void* ptr, void* i);
void* QmlBridge___dynamicPropertyNames_newList(void* ptr);
void* QmlBridge___findChildren_atList2(void* ptr, int i);
void QmlBridge___findChildren_setList2(void* ptr, void* i);
void* QmlBridge___findChildren_newList2(void* ptr);
void* QmlBridge___findChildren_atList3(void* ptr, int i);
void QmlBridge___findChildren_setList3(void* ptr, void* i);
void* QmlBridge___findChildren_newList3(void* ptr);
void* QmlBridge___findChildren_atList(void* ptr, int i);
void QmlBridge___findChildren_setList(void* ptr, void* i);
void* QmlBridge___findChildren_newList(void* ptr);
void* QmlBridge___children_atList(void* ptr, int i);
void QmlBridge___children_setList(void* ptr, void* i);
void* QmlBridge___children_newList(void* ptr);
void* QmlBridge_NewQmlBridge(void* parent);
void QmlBridge_DestroyQmlBridge(void* ptr);
void QmlBridge_DestroyQmlBridgeDefault(void* ptr);
char QmlBridge_EventDefault(void* ptr, void* e);
char QmlBridge_EventFilterDefault(void* ptr, void* watched, void* event);
void QmlBridge_ChildEventDefault(void* ptr, void* event);
void QmlBridge_ConnectNotifyDefault(void* ptr, void* sign);
void QmlBridge_CustomEventDefault(void* ptr, void* event);
void QmlBridge_DeleteLaterDefault(void* ptr);
void QmlBridge_DisconnectNotifyDefault(void* ptr, void* sign);
void QmlBridge_TimerEventDefault(void* ptr, void* event);
;

#ifdef __cplusplus
}
#endif

#endif