#!/bin/sh

go build -o mysql_plugin/mysql.plugin mysql_plugin/mysql_plugin.go
go run client/client.go $*
