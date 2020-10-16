# AiGateway.NotebookApi

All URIs are relative to *http://localhost:9000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**deleteNotebook**](NotebookApi.md#deleteNotebook) | **DELETE** /notebooks/{id} | 
[**getIPythonNotebooks**](NotebookApi.md#getIPythonNotebooks) | **GET** /notebooks/{id}/ipynbs | 
[**getNotebookDetails**](NotebookApi.md#getNotebookDetails) | **GET** /notebooks/{id} | 
[**getNotebooks**](NotebookApi.md#getNotebooks) | **GET** /notebooks | 
[**modifyNotebook**](NotebookApi.md#modifyNotebook) | **PATCH** /notebooks/{id} | 
[**postNewNotebook**](NotebookApi.md#postNewNotebook) | **POST** /notebooks | 


<a name="deleteNotebook"></a>
# **deleteNotebook**
> deleteNotebook(id)



delete a specified notebook 

### Example
```javascript
var AiGateway = require('ai_gateway');

var apiInstance = new AiGateway.NotebookApi();

var id = "id_example"; // String | Notebook container ID


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.deleteNotebook(id, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| Notebook container ID | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getIPythonNotebooks"></a>
# **getIPythonNotebooks**
> [IPythonNotebook] getIPythonNotebooks(id)



returns ipynb files on the specified notebook 

### Example
```javascript
var AiGateway = require('ai_gateway');

var apiInstance = new AiGateway.NotebookApi();

var id = "id_example"; // String | Notebook container ID


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getIPythonNotebooks(id, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| Notebook container ID | 

### Return type

[**[IPythonNotebook]**](IPythonNotebook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getNotebookDetails"></a>
# **getNotebookDetails**
> NotebookDetail getNotebookDetails(id)



returns Jupyter notebook detail information 

### Example
```javascript
var AiGateway = require('ai_gateway');

var apiInstance = new AiGateway.NotebookApi();

var id = "id_example"; // String | Notebook container ID


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getNotebookDetails(id, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| Notebook container ID | 

### Return type

[**NotebookDetail**](NotebookDetail.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getNotebooks"></a>
# **getNotebooks**
> [Notebook] getNotebooks()



returns Jupyter notebook information 

### Example
```javascript
var AiGateway = require('ai_gateway');

var apiInstance = new AiGateway.NotebookApi();

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getNotebooks(callback);
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**[Notebook]**](Notebook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="modifyNotebook"></a>
# **modifyNotebook**
> modifyNotebook(id, body)



modify the notebook status 

### Example
```javascript
var AiGateway = require('ai_gateway');

var apiInstance = new AiGateway.NotebookApi();

var id = "id_example"; // String | Notebook container ID

var body = new AiGateway.NotebookAttrs(); // NotebookAttrs | 


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.modifyNotebook(id, body, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| Notebook container ID | 
 **body** | [**NotebookAttrs**](NotebookAttrs.md)|  | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="postNewNotebook"></a>
# **postNewNotebook**
> postNewNotebook(body)



creates Jupyter notebook container 

### Example
```javascript
var AiGateway = require('ai_gateway');

var apiInstance = new AiGateway.NotebookApi();

var body = new AiGateway.ImageName2(); // ImageName2 | 


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.postNewNotebook(body, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ImageName2**](ImageName2.md)|  | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

