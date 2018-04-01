import QtQuick 2.0
import "../controls" as Awesome
import "../fonts"
Item {
    id: button
    property real dp: 1 //can set a dp from dpi from screen resolution

    width: 24 * dp
    height: 24 * dp

    property alias rippleColor: ripple.color
    property string iconSource

    signal clicked
    FontAwesome {
        id: awesome
        // resource: "qrc:///resource/fontawesome-webfont.ttf"
        resource: "http://maxcdn.bootstrapcdn.com/font-awesome/4.1.0/fonts/fontawesome-webfont.ttf"
    }
    PaperRipple {
        id: ripple
        anchors {
            fill: undefined
            centerIn: parent
        }
        width: 40 * dp
        height: 40 * dp
        radius: 20 * dp
        mouseArea: mouseArea
    }

    Awesome.Text {
        font.pointSize: 30
        icon: awesome.icons.fa_money
    }
//   Image {
//       id: icon
//       anchors.centerIn: parent
//       width: 24 * dp
//       height: 24 * dp
//       sourceSize.width: width
//       sourceSize.height: height
//       opacity: button.enabled ? 1 : 0.62

//       Behavior on opacity {
//           NumberAnimation {
//               duration: 200
//               easing.type: Easing.Bezier; easing.bezierCurve: [0.4, 0, 0.2, 1, 1, 1]
//           }
//       }
//   }

    MouseArea {
        id: mouseArea
        anchors.fill: ripple
        enabled: button.enabled
        onClicked: button.clicked()
    }
}
