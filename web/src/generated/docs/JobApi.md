# ScaleShift.JobApi

All URIs are relative to *http://localhost:9000/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**deleteJob**](JobApi.md#deleteJob) | **DELETE** /jobs/{id} | 
[**getJobDetail**](JobApi.md#getJobDetail) | **GET** /jobs/{id} | 
[**getJobFiles**](JobApi.md#getJobFiles) | **GET** /jobs/{id}/files | 
[**getJobLogs**](JobApi.md#getJobLogs) | **GET** /jobs/{id}/logs | 
[**getJobs**](JobApi.md#getJobs) | **GET** /jobs | 
[**modifyJob**](JobApi.md#modifyJob) | **PATCH** /jobs/{id} | 
[**postNewJob**](JobApi.md#postNewJob) | **POST** /jobs | 


<a name="deleteJob"></a>
# **deleteJob**
> deleteJob(id)



delete a job 

### Example
```javascript
var ScaleShift = require('scale_shift');
var defaultClient = ScaleShift.ApiClient.instance;

// Configure API key authorization: api-authorizer
var api-authorizer = defaultClient.authentications['api-authorizer'];
api-authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-authorizer.apiKeyPrefix = 'Token';

var apiInstance = new ScaleShift.JobApi();

var id = "id_example"; // String | Job ID


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.deleteJob(id, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| Job ID | 

### Return type

null (empty response body)

### Authorization

[api-authorizer](../README.md#api-authorizer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getJobDetail"></a>
# **getJobDetail**
> JobDetail getJobDetail(id)



returns the details of a job 

### Example
```javascript
var ScaleShift = require('scale_shift');
var defaultClient = ScaleShift.ApiClient.instance;

// Configure API key authorization: api-authorizer
var api-authorizer = defaultClient.authentications['api-authorizer'];
api-authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-authorizer.apiKeyPrefix = 'Token';

var apiInstance = new ScaleShift.JobApi();

var id = "id_example"; // String | Job ID


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getJobDetail(id, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| Job ID | 

### Return type

[**JobDetail**](JobDetail.md)

### Authorization

[api-authorizer](../README.md#api-authorizer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getJobFiles"></a>
# **getJobFiles**
> JobFiles getJobFiles(id)



returns the list of output files 

### Example
```javascript
var ScaleShift = require('scale_shift');
var defaultClient = ScaleShift.ApiClient.instance;

// Configure API key authorization: api-authorizer
var api-authorizer = defaultClient.authentications['api-authorizer'];
api-authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-authorizer.apiKeyPrefix = 'Token';

var apiInstance = new ScaleShift.JobApi();

var id = "id_example"; // String | Job ID


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getJobFiles(id, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| Job ID | 

### Return type

[**JobFiles**](JobFiles.md)

### Authorization

[api-authorizer](../README.md#api-authorizer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getJobLogs"></a>
# **getJobLogs**
> JobLogs getJobLogs(id)



returns the logs of a job 

### Example
```javascript
var ScaleShift = require('scale_shift');
var defaultClient = ScaleShift.ApiClient.instance;

// Configure API key authorization: api-authorizer
var api-authorizer = defaultClient.authentications['api-authorizer'];
api-authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-authorizer.apiKeyPrefix = 'Token';

var apiInstance = new ScaleShift.JobApi();

var id = "id_example"; // String | Job ID


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getJobLogs(id, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| Job ID | 

### Return type

[**JobLogs**](JobLogs.md)

### Authorization

[api-authorizer](../README.md#api-authorizer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getJobs"></a>
# **getJobs**
> [Job] getJobs()



returns training jobs on cloud 

### Example
```javascript
var ScaleShift = require('scale_shift');
var defaultClient = ScaleShift.ApiClient.instance;

// Configure API key authorization: api-authorizer
var api-authorizer = defaultClient.authentications['api-authorizer'];
api-authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-authorizer.apiKeyPrefix = 'Token';

var apiInstance = new ScaleShift.JobApi();

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getJobs(callback);
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**[Job]**](Job.md)

### Authorization

[api-authorizer](../README.md#api-authorizer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="modifyJob"></a>
# **modifyJob**
> modifyJob(id, body)



modify the job status 

### Example
```javascript
var ScaleShift = require('scale_shift');
var defaultClient = ScaleShift.ApiClient.instance;

// Configure API key authorization: api-authorizer
var api-authorizer = defaultClient.authentications['api-authorizer'];
api-authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-authorizer.apiKeyPrefix = 'Token';

var apiInstance = new ScaleShift.JobApi();

var id = "id_example"; // String | Job ID

var body = new ScaleShift.JobAttrs1(); // JobAttrs1 | 


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.modifyJob(id, body, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| Job ID | 
 **body** | [**JobAttrs1**](JobAttrs1.md)|  | 

### Return type

null (empty response body)

### Authorization

[api-authorizer](../README.md#api-authorizer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="postNewJob"></a>
# **postNewJob**
> NewJobID postNewJob(body)



Submit a job with the specified image 

### Example
```javascript
var ScaleShift = require('scale_shift');
var defaultClient = ScaleShift.ApiClient.instance;

// Configure API key authorization: api-authorizer
var api-authorizer = defaultClient.authentications['api-authorizer'];
api-authorizer.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//api-authorizer.apiKeyPrefix = 'Token';

var apiInstance = new ScaleShift.JobApi();

var body = new ScaleShift.JobAttrs(); // JobAttrs | 


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.postNewJob(body, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**JobAttrs**](JobAttrs.md)|  | 

### Return type

[**NewJobID**](NewJobID.md)

### Authorization

[api-authorizer](../README.md#api-authorizer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

