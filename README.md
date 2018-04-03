## Go + QT GUI Framework

This is a framework to make desktop/mobile applications in Go with a GUI written in QT Qml. Both of these languages are cross platform. Go is an open source programming language, however QT is not open source. The license is outlined however below.
I have made this because I:

* Wanted to make a Desktop Application that was as easy to use for anyone who could use a computer.
* Wanted to write it in a language that as cross platform so that I would only need to write it once.
* Use a language I was already familiar with.
* From research, there were a few options for the UI, for instance HTML/CSS/JS. However this required use of the installed default web browser and I didn't want to rely on that, but also it meant that the UI files had to be shipped around with the application. Felt a little messy and like using the wrong tool for the job.
* Qml seemed very well documented and when I discovered the work of [therecipe](https://githhub.com/therecipe/qt) I felt like I was solving alot of problems related to cross platform app development in one go.
* With go compiling the qml into the application as binary, the UI is very fast.
* The Go Material UI theme is included in QT which sped up UI development.

This is a combination of a lot of work done by alot of people around the web. I am very pleased that I have now, what I think, is a very workable framework for building apps cross platform.

## Installation

I am assuming you are a Go developer already and have Go installed with your GOPATH set, etc etc...

### Installing QT

You can install QT by following the instructions [here(look it up)](here), however in essence:

`brew install qt5`

--> look up the full process you went through

### Installing QT Go-Binding

You can install the binding by following the instructions [here(therecipe)](here), however in essence:

--> look this up

### Installing Material theme

(do they need to do this?)

### Installing this project

* `export PROJECTNAME=myproject`
* `git clone https://github.com/amlwwalker/qt-hotreloader`
* `mv qt-hotreloader $PROJECTNAME`
* `cd $PROJECTNAME`
* `qtdeploy test desktop`

At this stage, the application will build and run. It will take a bit of time.

## Using the framework

The issue I was having was that recompiling the go code each time I was changing the UI code was taking a very long time between builds. It can be sped up, however I wanted to develop my UI faster. I could have used qt-creator, however it was another IDE, I couldn't build along side my Go code easily... I just wanted to stay in one environment

So I added hotloading to the qml files. There is a configuration file that tells the application whether to run in hotloading mode. If it is run in hotloading mode, then it will continually monitor changes to its own qml files and update the user interface immediately. Makes for real time development.

My recommended approach is to build a UI using this, then decide how it needs to interface with the Go code (backend). Once you know those requirements, you can add in the listeners in the Go code and have basic interaction with the front end, proving to your self you have two way comms.

Then, move to building a console UI. The framework gives a great example of interacting with the project for a second, console based UI. This allows for rapid development of the backend business logic, with fast compile times.

Once the console UI you have built demonstrates the functionality required by the go code, you can then wire up your GUI to the same functions (as the two UI's are just top level interfaces to the backend), and save yourself a lot of compilation time.

## Steps

### Development

1. In config.json, make sure that hotloading is set to true
2. Run the application with `deploy/$(OSTYPE)/$(PROJECTNAME).app/Resources/.../$(PROJECTNAME)`
3. Navigate with your editor of choice to `qml/` - this is where all the UI files are kept.
4. For now you will be editing loader.qml. This file is configured for hotloading. loader-production.qml is for when the code is to be compiled into the application for deployment. This is necessary because the finder app runs under a different user space to you.
5. Start editing loader.qml. You will see the UI update as you save.

### Production

6. When you are ready for production, copy loader.qml to loader-production.qml, set hotloading to false in config.json and recompile. You can recompile now with `qtdeploy build desktop`. Navigate in Finder to your `deploy/$(OSTYPE)` directory and double click the application. It should start up with your changes with your changes compiled in.



