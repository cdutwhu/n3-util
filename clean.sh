#!/bin/bash

cd ./preprocess && ./clean.sh && cd ..
cd ./jkv && ./clean.sh && cd ..
cd ./common && rm -f *.txt *.log && cd ..
cd ./csv && ./clean.sh && cd ..
cd ./xml && ./clean.sh && cd ..