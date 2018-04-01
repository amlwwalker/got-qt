/****************************************************************************
**
**
** Copyright (c) 2014 Ricardo do Valle Flores de Oliveira
**
** $BEGIN_LICENSE:MIT$
** Permission is hereby granted, free of charge, to any person obtaining a copy
** of this software and associated documentation files (the "Software"), to deal
** in the Software without restriction, including without limitation the rights
** to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
** copies of the Software, and to permit persons to whom the Software is
** furnished to do so, subject to the following conditions:
**
** The above copyright notice and this permission notice shall be included in
** all copies or substantial portions of the Software.
**
** THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
** IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
** FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
** AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
** LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
** OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
** SOFTWARE.
**
**
****************************************************************************/

import QtQuick 2.0
import QtQuick.Controls 1.0
import QtQuick.Controls.Styles 1.0
import QtQuick.Layouts 1.0

Button {
    id: button
    property string icon
    property color color: "black"
    property font font

    style: ButtonStyle {
        id: buttonstyle
        property font font: button.font
        property color foregroundColor: button.color

        background: Item {
            Rectangle {
                id: baserect
                anchors.fill: parent
                color: "transparent"
            }
        }

        label: Item {
            implicitWidth: row.implicitWidth
            implicitHeight: row.implicitHeight

            RowLayout {
                id: row
                anchors.centerIn: parent
                spacing: 15

                Text {
                    color: buttonstyle.foregroundColor
                    font.pointSize: buttonstyle.font.pointSize * 2
                    font.family: awesome.family
                    renderType: Text.NativeRendering
                    text: awesome.loaded ? icon : ""
                    visible: !(icon === "")
                }
                Text {
                    color: buttonstyle.foregroundColor
                    font: buttonstyle.font
                    renderType: Text.NativeRendering
                    text: control.text
                    visible: !(control.text === "")

                    Layout.alignment: Qt.AlignBottom
                }
            }
        }
    }
}
