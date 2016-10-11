package collectors

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/cloudfoundry-community/firehose_exporter/metrics"
)

type containerMetricsCollector struct {
	namespace                  string
	metricsStore               *metrics.Store
	deploymentsFilter          map[string]struct{}
	cpuPercentageMetricDesc    *prometheus.Desc
	memoryBytesMetricDesc      *prometheus.Desc
	diskBytesMetricDesc        *prometheus.Desc
	memoryBytesQuotaMetricDesc *prometheus.Desc
	diskBytesQuotaMetricDesc   *prometheus.Desc
}

func NewContainerMetricsCollector(
	namespace string,
	metricsStore *metrics.Store,
	dopplerDeployments []string,
) *containerMetricsCollector {
	cpuPercentageMetricDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, container_metrics_subsystem, "cpu_percentage"),
		"Cloud Foundry Firehose container metric: CPU used, on a scale of 0 to 100.",
		[]string{"origin", "bosh_deployment", "bosh_job", "bosh_index", "bosh_ip", "application_id", "instance_id"},
		nil,
	)

	memoryBytesMetricDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, container_metrics_subsystem, "memory_bytes"),
		"Cloud Foundry Firehose container metric: bytes of memory used.",
		[]string{"origin", "bosh_deployment", "bosh_job", "bosh_index", "bosh_ip", "application_id", "instance_id"},
		nil,
	)

	diskBytesMetricDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, container_metrics_subsystem, "disk_bytes"),
		"Cloud Foundry Firehose container metric: bytes of disk used.",
		[]string{"origin", "bosh_deployment", "bosh_job", "bosh_index", "bosh_ip", "application_id", "instance_id"},
		nil,
	)

	memoryBytesQuotaMetricDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, container_metrics_subsystem, "memory_bytes_quota"),
		"Cloud Foundry Firehose container metric: maximum bytes of memory allocated to container.",
		[]string{"origin", "bosh_deployment", "bosh_job", "bosh_index", "bosh_ip", "application_id", "instance_id"},
		nil,
	)

	diskBytesQuotaMetricDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, container_metrics_subsystem, "disk_bytes_quota"),
		"Cloud Foundry Firehose container metric: maximum bytes of disk allocated to container.",
		[]string{"origin", "bosh_deployment", "bosh_job", "bosh_index", "bosh_ip", "application_id", "instance_id"},
		nil,
	)

	deploymentsFilter := map[string]struct{}{}
	for _, deployment := range dopplerDeployments {
		deploymentsFilter[deployment] = struct{}{}
	}

	collector := &containerMetricsCollector{
		namespace:                  namespace,
		metricsStore:               metricsStore,
		deploymentsFilter:          deploymentsFilter,
		cpuPercentageMetricDesc:    cpuPercentageMetricDesc,
		memoryBytesMetricDesc:      memoryBytesMetricDesc,
		diskBytesMetricDesc:        diskBytesMetricDesc,
		memoryBytesQuotaMetricDesc: memoryBytesQuotaMetricDesc,
		diskBytesQuotaMetricDesc:   diskBytesQuotaMetricDesc,
	}
	return collector
}

func (c containerMetricsCollector) Collect(ch chan<- prometheus.Metric) {
	for _, containerMetric := range c.metricsStore.GetContainerMetrics() {
		_, ok := c.deploymentsFilter[containerMetric.Deployment]
		if len(c.deploymentsFilter) == 0 || ok {
			ch <- prometheus.MustNewConstMetric(
				c.cpuPercentageMetricDesc,
				prometheus.GaugeValue,
				containerMetric.CpuPercentage,
				containerMetric.Origin,
				containerMetric.Deployment,
				containerMetric.Job,
				containerMetric.Index,
				containerMetric.IP,
				containerMetric.ApplicationId,
				strconv.Itoa(int(containerMetric.InstanceIndex)),
			)
			ch <- prometheus.MustNewConstMetric(
				c.memoryBytesMetricDesc,
				prometheus.GaugeValue,
				float64(containerMetric.MemoryBytes),
				containerMetric.Origin,
				containerMetric.Deployment,
				containerMetric.Job,
				containerMetric.Index,
				containerMetric.IP,
				containerMetric.ApplicationId,
				strconv.Itoa(int(containerMetric.InstanceIndex)),
			)
			ch <- prometheus.MustNewConstMetric(
				c.diskBytesMetricDesc,
				prometheus.GaugeValue,
				float64(containerMetric.DiskBytes),
				containerMetric.Origin,
				containerMetric.Deployment,
				containerMetric.Job,
				containerMetric.Index,
				containerMetric.IP,
				containerMetric.ApplicationId,
				strconv.Itoa(int(containerMetric.InstanceIndex)),
			)
			ch <- prometheus.MustNewConstMetric(
				c.memoryBytesQuotaMetricDesc,
				prometheus.GaugeValue,
				float64(containerMetric.MemoryBytesQuota),
				containerMetric.Origin,
				containerMetric.Deployment,
				containerMetric.Job,
				containerMetric.Index,
				containerMetric.IP,
				containerMetric.ApplicationId,
				strconv.Itoa(int(containerMetric.InstanceIndex)),
			)
			ch <- prometheus.MustNewConstMetric(
				c.diskBytesQuotaMetricDesc,
				prometheus.GaugeValue,
				float64(containerMetric.DiskBytesQuota),
				containerMetric.Origin,
				containerMetric.Deployment,
				containerMetric.Job,
				containerMetric.Index,
				containerMetric.IP,
				containerMetric.ApplicationId,
				strconv.Itoa(int(containerMetric.InstanceIndex)),
			)
		}
	}
}

func (c containerMetricsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.cpuPercentageMetricDesc
	ch <- c.memoryBytesMetricDesc
	ch <- c.diskBytesMetricDesc
	ch <- c.memoryBytesQuotaMetricDesc
	ch <- c.diskBytesQuotaMetricDesc
}
