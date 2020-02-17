#!/usr/bin/env bash
dir="$(cd "$(dirname "$0")/.." || exit && pwd -P)"
if hash wslpath 2>/dev/null; then
  dir="$(wslpath -w "dir")"
fi

echo "Found docker volume at ${dir}"

docker run --rm -v "${dir}:/data" dprslt/mdpdf 'bash ./scripts/compile.sh'
