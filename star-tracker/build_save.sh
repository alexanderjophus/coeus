#!/bin/bash
# create version
ver=$(date '+%s')

# building docker image
docker build -t star-tracker:0.0.0.$ver -f Dockerfile .

# saving docker image
### Is this needed?
docker save -o bin/star-tracker-pachd.img star-tracker:0.0.0.$ver | gzip -c > bin/star-tracker-pachd.tar.gz

# updating star-tracker to use docker image
sed "s/VERSION/0.0.0.$ver/g" star-tracker.json.tmpl > star-tracker.json

echo "deleting star-tracker pipeline/repo"
pachctl delete-pipeline star-tracker

echo "create star=tracker pipeline"
pachctl create-pipeline -f star-tracker.json
