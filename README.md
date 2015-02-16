Poker Bot in GO
===============

Requirements
------------
NodeJS
GO 1.4.1

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
~~~

start with: `npm run server-open`

Server URL:
`http://localhost:8080/`

Poker table (requires Flash):
`http://localhost:8080/playViewer.html`


Bots
----
On MacOS, you need to open firewall for bot.

### Rodribot
~~~
git clone https://github.com/reyesr/rodribot-poker.git
cd rodribot-poker
npm install
/opt/local/bin/node index.js
~~~

Bot URL:
`http://localhost:5000/bot/perfect-rodribot`


### GOd of Gamblers
~~~
git clone https://github.com/tischda/gog.git
cd gog
go build -o bin/gog .
bin/gog
~~~

Bot URL:
`http://localhost:8081/bot/gog`


Poker GO libraries
------------------
* https://github.com/loganjspears/joker/blob/master/hand/hand.go
* https://github.com/cmccabe/poker-odds
