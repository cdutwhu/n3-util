#!/bin/bash

cd ./preprocess && ./clean.sh && cd ..
cd ./jkv && ./clean.sh && cd ..
cd ./common && rm *.txt && cd ..
