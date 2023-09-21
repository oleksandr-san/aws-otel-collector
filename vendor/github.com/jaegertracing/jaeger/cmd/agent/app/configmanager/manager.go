// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package configmanager

import (
	"context"

<<<<<<< HEAD
	"github.com/jaegertracing/jaeger/thrift-gen/baggage"
	"github.com/jaegertracing/jaeger/thrift-gen/sampling"
)

=======
	"github.com/jaegertracing/jaeger/proto-gen/api_v2"
	"github.com/jaegertracing/jaeger/thrift-gen/baggage"
)

// TODO this interface could be moved to pkg/clientcfg, along with grpc proxy,
// but not the metrics wrapper (because its metric names are specific to agent).

>>>>>>> main
// ClientConfigManager decides:
// 1) which sampling strategy a given service should be using
// 2) which baggage restrictions a given service should be using.
type ClientConfigManager interface {
<<<<<<< HEAD
	GetSamplingStrategy(ctx context.Context, serviceName string) (*sampling.SamplingStrategyResponse, error)
=======
	GetSamplingStrategy(ctx context.Context, serviceName string) (*api_v2.SamplingStrategyResponse, error)
>>>>>>> main
	GetBaggageRestrictions(ctx context.Context, serviceName string) ([]*baggage.BaggageRestriction, error)
}
