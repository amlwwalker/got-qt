

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


class QmlBridge: public QObject
{
Q_OBJECT
public:
	QmlBridge(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QmlBridge_QmlBridge_QRegisterMetaType();QmlBridge_QmlBridge_QRegisterMetaTypes();callbackQmlBridge_Constructor(this);};
	void Signal_UpdateLoader(QString p) { QByteArray t516b97 = p.toUtf8(); Moc_PackedString pPacked = { const_cast<char*>(t516b97.prepend("WHITESPACE").constData()+10), t516b97.size()-10 };callbackQmlBridge_UpdateLoader(this, pPacked); };
	void Signal_UpdateSettings(QString author, QString mode, QString date, QString host, QString version, QString port, bool hotload) { QByteArray tf64cd8 = author.toUtf8(); Moc_PackedString authorPacked = { const_cast<char*>(tf64cd8.prepend("WHITESPACE").constData()+10), tf64cd8.size()-10 };QByteArray te78fe7 = mode.toUtf8(); Moc_PackedString modePacked = { const_cast<char*>(te78fe7.prepend("WHITESPACE").constData()+10), te78fe7.size()-10 };QByteArray te927d0 = date.toUtf8(); Moc_PackedString datePacked = { const_cast<char*>(te927d0.prepend("WHITESPACE").constData()+10), te927d0.size()-10 };QByteArray t86dd1c = host.toUtf8(); Moc_PackedString hostPacked = { const_cast<char*>(t86dd1c.prepend("WHITESPACE").constData()+10), t86dd1c.size()-10 };QByteArray tc69227 = version.toUtf8(); Moc_PackedString versionPacked = { const_cast<char*>(tc69227.prepend("WHITESPACE").constData()+10), tc69227.size()-10 };QByteArray t062d8a = port.toUtf8(); Moc_PackedString portPacked = { const_cast<char*>(t062d8a.prepend("WHITESPACE").constData()+10), t062d8a.size()-10 };callbackQmlBridge_UpdateSettings(this, authorPacked, modePacked, datePacked, hostPacked, versionPacked, portPacked, hotload); };
	void Signal_SendTime(QString data) { QByteArray ta17c9a = data.toUtf8(); Moc_PackedString dataPacked = { const_cast<char*>(ta17c9a.prepend("WHITESPACE").constData()+10), ta17c9a.size()-10 };callbackQmlBridge_SendTime(this, dataPacked); };
	 ~QmlBridge() { callbackQmlBridge_DestroyQmlBridge(this); };
	bool event(QEvent * e) { return callbackQmlBridge_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQmlBridge_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQmlBridge_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQmlBridge_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQmlBridge_CustomEvent(this, event); };
	void deleteLater() { callbackQmlBridge_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQmlBridge_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQmlBridge_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQmlBridge_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQmlBridge_TimerEvent(this, event); };
	
signals:
	void updateLoader(QString p);
	void updateSettings(QString author, QString mode, QString date, QString host, QString version, QString port, bool hotload);
	void sendTime(QString data);
public slots:
	QString calculator(QString number1, QString number2) { QByteArray t97e313 = number1.toUtf8(); Moc_PackedString number1Packed = { const_cast<char*>(t97e313.prepend("WHITESPACE").constData()+10), t97e313.size()-10 };QByteArray t0bc97d = number2.toUtf8(); Moc_PackedString number2Packed = { const_cast<char*>(t0bc97d.prepend("WHITESPACE").constData()+10), t0bc97d.size()-10 };return ({ Moc_PackedString tempVal = callbackQmlBridge_Calculator(this, number1Packed, number2Packed); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
private:
};

Q_DECLARE_METATYPE(QmlBridge*)


void QmlBridge_QmlBridge_QRegisterMetaTypes() {
}

void QmlBridge_ConnectUpdateLoader(void* ptr)
{
	QObject::connect(static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString)>(&QmlBridge::updateLoader), static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString)>(&QmlBridge::Signal_UpdateLoader));
}

void QmlBridge_DisconnectUpdateLoader(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString)>(&QmlBridge::updateLoader), static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString)>(&QmlBridge::Signal_UpdateLoader));
}

void QmlBridge_UpdateLoader(void* ptr, struct Moc_PackedString p)
{
	static_cast<QmlBridge*>(ptr)->updateLoader(QString::fromUtf8(p.data, p.len));
}

void QmlBridge_ConnectUpdateSettings(void* ptr)
{
	QObject::connect(static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString, QString, QString, QString, QString, QString, bool)>(&QmlBridge::updateSettings), static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString, QString, QString, QString, QString, QString, bool)>(&QmlBridge::Signal_UpdateSettings));
}

void QmlBridge_DisconnectUpdateSettings(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString, QString, QString, QString, QString, QString, bool)>(&QmlBridge::updateSettings), static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString, QString, QString, QString, QString, QString, bool)>(&QmlBridge::Signal_UpdateSettings));
}

void QmlBridge_UpdateSettings(void* ptr, struct Moc_PackedString author, struct Moc_PackedString mode, struct Moc_PackedString date, struct Moc_PackedString host, struct Moc_PackedString version, struct Moc_PackedString port, char hotload)
{
	static_cast<QmlBridge*>(ptr)->updateSettings(QString::fromUtf8(author.data, author.len), QString::fromUtf8(mode.data, mode.len), QString::fromUtf8(date.data, date.len), QString::fromUtf8(host.data, host.len), QString::fromUtf8(version.data, version.len), QString::fromUtf8(port.data, port.len), hotload != 0);
}

void QmlBridge_ConnectSendTime(void* ptr)
{
	QObject::connect(static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString)>(&QmlBridge::sendTime), static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString)>(&QmlBridge::Signal_SendTime));
}

void QmlBridge_DisconnectSendTime(void* ptr)
{
	QObject::disconnect(static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString)>(&QmlBridge::sendTime), static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString)>(&QmlBridge::Signal_SendTime));
}

void QmlBridge_SendTime(void* ptr, struct Moc_PackedString data)
{
	static_cast<QmlBridge*>(ptr)->sendTime(QString::fromUtf8(data.data, data.len));
}

struct Moc_PackedString QmlBridge_Calculator(void* ptr, struct Moc_PackedString number1, struct Moc_PackedString number2)
{
	QString returnArg;
	QMetaObject::invokeMethod(static_cast<QmlBridge*>(ptr), "calculator", Q_RETURN_ARG(QString, returnArg), Q_ARG(QString, QString::fromUtf8(number1.data, number1.len)), Q_ARG(QString, QString::fromUtf8(number2.data, number2.len)));
	return ({ QByteArray t8e5b69 = returnArg.toUtf8(); Moc_PackedString { const_cast<char*>(t8e5b69.prepend("WHITESPACE").constData()+10), t8e5b69.size()-10 }; });
}

int QmlBridge_QmlBridge_QRegisterMetaType()
{
	return qRegisterMetaType<QmlBridge*>();
}

int QmlBridge_QmlBridge_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QmlBridge*>(const_cast<const char*>(typeName));
}

int QmlBridge_QmlBridge_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QmlBridge>();
#else
	return 0;
#endif
}

int QmlBridge_QmlBridge_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QmlBridge>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QmlBridge___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QmlBridge___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QmlBridge___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QmlBridge___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QmlBridge___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QmlBridge___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QmlBridge___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QmlBridge___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QmlBridge___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QmlBridge___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QmlBridge___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

void* QmlBridge_NewQmlBridge(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QmlBridge(static_cast<QWindow*>(parent));
	} else {
		return new QmlBridge(static_cast<QObject*>(parent));
	}
}

void QmlBridge_DestroyQmlBridge(void* ptr)
{
	static_cast<QmlBridge*>(ptr)->~QmlBridge();
}

void QmlBridge_DestroyQmlBridgeDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char QmlBridge_EventDefault(void* ptr, void* e)
{
	return static_cast<QmlBridge*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QmlBridge_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QmlBridge*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QmlBridge_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QmlBridge*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QmlBridge_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QmlBridge*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QmlBridge_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QmlBridge*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QmlBridge_DeleteLaterDefault(void* ptr)
{
	static_cast<QmlBridge*>(ptr)->QObject::deleteLater();
}

void QmlBridge_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QmlBridge*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QmlBridge_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QmlBridge*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



#include "moc_moc.h"
