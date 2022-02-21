#!/bin/bash

wget https://sca-downloads.s3.amazonaws.com/cli/latest/ScaResolver-linux64.tar.gz
tar -xzvf ScaResolver-linux64.tar.gz
chmod +x ScaResolver
rm -rf ScaResolver-linux64.tar.gz

go test ./... -coverprofile cover.out
