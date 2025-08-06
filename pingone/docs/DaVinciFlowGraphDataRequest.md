# DaVinciFlowGraphDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoxSelectionEnabled** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Elements** | Pointer to [**DaVinciFlowGraphDataRequestElements**](DaVinciFlowGraphDataRequestElements.md) |  | [optional] 
**MaxZoom** | Pointer to **float32** |  | [optional] 
**MinZoom** | Pointer to **float32** |  | [optional] 
**Pan** | Pointer to [**DaVinciFlowGraphDataRequestPan**](DaVinciFlowGraphDataRequestPan.md) |  | [optional] 
**PanningEnabled** | Pointer to **bool** |  | [optional] 
**Renderer** | Pointer to **map[string]interface{}** |  | [optional] 
**UserPanningEnabled** | Pointer to **bool** |  | [optional] 
**UserZoomingEnabled** | Pointer to **bool** |  | [optional] 
**Zoom** | Pointer to **int32** |  | [optional] 
**ZoomingEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewDaVinciFlowGraphDataRequest

`func NewDaVinciFlowGraphDataRequest() *DaVinciFlowGraphDataRequest`

NewDaVinciFlowGraphDataRequest instantiates a new DaVinciFlowGraphDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciFlowGraphDataRequestWithDefaults

`func NewDaVinciFlowGraphDataRequestWithDefaults() *DaVinciFlowGraphDataRequest`

NewDaVinciFlowGraphDataRequestWithDefaults instantiates a new DaVinciFlowGraphDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoxSelectionEnabled

`func (o *DaVinciFlowGraphDataRequest) GetBoxSelectionEnabled() bool`

GetBoxSelectionEnabled returns the BoxSelectionEnabled field if non-nil, zero value otherwise.

### GetBoxSelectionEnabledOk

`func (o *DaVinciFlowGraphDataRequest) GetBoxSelectionEnabledOk() (*bool, bool)`

GetBoxSelectionEnabledOk returns a tuple with the BoxSelectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxSelectionEnabled

`func (o *DaVinciFlowGraphDataRequest) SetBoxSelectionEnabled(v bool)`

SetBoxSelectionEnabled sets BoxSelectionEnabled field to given value.

### HasBoxSelectionEnabled

`func (o *DaVinciFlowGraphDataRequest) HasBoxSelectionEnabled() bool`

HasBoxSelectionEnabled returns a boolean if a field has been set.

### GetData

`func (o *DaVinciFlowGraphDataRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DaVinciFlowGraphDataRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DaVinciFlowGraphDataRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *DaVinciFlowGraphDataRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetElements

`func (o *DaVinciFlowGraphDataRequest) GetElements() DaVinciFlowGraphDataRequestElements`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *DaVinciFlowGraphDataRequest) GetElementsOk() (*DaVinciFlowGraphDataRequestElements, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *DaVinciFlowGraphDataRequest) SetElements(v DaVinciFlowGraphDataRequestElements)`

SetElements sets Elements field to given value.

### HasElements

`func (o *DaVinciFlowGraphDataRequest) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetMaxZoom

`func (o *DaVinciFlowGraphDataRequest) GetMaxZoom() float32`

GetMaxZoom returns the MaxZoom field if non-nil, zero value otherwise.

### GetMaxZoomOk

`func (o *DaVinciFlowGraphDataRequest) GetMaxZoomOk() (*float32, bool)`

GetMaxZoomOk returns a tuple with the MaxZoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxZoom

`func (o *DaVinciFlowGraphDataRequest) SetMaxZoom(v float32)`

SetMaxZoom sets MaxZoom field to given value.

### HasMaxZoom

`func (o *DaVinciFlowGraphDataRequest) HasMaxZoom() bool`

HasMaxZoom returns a boolean if a field has been set.

### GetMinZoom

`func (o *DaVinciFlowGraphDataRequest) GetMinZoom() float32`

GetMinZoom returns the MinZoom field if non-nil, zero value otherwise.

### GetMinZoomOk

`func (o *DaVinciFlowGraphDataRequest) GetMinZoomOk() (*float32, bool)`

GetMinZoomOk returns a tuple with the MinZoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinZoom

`func (o *DaVinciFlowGraphDataRequest) SetMinZoom(v float32)`

SetMinZoom sets MinZoom field to given value.

### HasMinZoom

`func (o *DaVinciFlowGraphDataRequest) HasMinZoom() bool`

HasMinZoom returns a boolean if a field has been set.

### GetPan

`func (o *DaVinciFlowGraphDataRequest) GetPan() DaVinciFlowGraphDataRequestPan`

GetPan returns the Pan field if non-nil, zero value otherwise.

### GetPanOk

`func (o *DaVinciFlowGraphDataRequest) GetPanOk() (*DaVinciFlowGraphDataRequestPan, bool)`

GetPanOk returns a tuple with the Pan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPan

`func (o *DaVinciFlowGraphDataRequest) SetPan(v DaVinciFlowGraphDataRequestPan)`

SetPan sets Pan field to given value.

### HasPan

`func (o *DaVinciFlowGraphDataRequest) HasPan() bool`

HasPan returns a boolean if a field has been set.

### GetPanningEnabled

`func (o *DaVinciFlowGraphDataRequest) GetPanningEnabled() bool`

GetPanningEnabled returns the PanningEnabled field if non-nil, zero value otherwise.

### GetPanningEnabledOk

`func (o *DaVinciFlowGraphDataRequest) GetPanningEnabledOk() (*bool, bool)`

GetPanningEnabledOk returns a tuple with the PanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanningEnabled

`func (o *DaVinciFlowGraphDataRequest) SetPanningEnabled(v bool)`

SetPanningEnabled sets PanningEnabled field to given value.

### HasPanningEnabled

`func (o *DaVinciFlowGraphDataRequest) HasPanningEnabled() bool`

HasPanningEnabled returns a boolean if a field has been set.

### GetRenderer

`func (o *DaVinciFlowGraphDataRequest) GetRenderer() map[string]interface{}`

GetRenderer returns the Renderer field if non-nil, zero value otherwise.

### GetRendererOk

`func (o *DaVinciFlowGraphDataRequest) GetRendererOk() (*map[string]interface{}, bool)`

GetRendererOk returns a tuple with the Renderer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderer

`func (o *DaVinciFlowGraphDataRequest) SetRenderer(v map[string]interface{})`

SetRenderer sets Renderer field to given value.

### HasRenderer

`func (o *DaVinciFlowGraphDataRequest) HasRenderer() bool`

HasRenderer returns a boolean if a field has been set.

### GetUserPanningEnabled

`func (o *DaVinciFlowGraphDataRequest) GetUserPanningEnabled() bool`

GetUserPanningEnabled returns the UserPanningEnabled field if non-nil, zero value otherwise.

### GetUserPanningEnabledOk

`func (o *DaVinciFlowGraphDataRequest) GetUserPanningEnabledOk() (*bool, bool)`

GetUserPanningEnabledOk returns a tuple with the UserPanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPanningEnabled

`func (o *DaVinciFlowGraphDataRequest) SetUserPanningEnabled(v bool)`

SetUserPanningEnabled sets UserPanningEnabled field to given value.

### HasUserPanningEnabled

`func (o *DaVinciFlowGraphDataRequest) HasUserPanningEnabled() bool`

HasUserPanningEnabled returns a boolean if a field has been set.

### GetUserZoomingEnabled

`func (o *DaVinciFlowGraphDataRequest) GetUserZoomingEnabled() bool`

GetUserZoomingEnabled returns the UserZoomingEnabled field if non-nil, zero value otherwise.

### GetUserZoomingEnabledOk

`func (o *DaVinciFlowGraphDataRequest) GetUserZoomingEnabledOk() (*bool, bool)`

GetUserZoomingEnabledOk returns a tuple with the UserZoomingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserZoomingEnabled

`func (o *DaVinciFlowGraphDataRequest) SetUserZoomingEnabled(v bool)`

SetUserZoomingEnabled sets UserZoomingEnabled field to given value.

### HasUserZoomingEnabled

`func (o *DaVinciFlowGraphDataRequest) HasUserZoomingEnabled() bool`

HasUserZoomingEnabled returns a boolean if a field has been set.

### GetZoom

`func (o *DaVinciFlowGraphDataRequest) GetZoom() int32`

GetZoom returns the Zoom field if non-nil, zero value otherwise.

### GetZoomOk

`func (o *DaVinciFlowGraphDataRequest) GetZoomOk() (*int32, bool)`

GetZoomOk returns a tuple with the Zoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoom

`func (o *DaVinciFlowGraphDataRequest) SetZoom(v int32)`

SetZoom sets Zoom field to given value.

### HasZoom

`func (o *DaVinciFlowGraphDataRequest) HasZoom() bool`

HasZoom returns a boolean if a field has been set.

### GetZoomingEnabled

`func (o *DaVinciFlowGraphDataRequest) GetZoomingEnabled() bool`

GetZoomingEnabled returns the ZoomingEnabled field if non-nil, zero value otherwise.

### GetZoomingEnabledOk

`func (o *DaVinciFlowGraphDataRequest) GetZoomingEnabledOk() (*bool, bool)`

GetZoomingEnabledOk returns a tuple with the ZoomingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoomingEnabled

`func (o *DaVinciFlowGraphDataRequest) SetZoomingEnabled(v bool)`

SetZoomingEnabled sets ZoomingEnabled field to given value.

### HasZoomingEnabled

`func (o *DaVinciFlowGraphDataRequest) HasZoomingEnabled() bool`

HasZoomingEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


