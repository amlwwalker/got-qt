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
* Build the gui app (only configured for OSX so far): `make build`
* Run the binary in development (hotloading) mode: `make darwinhotload` for OSX, or `make lxhotload` for linux. (Not tested on Windows).
* You can now edit qml and it will hot load in your new project.

#### Device building

When building for devices such as Android or iOS(untested), hotloading must be disabled.

### Updating the production loader qml file

At this point this is quite crude. So that you can be certain you will get the latest qml changes when hotloading is disabled, that you have copied loader.qml to loader-production.qml and updated the paths to the compiled paths. More information [in the Readme under the production section](https://github.com/amlwwalker/got-qt/#production)