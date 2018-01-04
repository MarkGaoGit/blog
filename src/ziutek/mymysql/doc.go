// Package mymysql provides MySQL client API and database/sql driver.
//
// It can be used as a library or as a database/sql driver.
//
// Using as a library
//
// Import native or thrsafe engine. Optionally import autorc for autoreconnect connections.
//
//	import (
//		"ziutek/mymysql/mysql"
//		_ "ziutek/mymysql/thrsafe" // OR native
//		// _ "ziutek/mymysql/native"
//		"ziutek/mymysql/autorc" // for autoreconnect
//	)
//
//
//
// Using as a Go sql driver
//
// Import Go standard sql package and godrv driver.
//
//	import (
//		"database/sql"
//		_ "ziutek/mymysql/godrv"
//	)
//
//
//
package mymysql
