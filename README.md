Hacking Monday : Poker Robot
============================

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
    Project Name: `gog`
    , Finish

Run, Edit configurations...
    +, Go Application
    Name: `gog`
    Script: `/Users/daniel/src/hacking/poker/gog/randomhands.go`
    Working directory: `/Users/daniel/src/hacking/poker/gog`
    [x] Build before run, Output directory: `/Users/daniel/src/hacking/poker/gog/bin`

File, Project Structure, Platform Settings, SDKs
https://github.com/go-lang-plugin-org/go-lang-idea-plugin/issues/1146
    - add Go SDK (Go SDK go1.4.1 darwin/amd64 in /usr/local/go)
    - add in ClassPath tab:
        `/usr/local/go/pkg/darwin_amd64`
        `/usr/local/go/pkg`
        `/Users/daniel/src/hacking/poker/gog`
        `/usr/local/go/src`
    - add in SourcePath tab:
        `/usr/local/go/src`
        `/Users/Daniel/src/hacking/poker/gog`
        `/usr/local/go/pkg/darwin_amd64`
        `/Users/daniel/src/go/src`

Fix warnings by executing:
https://github.com/go-lang-plugin-org/go-lang-idea-plugin/issues/318#issuecomment-31303939

~~~
launchctl setenv GOROOT $GOROOT
launchctl setenv GOPATH $GOPATH
~~~

Restart IdeaJ





File extension .go would detect plugin, but Go Plugin 0.9.15.3 has issue GOPATH
not detected.

This does not work (not sure anymore, message dissapeard after running `launchctl` from command line):



Install 0.9.16-alpha.9 instead:
https://github.com/go-lang-plugin-org/go-lang-idea-plugin/releases

Preferences, Plugins, Install plugin from disk...
`google-go-language.jar`

Create Project, Go
File, Project Structure, Platform Settings, SDKs
    - remove Java
    - add Go SDK (Go SDK go1.4.1 darwin/amd64 in /usr/local/go)
    - add in ClassPath tab: `/usr/local/go/pkg`, `/Users/Daniel/src/go/pkg`
    - add in SourcePath tab: `/usr/local/go/src`, /Users/Daniel/src/go/src`
      (see: https://github.com/go-lang-plugin-org/go-lang-idea-plugin/issues/1146)
File, Project Structure, Project Settings, Project
    - select project SDK

Run, Edit Configuration
    Go Application
        Environment variables: `GOPATH=/Users/daniel/src/go`
        (o) Go file: `/Users/daniel/src/hacking/poker/gog/test.go`
        [x] Build before run


OK, you need to create a module in IDEAJ project for this to work. 0.9.16-alpha.9 obliges you
to create src/pkg directories but we don't want this. 0.9.15.3 seems to be a better option
here, but you need to put main.go in main package and folder. Same as joker project.

Still need to figure out if we need launchctl and what is needed in classpath and sourcepath.

OK, I can't compile anything from IDEAJ with this setup...

