#!/bin/bash

[[ ! -d .git ]] && [[ $AUTO_UPDATE == "1" ]] && git clone $GIT_ADDRESS . || [[ -d .git ]] && [[ $AUTO_UPDATE == "1" ]] && git pull
[[ ! -z $NODE_PACKAGES ]] && /usr/local/bin/npm install $NODE_PACKAGES
[[ ! -z $UNNODE_PACKAGES ]] && /usr/local/bin/npm uninstall $UNNODE_PACKAGES
[[ -f /home/container/package.json ]] && /usr/local/bin/npm install
[[ $MAIN_FILE == "*.js" ]] && /usr/local/bin/node "/home/container/$MAIN_FILE" $NODE_ARGS || /usr/local/bin/ts-node --esm "/home/container/$MAIN_FILE" $NODE_ARGS
