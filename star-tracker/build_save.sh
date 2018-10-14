#!/bin/bash
# create version
ver=$(date '+%s')

# building docker image
docker build -t trelore/star-tracker:0.0.0.$ver -f Dockerfile .

# saving docker image
docker save -o bin/star-tracker.img trelore/star-tracker:0.0.0.$ver | gzip -c > bin/star-tracker.tar.gz

# updating star-tracker to use docker image
sed "s/VERSION/0.0.0.$ver/g" star-tracker.json.tmpl > star-tracker.json

pachctl delete-pipeline star-tracker

pachctl create-pipeline -f star-tracker.json
