GO editor configuration
=======================

IdeaJ
-----
https://github.com/go-lang-plugin-org/go-lang-idea-plugin/releases

### MacOS

Install Go Plugin 1.0.0.alpha#134:
    Preferences, Plugins, Install plugin from disk..., select `Go.134.zip`, Restart

Create New Project (from Welcome screen)
    Go, Next
    (o) Do not create source directory, Next
    Project SDK: Go SDK go1.4.2 darwin/amd64, Next
    Project Name: `MachinePoker-goBot`, Finish

Run, Edit configurations...
    +, Go Application
    Name: `gog-app`
    File: `<full-path-to>/MachinePoker-goBot/main.go`
    Before launch: Make

File, Project Structure, Platform Settings, SDKs
    - add Go SDK (Go SDK go1.4.2 darwin/amd64 in /usr/local/go)
    - add in SourcePath tab:
        `/usr/local/go/src`

File, Project Structure, Modules, Sources -> check sources marked in blue

### Windows

Drives other than c: are not detected/used correctly:
https://github.com/go-lang-plugin-org/go-lang-idea-plugin/issues/1240

Workaround:

~~~
c:
cd \
mklink /d src u:\src
~~~


Sublime Text
------------

Tools, Command Palette, Package Control:Install Package, GoSublime
https://github.com/DisposaBoy/GoSublime
Problem -> Key-bindings fail on FR keyboard ! (super+.)

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

