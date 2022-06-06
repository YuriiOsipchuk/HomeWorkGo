#!/bin/bash

FILENAME="main.go"
PROGRAMNAME="main"

echo "***** START Test: **********";
if [[ -f $FILENAME ]]; then

  go build main.go

  if [[ -f $PROGRAMNAME ]]; then
    echo {1..150} | ./main
  else
    exit 1
  fi

else
  echo "File $FILENAME not found"
  exit 1
fi

if [[ -f $PROGRAMNAME ]]; then
  rm $PROGRAMNAME
fi
echo "***** END Test:   **********";
