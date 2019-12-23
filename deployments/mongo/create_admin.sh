#!/usr/bin/env sh

base64Psw=$(htpasswd -bnBC 12 "" back-challenge | tr -d ':\n' | sed 's/$2y/$2a/' | base64 -w 0)
insertCommand="db.users.insertOne({uuid: UUID(), name: \"admin\", password: BinData(0, \"$base64Psw\")})"
mongo --host mongo:27017 -u back -p back12345 --authenticationDatabase admin romson-back --eval "$insertCommand"