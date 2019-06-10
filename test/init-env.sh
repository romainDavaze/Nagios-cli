#!/bin/bash

readonly docker=/usr/bin/docker
readonly jq=/usr/bin/jq

readonly apiKeyQuery="SELECT api_key FROM xi_users WHERE username = 'nagiosadmin'"
readonly container="nagiosxi"
readonly image="docker.io/tgoetheyn/docker-nagiosxi"
readonly tmpFile=$(mktemp)

$docker inspect $container &> $tmpFile

if [[ $? -ne 0 ]]; then
    echo "NagiosXI container not found. Creating it..."
    $docker run -d --name $container -p 0.0.0.0:8080:80 $image > /dev/null
    # Waiting for container to be ready
    sleep 5
    # Sometimes nagiosxi does not start for some reason, restarting container to make sure it does
    docker restart $container > /dev/null
    sleep 10
    $docker inspect $container &> $tmpFile
fi

running=$($jq '.[0].State.Running' $tmpFile)

if [[ "${running}" -ne "true" ]]; then
    echo "NagiosXI container not running. Starting it..."
    $docker start $container > /dev/null
    # Waiting for container to be ready
    sleep 10
fi

export API_KEY=$($docker exec $container bash -c "mysql -s -u root -pnagiosxi -D nagiosxi -e \"${apiKeyQuery}\" | grep -v api_key")