module geeorm/session

go 1.17

replace geeorm/log => ../log

require (
	geeorm/clause v0.0.0-00010101000000-000000000000
	geeorm/dialect v0.0.0-00010101000000-000000000000
	geeorm/log v0.0.0-00010101000000-000000000000
	geeorm/schema v0.0.0-00010101000000-000000000000
)

replace geeorm/dialect => ../dialect

replace geeorm/schema => ../schema

replace geeorm/clause => ../clause
