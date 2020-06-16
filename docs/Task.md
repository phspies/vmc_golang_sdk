# Task

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Updated** | [**time.Time**](time.Time.md) |  | [default to null]
**UserId** | **string** | User id that last updated this record | [default to null]
**UpdatedByUserId** | **string** | User id that last updated this record | [default to null]
**Created** | [**time.Time**](time.Time.md) |  | [default to null]
**Version** | **int32** | Version of this entity | [default to null]
**UpdatedByUserName** | **string** | User name that last updated this record | [optional] [default to null]
**UserName** | **string** | User name that last updated this record | [default to null]
**Id** | **string** | Unique ID for this entity | [default to null]
**Status** | **string** |  | [optional] [default to null]
**LocalizedErrorMessage** | **string** |  | [optional] [default to null]
**ResourceId** | **string** | UUID of the resource the task is acting upon | [optional] [default to null]
**ParentTaskId** | **string** | If this task was created by another task - this provides the linkage. Mostly for debugging. | [optional] [default to null]
**TaskVersion** | **string** |  | [optional] [default to null]
**CorrelationId** | **string** | (Optional) Client provided uniqifier to make task creation idempotent. Be aware not all tasks support this. For tasks that do - supplying the same correlation Id, for the same task type, within a predefined window will ensure the operation happens at most once. | [optional] [default to null]
**StartResourceEntityVersion** | **int32** | Entity version of the resource at the start of the task. This is only set for some task types. | [optional] [default to null]
**SubStatus** | **string** |  | [optional] [default to null]
**TaskType** | **string** |  | [optional] [default to null]
**StartTime** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**TaskProgressPhases** | [**[]TaskProgressPhase**](TaskProgressPhase.md) | Task progress phases involved in current task execution | [optional] [default to null]
**ErrorMessage** | **string** |  | [optional] [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**EndResourceEntityVersion** | **int32** | Entity version of the resource at the end of the task. This is only set for some task types. | [optional] [default to null]
**ServiceErrors** | [**[]ServiceError**](ServiceError.md) | Service errors returned from SDDC services. | [optional] [default to null]
**OrgType** | **string** |  | [optional] [default to null]
**EstimatedRemainingMinutes** | **int32** | Estimated remaining time in minute of the task execution, &lt; 0 means no estimation for the task. | [optional] [default to null]
**Params** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**ProgressPercent** | **int32** | Estimated progress percentage the task executed | [optional] [default to null]
**PhaseInProgress** | **string** | The current in progress phase ID in the task execution, if none in progress, empty string returned. | [optional] [default to null]
**ResourceType** | **string** | Type of resource being acted upon | [optional] [default to null]
**EndTime** | [**time.Time**](time.Time.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

