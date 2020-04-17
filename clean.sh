#!/bin/bash

cd ./preprocess && ./clean.sh && cd ..
cd ./jkv && ./clean.sh && cd ..
cd ./common && rm -f *.txt *.log && cd ..
cd ./n3-csv && ./clean.sh && cd ..
cd ./n3-xml && ./clean.sh && cd ..
cd ./n3-json && ./clean.sh && cd ..