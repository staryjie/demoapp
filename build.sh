#!/bin/bash

read -p "Please enter the release num: "  release

docker build -t staryjie/demoapp:$release . && \
    docker tag staryjie/demoapp:$release hub.p8s.cn/kubevisual/demoapp:$release && \
    docker tag staryjie/demoapp:$release registry.cn-hangzhou.aliyuncs.com/staryjie/demoapp:$release
