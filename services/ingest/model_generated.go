/*
 * Copyright © 2020 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 *
 * Ingest API
 *
 * Use the Ingest service in Splunk Cloud Services to send event and metrics data, or upload a static file, to Splunk Cloud Services.
 *
 * API version: v1beta2.9 (recommended default)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
 */

package ingest

type Error struct {
	Code    *string                `json:"code,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
	Message *string                `json:"message,omitempty"`
}

type Event struct {
	// JSON object for the event.
	Body interface{} `json:"body"`
	// Specifies a JSON object that contains explicit custom fields to be defined at index time.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// The host value assigned to the event data. Typically, this is the hostname of the client from which you are sending data.
	Host *string `json:"host,omitempty"`
	// An optional ID that uniquely identifies the event data. It is used to deduplicate the data if same data is set multiple times. If ID is not specified, it will be assigned by the system.
	Id *string `json:"id,omitempty"`
	// Optional nanoseconds part of the timestamp.
	Nanos *int32 `json:"nanos,omitempty"`
	// The source value to assign to the event data. For example, if you are sending data from an app that you are developing, set this key to the name of the app.
	Source *string `json:"source,omitempty"`
	// The sourcetype value assigned to the event data.
	Sourcetype *string `json:"sourcetype,omitempty"`
	// Epoch time in milliseconds.
	Timestamp *int64 `json:"timestamp,omitempty"`
}

type HttpResponse struct {
	Code    *string                `json:"code,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
	Message *string                `json:"message,omitempty"`
}

type Metric struct {
	// Name of the metric e.g. CPU, Memory etc.
	Name string `json:"name"`
	// Dimensions allow metrics to be classified e.g. {\"Server\":\"nginx\", \"Region\":\"us-west-1\", ...}
	Dimensions map[string]string `json:"dimensions,omitempty"`
	// Type of metric. Default is g for gauge.
	Type *string `json:"type,omitempty"`
	// Unit of the metric e.g. percent, megabytes, seconds etc.
	Unit *string `json:"unit,omitempty"`
	// Value of the metric. If not specified, it will be defaulted to 0.
	Value *float64 `json:"value,omitempty"`
}

type MetricAttribute struct {
	// Optional. If set, individual metrics inherit these dimensions and can override any and/or all of them.
	DefaultDimensions map[string]string `json:"defaultDimensions,omitempty"`
	// Optional. If set, individual metrics inherit this type and can optionally override.
	DefaultType *string `json:"defaultType,omitempty"`
	// Optional. If set, individual metrics inherit this unit and can optionally override.
	DefaultUnit *string `json:"defaultUnit,omitempty"`
}

type MetricEvent struct {
	// Specifies multiple related metrics e.g. Memory, CPU etc.
	Body       []Metric         `json:"body"`
	Attributes *MetricAttribute `json:"attributes,omitempty"`
	// The host value assigned to the event data. Typically, this is the hostname of the client from which you are sending data.
	Host *string `json:"host,omitempty"`
	// An optional ID that uniquely identifies the metric data. It is used to deduplicate the data if same data is set multiple times. If ID is not specified, it will be assigned by the system.
	Id *string `json:"id,omitempty"`
	// Optional nanoseconds part of the timestamp.
	Nanos *int32 `json:"nanos,omitempty"`
	// The source value to assign to the event data. For example, if you are sending data from an app that you are developing, set this key to the name of the app.
	Source *string `json:"source,omitempty"`
	// The sourcetype value assigned to the event data.
	Sourcetype *string `json:"sourcetype,omitempty"`
	// Epoch time in milliseconds.
	Timestamp *int64 `json:"timestamp,omitempty"`
}
