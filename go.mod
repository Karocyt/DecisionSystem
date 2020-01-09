module github.com/Karocyt/expertsystem

go 1.12

replace github.com/Karocyt/expertsystem/internal/parser => ./internal/parser

replace github.com/Karocyt/expertsystem/internal/lexer => ./internal/lexer

require (
	github.com/Karocyt/expertsystem/internal/lexer v0.0.0-00010101000000-000000000000
	github.com/Karocyt/expertsystem/internal/parser v0.0.0-00010101000000-000000000000
)
