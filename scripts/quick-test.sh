#!/bin/sh

# https://stackoverflow.com/a/246128/210827
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
GOCEAN_PATH=$SCRIPT_DIR/../gocean

$GOCEAN_PATH coops air_gap -station 8545556 -count 1
sleep 1

$GOCEAN_PATH coops air_pressure -station 9410230 -count 1
sleep 1

$GOCEAN_PATH coops air_temperature --station 9410230 --count 1
sleep 1

$GOCEAN_PATH coops conductivity -station 8447386 -count 1
sleep 1

$GOCEAN_PATH coops visibility -station 8447412 -count 1
sleep 1

$GOCEAN_PATH coops predictions -station 9414523 -count 1
sleep 1

$GOCEAN_PATH coops water_level -station 9410230 -count 1
sleep 1

$GOCEAN_PATH coops water_temperature --station 9410230 --count 1
sleep 1

$GOCEAN_PATH coops wind -station 9410230 -count 1
sleep 1
