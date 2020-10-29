#!/bin/bash
# create version
ver=$(date '+%s')

# building docker image
docker build -t player-stats-gamelog:0.0.0.$ver -f Dockerfile .

# saving docker image
### Is this needed?
docker save -o bin/player-stats-gamelog-pachd.img player-stats-gamelog:0.0.0.$ver | gzip -c > bin/player-stats-gamelog-pachd.tar.gz

# updating player-stats-gamelog to use docker image
sed "s/VERSION/0.0.0.$ver/g" player-stats-gamelog.json.tmpl > player-stats-gamelog.json

echo "deleting player-stats-gamelog pipeline/repo"
pachctl delete pipeline player-stats-gamelog

echo "create player-stats-gamelog pipeline"
pachctl create pipeline -f player-stats-gamelog.json
