# ScaleShift.RepositoryApi

All URIs are relative to *http://localhost:9000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getNgcImages**](RepositoryApi.md#getNgcImages) | **GET** /nvidia/repositories/{namespace}/images/{id} | 
[**getNgcRepositories**](RepositoryApi.md#getNgcRepositories) | **GET** /nvidia/repositories | 


<a name="getNgcImages"></a>
# **getNgcImages**
> [NgcImage] getNgcImages(namespace, id)



returns NGC images 

### Example
```javascript
var ScaleShift = require('scale_shift');
var defaultClient = ScaleShift.ApiClient.instance;

// Configure API key authorization: api-authorizer
var api-authorizer = defaultClient.authentications['api-authorizer'];
api-authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-authorizer.apiKeyPrefix = 'Token';

var apiInstance = new ScaleShift.RepositoryApi();

var namespace = "namespace_example"; // String | Docker repositry namespace

var id = "id_example"; // String | Docker image name


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getNgcImages(namespace, id, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **String**| Docker repositry namespace | 
 **id** | **String**| Docker image name | 

### Return type

[**[NgcImage]**](NgcImage.md)

### Authorization

[api-authorizer](../README.md#api-authorizer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getNgcRepositories"></a>
# **getNgcRepositories**
> [Repository] getNgcRepositories()



returns NGC repositories 

### Example
```javascript
var ScaleShift = require('scale_shift');
var defaultClient = ScaleShift.ApiClient.instance;

// Configure API key authorization: api-authorizer
var api-authorizer = defaultClient.authentications['api-authorizer'];
api-authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-authorizer.apiKeyPrefix = 'Token';

var apiInstance = new ScaleShift.RepositoryApi();

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getNgcRepositories(callback);
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**[Repository]**](Repository.md)

### Authorization

[api-authorizer](../README.md#api-authorizer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

