MachinePoker Bot in GO
======================

Requirements
------------
* NodeJS
* GO 1.4.2

Server
------
Install:

~~~
git clone https://github.com/HackingMondays/MachinePoker.git
cd MachinePoker
sudo npm install -g coffee-script
sudo npm install -g bower
cd src/public/
bower install
cd ../..
~~~

start with: `npm run server-open` (on Windows, you need to run this command within cygwin bash)

Server URL:
`http://localhost:8080/`

Poker table (requires Flash):
`http://localhost:8080/playViewer.html`


Bots
----
On MacOS, you need to open firewall for bot.

### GOd of Gamblers
~~~
git clone https://github.com/tischda/MachinePoker-goBot.git
cd MachinePoker-goBot
go get github.com/loganjspears/joker/hand
go build -o bin/gog .
bin/gog -port=":5000" -name="GOd of Gamblers"
~~~

Bot URL:
`http://localhost:5000/bot/gog`

### Rodribot
This is another bot written in Node.js

~~~
git clone https://github.com/reyesr/rodribot-poker.git
cd rodribot-poker
npm install
/opt/local/bin/node index.js
~~~

Bot URL:
`http://localhost:5000/bot/perfect-rodribot`


GO libraries
------------

### Poker

* https://github.com/loganjspears/joker/hand
* https://github.com/cmccabe/poker-odds

### Logging

* http://stackoverflow.com/questions/16895651/how-to-implement-level-based-logging-in-golang


GO language stuff
-----------------
* http://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go
* https://github.com/axw/gocov
* https://github.com/stretchr/testify
* http://labix.org/gocheck
* https://github.com/franela/goblin
* http://goconvey.co/
* https://www.reddit.com/r/golang/comments/2pdp15/now_test_package_supports_setup_and_teardown_look/
* https://github.com/go-ini/ini
* http://grokbase.com/t/gg/golang-nuts/14384efsqy/go-nuts-best-orm-for-golang-and-their-framework-like-revel-and-martini