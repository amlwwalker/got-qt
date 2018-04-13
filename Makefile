# Borrowed from:
# https://github.com/silven/go-example/blob/master/Makefile
# https://vic.demuzere.be/articles/golang-makefile-crosscompile/

BINARY = got-qt

# Symlink into GOPATH
GITHUB_USERNAME=amlwwalker
GOT_QT_PROJECTS=got-qt-projects
CONSOLE_DIR=cuiinterface
DEV_SOURCE=${GOPATH}/src/github.com
GUI_DIR=qt
BUILD_DIR=${DEV_SOURCE}/${GITHUB_USERNAME}/${BINARY}
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
	cd ${DEV_SOURCE}; \
	#now make a directory to store the got-qt
	mkdir ${GITHUB_USERNAME}; \
	#now clone got-qt into there
	git clone ${GOT_QT_REPO} ${GITHUB_USERNAME}/ \
	#now make sure the user has therecipe/qt installed
	go get -u -v github.com/therecipe/qt/cmd/...; \
	$GOPATH/bin/qtsetup;

#var topLevel = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "amlwwalker", "got-qt", "qt", "qml")
createproject:
	#create the directory in github.com
	cd ${DEV_SOURCE}; \
	mkdir -p ${GOT_QT_PROJECTS}/${PROJECTNAME}; \
	cd ${GOT_QT_PROJECTS}/${PROJECTNAME}; \
	rsync -av --progress --exclude=".git" ${DEV_SOURCE}/${GITHUB_USERNAME}/got-qt/* .
	# cp -r ${DEV_SOURCE}/${GITHUB_USERNAME}/got-qt `ls -A | grep -v ".git"` . \

	# cd ${PROJECTNAME}; \ #we are going to update the toplevel domain for hotloading
	sed -i.bak s/"amlwwalker"/"${got-qt-projects}"/g qt/main.go \
	sed -i.bak s/"got-qt"/"${PROJECTNAME}"/g qt/main.go

console:
	cd ${BUILD_DIR}; \
	cd ${CONSOLE_DIR}; \
	go build -o ${BINARY}-console . ; \
	cd - >/dev/null

linux:
	cd ${BUILD_DIR}; \
	#GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}-linux-${GOARCH} . ; \

	cd - >/dev/null

darwin:
	cd ${BUILD_DIR}; \
	cd ${GUI_DIR}; \
	# GOOS=darwin go build ${LDFLAGS} -o ${BINARY}-darwin-${GOARCH} . ; \

	cd - >/dev/null

windows:
	cd ${BUILD_DIR}; \
	cd ${GUI_DIR}; \
	#GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}-windows-${GOARCH}.exe . ; \

	cd - >/dev/null

fmt:
	cd ${BUILD_DIR}; \
	go fmt $$(go list ./... | grep -v /vendor/) ; \
	cd - >/dev/null

clean:
	cd ${BUILD_DIR}; \
	rm -rf ${CONSOLE_DIR}/${BINARY}-* \
	rm -rf ${GUI_DIR}/deploy