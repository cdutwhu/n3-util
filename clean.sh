#!/bin/bash

rm -rf .vscode
cd ./external && ./clean.sh && cd -
cd ./jkv && ./clean.sh && cd -
cd ./common && rm -f *.txt *.log *.csv && cd -
cd ./n3csv && ./clean.sh && cd -
cd ./n3xml && ./clean.sh && cd -
cd ./n3json && ./clean.sh && cd -
cd ./n3toml && ./clean.sh && cd -

# delete all binary files
find . -type f -executable -exec sh -c "file -i '{}' | grep -q 'x-executable; charset=binary'" \; -print | xargs rm -f
for f in $(find ./ -name '*.log' -or -name '*.doc'); do rm $f; done