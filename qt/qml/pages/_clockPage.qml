import QtQuick 2.6
import QtQuick.Layouts 1.1
import QtQuick.Controls 2.0
import QtQuick.Controls.Material 2.0

Pane {
    padding: 0
    id: page
    Column {
        spacing: 40
        width: parent.width

        Label {
            width: parent.width
            wrapMode: Label.Wrap
            horizontalAlignment: Qt.AlignHCenter
            text: "Demonstrates receving information from Go backend"
        }

        GroupBox {
            title: "Clock"
            anchors.horizontalCenter: parent.horizontalCenter

            Column {
                spacing: 20
                width: page.itemWidth

                Text
                {
                    id: clock

                    color: "black"
                    font.pointSize: 16
                    opacity: 0.75

                    Connections
                    {
                        target: QmlBridge
                        onSendTime:
                        {
                            clock.text = data
                        }
                    }
                }
            }
        }
    }

}