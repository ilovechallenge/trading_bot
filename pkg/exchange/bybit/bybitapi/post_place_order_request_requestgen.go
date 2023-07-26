// Code generated by "requestgen -method POST -responseType .APIResponse -responseDataField Result -url /v5/order/create -type PostPlaceOrderRequest -responseDataType .PlaceOrderResponse"; DO NOT EDIT.

package bybitapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
)

func (p *PostPlaceOrderRequest) Category(category Category) *PostPlaceOrderRequest {
	p.category = category
	return p
}

func (p *PostPlaceOrderRequest) Symbol(symbol string) *PostPlaceOrderRequest {
	p.symbol = symbol
	return p
}

func (p *PostPlaceOrderRequest) Side(side Side) *PostPlaceOrderRequest {
	p.side = side
	return p
}

func (p *PostPlaceOrderRequest) OrderType(orderType OrderType) *PostPlaceOrderRequest {
	p.orderType = orderType
	return p
}

func (p *PostPlaceOrderRequest) Qty(qty string) *PostPlaceOrderRequest {
	p.qty = qty
	return p
}

func (p *PostPlaceOrderRequest) OrderLinkId(orderLinkId string) *PostPlaceOrderRequest {
	p.orderLinkId = orderLinkId
	return p
}

func (p *PostPlaceOrderRequest) TimeInForce(timeInForce TimeInForce) *PostPlaceOrderRequest {
	p.timeInForce = timeInForce
	return p
}

func (p *PostPlaceOrderRequest) IsLeverage(isLeverage bool) *PostPlaceOrderRequest {
	p.isLeverage = &isLeverage
	return p
}

func (p *PostPlaceOrderRequest) Price(price string) *PostPlaceOrderRequest {
	p.price = &price
	return p
}

func (p *PostPlaceOrderRequest) TriggerDirection(triggerDirection int) *PostPlaceOrderRequest {
	p.triggerDirection = &triggerDirection
	return p
}

func (p *PostPlaceOrderRequest) OrderFilter(orderFilter string) *PostPlaceOrderRequest {
	p.orderFilter = &orderFilter
	return p
}

func (p *PostPlaceOrderRequest) TriggerPrice(triggerPrice string) *PostPlaceOrderRequest {
	p.triggerPrice = &triggerPrice
	return p
}

func (p *PostPlaceOrderRequest) TriggerBy(triggerBy string) *PostPlaceOrderRequest {
	p.triggerBy = &triggerBy
	return p
}

func (p *PostPlaceOrderRequest) OrderIv(orderIv string) *PostPlaceOrderRequest {
	p.orderIv = &orderIv
	return p
}

func (p *PostPlaceOrderRequest) PositionIdx(positionIdx string) *PostPlaceOrderRequest {
	p.positionIdx = &positionIdx
	return p
}

func (p *PostPlaceOrderRequest) TakeProfit(takeProfit string) *PostPlaceOrderRequest {
	p.takeProfit = &takeProfit
	return p
}

func (p *PostPlaceOrderRequest) StopLoss(stopLoss string) *PostPlaceOrderRequest {
	p.stopLoss = &stopLoss
	return p
}

func (p *PostPlaceOrderRequest) TpTriggerBy(tpTriggerBy string) *PostPlaceOrderRequest {
	p.tpTriggerBy = &tpTriggerBy
	return p
}

func (p *PostPlaceOrderRequest) SlTriggerBy(slTriggerBy string) *PostPlaceOrderRequest {
	p.slTriggerBy = &slTriggerBy
	return p
}

func (p *PostPlaceOrderRequest) ReduceOnly(reduceOnly bool) *PostPlaceOrderRequest {
	p.reduceOnly = &reduceOnly
	return p
}

func (p *PostPlaceOrderRequest) CloseOnTrigger(closeOnTrigger bool) *PostPlaceOrderRequest {
	p.closeOnTrigger = &closeOnTrigger
	return p
}

func (p *PostPlaceOrderRequest) SmpType(smpType string) *PostPlaceOrderRequest {
	p.smpType = &smpType
	return p
}

func (p *PostPlaceOrderRequest) Mmp(mmp bool) *PostPlaceOrderRequest {
	p.mmp = &mmp
	return p
}

func (p *PostPlaceOrderRequest) TpslMode(tpslMode string) *PostPlaceOrderRequest {
	p.tpslMode = &tpslMode
	return p
}

func (p *PostPlaceOrderRequest) TpLimitPrice(tpLimitPrice string) *PostPlaceOrderRequest {
	p.tpLimitPrice = &tpLimitPrice
	return p
}

func (p *PostPlaceOrderRequest) SlLimitPrice(slLimitPrice string) *PostPlaceOrderRequest {
	p.slLimitPrice = &slLimitPrice
	return p
}

func (p *PostPlaceOrderRequest) TpOrderType(tpOrderType string) *PostPlaceOrderRequest {
	p.tpOrderType = &tpOrderType
	return p
}

func (p *PostPlaceOrderRequest) SlOrderType(slOrderType string) *PostPlaceOrderRequest {
	p.slOrderType = &slOrderType
	return p
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (p *PostPlaceOrderRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (p *PostPlaceOrderRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check category field -> json key category
	category := p.category

	// TEMPLATE check-valid-values
	switch category {
	case "spot":
		params["category"] = category

	default:
		return nil, fmt.Errorf("category value %v is invalid", category)

	}
	// END TEMPLATE check-valid-values

	// assign parameter of category
	params["category"] = category
	// check symbol field -> json key symbol
	symbol := p.symbol

	// assign parameter of symbol
	params["symbol"] = symbol
	// check side field -> json key side
	side := p.side

	// TEMPLATE check-valid-values
	switch side {
	case "Buy", "Sell":
		params["side"] = side

	default:
		return nil, fmt.Errorf("side value %v is invalid", side)

	}
	// END TEMPLATE check-valid-values

	// assign parameter of side
	params["side"] = side
	// check orderType field -> json key orderType
	orderType := p.orderType

	// TEMPLATE check-valid-values
	switch orderType {
	case "Market", "Limit":
		params["orderType"] = orderType

	default:
		return nil, fmt.Errorf("orderType value %v is invalid", orderType)

	}
	// END TEMPLATE check-valid-values

	// assign parameter of orderType
	params["orderType"] = orderType
	// check qty field -> json key qty
	qty := p.qty

	// assign parameter of qty
	params["qty"] = qty
	// check orderLinkId field -> json key orderLinkId
	orderLinkId := p.orderLinkId

	// assign parameter of orderLinkId
	params["orderLinkId"] = orderLinkId
	// check timeInForce field -> json key timeInForce
	timeInForce := p.timeInForce

	// TEMPLATE check-valid-values
	switch timeInForce {
	case TimeInForceGTC, TimeInForceIOC, TimeInForceFOK:
		params["timeInForce"] = timeInForce

	default:
		return nil, fmt.Errorf("timeInForce value %v is invalid", timeInForce)

	}
	// END TEMPLATE check-valid-values

	// assign parameter of timeInForce
	params["timeInForce"] = timeInForce
	// check isLeverage field -> json key isLeverage
	if p.isLeverage != nil {
		isLeverage := *p.isLeverage

		// assign parameter of isLeverage
		params["isLeverage"] = isLeverage
	} else {
	}
	// check price field -> json key price
	if p.price != nil {
		price := *p.price

		// assign parameter of price
		params["price"] = price
	} else {
	}
	// check triggerDirection field -> json key triggerDirection
	if p.triggerDirection != nil {
		triggerDirection := *p.triggerDirection

		// assign parameter of triggerDirection
		params["triggerDirection"] = triggerDirection
	} else {
	}
	// check orderFilter field -> json key orderFilter
	if p.orderFilter != nil {
		orderFilter := *p.orderFilter

		// assign parameter of orderFilter
		params["orderFilter"] = orderFilter
	} else {
	}
	// check triggerPrice field -> json key triggerPrice
	if p.triggerPrice != nil {
		triggerPrice := *p.triggerPrice

		// assign parameter of triggerPrice
		params["triggerPrice"] = triggerPrice
	} else {
	}
	// check triggerBy field -> json key triggerBy
	if p.triggerBy != nil {
		triggerBy := *p.triggerBy

		// assign parameter of triggerBy
		params["triggerBy"] = triggerBy
	} else {
	}
	// check orderIv field -> json key orderIv
	if p.orderIv != nil {
		orderIv := *p.orderIv

		// assign parameter of orderIv
		params["orderIv"] = orderIv
	} else {
	}
	// check positionIdx field -> json key positionIdx
	if p.positionIdx != nil {
		positionIdx := *p.positionIdx

		// assign parameter of positionIdx
		params["positionIdx"] = positionIdx
	} else {
	}
	// check takeProfit field -> json key takeProfit
	if p.takeProfit != nil {
		takeProfit := *p.takeProfit

		// assign parameter of takeProfit
		params["takeProfit"] = takeProfit
	} else {
	}
	// check stopLoss field -> json key stopLoss
	if p.stopLoss != nil {
		stopLoss := *p.stopLoss

		// assign parameter of stopLoss
		params["stopLoss"] = stopLoss
	} else {
	}
	// check tpTriggerBy field -> json key tpTriggerBy
	if p.tpTriggerBy != nil {
		tpTriggerBy := *p.tpTriggerBy

		// assign parameter of tpTriggerBy
		params["tpTriggerBy"] = tpTriggerBy
	} else {
	}
	// check slTriggerBy field -> json key slTriggerBy
	if p.slTriggerBy != nil {
		slTriggerBy := *p.slTriggerBy

		// assign parameter of slTriggerBy
		params["slTriggerBy"] = slTriggerBy
	} else {
	}
	// check reduceOnly field -> json key reduceOnly
	if p.reduceOnly != nil {
		reduceOnly := *p.reduceOnly

		// assign parameter of reduceOnly
		params["reduceOnly"] = reduceOnly
	} else {
	}
	// check closeOnTrigger field -> json key closeOnTrigger
	if p.closeOnTrigger != nil {
		closeOnTrigger := *p.closeOnTrigger

		// assign parameter of closeOnTrigger
		params["closeOnTrigger"] = closeOnTrigger
	} else {
	}
	// check smpType field -> json key smpType
	if p.smpType != nil {
		smpType := *p.smpType

		// assign parameter of smpType
		params["smpType"] = smpType
	} else {
	}
	// check mmp field -> json key mmp
	if p.mmp != nil {
		mmp := *p.mmp

		// assign parameter of mmp
		params["mmp"] = mmp
	} else {
	}
	// check tpslMode field -> json key tpslMode
	if p.tpslMode != nil {
		tpslMode := *p.tpslMode

		// assign parameter of tpslMode
		params["tpslMode"] = tpslMode
	} else {
	}
	// check tpLimitPrice field -> json key tpLimitPrice
	if p.tpLimitPrice != nil {
		tpLimitPrice := *p.tpLimitPrice

		// assign parameter of tpLimitPrice
		params["tpLimitPrice"] = tpLimitPrice
	} else {
	}
	// check slLimitPrice field -> json key slLimitPrice
	if p.slLimitPrice != nil {
		slLimitPrice := *p.slLimitPrice

		// assign parameter of slLimitPrice
		params["slLimitPrice"] = slLimitPrice
	} else {
	}
	// check tpOrderType field -> json key tpOrderType
	if p.tpOrderType != nil {
		tpOrderType := *p.tpOrderType

		// assign parameter of tpOrderType
		params["tpOrderType"] = tpOrderType
	} else {
	}
	// check slOrderType field -> json key slOrderType
	if p.slOrderType != nil {
		slOrderType := *p.slOrderType

		// assign parameter of slOrderType
		params["slOrderType"] = slOrderType
	} else {
	}

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (p *PostPlaceOrderRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := p.GetParameters()
	if err != nil {
		return query, err
	}

	for _k, _v := range params {
		if p.isVarSlice(_v) {
			p.iterateSlice(_v, func(it interface{}) {
				query.Add(_k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(_k, fmt.Sprintf("%v", _v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (p *PostPlaceOrderRequest) GetParametersJSON() ([]byte, error) {
	params, err := p.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (p *PostPlaceOrderRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

func (p *PostPlaceOrderRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (p *PostPlaceOrderRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (p *PostPlaceOrderRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (p *PostPlaceOrderRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := p.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for _k, _v := range params {
		slugs[_k] = fmt.Sprintf("%v", _v)
	}

	return slugs, nil
}

func (p *PostPlaceOrderRequest) Do(ctx context.Context) (*PlaceOrderResponse, error) {

	params, err := p.GetParameters()
	if err != nil {
		return nil, err
	}
	query := url.Values{}

	apiURL := "/v5/order/create"

	req, err := p.client.NewAuthenticatedRequest(ctx, "POST", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := p.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse APIResponse
	if err := response.DecodeJSON(&apiResponse); err != nil {
		return nil, err
	}
	var data PlaceOrderResponse
	if err := json.Unmarshal(apiResponse.Result, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
