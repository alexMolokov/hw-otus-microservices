package fasthttpprom

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fasthttp/router"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

// FastHTTPrometheus ...
type FastHTTPrometheus struct {
	requestsTotal   *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
	requestInFlight *prometheus.GaugeVec
	router          *router.Router
	defaultURL      string
}

func create(registry prometheus.Registerer, serviceName, namespace, subsystem string, labels map[string]string,
) *FastHTTPrometheus {
	constLabels := make(prometheus.Labels)
	if serviceName != "" {
		constLabels["service"] = serviceName
	}
	for label, value := range labels {
		constLabels[label] = value
	}

	counter := promauto.With(registry).NewCounterVec(
		prometheus.CounterOpts{
			Name:        prometheus.BuildFQName(namespace, subsystem, "requests_total"),
			Help:        "Count all http requests by status code, method and path.",
			ConstLabels: constLabels,
		},
		[]string{"code", "method", "url"},
	)
	histogram := promauto.With(registry).NewHistogramVec(prometheus.HistogramOpts{
		Name:        prometheus.BuildFQName(namespace, subsystem, "request_duration_seconds"),
		Help:        "Duration of all HTTP requests by status code, method and path.",
		ConstLabels: constLabels,
		Buckets: []float64{
			0.001, // 1ms
			0.002,
			0.005,
			0.01, // 10ms
			0.02,
			0.05,
			0.1, // 100 ms
			0.2,
			0.5,
			1.0, // 1s
			2.0,
			5.0,
			10.0, // 10s
		},
	},
		[]string{"code", "method", "url"},
	)

	gauge := promauto.With(registry).NewGaugeVec(prometheus.GaugeOpts{
		Name:        prometheus.BuildFQName(namespace, subsystem, "requests_in_progress_total"),
		Help:        "All the requests in progress",
		ConstLabels: constLabels,
	}, []string{"method"})

	return &FastHTTPrometheus{
		requestsTotal:   counter,
		requestDuration: histogram,
		requestInFlight: gauge,
		defaultURL:      "/metrics",
	}
}

// New creates a new instance of FastHTTPrometheus middleware
// serviceName is available as a const label.
func New(serviceName string) *FastHTTPrometheus {
	return create(prometheus.DefaultRegisterer, serviceName, "http", "", nil)
}

// NewWith creates a new instance of FastHTTPrometheus middleware but with an ability
// to pass namespace and a custom subsystem
// Here serviceName is created as a constant-label for the metrics
// Namespace, subsystem get prefixed to the metrics.
//
// For e.g. namespace = "my_app", subsystem = "http" then metrics would be.
// `my_app_http_requests_total{...,service= "serviceName"}`.
func NewWith(serviceName, namespace, subsystem string) *FastHTTPrometheus {
	return create(prometheus.DefaultRegisterer, serviceName, namespace, subsystem, nil)
}

// NewWithLabels creates a new instance of FastHTTPrometheus middleware but with an ability
// to pass namespace and a custom subsystem
// Here labels are created as a constant-labels for the metrics
// Namespace, subsystem get prefixed to the metrics.
//
// For e.g. namespace = "my_app", subsystem = "http" and labels = map[string]string{"key1": "value1", "key2":"value2"}
// then then metrics would become.
// `my_app_http_requests_total{...,key1= "value1", key2= "value2" }`.
func NewWithLabels(labels map[string]string, namespace, subsystem string) *FastHTTPrometheus {
	return create(prometheus.DefaultRegisterer, "", namespace, subsystem, labels)
}

// NewWithRegistry creates a new instance of FastHTTPrometheus middleware but with an ability
// to pass a custom registry, serviceName, namespace, subsystem and labels
// Here labels are created as a constant-labels for the metrics
// Namespace, subsystem get prefixed to the metrics.
//
// For e.g. namespace = "my_app", subsystem = "http" and labels = map[string]string{"key1": "value1", "key2":"value2"}
// then then metrics would become
// `my_app_http_requests_total{...,key1= "value1", key2= "value2" }`.
func NewWithRegistry(registry prometheus.Registerer, serviceName, namespace, subsystem string,
	labels map[string]string,
) *FastHTTPrometheus {
	return create(registry, serviceName, namespace, subsystem, labels)
}

func (ps *FastHTTPrometheus) Register(r *router.Router) {
	r.GET(ps.defaultURL, prometheusHandler())
	ps.router = r
}

func (ps *FastHTTPrometheus) Middleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		start := time.Now()
		method := string(ctx.Method())
		req := &ctx.Request

		if string(req.RequestURI()) == ps.defaultURL {
			next(ctx)
			return
		}

		ps.requestInFlight.WithLabelValues(method).Inc()
		defer func() {
			ps.requestInFlight.WithLabelValues(method).Dec()
		}()

		next(ctx)

		endpoint := string(ctx.URI().Path())
		routeList := ps.router.List()
		paths, ok := routeList[method]
		handler, _ := ps.router.Lookup(method, string(ctx.Path()), ctx)
		if ok {
			for _, v := range paths {
				tmp, _ := ps.router.Lookup(string(ctx.Request.Header.Method()), v, ctx)
				if fmt.Sprintf("%v", tmp) == fmt.Sprintf("%v", handler) {
					endpoint = v
					break
				}
			}
		}

		httpStatus := strconv.Itoa(ctx.Response.StatusCode())
		ps.requestsTotal.WithLabelValues(httpStatus, method, endpoint).Inc()

		elapsed := float64(time.Since(start).Nanoseconds()) / 1e9
		ps.requestDuration.WithLabelValues(httpStatus, method, endpoint).Observe(elapsed)
	}
}

func prometheusHandler() fasthttp.RequestHandler {
	return fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())
}
