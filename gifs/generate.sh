#!/bin/sh
#
hcho "Generating all gifs!!"

for filename in install upgrade list remove; do
    vhs "tapes/$filename.tape"
done
