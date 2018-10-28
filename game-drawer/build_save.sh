#!/bin/bash
# create version
ver=$(date '+%s')

# building docker image
docker build -t game-drawer:0.0.0.$ver -f Dockerfile .

# saving docker image
### Is this needed?
docker save -o bin/game-drawer-pachd.img game-drawer:0.0.0.$ver | gzip -c > bin/game-drawer-pachd.tar.gz

# updating game-drawer to use docker image
sed "s/VERSION/0.0.0.$ver/g" game-drawer.json.tmpl > game-drawer.json

echo "deleting game-drawer pipeline/repo"
pachctl delete-pipeline game-drawer

echo "create game-drawer pipeline"
pachctl create-pipeline -f game-drawer.json
