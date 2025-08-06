# DaVinciFlowGraphDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoxSelectionEnabled** | **bool** |  | 
**Data** | **interface{}** |  | 
**Elements** | [**DaVinciFlowGraphDataResponseElements**](DaVinciFlowGraphDataResponseElements.md) |  | 
**MaxZoom** | **float32** |  | 
**MinZoom** | **float32** |  | 
**Pan** | [**DaVinciFlowGraphDataResponsePan**](DaVinciFlowGraphDataResponsePan.md) |  | 
**PanningEnabled** | **bool** |  | 
**Renderer** | [**DaVinciFlowGraphDataResponseRenderer**](DaVinciFlowGraphDataResponseRenderer.md) |  | 
**UserPanningEnabled** | **bool** |  | 
**UserZoomingEnabled** | **bool** |  | 
**Zoom** | **float32** |  | 
**ZoomingEnabled** | **bool** |  | 
**AllLinterErrors** | Pointer to [**[]DaVinciFlowGraphDataResponseAllLinterError**](DaVinciFlowGraphDataResponseAllLinterError.md) |  | [optional] 

## Methods

### NewDaVinciFlowGraphDataResponse

`func NewDaVinciFlowGraphDataResponse(boxSelectionEnabled bool, data interface{}, elements DaVinciFlowGraphDataResponseElements, maxZoom float32, minZoom float32, pan DaVinciFlowGraphDataResponsePan, panningEnabled bool, renderer DaVinciFlowGraphDataResponseRenderer, userPanningEnabled bool, userZoomingEnabled bool, zoom float32, zoomingEnabled bool, ) *DaVinciFlowGraphDataResponse`

NewDaVinciFlowGraphDataResponse instantiates a new DaVinciFlowGraphDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciFlowGraphDataResponseWithDefaults

`func NewDaVinciFlowGraphDataResponseWithDefaults() *DaVinciFlowGraphDataResponse`

NewDaVinciFlowGraphDataResponseWithDefaults instantiates a new DaVinciFlowGraphDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoxSelectionEnabled

`func (o *DaVinciFlowGraphDataResponse) GetBoxSelectionEnabled() bool`

GetBoxSelectionEnabled returns the BoxSelectionEnabled field if non-nil, zero value otherwise.

### GetBoxSelectionEnabledOk

`func (o *DaVinciFlowGraphDataResponse) GetBoxSelectionEnabledOk() (*bool, bool)`

GetBoxSelectionEnabledOk returns a tuple with the BoxSelectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxSelectionEnabled

`func (o *DaVinciFlowGraphDataResponse) SetBoxSelectionEnabled(v bool)`

SetBoxSelectionEnabled sets BoxSelectionEnabled field to given value.


### GetData

`func (o *DaVinciFlowGraphDataResponse) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DaVinciFlowGraphDataResponse) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DaVinciFlowGraphDataResponse) SetData(v interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *DaVinciFlowGraphDataResponse) SetDataNil(b bool)`

SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *DaVinciFlowGraphDataResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetElements

`func (o *DaVinciFlowGraphDataResponse) GetElements() DaVinciFlowGraphDataResponseElements`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *DaVinciFlowGraphDataResponse) GetElementsOk() (*DaVinciFlowGraphDataResponseElements, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *DaVinciFlowGraphDataResponse) SetElements(v DaVinciFlowGraphDataResponseElements)`

SetElements sets Elements field to given value.


### GetMaxZoom

`func (o *DaVinciFlowGraphDataResponse) GetMaxZoom() float32`

GetMaxZoom returns the MaxZoom field if non-nil, zero value otherwise.

### GetMaxZoomOk

`func (o *DaVinciFlowGraphDataResponse) GetMaxZoomOk() (*float32, bool)`

GetMaxZoomOk returns a tuple with the MaxZoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxZoom

`func (o *DaVinciFlowGraphDataResponse) SetMaxZoom(v float32)`

SetMaxZoom sets MaxZoom field to given value.


### GetMinZoom

`func (o *DaVinciFlowGraphDataResponse) GetMinZoom() float32`

GetMinZoom returns the MinZoom field if non-nil, zero value otherwise.

### GetMinZoomOk

`func (o *DaVinciFlowGraphDataResponse) GetMinZoomOk() (*float32, bool)`

GetMinZoomOk returns a tuple with the MinZoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinZoom

`func (o *DaVinciFlowGraphDataResponse) SetMinZoom(v float32)`

SetMinZoom sets MinZoom field to given value.


### GetPan

`func (o *DaVinciFlowGraphDataResponse) GetPan() DaVinciFlowGraphDataResponsePan`

GetPan returns the Pan field if non-nil, zero value otherwise.

### GetPanOk

`func (o *DaVinciFlowGraphDataResponse) GetPanOk() (*DaVinciFlowGraphDataResponsePan, bool)`

GetPanOk returns a tuple with the Pan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPan

`func (o *DaVinciFlowGraphDataResponse) SetPan(v DaVinciFlowGraphDataResponsePan)`

SetPan sets Pan field to given value.


### GetPanningEnabled

`func (o *DaVinciFlowGraphDataResponse) GetPanningEnabled() bool`

GetPanningEnabled returns the PanningEnabled field if non-nil, zero value otherwise.

### GetPanningEnabledOk

`func (o *DaVinciFlowGraphDataResponse) GetPanningEnabledOk() (*bool, bool)`

GetPanningEnabledOk returns a tuple with the PanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanningEnabled

`func (o *DaVinciFlowGraphDataResponse) SetPanningEnabled(v bool)`

SetPanningEnabled sets PanningEnabled field to given value.


### GetRenderer

`func (o *DaVinciFlowGraphDataResponse) GetRenderer() DaVinciFlowGraphDataResponseRenderer`

GetRenderer returns the Renderer field if non-nil, zero value otherwise.

### GetRendererOk

`func (o *DaVinciFlowGraphDataResponse) GetRendererOk() (*DaVinciFlowGraphDataResponseRenderer, bool)`

GetRendererOk returns a tuple with the Renderer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderer

`func (o *DaVinciFlowGraphDataResponse) SetRenderer(v DaVinciFlowGraphDataResponseRenderer)`

SetRenderer sets Renderer field to given value.


### GetUserPanningEnabled

`func (o *DaVinciFlowGraphDataResponse) GetUserPanningEnabled() bool`

GetUserPanningEnabled returns the UserPanningEnabled field if non-nil, zero value otherwise.

### GetUserPanningEnabledOk

`func (o *DaVinciFlowGraphDataResponse) GetUserPanningEnabledOk() (*bool, bool)`

GetUserPanningEnabledOk returns a tuple with the UserPanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPanningEnabled

`func (o *DaVinciFlowGraphDataResponse) SetUserPanningEnabled(v bool)`

SetUserPanningEnabled sets UserPanningEnabled field to given value.


### GetUserZoomingEnabled

`func (o *DaVinciFlowGraphDataResponse) GetUserZoomingEnabled() bool`

GetUserZoomingEnabled returns the UserZoomingEnabled field if non-nil, zero value otherwise.

### GetUserZoomingEnabledOk

`func (o *DaVinciFlowGraphDataResponse) GetUserZoomingEnabledOk() (*bool, bool)`

GetUserZoomingEnabledOk returns a tuple with the UserZoomingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserZoomingEnabled

`func (o *DaVinciFlowGraphDataResponse) SetUserZoomingEnabled(v bool)`

SetUserZoomingEnabled sets UserZoomingEnabled field to given value.


### GetZoom

`func (o *DaVinciFlowGraphDataResponse) GetZoom() float32`

GetZoom returns the Zoom field if non-nil, zero value otherwise.

### GetZoomOk

`func (o *DaVinciFlowGraphDataResponse) GetZoomOk() (*float32, bool)`

GetZoomOk returns a tuple with the Zoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoom

`func (o *DaVinciFlowGraphDataResponse) SetZoom(v float32)`

SetZoom sets Zoom field to given value.


### GetZoomingEnabled

`func (o *DaVinciFlowGraphDataResponse) GetZoomingEnabled() bool`

GetZoomingEnabled returns the ZoomingEnabled field if non-nil, zero value otherwise.

### GetZoomingEnabledOk

`func (o *DaVinciFlowGraphDataResponse) GetZoomingEnabledOk() (*bool, bool)`

GetZoomingEnabledOk returns a tuple with the ZoomingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoomingEnabled

`func (o *DaVinciFlowGraphDataResponse) SetZoomingEnabled(v bool)`

SetZoomingEnabled sets ZoomingEnabled field to given value.


### GetAllLinterErrors

`func (o *DaVinciFlowGraphDataResponse) GetAllLinterErrors() []DaVinciFlowGraphDataResponseAllLinterError`

GetAllLinterErrors returns the AllLinterErrors field if non-nil, zero value otherwise.

### GetAllLinterErrorsOk

`func (o *DaVinciFlowGraphDataResponse) GetAllLinterErrorsOk() (*[]DaVinciFlowGraphDataResponseAllLinterError, bool)`

GetAllLinterErrorsOk returns a tuple with the AllLinterErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLinterErrors

`func (o *DaVinciFlowGraphDataResponse) SetAllLinterErrors(v []DaVinciFlowGraphDataResponseAllLinterError)`

SetAllLinterErrors sets AllLinterErrors field to given value.

### HasAllLinterErrors

`func (o *DaVinciFlowGraphDataResponse) HasAllLinterErrors() bool`

HasAllLinterErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


