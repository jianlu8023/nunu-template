#!/bin/bash


pwd=$(pwd)

clean(){
  if [[ -d "${pwd}/../bin" ]]; then
    rm -rf "${pwd}/../bin"

  fi
  echo "clean binary success"
}

clean