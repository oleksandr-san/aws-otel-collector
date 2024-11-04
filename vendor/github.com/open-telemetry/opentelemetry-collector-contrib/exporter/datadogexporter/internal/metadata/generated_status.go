// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/component"
)

var (
	Type      = component.MustNewType("datadog")
	ScopeName = "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter"
)

const (
	LogsStability    = component.StabilityLevelAlpha
	TracesStability  = component.StabilityLevelBeta
	MetricsStability = component.StabilityLevelBeta
)