//source: http://doc.qt.io/qt-5/qtquickcontrols2-material.html

import QtQuick 2.0
import QtQuick.Controls 2.0
import QtQuick.Controls.Material 2.0
import QtQuick.Controls.Universal 2.0
import "variables/colors.js" as Colors
import "material"
import "assets"
//import "variables/fontawesome.js" as FontAwesome
import "controls" as Awesome
import "fonts"


Rectangle {
    id: root
    width: 360 //isMobile ? 320 : 360
    height: 720 //isMobile ? 480 : 720
    visible: true
    color: "white"
    property real dp: 1 //can set a dp from dpi from screen resolution
    FontAwesome {
        id: awesome
        // resource: "qrc:///resource/fontawesome-webfont.ttf"
        resource: "http://maxcdn.bootstrapcdn.com/font-awesome/4.1.0/fonts/fontawesome-webfont.ttf"
    }
    ActionBar {
        id: actionBar
        raised: contents.contentY > height
        color: "#455ede"
        text: "Qt Material"
        z: 2

        IconButton {
            id: menuButton
            anchors.left: parent.left
            anchors.leftMargin: 16 * dp
            anchors.verticalCenter: parent.verticalCenter
            iconSource: awesome.icons.fa_money
            onClicked: menu.toggle()
        }

        RefreshButton {
            id: refreshButton
            anchors.right: parent.right
            anchors.rightMargin: 16 * dp
            anchors.verticalCenter: parent.verticalCenter
            loading: false
            onClicked: loading = !loading
        }
    }

    Flickable {
        id: contents
        anchors.fill: parent
        anchors.topMargin: actionBar.height
        contentWidth: width
        contentHeight: flow.height + 52 * dp

        Column {
            id: flow
            anchors.top: parent.top
            anchors.left: parent.left
            anchors.right: parent.right
            anchors.margins: 16 * dp
            spacing: 16 * dp

            Label {
                width: parent.width
                font.pointSize: UIConstants.subheadFontSize
                text: "Flat buttons"
            }

            Row {
                spacing: 8 * dp

                FlatButton {
                    text: "Button"
                }

                FlatButton {
                    text: "Colored"
                    textColor: "#5677fc"
                }

                FlatButton {
                    text: "Disabled"
                    enabled: false
                }
            }

            Label {
                width: parent.width
                font.pointSize: UIConstants.subheadFontSize
                text: "Raised buttons"
            }

            Row {
                spacing: 8 * dp

                RaisedButton {
                    text: "Button"
                }

                RaisedButton {
                    text: "Colored"
                    color: "#5677fc"
                    textColor: "white"
                    rippleColor: "#deffffff"
                }

                RaisedButton {
                    text: "Disabled"
                    enabled: false
                }
            }

            Label {
                width: parent.width
                font.pointSize: UIConstants.subheadFontSize
                text: "Text field"
            }

            TextField {
                width: parent.width
                hint: "Type some text"
                Keys.onEscapePressed: focus = false
            }
          Text {
              id: text
              font.pointSize: 180
              font.family: awesome.family
              text: awesome.loaded ? awesome.icons.fa_money : ""
          }
            Label {
                width: parent.width
                font.pointSize: UIConstants.subheadFontSize
                //text: "Ripples"
                //text: "\uf042"
                Awesome.Text {
                icon: awesome.icons.fa_money
                }
            }


            Item {
                width: parent.width
                height: parent.width * 0.67

                Image {
                    id: image
                    anchors.fill: parent
                    anchors.margins: -2 * dp
                    source: "qrc://./qml/assets/lawyer_cat"
                    fillMode: Image.PreserveAspectCrop
                    clip: true
                    visible: false
                }

                PaperShadow {
                    id: shadow
                    anchors.fill: image
                    source: image
                    depth: imageArea.pressed ? 3 : 1
                }

                PaperRipple {
                    anchors.fill: image
                    color: "#deffffff"
                    mouseArea: imageArea
                }

                MouseArea {
                    id: imageArea
                    anchors.fill: image
                }
            }
        }
    }

    FloatingActionButton {
        id: addButton
        anchors {
            right: parent.right
            bottom: parent.bottom
            margins: 16 * dp
        }
        transform: Translate {
            y: dialog.visible ? 72 * dp : 0

            Behavior on y {
                NumberAnimation {
                    duration: 200
                    easing.type: Easing.Bezier; easing.bezierCurve: [0.4, 0, 0.2, 1, 1, 1]
                }
            }
        }

        color: "#ff5177"
        iconSource: awesome.icons.fa_money
        onClicked: dialog.open()
    }
    NavigationDrawer {
        id: menu
    }
    ExampleDialog {
        id: dialog
    }
}