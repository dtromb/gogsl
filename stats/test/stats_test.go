package stats 

import (
	"fmt"
	"testing"
	"github.com/dtromb/gogsl/stats"
	"github.com/dtromb/gogsl/sort"
)

func TestStats(t *testing.T) {
	var data []float64 = []float64{17.2, 18.1, 16.5, 18.3, 12.6}
	
	mean := stats.Mean(data, 1, 5)
	variance := stats.Variance(data, 1, 5)
	largest := stats.Max(data, 1, 5)
	smallest := stats.Min(data, 1, 5)
	fmt.Printf("The dataset is %g, %g, %g, %g, %g\n",
               data[0], data[1], data[2], data[3], data[4])
  	fmt.Printf ("The sample mean is %g\n", mean)
  	fmt.Printf ("The estimated variance is %g\n", variance)
  	fmt.Printf ("The largest value is %g\n", largest)
  	fmt.Printf ("The smallest value is %g\n", smallest)		
}

// XXX - Is this performing ridiculously slowly?   Look into it...
func TestStatsSorted(t *testing.T) {
	
 	var data []float64 = []float64{17.2, 18.1, 16.5, 18.3, 12.6}
 	fmt.Printf("Original dataset:  %g, %g, %g, %g, %g\n",
         	    data[0], data[1], data[2], data[3], data[4])
	sort.Sort(data, 1, 5)
	fmt.Printf("Sorted dataset: %g, %g, %g, %g, %g\n",
         	    data[0], data[1], data[2], data[3], data[4])
    median := stats.MedianFromSortedData(data, 1, 5)
	upperq := stats.QuantileFromSortedData(data, 1, 5, 0.75)
    lowerq := stats.QuantileFromSortedData(data, 1, 5, 0.25)
  	fmt.Printf("The median is %g\n", median);
  	fmt.Printf("The upper quartile is %g\n", upperq);
  	fmt.Printf("The lower quartile is %g\n", lowerq);
}