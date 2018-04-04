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
                text: "Login using Google Authentication. Make sure you have created an account first here: <a href='http://website.co.uk'>website.co.uk</a>"
                onLinkActivated: Qt.openUrlExternally(link)
                MouseArea {
                    anchors.fill: parent
                    acceptedButtons: Qt.NoButton
                    cursorShape: parent.hoveredLink ? Qt.PointingHandCursor : Qt.ArrowCursor
                }
            }
            Button {
                id: button
                text: "Login"
                anchors.horizontalCenter: parent.horizontalCenter
                width: Math.max(implicitWidth, Math.min(implicitWidth * 2, page.availableWidth / 3))

                onClicked: function() {
                    globalToast.open()
                    globalToast.start("Attempting to log in")
            }
        }
    }
}
