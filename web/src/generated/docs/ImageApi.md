# AiGateway.ImageApi

All URIs are relative to *http://localhost:9000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**deleteImage**](ImageApi.md#deleteImage) | **DELETE** /images | 
[**getImages**](ImageApi.md#getImages) | **GET** /images | 
[**postNewImage**](ImageApi.md#postNewImage) | **POST** /images | 


<a name="deleteImage"></a>
# **deleteImage**
> deleteImage(body)



delete a specified local image 

### Example
```javascript
var AiGateway = require('ai_gateway');

var apiInstance = new AiGateway.ImageApi();

var body = new AiGateway.ImageName1(); // ImageName1 | 


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.deleteImage(body, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ImageName1**](ImageName1.md)|  | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getImages"></a>
# **getImages**
> [Image] getImages()



returns local images 

### Example
```javascript
var AiGateway = require('ai_gateway');

var apiInstance = new AiGateway.ImageApi();

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getImages(callback);
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**[Image]**](Image.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="postNewImage"></a>
# **postNewImage**
> postNewImage(body)



pull a specified image from Docker registry 

### Example
```javascript
var AiGateway = require('ai_gateway');

var apiInstance = new AiGateway.ImageApi();

var body = new AiGateway.ImageName(); // ImageName | 


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.postNewImage(body, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ImageName**](ImageName.md)|  | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

