#!/bin/bash
set -e

rm -rf .vscode

oripath=`pwd`

cd ./data/toml && rm -rf *.go && cd $oripath
cd ./external && ./clean.sh && cd $oripath
cd ./jkv && ./clean.sh && cd $oripath
cd ./n3csv && ./clean.sh && cd $oripath
cd ./n3xml && ./clean.sh && cd $oripath
cd ./n3json && ./clean.sh && cd $oripath
cd ./n3cfg && ./clean.sh && cd $oripath

# delete all binary files
find . -type f -executable -exec sh -c "file -i '{}' | grep -q 'x-executable; charset=binary'" \; -print | xargs rm -f
for f in $(find ./ -name '*.log' -or -name '*.doc'); do rm $f; done