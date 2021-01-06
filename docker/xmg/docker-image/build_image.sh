#!/bin/bash

LANG=en_US.UTF-8
LANGUAGE=en_US.UTF-8

set -x

REGISTRY="127.0.0.1:5000"


BUILD_DIR=`pwd`
last_change_time=`date`
SERVER_DIR=/data/server/ting_server/src/app/yut


build_image() {
	cd "$BUILD_DIR"
	image_name="xmg:zjw"
	dockerfile=Dockerfile/xmg
	
	cp -r "$SERVER_DIR"/bin/*  ${BUILD_DIR}/BuildContext/
	docker image rm "$image_name"
	docker image rm $REGISTRY/$image_name
	
	docker build \
		-f $dockerfile \
		-t $image_name \
		--label last_change_time="$last_change_time" \
		./BuildContext
	
	docker image tag $image_name $REGISTRY/$image_name
	rm -rf $BUILD_DIR/BuildContext/*
	docker push $REGISTRY/$image_name
}

build_image
