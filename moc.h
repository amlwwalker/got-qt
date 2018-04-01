

#pragma once

#ifndef GO_MOC_H
#define GO_MOC_H

#include <stdint.h>

#ifdef __cplusplus
class ApplicationUI;
void ApplicationUI_ApplicationUI_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; };
struct Moc_PackedList { void* data; long long len; };
struct Moc_PackedString ApplicationUI_SwapThemePalette(void* ptr);
struct Moc_PackedString ApplicationUI_DefaultThemePalette(void* ptr);
struct Moc_PackedString ApplicationUI_PrimaryPalette(void* ptr, int paletteIndex);
struct Moc_PackedString ApplicationUI_AccentPalette(void* ptr, int paletteIndex);
struct Moc_PackedString ApplicationUI_DefaultPrimaryPalette(void* ptr);
struct Moc_PackedString ApplicationUI_DefaultAccentPalette(void* ptr);
int ApplicationUI_ApplicationUI_QRegisterMetaType();
int ApplicationUI_ApplicationUI_QRegisterMetaType2(char* typeName);
int ApplicationUI_ApplicationUI_QmlRegisterType();
int ApplicationUI_ApplicationUI_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* ApplicationUI___dynamicPropertyNames_atList(void* ptr, int i);
void ApplicationUI___dynamicPropertyNames_setList(void* ptr, void* i);
void* ApplicationUI___dynamicPropertyNames_newList(void* ptr);
void* ApplicationUI___findChildren_atList2(void* ptr, int i);
void ApplicationUI___findChildren_setList2(void* ptr, void* i);
void* ApplicationUI___findChildren_newList2(void* ptr);
void* ApplicationUI___findChildren_atList3(void* ptr, int i);
void ApplicationUI___findChildren_setList3(void* ptr, void* i);
void* ApplicationUI___findChildren_newList3(void* ptr);
void* ApplicationUI___findChildren_atList(void* ptr, int i);
void ApplicationUI___findChildren_setList(void* ptr, void* i);
void* ApplicationUI___findChildren_newList(void* ptr);
void* ApplicationUI___children_atList(void* ptr, int i);
void ApplicationUI___children_setList(void* ptr, void* i);
void* ApplicationUI___children_newList(void* ptr);
void* ApplicationUI_NewApplicationUI(void* parent);
void ApplicationUI_DestroyApplicationUI(void* ptr);
void ApplicationUI_DestroyApplicationUIDefault(void* ptr);
char ApplicationUI_EventDefault(void* ptr, void* e);
char ApplicationUI_EventFilterDefault(void* ptr, void* watched, void* event);
void ApplicationUI_ChildEventDefault(void* ptr, void* event);
void ApplicationUI_ConnectNotifyDefault(void* ptr, void* sign);
void ApplicationUI_CustomEventDefault(void* ptr, void* event);
void ApplicationUI_DeleteLaterDefault(void* ptr);
void ApplicationUI_DisconnectNotifyDefault(void* ptr, void* sign);
void ApplicationUI_TimerEventDefault(void* ptr, void* event);
;

#ifdef __cplusplus
}
#endif

#endif