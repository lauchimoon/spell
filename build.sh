#!/usr/bin/sh
wget https://raw.githubusercontent.com/dwyl/english-words/master/words.txt
go build -o spell spell.go
