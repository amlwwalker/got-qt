import QtQuick 2.0
import QtQuick.Controls 1.0
import QtQuick.Layouts 1.0

Text {
    id: root

    property alias  spacing: row.spacing
    property alias  text: content.text
    property color  color: "black"
    property font   font
    property string icon

    RowLayout {
        id: row

        Text {
            color: root.color
            font.pointSize: root.font.pointSize
            font.family: awesome.family
            renderType: Text.NativeRendering
            text: awesome.loaded ? icon : ""
        }

        Text {
            id: content
            color: root.color
            font.pointSize: root.font.pointSize
            renderType: Text.NativeRendering
        }
    }
}