// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package connector // import "go.opentelemetry.io/collector/connector"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector/internal"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/internal/fanoutconsumer"
)

// TracesRouterAndConsumer feeds the first consumer.Traces in each of the specified pipelines.
type TracesRouterAndConsumer interface {
	consumer.Traces
	Consumer(...component.ID) (consumer.Traces, error)
	PipelineIDs() []component.ID
	privateFunc()
}

type tracesRouter struct {
	consumer.Traces
	internal.BaseRouter[consumer.Traces]
}

func NewTracesRouter(cm map[component.ID]consumer.Traces) TracesRouterAndConsumer {
	consumers := make([]consumer.Traces, 0, len(cm))
	for _, cons := range cm {
		consumers = append(consumers, cons)
	}
	return &tracesRouter{
		Traces:     fanoutconsumer.NewTraces(consumers),
		BaseRouter: internal.NewBaseRouter(fanoutconsumer.NewTraces, cm),
	}
}

func (r *tracesRouter) privateFunc() {}