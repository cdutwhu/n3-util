#!/bin/bash

cd ./preprocess && ./clean.sh && cd ..
cd ./jkv && ./clean.sh && cd ..
cd ./common && rm -f *.txt *.log && cd ..
cd ./n3csv && ./clean.sh && cd ..
cd ./n3xml && ./clean.sh && cd ..
cd ./n3json && ./clean.sh && cd ..