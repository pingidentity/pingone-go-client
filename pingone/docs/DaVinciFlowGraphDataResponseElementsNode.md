# DaVinciFlowGraphDataResponseElementsNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**DaVinciFlowGraphDataResponseElementsNodeData**](DaVinciFlowGraphDataResponseElementsNodeData.md) |  | 
**Grabbable** | **bool** |  | 
**Group** | **string** |  | 
**Locked** | **bool** |  | 
**Pannable** | **bool** |  | 
**Position** | [**DaVinciFlowGraphDataResponseElementsNodePosition**](DaVinciFlowGraphDataResponseElementsNodePosition.md) |  | 
**Removed** | **bool** |  | 
**Selectable** | **bool** |  | 
**Selected** | **bool** |  | 
**Classes** | Pointer to **string** |  | [optional] 

## Methods

### NewDaVinciFlowGraphDataResponseElementsNode

`func NewDaVinciFlowGraphDataResponseElementsNode(data DaVinciFlowGraphDataResponseElementsNodeData, grabbable bool, group string, locked bool, pannable bool, position DaVinciFlowGraphDataResponseElementsNodePosition, removed bool, selectable bool, selected bool, ) *DaVinciFlowGraphDataResponseElementsNode`

NewDaVinciFlowGraphDataResponseElementsNode instantiates a new DaVinciFlowGraphDataResponseElementsNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciFlowGraphDataResponseElementsNodeWithDefaults

`func NewDaVinciFlowGraphDataResponseElementsNodeWithDefaults() *DaVinciFlowGraphDataResponseElementsNode`

NewDaVinciFlowGraphDataResponseElementsNodeWithDefaults instantiates a new DaVinciFlowGraphDataResponseElementsNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DaVinciFlowGraphDataResponseElementsNode) GetData() DaVinciFlowGraphDataResponseElementsNodeData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DaVinciFlowGraphDataResponseElementsNode) GetDataOk() (*DaVinciFlowGraphDataResponseElementsNodeData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DaVinciFlowGraphDataResponseElementsNode) SetData(v DaVinciFlowGraphDataResponseElementsNodeData)`

SetData sets Data field to given value.


### GetGrabbable

`func (o *DaVinciFlowGraphDataResponseElementsNode) GetGrabbable() bool`

GetGrabbable returns the Grabbable field if non-nil, zero value otherwise.

### GetGrabbableOk

`func (o *DaVinciFlowGraphDataResponseElementsNode) GetGrabbableOk() (*bool, bool)`

GetGrabbableOk returns a tuple with the Grabbable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrabbable

`func (o *DaVinciFlowGraphDataResponseElementsNode) SetGrabbable(v bool)`

SetGrabbable sets Grabbable field to given value.


### GetGroup

`func (o *DaVinciFlowGraphDataResponseElementsNode) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DaVinciFlowGraphDataResponseElementsNode) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DaVinciFlowGraphDataResponseElementsNode) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetLocked

`func (o *DaVinciFlowGraphDataResponseElementsNode) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *DaVinciFlowGraphDataResponseElementsNode) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *DaVinciFlowGraphDataResponseElementsNode) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetPannable

`func (o *DaVinciFlowGraphDataResponseElementsNode) GetPannable() bool`

GetPannable returns the Pannable field if non-nil, zero value otherwise.

### GetPannableOk

`func (o *DaVinciFlowGraphDataResponseElementsNode) GetPannableOk() (*bool, bool)`

GetPannableOk returns a tuple with the Pannable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPannable

`func (o *DaVinciFlowGraphDataResponseElementsNode) SetPannable(v bool)`

SetPannable sets Pannable field to given value.


### GetPosition

`func (o *DaVinciFlowGraphDataResponseElementsNode) GetPosition() DaVinciFlowGraphDataResponseElementsNodePosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *DaVinciFlowGraphDataResponseElementsNode) GetPositionOk() (*DaVinciFlowGraphDataResponseElementsNodePosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *DaVinciFlowGraphDataResponseElementsNode) SetPosition(v DaVinciFlowGraphDataResponseElementsNodePosition)`

SetPosition sets Position field to given value.


### GetRemoved

`func (o *DaVinciFlowGraphDataResponseElementsNode) GetRemoved() bool`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *DaVinciFlowGraphDataResponseElementsNode) GetRemovedOk() (*bool, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *DaVinciFlowGraphDataResponseElementsNode) SetRemoved(v bool)`

SetRemoved sets Removed field to given value.


### GetSelectable

`func (o *DaVinciFlowGraphDataResponseElementsNode) GetSelectable() bool`

GetSelectable returns the Selectable field if non-nil, zero value otherwise.

### GetSelectableOk

`func (o *DaVinciFlowGraphDataResponseElementsNode) GetSelectableOk() (*bool, bool)`

GetSelectableOk returns a tuple with the Selectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectable

`func (o *DaVinciFlowGraphDataResponseElementsNode) SetSelectable(v bool)`

SetSelectable sets Selectable field to given value.


### GetSelected

`func (o *DaVinciFlowGraphDataResponseElementsNode) GetSelected() bool`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *DaVinciFlowGraphDataResponseElementsNode) GetSelectedOk() (*bool, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *DaVinciFlowGraphDataResponseElementsNode) SetSelected(v bool)`

SetSelected sets Selected field to given value.


### GetClasses

`func (o *DaVinciFlowGraphDataResponseElementsNode) GetClasses() string`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *DaVinciFlowGraphDataResponseElementsNode) GetClassesOk() (*string, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasses

`func (o *DaVinciFlowGraphDataResponseElementsNode) SetClasses(v string)`

SetClasses sets Classes field to given value.

### HasClasses

`func (o *DaVinciFlowGraphDataResponseElementsNode) HasClasses() bool`

HasClasses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


