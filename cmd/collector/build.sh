#!/bin/bash
rm -f collector
go clean -i -cache
go build
