package dtos

type ValidatorType string

const (
	RowValidator            ValidatorType = "RowValidator"
	ColumnValidator         ValidatorType = "ColumnValidator"
	FirstDiagonalValidator  ValidatorType = "FirstDiagonalValidator"
	SecondDiagonalValidator ValidatorType = "SecondDiagonalValidator"
)
