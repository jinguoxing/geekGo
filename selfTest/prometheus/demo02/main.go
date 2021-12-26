package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"time"
)


var (
	// counter 监控类型  只支持自增
	//http_request_total = promauto.NewCounter(
	//			prometheus.CounterOpts{
	//				Name :"http_request_total",
	//				Help : "统计http的请求总数",
	//			},
	//	)

	http_request_total = promauto.NewCounterVec(
				prometheus.CounterOpts{
					Name :"http_request_total",
					Help : "统计http的请求总数",
				},
				[]string{"path"},
		)

	// Gange 监控类型  支持 加 和 减的操作
	http_request_in_flight = promauto.NewGauge(
			prometheus.GaugeOpts{
			 Name : "http_request_in_flight",
			 Help : "http request flight统计",

			})
    // 直方图
	http_request_druation_seconds = promauto.NewHistogram(
				prometheus.HistogramOpts{
					Name:		"http_request_duration_seconds",
					Help:		" HTTP requests 延时统计",
					 Buckets:	[]float64{.1, .2, .4, 1, 3, 8, 20, 60, 120},
				},
		)

	http_request_summary_seconds = promauto.NewSummary(
			prometheus.SummaryOpts{
				Name : "http_request_summary_seconds",
				Help : "HTTP requests 摘要统计",
				Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},

			},
		)

)

func main(){

	http.HandleFunc("/test", func(response http.ResponseWriter,request  *http.Request) {

		now := time.Now()
		response.Write([]byte("你好"))
		//http_request_total.Inc()

		http_request_total.WithLabelValues("test").Inc()


		http_request_in_flight.Inc()

		time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)

		http_request_druation_seconds.Observe(time.Since(now).Seconds())
		http_request_summary_seconds.Observe(time.Since(now).Seconds())

	})

	http.HandleFunc("/foo", func(response http.ResponseWriter,request  *http.Request) {

		now := time.Now()
		response.Write([]byte("foo"))


		http_request_total.With(prometheus.Labels{"path":"foo"}).Inc()

		http_request_in_flight.Inc()

		time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)

		http_request_druation_seconds.Observe(time.Since(now).Seconds())
		http_request_summary_seconds.Observe(time.Since(now).Seconds())

	})


	http.Handle("/metrics",promhttp.Handler())

	http.ListenAndServe(":8989",nil)

}
