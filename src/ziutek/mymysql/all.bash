#!/usr/bin/env bash
p=ziutek/mymysql

go $* $p/mysql $p/native $p/thrsafe $p/autorc $p/godrv
