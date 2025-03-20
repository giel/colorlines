#!/usr/bin/env bash

function ShowInfo {
	echo "--------------"
	echo "--" $1
	echo "--------------"
}

ShowInfo "Check directory"
if [[ "$(basename "$PWD")" == "colorlines" ]]; then
	echo "Running in the 'colorlines' directory."
else
	echo "Not running in the 'colorlines' directory."
	cd ~/git/colorlines
fi

ShowInfo "Go fmt"
go fmt

if [[ -z $GITHUB_ACTIONS ]]; then
	version=$(git describe)
	# version=1.2.3
fi

ShowInfo "Version $version"

ShowInfo "Build"
