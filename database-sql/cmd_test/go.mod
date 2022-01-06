module github.com/NDzuki/geeklearn/database-sql/cmd_test

go 1.17

replace geeorm => ../geeorm

replace geeorm/log => ../log

replace geeorm/session => ../session

require (
	geeorm v0.0.0-00010101000000-000000000000
	geeorm/log v0.0.0-00010101000000-000000000000
	github.com/mattn/go-sqlite3 v1.14.10
)

require geeorm/session v0.0.0-00010101000000-000000000000 // indirect
