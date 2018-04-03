

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QAbstractItemModel>
#include <QAbstractListModel>
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
#include <QHash>
#include <QLayout>
#include <QList>
#include <QMap>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMimeData>
#include <QModelIndex>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QPersistentModelIndex>
#include <QQuickItem>
#include <QRadioData>
#include <QSignalSpy>
#include <QSize>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>
#include <QVector>
#include <QWidget>
#include <QWindow>


typedef QHash<qint32, QByteArray> typed01680;
class Person: public QObject
{
Q_OBJECT
Q_PROPERTY(QString firstName READ firstName WRITE setFirstName NOTIFY firstNameChanged)
Q_PROPERTY(QString lastName READ lastName WRITE setLastName NOTIFY lastNameChanged)
Q_PROPERTY(QString email READ email WRITE setEmail NOTIFY emailChanged)
public:
	Person(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");Person_Person_QRegisterMetaType();Person_Person_QRegisterMetaTypes();callbackPerson_Constructor(this);};
	QString firstName() { return ({ Moc_PackedString tempVal = callbackPerson_FirstName(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setFirstName(QString firstName) { QByteArray te57915 = firstName.toUtf8(); Moc_PackedString firstNamePacked = { const_cast<char*>(te57915.prepend("WHITESPACE").constData()+10), te57915.size()-10 };callbackPerson_SetFirstName(this, firstNamePacked); };
	void Signal_FirstNameChanged(QString firstName) { QByteArray te57915 = firstName.toUtf8(); Moc_PackedString firstNamePacked = { const_cast<char*>(te57915.prepend("WHITESPACE").constData()+10), te57915.size()-10 };callbackPerson_FirstNameChanged(this, firstNamePacked); };
	QString lastName() { return ({ Moc_PackedString tempVal = callbackPerson_LastName(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setLastName(QString lastName) { QByteArray t9b831e = lastName.toUtf8(); Moc_PackedString lastNamePacked = { const_cast<char*>(t9b831e.prepend("WHITESPACE").constData()+10), t9b831e.size()-10 };callbackPerson_SetLastName(this, lastNamePacked); };
	void Signal_LastNameChanged(QString lastName) { QByteArray t9b831e = lastName.toUtf8(); Moc_PackedString lastNamePacked = { const_cast<char*>(t9b831e.prepend("WHITESPACE").constData()+10), t9b831e.size()-10 };callbackPerson_LastNameChanged(this, lastNamePacked); };
	QString email() { return ({ Moc_PackedString tempVal = callbackPerson_Email(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setEmail(QString email) { QByteArray ta88b7d = email.toUtf8(); Moc_PackedString emailPacked = { const_cast<char*>(ta88b7d.prepend("WHITESPACE").constData()+10), ta88b7d.size()-10 };callbackPerson_SetEmail(this, emailPacked); };
	void Signal_EmailChanged(QString email) { QByteArray ta88b7d = email.toUtf8(); Moc_PackedString emailPacked = { const_cast<char*>(ta88b7d.prepend("WHITESPACE").constData()+10), ta88b7d.size()-10 };callbackPerson_EmailChanged(this, emailPacked); };
	 ~Person() { callbackPerson_DestroyPerson(this); };
	bool event(QEvent * e) { return callbackPerson_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackPerson_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackPerson_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackPerson_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackPerson_CustomEvent(this, event); };
	void deleteLater() { callbackPerson_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackPerson_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackPerson_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackPerson_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackPerson_TimerEvent(this, event); };
	
	QString firstNameDefault() { return _firstName; };
	void setFirstNameDefault(QString p) { if (p != _firstName) { _firstName = p; firstNameChanged(_firstName); } };
	QString lastNameDefault() { return _lastName; };
	void setLastNameDefault(QString p) { if (p != _lastName) { _lastName = p; lastNameChanged(_lastName); } };
	QString emailDefault() { return _email; };
	void setEmailDefault(QString p) { if (p != _email) { _email = p; emailChanged(_email); } };
signals:
	void firstNameChanged(QString firstName);
	void lastNameChanged(QString lastName);
	void emailChanged(QString email);
public slots:
private:
	QString _firstName;
	QString _lastName;
	QString _email;
};

Q_DECLARE_METATYPE(Person*)


void Person_Person_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class PersonModel: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(typed01680 roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<Person*> people READ people WRITE setPeople NOTIFY peopleChanged)
public:
	PersonModel(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");PersonModel_PersonModel_QRegisterMetaType();PersonModel_PersonModel_QRegisterMetaTypes();callbackPersonModel_Constructor(this);};
	typed01680 roles() { return *static_cast<QHash<qint32, QByteArray>*>(callbackPersonModel_Roles(this)); };
	void setRoles(typed01680 roles) { callbackPersonModel_SetRoles(this, ({ QHash<qint32, QByteArray>* tmpValue = new QHash<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_RolesChanged(typed01680 roles) { callbackPersonModel_RolesChanged(this, ({ QHash<qint32, QByteArray>* tmpValue = new QHash<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QList<Person*> people() { return *static_cast<QList<Person*>*>(callbackPersonModel_People(this)); };
	void setPeople(QList<Person*> people) { callbackPersonModel_SetPeople(this, ({ QList<Person*>* tmpValue = new QList<Person*>(people); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_PeopleChanged(QList<Person*> people) { callbackPersonModel_PeopleChanged(this, ({ QList<Person*>* tmpValue = new QList<Person*>(people); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackPersonModel_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackPersonModel_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackPersonModel_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackPersonModel_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackPersonModel_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackPersonModel_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackPersonModel_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackPersonModel_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackPersonModel_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackPersonModel_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackPersonModel_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackPersonModel_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackPersonModel_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = const_cast<QMap<int, QVariant>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	bool submit() { return callbackPersonModel_Submit(this) != 0; };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackPersonModel_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackPersonModel_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackPersonModel_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackPersonModel_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackPersonModel_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackPersonModel_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackPersonModel_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = const_cast<QVector<int>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackPersonModel_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackPersonModel_HeaderDataChanged(this, orientation, first, last); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackPersonModel_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackPersonModel_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_ModelAboutToBeReset() { callbackPersonModel_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackPersonModel_ModelReset(this); };
	void resetInternalData() { callbackPersonModel_ResetInternalData(this); };
	void revert() { callbackPersonModel_Revert(this); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackPersonModel_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackPersonModel_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackPersonModel_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackPersonModel_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackPersonModel_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackPersonModel_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void sort(int column, Qt::SortOrder order) { callbackPersonModel_Sort(this, column, order); };
	QHash<int, QByteArray> roleNames() const { return *static_cast<QHash<int, QByteArray>*>(callbackPersonModel_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return *static_cast<QMap<int, QVariant>*>(callbackPersonModel_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackPersonModel_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackPersonModel_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackPersonModel_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return *static_cast<QList<QModelIndex>*>(callbackPersonModel_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackPersonModel_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackPersonModel_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackPersonModel_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackPersonModel_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackPersonModel_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackPersonModel_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackPersonModel_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackPersonModel_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	bool hasChildren(const QModelIndex & parent) const { return callbackPersonModel_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackPersonModel_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	int rowCount(const QModelIndex & parent) const { return callbackPersonModel_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool event(QEvent * e) { return callbackPersonModel_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackPersonModel_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackPersonModel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackPersonModel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackPersonModel_CustomEvent(this, event); };
	void deleteLater() { callbackPersonModel_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackPersonModel_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackPersonModel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackPersonModel_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackPersonModel_TimerEvent(this, event); };
	
	typed01680 rolesDefault() { return _roles; };
	void setRolesDefault(typed01680 p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<Person*> peopleDefault() { return _people; };
	void setPeopleDefault(QList<Person*> p) { if (p != _people) { _people = p; peopleChanged(_people); } };
signals:
	void rolesChanged(typed01680 roles);
	void peopleChanged(QList<Person*> people);
public slots:
	void addPerson(Person* v0) { callbackPersonModel_AddPerson(this, v0); };
	void editPerson(qint32 row, QString firstName, QString lastName, QString email) { QByteArray te57915 = firstName.toUtf8(); Moc_PackedString firstNamePacked = { const_cast<char*>(te57915.prepend("WHITESPACE").constData()+10), te57915.size()-10 };QByteArray t9b831e = lastName.toUtf8(); Moc_PackedString lastNamePacked = { const_cast<char*>(t9b831e.prepend("WHITESPACE").constData()+10), t9b831e.size()-10 };QByteArray ta88b7d = email.toUtf8(); Moc_PackedString emailPacked = { const_cast<char*>(ta88b7d.prepend("WHITESPACE").constData()+10), ta88b7d.size()-10 };callbackPersonModel_EditPerson(this, row, firstNamePacked, lastNamePacked, emailPacked); };
	void removePerson(qint32 row) { callbackPersonModel_RemovePerson(this, row); };
private:
	typed01680 _roles;
	QList<Person*> _people;
};

Q_DECLARE_METATYPE(PersonModel*)


void PersonModel_PersonModel_QRegisterMetaTypes() {
	qRegisterMetaType<typed01680>("typed01680");
}

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

struct Moc_PackedString Person_FirstName(void* ptr)
{
	return ({ QByteArray td2d56b = static_cast<Person*>(ptr)->firstName().toUtf8(); Moc_PackedString { const_cast<char*>(td2d56b.prepend("WHITESPACE").constData()+10), td2d56b.size()-10 }; });
}

struct Moc_PackedString Person_FirstNameDefault(void* ptr)
{
	return ({ QByteArray te9c1fc = static_cast<Person*>(ptr)->firstNameDefault().toUtf8(); Moc_PackedString { const_cast<char*>(te9c1fc.prepend("WHITESPACE").constData()+10), te9c1fc.size()-10 }; });
}

void Person_SetFirstName(void* ptr, struct Moc_PackedString firstName)
{
	static_cast<Person*>(ptr)->setFirstName(QString::fromUtf8(firstName.data, firstName.len));
}

void Person_SetFirstNameDefault(void* ptr, struct Moc_PackedString firstName)
{
	static_cast<Person*>(ptr)->setFirstNameDefault(QString::fromUtf8(firstName.data, firstName.len));
}

void Person_ConnectFirstNameChanged(void* ptr)
{
	QObject::connect(static_cast<Person*>(ptr), static_cast<void (Person::*)(QString)>(&Person::firstNameChanged), static_cast<Person*>(ptr), static_cast<void (Person::*)(QString)>(&Person::Signal_FirstNameChanged));
}

void Person_DisconnectFirstNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<Person*>(ptr), static_cast<void (Person::*)(QString)>(&Person::firstNameChanged), static_cast<Person*>(ptr), static_cast<void (Person::*)(QString)>(&Person::Signal_FirstNameChanged));
}

void Person_FirstNameChanged(void* ptr, struct Moc_PackedString firstName)
{
	static_cast<Person*>(ptr)->firstNameChanged(QString::fromUtf8(firstName.data, firstName.len));
}

struct Moc_PackedString Person_LastName(void* ptr)
{
	return ({ QByteArray t2d35c4 = static_cast<Person*>(ptr)->lastName().toUtf8(); Moc_PackedString { const_cast<char*>(t2d35c4.prepend("WHITESPACE").constData()+10), t2d35c4.size()-10 }; });
}

struct Moc_PackedString Person_LastNameDefault(void* ptr)
{
	return ({ QByteArray tdd9162 = static_cast<Person*>(ptr)->lastNameDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tdd9162.prepend("WHITESPACE").constData()+10), tdd9162.size()-10 }; });
}

void Person_SetLastName(void* ptr, struct Moc_PackedString lastName)
{
	static_cast<Person*>(ptr)->setLastName(QString::fromUtf8(lastName.data, lastName.len));
}

void Person_SetLastNameDefault(void* ptr, struct Moc_PackedString lastName)
{
	static_cast<Person*>(ptr)->setLastNameDefault(QString::fromUtf8(lastName.data, lastName.len));
}

void Person_ConnectLastNameChanged(void* ptr)
{
	QObject::connect(static_cast<Person*>(ptr), static_cast<void (Person::*)(QString)>(&Person::lastNameChanged), static_cast<Person*>(ptr), static_cast<void (Person::*)(QString)>(&Person::Signal_LastNameChanged));
}

void Person_DisconnectLastNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<Person*>(ptr), static_cast<void (Person::*)(QString)>(&Person::lastNameChanged), static_cast<Person*>(ptr), static_cast<void (Person::*)(QString)>(&Person::Signal_LastNameChanged));
}

void Person_LastNameChanged(void* ptr, struct Moc_PackedString lastName)
{
	static_cast<Person*>(ptr)->lastNameChanged(QString::fromUtf8(lastName.data, lastName.len));
}

struct Moc_PackedString Person_Email(void* ptr)
{
	return ({ QByteArray taba31d = static_cast<Person*>(ptr)->email().toUtf8(); Moc_PackedString { const_cast<char*>(taba31d.prepend("WHITESPACE").constData()+10), taba31d.size()-10 }; });
}

struct Moc_PackedString Person_EmailDefault(void* ptr)
{
	return ({ QByteArray tc63646 = static_cast<Person*>(ptr)->emailDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tc63646.prepend("WHITESPACE").constData()+10), tc63646.size()-10 }; });
}

void Person_SetEmail(void* ptr, struct Moc_PackedString email)
{
	static_cast<Person*>(ptr)->setEmail(QString::fromUtf8(email.data, email.len));
}

void Person_SetEmailDefault(void* ptr, struct Moc_PackedString email)
{
	static_cast<Person*>(ptr)->setEmailDefault(QString::fromUtf8(email.data, email.len));
}

void Person_ConnectEmailChanged(void* ptr)
{
	QObject::connect(static_cast<Person*>(ptr), static_cast<void (Person::*)(QString)>(&Person::emailChanged), static_cast<Person*>(ptr), static_cast<void (Person::*)(QString)>(&Person::Signal_EmailChanged));
}

void Person_DisconnectEmailChanged(void* ptr)
{
	QObject::disconnect(static_cast<Person*>(ptr), static_cast<void (Person::*)(QString)>(&Person::emailChanged), static_cast<Person*>(ptr), static_cast<void (Person::*)(QString)>(&Person::Signal_EmailChanged));
}

void Person_EmailChanged(void* ptr, struct Moc_PackedString email)
{
	static_cast<Person*>(ptr)->emailChanged(QString::fromUtf8(email.data, email.len));
}

int Person_Person_QRegisterMetaType()
{
	return qRegisterMetaType<Person*>();
}

int Person_Person_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<Person*>(const_cast<const char*>(typeName));
}

int Person_Person_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Person>();
#else
	return 0;
#endif
}

int Person_Person_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Person>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* Person___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void Person___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* Person___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* Person___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void Person___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Person___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* Person___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void Person___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Person___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* Person___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void Person___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Person___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* Person___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void Person___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Person___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

void* Person_NewPerson(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new Person(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new Person(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new Person(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new Person(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new Person(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new Person(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new Person(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new Person(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new Person(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new Person(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new Person(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new Person(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new Person(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new Person(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new Person(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new Person(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new Person(static_cast<QWindow*>(parent));
	} else {
		return new Person(static_cast<QObject*>(parent));
	}
}

void Person_DestroyPerson(void* ptr)
{
	static_cast<Person*>(ptr)->~Person();
}

void Person_DestroyPersonDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char Person_EventDefault(void* ptr, void* e)
{
	return static_cast<Person*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char Person_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<Person*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void Person_ChildEventDefault(void* ptr, void* event)
{
	static_cast<Person*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void Person_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Person*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void Person_CustomEventDefault(void* ptr, void* event)
{
	static_cast<Person*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void Person_DeleteLaterDefault(void* ptr)
{
	static_cast<Person*>(ptr)->QObject::deleteLater();
}

void Person_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Person*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void Person_TimerEventDefault(void* ptr, void* event)
{
	static_cast<Person*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



void PersonModel_AddPerson(void* ptr, void* v0)
{
	QMetaObject::invokeMethod(static_cast<PersonModel*>(ptr), "addPerson", Q_ARG(Person*, static_cast<Person*>(v0)));
}

void PersonModel_EditPerson(void* ptr, int row, struct Moc_PackedString firstName, struct Moc_PackedString lastName, struct Moc_PackedString email)
{
	QMetaObject::invokeMethod(static_cast<PersonModel*>(ptr), "editPerson", Q_ARG(qint32, row), Q_ARG(QString, QString::fromUtf8(firstName.data, firstName.len)), Q_ARG(QString, QString::fromUtf8(lastName.data, lastName.len)), Q_ARG(QString, QString::fromUtf8(email.data, email.len)));
}

void PersonModel_RemovePerson(void* ptr, int row)
{
	QMetaObject::invokeMethod(static_cast<PersonModel*>(ptr), "removePerson", Q_ARG(qint32, row));
}

struct Moc_PackedList PersonModel_Roles(void* ptr)
{
	return ({ QHash<qint32, QByteArray>* tmpValue = new QHash<qint32, QByteArray>(static_cast<PersonModel*>(ptr)->roles()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList PersonModel_RolesDefault(void* ptr)
{
	return ({ QHash<qint32, QByteArray>* tmpValue = new QHash<qint32, QByteArray>(static_cast<PersonModel*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void PersonModel_SetRoles(void* ptr, void* roles)
{
	static_cast<PersonModel*>(ptr)->setRoles(*static_cast<QHash<qint32, QByteArray>*>(roles));
}

void PersonModel_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<PersonModel*>(ptr)->setRolesDefault(*static_cast<QHash<qint32, QByteArray>*>(roles));
}

void PersonModel_ConnectRolesChanged(void* ptr)
{
	QObject::connect(static_cast<PersonModel*>(ptr), static_cast<void (PersonModel::*)(QHash<qint32, QByteArray>)>(&PersonModel::rolesChanged), static_cast<PersonModel*>(ptr), static_cast<void (PersonModel::*)(QHash<qint32, QByteArray>)>(&PersonModel::Signal_RolesChanged));
}

void PersonModel_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<PersonModel*>(ptr), static_cast<void (PersonModel::*)(QHash<qint32, QByteArray>)>(&PersonModel::rolesChanged), static_cast<PersonModel*>(ptr), static_cast<void (PersonModel::*)(QHash<qint32, QByteArray>)>(&PersonModel::Signal_RolesChanged));
}

void PersonModel_RolesChanged(void* ptr, void* roles)
{
	static_cast<PersonModel*>(ptr)->rolesChanged(*static_cast<QHash<qint32, QByteArray>*>(roles));
}

struct Moc_PackedList PersonModel_People(void* ptr)
{
	return ({ QList<Person*>* tmpValue = new QList<Person*>(static_cast<PersonModel*>(ptr)->people()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList PersonModel_PeopleDefault(void* ptr)
{
	return ({ QList<Person*>* tmpValue = new QList<Person*>(static_cast<PersonModel*>(ptr)->peopleDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void PersonModel_SetPeople(void* ptr, void* people)
{
	static_cast<PersonModel*>(ptr)->setPeople(*static_cast<QList<Person*>*>(people));
}

void PersonModel_SetPeopleDefault(void* ptr, void* people)
{
	static_cast<PersonModel*>(ptr)->setPeopleDefault(*static_cast<QList<Person*>*>(people));
}

void PersonModel_ConnectPeopleChanged(void* ptr)
{
	QObject::connect(static_cast<PersonModel*>(ptr), static_cast<void (PersonModel::*)(QList<Person*>)>(&PersonModel::peopleChanged), static_cast<PersonModel*>(ptr), static_cast<void (PersonModel::*)(QList<Person*>)>(&PersonModel::Signal_PeopleChanged));
}

void PersonModel_DisconnectPeopleChanged(void* ptr)
{
	QObject::disconnect(static_cast<PersonModel*>(ptr), static_cast<void (PersonModel::*)(QList<Person*>)>(&PersonModel::peopleChanged), static_cast<PersonModel*>(ptr), static_cast<void (PersonModel::*)(QList<Person*>)>(&PersonModel::Signal_PeopleChanged));
}

void PersonModel_PeopleChanged(void* ptr, void* people)
{
	static_cast<PersonModel*>(ptr)->peopleChanged(*static_cast<QList<Person*>*>(people));
}

int PersonModel_PersonModel_QRegisterMetaType()
{
	return qRegisterMetaType<PersonModel*>();
}

int PersonModel_PersonModel_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<PersonModel*>(const_cast<const char*>(typeName));
}

int PersonModel_PersonModel_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<PersonModel>();
#else
	return 0;
#endif
}

int PersonModel_PersonModel_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<PersonModel>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int PersonModel_____setItemData_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void PersonModel_____setItemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* PersonModel_____setItemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

int PersonModel_____roleNames_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void PersonModel_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* PersonModel_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

int PersonModel_____itemData_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void PersonModel_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* PersonModel_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

void* PersonModel___setItemData_roles_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QMap<int, QVariant>*>(ptr)->value(i));
}

void PersonModel___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* PersonModel___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>;
}

struct Moc_PackedList PersonModel___setItemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* PersonModel___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void PersonModel___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* PersonModel___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

void* PersonModel___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void PersonModel___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* PersonModel___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

int PersonModel___dataChanged_roles_atList(void* ptr, int i)
{
	return static_cast<QVector<int>*>(ptr)->at(i);
}

void PersonModel___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* PersonModel___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>;
}

void* PersonModel___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i));
}

void PersonModel___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* PersonModel___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>;
}

void* PersonModel___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i));
}

void PersonModel___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* PersonModel___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>;
}

void* PersonModel___roleNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QHash<int, QByteArray>*>(ptr)->value(i));
}

void PersonModel___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* PersonModel___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>;
}

struct Moc_PackedList PersonModel___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* PersonModel___itemData_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QMap<int, QVariant>*>(ptr)->value(i));
}

void PersonModel___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* PersonModel___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>;
}

struct Moc_PackedList PersonModel___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* PersonModel___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void PersonModel___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* PersonModel___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

void* PersonModel___match_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void PersonModel___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* PersonModel___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

void* PersonModel___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void PersonModel___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* PersonModel___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

int PersonModel_____doSetRoleNames_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void PersonModel_____doSetRoleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* PersonModel_____doSetRoleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

int PersonModel_____setRoleNames_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void PersonModel_____setRoleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* PersonModel_____setRoleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

void* PersonModel___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void PersonModel___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* PersonModel___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* PersonModel___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void PersonModel___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* PersonModel___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* PersonModel___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void PersonModel___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* PersonModel___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* PersonModel___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void PersonModel___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* PersonModel___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* PersonModel___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void PersonModel___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* PersonModel___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

void* PersonModel___roles_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QHash<qint32, QByteArray>*>(ptr)->value(i));
}

void PersonModel___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* PersonModel___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<qint32, QByteArray>;
}

struct Moc_PackedList PersonModel___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QHash<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* PersonModel___setRoles_roles_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QHash<qint32, QByteArray>*>(ptr)->value(i));
}

void PersonModel___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* PersonModel___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<qint32, QByteArray>;
}

struct Moc_PackedList PersonModel___setRoles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QHash<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* PersonModel___rolesChanged_roles_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QHash<qint32, QByteArray>*>(ptr)->value(i));
}

void PersonModel___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* PersonModel___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<qint32, QByteArray>;
}

struct Moc_PackedList PersonModel___rolesChanged_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QHash<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* PersonModel___people_atList(void* ptr, int i)
{
	return const_cast<Person*>(static_cast<QList<Person*>*>(ptr)->at(i));
}

void PersonModel___people_setList(void* ptr, void* i)
{
	static_cast<QList<Person*>*>(ptr)->append(static_cast<Person*>(i));
}

void* PersonModel___people_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Person*>;
}

void* PersonModel___setPeople_people_atList(void* ptr, int i)
{
	return const_cast<Person*>(static_cast<QList<Person*>*>(ptr)->at(i));
}

void PersonModel___setPeople_people_setList(void* ptr, void* i)
{
	static_cast<QList<Person*>*>(ptr)->append(static_cast<Person*>(i));
}

void* PersonModel___setPeople_people_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Person*>;
}

void* PersonModel___peopleChanged_people_atList(void* ptr, int i)
{
	return const_cast<Person*>(static_cast<QList<Person*>*>(ptr)->at(i));
}

void PersonModel___peopleChanged_people_setList(void* ptr, void* i)
{
	static_cast<QList<Person*>*>(ptr)->append(static_cast<Person*>(i));
}

void* PersonModel___peopleChanged_people_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Person*>;
}

int PersonModel_____roles_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<qint32>*>(ptr)->at(i);
}

void PersonModel_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* PersonModel_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>;
}

int PersonModel_____setRoles_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<qint32>*>(ptr)->at(i);
}

void PersonModel_____setRoles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* PersonModel_____setRoles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>;
}

int PersonModel_____rolesChanged_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<qint32>*>(ptr)->at(i);
}

void PersonModel_____rolesChanged_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* PersonModel_____rolesChanged_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>;
}

void* PersonModel_NewPersonModel(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new PersonModel(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new PersonModel(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new PersonModel(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new PersonModel(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new PersonModel(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new PersonModel(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new PersonModel(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new PersonModel(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new PersonModel(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new PersonModel(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new PersonModel(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new PersonModel(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new PersonModel(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new PersonModel(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new PersonModel(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new PersonModel(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new PersonModel(static_cast<QWindow*>(parent));
	} else {
		return new PersonModel(static_cast<QObject*>(parent));
	}
}

void PersonModel_DestroyPersonModel(void* ptr)
{
	static_cast<PersonModel*>(ptr)->~PersonModel();
}

char PersonModel_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<PersonModel*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void* PersonModel_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<PersonModel*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* PersonModel_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<PersonModel*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

long long PersonModel_FlagsDefault(void* ptr, void* index)
{
	return static_cast<PersonModel*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}

char PersonModel_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<PersonModel*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char PersonModel_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<PersonModel*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char PersonModel_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<PersonModel*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char PersonModel_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<PersonModel*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char PersonModel_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<PersonModel*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char PersonModel_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<PersonModel*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

char PersonModel_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<PersonModel*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char PersonModel_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<PersonModel*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char PersonModel_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<PersonModel*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

char PersonModel_SubmitDefault(void* ptr)
{
	return static_cast<PersonModel*>(ptr)->QAbstractListModel::submit();
}

void PersonModel_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<PersonModel*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void PersonModel_ResetInternalDataDefault(void* ptr)
{
	static_cast<PersonModel*>(ptr)->QAbstractListModel::resetInternalData();
}

void PersonModel_RevertDefault(void* ptr)
{
	static_cast<PersonModel*>(ptr)->QAbstractListModel::revert();
}

void PersonModel_SortDefault(void* ptr, int column, long long order)
{
	static_cast<PersonModel*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

struct Moc_PackedList PersonModel_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<PersonModel*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList PersonModel_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<PersonModel*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* PersonModel_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<PersonModel*>(ptr)->QAbstractListModel::mimeData(*static_cast<QList<QModelIndex>*>(indexes));
}

void* PersonModel_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<PersonModel*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

void* PersonModel_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<PersonModel*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

struct Moc_PackedList PersonModel_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<PersonModel*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* PersonModel_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<PersonModel*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct Moc_PackedString PersonModel_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t3895e7 = static_cast<PersonModel*>(ptr)->QAbstractListModel::mimeTypes().join("|").toUtf8(); Moc_PackedString { const_cast<char*>(t3895e7.prepend("WHITESPACE").constData()+10), t3895e7.size()-10 }; });
}

void* PersonModel_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void* PersonModel_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<PersonModel*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

long long PersonModel_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<PersonModel*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long PersonModel_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<PersonModel*>(ptr)->QAbstractListModel::supportedDropActions();
}

char PersonModel_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<PersonModel*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char PersonModel_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<PersonModel*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

char PersonModel_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<PersonModel*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

int PersonModel_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<PersonModel*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

int PersonModel_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char PersonModel_EventDefault(void* ptr, void* e)
{
	return static_cast<PersonModel*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char PersonModel_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<PersonModel*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void PersonModel_ChildEventDefault(void* ptr, void* event)
{
	static_cast<PersonModel*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void PersonModel_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<PersonModel*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void PersonModel_CustomEventDefault(void* ptr, void* event)
{
	static_cast<PersonModel*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void PersonModel_DeleteLaterDefault(void* ptr)
{
	static_cast<PersonModel*>(ptr)->QAbstractListModel::deleteLater();
}

void PersonModel_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<PersonModel*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void PersonModel_TimerEventDefault(void* ptr, void* event)
{
	static_cast<PersonModel*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
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
