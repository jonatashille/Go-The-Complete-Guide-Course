package prices

import (
	"Calculator/conversion"
	"Calculator/filemanager"
	"fmt"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager `json:"-"`
	TaxRate           float64                 `json:"tax_rate"`
	InputPrices       []float64               `json:"input_prices"`
	TaxIncludedPrices map[string]string       `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices

	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {

	err := job.LoadData()
	//job.LoadData()

	if err != nil {
		//return err
		errorChan <- err
		return
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	job.IOManager.WriteResult(job)

	doneChan <- true
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   fm,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
