/*
    Copyright (c) 2014 Cutehacks A/S

    Permission is hereby granted, free of charge, to any person obtaining
    a copy of this software and associated documentation files (the
    "Software"), to deal in the Software without restriction, including
    without limitation the rights to use, copy, modify, merge, publish,
    distribute, sublicense, and/or sell copies of the Software, and to
    permit persons to whom the Software is furnished to do so, subject to
    the following conditions:

    The above copyright notice and this permission notice shall be
    included in all copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
    EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
    MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
    IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
    CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
    TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
    SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.


    About:

    This is a Qt Quick implementation of the Android Navigation drawer
    Questions, suggestions or requests can be directed to jens@cutehacks.com

    www.cutehacks.com
*/


import QtQuick 2.2

Rectangle {
    id: panel

    property bool open: false                     // The open or close state of the drawer
    property int position: Qt.LeftEdge            // Which side of the screen the drawer is on, can be Qt.LeftEdge or Qt.RightEdge
    property Item visualParent: parent            // The item the drawer should cover, by default the parent of the Drawer

    // The fraction showing how open the drawer is
    readonly property real panelProgress:  (panel.x - _minimumX) / (_maximumX - _minimumX)

    function show() { open = true; }
    function hide() { open = false; }

    function toggle() {
        if (open) open = false;
        else open = true;
    }

    // Internal

    default property alias data: contentItem.data
    readonly property real expandedFraction: 0.78  // How big fraction of the screen realesatate that is covered by an open menu
    readonly property real _scaleFactor: _rootItem.width / 320 // Note, this should really be application global
    readonly property int _pullThreshold: panel.width/2
    readonly property int _slideDuration: 260
    readonly property int _collapsedX: _rightEdge ? _rootItem.width :  - panel.width
    readonly property int _expandedWidth: expandedFraction * _rootItem.width
    readonly property int _expandedX: _rightEdge ? _rootItem.width - width : 0
    readonly property bool _rightEdge: position === Qt.RightEdge
    readonly property int _minimumX:  _rightEdge ?  _rootItem.width - panel.width : -panel.width
    readonly property int _maximumX: _rightEdge ? _rootItem.width : 0
    readonly property int _openMarginSize: 20 * _scaleFactor

    property real _velocity: 0
    property real _oldMouseX: -1

    function _findRootItem() {
        var p = panel;
        while (p.parent != null)
            p = p.parent;
        return p;
    }

    property Item _rootItem: _findRootItem()
    height: parent.height
    on_RightEdgeChanged: _setupAnchors()
    onOpenChanged: completeSlideDirection()
    width: _expandedWidth
    x: _collapsedX
    z: 10

    function _setupAnchors() {     // Note that we can't reliably apply anchors using bindings
        _rootItem = _findRootItem();

        shadow.anchors.right = undefined;
        shadow.anchors.left = undefined;

        mouse.anchors.left = undefined;
        mouse.anchors.right = undefined;

        if (_rightEdge) {
            mouse.anchors.right = mouse.parent.right;
            shadow.anchors.right = panel.left;
        } else {
            mouse.anchors.left = mouse.parent.left;
            shadow.anchors.left = panel.right;
        }

        slideAnimation.enabled = false;
        panel.x =_rightEdge ? _rootItem.width :  - panel.width;
        slideAnimation.enabled = true;
    }

    function completeSlideDirection() {
        if (open) {
            panel.x = _expandedX;
        } else {
            panel.x = _collapsedX;
            Qt.inputMethod.hide();
        }
    }

    function handleRelease() {
        var velocityThreshold = 5 * _scaleFactor;
        if ((_rightEdge && _velocity > velocityThreshold) ||
                (!_rightEdge && _velocity < -velocityThreshold)) {
            panel.open = false;
            completeSlideDirection()
        } else if ((_rightEdge && _velocity < -velocityThreshold) ||
                   (!_rightEdge && _velocity > velocityThreshold)) {
            panel.open = true;
            completeSlideDirection()
        } else if ((_rightEdge && panel.x < _expandedX + _pullThreshold) ||
                   (!_rightEdge && panel.x > _expandedX - _pullThreshold) ) {
            panel.open = true;
            panel.x = _expandedX;
        } else {
            panel.open = false;
            panel.x = _collapsedX;
        }
    }

    function handleClick(mouse) {
        if ((_rightEdge && mouse.x < panel.x ) || mouse.x > panel.width) {
            open = false;
        }
    }

    onPositionChanged: {
        if (! (position === Qt.RightEdge || position === Qt.LeftEdge ) ) {
            console.warn("SlidePanel: Unsupported position.")
        }
    }

    Behavior on x {
        id: slideAnimation
        enabled: !mouse.drag.active
        NumberAnimation {
            duration: _slideDuration
            easing.type: Easing.OutCubic
        }
    }

    Connections {
        target: _rootItem
        onWidthChanged: {
            slideAnimation.enabled = false
            panel.completeSlideDirection()
            slideAnimation.enabled = true
        }
    }

    NumberAnimation on x {
        id: holdAnimation
        to: _collapsedX + (_openMarginSize * (_rightEdge ? -1 : 1))
        running : false
        easing.type: Easing.OutCubic
        duration: 200
    }

    MouseArea {
        id: mouse
        parent: _rootItem

        y: visualParent.y
        width: open ? _rootItem.width : _openMarginSize
        height: visualParent.height
        onPressed:  if (!open) holdAnimation.restart();
        onClicked: handleClick(mouse)
        drag.target: panel
        drag.minimumX: _minimumX
        drag.maximumX: _maximumX
        drag.axis: Qt.Horizontal
        drag.onActiveChanged: if (active) holdAnimation.stop()
        onReleased: handleRelease()
        z: open ? 1 : 0
        onMouseXChanged: {
            _velocity = (mouse.x - _oldMouseX);
            _oldMouseX = mouse.x;
        }
    }

    Rectangle {
        id: backgroundDimmer
        parent: visualParent
        anchors.fill: parent
        opacity: 0.5 * Math.min(1, Math.abs(panel.x - _collapsedX) / _rootItem.width/2)
        color: "black"
    }

    Item {
        id: contentItem
        parent: _rootItem
        width: panel.width
        height: panel.height
        x: panel.x
        y: panel.y
        z: open ? 5 : 0
        clip: true
    }

    Item {
        id: shadow
        anchors.left: panel.right
        anchors.leftMargin: _rightEdge ? 0 : 4 * _scaleFactor
        height: parent.height
        Rectangle {
            height: 4 * _scaleFactor
            width: panel.height
            rotation: 90
            opacity: Math.min(1, Math.abs(panel.x - _collapsedX)/_openMarginSize)
            transformOrigin: Item.TopLeft
            gradient: Gradient{
                GradientStop { position: _rightEdge ? 1 : 0 ; color: "#00000000"}
                GradientStop { position: _rightEdge ? 0 : 1 ; color: "#2c000000"}
            }
        }
    }
}