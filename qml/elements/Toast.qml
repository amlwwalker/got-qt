import QtQuick 2.6
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0
import QtQuick.Controls.Material 2.0

Popup {
    id: popup
    closePolicy: Popup.NoAutoClose
    bottomMargin: 80
    background: Rectangle{
        color: "#ddd"
        radius: 10
    }
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
        font.pixelSize: 16
        color: "black"
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
