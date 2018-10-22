#!/bin/bash
# create version
ver=$(date '+%s')

# building docker image
docker build -t play-tracker:0.0.0.$ver -f Dockerfile .

# saving docker image
### Is this needed?
docker save -o bin/play-tracker-pachd.img play-tracker:0.0.0.$ver | gzip -c > bin/play-tracker-pachd.tar.gz

# updating play-tracker to use docker image
sed "s/VERSION/0.0.0.$ver/g" play-tracker.json.tmpl > play-tracker.json

echo "deleting play-tracker pipeline/repo"
pachctl delete-pipeline play-tracker

echo "create play-tracker pipeline"
pachctl create-pipeline -f play-tracker.json
