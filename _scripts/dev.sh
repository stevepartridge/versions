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

# if [[ "$1" == "deps" ]]; then
#   ./_scripts/internal_deps.sh
# fi

echo "Building..."
go run cmd/service/*.go $@