#!/bin/bash

grpc_cli ls localhost:8180
grpc_cli type localhost:8180 pb.PostAccountRequest
grpc_cli call localhost:8180 AccountService.PostAccount 'name: "inu"'
grpc_cli call localhost:8180 AccountService.GetAccount 'id: "172r8K9zSq5pgTewuyYjuOmjaQm"'
