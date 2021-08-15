package hampel

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunningMedianAndSigma(t *testing.T) {
	// N = 3, windowSize = 3
	medians, sigmas := runningMedianAndSigma([]float64{
		5232.200195,
		5277.899902,
		5281.799805,
	}, 3)

	require.Equal(t, []float64{
		5277.899902,
		5277.899902,
		5277.899902,
	}, medians)

	require.Equal(t, []float64{
		5.781996187799318,
		5.781996187799318,
		5.781996187799318,
	}, sigmas)

	// N = 4, windowSize = 4
	medians, sigmas = runningMedianAndSigma([]float64{
		5232.200195,
		5277.899902,
		5281.799805,
		5263.100098,
	}, 4)

	require.Equal(t, []float64{
		5270.5,
		5270.5,
		5270.5,
		5270.5,
	}, medians)

	require.Equal(t, []float64{
		13.862092799099823,
		13.862092799099823,
		13.862092799099823,
		13.862092799099823,
	}, sigmas)

	// N = 10, windowSize = 5
	medians, sigmas = runningMedianAndSigma([]float64{
		5232.200195,
		5277.899902,
		5281.799805,
		5263.100098,
		5244.750000,
		5249.399902,
		5210.399902,
		5233.950195,
		5259.899902,
		5252.200195}, 5)

	require.Equal(t, []float64{
		5263.100098,
		5263.100098,
		5263.100098,
		5263.100098,
		5249.399902,
		5244.750000,
		5244.750000,
		5249.399902,
		5249.399902,
		5249.399902,
	}, medians)

	require.Equal(t, []float64{
		27.205855294799832,
		27.205855294799832,
		27.205855294799832,
		21.942189410400328,
		20.31191058959967,
		16.011790892999482,
		16.011790892999482,
		15.567300,
		15.567300,
		15.567300,
	}, sigmas)

	// N = 11, windowSize = 5
	medians, sigmas = runningMedianAndSigma([]float64{
		5232.200195,
		5277.899902,
		5281.799805,
		5263.100098,
		5244.750000,
		5249.399902,
		5210.399902,
		5233.950195,
		5259.899902,
		5252.200195,
		5274.850098}, 5)
	require.Equal(t, []float64{
		5263.100098,
		5263.100098,
		5263.100098,
		5263.100098,
		5249.399902,
		5244.750000,
		5244.750000,
		5249.399902,
		5252.200195,
		5252.200195,
		5252.200195,
	}, medians)

	require.Equal(t, []float64{
		27.205855294799832,
		27.205855294799832,
		27.205855294799832,
		21.942189410400328,
		20.31191058959967,
		16.011790892999482,
		16.011790892999482,
		15.567299999999999,
		27.057449999999999,
		27.057449999999999,
		27.057449999999999,
	}, sigmas)

	// N = 10, windowSize = 10
	medians, sigmas = runningMedianAndSigma([]float64{
		5232.200195,
		5277.899902,
		5281.799805,
		5263.100098,
		5244.750000,
		5249.399902,
		5210.399902,
		5233.950195,
		5259.899902,
		5252.200195}, 10)

	require.Equal(t, []float64{
		5250.8000485,
		5250.8000485,
		5250.8000485,
		5250.8000485,
		5250.8000485,
		5250.8000485,
		5250.8000485,
		5250.8000485,
		5250.8000485,
		5250.8000485,
	}, medians)

	require.Equal(t, []float64{
		21.60882309389966,
		21.60882309389966,
		21.60882309389966,
		21.60882309389966,
		21.60882309389966,
		21.60882309389966,
		21.60882309389966,
		21.60882309389966,
		21.60882309389966,
		21.60882309389966,
	}, sigmas)

	// N = 11, windowSize = 10
	medians, sigmas = runningMedianAndSigma([]float64{
		5232.200195,
		5277.899902,
		5281.799805,
		5263.100098,
		5244.750000,
		5249.399902,
		5210.399902,
		5233.950195,
		5259.899902,
		5252.200195,
		5274.850098}, 10)
	require.Equal(t, []float64{
		5250.8000485,
		5250.8000485,
		5250.8000485,
		5250.8000485,
		5250.8000485,
		5250.8000485,
		5256.0500485,
		5256.0500485,
		5256.0500485,
		5256.0500485,
		5256.0500485,
	}, medians)

	require.Equal(t, []float64{
		21.60882309389966,
		21.60882309389966,
		21.60882309389966,
		21.60882309389966,
		21.60882309389966,
		21.60882309389966,
		22.313202647399915,
		22.313202647399915,
		22.313202647399915,
		22.313202647399915,
		22.313202647399915,
	}, sigmas)

	// N = 25, windowSize = 10
	medians, sigmas = runningMedianAndSigma([]float64{
		5232.200195,
		5277.899902,
		5281.799805,
		5263.100098,
		5244.750000,
		5249.399902,
		5210.399902,
		5233.950195,
		5259.899902,
		5252.200195,
		5274.850098,
		5225.649902,
		5221.700195,
		5094.149902,
		5036.000000,
		5007.899902,
		4853.100098,
		4867.250000,
		4882.049805,
		4899.700195,
		4830.100098,
		4931.850098,
		4845.350098,
		4718.649902,
		4760.399902,
	}, 10)

	require.Equal(t, []float64{
		5250.8000485,
		5250.8000485,
		5250.8000485,
		5250.8000485,
		5250.8000485,
		5250.8000485,
		5256.0500485,
		5250.8000485,
		5247.0749510000005,
		5239.350097500001,
		5229.8000485,
		5223.6750485,
		5223.6750485,
		5157.9250485,
		5065.0749510000005,
		5021.9499510000005,
		4953.8000485,
		4915.7751465,
		4890.875000,
		4874.649902499999,
		4860.1750489999995,
		4860.1750489999995,
		4860.1750489999995,
		4860.1750489999995,
		4860.1750489999995,
	}, medians)

	require.Equal(t, []float64{
		21.60882309389966,
		21.60882309389966,
		21.60882309389966,
		21.60882309389966,
		21.60882309389966,
		21.60882309389966,
		22.313202647399915,
		21.60882309389966,
		21.60882309389966,
		23.239827647400592,
		31.1346,
		47.999174999999994,
		64.78954809389965,
		162.27049809389965,
		254.7107519061003,
		218.387051906101,
		138.80842425870043,
		109.97192690610034,
		64.12245,
		54.74471515169923,
		51.594551906100335,
		51.594551906100335,
		51.594551906100335,
		51.594551906100335,
		51.594551906100335,
	}, sigmas)
}

func TestFilter(t *testing.T) {
	data := []float64{1, 2, 1, 1, 40, 2, 1, 1, 30, 40, 1, 1, 2, 1}
	outliers := Filter(data, DefaultWindow, DefaultN)
	require.Equal(t, []int{
		4, 7, 8, 9,
	}, outliers)

	data = FilterImpute(data, DefaultWindow, DefaultN)
	require.Equal(t, []float64{1.0, 2.0, 1.0, 1.0, 1.5, 2.0, 1.0, 1.0, 1.5, 1.5, 1.0, 1.0, 2.0, 1.0}, data)

	data = []float64{1, 2, 1, 1, 40, 2, 1, 1, 30, 40, 1, 1, 2, 1, 1}
	FilterImputeInPlace(data, DefaultWindow, DefaultN)
	require.Equal(t, []float64{
		1, 2, 1, 1, 1.5, 2, 1, 1, 1.5, 1.5, 1, 1, 1, 1, 1,
	}, data)
}
