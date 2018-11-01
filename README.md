# Spring Initializr Hijack

Repository has demo grade code to hijack [Spring Initializr](https://start.spring.io)'s Form call to its backend. 

When a user presses on the "Generate Project" button on Spring Initializr, get action is submitted to the backend i.e. *https://start.spring.io/starter.zip* with User selected criteria passed on as http parameteres to starter.zip url. 

This project has tooling to capture user input and run a concourse pipeline with it.


## Repo Contents

Below is the folder structure of the Repository and its contents

- ***infra*** : Terraform specs to setup concourse on GCP. Currently only setups a CoreOS machine on GCP. You can download `docker-compose` binary, `wget https://raw.githubusercontent.com/concourse/concourse-docker/master/docker-compose.yml`  and execute `docker-compose up -d` to bring up Concourse
- ***pipeline*** : Concourse pipeline that gets deployed
- ***templates*** : Template folder gets coppied into GitHub Repo that's created
- ***tools/github_create_repository*** : Is a small golang tool that can create an empty GitHub repository. This will be moved off to; `https://github.com/alekssaul/github-create-repository`
- ***tools/router*** : Small go app that is intended to hijack the call Spring Initializr's web UI makes from the Forum
- ***tools/webhook*** : router calls webhook to run a small shell script to push the pipeline to Concourse. You can see the concourse_sp.sh script in its folder
- ***vendor*** : Go dependencies



## Environmental Variables

```sh
DEMOAPP-INITIALIZR-GITHUBTOKEN="YOURGITHUBTOKEN"
```
