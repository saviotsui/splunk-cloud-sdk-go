/*
 * Copyright © 2019 Splunk, Inc.
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
 * KV Store API
 *
 * With the Splunk Cloud KV store service in Splunk Cloud Services, you can save and retrieve data within your Splunk Cloud apps, enabling you to manage and maintain state in your application.
 *
 * API version: v1beta1.2 (recommended default)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
 */

package kvstore

import (
	"net/http"

	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

const serviceCluster = "api"

type Service services.BaseService

// NewService creates a new kvstore service client from the given Config
func NewService(config *services.Config) (*Service, error) {
	baseClient, err := services.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Service{Client: baseClient}, nil
}

/*
	CreateIndex - Creates an index on a collection.
	Parameters:
		collection: The name of the collection.
		indexDefinition
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) CreateIndex(collection string, indexDefinition IndexDefinition, resp ...*http.Response) (*IndexDescription, error) {
	pp := struct {
		Collection string
	}{
		Collection: collection,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/kvstore/v1beta1/collections/{{.Collection}}/indexes`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: indexDefinition})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb IndexDescription
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	DeleteIndex - Removes an index from a collection.
	Parameters:
		collection: The name of the collection.
		index: The name of the index.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeleteIndex(collection string, index string, resp ...*http.Response) error {
	pp := struct {
		Collection string
		Index      string
	}{
		Collection: collection,
		Index:      index,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/kvstore/v1beta1/collections/{{.Collection}}/indexes/{{.Index}}`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	return err
}

/*
	DeleteRecordByKey - Deletes a record with a given key.
	Parameters:
		collection: The name of the collection.
		key: The key of the record.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeleteRecordByKey(collection string, key string, resp ...*http.Response) error {
	pp := struct {
		Collection string
		Key        string
	}{
		Collection: collection,
		Key:        key,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/kvstore/v1beta1/collections/{{.Collection}}/records/{{.Key}}`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	return err
}

/*
	DeleteRecords - Removes records in a collection that match the query.
	Parameters:
		collection: The name of the collection.
		query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeleteRecords(collection string, query *DeleteRecordsQueryParams, resp ...*http.Response) error {
	values := util.ParseURLParams(query)
	pp := struct {
		Collection string
	}{
		Collection: collection,
	}
	u, err := s.Client.BuildURLFromPathParams(values, serviceCluster, `/kvstore/v1beta1/collections/{{.Collection}}/query`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	return err
}

/*
	GetRecordByKey - Returns a record with a given key.
	Parameters:
		collection: The name of the collection.
		key: The key of the record.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetRecordByKey(collection string, key string, resp ...*http.Response) (*map[string]interface{}, error) {
	pp := struct {
		Collection string
		Key        string
	}{
		Collection: collection,
		Key:        key,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/kvstore/v1beta1/collections/{{.Collection}}/records/{{.Key}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb map[string]interface{}
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	InsertRecord - Inserts a record into a collection.
	Parameters:
		collection: The name of the collection.
		body: Record to add to the collection, formatted as a JSON object.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) InsertRecord(collection string, body map[string]interface{}, resp ...*http.Response) (*Key, error) {
	pp := struct {
		Collection string
	}{
		Collection: collection,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/kvstore/v1beta1/collections/{{.Collection}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: body})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb Key
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	InsertRecords - Inserts multiple records in a single request.
	Parameters:
		collection: The name of the collection.
		requestBody: Array of records to insert.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) InsertRecords(collection string, requestBody []map[string]interface{}, resp ...*http.Response) ([]string, error) {
	pp := struct {
		Collection string
	}{
		Collection: collection,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/kvstore/v1beta1/collections/{{.Collection}}/batch`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: requestBody})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb []string
	err = util.ParseResponse(&rb, response)
	return rb, err
}

/*
	ListIndexes - Returns a list of all indexes on a collection.
	Parameters:
		collection: The name of the collection.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListIndexes(collection string, resp ...*http.Response) ([]IndexDefinition, error) {
	pp := struct {
		Collection string
	}{
		Collection: collection,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/kvstore/v1beta1/collections/{{.Collection}}/indexes`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb []IndexDefinition
	err = util.ParseResponse(&rb, response)
	return rb, err
}

/*
	ListRecords - Returns a list of records in a collection with basic filtering, sorting, pagination and field projection.
	Use key-value query parameters to filter fields. Fields are implicitly ANDed and values for the same field are implicitly ORed.
	Parameters:
		collection: The name of the collection.
		query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListRecords(collection string, query *ListRecordsQueryParams, resp ...*http.Response) ([]map[string]interface{}, error) {
	values := util.ParseURLParams(query)
	pp := struct {
		Collection string
	}{
		Collection: collection,
	}
	u, err := s.Client.BuildURLFromPathParams(values, serviceCluster, `/kvstore/v1beta1/collections/{{.Collection}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb []map[string]interface{}
	err = util.ParseResponse(&rb, response)
	return rb, err
}

/*
	Ping - Returns the health status from the database.
	Parameters:
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) Ping(resp ...*http.Response) (*PingResponse, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/kvstore/v1beta1/ping`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb PingResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	PutRecord - Updates the record with a given key, either by inserting or replacing the record.
	Parameters:
		collection: The name of the collection.
		key: The key of the record.
		body: Record to add to the collection, formatted as a JSON object.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) PutRecord(collection string, key string, body map[string]interface{}, resp ...*http.Response) (*Key, error) {
	pp := struct {
		Collection string
		Key        string
	}{
		Collection: collection,
		Key:        key,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/kvstore/v1beta1/collections/{{.Collection}}/records/{{.Key}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Put(services.RequestParams{URL: u, Body: body})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb Key
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	QueryRecords - Returns a list of query records in a collection.
	Parameters:
		collection: The name of the collection.
		query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) QueryRecords(collection string, query *QueryRecordsQueryParams, resp ...*http.Response) ([]map[string]interface{}, error) {
	values := util.ParseURLParams(query)
	pp := struct {
		Collection string
	}{
		Collection: collection,
	}
	u, err := s.Client.BuildURLFromPathParams(values, serviceCluster, `/kvstore/v1beta1/collections/{{.Collection}}/query`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb []map[string]interface{}
	err = util.ParseResponse(&rb, response)
	return rb, err
}
