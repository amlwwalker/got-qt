import QtQuick 2.1
import QtQuick.Controls 1.0

ApplicationWindow {
    title: "Hello World"
    width: 600
    height: 300

    FileSaveDialog {
        id: saveFile
        title: "Save file"
        filename: "download.png"

        onAccepted: {
            output.text = "File selected: " + saveFile.fileUrl
        }
        onRejected: {
            output.text = "File selected: –"
        }
    }

    Column {
        Button {
            text: "Select File"
            onClicked: { saveFile.open(); }
        }

        Text { id: output }
    }
}