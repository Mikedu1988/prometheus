#!/bin/sh
cd ..
docker build -t duliyang/testmetric:0.1.0 .
docker push duliyang/testmetric:0.1.0