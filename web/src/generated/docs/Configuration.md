# AiGateway.Configuration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**mustSignedIn** | **String** | Users should be signed in | 
**usePrivateRegistry** | **String** | Private registry will be used or not | [optional] 
**useNgc** | **String** | NGC will be used or not | [optional] 
**useK8s** | **String** | Kubernetes will be used or not | [optional] 
**useRescale** | **String** | Rescale will be used or not | [optional] 
**dockerRegistry** | **String** | Docker Registry endpoint | [optional] 
**dockerHostname** | **String** | Hostname for the private Docker registry | [optional] 
**dockerUsername** | **String** | Username for the private Docker registry | [optional] 
**dockerPassword** | **String** | Fist 3 chars of the password of the private Docker registry | [optional] 
**ngcEmail** | **String** | E-mail address for NGC console | [optional] 
**ngcPassword** | **String** | Fist 3 chars of the password for NGC console | [optional] 
**ngcApikey** | **String** | Fist 5 chars of NGC API Key | [optional] 
**k8sConfig** | **String** | kubecfg | [optional] 
**rescalePlatform** | **String** |  | [optional] 
**rescaleKey** | **String** | Fist 5 chars of Rescal API Key | [optional] 
**localGpus** | **Number** | Number of the host GPUs | 
**localGpusPerContainer** | **Number** | Number of GPUs per container | 


<a name="MustSignedInEnum"></a>
## Enum: MustSignedInEnum


* `yes` (value: `"yes"`)

* `no` (value: `"no"`)




<a name="UsePrivateRegistryEnum"></a>
## Enum: UsePrivateRegistryEnum


* `yes` (value: `"yes"`)

* `no` (value: `"no"`)




<a name="UseNgcEnum"></a>
## Enum: UseNgcEnum


* `yes` (value: `"yes"`)

* `no` (value: `"no"`)




<a name="UseK8sEnum"></a>
## Enum: UseK8sEnum


* `yes` (value: `"yes"`)

* `no` (value: `"no"`)




<a name="UseRescaleEnum"></a>
## Enum: UseRescaleEnum


* `yes` (value: `"yes"`)

* `no` (value: `"no"`)




<a name="RescalePlatformEnum"></a>
## Enum: RescalePlatformEnum


* `platform.rescale.com` (value: `"https://platform.rescale.com"`)

* `platform.rescale.jp` (value: `"https://platform.rescale.jp"`)

* `kr.rescale.com` (value: `"https://kr.rescale.com"`)

* `eu.rescale.com` (value: `"https://eu.rescale.com"`)




