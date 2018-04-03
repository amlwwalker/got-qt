import QtQuick 2.6
import QtQuick.Controls 2.0
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

            Button {
                id: button
                text: "Show Toast"
                anchors.horizontalCenter: parent.horizontalCenter
                width: Math.max(implicitWidth, Math.min(implicitWidth * 2, page.availableWidth / 3))

                onClicked: function() {
                    globalToast.open()
                    globalToast.start("hello from toast")
                }

            }
    }
}
