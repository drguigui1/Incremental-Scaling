package main

type MinMaxScaler struct {
	dim uint
	mins []*float32
	maxs []*float32
}

/*
** Init the MinMaxScaler object
*/
func InitMinMaxScaler(dim uint) *MinMaxScaler {
	mins := make([]*float32, dim)
	maxs := make([]*float32, dim)
	return &MinMaxScaler{dim, mins, maxs}
}

/*
** Partial Fit on specific batch of data
** Find the mins and maxs for the current batch
*/
func (scaler *MinMaxScaler) PartialFit(data *[][]float32) {
	dim := len((*data)[0])
	n_sample := len((*data))

	for i := 0; i < dim; i++ {
		curr_min := (*data)[0][i]
		curr_max := (*data)[0][i]
		for j := 1; j < n_sample; j++ {
			elm := (*data)[j][i]
			if elm < curr_min {
				// new min
				curr_min = elm
			}
			if elm > curr_max {
				// new max
				curr_max = elm
			}
		}
		// set the new min and the new max
		if scaler.mins[i] == nil || curr_min < *scaler.mins[i] {
			scaler.mins[i] = &curr_min
		}
		if scaler.maxs[i] == nil || curr_max > *scaler.maxs[i] {
			scaler.maxs[i] = &curr_max
		}
	}
}

/*
** Transform the data using he preconputed mins and maxs of the scaler
*/
func (scaler *MinMaxScaler) Transform(data *[][]float32) *[][]float32 {
	n_samples, dim := len(*data), len((*data)[0])
	scaledData := InitSlices(n_samples, dim)

	for i := 0; i < dim; i++ {
		for j := 0; j < n_samples; j++ {
			(*scaledData)[j][i] = ((*data)[j][i] - *scaler.mins[i]) / (*scaler.maxs[i] - *scaler.mins[i])
		}
	}

	return scaledData
}