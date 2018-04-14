#!/usr/bin/env bash

package_name="lhman"

platforms=("windows/amd64" "windows/386" "darwin/amd64" "darwin/386" "linux/386" "linux/amd64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$package_name'-'$GOOS'-'$GOARCH

    filename=$package_name
    if [ $GOOS = "windows" ]; then
        filename+='.exe'
    fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -o bin/$output_name/$filename
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    else
        echo 'Builded: ' ${platform}
    fi
done