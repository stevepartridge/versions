#!/usr/bin/env bash
set -e

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"

if [ -f "$BASE_DIR/_scripts/local.env" ]; then
  set -o allexport
  source "${BASE_DIR}/_scripts/local.env"
  set +o allexport
  # printenv
fi

cd $BASE_DIR

# printf "SQL Files..."
# ./_sql/build_go_file.sh
# printf "done.\n"

printf "Protos..."
./_scripts/protos.sh
printf "done.\n"

printf "Swagger Docs..."
./_scripts/swagger.sh
printf "done.\n"

echo $1
case "$1" in 
  *deps\=*)
    ./_scripts/internal_deps.sh ${1//deps=/}
    shift 
    ;;
  *deps*)
  ./_scripts/internal_deps.sh
    shift 
    ;;
esac
# if [[ "$1" == "deps" ]]; then
# fi


echo "Building..."
# go run cmd/service/*.go $@
go run $(ls -1 cmd/service/*.go | grep -v _test.go) $@