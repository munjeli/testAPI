#!/bin/bash

mkdir package/testAPI
cp testAPI package/testAPI/
cp config/* package/testAPI/
cd package || exit
zip -m -r testAPI.zip testAPI