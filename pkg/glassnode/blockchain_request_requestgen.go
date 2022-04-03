// Code generated by "requestgen -method GET -type BlockchainRequest -url /v1/metrics/blockchain/:metric -responseType Response"; DO NOT EDIT.

package glassnode

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
)

func (b *BlockchainRequest) SetAsset(Asset string) *BlockchainRequest {
	b.Asset = Asset
	return b
}

func (b *BlockchainRequest) SetSince(Since int64) *BlockchainRequest {
	b.Since = Since
	return b
}

func (b *BlockchainRequest) SetUntil(Until int64) *BlockchainRequest {
	b.Until = Until
	return b
}

func (b *BlockchainRequest) SetInterval(Interval Interval) *BlockchainRequest {
	b.Interval = Interval
	return b
}

func (b *BlockchainRequest) SetFormat(Format Format) *BlockchainRequest {
	b.Format = Format
	return b
}

func (b *BlockchainRequest) SetTimestampFormat(TimestampFormat string) *BlockchainRequest {
	b.TimestampFormat = TimestampFormat
	return b
}

func (b *BlockchainRequest) SetMetric(Metric string) *BlockchainRequest {
	b.Metric = Metric
	return b
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (b *BlockchainRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}
	// check Asset field -> json key a
	Asset := b.Asset

	// TEMPLATE check-required
	if len(Asset) == 0 {
		return nil, fmt.Errorf("a is required, empty string given")
	}
	// END TEMPLATE check-required

	// assign parameter of Asset
	params["a"] = Asset
	// check Since field -> json key s
	Since := b.Since

	// assign parameter of Since
	params["s"] = Since
	// check Until field -> json key u
	Until := b.Until

	// assign parameter of Until
	params["u"] = Until
	// check Interval field -> json key i
	Interval := b.Interval

	// assign parameter of Interval
	params["i"] = Interval
	// check Format field -> json key f
	Format := b.Format

	// assign parameter of Format
	params["f"] = Format
	// check TimestampFormat field -> json key timestamp_format
	TimestampFormat := b.TimestampFormat

	// assign parameter of TimestampFormat
	params["timestamp_format"] = TimestampFormat

	query := url.Values{}
	for k, v := range params {
		query.Add(k, fmt.Sprintf("%v", v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (b *BlockchainRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (b *BlockchainRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := b.GetParameters()
	if err != nil {
		return query, err
	}

	for k, v := range params {
		query.Add(k, fmt.Sprintf("%v", v))
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (b *BlockchainRequest) GetParametersJSON() ([]byte, error) {
	params, err := b.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (b *BlockchainRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check Metric field -> json key metric
	Metric := b.Metric

	// assign parameter of Metric
	params["metric"] = Metric

	return params, nil
}

func (b *BlockchainRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for k, v := range slugs {
		needleRE := regexp.MustCompile(":" + k + "\\b")
		url = needleRE.ReplaceAllString(url, v)
	}

	return url
}

func (b *BlockchainRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := b.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for k, v := range params {
		slugs[k] = fmt.Sprintf("%v", v)
	}

	return slugs, nil
}

func (b *BlockchainRequest) Do(ctx context.Context) (Response, error) {

	// no body params
	var params interface{}
	query, err := b.GetQueryParameters()
	if err != nil {
		return nil, err
	}

	apiURL := "/v1/metrics/blockchain/:metric"
	slugs, err := b.GetSlugsMap()
	if err != nil {
		return nil, err
	}

	apiURL = b.applySlugsToUrl(apiURL, slugs)

	req, err := b.Client.NewAuthenticatedRequest(ctx, "GET", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := b.Client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse Response
	if err := response.DecodeJSON(&apiResponse); err != nil {
		return nil, err
	}
	return apiResponse, nil
}
