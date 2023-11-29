#!/bin/sh
#
echo "Generating all gifs!!"

for filename in install update list remove; do
    vhs "tapes/$filename.tape"
done
