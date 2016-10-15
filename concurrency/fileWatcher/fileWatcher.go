//This package watches for csv files in a specified path.
//When file is added, it reads the csv prints the data and then removes the file.
package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const watchedPath = "./source"

type Invoice struct {
	Number              string
	Amount              float64
	PurchaseOrderNumber int
	InvoiceDate         time.Time
}

func main() {
	runtime.GOMAXPROCS(4)
	for {
		d, _ := os.Open(watchedPath)
		//fmt.Println(a)
		files, _ := d.Readdir(-1)
		//fmt.Println(files)
		for _, fi := range files {
			filePath := watchedPath + "/" + fi.Name()
			f, _ := os.Open(filePath)
			data, _ := ioutil.ReadAll(f)
			f.Close()
			os.Remove(filePath)
			go func(data string) {
				reader := csv.NewReader(strings.NewReader(data))
				records, _ := reader.ReadAll()
				for _, r := range records {
					invoice := new(Invoice)
					invoice.Number = r[0]
					invoice.Amount, _ = strconv.ParseFloat(r[1], 64)
					invoice.PurchaseOrderNumber, _ = strconv.Atoi(r[2])
					unixTime, _ := strconv.ParseInt(r[3], 10, 64)
					invoice.InvoiceDate = time.Unix(unixTime, 0)
					fmt.Printf("Received Invoice '%v' for $%.2f and submitted for processing\n", invoice.Number, invoice.Amount)
				}
			}(string(data))
			d.Close()
			time.Sleep(100 * time.Millisecond)
		}

	}

}
