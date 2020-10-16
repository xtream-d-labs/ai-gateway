# AiGateway.RescaleApi

All URIs are relative to *http://localhost:9000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getRescaleApplication**](RescaleApi.md#getRescaleApplication) | **GET** /rescale/applications/{code}/ | 
[**getRescaleApplicationVersion**](RescaleApi.md#getRescaleApplicationVersion) | **GET** /rescale/applications/{code}/{version}/ | 
[**getRescaleCoreTypes**](RescaleApi.md#getRescaleCoreTypes) | **GET** /rescale/coretypes | 


<a name="getRescaleApplication"></a>
# **getRescaleApplication**
> RescaleApplication getRescaleApplication(code)



returns a Rescale application 

### Example
```javascript
var AiGateway = require('ai_gateway');
var defaultClient = AiGateway.ApiClient.instance;

// Configure API key authorization: api-authorizer
var api-authorizer = defaultClient.authentications['api-authorizer'];
api-authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-authorizer.apiKeyPrefix = 'Token';

var apiInstance = new AiGateway.RescaleApi();

var code = "code_example"; // String | application code


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getRescaleApplication(code, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **String**| application code | 

### Return type

[**RescaleApplication**](RescaleApplication.md)

### Authorization

[api-authorizer](../README.md#api-authorizer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getRescaleApplicationVersion"></a>
# **getRescaleApplicationVersion**
> RescaleApplicationVersion getRescaleApplicationVersion(code, version)



returns version information of a specified Rescale application 

### Example
```javascript
var AiGateway = require('ai_gateway');
var defaultClient = AiGateway.ApiClient.instance;

// Configure API key authorization: api-authorizer
var api-authorizer = defaultClient.authentications['api-authorizer'];
api-authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-authorizer.apiKeyPrefix = 'Token';

var apiInstance = new AiGateway.RescaleApi();

var code = "code_example"; // String | Rescale application code

var version = "version_example"; // String | Rescale application version


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getRescaleApplicationVersion(code, version, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **String**| Rescale application code | 
 **version** | **String**| Rescale application version | 

### Return type

[**RescaleApplicationVersion**](RescaleApplicationVersion.md)

### Authorization

[api-authorizer](../README.md#api-authorizer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getRescaleCoreTypes"></a>
# **getRescaleCoreTypes**
> [RescaleCoreType] getRescaleCoreTypes(opts)



returns Rescale CoreTypes 

### Example
```javascript
var AiGateway = require('ai_gateway');
var defaultClient = AiGateway.ApiClient.instance;

// Configure API key authorization: api-authorizer
var api-authorizer = defaultClient.authentications['api-authorizer'];
api-authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-authorizer.apiKeyPrefix = 'Token';

var apiInstance = new AiGateway.RescaleApi();

var opts = { 
  'appVer': "appVer_example", // String | Rescale Application version
  'minGpus': 56 // Number | Required number of GPUs
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getRescaleCoreTypes(opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appVer** | **String**| Rescale Application version | [optional] 
 **minGpus** | **Number**| Required number of GPUs | [optional] 

### Return type

[**[RescaleCoreType]**](RescaleCoreType.md)

### Authorization

[api-authorizer](../README.md#api-authorizer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

