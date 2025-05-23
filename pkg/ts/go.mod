module github.com/mentat/graphc/pkg/ts

go 1.23.6

require (
	github.com/flosch/pongo2 v0.0.0-20200913210552-0d938eb266f3
	github.com/mentat/graphc/pkg/common v0.0.0-unpublished
)

require gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect

replace github.com/mentat/graphc/pkg/common v0.0.0-unpublished => ../common
