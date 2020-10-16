# AiGateway.AppErrorsApi

All URIs are relative to *http://localhost:9000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getAppErrors**](AppErrorsApi.md#getAppErrors) | **GET** /errors | 


<a name="getAppErrors"></a>
# **getAppErrors**
> [AppError] getAppErrors()



returns the list of application errors 

### Example
```javascript
var AiGateway = require('ai_gateway');

var apiInstance = new AiGateway.AppErrorsApi();

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getAppErrors(callback);
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**[AppError]**](AppError.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

