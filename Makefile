# Borrowed from:
# https://github.com/silven/go-example/blob/master/Makefile
# https://vic.demuzere.be/articles/golang-makefile-crosscompile/

BINARY = got-qt

# Symlink into GOPATH
GITHUB_USERNAME=amlwwalker
GOT_QT_PROJECTS=got-qt-projects
CONSOLE_DIR=cuiinterface
GUI_DIR=qt
DEV_SOURCE=${GOPATH}/src/github.com
# BUILD_DIR=${DEV_SOURCE}/${GOT_QT_PROJECTS}/${BINARY}
CURRENT_DIR=$(shell pwd)

#install parameters
GOT_QT_REPO=git@github.com:amlwwalker/got-qt.git

# Build the project
all: clean console darwin #linux windows

hardinstall: setup install

setup:
	#install everything including qt
	brew install qt5
#configure the setup for this
install:
	#first go to build dir
	cd ${DEV_SOURCE}; \ #now make a directory to store the got-qt
	mkdir ${GITHUB_USERNAME}; \ #now clone got-qt into there
	git clone ${GOT_QT_REPO} ${GITHUB_USERNAME}/ \
	mkdir 0xAX; \ #you will need notifier for notifications
	git clone git@github.com:0xAX/notificator.git 0xAX; \
	go get -u github.com/gobuffalo/packr/...; \ #you will need packr for configuration files
	go get github.com/radovskyb/watcher; \ #you will need watcher by radovskyb for hotloading
	go get -u -v github.com/therecipe/qt/cmd/...; \ #now make sure the user has therecipe/qt installed
	$GOPATH/bin/qtsetup;

#var topLevel = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "amlwwalker", "got-qt", "qt", "qml")
createproject: #pass this PROJECTNAME, e.g: make createproject PROJECTNAME=testproject
	#create the directory in github.com
	cd ${DEV_SOURCE}; \
	mkdir -p ${GOT_QT_PROJECTS}/${PROJECTNAME}; \
	cd ${GOT_QT_PROJECTS}/${PROJECTNAME}; \
	rsync -av --exclude=".git" ${DEV_SOURCE}/${GITHUB_USERNAME}/got-qt/* . ; \
	sed -i.bak s/"got-qt"/"${PROJECTNAME}"/g qt/main.go; \
	sed -i.bak s/"amlwwalker"/"${GOT_QT_PROJECTS}"/g qt/main.go;


console:
	cd ${CONSOLE_DIR}; \
	go build -o consoleApp; \
	cd - >/dev/null

build:
	cd ${GUI_DIR}; \
	qtdeploy build desktop; \
	cd - >/dev/null

darwinhotload:
	cd ${GUI_DIR}; \
	deploy/darwin/${GUI_DIR}.app/Contents/MacOS/${GUI_DIR}
lxhotload:
	cd ${GUI_DIR}; \
	deploy/linux/${GUI_DIR}.sh
