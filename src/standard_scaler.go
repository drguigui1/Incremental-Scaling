package main

import (
	"math"
)

/*
** Standard scaler object
**
** 'dim_': dimension of the data
** 'idxSample_': index of the current sample (to know how many data have been computed)
** 'mean_': Slice of mean (for each dimension)
** 'var_': Slice of variance (for each dimension)
*/
type StandardScaler struct {
	dim_ uint
	idxSample_ uint
	mean_ []*float32
	var_ []*float32
	sumSquareDiff_ []*float32
}

func InitStandardScaler(dim uint) *StandardScaler {
	mean_ := make([]*float32, dim)
	var_ := make([]*float32, dim)
	sumSquareDiff_ := make([]*float32, dim)
	return &StandardScaler{dim, 1, mean_, var_, sumSquareDiff_}
}

func (scaler *StandardScaler) PartialFit(data *[][]float32) {
	n_samples := len((*data))
	n_features := scaler.dim_

	for j := 0; j < n_samples; j++ {
		for i := 0; i < int(n_features); i++ {
			elm := (*data)[j][i]

			// not already set
			if scaler.mean_[i] == nil && scaler.var_[i] == nil {
				// set the default mean
				scaler.mean_[i] = &elm

				// set the default sum square diff
				var defaultSumSquareDiff float32 = 0.
				scaler.sumSquareDiff_[i] = &defaultSumSquareDiff
			} else{
				// Incrementaly update mean
				newMean := *scaler.mean_[i] + (elm - *scaler.mean_[i]) / float32(scaler.idxSample_)

				// Incrementaly update sum square diff
				newSumSquare := *scaler.sumSquareDiff_[i] + (elm - *scaler.mean_[i]) * (elm - newMean)

				// set the mean after
				scaler.mean_[i] = &newMean
				scaler.sumSquareDiff_[i] = &newSumSquare

				// update the var
				newVar := *scaler.sumSquareDiff_[i] / float32(scaler.idxSample_)
				scaler.var_[i] = &newVar
			}
		}
		scaler.idxSample_++
	}
}

func (scaler *StandardScaler) Transform(data *[][]float32) *[][]float32 {
	n_samples, dim := len(*data), len((*data)[0])
	scaledData := InitSlices(n_samples, dim)

	// transform the data
	for i := 0; i < dim; i++ {
		sd := float32(math.Sqrt(float64(*scaler.var_[i])))
		for j := 0; j < n_samples; j++ {
			(*scaledData)[j][i] = ((*data)[j][i] - *scaler.mean_[i]) / sd
		}
	}

	return scaledData
}