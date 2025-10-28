# DaVinciFlowSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Csp** | Pointer to **string** |  | [optional] 
**Css** | Pointer to **string** |  | [optional] 
**CssLinks** | Pointer to **[]string** |  | [optional] 
**CustomErrorScreenBrandLogoUrl** | Pointer to **string** |  | [optional] 
**CustomErrorShowFooter** | Pointer to **bool** |  | [optional] 
**CustomFaviconLink** | Pointer to **string** |  | [optional] 
**CustomLogoURLSelection** | Pointer to **int32** |  | [optional] 
**CustomTitle** | Pointer to **string** |  | [optional] 
**DefaultErrorScreenBrandLogo** | Pointer to **bool** |  | [optional] 
**FlowHttpTimeoutInSeconds** | Pointer to **int32** |  | [optional] 
**FlowTimeoutInSeconds** | Pointer to **float32** |  | [optional] 
**IntermediateLoadingScreenCSS** | Pointer to [**DaVinciFlowSettingsResponseIntermediateLoadingScreenCSS**](DaVinciFlowSettingsResponseIntermediateLoadingScreenCSS.md) |  | [optional] 
**IntermediateLoadingScreenHTML** | Pointer to [**DaVinciFlowSettingsResponseIntermediateLoadingScreenHTML**](DaVinciFlowSettingsResponseIntermediateLoadingScreenHTML.md) |  | [optional] 
**JsCustomFlowPlayer** | Pointer to **string** |  | [optional] 
**JsLinks** | Pointer to [**[]DaVinciFlowSettingsResponseJsLink**](DaVinciFlowSettingsResponseJsLink.md) |  | [optional] 
**LogLevel** | Pointer to **int32** |  | [optional] 
**RequireAuthenticationToInitiate** | Pointer to **bool** |  | [optional] 
**ScrubSensitiveInfo** | Pointer to **bool** |  | [optional] 
**SensitiveInfoFields** | Pointer to **[]string** |  | [optional] 
**UseCSP** | Pointer to **bool** |  | [optional] 
**UseCustomCSS** | Pointer to **bool** |  | [optional] 
**UseCustomFlowPlayer** | Pointer to **bool** |  | [optional] 
**UseCustomScript** | Pointer to **bool** |  | [optional] 
**UseIntermediateLoadingScreen** | Pointer to **bool** |  | [optional] 
**ValidateOnSave** | Pointer to **bool** |  | [optional] 

## Methods

### NewDaVinciFlowSettingsResponse

`func NewDaVinciFlowSettingsResponse() *DaVinciFlowSettingsResponse`

NewDaVinciFlowSettingsResponse instantiates a new DaVinciFlowSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciFlowSettingsResponseWithDefaults

`func NewDaVinciFlowSettingsResponseWithDefaults() *DaVinciFlowSettingsResponse`

NewDaVinciFlowSettingsResponseWithDefaults instantiates a new DaVinciFlowSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsp

`func (o *DaVinciFlowSettingsResponse) GetCsp() string`

GetCsp returns the Csp field if non-nil, zero value otherwise.

### GetCspOk

`func (o *DaVinciFlowSettingsResponse) GetCspOk() (*string, bool)`

GetCspOk returns a tuple with the Csp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsp

`func (o *DaVinciFlowSettingsResponse) SetCsp(v string)`

SetCsp sets Csp field to given value.

### HasCsp

`func (o *DaVinciFlowSettingsResponse) HasCsp() bool`

HasCsp returns a boolean if a field has been set.

### GetCss

`func (o *DaVinciFlowSettingsResponse) GetCss() string`

GetCss returns the Css field if non-nil, zero value otherwise.

### GetCssOk

`func (o *DaVinciFlowSettingsResponse) GetCssOk() (*string, bool)`

GetCssOk returns a tuple with the Css field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCss

`func (o *DaVinciFlowSettingsResponse) SetCss(v string)`

SetCss sets Css field to given value.

### HasCss

`func (o *DaVinciFlowSettingsResponse) HasCss() bool`

HasCss returns a boolean if a field has been set.

### GetCssLinks

`func (o *DaVinciFlowSettingsResponse) GetCssLinks() []string`

GetCssLinks returns the CssLinks field if non-nil, zero value otherwise.

### GetCssLinksOk

`func (o *DaVinciFlowSettingsResponse) GetCssLinksOk() (*[]string, bool)`

GetCssLinksOk returns a tuple with the CssLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssLinks

`func (o *DaVinciFlowSettingsResponse) SetCssLinks(v []string)`

SetCssLinks sets CssLinks field to given value.

### HasCssLinks

`func (o *DaVinciFlowSettingsResponse) HasCssLinks() bool`

HasCssLinks returns a boolean if a field has been set.

### GetCustomErrorScreenBrandLogoUrl

`func (o *DaVinciFlowSettingsResponse) GetCustomErrorScreenBrandLogoUrl() string`

GetCustomErrorScreenBrandLogoUrl returns the CustomErrorScreenBrandLogoUrl field if non-nil, zero value otherwise.

### GetCustomErrorScreenBrandLogoUrlOk

`func (o *DaVinciFlowSettingsResponse) GetCustomErrorScreenBrandLogoUrlOk() (*string, bool)`

GetCustomErrorScreenBrandLogoUrlOk returns a tuple with the CustomErrorScreenBrandLogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomErrorScreenBrandLogoUrl

`func (o *DaVinciFlowSettingsResponse) SetCustomErrorScreenBrandLogoUrl(v string)`

SetCustomErrorScreenBrandLogoUrl sets CustomErrorScreenBrandLogoUrl field to given value.

### HasCustomErrorScreenBrandLogoUrl

`func (o *DaVinciFlowSettingsResponse) HasCustomErrorScreenBrandLogoUrl() bool`

HasCustomErrorScreenBrandLogoUrl returns a boolean if a field has been set.

### GetCustomErrorShowFooter

`func (o *DaVinciFlowSettingsResponse) GetCustomErrorShowFooter() bool`

GetCustomErrorShowFooter returns the CustomErrorShowFooter field if non-nil, zero value otherwise.

### GetCustomErrorShowFooterOk

`func (o *DaVinciFlowSettingsResponse) GetCustomErrorShowFooterOk() (*bool, bool)`

GetCustomErrorShowFooterOk returns a tuple with the CustomErrorShowFooter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomErrorShowFooter

`func (o *DaVinciFlowSettingsResponse) SetCustomErrorShowFooter(v bool)`

SetCustomErrorShowFooter sets CustomErrorShowFooter field to given value.

### HasCustomErrorShowFooter

`func (o *DaVinciFlowSettingsResponse) HasCustomErrorShowFooter() bool`

HasCustomErrorShowFooter returns a boolean if a field has been set.

### GetCustomFaviconLink

`func (o *DaVinciFlowSettingsResponse) GetCustomFaviconLink() string`

GetCustomFaviconLink returns the CustomFaviconLink field if non-nil, zero value otherwise.

### GetCustomFaviconLinkOk

`func (o *DaVinciFlowSettingsResponse) GetCustomFaviconLinkOk() (*string, bool)`

GetCustomFaviconLinkOk returns a tuple with the CustomFaviconLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFaviconLink

`func (o *DaVinciFlowSettingsResponse) SetCustomFaviconLink(v string)`

SetCustomFaviconLink sets CustomFaviconLink field to given value.

### HasCustomFaviconLink

`func (o *DaVinciFlowSettingsResponse) HasCustomFaviconLink() bool`

HasCustomFaviconLink returns a boolean if a field has been set.

### GetCustomLogoURLSelection

`func (o *DaVinciFlowSettingsResponse) GetCustomLogoURLSelection() int32`

GetCustomLogoURLSelection returns the CustomLogoURLSelection field if non-nil, zero value otherwise.

### GetCustomLogoURLSelectionOk

`func (o *DaVinciFlowSettingsResponse) GetCustomLogoURLSelectionOk() (*int32, bool)`

GetCustomLogoURLSelectionOk returns a tuple with the CustomLogoURLSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLogoURLSelection

`func (o *DaVinciFlowSettingsResponse) SetCustomLogoURLSelection(v int32)`

SetCustomLogoURLSelection sets CustomLogoURLSelection field to given value.

### HasCustomLogoURLSelection

`func (o *DaVinciFlowSettingsResponse) HasCustomLogoURLSelection() bool`

HasCustomLogoURLSelection returns a boolean if a field has been set.

### GetCustomTitle

`func (o *DaVinciFlowSettingsResponse) GetCustomTitle() string`

GetCustomTitle returns the CustomTitle field if non-nil, zero value otherwise.

### GetCustomTitleOk

`func (o *DaVinciFlowSettingsResponse) GetCustomTitleOk() (*string, bool)`

GetCustomTitleOk returns a tuple with the CustomTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTitle

`func (o *DaVinciFlowSettingsResponse) SetCustomTitle(v string)`

SetCustomTitle sets CustomTitle field to given value.

### HasCustomTitle

`func (o *DaVinciFlowSettingsResponse) HasCustomTitle() bool`

HasCustomTitle returns a boolean if a field has been set.

### GetDefaultErrorScreenBrandLogo

`func (o *DaVinciFlowSettingsResponse) GetDefaultErrorScreenBrandLogo() bool`

GetDefaultErrorScreenBrandLogo returns the DefaultErrorScreenBrandLogo field if non-nil, zero value otherwise.

### GetDefaultErrorScreenBrandLogoOk

`func (o *DaVinciFlowSettingsResponse) GetDefaultErrorScreenBrandLogoOk() (*bool, bool)`

GetDefaultErrorScreenBrandLogoOk returns a tuple with the DefaultErrorScreenBrandLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultErrorScreenBrandLogo

`func (o *DaVinciFlowSettingsResponse) SetDefaultErrorScreenBrandLogo(v bool)`

SetDefaultErrorScreenBrandLogo sets DefaultErrorScreenBrandLogo field to given value.

### HasDefaultErrorScreenBrandLogo

`func (o *DaVinciFlowSettingsResponse) HasDefaultErrorScreenBrandLogo() bool`

HasDefaultErrorScreenBrandLogo returns a boolean if a field has been set.

### GetFlowHttpTimeoutInSeconds

`func (o *DaVinciFlowSettingsResponse) GetFlowHttpTimeoutInSeconds() int32`

GetFlowHttpTimeoutInSeconds returns the FlowHttpTimeoutInSeconds field if non-nil, zero value otherwise.

### GetFlowHttpTimeoutInSecondsOk

`func (o *DaVinciFlowSettingsResponse) GetFlowHttpTimeoutInSecondsOk() (*int32, bool)`

GetFlowHttpTimeoutInSecondsOk returns a tuple with the FlowHttpTimeoutInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowHttpTimeoutInSeconds

`func (o *DaVinciFlowSettingsResponse) SetFlowHttpTimeoutInSeconds(v int32)`

SetFlowHttpTimeoutInSeconds sets FlowHttpTimeoutInSeconds field to given value.

### HasFlowHttpTimeoutInSeconds

`func (o *DaVinciFlowSettingsResponse) HasFlowHttpTimeoutInSeconds() bool`

HasFlowHttpTimeoutInSeconds returns a boolean if a field has been set.

### GetFlowTimeoutInSeconds

`func (o *DaVinciFlowSettingsResponse) GetFlowTimeoutInSeconds() float32`

GetFlowTimeoutInSeconds returns the FlowTimeoutInSeconds field if non-nil, zero value otherwise.

### GetFlowTimeoutInSecondsOk

`func (o *DaVinciFlowSettingsResponse) GetFlowTimeoutInSecondsOk() (*float32, bool)`

GetFlowTimeoutInSecondsOk returns a tuple with the FlowTimeoutInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowTimeoutInSeconds

`func (o *DaVinciFlowSettingsResponse) SetFlowTimeoutInSeconds(v float32)`

SetFlowTimeoutInSeconds sets FlowTimeoutInSeconds field to given value.

### HasFlowTimeoutInSeconds

`func (o *DaVinciFlowSettingsResponse) HasFlowTimeoutInSeconds() bool`

HasFlowTimeoutInSeconds returns a boolean if a field has been set.

### GetIntermediateLoadingScreenCSS

`func (o *DaVinciFlowSettingsResponse) GetIntermediateLoadingScreenCSS() DaVinciFlowSettingsResponseIntermediateLoadingScreenCSS`

GetIntermediateLoadingScreenCSS returns the IntermediateLoadingScreenCSS field if non-nil, zero value otherwise.

### GetIntermediateLoadingScreenCSSOk

`func (o *DaVinciFlowSettingsResponse) GetIntermediateLoadingScreenCSSOk() (*DaVinciFlowSettingsResponseIntermediateLoadingScreenCSS, bool)`

GetIntermediateLoadingScreenCSSOk returns a tuple with the IntermediateLoadingScreenCSS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediateLoadingScreenCSS

`func (o *DaVinciFlowSettingsResponse) SetIntermediateLoadingScreenCSS(v DaVinciFlowSettingsResponseIntermediateLoadingScreenCSS)`

SetIntermediateLoadingScreenCSS sets IntermediateLoadingScreenCSS field to given value.

### HasIntermediateLoadingScreenCSS

`func (o *DaVinciFlowSettingsResponse) HasIntermediateLoadingScreenCSS() bool`

HasIntermediateLoadingScreenCSS returns a boolean if a field has been set.

### GetIntermediateLoadingScreenHTML

`func (o *DaVinciFlowSettingsResponse) GetIntermediateLoadingScreenHTML() DaVinciFlowSettingsResponseIntermediateLoadingScreenHTML`

GetIntermediateLoadingScreenHTML returns the IntermediateLoadingScreenHTML field if non-nil, zero value otherwise.

### GetIntermediateLoadingScreenHTMLOk

`func (o *DaVinciFlowSettingsResponse) GetIntermediateLoadingScreenHTMLOk() (*DaVinciFlowSettingsResponseIntermediateLoadingScreenHTML, bool)`

GetIntermediateLoadingScreenHTMLOk returns a tuple with the IntermediateLoadingScreenHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediateLoadingScreenHTML

`func (o *DaVinciFlowSettingsResponse) SetIntermediateLoadingScreenHTML(v DaVinciFlowSettingsResponseIntermediateLoadingScreenHTML)`

SetIntermediateLoadingScreenHTML sets IntermediateLoadingScreenHTML field to given value.

### HasIntermediateLoadingScreenHTML

`func (o *DaVinciFlowSettingsResponse) HasIntermediateLoadingScreenHTML() bool`

HasIntermediateLoadingScreenHTML returns a boolean if a field has been set.

### GetJsCustomFlowPlayer

`func (o *DaVinciFlowSettingsResponse) GetJsCustomFlowPlayer() string`

GetJsCustomFlowPlayer returns the JsCustomFlowPlayer field if non-nil, zero value otherwise.

### GetJsCustomFlowPlayerOk

`func (o *DaVinciFlowSettingsResponse) GetJsCustomFlowPlayerOk() (*string, bool)`

GetJsCustomFlowPlayerOk returns a tuple with the JsCustomFlowPlayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsCustomFlowPlayer

`func (o *DaVinciFlowSettingsResponse) SetJsCustomFlowPlayer(v string)`

SetJsCustomFlowPlayer sets JsCustomFlowPlayer field to given value.

### HasJsCustomFlowPlayer

`func (o *DaVinciFlowSettingsResponse) HasJsCustomFlowPlayer() bool`

HasJsCustomFlowPlayer returns a boolean if a field has been set.

### GetJsLinks

`func (o *DaVinciFlowSettingsResponse) GetJsLinks() []DaVinciFlowSettingsResponseJsLink`

GetJsLinks returns the JsLinks field if non-nil, zero value otherwise.

### GetJsLinksOk

`func (o *DaVinciFlowSettingsResponse) GetJsLinksOk() (*[]DaVinciFlowSettingsResponseJsLink, bool)`

GetJsLinksOk returns a tuple with the JsLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsLinks

`func (o *DaVinciFlowSettingsResponse) SetJsLinks(v []DaVinciFlowSettingsResponseJsLink)`

SetJsLinks sets JsLinks field to given value.

### HasJsLinks

`func (o *DaVinciFlowSettingsResponse) HasJsLinks() bool`

HasJsLinks returns a boolean if a field has been set.

### GetLogLevel

`func (o *DaVinciFlowSettingsResponse) GetLogLevel() int32`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *DaVinciFlowSettingsResponse) GetLogLevelOk() (*int32, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *DaVinciFlowSettingsResponse) SetLogLevel(v int32)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *DaVinciFlowSettingsResponse) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetRequireAuthenticationToInitiate

`func (o *DaVinciFlowSettingsResponse) GetRequireAuthenticationToInitiate() bool`

GetRequireAuthenticationToInitiate returns the RequireAuthenticationToInitiate field if non-nil, zero value otherwise.

### GetRequireAuthenticationToInitiateOk

`func (o *DaVinciFlowSettingsResponse) GetRequireAuthenticationToInitiateOk() (*bool, bool)`

GetRequireAuthenticationToInitiateOk returns a tuple with the RequireAuthenticationToInitiate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAuthenticationToInitiate

`func (o *DaVinciFlowSettingsResponse) SetRequireAuthenticationToInitiate(v bool)`

SetRequireAuthenticationToInitiate sets RequireAuthenticationToInitiate field to given value.

### HasRequireAuthenticationToInitiate

`func (o *DaVinciFlowSettingsResponse) HasRequireAuthenticationToInitiate() bool`

HasRequireAuthenticationToInitiate returns a boolean if a field has been set.

### GetScrubSensitiveInfo

`func (o *DaVinciFlowSettingsResponse) GetScrubSensitiveInfo() bool`

GetScrubSensitiveInfo returns the ScrubSensitiveInfo field if non-nil, zero value otherwise.

### GetScrubSensitiveInfoOk

`func (o *DaVinciFlowSettingsResponse) GetScrubSensitiveInfoOk() (*bool, bool)`

GetScrubSensitiveInfoOk returns a tuple with the ScrubSensitiveInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrubSensitiveInfo

`func (o *DaVinciFlowSettingsResponse) SetScrubSensitiveInfo(v bool)`

SetScrubSensitiveInfo sets ScrubSensitiveInfo field to given value.

### HasScrubSensitiveInfo

`func (o *DaVinciFlowSettingsResponse) HasScrubSensitiveInfo() bool`

HasScrubSensitiveInfo returns a boolean if a field has been set.

### GetSensitiveInfoFields

`func (o *DaVinciFlowSettingsResponse) GetSensitiveInfoFields() []string`

GetSensitiveInfoFields returns the SensitiveInfoFields field if non-nil, zero value otherwise.

### GetSensitiveInfoFieldsOk

`func (o *DaVinciFlowSettingsResponse) GetSensitiveInfoFieldsOk() (*[]string, bool)`

GetSensitiveInfoFieldsOk returns a tuple with the SensitiveInfoFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveInfoFields

`func (o *DaVinciFlowSettingsResponse) SetSensitiveInfoFields(v []string)`

SetSensitiveInfoFields sets SensitiveInfoFields field to given value.

### HasSensitiveInfoFields

`func (o *DaVinciFlowSettingsResponse) HasSensitiveInfoFields() bool`

HasSensitiveInfoFields returns a boolean if a field has been set.

### GetUseCSP

`func (o *DaVinciFlowSettingsResponse) GetUseCSP() bool`

GetUseCSP returns the UseCSP field if non-nil, zero value otherwise.

### GetUseCSPOk

`func (o *DaVinciFlowSettingsResponse) GetUseCSPOk() (*bool, bool)`

GetUseCSPOk returns a tuple with the UseCSP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCSP

`func (o *DaVinciFlowSettingsResponse) SetUseCSP(v bool)`

SetUseCSP sets UseCSP field to given value.

### HasUseCSP

`func (o *DaVinciFlowSettingsResponse) HasUseCSP() bool`

HasUseCSP returns a boolean if a field has been set.

### GetUseCustomCSS

`func (o *DaVinciFlowSettingsResponse) GetUseCustomCSS() bool`

GetUseCustomCSS returns the UseCustomCSS field if non-nil, zero value otherwise.

### GetUseCustomCSSOk

`func (o *DaVinciFlowSettingsResponse) GetUseCustomCSSOk() (*bool, bool)`

GetUseCustomCSSOk returns a tuple with the UseCustomCSS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCustomCSS

`func (o *DaVinciFlowSettingsResponse) SetUseCustomCSS(v bool)`

SetUseCustomCSS sets UseCustomCSS field to given value.

### HasUseCustomCSS

`func (o *DaVinciFlowSettingsResponse) HasUseCustomCSS() bool`

HasUseCustomCSS returns a boolean if a field has been set.

### GetUseCustomFlowPlayer

`func (o *DaVinciFlowSettingsResponse) GetUseCustomFlowPlayer() bool`

GetUseCustomFlowPlayer returns the UseCustomFlowPlayer field if non-nil, zero value otherwise.

### GetUseCustomFlowPlayerOk

`func (o *DaVinciFlowSettingsResponse) GetUseCustomFlowPlayerOk() (*bool, bool)`

GetUseCustomFlowPlayerOk returns a tuple with the UseCustomFlowPlayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCustomFlowPlayer

`func (o *DaVinciFlowSettingsResponse) SetUseCustomFlowPlayer(v bool)`

SetUseCustomFlowPlayer sets UseCustomFlowPlayer field to given value.

### HasUseCustomFlowPlayer

`func (o *DaVinciFlowSettingsResponse) HasUseCustomFlowPlayer() bool`

HasUseCustomFlowPlayer returns a boolean if a field has been set.

### GetUseCustomScript

`func (o *DaVinciFlowSettingsResponse) GetUseCustomScript() bool`

GetUseCustomScript returns the UseCustomScript field if non-nil, zero value otherwise.

### GetUseCustomScriptOk

`func (o *DaVinciFlowSettingsResponse) GetUseCustomScriptOk() (*bool, bool)`

GetUseCustomScriptOk returns a tuple with the UseCustomScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCustomScript

`func (o *DaVinciFlowSettingsResponse) SetUseCustomScript(v bool)`

SetUseCustomScript sets UseCustomScript field to given value.

### HasUseCustomScript

`func (o *DaVinciFlowSettingsResponse) HasUseCustomScript() bool`

HasUseCustomScript returns a boolean if a field has been set.

### GetUseIntermediateLoadingScreen

`func (o *DaVinciFlowSettingsResponse) GetUseIntermediateLoadingScreen() bool`

GetUseIntermediateLoadingScreen returns the UseIntermediateLoadingScreen field if non-nil, zero value otherwise.

### GetUseIntermediateLoadingScreenOk

`func (o *DaVinciFlowSettingsResponse) GetUseIntermediateLoadingScreenOk() (*bool, bool)`

GetUseIntermediateLoadingScreenOk returns a tuple with the UseIntermediateLoadingScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIntermediateLoadingScreen

`func (o *DaVinciFlowSettingsResponse) SetUseIntermediateLoadingScreen(v bool)`

SetUseIntermediateLoadingScreen sets UseIntermediateLoadingScreen field to given value.

### HasUseIntermediateLoadingScreen

`func (o *DaVinciFlowSettingsResponse) HasUseIntermediateLoadingScreen() bool`

HasUseIntermediateLoadingScreen returns a boolean if a field has been set.

### GetValidateOnSave

`func (o *DaVinciFlowSettingsResponse) GetValidateOnSave() bool`

GetValidateOnSave returns the ValidateOnSave field if non-nil, zero value otherwise.

### GetValidateOnSaveOk

`func (o *DaVinciFlowSettingsResponse) GetValidateOnSaveOk() (*bool, bool)`

GetValidateOnSaveOk returns a tuple with the ValidateOnSave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateOnSave

`func (o *DaVinciFlowSettingsResponse) SetValidateOnSave(v bool)`

SetValidateOnSave sets ValidateOnSave field to given value.

### HasValidateOnSave

`func (o *DaVinciFlowSettingsResponse) HasValidateOnSave() bool`

HasValidateOnSave returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


