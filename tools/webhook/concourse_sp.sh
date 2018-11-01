#!/bin/sh

export reponame=$1
echo "$BIGSECRET" > credentials.yml
echo PAS_APP_SUFFIX: $reponame >> credentials.yml
echo PAS_SPACE: $reponame >> credentials.yml
echo REPONAME: $reponame  >> credentials.yml
cat credentials.yml

alias fly="/app/fly"

CONCOURSE_USERNAME=$(cat credentials.yml | grep CONCOURSE_USERNAME | awk 'BEGIN { FS = ":" }; { print $2 }' | tr -d " ")
CONCOURSE_PASSWORD=$(cat credentials.yml | grep CONCOURSE_PASSWORD | awk 'BEGIN { FS = ":" }; { print $2 }' | tr -d " " | tr -d "\"")
CONCOURSEURL=$(cat credentials.yml | grep CONCOURSEURL | awk 'BEGIN { FS = ": " }; { print $2 }'| tr -d " " | tr -d "\"")
fly login -t concourse -c $CONCOURSEURL -k -u $CONCOURSE_USERNAME -p $CONCOURSE_PASSWORD
export pipelinename=$reponame-$RANDOM
fly -t concourse set-pipeline -p $pipelinename -c pipeline.yml -l credentials.yml -n
fly -t concourse unpause-pipeline -p $pipelinename
