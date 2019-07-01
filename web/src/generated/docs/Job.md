# ScaleShift.Job

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **String** | Job ID | 
**platform** | **String** | platform | [optional] 
**status** | **String** | the status of the job | [optional] 
**image** | **String** | the image ID | [optional] 
**mounts** | **[String]** | the container labels | [optional] 
**commands** | **[String]** | the container labels | [optional] 
**started** | **Date** | started unix timestamp | [optional] 
**ended** | **Date** | ended unix timestamp | [optional] 


<a name="PlatformEnum"></a>
## Enum: PlatformEnum


* `kubernetes` (value: `"kubernetes"`)

* `rescale` (value: `"rescale"`)




