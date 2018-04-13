## Automated Setup and Building

I have added a makefile to this project to ease configuration.

### Instructions

#### Pre requisite

* Clone this directory:  ` git clone git@github.com:amlwwalker/got-qt.git $(GOPATH)/src/github.com/amlwwalker/got-qt`
* Go to the directory: `cd $(GOPATH)/src/github.com/amlwwalker/got-qt`

#### Installation

* Run setup to install qt5 `make setup`
* Run install to install **therecipe/qt** and **got-qt**: `make install`

#### Creating your project

* It will automatically create a directory for projects at `$(GOPATH)/src/github.com/got-qt-projects`
	* all projects created using this makefile will end up here
* Run make to create a new project, setting the project name as an argument: `make createproject PROJECTNAME=testproject`
* Go to your project directory, `cd $(GOPATH)/src/github.com/got-qt-projects/testproject`
* Build the console app: `make console`. The binary will appear inside cuiinterface
* Build the gui app (only configured for OSX so far): `make darwin`
* Run the binary in development (hotloading) mode: `make hotload`
* You can now edit qml and it will hot load in your new project.