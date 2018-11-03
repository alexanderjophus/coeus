# Project Coeus

This is a collection of services designed to work with [stats-importer](https://github.com/alexanderjosephtrelore/stats-importer).
The idea being stats-importer collates information from [NHL stats](https://gitlab.com/dword4/nhlapi/blob/master/stats-api.md), and imports them into pachyderm.
These services then reads off that pachyderm repo and transform/manipulate that data in multiple ways.

## Current goals

- Explore Markov Chains using highlight descriptions
- Have a chain of services (currently `play-tracker` feeds into `game-drawer`, but this is being deleted)
- Explore machine learning within pachyderm
- More visualization of hockey data (geomap of teams origins?)

## Running the services

### Prerequisites

Assuming minikube is running locally with [pachyderm](https://pachyderm.readthedocs.io/en/latest/) deployed on it.
Assuming pachctl is installed and there is a statspi repo (possibly populated by stats-importer project).

This will set env variables allowing you build directly on the minikube docker daemon

```
eval $(minikube docker-env)
```

This will allow you to connect to pachyderm on your local machine.

```
export ADDRESS=$(minikube ip):30650
pachctl port-forward &
```

To check this is all working `pachctl list-repo` should show you the statsapi repo.

### run

Each service should have a make file (or documentation explaining otherwise), type `make`, the service should build, deploy docker image, deploy pipeline all automatically.
