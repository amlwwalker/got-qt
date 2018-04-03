import QtQuick 2.6
import QtQuick.Layouts 1.1
import QtQuick.Controls 2.0
import QtQuick.Controls.Material 2.0
Pane {
    padding: 0
    property var delegateComponentMap: {
        "ItemDelegate": itemDelegateComponent
    }

    Component {
        id: itemDelegateComponent
        ItemDelegate {
            text: labelText
            width: parent.width
            Material.foreground: Material.BlueGrey
            MouseArea {
                anchors.fill: parent
                cursorShape: Qt.PointingHandCursor
                onClicked: {
                    footerLabel.text = "Clicked on " + labelText

                }

            }
            ToolTip.timeout: 5000
            ToolTip.visible: hovered
            ToolTip.text: "Click to add contact to your contacts"
        }
    }
    ColumnLayout {
        spacing: 10
        anchors.fill: parent
        anchors.topMargin: 20

        Label {
            Layout.fillWidth: true
            wrapMode: Label.Wrap
            horizontalAlignment: Qt.AlignHCenter
            padding: 20
            topPadding: 0
            text: "Search for new contacts here. Click on a search result to add them to your contacts"
        }
        TextField {
            id: searchContactField
            Layout.fillWidth: true
            placeholderText: "James Bond"
            horizontalAlignment: Qt.AlignHCenter
            Keys.onReturnPressed: {
                console.log("the user is searching for " + searchContactField.text)
                footerLabel.text = "searching for " + searchContactField.text
                searchContactField.text = ""
            }
            background: Rectangle {
                border.color: "grey"
                radius: 2
            }
        }
        Button {
            id: searchButton
            text: "Search"
            Layout.fillWidth: true
            onClicked: function() {
                footerLabel.text = "searching for " + searchContactField.text
                searchContactField.text = ""
                QmlBridge.searchFor(searchContactField.text)
            }
        }
        ListView {
            id: listView
            Layout.fillWidth: true
            Layout.fillHeight: true
            clip: true
            spacing: 2
            model: SearchModel
            delegate: Loader {
                id: delegateLoader
                width: listView.width
                sourceComponent: delegateComponentMap["ItemDelegate"]

                property string labelText: email
                property ListView view: listView
                property int ourIndex: index
            }
        }
    }
}
