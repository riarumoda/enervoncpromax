package models

type ProductComposition struct {
	ProductDescription string
	VitaminC           float64
	VitaminB1          float64
	VitaminB2          float64
	VitaminB6          float64
	VitaminB12         float64
	VitaminD           float64
}

const MAXDATA int = 20

type ProductData [MAXDATA]ProductComposition
