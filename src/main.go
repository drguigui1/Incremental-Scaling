package main

import (
	"fmt"
	//"strconv"
)

func main() {
	/*path := "./dataset/data/dataset_usefull_batch_nb_"
	saved_path := "./dataset/scaled_data/dataset_scaled_nb_"

	// number of file
	nb_files := 21

	// create the min max model
	// to change (get the dim dynamicaly)
	var dim uint = 93
	scaler := InitMinMaxScaler(dim)

	// fit the model
	for idx := 1; idx <= nb_files; idx++ {
		// load the file
		data_str := FromCSV(path + strconv.Itoa(idx))
		data := ToFloatSlices(&data_str)

		fmt.Println("Fitting batch nb ", idx)
		// partial fit model
		scaler.PartialFit(data)
	}

	// transform the data
	for idx := 1; idx <= nb_files; idx++ {
		fmt.Println("Saving batch nb ", idx)

		// load the file
		data_str := FromCSV(path + strconv.Itoa(idx))
		data := ToFloatSlices(&data_str)

		// transform the data using the scaler
		scaled_data := scaler.Transform(data)

		// convert scaled data into string slice
		scaled_data_str := ToStringSlices(scaled_data)

		// save the transformed data
		ToCSV(*scaled_data_str, saved_path + strconv.Itoa(idx))
	}
	*/
	test1 := [][]float32{
		{33., 12.},
		{1., 11.},
	}
	test2 := [][]float32{
		{1., 0.},
		{4., 1.},
	}

	/*ref1 := [][]float32{
		{-1., -1.},
		{-1., -1.},
	}

	ref2 := [][]float32{
		{1., 1.},
		{1., 1.},
	}*/

	scaler := InitStandardScaler(2)

	scaler.PartialFit(&test1)
	scaler.PartialFit(&test2)

	res1 := scaler.Transform(&test1)
	res2 := scaler.Transform(&test2)

	RoundAllElm(res1)
	RoundAllElm(res1)

	for _, elm := range scaler.mean_ {
		fmt.Printf("%v ", *elm)
	}
	fmt.Printf("\n")
	for _, elm := range scaler.var_ {
		fmt.Printf("%v ", *elm)
	}
	fmt.Printf("\n")

	fmt.Println(res1)
	fmt.Println(res2)
}