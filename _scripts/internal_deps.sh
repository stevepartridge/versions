#!/usr/bin/env bash
set -e

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"
CUR_DIR=$(pwd)

PKGS=go,env,service
if [ ! -z ${1+x} ]; then
  PKGS=$1
fi

function cleanPath() {
  echo $( cd "$( dirname "$1" )" && pwd )
}

rm -Rf "vendor/github.com/stevepartridge"
mkdir "vendor/github.com/stevepartridge"

cd "$BASE_DIR/.."
pwd
IFS=',' read -r -a packages <<< "$PKGS"
for pkg in "${packages[@]}"
do
  pkgPath=$(cleanPath "./$pkg")
  pkgPath="$pkgPath/$pkg"
  echo "$pkg > $pkgPath"
  echo "name: $pkg"
  # continue

  echo "pkg: $pkgPath"
  if [[ -d "$pkgPath" ]]; then
    echo "Checking for $BASE_DIR/vendor/github.com/stevepartridge/$pkg"
    if [ -d "$BASE_DIR/vendor/github.com/stevepartridge/$pkg" ]; then
      echo "Clearing $pkg vendor..."
      rm -Rf "$BASE_DIR/vendor/github.com/stevepartridge/$pkg/"
    fi
    echo "Updating vendor $pkg with local..."

    for f in $(find "$pkg" -not -path "*/vendor*" -not -path "*/.git*" -print); do

      if [[ "$f" == "$pkg" ]]; then
        mkdir "$BASE_DIR/vendor/github.com/stevepartridge/$pkg"
        continue
      fi
      if [[ "$f" == "$pkg/.DS_Store" ]]; then
        continue
      fi

      rf=$( cut -d '/' -f 2- <<< "$f" )

      if [[ -d "$f" ]]; then 
        cp -R $f "$BASE_DIR/vendor/github.com/stevepartridge/$pkg/$rf"
      else 

        rf=$( cut -d '/' -f 2- <<< "$f" )

        if [[ ! -f "$BASE_DIR/vendor/github.com/stevepartridge/$pkg/$rf" ]]; then
          cp "$pkg/$rf" "$BASE_DIR/vendor/github.com/stevepartridge/$pkg/$rf"
        fi
      fi 
    done

  fi
done

cd $CUR_DIR

echo "complete."