GO editor configuration
=======================

IdeaJ (TODO: review for 1.0.0)
------------------------------
https://github.com/go-lang-plugin-org/go-lang-idea-plugin/releases

### MacOS

Install Go Plugin 1.0.0.alpha#92:
    Preferences, Plugins, Install plugin from disk..., select `Go.92.zip`, Restart

Create New Project (from Welcome screen)
    Go, Next
    (o) Do not create source directory, Next
    Project SDK: Go SDK go1.4.2 darwin/amd64, Next
    Project Name: `gog`, Finish

Run, Edit configurations...
    +, Go Application
    Name: `gog-app`
    File: `/Users/daniel/src/hacking/poker/gog/randomhands.go`

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
    - add Go SDK (Go SDK go1.4.2 darwin/amd64 in /usr/local/go)
    - add in SourcePath tab:
        `/usr/local/go/src`
        `/Users/daniel/src/go/src`

File, Project Structure, Modules, Sources -> check sources marked in blue

### Windows

Different drives are not correctly detected:
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

