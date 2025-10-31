# DaVinciFlowSettingsRequest

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
**FlowTimeoutInSeconds** | Pointer to **int32** |  | [optional] 
**IntermediateLoadingScreenCSS** | Pointer to [**DaVinciFlowSettingsRequestIntermediateLoadingScreenCSS**](DaVinciFlowSettingsRequestIntermediateLoadingScreenCSS.md) |  | [optional] 
**IntermediateLoadingScreenHTML** | Pointer to [**DaVinciFlowSettingsRequestIntermediateLoadingScreenHTML**](DaVinciFlowSettingsRequestIntermediateLoadingScreenHTML.md) |  | [optional] 
**JsCustomFlowPlayer** | Pointer to **string** |  | [optional] 
**JsLinks** | Pointer to [**[]DaVinciFlowSettingsRequestJsLink**](DaVinciFlowSettingsRequestJsLink.md) |  | [optional] 
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

### NewDaVinciFlowSettingsRequest

`func NewDaVinciFlowSettingsRequest() *DaVinciFlowSettingsRequest`

NewDaVinciFlowSettingsRequest instantiates a new DaVinciFlowSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciFlowSettingsRequestWithDefaults

`func NewDaVinciFlowSettingsRequestWithDefaults() *DaVinciFlowSettingsRequest`

NewDaVinciFlowSettingsRequestWithDefaults instantiates a new DaVinciFlowSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsp

`func (o *DaVinciFlowSettingsRequest) GetCsp() string`

GetCsp returns the Csp field if non-nil, zero value otherwise.

### GetCspOk

`func (o *DaVinciFlowSettingsRequest) GetCspOk() (*string, bool)`

GetCspOk returns a tuple with the Csp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsp

`func (o *DaVinciFlowSettingsRequest) SetCsp(v string)`

SetCsp sets Csp field to given value.

### HasCsp

`func (o *DaVinciFlowSettingsRequest) HasCsp() bool`

HasCsp returns a boolean if a field has been set.

### GetCss

`func (o *DaVinciFlowSettingsRequest) GetCss() string`

GetCss returns the Css field if non-nil, zero value otherwise.

### GetCssOk

`func (o *DaVinciFlowSettingsRequest) GetCssOk() (*string, bool)`

GetCssOk returns a tuple with the Css field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCss

`func (o *DaVinciFlowSettingsRequest) SetCss(v string)`

SetCss sets Css field to given value.

### HasCss

`func (o *DaVinciFlowSettingsRequest) HasCss() bool`

HasCss returns a boolean if a field has been set.

### GetCssLinks

`func (o *DaVinciFlowSettingsRequest) GetCssLinks() []string`

GetCssLinks returns the CssLinks field if non-nil, zero value otherwise.

### GetCssLinksOk

`func (o *DaVinciFlowSettingsRequest) GetCssLinksOk() (*[]string, bool)`

GetCssLinksOk returns a tuple with the CssLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssLinks

`func (o *DaVinciFlowSettingsRequest) SetCssLinks(v []string)`

SetCssLinks sets CssLinks field to given value.

### HasCssLinks

`func (o *DaVinciFlowSettingsRequest) HasCssLinks() bool`

HasCssLinks returns a boolean if a field has been set.

### GetCustomErrorScreenBrandLogoUrl

`func (o *DaVinciFlowSettingsRequest) GetCustomErrorScreenBrandLogoUrl() string`

GetCustomErrorScreenBrandLogoUrl returns the CustomErrorScreenBrandLogoUrl field if non-nil, zero value otherwise.

### GetCustomErrorScreenBrandLogoUrlOk

`func (o *DaVinciFlowSettingsRequest) GetCustomErrorScreenBrandLogoUrlOk() (*string, bool)`

GetCustomErrorScreenBrandLogoUrlOk returns a tuple with the CustomErrorScreenBrandLogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomErrorScreenBrandLogoUrl

`func (o *DaVinciFlowSettingsRequest) SetCustomErrorScreenBrandLogoUrl(v string)`

SetCustomErrorScreenBrandLogoUrl sets CustomErrorScreenBrandLogoUrl field to given value.

### HasCustomErrorScreenBrandLogoUrl

`func (o *DaVinciFlowSettingsRequest) HasCustomErrorScreenBrandLogoUrl() bool`

HasCustomErrorScreenBrandLogoUrl returns a boolean if a field has been set.

### GetCustomErrorShowFooter

`func (o *DaVinciFlowSettingsRequest) GetCustomErrorShowFooter() bool`

GetCustomErrorShowFooter returns the CustomErrorShowFooter field if non-nil, zero value otherwise.

### GetCustomErrorShowFooterOk

`func (o *DaVinciFlowSettingsRequest) GetCustomErrorShowFooterOk() (*bool, bool)`

GetCustomErrorShowFooterOk returns a tuple with the CustomErrorShowFooter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomErrorShowFooter

`func (o *DaVinciFlowSettingsRequest) SetCustomErrorShowFooter(v bool)`

SetCustomErrorShowFooter sets CustomErrorShowFooter field to given value.

### HasCustomErrorShowFooter

`func (o *DaVinciFlowSettingsRequest) HasCustomErrorShowFooter() bool`

HasCustomErrorShowFooter returns a boolean if a field has been set.

### GetCustomFaviconLink

`func (o *DaVinciFlowSettingsRequest) GetCustomFaviconLink() string`

GetCustomFaviconLink returns the CustomFaviconLink field if non-nil, zero value otherwise.

### GetCustomFaviconLinkOk

`func (o *DaVinciFlowSettingsRequest) GetCustomFaviconLinkOk() (*string, bool)`

GetCustomFaviconLinkOk returns a tuple with the CustomFaviconLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFaviconLink

`func (o *DaVinciFlowSettingsRequest) SetCustomFaviconLink(v string)`

SetCustomFaviconLink sets CustomFaviconLink field to given value.

### HasCustomFaviconLink

`func (o *DaVinciFlowSettingsRequest) HasCustomFaviconLink() bool`

HasCustomFaviconLink returns a boolean if a field has been set.

### GetCustomLogoURLSelection

`func (o *DaVinciFlowSettingsRequest) GetCustomLogoURLSelection() int32`

GetCustomLogoURLSelection returns the CustomLogoURLSelection field if non-nil, zero value otherwise.

### GetCustomLogoURLSelectionOk

`func (o *DaVinciFlowSettingsRequest) GetCustomLogoURLSelectionOk() (*int32, bool)`

GetCustomLogoURLSelectionOk returns a tuple with the CustomLogoURLSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLogoURLSelection

`func (o *DaVinciFlowSettingsRequest) SetCustomLogoURLSelection(v int32)`

SetCustomLogoURLSelection sets CustomLogoURLSelection field to given value.

### HasCustomLogoURLSelection

`func (o *DaVinciFlowSettingsRequest) HasCustomLogoURLSelection() bool`

HasCustomLogoURLSelection returns a boolean if a field has been set.

### GetCustomTitle

`func (o *DaVinciFlowSettingsRequest) GetCustomTitle() string`

GetCustomTitle returns the CustomTitle field if non-nil, zero value otherwise.

### GetCustomTitleOk

`func (o *DaVinciFlowSettingsRequest) GetCustomTitleOk() (*string, bool)`

GetCustomTitleOk returns a tuple with the CustomTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTitle

`func (o *DaVinciFlowSettingsRequest) SetCustomTitle(v string)`

SetCustomTitle sets CustomTitle field to given value.

### HasCustomTitle

`func (o *DaVinciFlowSettingsRequest) HasCustomTitle() bool`

HasCustomTitle returns a boolean if a field has been set.

### GetDefaultErrorScreenBrandLogo

`func (o *DaVinciFlowSettingsRequest) GetDefaultErrorScreenBrandLogo() bool`

GetDefaultErrorScreenBrandLogo returns the DefaultErrorScreenBrandLogo field if non-nil, zero value otherwise.

### GetDefaultErrorScreenBrandLogoOk

`func (o *DaVinciFlowSettingsRequest) GetDefaultErrorScreenBrandLogoOk() (*bool, bool)`

GetDefaultErrorScreenBrandLogoOk returns a tuple with the DefaultErrorScreenBrandLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultErrorScreenBrandLogo

`func (o *DaVinciFlowSettingsRequest) SetDefaultErrorScreenBrandLogo(v bool)`

SetDefaultErrorScreenBrandLogo sets DefaultErrorScreenBrandLogo field to given value.

### HasDefaultErrorScreenBrandLogo

`func (o *DaVinciFlowSettingsRequest) HasDefaultErrorScreenBrandLogo() bool`

HasDefaultErrorScreenBrandLogo returns a boolean if a field has been set.

### GetFlowHttpTimeoutInSeconds

`func (o *DaVinciFlowSettingsRequest) GetFlowHttpTimeoutInSeconds() int32`

GetFlowHttpTimeoutInSeconds returns the FlowHttpTimeoutInSeconds field if non-nil, zero value otherwise.

### GetFlowHttpTimeoutInSecondsOk

`func (o *DaVinciFlowSettingsRequest) GetFlowHttpTimeoutInSecondsOk() (*int32, bool)`

GetFlowHttpTimeoutInSecondsOk returns a tuple with the FlowHttpTimeoutInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowHttpTimeoutInSeconds

`func (o *DaVinciFlowSettingsRequest) SetFlowHttpTimeoutInSeconds(v int32)`

SetFlowHttpTimeoutInSeconds sets FlowHttpTimeoutInSeconds field to given value.

### HasFlowHttpTimeoutInSeconds

`func (o *DaVinciFlowSettingsRequest) HasFlowHttpTimeoutInSeconds() bool`

HasFlowHttpTimeoutInSeconds returns a boolean if a field has been set.

### GetFlowTimeoutInSeconds

`func (o *DaVinciFlowSettingsRequest) GetFlowTimeoutInSeconds() int32`

GetFlowTimeoutInSeconds returns the FlowTimeoutInSeconds field if non-nil, zero value otherwise.

### GetFlowTimeoutInSecondsOk

`func (o *DaVinciFlowSettingsRequest) GetFlowTimeoutInSecondsOk() (*int32, bool)`

GetFlowTimeoutInSecondsOk returns a tuple with the FlowTimeoutInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowTimeoutInSeconds

`func (o *DaVinciFlowSettingsRequest) SetFlowTimeoutInSeconds(v int32)`

SetFlowTimeoutInSeconds sets FlowTimeoutInSeconds field to given value.

### HasFlowTimeoutInSeconds

`func (o *DaVinciFlowSettingsRequest) HasFlowTimeoutInSeconds() bool`

HasFlowTimeoutInSeconds returns a boolean if a field has been set.

### GetIntermediateLoadingScreenCSS

`func (o *DaVinciFlowSettingsRequest) GetIntermediateLoadingScreenCSS() DaVinciFlowSettingsRequestIntermediateLoadingScreenCSS`

GetIntermediateLoadingScreenCSS returns the IntermediateLoadingScreenCSS field if non-nil, zero value otherwise.

### GetIntermediateLoadingScreenCSSOk

`func (o *DaVinciFlowSettingsRequest) GetIntermediateLoadingScreenCSSOk() (*DaVinciFlowSettingsRequestIntermediateLoadingScreenCSS, bool)`

GetIntermediateLoadingScreenCSSOk returns a tuple with the IntermediateLoadingScreenCSS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediateLoadingScreenCSS

`func (o *DaVinciFlowSettingsRequest) SetIntermediateLoadingScreenCSS(v DaVinciFlowSettingsRequestIntermediateLoadingScreenCSS)`

SetIntermediateLoadingScreenCSS sets IntermediateLoadingScreenCSS field to given value.

### HasIntermediateLoadingScreenCSS

`func (o *DaVinciFlowSettingsRequest) HasIntermediateLoadingScreenCSS() bool`

HasIntermediateLoadingScreenCSS returns a boolean if a field has been set.

### GetIntermediateLoadingScreenHTML

`func (o *DaVinciFlowSettingsRequest) GetIntermediateLoadingScreenHTML() DaVinciFlowSettingsRequestIntermediateLoadingScreenHTML`

GetIntermediateLoadingScreenHTML returns the IntermediateLoadingScreenHTML field if non-nil, zero value otherwise.

### GetIntermediateLoadingScreenHTMLOk

`func (o *DaVinciFlowSettingsRequest) GetIntermediateLoadingScreenHTMLOk() (*DaVinciFlowSettingsRequestIntermediateLoadingScreenHTML, bool)`

GetIntermediateLoadingScreenHTMLOk returns a tuple with the IntermediateLoadingScreenHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediateLoadingScreenHTML

`func (o *DaVinciFlowSettingsRequest) SetIntermediateLoadingScreenHTML(v DaVinciFlowSettingsRequestIntermediateLoadingScreenHTML)`

SetIntermediateLoadingScreenHTML sets IntermediateLoadingScreenHTML field to given value.

### HasIntermediateLoadingScreenHTML

`func (o *DaVinciFlowSettingsRequest) HasIntermediateLoadingScreenHTML() bool`

HasIntermediateLoadingScreenHTML returns a boolean if a field has been set.

### GetJsCustomFlowPlayer

`func (o *DaVinciFlowSettingsRequest) GetJsCustomFlowPlayer() string`

GetJsCustomFlowPlayer returns the JsCustomFlowPlayer field if non-nil, zero value otherwise.

### GetJsCustomFlowPlayerOk

`func (o *DaVinciFlowSettingsRequest) GetJsCustomFlowPlayerOk() (*string, bool)`

GetJsCustomFlowPlayerOk returns a tuple with the JsCustomFlowPlayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsCustomFlowPlayer

`func (o *DaVinciFlowSettingsRequest) SetJsCustomFlowPlayer(v string)`

SetJsCustomFlowPlayer sets JsCustomFlowPlayer field to given value.

### HasJsCustomFlowPlayer

`func (o *DaVinciFlowSettingsRequest) HasJsCustomFlowPlayer() bool`

HasJsCustomFlowPlayer returns a boolean if a field has been set.

### GetJsLinks

`func (o *DaVinciFlowSettingsRequest) GetJsLinks() []DaVinciFlowSettingsRequestJsLink`

GetJsLinks returns the JsLinks field if non-nil, zero value otherwise.

### GetJsLinksOk

`func (o *DaVinciFlowSettingsRequest) GetJsLinksOk() (*[]DaVinciFlowSettingsRequestJsLink, bool)`

GetJsLinksOk returns a tuple with the JsLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsLinks

`func (o *DaVinciFlowSettingsRequest) SetJsLinks(v []DaVinciFlowSettingsRequestJsLink)`

SetJsLinks sets JsLinks field to given value.

### HasJsLinks

`func (o *DaVinciFlowSettingsRequest) HasJsLinks() bool`

HasJsLinks returns a boolean if a field has been set.

### GetLogLevel

`func (o *DaVinciFlowSettingsRequest) GetLogLevel() int32`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *DaVinciFlowSettingsRequest) GetLogLevelOk() (*int32, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *DaVinciFlowSettingsRequest) SetLogLevel(v int32)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *DaVinciFlowSettingsRequest) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetRequireAuthenticationToInitiate

`func (o *DaVinciFlowSettingsRequest) GetRequireAuthenticationToInitiate() bool`

GetRequireAuthenticationToInitiate returns the RequireAuthenticationToInitiate field if non-nil, zero value otherwise.

### GetRequireAuthenticationToInitiateOk

`func (o *DaVinciFlowSettingsRequest) GetRequireAuthenticationToInitiateOk() (*bool, bool)`

GetRequireAuthenticationToInitiateOk returns a tuple with the RequireAuthenticationToInitiate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAuthenticationToInitiate

`func (o *DaVinciFlowSettingsRequest) SetRequireAuthenticationToInitiate(v bool)`

SetRequireAuthenticationToInitiate sets RequireAuthenticationToInitiate field to given value.

### HasRequireAuthenticationToInitiate

`func (o *DaVinciFlowSettingsRequest) HasRequireAuthenticationToInitiate() bool`

HasRequireAuthenticationToInitiate returns a boolean if a field has been set.

### GetScrubSensitiveInfo

`func (o *DaVinciFlowSettingsRequest) GetScrubSensitiveInfo() bool`

GetScrubSensitiveInfo returns the ScrubSensitiveInfo field if non-nil, zero value otherwise.

### GetScrubSensitiveInfoOk

`func (o *DaVinciFlowSettingsRequest) GetScrubSensitiveInfoOk() (*bool, bool)`

GetScrubSensitiveInfoOk returns a tuple with the ScrubSensitiveInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrubSensitiveInfo

`func (o *DaVinciFlowSettingsRequest) SetScrubSensitiveInfo(v bool)`

SetScrubSensitiveInfo sets ScrubSensitiveInfo field to given value.

### HasScrubSensitiveInfo

`func (o *DaVinciFlowSettingsRequest) HasScrubSensitiveInfo() bool`

HasScrubSensitiveInfo returns a boolean if a field has been set.

### GetSensitiveInfoFields

`func (o *DaVinciFlowSettingsRequest) GetSensitiveInfoFields() []string`

GetSensitiveInfoFields returns the SensitiveInfoFields field if non-nil, zero value otherwise.

### GetSensitiveInfoFieldsOk

`func (o *DaVinciFlowSettingsRequest) GetSensitiveInfoFieldsOk() (*[]string, bool)`

GetSensitiveInfoFieldsOk returns a tuple with the SensitiveInfoFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveInfoFields

`func (o *DaVinciFlowSettingsRequest) SetSensitiveInfoFields(v []string)`

SetSensitiveInfoFields sets SensitiveInfoFields field to given value.

### HasSensitiveInfoFields

`func (o *DaVinciFlowSettingsRequest) HasSensitiveInfoFields() bool`

HasSensitiveInfoFields returns a boolean if a field has been set.

### GetUseCSP

`func (o *DaVinciFlowSettingsRequest) GetUseCSP() bool`

GetUseCSP returns the UseCSP field if non-nil, zero value otherwise.

### GetUseCSPOk

`func (o *DaVinciFlowSettingsRequest) GetUseCSPOk() (*bool, bool)`

GetUseCSPOk returns a tuple with the UseCSP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCSP

`func (o *DaVinciFlowSettingsRequest) SetUseCSP(v bool)`

SetUseCSP sets UseCSP field to given value.

### HasUseCSP

`func (o *DaVinciFlowSettingsRequest) HasUseCSP() bool`

HasUseCSP returns a boolean if a field has been set.

### GetUseCustomCSS

`func (o *DaVinciFlowSettingsRequest) GetUseCustomCSS() bool`

GetUseCustomCSS returns the UseCustomCSS field if non-nil, zero value otherwise.

### GetUseCustomCSSOk

`func (o *DaVinciFlowSettingsRequest) GetUseCustomCSSOk() (*bool, bool)`

GetUseCustomCSSOk returns a tuple with the UseCustomCSS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCustomCSS

`func (o *DaVinciFlowSettingsRequest) SetUseCustomCSS(v bool)`

SetUseCustomCSS sets UseCustomCSS field to given value.

### HasUseCustomCSS

`func (o *DaVinciFlowSettingsRequest) HasUseCustomCSS() bool`

HasUseCustomCSS returns a boolean if a field has been set.

### GetUseCustomFlowPlayer

`func (o *DaVinciFlowSettingsRequest) GetUseCustomFlowPlayer() bool`

GetUseCustomFlowPlayer returns the UseCustomFlowPlayer field if non-nil, zero value otherwise.

### GetUseCustomFlowPlayerOk

`func (o *DaVinciFlowSettingsRequest) GetUseCustomFlowPlayerOk() (*bool, bool)`

GetUseCustomFlowPlayerOk returns a tuple with the UseCustomFlowPlayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCustomFlowPlayer

`func (o *DaVinciFlowSettingsRequest) SetUseCustomFlowPlayer(v bool)`

SetUseCustomFlowPlayer sets UseCustomFlowPlayer field to given value.

### HasUseCustomFlowPlayer

`func (o *DaVinciFlowSettingsRequest) HasUseCustomFlowPlayer() bool`

HasUseCustomFlowPlayer returns a boolean if a field has been set.

### GetUseCustomScript

`func (o *DaVinciFlowSettingsRequest) GetUseCustomScript() bool`

GetUseCustomScript returns the UseCustomScript field if non-nil, zero value otherwise.

### GetUseCustomScriptOk

`func (o *DaVinciFlowSettingsRequest) GetUseCustomScriptOk() (*bool, bool)`

GetUseCustomScriptOk returns a tuple with the UseCustomScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCustomScript

`func (o *DaVinciFlowSettingsRequest) SetUseCustomScript(v bool)`

SetUseCustomScript sets UseCustomScript field to given value.

### HasUseCustomScript

`func (o *DaVinciFlowSettingsRequest) HasUseCustomScript() bool`

HasUseCustomScript returns a boolean if a field has been set.

### GetUseIntermediateLoadingScreen

`func (o *DaVinciFlowSettingsRequest) GetUseIntermediateLoadingScreen() bool`

GetUseIntermediateLoadingScreen returns the UseIntermediateLoadingScreen field if non-nil, zero value otherwise.

### GetUseIntermediateLoadingScreenOk

`func (o *DaVinciFlowSettingsRequest) GetUseIntermediateLoadingScreenOk() (*bool, bool)`

GetUseIntermediateLoadingScreenOk returns a tuple with the UseIntermediateLoadingScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIntermediateLoadingScreen

`func (o *DaVinciFlowSettingsRequest) SetUseIntermediateLoadingScreen(v bool)`

SetUseIntermediateLoadingScreen sets UseIntermediateLoadingScreen field to given value.

### HasUseIntermediateLoadingScreen

`func (o *DaVinciFlowSettingsRequest) HasUseIntermediateLoadingScreen() bool`

HasUseIntermediateLoadingScreen returns a boolean if a field has been set.

### GetValidateOnSave

`func (o *DaVinciFlowSettingsRequest) GetValidateOnSave() bool`

GetValidateOnSave returns the ValidateOnSave field if non-nil, zero value otherwise.

### GetValidateOnSaveOk

`func (o *DaVinciFlowSettingsRequest) GetValidateOnSaveOk() (*bool, bool)`

GetValidateOnSaveOk returns a tuple with the ValidateOnSave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateOnSave

`func (o *DaVinciFlowSettingsRequest) SetValidateOnSave(v bool)`

SetValidateOnSave sets ValidateOnSave field to given value.

### HasValidateOnSave

`func (o *DaVinciFlowSettingsRequest) HasValidateOnSave() bool`

HasValidateOnSave returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


