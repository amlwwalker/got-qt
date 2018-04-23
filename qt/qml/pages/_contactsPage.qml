import QtQuick 2.6
import QtQuick.Layouts 1.1
import QtQuick.Controls 2.0
import QtQuick.Controls.Material 2.0
import QtQuick.Dialogs 1.0
import "../elements"
Pane {
    id: pane
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
            ToolTip.timeout: 5000
            ToolTip.visible: hovered
            ToolTip.text: "Click to choose a file, or drag a file here, to send"
            MouseArea {
                anchors.fill: parent
                cursorShape: Qt.PointingHandCursor
                acceptedButtons: Qt.LeftButton | Qt.RightButton
                onClicked: {
                    console.log("Click")
                    if (mouse.button == Qt.LeftButton)
                    {
                        console.log("Left")
                        fPicker.visible = true
                    }
                    else if (mouse.button == Qt.RightButton)
                    {
                        fileDialog.open()
                        footerLabel.text = "Clicked on " + labelText
                    }
                }
            }
            Popup {
                id: fPicker
                visible: false
                modal: true
                focus: true
                width: window.height / 3 * 2
                x: (window.width - width) / 2
                y: -200
                contentHeight: 400
                FilePicker {
                    anchors.fill: parent
                    showDotAndDotDot: true
                    onFileSelected: {
                    }
                }
            }
            DropArea {
                id: drg
                anchors.fill: parent
                onDropped: {
                    //ctrl.processFileForSharing(ctrl.pendingFile)
                    //badgeArea.color = "transparent"
                    //badge.color = Colors.color.steelblue
                    //badge.text = ctrl.person(index).len
                }
                onEntered: {
                    console.log("hello world")
                    //badge.color = "white"
                    //badge.text = FontAwesome.icons.fa_upload
                    //badgeArea.color = Colors.color.balanced
                    //ctrl.pendingFile = drag.urls.toString()
                    //ctrl.pendingContact = txt.text
                }
                onExited: {
                    //badgeArea.color = "transparent"
                    //badge.color = Colors.color.steelblue
                    //badge.text = ctrl.person(index).len
                }
            }
        }
    }

    ColumnLayout {
        spacing: 10
        anchors.fill: parent
        anchors.topMargin: 20

        Label {
            Layout.fillWidth: true
            wrapMode: Label.Wrap
            padding: 20
            topPadding: 0
            horizontalAlignment: Qt.AlignHLeft
            text: "These are your contacts. <br><ul><li>Left click to see files from them</li><li>Right click to choose a file to send them</li><li>You can also drag a file onto a contact to send it to them.</li></ul>"
        }
        ListView {
            id: listView
            Layout.fillWidth: true
            Layout.fillHeight: true
            clip: true
            model: ContactsModel
            delegate: Loader {
                id: delegateLoader
                width: listView.width
                sourceComponent: delegateComponentMap["ItemDelegate"]

                property string labelText: email
                property ListView view: listView
                property int ourIndex: index

                // Can't find a way to do this in the SwipeDelegate component itself,
                // so do it here instead.
                ListView.onRemove: SequentialAnimation {
                    PropertyAction {
                        target: delegateLoader
                        property: "ListView.delayRemove"
                        value: true
                    }
                    NumberAnimation {
                        target: item
                        property: "height"
                        to: 0
                        easing.type: Easing.InOutQuad
                    }
                    PropertyAction {
                        target: delegateLoader
                        property: "ListView.delayRemove"
                        value: false
                    }
                }
            }
        }
    }
}
