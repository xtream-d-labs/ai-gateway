# ScaleShift.JobAttrs

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**platformId** | **String** | Platform ID | [optional] 
**notebookId** | **String** | Notebook container ID | [optional] 
**entrypointFile** | **String** | The entrypoint file of the job | [optional] 
**commands** | **[String]** | commands to be excuted after the entrypoint | [optional] 
**cpu** | **Number** | Requesting millicores of CPU | [optional] 
**mem** | **Number** | Requesting bytes of memory | [optional] 
**gpu** | **Number** | Requesting number of GPU | [optional] 
**coretype** | **String** | Rescale CoreType as its infrastructure | [optional] 
**cores** | **Number** | The number of CPU cores | [optional] 


<a name="PlatformIdEnum"></a>
## Enum: PlatformIdEnum


* `kubernetes` (value: `"kubernetes"`)

* `rescale` (value: `"rescale"`)




