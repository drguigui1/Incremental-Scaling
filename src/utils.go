package main

import (
	"io/ioutil"
	"encoding/csv"
	"os"
	"strings"
    "path/filepath"
    "strconv"
    "fmt"
	"math"
)

/*
** Read csv
**
** param:
** 'csvPath': path of the csv
**
** return:
** Data structure containing the csv data
*/
func FromCSV(csvPath string) [][]string {
    // read the file
    data, err := ioutil.ReadFile(csvPath)
    if err != nil {
        panic(err)
    }

    r := csv.NewReader(strings.NewReader(string(data)))
    records, err := r.ReadAll()
    if err != nil {
        panic(err)
    }

    return records
}

/*
** Save the data into csv
**
** params:
** 'data': data structure containing the data to save
** 'csv_path': path to the csv
*/
func ToCSV(data [][]string, csv_path string) {
    // create a new file
    file, err := os.Create(csv_path)
    if err != nil {
        panic(err)
    }

    // create a new writer
    w := csv.NewWriter(file)

    // write data
    w.WriteAll(data)
    file.Close()
}

/*
** Init Slice of slice
*/
func InitSlices(dim1, dim2 int) *[][]float32 {
	newSlices := make([][]float32, dim1)
	for i := 0; i < dim1; i++ {
		newSlices[i] = make([]float32, dim2)
	}
	return &newSlices
}

/*
** Init Slice of slice string
*/
func InitSlicesStr(dim1, dim2 int) *[][]string {
	newSlices := make([][]string, dim1)
	for i := 0; i < dim1; i++ {
		newSlices[i] = make([]string, dim2)
	}
	return &newSlices
}

/*
** From string slices to float32
*/
func ToFloatSlices(slice *[][]string) *[][]float32 {
    n_samples := len(*slice)
    dim := len((*slice)[0])
    res := InitSlices(n_samples, dim)

    for i := 0; i < n_samples; i++ {
        for j := 0; j < dim; j++ {
            tmp, err := strconv.ParseFloat((*slice)[i][j], 32)

            if err != nil {
                panic(err)
            }

            (*res)[i][j] = float32(tmp)
        }
    }
    return res
}

/*
** From Float slices to string
*/
func ToStringSlices(slice *[][]float32) *[][]string {
    n_samples := len(*slice)
    dim := len((*slice)[0])
    res := InitSlicesStr(n_samples, dim)

    for i := 0; i < n_samples; i++ {
        for j := 0; j < dim; j++ {
            (*res)[i][j] = fmt.Sprintf("%f", (*slice)[i][j])
        }
    }
    return res
}

/*
** List all files from specific dir
*/
func ListFiles(path string) ([]string, error) {
    var files []string
    err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
        if !info.IsDir() {
            files = append(files, p)
        }
        return nil
    })
    return files, err
}

func RoundAllElm(slice *[][]float32) {
	n, m := len(*slice), len((*slice)[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
            x := (*slice)[i][j]
            (*slice)[i][j] = float32(math.Round(float64(x*1000))/1000.)
		}
	}
}