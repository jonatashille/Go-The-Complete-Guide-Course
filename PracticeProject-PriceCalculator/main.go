package main

import (
	"Calculator/filemanager"
	"Calculator/prices"
	"fmt"
)

func main() {
	var taxRates []float64 = []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("pricess.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}

}
