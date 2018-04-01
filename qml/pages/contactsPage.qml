import QtQuick 2.6
import QtQuick.Layouts 1.1
import QtQuick.Controls 2.0

Pane {
    padding: 0

    property var delegateComponentMap: {
        "ItemDelegate": itemDelegateComponent,
        "SwipeDelegate": swipeDelegateComponent,
        "CheckDelegate": checkDelegateComponent,
        "RadioDelegate": radioDelegateComponent,
        "SwitchDelegate": switchDelegateComponent
    }

    Component {
        id: itemDelegateComponent

        ItemDelegate {
            text: labelText
            width: parent.width
        }
    }

    Component {
        id: swipeDelegateComponent

        SwipeDelegate {
            id: swipeDelegate
            text: labelText
            width: parent.width

            onClicked: if (swipe.complete) view.model.remove(ourIndex)

            Component {
                id: removeComponent

                Rectangle {
                    color: swipeDelegate.swipe.complete && swipeDelegate.pressed ? "#333" : "#444"
                    width: parent.width
                    height: parent.height
                    clip: true

                    Label {
                        font.pixelSize: swipeDelegate.font.pixelSize
                        text: "Remove"
                        color: "white"
                        anchors.centerIn: parent
                    }
                }
            }

            swipe.left: removeComponent
            swipe.right: removeComponent
        }
    }

    Component {
        id: checkDelegateComponent

        CheckDelegate {
            text: labelText
            width: parent.width
        }
    }

    ButtonGroup {
        id: radioButtonGroup
    }

    Component {
        id: radioDelegateComponent

        RadioDelegate {
            text: labelText
            ButtonGroup.group: radioButtonGroup
        }
    }

    Component {
        id: switchDelegateComponent

        SwitchDelegate {
            text: labelText
        }
    }

    ColumnLayout {
        id: column
        spacing: 40
        anchors.fill: parent
        anchors.topMargin: 20

        Label {
            Layout.fillWidth: true
            wrapMode: Label.Wrap
            horizontalAlignment: Qt.AlignHCenter
            text: "Delegate controls are used as delegates in views such as ListView."
        }

        ListView {
            id: listView
            Layout.fillWidth: true
            Layout.fillHeight: true
            clip: true
            model: ListModel {
                ListElement { type: "RadioDelegate"; text: "Contacts" }
            }

            section.property: "type"
            section.delegate: Pane {
                width: listView.width
                height: sectionLabel.implicitHeight + 20

                Label {
                    id: sectionLabel
                    text: section
                    anchors.centerIn: parent
                }
            }

            delegate: Loader {
                id: delegateLoader
                width: listView.width
                sourceComponent: delegateComponentMap[text]

                property string labelText: text
                property ListView view: listView
                property int ourIndex: index

                // Can't find a way to do this in the SwipeDelegate component itself,
                // so do it here instead.
                ListView.onRemove: SequentialAnimation {
                    PropertyAction {
                        target: delegateLoader
                        property: "ListView.delayRemove"
                        value: true
                    }
                    NumberAnimation {
                        target: item
                        property: "height"
                        to: 0
                        easing.type: Easing.InOutQuad
                    }
                    PropertyAction {
                        target: delegateLoader
                        property: "ListView.delayRemove"
                        value: false
                    }
                }
            }
        }
    }
}
