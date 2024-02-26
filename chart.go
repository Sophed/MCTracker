package main

import (
	"fmt"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

const PORT = 3000

func lineItems(server Server) []opts.LineData {
	items := make([]opts.LineData, 0)

	for _, count := range server.Counts {
		items = append(items, opts.LineData{Value: count})
	}

	return items
}

func get(w http.ResponseWriter, r *http.Request) {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithYAxisOpts(opts.YAxis{
			Type: "value",
		}),
	)

	var xValues []string

	for timestamp := range serverList[0].Counts {
		xValues = append(xValues, timestamp)
	}
	line.SetXAxis(xValues)

	for _, server := range serverList {
		line.AddSeries(server.Name, lineItems(server))
	}

	line.Render(w)
}

func server() {
	port := fmt.Sprintf(":%d", PORT)
	fmt.Println("Server started at http://localhost" + port + "/")
	http.HandleFunc("/", get)
	http.ListenAndServe(port, nil)
}
