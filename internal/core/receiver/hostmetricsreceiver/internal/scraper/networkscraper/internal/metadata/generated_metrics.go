// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/model/pdata"
)

// Type is the component type name.
const Type config.Type = "network"

// MetricIntf is an interface to generically interact with generated metric.
type MetricIntf interface {
	Name() string
	New() pdata.Metric
	Init(metric pdata.Metric)
}

// Intentionally not exposing this so that it is opaque and can change freely.
type metricImpl struct {
	name     string
	initFunc func(pdata.Metric)
}

// Name returns the metric name.
func (m *metricImpl) Name() string {
	return m.name
}

// New creates a metric object preinitialized.
func (m *metricImpl) New() pdata.Metric {
	metric := pdata.NewMetric()
	m.Init(metric)
	return metric
}

// Init initializes the provided metric object.
func (m *metricImpl) Init(metric pdata.Metric) {
	m.initFunc(metric)
}

type metricStruct struct {
	SystemNetworkConnections MetricIntf
	SystemNetworkDropped     MetricIntf
	SystemNetworkErrors      MetricIntf
	SystemNetworkIo          MetricIntf
	SystemNetworkPackets     MetricIntf
}

// Names returns a list of all the metric name strings.
func (m *metricStruct) Names() []string {
	return []string{
		"system.network.connections",
		"system.network.dropped",
		"system.network.errors",
		"system.network.io",
		"system.network.packets",
	}
}

var metricsByName = map[string]MetricIntf{
	"system.network.connections": Metrics.SystemNetworkConnections,
	"system.network.dropped":     Metrics.SystemNetworkDropped,
	"system.network.errors":      Metrics.SystemNetworkErrors,
	"system.network.io":          Metrics.SystemNetworkIo,
	"system.network.packets":     Metrics.SystemNetworkPackets,
}

func (m *metricStruct) ByName(n string) MetricIntf {
	return metricsByName[n]
}

// Metrics contains a set of methods for each metric that help with
// manipulating those metrics.
var Metrics = &metricStruct{
	&metricImpl{
		"system.network.connections",
		func(metric pdata.Metric) {
			metric.SetName("system.network.connections")
			metric.SetDescription("The number of connections.")
			metric.SetUnit("{connections}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.network.dropped",
		func(metric pdata.Metric) {
			metric.SetName("system.network.dropped")
			metric.SetDescription("The number of packets dropped.")
			metric.SetUnit("{packets}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.network.errors",
		func(metric pdata.Metric) {
			metric.SetName("system.network.errors")
			metric.SetDescription("The number of errors encountered.")
			metric.SetUnit("{errors}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.network.io",
		func(metric pdata.Metric) {
			metric.SetName("system.network.io")
			metric.SetDescription("The number of bytes transmitted and received.")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.network.packets",
		func(metric pdata.Metric) {
			metric.SetName("system.network.packets")
			metric.SetDescription("The number of packets transferred.")
			metric.SetUnit("{packets}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
}

// M contains a set of methods for each metric that help with
// manipulating those metrics. M is an alias for Metrics
var M = Metrics

// Labels contains the possible metric labels that can be used.
var Labels = struct {
	// Device (Name of the network interface.)
	Device string
	// Direction (Direction of flow of bytes/opertations (receive or transmit).)
	Direction string
	// Protocol (Network protocol, e.g. TCP or UDP.)
	Protocol string
	// State (State of the network connection.)
	State string
}{
	"device",
	"direction",
	"protocol",
	"state",
}

// L contains the possible metric labels that can be used. L is an alias for
// Labels.
var L = Labels

// LabelDirection are the possible values that the label "direction" can have.
var LabelDirection = struct {
	Receive  string
	Transmit string
}{
	"receive",
	"transmit",
}

// LabelProtocol are the possible values that the label "protocol" can have.
var LabelProtocol = struct {
	Tcp string
}{
	"tcp",
}
