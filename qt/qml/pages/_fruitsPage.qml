import QtQuick 2.6
import QtQuick.Layouts 1.1
import QtQuick.Controls 2.0
import QtQuick.Controls.Material 2.0
import "../elements"
import "../images"
Rectangle {
    id: container
    width: 500; height: 400
    color: "#343434"

    // The model:
    ListModel {
        id: fruitModel

        ListElement {
            name: "Apple"; cost: 2.45
            attributes: [
                ListElement { description: "Core" },
                ListElement { description: "Deciduous" }
            ]
        }
        ListElement {
            name: "Banana"; cost: 1.95
            attributes: [
                ListElement { description: "Tropical" },
                ListElement { description: "Seedless" }
            ]
        }
        ListElement {
            name: "Cumquat"; cost: 3.25
            attributes: [
                ListElement { description: "Citrus" }
            ]
        }
        ListElement {
            name: "Durian"; cost: 9.95
            attributes: [
                ListElement { description: "Tropical" },
                ListElement { description: "Smelly" }
            ]
        }
    }

    // The delegate for each fruit in the model:
    Component {
        id: listDelegate

        Item {
            id: delegateItem
            width: listView.width - 10
            height: 55
            clip: true

            Row {
                anchors.verticalCenter: parent.verticalCenter
                spacing: 10

                Column {
                    Image {
                        source: "../images/FA/white/png/22/arrow-up.png"
                        MouseArea { anchors.fill: parent; onClicked: fruitModel.move(index, index-1, 1) }
                    }
                    Image { source: "../images/FA/white/png/22/arrow-down.png"
                        MouseArea { anchors.fill: parent; onClicked: fruitModel.move(index, index+1, 1) }
                    }
                }

                Column {
                    anchors.verticalCenter: parent.verticalCenter

                    Text {
                        text: name
                        font.pixelSize: 15
                        color: "white"
                    }
                    Row {
                        spacing: 5
                        Repeater {
                            model: attributes
                            Text { text: description; color: "White" }
                        }
                    }
                }
            }

            Row {
                anchors.verticalCenter: parent.verticalCenter
                anchors.right: parent.right
                spacing: 10

                PressAndHoldButton {
                    anchors.verticalCenter: parent.verticalCenter
                    source: "../images/FA/white/png/22/minus.png"
                    onClicked: fruitModel.setProperty(index, "cost", cost + 0.25)
                }

                Text {
                    id: costText
                    anchors.verticalCenter: parent.verticalCenter
                    text: '$' + Number(cost).toFixed(2)
                    font.pixelSize: 15
                    color: "white"
                    font.bold: true
                }

                PressAndHoldButton {
                    anchors.verticalCenter: parent.verticalCenter
                    source: "../images/FA/white/png/22/plus.png"
                    onClicked: fruitModel.setProperty(index, "cost", Math.max(0,cost-0.25))
                }

                Image {
                    source: "../images/FA/white/png/22/remove.png"
                    MouseArea { anchors.fill:parent; onClicked: fruitModel.remove(index) }
                }
            }

            // Animate adding and removing of items:

            ListView.onAdd: SequentialAnimation {
                PropertyAction { target: delegateItem; property: "height"; value: 0 }
                NumberAnimation { target: delegateItem; property: "height"; to: 55; duration: 250; easing.type: Easing.InOutQuad }
            }

            ListView.onRemove: SequentialAnimation {
                PropertyAction { target: delegateItem; property: "ListView.delayRemove"; value: true }
                NumberAnimation { target: delegateItem; property: "height"; to: 0; duration: 250; easing.type: Easing.InOutQuad }

                // Make sure delayRemove is set back to false so that the item can be destroyed
                PropertyAction { target: delegateItem; property: "ListView.delayRemove"; value: false }
            }
        }
    }

    // The view:
    ListView {
        id: listView
        anchors.fill: parent; anchors.margins: 20
        model: fruitModel
        delegate: listDelegate
    }

    Row {
        anchors { left: parent.left; bottom: parent.bottom; margins: 20 }
        spacing: 10

        TextButton {
            text: "Add an item"
            onClicked: {
                fruitModel.append({
                    "name": "Pizza Margarita",
                    "cost": 5.95,
                    "attributes": [{"description": "Cheese"}, {"description": "Tomato"}]
                })
            }
        }

        TextButton {
            text: "Remove all items"
            onClicked: fruitModel.clear()
        }
    }
}