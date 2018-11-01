#!/bin/bash
# create version
ver=$(date '+%s')

# building docker image
docker build -t game-drawerv2:0.0.0.$ver -f Dockerfile .

# saving docker image
### Is this needed?
docker save -o bin/game-drawerv2-pachd.img game-drawerv2:0.0.0.$ver | gzip -c > bin/game-drawerv2-pachd.tar.gz

# updating game-drawerv2 to use docker image
sed "s/VERSION/0.0.0.$ver/g" game-drawerv2.jsonnet.tmpl > game-drawerv2.jsonnet

mkdir pipelines
jsonnet -m . game-drawerv2.jsonnet
