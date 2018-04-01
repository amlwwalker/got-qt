

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QCamera>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLayout>
#include <QList>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QRadioData>
#include <QSignalSpy>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QWidget>
#include <QWindow>


class ApplicationUI: public QObject
{
Q_OBJECT
public:
	ApplicationUI(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");ApplicationUI_ApplicationUI_QRegisterMetaType();ApplicationUI_ApplicationUI_QRegisterMetaTypes();callbackApplicationUI_Constructor(this);};
	 ~ApplicationUI() { callbackApplicationUI_DestroyApplicationUI(this); };
	bool event(QEvent * e) { return callbackApplicationUI_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackApplicationUI_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackApplicationUI_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackApplicationUI_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackApplicationUI_CustomEvent(this, event); };
	void deleteLater() { callbackApplicationUI_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackApplicationUI_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackApplicationUI_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackApplicationUI_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackApplicationUI_TimerEvent(this, event); };
	
signals:
public slots:
	QStringList swapThemePalette() { return ({ Moc_PackedString tempVal = callbackApplicationUI_SwapThemePalette(this); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QStringList defaultThemePalette() { return ({ Moc_PackedString tempVal = callbackApplicationUI_DefaultThemePalette(this); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QStringList primaryPalette(qint32 paletteIndex) { return ({ Moc_PackedString tempVal = callbackApplicationUI_PrimaryPalette(this, paletteIndex); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QStringList accentPalette(qint32 paletteIndex) { return ({ Moc_PackedString tempVal = callbackApplicationUI_AccentPalette(this, paletteIndex); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QStringList defaultPrimaryPalette() { return ({ Moc_PackedString tempVal = callbackApplicationUI_DefaultPrimaryPalette(this); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QStringList defaultAccentPalette() { return ({ Moc_PackedString tempVal = callbackApplicationUI_DefaultAccentPalette(this); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
private:
};

Q_DECLARE_METATYPE(ApplicationUI*)


void ApplicationUI_ApplicationUI_QRegisterMetaTypes() {
}

struct Moc_PackedString ApplicationUI_SwapThemePalette(void* ptr)
{
	QStringList returnArg;
	QMetaObject::invokeMethod(static_cast<ApplicationUI*>(ptr), "swapThemePalette", Q_RETURN_ARG(QStringList, returnArg));
	return ({ QByteArray t8e5b69 = returnArg.join("|").toUtf8(); Moc_PackedString { const_cast<char*>(t8e5b69.prepend("WHITESPACE").constData()+10), t8e5b69.size()-10 }; });
}

struct Moc_PackedString ApplicationUI_DefaultThemePalette(void* ptr)
{
	QStringList returnArg;
	QMetaObject::invokeMethod(static_cast<ApplicationUI*>(ptr), "defaultThemePalette", Q_RETURN_ARG(QStringList, returnArg));
	return ({ QByteArray t8e5b69 = returnArg.join("|").toUtf8(); Moc_PackedString { const_cast<char*>(t8e5b69.prepend("WHITESPACE").constData()+10), t8e5b69.size()-10 }; });
}

struct Moc_PackedString ApplicationUI_PrimaryPalette(void* ptr, int paletteIndex)
{
	QStringList returnArg;
	QMetaObject::invokeMethod(static_cast<ApplicationUI*>(ptr), "primaryPalette", Q_RETURN_ARG(QStringList, returnArg), Q_ARG(qint32, paletteIndex));
	return ({ QByteArray t8e5b69 = returnArg.join("|").toUtf8(); Moc_PackedString { const_cast<char*>(t8e5b69.prepend("WHITESPACE").constData()+10), t8e5b69.size()-10 }; });
}

struct Moc_PackedString ApplicationUI_AccentPalette(void* ptr, int paletteIndex)
{
	QStringList returnArg;
	QMetaObject::invokeMethod(static_cast<ApplicationUI*>(ptr), "accentPalette", Q_RETURN_ARG(QStringList, returnArg), Q_ARG(qint32, paletteIndex));
	return ({ QByteArray t8e5b69 = returnArg.join("|").toUtf8(); Moc_PackedString { const_cast<char*>(t8e5b69.prepend("WHITESPACE").constData()+10), t8e5b69.size()-10 }; });
}

struct Moc_PackedString ApplicationUI_DefaultPrimaryPalette(void* ptr)
{
	QStringList returnArg;
	QMetaObject::invokeMethod(static_cast<ApplicationUI*>(ptr), "defaultPrimaryPalette", Q_RETURN_ARG(QStringList, returnArg));
	return ({ QByteArray t8e5b69 = returnArg.join("|").toUtf8(); Moc_PackedString { const_cast<char*>(t8e5b69.prepend("WHITESPACE").constData()+10), t8e5b69.size()-10 }; });
}

struct Moc_PackedString ApplicationUI_DefaultAccentPalette(void* ptr)
{
	QStringList returnArg;
	QMetaObject::invokeMethod(static_cast<ApplicationUI*>(ptr), "defaultAccentPalette", Q_RETURN_ARG(QStringList, returnArg));
	return ({ QByteArray t8e5b69 = returnArg.join("|").toUtf8(); Moc_PackedString { const_cast<char*>(t8e5b69.prepend("WHITESPACE").constData()+10), t8e5b69.size()-10 }; });
}

int ApplicationUI_ApplicationUI_QRegisterMetaType()
{
	return qRegisterMetaType<ApplicationUI*>();
}

int ApplicationUI_ApplicationUI_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<ApplicationUI*>(const_cast<const char*>(typeName));
}

int ApplicationUI_ApplicationUI_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ApplicationUI>();
#else
	return 0;
#endif
}

int ApplicationUI_ApplicationUI_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ApplicationUI>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* ApplicationUI___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void ApplicationUI___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* ApplicationUI___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* ApplicationUI___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void ApplicationUI___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ApplicationUI___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* ApplicationUI___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void ApplicationUI___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ApplicationUI___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* ApplicationUI___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void ApplicationUI___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ApplicationUI___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* ApplicationUI___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void ApplicationUI___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ApplicationUI___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

void* ApplicationUI_NewApplicationUI(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new ApplicationUI(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new ApplicationUI(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new ApplicationUI(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new ApplicationUI(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new ApplicationUI(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new ApplicationUI(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new ApplicationUI(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new ApplicationUI(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new ApplicationUI(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new ApplicationUI(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new ApplicationUI(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new ApplicationUI(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new ApplicationUI(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new ApplicationUI(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new ApplicationUI(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new ApplicationUI(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new ApplicationUI(static_cast<QWindow*>(parent));
	} else {
		return new ApplicationUI(static_cast<QObject*>(parent));
	}
}

void ApplicationUI_DestroyApplicationUI(void* ptr)
{
	static_cast<ApplicationUI*>(ptr)->~ApplicationUI();
}

void ApplicationUI_DestroyApplicationUIDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char ApplicationUI_EventDefault(void* ptr, void* e)
{
	return static_cast<ApplicationUI*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char ApplicationUI_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<ApplicationUI*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void ApplicationUI_ChildEventDefault(void* ptr, void* event)
{
	static_cast<ApplicationUI*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void ApplicationUI_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ApplicationUI*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void ApplicationUI_CustomEventDefault(void* ptr, void* event)
{
	static_cast<ApplicationUI*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void ApplicationUI_DeleteLaterDefault(void* ptr)
{
	static_cast<ApplicationUI*>(ptr)->QObject::deleteLater();
}

void ApplicationUI_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ApplicationUI*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void ApplicationUI_TimerEventDefault(void* ptr, void* event)
{
	static_cast<ApplicationUI*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



#include "moc_moc.h"
