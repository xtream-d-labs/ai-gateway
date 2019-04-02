# ScaleShift.SessionApi

All URIs are relative to *http://localhost:9000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**postNewSession**](SessionApi.md#postNewSession) | **POST** /sessions | 


<a name="postNewSession"></a>
# **postNewSession**
> Session postNewSession(body)



login 

### Example
```javascript
var ScaleShift = require('scale_shift');

var apiInstance = new ScaleShift.SessionApi();

var body = new ScaleShift.AccountInfo(); // AccountInfo | 


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.postNewSession(body, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AccountInfo**](AccountInfo.md)|  | 

### Return type

[**Session**](Session.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

