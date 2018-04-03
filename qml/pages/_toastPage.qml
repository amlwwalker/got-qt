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
            text: "This page demos a toast element"
        }

        Toast {
            id: popup
            x: parent.width / 10
            y: (page.height * 4) / 5
            width: (page.width * 4) / 5
        }
        Button {
            id: button
            text: "Show Toast"
            anchors.horizontalCenter: parent.horizontalCenter
            width: Math.max(implicitWidth, Math.min(implicitWidth * 2, page.availableWidth / 3))

            onClicked: function() {
                popup.open()
                popup.start("hello from toast")
        }


        }
    }
}
