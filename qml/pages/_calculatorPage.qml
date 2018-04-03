import QtQuick 2.6
import QtQuick.Layouts 1.1
import QtQuick.Controls 2.0
import QtQuick.Controls.Material 2.0

Pane {
    padding: 0
    Column {
        spacing: 40
        width: parent.width

        Label {
            width: parent.width
            wrapMode: Label.Wrap
            horizontalAlignment: Qt.AlignHCenter
            text: "This page demonstrates a simple calculator, using Go to add two values together. The answer is returned to the UI"
        }
        Row {
            id: row
            anchors.left: parent.left
            anchors.right: parent.right

            TextField {
                id: prefix
                placeholderText: "10"
                horizontalAlignment: Qt.AlignHCenter
                width: (row.width/2) - 10
            }
            Label {
                wrapMode: Label.Wrap
                horizontalAlignment: Qt.AlignHCenter
                text: " + "
            }
            TextField {
                id: suffix
                placeholderText: "5"
                horizontalAlignment: Qt.AlignHCenter
                width: (row.width/2) - 10
            }
        }
        Button {
            width: parent.width
            text: "Get Result!"
            onClicked: function() {
            console.log('updating answer')
                answer.text = QmlBridge.calculator(Number(suffix.text), Number(prefix.text))
            }
        }
        Label {
            id: answer
            width: parent.width
            wrapMode: Label.Wrap
            horizontalAlignment: Qt.AlignHCenter
            text: ""
        }
    }
}