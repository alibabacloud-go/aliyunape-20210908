// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ApeInnerCommonApiRequest struct {
	// appName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// channel
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// endTime
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// lat
	Lat *string `json:"Lat,omitempty" xml:"Lat,omitempty"`
	// lon
	Lon *string `json:"Lon,omitempty" xml:"Lon,omitempty"`
	// pageNum
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// pageSize
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// spCode
	SpCode *string `json:"SpCode,omitempty" xml:"SpCode,omitempty"`
	// startTime
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// station
	Station *string `json:"Station,omitempty" xml:"Station,omitempty"`
}

func (s ApeInnerCommonApiRequest) String() string {
	return tea.Prettify(s)
}

func (s ApeInnerCommonApiRequest) GoString() string {
	return s.String()
}

func (s *ApeInnerCommonApiRequest) SetAppName(v string) *ApeInnerCommonApiRequest {
	s.AppName = &v
	return s
}

func (s *ApeInnerCommonApiRequest) SetChannel(v string) *ApeInnerCommonApiRequest {
	s.Channel = &v
	return s
}

func (s *ApeInnerCommonApiRequest) SetEndTime(v string) *ApeInnerCommonApiRequest {
	s.EndTime = &v
	return s
}

func (s *ApeInnerCommonApiRequest) SetLat(v string) *ApeInnerCommonApiRequest {
	s.Lat = &v
	return s
}

func (s *ApeInnerCommonApiRequest) SetLon(v string) *ApeInnerCommonApiRequest {
	s.Lon = &v
	return s
}

func (s *ApeInnerCommonApiRequest) SetPageNum(v int32) *ApeInnerCommonApiRequest {
	s.PageNum = &v
	return s
}

func (s *ApeInnerCommonApiRequest) SetPageSize(v int32) *ApeInnerCommonApiRequest {
	s.PageSize = &v
	return s
}

func (s *ApeInnerCommonApiRequest) SetSpCode(v string) *ApeInnerCommonApiRequest {
	s.SpCode = &v
	return s
}

func (s *ApeInnerCommonApiRequest) SetStartTime(v string) *ApeInnerCommonApiRequest {
	s.StartTime = &v
	return s
}

func (s *ApeInnerCommonApiRequest) SetStation(v string) *ApeInnerCommonApiRequest {
	s.Station = &v
	return s
}

type ApeInnerCommonApiResponseBody struct {
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// rt
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ApeInnerCommonApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApeInnerCommonApiResponseBody) GoString() string {
	return s.String()
}

func (s *ApeInnerCommonApiResponseBody) SetCode(v string) *ApeInnerCommonApiResponseBody {
	s.Code = &v
	return s
}

func (s *ApeInnerCommonApiResponseBody) SetData(v []map[string]interface{}) *ApeInnerCommonApiResponseBody {
	s.Data = v
	return s
}

func (s *ApeInnerCommonApiResponseBody) SetMessage(v string) *ApeInnerCommonApiResponseBody {
	s.Message = &v
	return s
}

func (s *ApeInnerCommonApiResponseBody) SetRequestId(v string) *ApeInnerCommonApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApeInnerCommonApiResponseBody) SetRt(v int64) *ApeInnerCommonApiResponseBody {
	s.Rt = &v
	return s
}

func (s *ApeInnerCommonApiResponseBody) SetSuccess(v bool) *ApeInnerCommonApiResponseBody {
	s.Success = &v
	return s
}

type ApeInnerCommonApiResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApeInnerCommonApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApeInnerCommonApiResponse) String() string {
	return tea.Prettify(s)
}

func (s ApeInnerCommonApiResponse) GoString() string {
	return s.String()
}

func (s *ApeInnerCommonApiResponse) SetHeaders(v map[string]*string) *ApeInnerCommonApiResponse {
	s.Headers = v
	return s
}

func (s *ApeInnerCommonApiResponse) SetBody(v *ApeInnerCommonApiResponseBody) *ApeInnerCommonApiResponse {
	s.Body = v
	return s
}

type ApeProvinceStationRefRequest struct {
	// adcode
	Adcode *int64 `json:"Adcode,omitempty" xml:"Adcode,omitempty"`
	// appName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// city
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// cnty
	Cnty *string `json:"Cnty,omitempty" xml:"Cnty,omitempty"`
	// country
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// offset
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// pageSize
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// provinceCode
	ProvinceCode *int64 `json:"ProvinceCode,omitempty" xml:"ProvinceCode,omitempty"`
	// provinceName
	ProvinceName *string `json:"ProvinceName,omitempty" xml:"ProvinceName,omitempty"`
	// stationName
	StationName *string `json:"StationName,omitempty" xml:"StationName,omitempty"`
}

func (s ApeProvinceStationRefRequest) String() string {
	return tea.Prettify(s)
}

func (s ApeProvinceStationRefRequest) GoString() string {
	return s.String()
}

func (s *ApeProvinceStationRefRequest) SetAdcode(v int64) *ApeProvinceStationRefRequest {
	s.Adcode = &v
	return s
}

func (s *ApeProvinceStationRefRequest) SetAppName(v string) *ApeProvinceStationRefRequest {
	s.AppName = &v
	return s
}

func (s *ApeProvinceStationRefRequest) SetCity(v string) *ApeProvinceStationRefRequest {
	s.City = &v
	return s
}

func (s *ApeProvinceStationRefRequest) SetCnty(v string) *ApeProvinceStationRefRequest {
	s.Cnty = &v
	return s
}

func (s *ApeProvinceStationRefRequest) SetCountry(v string) *ApeProvinceStationRefRequest {
	s.Country = &v
	return s
}

func (s *ApeProvinceStationRefRequest) SetOffset(v int32) *ApeProvinceStationRefRequest {
	s.Offset = &v
	return s
}

func (s *ApeProvinceStationRefRequest) SetPageSize(v int32) *ApeProvinceStationRefRequest {
	s.PageSize = &v
	return s
}

func (s *ApeProvinceStationRefRequest) SetProvinceCode(v int64) *ApeProvinceStationRefRequest {
	s.ProvinceCode = &v
	return s
}

func (s *ApeProvinceStationRefRequest) SetProvinceName(v string) *ApeProvinceStationRefRequest {
	s.ProvinceName = &v
	return s
}

func (s *ApeProvinceStationRefRequest) SetStationName(v string) *ApeProvinceStationRefRequest {
	s.StationName = &v
	return s
}

type ApeProvinceStationRefResponseBody struct {
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// rt
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ApeProvinceStationRefResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApeProvinceStationRefResponseBody) GoString() string {
	return s.String()
}

func (s *ApeProvinceStationRefResponseBody) SetCode(v string) *ApeProvinceStationRefResponseBody {
	s.Code = &v
	return s
}

func (s *ApeProvinceStationRefResponseBody) SetData(v interface{}) *ApeProvinceStationRefResponseBody {
	s.Data = v
	return s
}

func (s *ApeProvinceStationRefResponseBody) SetMessage(v string) *ApeProvinceStationRefResponseBody {
	s.Message = &v
	return s
}

func (s *ApeProvinceStationRefResponseBody) SetRequestId(v string) *ApeProvinceStationRefResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApeProvinceStationRefResponseBody) SetRt(v int64) *ApeProvinceStationRefResponseBody {
	s.Rt = &v
	return s
}

func (s *ApeProvinceStationRefResponseBody) SetSuccess(v bool) *ApeProvinceStationRefResponseBody {
	s.Success = &v
	return s
}

type ApeProvinceStationRefResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApeProvinceStationRefResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApeProvinceStationRefResponse) String() string {
	return tea.Prettify(s)
}

func (s ApeProvinceStationRefResponse) GoString() string {
	return s.String()
}

func (s *ApeProvinceStationRefResponse) SetHeaders(v map[string]*string) *ApeProvinceStationRefResponse {
	s.Headers = v
	return s
}

func (s *ApeProvinceStationRefResponse) SetBody(v *ApeProvinceStationRefResponseBody) *ApeProvinceStationRefResponse {
	s.Body = v
	return s
}

type HistoricalRequest struct {
	// endTime
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// ????????????--????????????--????????????--???????????????aliyunape_meteor12_public_cn-0ju2d2hh90b
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// pageNum
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// pageSize
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// startTime
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// ??????????????????????????????
	Station *string `json:"Station,omitempty" xml:"Station,omitempty"`
}

func (s HistoricalRequest) String() string {
	return tea.Prettify(s)
}

func (s HistoricalRequest) GoString() string {
	return s.String()
}

func (s *HistoricalRequest) SetEndTime(v string) *HistoricalRequest {
	s.EndTime = &v
	return s
}

func (s *HistoricalRequest) SetOrderId(v string) *HistoricalRequest {
	s.OrderId = &v
	return s
}

func (s *HistoricalRequest) SetPageNum(v int32) *HistoricalRequest {
	s.PageNum = &v
	return s
}

func (s *HistoricalRequest) SetPageSize(v int32) *HistoricalRequest {
	s.PageSize = &v
	return s
}

func (s *HistoricalRequest) SetStartTime(v string) *HistoricalRequest {
	s.StartTime = &v
	return s
}

func (s *HistoricalRequest) SetStation(v string) *HistoricalRequest {
	s.Station = &v
	return s
}

type HistoricalResponseBody struct {
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// rt
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s HistoricalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HistoricalResponseBody) GoString() string {
	return s.String()
}

func (s *HistoricalResponseBody) SetCode(v string) *HistoricalResponseBody {
	s.Code = &v
	return s
}

func (s *HistoricalResponseBody) SetData(v []map[string]interface{}) *HistoricalResponseBody {
	s.Data = v
	return s
}

func (s *HistoricalResponseBody) SetMessage(v string) *HistoricalResponseBody {
	s.Message = &v
	return s
}

func (s *HistoricalResponseBody) SetRequestId(v string) *HistoricalResponseBody {
	s.RequestId = &v
	return s
}

func (s *HistoricalResponseBody) SetRt(v int64) *HistoricalResponseBody {
	s.Rt = &v
	return s
}

func (s *HistoricalResponseBody) SetSuccess(v bool) *HistoricalResponseBody {
	s.Success = &v
	return s
}

type HistoricalResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *HistoricalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HistoricalResponse) String() string {
	return tea.Prettify(s)
}

func (s HistoricalResponse) GoString() string {
	return s.String()
}

func (s *HistoricalResponse) SetHeaders(v map[string]*string) *HistoricalResponse {
	s.Headers = v
	return s
}

func (s *HistoricalResponse) SetBody(v *HistoricalResponseBody) *HistoricalResponse {
	s.Body = v
	return s
}

type StationDayRequest struct {
	// ????????????--????????????--????????????--???????????????aliyunape_meteor12_public_cn-0ju2d2hh90b
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// ????????????????????????
	StartForecast *string `json:"StartForecast,omitempty" xml:"StartForecast,omitempty"`
	// ????????????????????????????????????
	Station *string `json:"Station,omitempty" xml:"Station,omitempty"`
}

func (s StationDayRequest) String() string {
	return tea.Prettify(s)
}

func (s StationDayRequest) GoString() string {
	return s.String()
}

func (s *StationDayRequest) SetOrderId(v string) *StationDayRequest {
	s.OrderId = &v
	return s
}

func (s *StationDayRequest) SetStartForecast(v string) *StationDayRequest {
	s.StartForecast = &v
	return s
}

func (s *StationDayRequest) SetStation(v string) *StationDayRequest {
	s.Station = &v
	return s
}

type StationDayResponseBody struct {
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// rt
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StationDayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StationDayResponseBody) GoString() string {
	return s.String()
}

func (s *StationDayResponseBody) SetCode(v string) *StationDayResponseBody {
	s.Code = &v
	return s
}

func (s *StationDayResponseBody) SetData(v []map[string]interface{}) *StationDayResponseBody {
	s.Data = v
	return s
}

func (s *StationDayResponseBody) SetMessage(v string) *StationDayResponseBody {
	s.Message = &v
	return s
}

func (s *StationDayResponseBody) SetRequestId(v string) *StationDayResponseBody {
	s.RequestId = &v
	return s
}

func (s *StationDayResponseBody) SetRt(v int64) *StationDayResponseBody {
	s.Rt = &v
	return s
}

func (s *StationDayResponseBody) SetSuccess(v bool) *StationDayResponseBody {
	s.Success = &v
	return s
}

type StationDayResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StationDayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StationDayResponse) String() string {
	return tea.Prettify(s)
}

func (s StationDayResponse) GoString() string {
	return s.String()
}

func (s *StationDayResponse) SetHeaders(v map[string]*string) *StationDayResponse {
	s.Headers = v
	return s
}

func (s *StationDayResponse) SetBody(v *StationDayResponseBody) *StationDayResponse {
	s.Body = v
	return s
}

type WeatherforecastRequest struct {
	// ?????????????????????15??N~59.95??N???
	Lat *string `json:"Lat,omitempty" xml:"Lat,omitempty"`
	// ?????????????????????70??E~139.96??E???
	Lon *string `json:"Lon,omitempty" xml:"Lon,omitempty"`
	// ????????????--????????????--????????????--???????????????aliyunape_meteor12_public_cn-0ju2d2hh90b
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// yyyymmdd080000???yyyymmdd200000
	StartForecast *string `json:"StartForecast,omitempty" xml:"StartForecast,omitempty"`
}

func (s WeatherforecastRequest) String() string {
	return tea.Prettify(s)
}

func (s WeatherforecastRequest) GoString() string {
	return s.String()
}

func (s *WeatherforecastRequest) SetLat(v string) *WeatherforecastRequest {
	s.Lat = &v
	return s
}

func (s *WeatherforecastRequest) SetLon(v string) *WeatherforecastRequest {
	s.Lon = &v
	return s
}

func (s *WeatherforecastRequest) SetOrderId(v string) *WeatherforecastRequest {
	s.OrderId = &v
	return s
}

func (s *WeatherforecastRequest) SetStartForecast(v string) *WeatherforecastRequest {
	s.StartForecast = &v
	return s
}

type WeatherforecastResponseBody struct {
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// rt
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WeatherforecastResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WeatherforecastResponseBody) GoString() string {
	return s.String()
}

func (s *WeatherforecastResponseBody) SetCode(v string) *WeatherforecastResponseBody {
	s.Code = &v
	return s
}

func (s *WeatherforecastResponseBody) SetData(v []map[string]interface{}) *WeatherforecastResponseBody {
	s.Data = v
	return s
}

func (s *WeatherforecastResponseBody) SetMessage(v string) *WeatherforecastResponseBody {
	s.Message = &v
	return s
}

func (s *WeatherforecastResponseBody) SetRequestId(v string) *WeatherforecastResponseBody {
	s.RequestId = &v
	return s
}

func (s *WeatherforecastResponseBody) SetRt(v int64) *WeatherforecastResponseBody {
	s.Rt = &v
	return s
}

func (s *WeatherforecastResponseBody) SetSuccess(v bool) *WeatherforecastResponseBody {
	s.Success = &v
	return s
}

type WeatherforecastResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *WeatherforecastResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s WeatherforecastResponse) String() string {
	return tea.Prettify(s)
}

func (s WeatherforecastResponse) GoString() string {
	return s.String()
}

func (s *WeatherforecastResponse) SetHeaders(v map[string]*string) *WeatherforecastResponse {
	s.Headers = v
	return s
}

func (s *WeatherforecastResponse) SetBody(v *WeatherforecastResponseBody) *WeatherforecastResponse {
	s.Body = v
	return s
}

type WeatherforecastTimeRequest struct {
	// 20210809090000
	CurHour *string `json:"CurHour,omitempty" xml:"CurHour,omitempty"`
	// ?????????????????????15??N~59.95??N
	Lat *string `json:"Lat,omitempty" xml:"Lat,omitempty"`
	// ?????????????????????70??E~139.96??E???
	Lon *string `json:"Lon,omitempty" xml:"Lon,omitempty"`
	// ????????????--????????????--????????????--???????????????aliyunape_meteor12_public_cn-0ju2d2hh90b
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s WeatherforecastTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s WeatherforecastTimeRequest) GoString() string {
	return s.String()
}

func (s *WeatherforecastTimeRequest) SetCurHour(v string) *WeatherforecastTimeRequest {
	s.CurHour = &v
	return s
}

func (s *WeatherforecastTimeRequest) SetLat(v string) *WeatherforecastTimeRequest {
	s.Lat = &v
	return s
}

func (s *WeatherforecastTimeRequest) SetLon(v string) *WeatherforecastTimeRequest {
	s.Lon = &v
	return s
}

func (s *WeatherforecastTimeRequest) SetOrderId(v string) *WeatherforecastTimeRequest {
	s.OrderId = &v
	return s
}

type WeatherforecastTimeResponseBody struct {
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// rt
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WeatherforecastTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WeatherforecastTimeResponseBody) GoString() string {
	return s.String()
}

func (s *WeatherforecastTimeResponseBody) SetCode(v string) *WeatherforecastTimeResponseBody {
	s.Code = &v
	return s
}

func (s *WeatherforecastTimeResponseBody) SetData(v []map[string]interface{}) *WeatherforecastTimeResponseBody {
	s.Data = v
	return s
}

func (s *WeatherforecastTimeResponseBody) SetMessage(v string) *WeatherforecastTimeResponseBody {
	s.Message = &v
	return s
}

func (s *WeatherforecastTimeResponseBody) SetRequestId(v string) *WeatherforecastTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *WeatherforecastTimeResponseBody) SetRt(v int64) *WeatherforecastTimeResponseBody {
	s.Rt = &v
	return s
}

func (s *WeatherforecastTimeResponseBody) SetSuccess(v bool) *WeatherforecastTimeResponseBody {
	s.Success = &v
	return s
}

type WeatherforecastTimeResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *WeatherforecastTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s WeatherforecastTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s WeatherforecastTimeResponse) GoString() string {
	return s.String()
}

func (s *WeatherforecastTimeResponse) SetHeaders(v map[string]*string) *WeatherforecastTimeResponse {
	s.Headers = v
	return s
}

func (s *WeatherforecastTimeResponse) SetBody(v *WeatherforecastTimeResponseBody) *WeatherforecastTimeResponse {
	s.Body = v
	return s
}

type WeathermonitorRequest struct {
	// ?????????????????? yyyymmddhh0000 ?????????????????????2021-08-16???????????????	20210817120000
	CurHour *string `json:"CurHour,omitempty" xml:"CurHour,omitempty"`
	// ????????????--????????????--????????????--???????????????aliyunape_meteor12_public_cn-0ju2d2hh90b
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// ??????
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// ????????????
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s WeathermonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s WeathermonitorRequest) GoString() string {
	return s.String()
}

func (s *WeathermonitorRequest) SetCurHour(v string) *WeathermonitorRequest {
	s.CurHour = &v
	return s
}

func (s *WeathermonitorRequest) SetOrderId(v string) *WeathermonitorRequest {
	s.OrderId = &v
	return s
}

func (s *WeathermonitorRequest) SetPageNum(v int32) *WeathermonitorRequest {
	s.PageNum = &v
	return s
}

func (s *WeathermonitorRequest) SetPageSize(v int32) *WeathermonitorRequest {
	s.PageSize = &v
	return s
}

type WeathermonitorResponseBody struct {
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// rt
	Rt *int64 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WeathermonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WeathermonitorResponseBody) GoString() string {
	return s.String()
}

func (s *WeathermonitorResponseBody) SetCode(v string) *WeathermonitorResponseBody {
	s.Code = &v
	return s
}

func (s *WeathermonitorResponseBody) SetData(v []map[string]interface{}) *WeathermonitorResponseBody {
	s.Data = v
	return s
}

func (s *WeathermonitorResponseBody) SetMessage(v string) *WeathermonitorResponseBody {
	s.Message = &v
	return s
}

func (s *WeathermonitorResponseBody) SetRequestId(v string) *WeathermonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *WeathermonitorResponseBody) SetRt(v int64) *WeathermonitorResponseBody {
	s.Rt = &v
	return s
}

func (s *WeathermonitorResponseBody) SetSuccess(v bool) *WeathermonitorResponseBody {
	s.Success = &v
	return s
}

type WeathermonitorResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *WeathermonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s WeathermonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s WeathermonitorResponse) GoString() string {
	return s.String()
}

func (s *WeathermonitorResponse) SetHeaders(v map[string]*string) *WeathermonitorResponse {
	s.Headers = v
	return s
}

func (s *WeathermonitorResponse) SetBody(v *WeathermonitorResponseBody) *WeathermonitorResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("aliyunape"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApeInnerCommonApiWithOptions(request *ApeInnerCommonApiRequest, runtime *util.RuntimeOptions) (_result *ApeInnerCommonApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApeInnerCommonApi"),
		Version:     tea.String("2021-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApeInnerCommonApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApeInnerCommonApi(request *ApeInnerCommonApiRequest) (_result *ApeInnerCommonApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApeInnerCommonApiResponse{}
	_body, _err := client.ApeInnerCommonApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApeProvinceStationRefWithOptions(request *ApeProvinceStationRefRequest, runtime *util.RuntimeOptions) (_result *ApeProvinceStationRefResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Adcode)) {
		query["Adcode"] = request.Adcode
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.City)) {
		query["City"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.Cnty)) {
		query["Cnty"] = request.Cnty
	}

	if !tea.BoolValue(util.IsUnset(request.Country)) {
		query["Country"] = request.Country
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["Offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProvinceCode)) {
		query["ProvinceCode"] = request.ProvinceCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProvinceName)) {
		query["ProvinceName"] = request.ProvinceName
	}

	if !tea.BoolValue(util.IsUnset(request.StationName)) {
		query["StationName"] = request.StationName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApeProvinceStationRef"),
		Version:     tea.String("2021-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApeProvinceStationRefResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApeProvinceStationRef(request *ApeProvinceStationRefRequest) (_result *ApeProvinceStationRefResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApeProvinceStationRefResponse{}
	_body, _err := client.ApeProvinceStationRefWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HistoricalWithOptions(request *HistoricalRequest, runtime *util.RuntimeOptions) (_result *HistoricalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Station)) {
		query["Station"] = request.Station
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Historical"),
		Version:     tea.String("2021-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &HistoricalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Historical(request *HistoricalRequest) (_result *HistoricalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HistoricalResponse{}
	_body, _err := client.HistoricalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StationDayWithOptions(request *StationDayRequest, runtime *util.RuntimeOptions) (_result *StationDayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.StartForecast)) {
		query["StartForecast"] = request.StartForecast
	}

	if !tea.BoolValue(util.IsUnset(request.Station)) {
		query["Station"] = request.Station
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StationDay"),
		Version:     tea.String("2021-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StationDayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StationDay(request *StationDayRequest) (_result *StationDayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StationDayResponse{}
	_body, _err := client.StationDayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WeatherforecastWithOptions(request *WeatherforecastRequest, runtime *util.RuntimeOptions) (_result *WeatherforecastResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lat)) {
		query["Lat"] = request.Lat
	}

	if !tea.BoolValue(util.IsUnset(request.Lon)) {
		query["Lon"] = request.Lon
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.StartForecast)) {
		query["StartForecast"] = request.StartForecast
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Weatherforecast"),
		Version:     tea.String("2021-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &WeatherforecastResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Weatherforecast(request *WeatherforecastRequest) (_result *WeatherforecastResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &WeatherforecastResponse{}
	_body, _err := client.WeatherforecastWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WeatherforecastTimeWithOptions(request *WeatherforecastTimeRequest, runtime *util.RuntimeOptions) (_result *WeatherforecastTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurHour)) {
		query["CurHour"] = request.CurHour
	}

	if !tea.BoolValue(util.IsUnset(request.Lat)) {
		query["Lat"] = request.Lat
	}

	if !tea.BoolValue(util.IsUnset(request.Lon)) {
		query["Lon"] = request.Lon
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("WeatherforecastTime"),
		Version:     tea.String("2021-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &WeatherforecastTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WeatherforecastTime(request *WeatherforecastTimeRequest) (_result *WeatherforecastTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &WeatherforecastTimeResponse{}
	_body, _err := client.WeatherforecastTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WeathermonitorWithOptions(request *WeathermonitorRequest, runtime *util.RuntimeOptions) (_result *WeathermonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurHour)) {
		query["CurHour"] = request.CurHour
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Weathermonitor"),
		Version:     tea.String("2021-09-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &WeathermonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Weathermonitor(request *WeathermonitorRequest) (_result *WeathermonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &WeathermonitorResponse{}
	_body, _err := client.WeathermonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
