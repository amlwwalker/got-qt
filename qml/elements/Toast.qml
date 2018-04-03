import QtQuick 2.6
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0
import QtQuick.Controls.Material 2.0

Popup {
    id: popup
    bottomMargin: 80
    background: Rectangle{
        color: "#ddd"
        radius: 3
    }

    height: 40
    modal: true
    focus: true
    closePolicy: Popup.CloseOnEscape | Popup.CloseOnPressOutsideParent
    Timer {
        id: toastTimer
        interval: 3000
        repeat: false
        onTriggered: {
            popup.close()
        }
    } // toastTimer
    Label {
        id: toastLabel
        leftPadding: 16
        rightPadding: 16
        font.pixelSize: 14
        color: "black"
        horizontalAlignment: Text.AlignHCenter
        verticalAlignment: Text.AlignVCenter
    } // toastLabel
    onAboutToShow: {
        toastTimer.start()
    }
    function start(toastText) {
        toastLabel.text = toastText
        if(!toastTimer.running) {
            open()
        } else {
            toastTimer.restart()
        }
    } // function start
} // popup toastPopup
