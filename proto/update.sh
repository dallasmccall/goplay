#!/bin/bash
protoc --proto_path=. --go_out=. schema/schema.proto
protoc --proto_path=. --go_out=. dataset/dataset.proto
#TODO - How do I setup an absolute path?
#protoc --proto_path=. --go_out=. packet.proto
