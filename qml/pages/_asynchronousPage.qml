import QtQuick 2.6
import QtQuick.Controls 2.0
import "../elements"
ScrollablePage {
    id: page

    Column {
        spacing: 40
        width: parent.width

        Label {
            width: parent.width
            wrapMode: Label.Wrap
            horizontalAlignment: Qt.AlignHCenter
            text: "This page demos using a progress bar (found on the layout page) to update the UI of the progres of a backend process (for instance downloading a file)"
        }

        Button {
            id: button
            text: "Start Process"
            anchors.horizontalCenter: parent.horizontalCenter
            width: Math.max(implicitWidth, Math.min(implicitWidth * 2, page.availableWidth / 3))

            onClicked: function() {
                footerLabel.visible = false
                progressIndicator.value = 0
                progressIndicator.visible = true
                QmlBridge.startAsynchronousProcess()
            }
            Connections {
                target: QmlBridge
//Progress bar update
                onUpdateProcessStatus: {
                    progressIndicator.value = c
                    if (c.toFixed(2) ==  1.0) {
                        //process complete
                        progressIndicator.visible = false
                        footerLabel.text = "process complete"
                        footerLabel.visible = true
                    }
                }
            }
        }
    }
}
