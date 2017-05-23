package talib4g

type commidityChannelIndexIndicator struct {
	series *TimeSeries
	window int
}

// Returns a new Commodity Channel Index Indicator
// http://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:commodity_channel_index_cci
func NewCCIIndicator(ts *TimeSeries, window int) Indicator {
	return commidityChannelIndexIndicator{
		series: ts,
		window: window,
	}
}

func (ccii commidityChannelIndexIndicator) Calculate(index int) float64 {
	typicalPrice := NewTypicalPriceIndicator(ccii.series)
	typicalPriceSma := NewSimpleMovingAverage(typicalPrice, ccii.window)
	meanDeviation := NewMeanDeviationIndicator(NewClosePriceIndicator(ccii.series), ccii.window)

	return (typicalPrice.Calculate(index) - typicalPriceSma.Calculate(index)) / (meanDeviation.Calculate(index) * 0.015)
}