#!/usr/bin/env bash
# https://github.com/hashicorp/vault/blob/master/scripts/gofmtcheck.sh

echo "==> Checking that code complies with gofmt requirements..."

gofmt_files=$(gofmt -l -s $1)
if [[ -n ${gofmt_files} ]]; then
    echo 'gofmt needs running on the following files:'
    echo " ===== "
    echo "${gofmt_files}"
    echo " ===== "
    echo "Please, use \`make fmt\` to reformat code."
    exit 1
fi
