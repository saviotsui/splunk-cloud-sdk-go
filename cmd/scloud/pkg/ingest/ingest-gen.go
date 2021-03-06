// Package ingest -- generated by scloudgen
// !! DO NOT EDIT !!
//
package ingest

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/flags"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/jsonx"
	model "github.com/splunk/splunk-cloud-sdk-go/services/ingest"
)

// PostEvents Sends events.
func PostEvents(cmd *cobra.Command, args []string) error {

	var err error

	// Parse all flags

	var attributes map[string]interface{}
	err = flags.ParseFlag(cmd.Flags(), "attributes", &attributes)
	if err != nil {
		return fmt.Errorf(`error parsing "attributes": ` + err.Error())
	}
	var format string
	err = flags.ParseFlag(cmd.Flags(), "format", &format)
	if err != nil {
		return fmt.Errorf(`error parsing "format": ` + err.Error())
	}
	var hostDefault string
	host := &hostDefault
	err = flags.ParseFlag(cmd.Flags(), "host", &host)
	if err != nil {
		return fmt.Errorf(`error parsing "host": ` + err.Error())
	}
	var idDefault string
	id := &idDefault
	err = flags.ParseFlag(cmd.Flags(), "id", &id)
	if err != nil {
		return fmt.Errorf(`error parsing "id": ` + err.Error())
	}
	var nanosDefault int32
	nanos := &nanosDefault
	err = flags.ParseFlag(cmd.Flags(), "nanos", &nanos)
	if err != nil {
		return fmt.Errorf(`error parsing "nanos": ` + err.Error())
	}
	var sourceDefault string
	source := &sourceDefault
	err = flags.ParseFlag(cmd.Flags(), "source", &source)
	if err != nil {
		return fmt.Errorf(`error parsing "source": ` + err.Error())
	}
	var sourcetypeDefault string
	sourcetype := &sourcetypeDefault
	err = flags.ParseFlag(cmd.Flags(), "sourcetype", &sourcetype)
	if err != nil {
		return fmt.Errorf(`error parsing "sourcetype": ` + err.Error())
	}
	var timestampDefault int64
	timestamp := &timestampDefault
	err = flags.ParseFlag(cmd.Flags(), "timestamp", &timestamp)
	if err != nil {
		return fmt.Errorf(`error parsing "timestamp": ` + err.Error())
	}
	// Form the request body
	generated_request_body := []model.Event{
		{
			Attributes: attributes,

			Host:       host,
			Id:         id,
			Nanos:      nanos,
			Source:     source,
			Sourcetype: sourcetype,
			Timestamp:  timestamp,
		},
	}

	resp, err := PostEventsOverride(generated_request_body, format)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// PostMetrics Sends metric events.
func PostMetrics(cmd *cobra.Command, args []string) error {

	var err error

	// Parse all flags

	var defaultDimensions map[string]string
	err = flags.ParseFlag(cmd.Flags(), "default-dimensions", &defaultDimensions)
	if err != nil {
		return fmt.Errorf(`error parsing "default-dimensions": ` + err.Error())
	}
	var defaultTypeDefault string
	defaultType := &defaultTypeDefault
	err = flags.ParseFlag(cmd.Flags(), "default-type", &defaultType)
	if err != nil {
		return fmt.Errorf(`error parsing "default-type": ` + err.Error())
	}
	var defaultUnitDefault string
	defaultUnit := &defaultUnitDefault
	err = flags.ParseFlag(cmd.Flags(), "default-unit", &defaultUnit)
	if err != nil {
		return fmt.Errorf(`error parsing "default-unit": ` + err.Error())
	}
	var hostDefault string
	host := &hostDefault
	err = flags.ParseFlag(cmd.Flags(), "host", &host)
	if err != nil {
		return fmt.Errorf(`error parsing "host": ` + err.Error())
	}
	var idDefault string
	id := &idDefault
	err = flags.ParseFlag(cmd.Flags(), "id", &id)
	if err != nil {
		return fmt.Errorf(`error parsing "id": ` + err.Error())
	}
	var nanosDefault int32
	nanos := &nanosDefault
	err = flags.ParseFlag(cmd.Flags(), "nanos", &nanos)
	if err != nil {
		return fmt.Errorf(`error parsing "nanos": ` + err.Error())
	}
	var sourceDefault string
	source := &sourceDefault
	err = flags.ParseFlag(cmd.Flags(), "source", &source)
	if err != nil {
		return fmt.Errorf(`error parsing "source": ` + err.Error())
	}
	var sourcetypeDefault string
	sourcetype := &sourcetypeDefault
	err = flags.ParseFlag(cmd.Flags(), "sourcetype", &sourcetype)
	if err != nil {
		return fmt.Errorf(`error parsing "sourcetype": ` + err.Error())
	}
	var timestampDefault int64
	timestamp := &timestampDefault
	err = flags.ParseFlag(cmd.Flags(), "timestamp", &timestamp)
	if err != nil {
		return fmt.Errorf(`error parsing "timestamp": ` + err.Error())
	}
	// Form the request body
	generated_request_body := []model.MetricEvent{
		{
			Attributes: &model.MetricAttribute{
				DefaultDimensions: defaultDimensions,
				DefaultType:       defaultType,
				DefaultUnit:       defaultUnit,
			},

			Host:       host,
			Id:         id,
			Nanos:      nanos,
			Source:     source,
			Sourcetype: sourcetype,
			Timestamp:  timestamp,
		},
	}

	resp, err := PostMetricsOverride(generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// UploadFiles Upload a CSV or text file that contains events.
func UploadFiles(cmd *cobra.Command, args []string) error {

	var err error

	// Parse all flags

	var fileName string
	err = flags.ParseFlag(cmd.Flags(), "file-name", &fileName)
	if err != nil {
		return fmt.Errorf(`error parsing "file-name": ` + err.Error())
	}

	resp, err := UploadFilesOverride(fileName)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}
