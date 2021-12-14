module github.com/test/case1

require (
	github.com/scott/sayhello v0.0.0
)
replace (
	github.com/scott/sayhello v0.0.0 => ../module1
)
go 1.16
