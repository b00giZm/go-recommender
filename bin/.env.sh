#!/usr/bin/env bash

export DOCKER_IMAGE_BEANSTALKD="dr.chefkoch.net/benutzer/beanstalkd"
export DOCKER_IMAGE_RECOMMENDER="dr.chefkoch.net/benutzer/recommender"

export DOCKER_CONTAINER_BEANSTALKD="beanstalkd"

export DOCKER_HOST_BEANSTALKED="${DOCKER_CONTAINER_BEANSTALKD}.local.dev.chefkoch.de"