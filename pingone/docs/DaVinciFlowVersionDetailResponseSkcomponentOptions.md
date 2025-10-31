# DaVinciFlowVersionDetailResponseSkcomponentOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppSessionId** | Pointer to **string** |  | [optional] 
**Appid** | Pointer to **string** |  | [optional] 
**Appsecret** | Pointer to **string** |  | [optional] 
**Browsertoken** | Pointer to **string** |  | [optional] 
**ButtonImage** | Pointer to **string** |  | [optional] 
**ButtonImageClass** | Pointer to **string** |  | [optional] 
**ButtonImagePlacement** | Pointer to [**DaVinciFlowVersionDetailResponseSkcomponentOptionsButtonImagePlacement**](DaVinciFlowVersionDetailResponseSkcomponentOptionsButtonImagePlacement.md) |  | [optional] 
**ButtonType** | Pointer to [**DaVinciFlowVersionDetailResponseSkcomponentOptionsButtonType**](DaVinciFlowVersionDetailResponseSkcomponentOptionsButtonType.md) |  | [optional] 
**Cdn** | Pointer to **string** |  | [optional] 
**Class** | Pointer to **string** |  | [optional] 
**CollectBehavioralData** | Pointer to **bool** |  | [optional] [default to true]
**ContentType** | Pointer to [**DaVinciFlowVersionDetailResponseSkcomponentOptionsContentType**](DaVinciFlowVersionDetailResponseSkcomponentOptionsContentType.md) |  | [optional] 
**CreateP1User** | Pointer to **string** |  | [optional] 
**CustomLoadingIndicator** | Pointer to **string** |  | [optional] 
**CustomLoadingIndicatorClass** | Pointer to **string** |  | [optional] 
**DefaultLoadingColor** | Pointer to **string** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**ErrorId** | Pointer to **string** |  | [optional] 
**Form** | Pointer to **string** |  | [optional] 
**Height** | Pointer to **string** |  | [optional] [default to "600px"]
**Host** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IdentityProviderId** | Pointer to **string** |  | [optional] 
**IdentityProviderIdEntry** | Pointer to **string** |  | [optional] 
**IdentityProviderType** | Pointer to **string** |  | [optional] 
**IdpConnector** | Pointer to **string** |  | [optional] 
**ImgUrl** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**MessageClass** | Pointer to **string** |  | [optional] 
**MicUrl** | Pointer to **string** |  | [optional] [default to "https://devsdk.singularkey.com/react-mic/latest/react-mic.js"]
**OtpInput** | Pointer to **string** |  | [optional] 
**PollChallengeStatus** | Pointer to **bool** |  | [optional] [default to true]
**PollInterval** | Pointer to **int32** |  | [optional] [default to 2000]
**PollRetries** | Pointer to **int32** |  | [optional] [default to 60]
**PopulationId** | Pointer to **string** |  | [optional] 
**PopulationIdEntry** | Pointer to **string** |  | [optional] 
**PreviewType** | Pointer to [**DaVinciFlowVersionDetailResponseSkcomponentOptionsPreviewType**](DaVinciFlowVersionDetailResponseSkcomponentOptionsPreviewType.md) |  | [optional] [default to DAVINCIFLOWVERSIONDETAILRESPONSESKCOMPONENTOPTIONSPREVIEWTYPE_IMAGE]
**Propertyname** | Pointer to **string** |  | [optional] 
**ReturnUrl** | Pointer to **string** |  | [optional] 
**ShowPreview** | Pointer to [**DaVinciFlowVersionDetailResponseSkcomponentOptionsShowPreview**](DaVinciFlowVersionDetailResponseSkcomponentOptionsShowPreview.md) |  | [optional] [default to DAVINCIFLOWVERSIONDETAILRESPONSESKCOMPONENTOPTIONSSHOWPREVIEW_YES]
**SiteKey** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Width** | Pointer to **string** |  | [optional] [default to "600px"]

## Methods

### NewDaVinciFlowVersionDetailResponseSkcomponentOptions

`func NewDaVinciFlowVersionDetailResponseSkcomponentOptions() *DaVinciFlowVersionDetailResponseSkcomponentOptions`

NewDaVinciFlowVersionDetailResponseSkcomponentOptions instantiates a new DaVinciFlowVersionDetailResponseSkcomponentOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciFlowVersionDetailResponseSkcomponentOptionsWithDefaults

`func NewDaVinciFlowVersionDetailResponseSkcomponentOptionsWithDefaults() *DaVinciFlowVersionDetailResponseSkcomponentOptions`

NewDaVinciFlowVersionDetailResponseSkcomponentOptionsWithDefaults instantiates a new DaVinciFlowVersionDetailResponseSkcomponentOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppSessionId

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetAppSessionId() string`

GetAppSessionId returns the AppSessionId field if non-nil, zero value otherwise.

### GetAppSessionIdOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetAppSessionIdOk() (*string, bool)`

GetAppSessionIdOk returns a tuple with the AppSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSessionId

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetAppSessionId(v string)`

SetAppSessionId sets AppSessionId field to given value.

### HasAppSessionId

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasAppSessionId() bool`

HasAppSessionId returns a boolean if a field has been set.

### GetAppid

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetAppid() string`

GetAppid returns the Appid field if non-nil, zero value otherwise.

### GetAppidOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetAppidOk() (*string, bool)`

GetAppidOk returns a tuple with the Appid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppid

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetAppid(v string)`

SetAppid sets Appid field to given value.

### HasAppid

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasAppid() bool`

HasAppid returns a boolean if a field has been set.

### GetAppsecret

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetAppsecret() string`

GetAppsecret returns the Appsecret field if non-nil, zero value otherwise.

### GetAppsecretOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetAppsecretOk() (*string, bool)`

GetAppsecretOk returns a tuple with the Appsecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppsecret

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetAppsecret(v string)`

SetAppsecret sets Appsecret field to given value.

### HasAppsecret

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasAppsecret() bool`

HasAppsecret returns a boolean if a field has been set.

### GetBrowsertoken

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetBrowsertoken() string`

GetBrowsertoken returns the Browsertoken field if non-nil, zero value otherwise.

### GetBrowsertokenOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetBrowsertokenOk() (*string, bool)`

GetBrowsertokenOk returns a tuple with the Browsertoken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowsertoken

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetBrowsertoken(v string)`

SetBrowsertoken sets Browsertoken field to given value.

### HasBrowsertoken

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasBrowsertoken() bool`

HasBrowsertoken returns a boolean if a field has been set.

### GetButtonImage

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetButtonImage() string`

GetButtonImage returns the ButtonImage field if non-nil, zero value otherwise.

### GetButtonImageOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetButtonImageOk() (*string, bool)`

GetButtonImageOk returns a tuple with the ButtonImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonImage

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetButtonImage(v string)`

SetButtonImage sets ButtonImage field to given value.

### HasButtonImage

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasButtonImage() bool`

HasButtonImage returns a boolean if a field has been set.

### GetButtonImageClass

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetButtonImageClass() string`

GetButtonImageClass returns the ButtonImageClass field if non-nil, zero value otherwise.

### GetButtonImageClassOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetButtonImageClassOk() (*string, bool)`

GetButtonImageClassOk returns a tuple with the ButtonImageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonImageClass

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetButtonImageClass(v string)`

SetButtonImageClass sets ButtonImageClass field to given value.

### HasButtonImageClass

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasButtonImageClass() bool`

HasButtonImageClass returns a boolean if a field has been set.

### GetButtonImagePlacement

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetButtonImagePlacement() DaVinciFlowVersionDetailResponseSkcomponentOptionsButtonImagePlacement`

GetButtonImagePlacement returns the ButtonImagePlacement field if non-nil, zero value otherwise.

### GetButtonImagePlacementOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetButtonImagePlacementOk() (*DaVinciFlowVersionDetailResponseSkcomponentOptionsButtonImagePlacement, bool)`

GetButtonImagePlacementOk returns a tuple with the ButtonImagePlacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonImagePlacement

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetButtonImagePlacement(v DaVinciFlowVersionDetailResponseSkcomponentOptionsButtonImagePlacement)`

SetButtonImagePlacement sets ButtonImagePlacement field to given value.

### HasButtonImagePlacement

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasButtonImagePlacement() bool`

HasButtonImagePlacement returns a boolean if a field has been set.

### GetButtonType

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetButtonType() DaVinciFlowVersionDetailResponseSkcomponentOptionsButtonType`

GetButtonType returns the ButtonType field if non-nil, zero value otherwise.

### GetButtonTypeOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetButtonTypeOk() (*DaVinciFlowVersionDetailResponseSkcomponentOptionsButtonType, bool)`

GetButtonTypeOk returns a tuple with the ButtonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonType

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetButtonType(v DaVinciFlowVersionDetailResponseSkcomponentOptionsButtonType)`

SetButtonType sets ButtonType field to given value.

### HasButtonType

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasButtonType() bool`

HasButtonType returns a boolean if a field has been set.

### GetCdn

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetCdn() string`

GetCdn returns the Cdn field if non-nil, zero value otherwise.

### GetCdnOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetCdnOk() (*string, bool)`

GetCdnOk returns a tuple with the Cdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdn

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetCdn(v string)`

SetCdn sets Cdn field to given value.

### HasCdn

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasCdn() bool`

HasCdn returns a boolean if a field has been set.

### GetClass

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetCollectBehavioralData

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetCollectBehavioralData() bool`

GetCollectBehavioralData returns the CollectBehavioralData field if non-nil, zero value otherwise.

### GetCollectBehavioralDataOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetCollectBehavioralDataOk() (*bool, bool)`

GetCollectBehavioralDataOk returns a tuple with the CollectBehavioralData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectBehavioralData

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetCollectBehavioralData(v bool)`

SetCollectBehavioralData sets CollectBehavioralData field to given value.

### HasCollectBehavioralData

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasCollectBehavioralData() bool`

HasCollectBehavioralData returns a boolean if a field has been set.

### GetContentType

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetContentType() DaVinciFlowVersionDetailResponseSkcomponentOptionsContentType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetContentTypeOk() (*DaVinciFlowVersionDetailResponseSkcomponentOptionsContentType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetContentType(v DaVinciFlowVersionDetailResponseSkcomponentOptionsContentType)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetCreateP1User

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetCreateP1User() string`

GetCreateP1User returns the CreateP1User field if non-nil, zero value otherwise.

### GetCreateP1UserOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetCreateP1UserOk() (*string, bool)`

GetCreateP1UserOk returns a tuple with the CreateP1User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateP1User

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetCreateP1User(v string)`

SetCreateP1User sets CreateP1User field to given value.

### HasCreateP1User

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasCreateP1User() bool`

HasCreateP1User returns a boolean if a field has been set.

### GetCustomLoadingIndicator

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetCustomLoadingIndicator() string`

GetCustomLoadingIndicator returns the CustomLoadingIndicator field if non-nil, zero value otherwise.

### GetCustomLoadingIndicatorOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetCustomLoadingIndicatorOk() (*string, bool)`

GetCustomLoadingIndicatorOk returns a tuple with the CustomLoadingIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLoadingIndicator

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetCustomLoadingIndicator(v string)`

SetCustomLoadingIndicator sets CustomLoadingIndicator field to given value.

### HasCustomLoadingIndicator

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasCustomLoadingIndicator() bool`

HasCustomLoadingIndicator returns a boolean if a field has been set.

### GetCustomLoadingIndicatorClass

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetCustomLoadingIndicatorClass() string`

GetCustomLoadingIndicatorClass returns the CustomLoadingIndicatorClass field if non-nil, zero value otherwise.

### GetCustomLoadingIndicatorClassOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetCustomLoadingIndicatorClassOk() (*string, bool)`

GetCustomLoadingIndicatorClassOk returns a tuple with the CustomLoadingIndicatorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLoadingIndicatorClass

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetCustomLoadingIndicatorClass(v string)`

SetCustomLoadingIndicatorClass sets CustomLoadingIndicatorClass field to given value.

### HasCustomLoadingIndicatorClass

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasCustomLoadingIndicatorClass() bool`

HasCustomLoadingIndicatorClass returns a boolean if a field has been set.

### GetDefaultLoadingColor

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetDefaultLoadingColor() string`

GetDefaultLoadingColor returns the DefaultLoadingColor field if non-nil, zero value otherwise.

### GetDefaultLoadingColorOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetDefaultLoadingColorOk() (*string, bool)`

GetDefaultLoadingColorOk returns a tuple with the DefaultLoadingColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLoadingColor

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetDefaultLoadingColor(v string)`

SetDefaultLoadingColor sets DefaultLoadingColor field to given value.

### HasDefaultLoadingColor

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasDefaultLoadingColor() bool`

HasDefaultLoadingColor returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetErrorId

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetErrorId() string`

GetErrorId returns the ErrorId field if non-nil, zero value otherwise.

### GetErrorIdOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetErrorIdOk() (*string, bool)`

GetErrorIdOk returns a tuple with the ErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorId

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetErrorId(v string)`

SetErrorId sets ErrorId field to given value.

### HasErrorId

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasErrorId() bool`

HasErrorId returns a boolean if a field has been set.

### GetForm

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetForm(v string)`

SetForm sets Form field to given value.

### HasForm

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetHeight

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetHeight() string`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetHeightOk() (*string, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetHeight(v string)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetHost

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityProviderId

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetIdentityProviderId() string`

GetIdentityProviderId returns the IdentityProviderId field if non-nil, zero value otherwise.

### GetIdentityProviderIdOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetIdentityProviderIdOk() (*string, bool)`

GetIdentityProviderIdOk returns a tuple with the IdentityProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderId

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetIdentityProviderId(v string)`

SetIdentityProviderId sets IdentityProviderId field to given value.

### HasIdentityProviderId

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasIdentityProviderId() bool`

HasIdentityProviderId returns a boolean if a field has been set.

### GetIdentityProviderIdEntry

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetIdentityProviderIdEntry() string`

GetIdentityProviderIdEntry returns the IdentityProviderIdEntry field if non-nil, zero value otherwise.

### GetIdentityProviderIdEntryOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetIdentityProviderIdEntryOk() (*string, bool)`

GetIdentityProviderIdEntryOk returns a tuple with the IdentityProviderIdEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderIdEntry

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetIdentityProviderIdEntry(v string)`

SetIdentityProviderIdEntry sets IdentityProviderIdEntry field to given value.

### HasIdentityProviderIdEntry

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasIdentityProviderIdEntry() bool`

HasIdentityProviderIdEntry returns a boolean if a field has been set.

### GetIdentityProviderType

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetIdentityProviderType() string`

GetIdentityProviderType returns the IdentityProviderType field if non-nil, zero value otherwise.

### GetIdentityProviderTypeOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetIdentityProviderTypeOk() (*string, bool)`

GetIdentityProviderTypeOk returns a tuple with the IdentityProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderType

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetIdentityProviderType(v string)`

SetIdentityProviderType sets IdentityProviderType field to given value.

### HasIdentityProviderType

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasIdentityProviderType() bool`

HasIdentityProviderType returns a boolean if a field has been set.

### GetIdpConnector

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetIdpConnector() string`

GetIdpConnector returns the IdpConnector field if non-nil, zero value otherwise.

### GetIdpConnectorOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetIdpConnectorOk() (*string, bool)`

GetIdpConnectorOk returns a tuple with the IdpConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpConnector

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetIdpConnector(v string)`

SetIdpConnector sets IdpConnector field to given value.

### HasIdpConnector

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasIdpConnector() bool`

HasIdpConnector returns a boolean if a field has been set.

### GetImgUrl

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetImgUrl() string`

GetImgUrl returns the ImgUrl field if non-nil, zero value otherwise.

### GetImgUrlOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetImgUrlOk() (*string, bool)`

GetImgUrlOk returns a tuple with the ImgUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgUrl

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetImgUrl(v string)`

SetImgUrl sets ImgUrl field to given value.

### HasImgUrl

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasImgUrl() bool`

HasImgUrl returns a boolean if a field has been set.

### GetLabel

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMessageClass

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetMessageClass() string`

GetMessageClass returns the MessageClass field if non-nil, zero value otherwise.

### GetMessageClassOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetMessageClassOk() (*string, bool)`

GetMessageClassOk returns a tuple with the MessageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageClass

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetMessageClass(v string)`

SetMessageClass sets MessageClass field to given value.

### HasMessageClass

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasMessageClass() bool`

HasMessageClass returns a boolean if a field has been set.

### GetMicUrl

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetMicUrl() string`

GetMicUrl returns the MicUrl field if non-nil, zero value otherwise.

### GetMicUrlOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetMicUrlOk() (*string, bool)`

GetMicUrlOk returns a tuple with the MicUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicUrl

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetMicUrl(v string)`

SetMicUrl sets MicUrl field to given value.

### HasMicUrl

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasMicUrl() bool`

HasMicUrl returns a boolean if a field has been set.

### GetOtpInput

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetOtpInput() string`

GetOtpInput returns the OtpInput field if non-nil, zero value otherwise.

### GetOtpInputOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetOtpInputOk() (*string, bool)`

GetOtpInputOk returns a tuple with the OtpInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpInput

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetOtpInput(v string)`

SetOtpInput sets OtpInput field to given value.

### HasOtpInput

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasOtpInput() bool`

HasOtpInput returns a boolean if a field has been set.

### GetPollChallengeStatus

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetPollChallengeStatus() bool`

GetPollChallengeStatus returns the PollChallengeStatus field if non-nil, zero value otherwise.

### GetPollChallengeStatusOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetPollChallengeStatusOk() (*bool, bool)`

GetPollChallengeStatusOk returns a tuple with the PollChallengeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollChallengeStatus

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetPollChallengeStatus(v bool)`

SetPollChallengeStatus sets PollChallengeStatus field to given value.

### HasPollChallengeStatus

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasPollChallengeStatus() bool`

HasPollChallengeStatus returns a boolean if a field has been set.

### GetPollInterval

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetPollInterval() int32`

GetPollInterval returns the PollInterval field if non-nil, zero value otherwise.

### GetPollIntervalOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetPollIntervalOk() (*int32, bool)`

GetPollIntervalOk returns a tuple with the PollInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollInterval

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetPollInterval(v int32)`

SetPollInterval sets PollInterval field to given value.

### HasPollInterval

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasPollInterval() bool`

HasPollInterval returns a boolean if a field has been set.

### GetPollRetries

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetPollRetries() int32`

GetPollRetries returns the PollRetries field if non-nil, zero value otherwise.

### GetPollRetriesOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetPollRetriesOk() (*int32, bool)`

GetPollRetriesOk returns a tuple with the PollRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollRetries

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetPollRetries(v int32)`

SetPollRetries sets PollRetries field to given value.

### HasPollRetries

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasPollRetries() bool`

HasPollRetries returns a boolean if a field has been set.

### GetPopulationId

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetPopulationId() string`

GetPopulationId returns the PopulationId field if non-nil, zero value otherwise.

### GetPopulationIdOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetPopulationIdOk() (*string, bool)`

GetPopulationIdOk returns a tuple with the PopulationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulationId

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetPopulationId(v string)`

SetPopulationId sets PopulationId field to given value.

### HasPopulationId

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasPopulationId() bool`

HasPopulationId returns a boolean if a field has been set.

### GetPopulationIdEntry

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetPopulationIdEntry() string`

GetPopulationIdEntry returns the PopulationIdEntry field if non-nil, zero value otherwise.

### GetPopulationIdEntryOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetPopulationIdEntryOk() (*string, bool)`

GetPopulationIdEntryOk returns a tuple with the PopulationIdEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulationIdEntry

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetPopulationIdEntry(v string)`

SetPopulationIdEntry sets PopulationIdEntry field to given value.

### HasPopulationIdEntry

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasPopulationIdEntry() bool`

HasPopulationIdEntry returns a boolean if a field has been set.

### GetPreviewType

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetPreviewType() DaVinciFlowVersionDetailResponseSkcomponentOptionsPreviewType`

GetPreviewType returns the PreviewType field if non-nil, zero value otherwise.

### GetPreviewTypeOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetPreviewTypeOk() (*DaVinciFlowVersionDetailResponseSkcomponentOptionsPreviewType, bool)`

GetPreviewTypeOk returns a tuple with the PreviewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewType

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetPreviewType(v DaVinciFlowVersionDetailResponseSkcomponentOptionsPreviewType)`

SetPreviewType sets PreviewType field to given value.

### HasPreviewType

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasPreviewType() bool`

HasPreviewType returns a boolean if a field has been set.

### GetPropertyname

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetPropertyname() string`

GetPropertyname returns the Propertyname field if non-nil, zero value otherwise.

### GetPropertynameOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetPropertynameOk() (*string, bool)`

GetPropertynameOk returns a tuple with the Propertyname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyname

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetPropertyname(v string)`

SetPropertyname sets Propertyname field to given value.

### HasPropertyname

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasPropertyname() bool`

HasPropertyname returns a boolean if a field has been set.

### GetReturnUrl

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetShowPreview

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetShowPreview() DaVinciFlowVersionDetailResponseSkcomponentOptionsShowPreview`

GetShowPreview returns the ShowPreview field if non-nil, zero value otherwise.

### GetShowPreviewOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetShowPreviewOk() (*DaVinciFlowVersionDetailResponseSkcomponentOptionsShowPreview, bool)`

GetShowPreviewOk returns a tuple with the ShowPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPreview

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetShowPreview(v DaVinciFlowVersionDetailResponseSkcomponentOptionsShowPreview)`

SetShowPreview sets ShowPreview field to given value.

### HasShowPreview

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasShowPreview() bool`

HasShowPreview returns a boolean if a field has been set.

### GetSiteKey

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetSiteKey() string`

GetSiteKey returns the SiteKey field if non-nil, zero value otherwise.

### GetSiteKeyOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetSiteKeyOk() (*string, bool)`

GetSiteKeyOk returns a tuple with the SiteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteKey

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetSiteKey(v string)`

SetSiteKey sets SiteKey field to given value.

### HasSiteKey

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasSiteKey() bool`

HasSiteKey returns a boolean if a field has been set.

### GetUserId

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetValue

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetWidth

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetWidth() string`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) GetWidthOk() (*string, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) SetWidth(v string)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *DaVinciFlowVersionDetailResponseSkcomponentOptions) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


