package procedures

import (
	f "Enervon-CProMax/functions"
	m "Enervon-CProMax/models"
	"fmt"
)

func DetailedHighestComposition(data m.ProductData) {
	var determinator, n int
	n = 0
	MenuDetailedHighestComposition()
	fmt.Print("Please Enter Menu: ")
	fmt.Scan(&determinator)
	fmt.Print("CProMaxDebug: data[determinator-1].ProductDescription value @ procedures/composition.go before entering for determinator: ", data[determinator-1].ProductDescription, "\n")
	for determinator != 7 {
		if determinator == 1 {
			//TODO: SOMEONE PLS FIX THIS BUGG ARRGHHHHH
			fmt.Print("CProMaxDebug: determinator, n value @ procedures/composition.go: ", determinator, n, "\n")
			idx := f.HighestCompositionVitaminC(data, n)
			fmt.Print("CProMaxDebug: idx value @ procedures/composition.go: ", idx, "\n")
			fmt.Print("CProMaxDebug: data[idx].ProductDescription value @ procedures/composition.go: ", data[idx].ProductDescription, "\n")
			fmt.Print("CProMaxDebug: data[idx+1].ProductDescription value @ procedures/composition.go: ", data[idx+1].ProductDescription, "\n")
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin C in: ", data[idx].VitaminC, "mg")
			}
		} else if determinator == 2 {
			fmt.Print("CProMaxDebug: determinator, n value @ procedures/composition.go: ", determinator, n, "\n")
			idx := f.HighestCompositionVitaminB1(data, n)
			fmt.Print("CProMaxDebug: idx value @ procedures/composition.go: ", idx, "\n")
			fmt.Print("CProMaxDebug: data[idx].ProductDescription value @ procedures/composition.go: ", data[idx].ProductDescription, "\n")
			fmt.Print("CProMaxDebug: data[idx+1].ProductDescription value @ procedures/composition.go: ", data[idx+1].ProductDescription, "\n")
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin C in: ", data[idx].VitaminB1, "mg")
			}
		} else if determinator == 3 {
			fmt.Print("CProMaxDebug: determinator, n value @ procedures/composition.go: ", determinator, n, "\n")
			idx := f.HighestCompositionVitaminB2(data, n)
			fmt.Print("CProMaxDebug: idx value @ procedures/composition.go: ", idx, "\n")
			fmt.Print("CProMaxDebug: data[idx].ProductDescription value @ procedures/composition.go: ", data[idx].ProductDescription, "\n")
			fmt.Print("CProMaxDebug: data[idx+1].ProductDescription value @ procedures/composition.go: ", data[idx+1].ProductDescription, "\n")
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin C in: ", data[idx].VitaminB2, "mg")
			}
		} else if determinator == 4 {
			fmt.Print("CProMaxDebug: determinator, n value @ procedures/composition.go: ", determinator, n, "\n")
			idx := f.HighestCompositionVitaminB6(data, n)
			fmt.Print("CProMaxDebug: idx value @ procedures/composition.go: ", idx, "\n")
			fmt.Print("CProMaxDebug: data[idx].ProductDescription value @ procedures/composition.go: ", data[idx].ProductDescription, "\n")
			fmt.Print("CProMaxDebug: data[idx+1].ProductDescription value @ procedures/composition.go: ", data[idx+1].ProductDescription, "\n")
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin C in: ", data[idx].VitaminB6, "mg")
			}
		} else if determinator == 5 {
			fmt.Print("CProMaxDebug: determinator, n value @ procedures/composition.go: ", determinator, n, "\n")
			idx := f.HighestCompositionVitaminB12(data, n)
			fmt.Print("CProMaxDebug: idx value @ procedures/composition.go: ", idx, "\n")
			fmt.Print("CProMaxDebug: data[idx].ProductDescription value @ procedures/composition.go: ", data[idx].ProductDescription, "\n")
			fmt.Print("CProMaxDebug: data[idx+1].ProductDescription value @ procedures/composition.go: ", data[idx+1].ProductDescription, "\n")
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin C in: ", data[idx].VitaminB12, "mg")
			}
		} else if determinator == 6 {
			fmt.Print("CProMaxDebug: determinator, n value @ procedures/composition.go: ", determinator, n, "\n")
			idx := f.HighestCompositionVitaminD(data, n)
			fmt.Print("CProMaxDebug: idx value @ procedures/composition.go: ", idx, "\n")
			fmt.Print("CProMaxDebug: data[idx].ProductDescription value @ procedures/composition.go: ", data[idx].ProductDescription, "\n")
			fmt.Print("CProMaxDebug: data[idx+1].ProductDescription value @ procedures/composition.go: ", data[idx+1].ProductDescription, "\n")
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin C in: ", data[idx].VitaminD, "mg")
			}
		} else {
			fmt.Print("Incorrect Menu. Please Enter Menu: ")
			fmt.Scan(&determinator)
		}
		MenuDetailedHighestComposition()
		fmt.Print("Please Enter Menu: ")
		fmt.Scan(&determinator)
	}
}


func DetailedLowestComposition(data m.ProductData) {
	var determinator, n int
	n = 0
	MenuDetailedLowestComposition()
	fmt.Print("Please Enter Menu: ")
	fmt.Scan(&determinator)
	fmt.Print("CProMaxDebug: data[determinator-1].ProductDescription value @ procedures/composition.go before entering for determinator: ", data[determinator-1].ProductDescription, "\n")
	for determinator != 7 {
		if determinator == 1 {
			fmt.Print("CProMaxDebug: determinator, n value @ procedures/composition.go: ", determinator, n, "\n")
			idx := f.LowestCompositionVitaminC(data, n)
			fmt.Print("CProMaxDebug: idx value @ procedures/composition.go: ", idx, "\n")
			fmt.Print("CProMaxDebug: data[idx].ProductDescription value @ procedures/composition.go: ", data[idx].ProductDescription, "\n")
			fmt.Print("CProMaxDebug: data[idx+1].ProductDescription value @ procedures/composition.go: ", data[idx+1].ProductDescription, "\n")
			if data[idx].ProductDescription == "" {
				fmt.Println("Empty Data")
			} else {
				fmt.Println("Product Name: ", data[idx].ProductDescription)
				fmt.Println("Amount Vitamin C in: ", data[idx].VitaminC, "mg")
			}
			/* Fix this plox
		} else if determinator == 2 {
			if f.LowestCompositionVitaminB1(data, n-1) == 0 {
				fmt.Println("Data Kosong")
			} else {
				fmt.Println(f.LowestCompositionVitaminB1(data, n-1))
			}
		} else if determinator == 3 {
			if f.LowestCompositionVitaminB2(data, n-1) == 0 {
				fmt.Println("Data Kosong")
			} else {
				fmt.Println(f.LowestCompositionVitaminB2(data, n-1))
			}
		} else if determinator == 4 {
			if f.LowestCompositionVitaminB6(data, n-1) == 0 {
				fmt.Println("Data Kosong")
			} else {
				fmt.Println(f.LowestCompositionVitaminB6(data, n-1))
			}
		} else if determinator == 5 {
			if f.LowestCompositionVitaminB12(data, n-1) == 0 {
				fmt.Println("Data Kosong")
			} else {
				fmt.Println(f.LowestCompositionVitaminB12(data, n-1))
			}
		} else if determinator == 6 {
			if f.LowestCompositionVitaminD(data, n-1) == 0 {
				fmt.Println("Data Kosong")
			} else {
				fmt.Println(f.LowestCompositionVitaminD(data, n-1))
			} */
		} else {
			fmt.Print("Incorrect Menu. Please Enter Menu: ")
			fmt.Scan(&determinator)
		}
		MenuDetailedLowestComposition()
		fmt.Print("Please Enter Menu: ")
		fmt.Scan(&determinator)
	}
}
