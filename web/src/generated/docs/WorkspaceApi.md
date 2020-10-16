# AiGateway.WorkspaceApi

All URIs are relative to *http://localhost:9000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**deleteWorkspace**](WorkspaceApi.md#deleteWorkspace) | **DELETE** /workspaces | 
[**getWorkspaces**](WorkspaceApi.md#getWorkspaces) | **GET** /workspaces | 


<a name="deleteWorkspace"></a>
# **deleteWorkspace**
> deleteWorkspace(body)



delete user&#39;s workspace 

### Example
```javascript
var AiGateway = require('ai_gateway');

var apiInstance = new AiGateway.WorkspaceApi();

var body = new AiGateway.Workspace(); // Workspace | 


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.deleteWorkspace(body, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Workspace**](Workspace.md)|  | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getWorkspaces"></a>
# **getWorkspaces**
> [Workspace] getWorkspaces()



returns user&#39;s workspaces 

### Example
```javascript
var AiGateway = require('ai_gateway');

var apiInstance = new AiGateway.WorkspaceApi();

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getWorkspaces(callback);
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**[Workspace]**](Workspace.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

