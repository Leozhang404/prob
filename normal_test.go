package distributions

import (
  "testing"
)

type normalTest struct {
  dist        Distribution
  mean        float64
  variance    float64
  stdDev      float64
  relStdDev   float64
  skewness    float64
  kurtosis    float64
  pdf         []inOut
  cdf         []inOut
}

//Test at http://keisan.casio.com/exec/system/1180573188
func Test_Normal(t *testing.T) {
  examples := []normalTest{
    normalTest{
      dist:       Normal{1.0, 4.0},
      mean:       1.0,
      variance:   16.0,
      stdDev:     4.0,
      relStdDev:  4.0,
      skewness:   0.0,
      kurtosis:   3.0,
      pdf: []inOut{
        inOut{ in: -4.0,  out: 0.04566227134725547624776 },
        inOut{ in: 0.5,   out: 0.09895942173618737103265 },
        inOut{ in: 12.0,  out: 0.002273390625397763192514 },
      },
      cdf: []inOut{
        inOut{ in: -4.0,  out: 0.1056497736668552576888 },
        inOut{ in: 0.5,   out: 0.4502617751698871070207 },
        inOut{ in: 12.0,  out: 0.9970202367649454432457 },
      },
    },
    normalTest{
      dist:       Normal{10.0, 2.0},
      mean:       10.0,
      variance:   4.0,
      stdDev:     2.0,
      relStdDev:  0.2,
      skewness:   0.0,
      kurtosis:   3.0,
      pdf: []inOut{
        inOut{ in: 4.0,   out: 0.002215924205969003587801 },
        inOut{ in: 6.0,   out: 0.02699548325659402597528 },
        inOut{ in: 16.0,  out: 0.002215924205969003587801 },
      },
      cdf: []inOut{
        inOut{ in: 4.0,   out: 0.001349898031630094526652 },
        inOut{ in: 6.0,   out: 0.02275013194817920720028 },
        inOut{ in: 16.0,  out: 0.9986501019683699054733 },
      },
    },
  }

  for _, example := range examples {
    mean, err := example.dist.Mean()
    if err != nil || mean != example.mean {
      t.Fatalf("\nMean:\n  Expected: %f\n  Got: %f\n", example.mean, mean)
    }
    variance, err := example.dist.Variance()
    if err != nil || variance != example.variance {
      t.Fatalf("\nVariance:\n  Expected: %f\n  Got: %f\n", example.variance, variance)
    }
    stdDev, err := example.dist.StdDev()
    if err != nil || stdDev != example.stdDev {
      t.Fatalf("\nStdDev:\n  Expected: %f\n  Got: %f\n", example.stdDev, stdDev)
    }
    relStdDev, err := example.dist.RelStdDev()
    if err != nil || relStdDev != example.relStdDev {
      t.Fatalf("\nRelStdDev:\n  Expected: %f\n  Got: %f\n", example.relStdDev, relStdDev)
    }
    skewness, err := example.dist.Skewness()
    if err != nil || skewness != example.skewness {
      t.Fatalf("\nSkewness:\n  Expected: %f\n  Got: %f\n", example.skewness, skewness)
    }
    kurtosis, err := example.dist.Kurtosis()
    if err != nil || kurtosis != example.kurtosis {
      t.Fatalf("\nKurtosis:\n  Expected: %f\n  Got: %f\n", example.kurtosis, kurtosis)
    }
    for _, pdf := range example.pdf {
      out, err := example.dist.Pdf(pdf.in)
      if err != nil || !floatsPicoEqual(out, pdf.out) {
        t.Fatalf("\nPdf of %f:\n  Expected: %f\n  Got: %f\n", pdf.in, pdf.out, out)
      }
    }
    for _, cdf := range example.cdf {
      out, err := example.dist.Cdf(cdf.in)
      if err != nil || !floatsPicoEqual(out, cdf.out) {
        t.Fatalf("\nCdf of %f:\n  Expected: %f\n  Got: %f\n", cdf.in, cdf.out, out)
      }
    }
    samples, err := example.dist.Sample(1000000)
    if err != nil {
      t.Fatalf("\nCould not generate 1,000,000 samples.")
    }
    sampleMean := averageFloats(samples)
    if !floatsDeciEqual(example.mean, sampleMean) {
      t.Fatalf("\nSample average:\n  Expected: %f\n  Got: %f\n", example.mean, sampleMean)
    }
    sampleVar := varianceFloats(samples, sampleMean)
    if !floatsDeciEqual(example.variance, sampleVar) {
      t.Fatalf("\nSample variance:\n  Expected: %f\n  Got: %f\n", example.variance, sampleVar)
    }
  }
}