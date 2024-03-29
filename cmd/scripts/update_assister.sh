#!/bin/bash

GITHUB_API_REST="pro-assistance/pro-assister"
last_tag=$(curl -s GET https://api.github.com/repos/pro-assistance/pro-assister/tags | jq -r '.[].name' | head -n1)

go get github.com/pro-assistance/pro-assister@${last_tag}
