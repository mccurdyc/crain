#!/bin/sh
# Generate test coverage statistics for Go packages.
#
# Works around the fact that `go test -coverprofile` currently does not work
# with multiple packages, see https://code.google.com/p/go/issues/detail?id=6909
#
# Usage: script/coverage [--html|--coveralls]
#
#     --html      Additionally create HTML report and open it in browser
#     --coveralls Push coverage statistics to coveralls.io
#
# credit: https://github.com/hashicorp/vault/blob/master/scripts/coverage.sh

set -e

workdir=.cover
profile="$workdir/cover.out"
mode=count

generate_cover_data() {
    rm -rf "$workdir"
    mkdir "$workdir"

    for pkg in "${@}"; do
      echo "===> Generating coverage for: $pkg"
        f="$workdir/$(echo $pkg | tr / -).cover"
        go test -covermode="$mode" -coverprofile="$f" "$pkg"
    done

    echo "mode: $mode" >"$profile"
    grep -h -v "^mode:" "$workdir"/*.cover >>"$profile"
}

case "$1" in
-html)
    generate_cover_data ${@:2}
    go tool cover -html="$profile"
    ;;
-push)
    echo >&2 "no test coverage tool connected"
    ;;
*)
    generate_cover_data $@
    go tool cover -func="$profile"
    ;;
esac
