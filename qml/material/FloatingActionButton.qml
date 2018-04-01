import QtQuick 2.0
import "../assets"
import "../controls" as Awesome
import "../fonts"
Item {
    id: button
    width: 56 * dp
    height: 56 * dp

    property alias color: background.color
    property alias rippleColor: ripple.color
    property alias iconSource: icon.icon

    signal clicked
    FontAwesome {
        id: awesome
        // resource: "qrc:///resource/fontawesome-webfont.ttf"
        resource: "http://maxcdn.bootstrapcdn.com/font-awesome/4.1.0/fonts/fontawesome-webfont.ttf"
    }
    Rectangle {
        id: background
        anchors.fill: parent
        radius: 28 * dp
        visible: false
    }

    PaperShadow {
        id: shadow
        source: background
        depth: button.enabled ? (mouseArea.pressed ? 4 : 2) : 0
    }

    PaperRipple {
        id: ripple
        radius: 28 * dp
        color: "#deffffff"
        mouseArea: mouseArea
    }
    Awesome.Text {
        id: icon
        font.pointSize: 30
        anchors {
            alignWhenCentered: false
            centerIn: parent
        }
        verticalAlignment: Text.AlignVCenter
        horizontalAlignment: Text.AlignHCenter
    }
//    Image {
//        id: icon
//        anchors.centerIn: parent
//        width: 24 * dp
//        height: 24 * dp
//        sourceSize.width: width
//        sourceSize.height: height
//    }

    MouseArea {
        id: mouseArea
        anchors.fill: parent
        enabled: button.enabled
        onClicked: button.clicked()
    }
}
