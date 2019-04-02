# ScaleShift.AppApi

All URIs are relative to *http://localhost:9000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getConfigurations**](AppApi.md#getConfigurations) | **GET** /configurations | 
[**getEndpoints**](AppApi.md#getEndpoints) | **GET** /endpoints | 
[**getVersions**](AppApi.md#getVersions) | **GET** /versions | 
[**postConfigurations**](AppApi.md#postConfigurations) | **POST** /configurations | 
[**postNewSession**](AppApi.md#postNewSession) | **POST** /sessions | 


<a name="getConfigurations"></a>
# **getConfigurations**
> Configuration getConfigurations()



returns app configurations 

### Example
```javascript
var ScaleShift = require('scale_shift');

var apiInstance = new ScaleShift.AppApi();

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getConfigurations(callback);
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**Configuration**](Configuration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getEndpoints"></a>
# **getEndpoints**
> Endpoints getEndpoints()



returns third-party endpoints 

### Example
```javascript
var ScaleShift = require('scale_shift');

var apiInstance = new ScaleShift.AppApi();

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getEndpoints(callback);
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**Endpoints**](Endpoints.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getVersions"></a>
# **getVersions**
> Versions getVersions()



returns application versions 

### Example
```javascript
var ScaleShift = require('scale_shift');

var apiInstance = new ScaleShift.AppApi();

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getVersions(callback);
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**Versions**](Versions.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="postConfigurations"></a>
# **postConfigurations**
> Session postConfigurations(body)



set app configurations 

### Example
```javascript
var ScaleShift = require('scale_shift');

var apiInstance = new ScaleShift.AppApi();

var body = new ScaleShift.AccountInfo(); // AccountInfo | 


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.postConfigurations(body, callback);
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

<a name="postNewSession"></a>
# **postNewSession**
> Session postNewSession(body)



login 

### Example
```javascript
var ScaleShift = require('scale_shift');

var apiInstance = new ScaleShift.AppApi();

var body = new ScaleShift.AccountInfo1(); // AccountInfo1 | 


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
 **body** | [**AccountInfo1**](AccountInfo1.md)|  | 

### Return type

[**Session**](Session.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

