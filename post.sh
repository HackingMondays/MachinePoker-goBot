#!/bin/bash

# http://stackoverflow.com/questions/7172784/how-to-post-json-data-with-curl-from-terminal-commandline-to-test-spring-rest

curl http://localhost:8081/bot/gog

curl -H "Content-Type: application/json" --data @data/data.json http://localhost:8081/bot/gog

