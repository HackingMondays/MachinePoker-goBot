Hacking Monday : Poker Robot
============================

* https://github.com/loganjspears/joker/blob/master/hand/hand.go
* https://github.com/cmccabe/poker-odds

19-01-2015 : Setup GO environment
---------------------------------

### Sublime Text

Tools, Command Palette, Package Control:Install Package, GoSublime
https://github.com/DisposaBoy/GoSublime
Problem -> Key-bindings do not work on FR keyboard ! (super+.)

You need to edit `~/Library/Application Support/Sublime Text 2/Packages/GoSublime/Default (OSX).sublime-keymap`

To see what keys were typed:

View, Show Console
~~~
sublime.log_input(True)
~~~

ST2 does not find GOPATH:
https://github.com/DisposaBoy/GoSublime/issues/71
https://github.com/DisposaBoy/GoSublime/issues/513

This fixes it:
    Quit Sublime Text 2
    cd ~/Library/Application\ Support/Sublime\ Text\ 2/Packages/
    rm -rf GoSublime
    git co https://github.com/FiloSottile/GoSublime.git
    Start Sublime Text 2

Edit project file and add:

    "settings": {
        "GoSublime": {
            "env": {
                "GOPATH": "/Users/daniel/src/go"
            }
        }
    },


### IdeaJ

Install Go Plugin 0.9.15.3:
    Preferences, Plugins, Browse, search for `golang`, Install, Restart

Create New Project (from Welcome screen)
    Go, Next
    (o) Do not create source directory, Next
    Project SDK: Go SDK go1.4.1 darwin/amd64, Next
    Project Name: `gog`, Finish

Run, Edit configurations...
    +, Go Application
    Name: `gog`
    Script: `/Users/daniel/src/hacking/poker/gog/randomhands.go`
    Working directory: `/Users/daniel/src/hacking/poker/gog`
    [x] Build before run, Output directory: `/Users/daniel/src/hacking/poker/gog/bin`

Fix warnings:
https://github.com/go-lang-plugin-org/go-lang-idea-plugin/issues/318#issuecomment-31303939
~~~
launchctl setenv GOROOT $GOROOT
launchctl setenv GOPATH $GOPATH
~~~
Restart IdeaJ

Fix auto-completion:
https://github.com/go-lang-plugin-org/go-lang-idea-plugin/issues/1146
File, Project Structure, Platform Settings, SDKs
    - add Go SDK (Go SDK go1.4.1 darwin/amd64 in /usr/local/go)
    - add in SourcePath tab:
        `/usr/local/go/src`
        `/Users/daniel/src/go/src`


Server
------
Install:
~~~
cd MachinePoker
sudo npm install -g coffee-script
sudo npm install -g bower
cd src/public/
bower install
~~~


go to MachinePoker and run start.sh

Server:
http://localhost:8080/

Graphical view (requires Flash)
http://localhost:8080/playViewer.html



