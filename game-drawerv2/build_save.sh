#!/bin/bash
# create version
ver=$(date '+%s')

# building docker image
docker build -t game-drawerv2:0.0.0.$ver -f Dockerfile .

# saving docker image
### Is this needed?
docker save -o bin/game-drawerv2-pachd.img game-drawerv2:0.0.0.$ver | gzip -c > bin/game-drawerv2-pachd.tar.gz

# updating game-drawerv2 to use docker image
sed "s/VERSION/0.0.0.$ver/g" game-drawerv2.json.tmpl > game-drawerv2.json

echo "deleting game-drawerv2 pipeline/repo"
pachctl delete-pipeline game-drawerv2

echo "create game-drawerv2 pipeline"
pachctl create-pipeline -f game-drawerv2.json
