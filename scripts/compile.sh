#!/usr/bin/env bash
outputDirectory="$(cd "$(dirname "$0")/.." || exit && pwd -P)/output"
if hash wslpath 2>/dev/null; then
  outputDirectory="$(wslpath -w "$outputDirectory")"
fi
echo "Moved to project root outputDirectory: $(pwd)"

if [[ -d ./output ]]; then
  rm -rf ./output
  echo "Deleted previous output directory"
fi
mkdir output

cat >./output/compile-script.sh <<EOL
for file in *.md; do
    [ -f "\$file" ] || continue
    mdpdf "\$file"
done
EOL
chmod u+x ./output/compile-script.sh
echo "Generated compile script"

#We use depth 3 here because README.md is in depth 1, and the template directory is depth 2
find . -mindepth 3 -name "*.md" -print0 | while read -r -d $'\0' file; do
  cleanFileName=${file:2:-3}
  pdfOutput="${cleanFileName//\//-}.md"

  cp "$file" "./output/$pdfOutput"
done
echo "Compied files"

if which mdpdf &> /dev/null; then
  pushd output || exit 1
  ./compile-script.sh
  popd || exit 1
elif which docker &> /dev/null; then
  docker run --rm -v "$outputDirectory:/data" enric1994/mdpdf /bin/sh -c 'cd /data; ./compile-script.sh'
else
  echo "Did not find docker or mdpdf binary!!"
  exit 1
fi
echo "Compiled pdfs"

for file in ./output/*.md; do
  rm "$file"
done
rm ./output/compile-script.sh
echo "Cleaned output directory"
