#!/bin/bash
# create version
ver=$(date '+%s')

# building docker image
docker build -t highlight-generator:0.0.0.$ver -f Dockerfile .

# saving docker image
### Is this needed?
docker save -o bin/highlight-generator-pachd.img highlight-generator:0.0.0.$ver | gzip -c > bin/highlight-generator-pachd.tar.gz

# updating highlight-generator to use docker image
sed "s/VERSION/0.0.0.$ver/g" highlight-generator.jsonnet.tmpl > highlight-generator.jsonnet

rm -rf pipelines
mkdir pipelines
jsonnet -m . highlight-generator.jsonnet
