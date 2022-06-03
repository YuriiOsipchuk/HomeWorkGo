#!/bin/bash

FILENAME="main.go"
PROGRAMNAME="main"

echo "***** START Test: **********";
if [[ -f $FILENAME ]]; then

	go build main.go

	if [[ -f $PROGRAMNAME ]]; then
		echo {1..10} | ./main
	else
		exit 1
	fi

else
	echo "File $FILENAME not found"
	exit 1
fi

echo "***** END Test:   **********";
