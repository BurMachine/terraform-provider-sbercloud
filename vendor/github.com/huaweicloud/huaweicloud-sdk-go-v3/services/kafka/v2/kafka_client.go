package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kafka/v2/model"
)

type KafkaClient struct {
	HcClient *http_client.HcHttpClient
}

func NewKafkaClient(hcClient *http_client.HcHttpClient) *KafkaClient {
	return &KafkaClient{HcClient: hcClient}
}

func KafkaClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// BatchCreateOrDeleteKafkaTag 批量添加或删除实例标签
//
// 批量添加或删除实例标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) BatchCreateOrDeleteKafkaTag(request *model.BatchCreateOrDeleteKafkaTagRequest) (*model.BatchCreateOrDeleteKafkaTagResponse, error) {
	requestDef := GenReqDefForBatchCreateOrDeleteKafkaTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateOrDeleteKafkaTagResponse), nil
	}
}

// BatchCreateOrDeleteKafkaTagInvoker 批量添加或删除实例标签
func (c *KafkaClient) BatchCreateOrDeleteKafkaTagInvoker(request *model.BatchCreateOrDeleteKafkaTagRequest) *BatchCreateOrDeleteKafkaTagInvoker {
	requestDef := GenReqDefForBatchCreateOrDeleteKafkaTag()
	return &BatchCreateOrDeleteKafkaTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteInstanceTopic Kafka实例批量删除Topic
//
// 该接口用于向Kafka实例批量删除Topic。批量删除多个消费组时，部分删除成功，部分失败，此时接口返回删除成功，并在返回中显示删除失败的消费组信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) BatchDeleteInstanceTopic(request *model.BatchDeleteInstanceTopicRequest) (*model.BatchDeleteInstanceTopicResponse, error) {
	requestDef := GenReqDefForBatchDeleteInstanceTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteInstanceTopicResponse), nil
	}
}

// BatchDeleteInstanceTopicInvoker Kafka实例批量删除Topic
func (c *KafkaClient) BatchDeleteInstanceTopicInvoker(request *model.BatchDeleteInstanceTopicRequest) *BatchDeleteInstanceTopicInvoker {
	requestDef := GenReqDefForBatchDeleteInstanceTopic()
	return &BatchDeleteInstanceTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteInstanceUsers 批量删除用户
//
// 批量删除Kafka实例的用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) BatchDeleteInstanceUsers(request *model.BatchDeleteInstanceUsersRequest) (*model.BatchDeleteInstanceUsersResponse, error) {
	requestDef := GenReqDefForBatchDeleteInstanceUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteInstanceUsersResponse), nil
	}
}

// BatchDeleteInstanceUsersInvoker 批量删除用户
func (c *KafkaClient) BatchDeleteInstanceUsersInvoker(request *model.BatchDeleteInstanceUsersRequest) *BatchDeleteInstanceUsersInvoker {
	requestDef := GenReqDefForBatchDeleteInstanceUsers()
	return &BatchDeleteInstanceUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRestartOrDeleteInstances 批量重启或删除实例
//
// 批量重启或删除实例。
//
// 在实例重启过程中，客户端的生产与消费消息等请求会被拒绝。
//
// 实例删除后，实例中原有的数据将被删除，且没有备份，请谨慎操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) BatchRestartOrDeleteInstances(request *model.BatchRestartOrDeleteInstancesRequest) (*model.BatchRestartOrDeleteInstancesResponse, error) {
	requestDef := GenReqDefForBatchRestartOrDeleteInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRestartOrDeleteInstancesResponse), nil
	}
}

// BatchRestartOrDeleteInstancesInvoker 批量重启或删除实例
func (c *KafkaClient) BatchRestartOrDeleteInstancesInvoker(request *model.BatchRestartOrDeleteInstancesRequest) *BatchRestartOrDeleteInstancesInvoker {
	requestDef := GenReqDefForBatchRestartOrDeleteInstances()
	return &BatchRestartOrDeleteInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConnector 创建实例的转储节点
//
// 创建实例的转储节点。
//
// **当前通过调用API，只支持按需实例创建转储节点。**
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) CreateConnector(request *model.CreateConnectorRequest) (*model.CreateConnectorResponse, error) {
	requestDef := GenReqDefForCreateConnector()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConnectorResponse), nil
	}
}

// CreateConnectorInvoker 创建实例的转储节点
func (c *KafkaClient) CreateConnectorInvoker(request *model.CreateConnectorRequest) *CreateConnectorInvoker {
	requestDef := GenReqDefForCreateConnector()
	return &CreateConnectorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstanceTopic Kafka实例创建Topic
//
// 该接口用于向Kafka实例创建Topic。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) CreateInstanceTopic(request *model.CreateInstanceTopicRequest) (*model.CreateInstanceTopicResponse, error) {
	requestDef := GenReqDefForCreateInstanceTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceTopicResponse), nil
	}
}

// CreateInstanceTopicInvoker Kafka实例创建Topic
func (c *KafkaClient) CreateInstanceTopicInvoker(request *model.CreateInstanceTopicRequest) *CreateInstanceTopicInvoker {
	requestDef := GenReqDefForCreateInstanceTopic()
	return &CreateInstanceTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstanceUser 创建用户
//
// 创建Kafka实例的用户，用户可连接开启SASL的Kafka实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) CreateInstanceUser(request *model.CreateInstanceUserRequest) (*model.CreateInstanceUserResponse, error) {
	requestDef := GenReqDefForCreateInstanceUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceUserResponse), nil
	}
}

// CreateInstanceUserInvoker 创建用户
func (c *KafkaClient) CreateInstanceUserInvoker(request *model.CreateInstanceUserRequest) *CreateInstanceUserInvoker {
	requestDef := GenReqDefForCreateInstanceUser()
	return &CreateInstanceUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePartition 新增Kafka实例指定Topic分区
//
// 新增Kafka实例指定Topic分区。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) CreatePartition(request *model.CreatePartitionRequest) (*model.CreatePartitionResponse, error) {
	requestDef := GenReqDefForCreatePartition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePartitionResponse), nil
	}
}

// CreatePartitionInvoker 新增Kafka实例指定Topic分区
func (c *KafkaClient) CreatePartitionInvoker(request *model.CreatePartitionRequest) *CreatePartitionInvoker {
	requestDef := GenReqDefForCreatePartition()
	return &CreatePartitionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePostPaidInstance 创建实例
//
// [创建按需计费类型的Kafka实例。](tag:hc,hk,hws,hws_hk,otc,hws_ocb,ctc,sbc,hk_sbc,cmcc,hws_eu)[创建kafka实例。](tag:ocb)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) CreatePostPaidInstance(request *model.CreatePostPaidInstanceRequest) (*model.CreatePostPaidInstanceResponse, error) {
	requestDef := GenReqDefForCreatePostPaidInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePostPaidInstanceResponse), nil
	}
}

// CreatePostPaidInstanceInvoker 创建实例
func (c *KafkaClient) CreatePostPaidInstanceInvoker(request *model.CreatePostPaidInstanceRequest) *CreatePostPaidInstanceInvoker {
	requestDef := GenReqDefForCreatePostPaidInstance()
	return &CreatePostPaidInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSinkTask 创建转储任务
//
// 创建转储任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) CreateSinkTask(request *model.CreateSinkTaskRequest) (*model.CreateSinkTaskResponse, error) {
	requestDef := GenReqDefForCreateSinkTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSinkTaskResponse), nil
	}
}

// CreateSinkTaskInvoker 创建转储任务
func (c *KafkaClient) CreateSinkTaskInvoker(request *model.CreateSinkTaskRequest) *CreateSinkTaskInvoker {
	requestDef := GenReqDefForCreateSinkTask()
	return &CreateSinkTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBackgroundTask 删除后台任务管理中的指定记录
//
// 删除后台任务管理中的指定记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) DeleteBackgroundTask(request *model.DeleteBackgroundTaskRequest) (*model.DeleteBackgroundTaskResponse, error) {
	requestDef := GenReqDefForDeleteBackgroundTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBackgroundTaskResponse), nil
	}
}

// DeleteBackgroundTaskInvoker 删除后台任务管理中的指定记录
func (c *KafkaClient) DeleteBackgroundTaskInvoker(request *model.DeleteBackgroundTaskRequest) *DeleteBackgroundTaskInvoker {
	requestDef := GenReqDefForDeleteBackgroundTask()
	return &DeleteBackgroundTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstance 删除指定的实例
//
// 删除指定的实例，释放该实例的所有资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

// DeleteInstanceInvoker 删除指定的实例
func (c *KafkaClient) DeleteInstanceInvoker(request *model.DeleteInstanceRequest) *DeleteInstanceInvoker {
	requestDef := GenReqDefForDeleteInstance()
	return &DeleteInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSinkTask 删除单个转储任务
//
// 删除单个转储任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) DeleteSinkTask(request *model.DeleteSinkTaskRequest) (*model.DeleteSinkTaskResponse, error) {
	requestDef := GenReqDefForDeleteSinkTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSinkTaskResponse), nil
	}
}

// DeleteSinkTaskInvoker 删除单个转储任务
func (c *KafkaClient) DeleteSinkTaskInvoker(request *model.DeleteSinkTaskRequest) *DeleteSinkTaskInvoker {
	requestDef := GenReqDefForDeleteSinkTask()
	return &DeleteSinkTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailableZones 查询可用区信息
//
// 在创建实例时，需要配置实例所在的可用区ID，可通过该接口查询可用区的ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ListAvailableZones(request *model.ListAvailableZonesRequest) (*model.ListAvailableZonesResponse, error) {
	requestDef := GenReqDefForListAvailableZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailableZonesResponse), nil
	}
}

// ListAvailableZonesInvoker 查询可用区信息
func (c *KafkaClient) ListAvailableZonesInvoker(request *model.ListAvailableZonesRequest) *ListAvailableZonesInvoker {
	requestDef := GenReqDefForListAvailableZones()
	return &ListAvailableZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBackgroundTasks 查询实例的后台任务列表
//
// 查询实例的后台任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ListBackgroundTasks(request *model.ListBackgroundTasksRequest) (*model.ListBackgroundTasksResponse, error) {
	requestDef := GenReqDefForListBackgroundTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackgroundTasksResponse), nil
	}
}

// ListBackgroundTasksInvoker 查询实例的后台任务列表
func (c *KafkaClient) ListBackgroundTasksInvoker(request *model.ListBackgroundTasksRequest) *ListBackgroundTasksInvoker {
	requestDef := GenReqDefForListBackgroundTasks()
	return &ListBackgroundTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEngineProducts 查询产品规格列表
//
// 查询产品规格列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ListEngineProducts(request *model.ListEngineProductsRequest) (*model.ListEngineProductsResponse, error) {
	requestDef := GenReqDefForListEngineProducts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEngineProductsResponse), nil
	}
}

// ListEngineProductsInvoker 查询产品规格列表
func (c *KafkaClient) ListEngineProductsInvoker(request *model.ListEngineProductsRequest) *ListEngineProductsInvoker {
	requestDef := GenReqDefForListEngineProducts()
	return &ListEngineProductsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceConsumerGroups 查询所有消费组
//
// 查询所有消费组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ListInstanceConsumerGroups(request *model.ListInstanceConsumerGroupsRequest) (*model.ListInstanceConsumerGroupsResponse, error) {
	requestDef := GenReqDefForListInstanceConsumerGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceConsumerGroupsResponse), nil
	}
}

// ListInstanceConsumerGroupsInvoker 查询所有消费组
func (c *KafkaClient) ListInstanceConsumerGroupsInvoker(request *model.ListInstanceConsumerGroupsRequest) *ListInstanceConsumerGroupsInvoker {
	requestDef := GenReqDefForListInstanceConsumerGroups()
	return &ListInstanceConsumerGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceTopics Kafka实例查询Topic
//
// 该接口用于查询指定Kafka实例的Topic详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ListInstanceTopics(request *model.ListInstanceTopicsRequest) (*model.ListInstanceTopicsResponse, error) {
	requestDef := GenReqDefForListInstanceTopics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceTopicsResponse), nil
	}
}

// ListInstanceTopicsInvoker Kafka实例查询Topic
func (c *KafkaClient) ListInstanceTopicsInvoker(request *model.ListInstanceTopicsRequest) *ListInstanceTopicsInvoker {
	requestDef := GenReqDefForListInstanceTopics()
	return &ListInstanceTopicsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstances 查询所有实例列表
//
// 查询租户的实例列表，支持按照条件查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

// ListInstancesInvoker 查询所有实例列表
func (c *KafkaClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProducts 查询产品规格列表
//
// 在创建kafka实例时，需要配置订购的产品ID（即product_id），可通过该接口查询产品规格。
//
// 例如，要订购按需计费、基准带宽为100MB的kafka实例，可从接口响应消息中，查找Hourly的消息体，然后找到bandwidth为100MB的记录对应的product_id，该product_id的值即是创建上述kafka实例时需要配置的产品ID。
//
// 同时，unavailable_zones字段表示资源不足的可用区列表，如果为空，则表示所有可用区都有资源，如果不为空，则表示字段值的可用区没有资源。所以必须确保您购买的资源所在的可用区有资源，不在该字段列表内。
//
// [例如，响应消息中bandwidth字段为1200MB的记录，unavailable_zones字段包含cn-east-2b、cn-east-2a和cn-east-2d，表示在华东-上海2的可用区1、可用区2、可用区3都没有该资源。](tag:hc,hws)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ListProducts(request *model.ListProductsRequest) (*model.ListProductsResponse, error) {
	requestDef := GenReqDefForListProducts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProductsResponse), nil
	}
}

// ListProductsInvoker 查询产品规格列表
func (c *KafkaClient) ListProductsInvoker(request *model.ListProductsRequest) *ListProductsInvoker {
	requestDef := GenReqDefForListProducts()
	return &ListProductsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSinkTasks 查询转储任务列表
//
// 查询转储任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ListSinkTasks(request *model.ListSinkTasksRequest) (*model.ListSinkTasksResponse, error) {
	requestDef := GenReqDefForListSinkTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSinkTasksResponse), nil
	}
}

// ListSinkTasksInvoker 查询转储任务列表
func (c *KafkaClient) ListSinkTasksInvoker(request *model.ListSinkTasksRequest) *ListSinkTasksInvoker {
	requestDef := GenReqDefForListSinkTasks()
	return &ListSinkTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetManagerPassword 重置Manager密码
//
// 重置Manager密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ResetManagerPassword(request *model.ResetManagerPasswordRequest) (*model.ResetManagerPasswordResponse, error) {
	requestDef := GenReqDefForResetManagerPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetManagerPasswordResponse), nil
	}
}

// ResetManagerPasswordInvoker 重置Manager密码
func (c *KafkaClient) ResetManagerPasswordInvoker(request *model.ResetManagerPasswordRequest) *ResetManagerPasswordInvoker {
	requestDef := GenReqDefForResetManagerPassword()
	return &ResetManagerPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetMessageOffset 重置消费组消费进度到指定位置
//
// Kafka实例不支持在线重置消费进度。在执行重置消费进度之前，必须停止被重置消费组客户端。
//
// &gt; 停止待重置消费组客户端，然后等待一段时间（即ConsumerConfig.SESSION_TIMEOUT_MS_CONFIG配置的时间，默认为1000毫秒）后，服务端才认为此消费组客户端已下线。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ResetMessageOffset(request *model.ResetMessageOffsetRequest) (*model.ResetMessageOffsetResponse, error) {
	requestDef := GenReqDefForResetMessageOffset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetMessageOffsetResponse), nil
	}
}

// ResetMessageOffsetInvoker 重置消费组消费进度到指定位置
func (c *KafkaClient) ResetMessageOffsetInvoker(request *model.ResetMessageOffsetRequest) *ResetMessageOffsetInvoker {
	requestDef := GenReqDefForResetMessageOffset()
	return &ResetMessageOffsetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetPassword 重置密码
//
// 重置密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ResetPassword(request *model.ResetPasswordRequest) (*model.ResetPasswordResponse, error) {
	requestDef := GenReqDefForResetPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetPasswordResponse), nil
	}
}

// ResetPasswordInvoker 重置密码
func (c *KafkaClient) ResetPasswordInvoker(request *model.ResetPasswordRequest) *ResetPasswordInvoker {
	requestDef := GenReqDefForResetPassword()
	return &ResetPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetUserPasswrod 重置用户密码
//
// 重置用户密码
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ResetUserPasswrod(request *model.ResetUserPasswrodRequest) (*model.ResetUserPasswrodResponse, error) {
	requestDef := GenReqDefForResetUserPasswrod()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetUserPasswrodResponse), nil
	}
}

// ResetUserPasswrodInvoker 重置用户密码
func (c *KafkaClient) ResetUserPasswrodInvoker(request *model.ResetUserPasswrodRequest) *ResetUserPasswrodInvoker {
	requestDef := GenReqDefForResetUserPasswrod()
	return &ResetUserPasswrodInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeInstance 实例规格变更
//
// 实例规格变更。
//
// **当前通过调用API，只支持按需实例进行实例规格变更。**
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ResizeInstance(request *model.ResizeInstanceRequest) (*model.ResizeInstanceResponse, error) {
	requestDef := GenReqDefForResizeInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeInstanceResponse), nil
	}
}

// ResizeInstanceInvoker 实例规格变更
func (c *KafkaClient) ResizeInstanceInvoker(request *model.ResizeInstanceRequest) *ResizeInstanceInvoker {
	requestDef := GenReqDefForResizeInstance()
	return &ResizeInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartManager 重启Manager
//
// 重启Manager。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) RestartManager(request *model.RestartManagerRequest) (*model.RestartManagerResponse, error) {
	requestDef := GenReqDefForRestartManager()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartManagerResponse), nil
	}
}

// RestartManagerInvoker 重启Manager
func (c *KafkaClient) RestartManagerInvoker(request *model.RestartManagerRequest) *RestartManagerInvoker {
	requestDef := GenReqDefForRestartManager()
	return &RestartManagerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBackgroundTask 查询后台任务管理中的指定记录
//
// 查询后台任务管理中的指定记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ShowBackgroundTask(request *model.ShowBackgroundTaskRequest) (*model.ShowBackgroundTaskResponse, error) {
	requestDef := GenReqDefForShowBackgroundTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackgroundTaskResponse), nil
	}
}

// ShowBackgroundTaskInvoker 查询后台任务管理中的指定记录
func (c *KafkaClient) ShowBackgroundTaskInvoker(request *model.ShowBackgroundTaskRequest) *ShowBackgroundTaskInvoker {
	requestDef := GenReqDefForShowBackgroundTask()
	return &ShowBackgroundTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCesHierarchy 查询实例在CES的监控层级关系
//
// 查询实例在CES的监控层级关系。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ShowCesHierarchy(request *model.ShowCesHierarchyRequest) (*model.ShowCesHierarchyResponse, error) {
	requestDef := GenReqDefForShowCesHierarchy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCesHierarchyResponse), nil
	}
}

// ShowCesHierarchyInvoker 查询实例在CES的监控层级关系
func (c *KafkaClient) ShowCesHierarchyInvoker(request *model.ShowCesHierarchyRequest) *ShowCesHierarchyInvoker {
	requestDef := GenReqDefForShowCesHierarchy()
	return &ShowCesHierarchyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCluster 查询Kafka集群元数据信息
//
// 查询Kafka集群元数据信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ShowCluster(request *model.ShowClusterRequest) (*model.ShowClusterResponse, error) {
	requestDef := GenReqDefForShowCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterResponse), nil
	}
}

// ShowClusterInvoker 查询Kafka集群元数据信息
func (c *KafkaClient) ShowClusterInvoker(request *model.ShowClusterRequest) *ShowClusterInvoker {
	requestDef := GenReqDefForShowCluster()
	return &ShowClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCoordinators 查询Kafka实例的协调器信息
//
// 查询Kafka实例的协调器信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ShowCoordinators(request *model.ShowCoordinatorsRequest) (*model.ShowCoordinatorsResponse, error) {
	requestDef := GenReqDefForShowCoordinators()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCoordinatorsResponse), nil
	}
}

// ShowCoordinatorsInvoker 查询Kafka实例的协调器信息
func (c *KafkaClient) ShowCoordinatorsInvoker(request *model.ShowCoordinatorsRequest) *ShowCoordinatorsInvoker {
	requestDef := GenReqDefForShowCoordinators()
	return &ShowCoordinatorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroups 查询消费组信息
//
// 查询消费组信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ShowGroups(request *model.ShowGroupsRequest) (*model.ShowGroupsResponse, error) {
	requestDef := GenReqDefForShowGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupsResponse), nil
	}
}

// ShowGroupsInvoker 查询消费组信息
func (c *KafkaClient) ShowGroupsInvoker(request *model.ShowGroupsRequest) *ShowGroupsInvoker {
	requestDef := GenReqDefForShowGroups()
	return &ShowGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstance 查询指定实例
//
// 查询指定实例的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ShowInstance(request *model.ShowInstanceRequest) (*model.ShowInstanceResponse, error) {
	requestDef := GenReqDefForShowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResponse), nil
	}
}

// ShowInstanceInvoker 查询指定实例
func (c *KafkaClient) ShowInstanceInvoker(request *model.ShowInstanceRequest) *ShowInstanceInvoker {
	requestDef := GenReqDefForShowInstance()
	return &ShowInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceExtendProductInfo 查询实例的扩容规格列表
//
// 查询实例的扩容规格列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ShowInstanceExtendProductInfo(request *model.ShowInstanceExtendProductInfoRequest) (*model.ShowInstanceExtendProductInfoResponse, error) {
	requestDef := GenReqDefForShowInstanceExtendProductInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceExtendProductInfoResponse), nil
	}
}

// ShowInstanceExtendProductInfoInvoker 查询实例的扩容规格列表
func (c *KafkaClient) ShowInstanceExtendProductInfoInvoker(request *model.ShowInstanceExtendProductInfoRequest) *ShowInstanceExtendProductInfoInvoker {
	requestDef := GenReqDefForShowInstanceExtendProductInfo()
	return &ShowInstanceExtendProductInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceMessages 查询消息
//
// 查询消息的偏移量和消息内容。
// 先根据时间戳查询消息的偏移量，再根据偏移量查询消息内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ShowInstanceMessages(request *model.ShowInstanceMessagesRequest) (*model.ShowInstanceMessagesResponse, error) {
	requestDef := GenReqDefForShowInstanceMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceMessagesResponse), nil
	}
}

// ShowInstanceMessagesInvoker 查询消息
func (c *KafkaClient) ShowInstanceMessagesInvoker(request *model.ShowInstanceMessagesRequest) *ShowInstanceMessagesInvoker {
	requestDef := GenReqDefForShowInstanceMessages()
	return &ShowInstanceMessagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceTopicDetail 查询Kafka实例Topic详细信息
//
// 查询Kafka实例Topic详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ShowInstanceTopicDetail(request *model.ShowInstanceTopicDetailRequest) (*model.ShowInstanceTopicDetailResponse, error) {
	requestDef := GenReqDefForShowInstanceTopicDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceTopicDetailResponse), nil
	}
}

// ShowInstanceTopicDetailInvoker 查询Kafka实例Topic详细信息
func (c *KafkaClient) ShowInstanceTopicDetailInvoker(request *model.ShowInstanceTopicDetailRequest) *ShowInstanceTopicDetailInvoker {
	requestDef := GenReqDefForShowInstanceTopicDetail()
	return &ShowInstanceTopicDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceUsers 查询用户列表
//
// 查询用户列表。
//
// Kafka实例开启SASL功能时，才支持多用户管理的功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ShowInstanceUsers(request *model.ShowInstanceUsersRequest) (*model.ShowInstanceUsersResponse, error) {
	requestDef := GenReqDefForShowInstanceUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceUsersResponse), nil
	}
}

// ShowInstanceUsersInvoker 查询用户列表
func (c *KafkaClient) ShowInstanceUsersInvoker(request *model.ShowInstanceUsersRequest) *ShowInstanceUsersInvoker {
	requestDef := GenReqDefForShowInstanceUsers()
	return &ShowInstanceUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowKafkaProjectTags 查询项目标签
//
// 查询项目标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ShowKafkaProjectTags(request *model.ShowKafkaProjectTagsRequest) (*model.ShowKafkaProjectTagsResponse, error) {
	requestDef := GenReqDefForShowKafkaProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKafkaProjectTagsResponse), nil
	}
}

// ShowKafkaProjectTagsInvoker 查询项目标签
func (c *KafkaClient) ShowKafkaProjectTagsInvoker(request *model.ShowKafkaProjectTagsRequest) *ShowKafkaProjectTagsInvoker {
	requestDef := GenReqDefForShowKafkaProjectTags()
	return &ShowKafkaProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowKafkaTags 查询实例标签
//
// 查询实例标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ShowKafkaTags(request *model.ShowKafkaTagsRequest) (*model.ShowKafkaTagsResponse, error) {
	requestDef := GenReqDefForShowKafkaTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKafkaTagsResponse), nil
	}
}

// ShowKafkaTagsInvoker 查询实例标签
func (c *KafkaClient) ShowKafkaTagsInvoker(request *model.ShowKafkaTagsRequest) *ShowKafkaTagsInvoker {
	requestDef := GenReqDefForShowKafkaTags()
	return &ShowKafkaTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowKafkaTopicPartitionDiskusage 查询topic的磁盘存储情况
//
// 查询topic在Broker上磁盘占用情况。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ShowKafkaTopicPartitionDiskusage(request *model.ShowKafkaTopicPartitionDiskusageRequest) (*model.ShowKafkaTopicPartitionDiskusageResponse, error) {
	requestDef := GenReqDefForShowKafkaTopicPartitionDiskusage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKafkaTopicPartitionDiskusageResponse), nil
	}
}

// ShowKafkaTopicPartitionDiskusageInvoker 查询topic的磁盘存储情况
func (c *KafkaClient) ShowKafkaTopicPartitionDiskusageInvoker(request *model.ShowKafkaTopicPartitionDiskusageRequest) *ShowKafkaTopicPartitionDiskusageInvoker {
	requestDef := GenReqDefForShowKafkaTopicPartitionDiskusage()
	return &ShowKafkaTopicPartitionDiskusageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMaintainWindows 查询维护时间窗时间段
//
// 查询维护时间窗开始时间和结束时间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ShowMaintainWindows(request *model.ShowMaintainWindowsRequest) (*model.ShowMaintainWindowsResponse, error) {
	requestDef := GenReqDefForShowMaintainWindows()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMaintainWindowsResponse), nil
	}
}

// ShowMaintainWindowsInvoker 查询维护时间窗时间段
func (c *KafkaClient) ShowMaintainWindowsInvoker(request *model.ShowMaintainWindowsRequest) *ShowMaintainWindowsInvoker {
	requestDef := GenReqDefForShowMaintainWindows()
	return &ShowMaintainWindowsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMessages 查询分区指定时间段的消息
//
// 查询分区指定时间段的消息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ShowMessages(request *model.ShowMessagesRequest) (*model.ShowMessagesResponse, error) {
	requestDef := GenReqDefForShowMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMessagesResponse), nil
	}
}

// ShowMessagesInvoker 查询分区指定时间段的消息
func (c *KafkaClient) ShowMessagesInvoker(request *model.ShowMessagesRequest) *ShowMessagesInvoker {
	requestDef := GenReqDefForShowMessages()
	return &ShowMessagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPartitionBeginningMessage 查询分区最早消息的位置
//
// 查询分区最早消息的位置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ShowPartitionBeginningMessage(request *model.ShowPartitionBeginningMessageRequest) (*model.ShowPartitionBeginningMessageResponse, error) {
	requestDef := GenReqDefForShowPartitionBeginningMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPartitionBeginningMessageResponse), nil
	}
}

// ShowPartitionBeginningMessageInvoker 查询分区最早消息的位置
func (c *KafkaClient) ShowPartitionBeginningMessageInvoker(request *model.ShowPartitionBeginningMessageRequest) *ShowPartitionBeginningMessageInvoker {
	requestDef := GenReqDefForShowPartitionBeginningMessage()
	return &ShowPartitionBeginningMessageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPartitionEndMessage 查询分区最新消息的位置
//
// 查询分区最新消息的位置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ShowPartitionEndMessage(request *model.ShowPartitionEndMessageRequest) (*model.ShowPartitionEndMessageResponse, error) {
	requestDef := GenReqDefForShowPartitionEndMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPartitionEndMessageResponse), nil
	}
}

// ShowPartitionEndMessageInvoker 查询分区最新消息的位置
func (c *KafkaClient) ShowPartitionEndMessageInvoker(request *model.ShowPartitionEndMessageRequest) *ShowPartitionEndMessageInvoker {
	requestDef := GenReqDefForShowPartitionEndMessage()
	return &ShowPartitionEndMessageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPartitionMessage 查询分区指定偏移量的消息
//
// 查询分区指定偏移量的消息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ShowPartitionMessage(request *model.ShowPartitionMessageRequest) (*model.ShowPartitionMessageResponse, error) {
	requestDef := GenReqDefForShowPartitionMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPartitionMessageResponse), nil
	}
}

// ShowPartitionMessageInvoker 查询分区指定偏移量的消息
func (c *KafkaClient) ShowPartitionMessageInvoker(request *model.ShowPartitionMessageRequest) *ShowPartitionMessageInvoker {
	requestDef := GenReqDefForShowPartitionMessage()
	return &ShowPartitionMessageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSinkTaskDetail 查询单个转储任务
//
// 查询单个转储任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ShowSinkTaskDetail(request *model.ShowSinkTaskDetailRequest) (*model.ShowSinkTaskDetailResponse, error) {
	requestDef := GenReqDefForShowSinkTaskDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSinkTaskDetailResponse), nil
	}
}

// ShowSinkTaskDetailInvoker 查询单个转储任务
func (c *KafkaClient) ShowSinkTaskDetailInvoker(request *model.ShowSinkTaskDetailRequest) *ShowSinkTaskDetailInvoker {
	requestDef := GenReqDefForShowSinkTaskDetail()
	return &ShowSinkTaskDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTopicAccessPolicy 查询用户权限
//
// 查询用户权限。
//
// Kafka实例开启SASL功能时，才支持多用户管理的功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) ShowTopicAccessPolicy(request *model.ShowTopicAccessPolicyRequest) (*model.ShowTopicAccessPolicyResponse, error) {
	requestDef := GenReqDefForShowTopicAccessPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTopicAccessPolicyResponse), nil
	}
}

// ShowTopicAccessPolicyInvoker 查询用户权限
func (c *KafkaClient) ShowTopicAccessPolicyInvoker(request *model.ShowTopicAccessPolicyRequest) *ShowTopicAccessPolicyInvoker {
	requestDef := GenReqDefForShowTopicAccessPolicy()
	return &ShowTopicAccessPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstance 修改实例信息
//
// 修改实例信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) UpdateInstance(request *model.UpdateInstanceRequest) (*model.UpdateInstanceResponse, error) {
	requestDef := GenReqDefForUpdateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceResponse), nil
	}
}

// UpdateInstanceInvoker 修改实例信息
func (c *KafkaClient) UpdateInstanceInvoker(request *model.UpdateInstanceRequest) *UpdateInstanceInvoker {
	requestDef := GenReqDefForUpdateInstance()
	return &UpdateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceAutoCreateTopic 开启或关闭实例自动创建topic功能
//
// 开启或关闭实例自动创建topic功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) UpdateInstanceAutoCreateTopic(request *model.UpdateInstanceAutoCreateTopicRequest) (*model.UpdateInstanceAutoCreateTopicResponse, error) {
	requestDef := GenReqDefForUpdateInstanceAutoCreateTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceAutoCreateTopicResponse), nil
	}
}

// UpdateInstanceAutoCreateTopicInvoker 开启或关闭实例自动创建topic功能
func (c *KafkaClient) UpdateInstanceAutoCreateTopicInvoker(request *model.UpdateInstanceAutoCreateTopicRequest) *UpdateInstanceAutoCreateTopicInvoker {
	requestDef := GenReqDefForUpdateInstanceAutoCreateTopic()
	return &UpdateInstanceAutoCreateTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceCrossVpcIp 修改实例跨VPC访问的内网IP
//
// 修改实例跨VPC访问的内网IP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) UpdateInstanceCrossVpcIp(request *model.UpdateInstanceCrossVpcIpRequest) (*model.UpdateInstanceCrossVpcIpResponse, error) {
	requestDef := GenReqDefForUpdateInstanceCrossVpcIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceCrossVpcIpResponse), nil
	}
}

// UpdateInstanceCrossVpcIpInvoker 修改实例跨VPC访问的内网IP
func (c *KafkaClient) UpdateInstanceCrossVpcIpInvoker(request *model.UpdateInstanceCrossVpcIpRequest) *UpdateInstanceCrossVpcIpInvoker {
	requestDef := GenReqDefForUpdateInstanceCrossVpcIp()
	return &UpdateInstanceCrossVpcIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceTopic 修改Kafka实例Topic
//
// 修改Kafka实例Topic
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) UpdateInstanceTopic(request *model.UpdateInstanceTopicRequest) (*model.UpdateInstanceTopicResponse, error) {
	requestDef := GenReqDefForUpdateInstanceTopic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceTopicResponse), nil
	}
}

// UpdateInstanceTopicInvoker 修改Kafka实例Topic
func (c *KafkaClient) UpdateInstanceTopicInvoker(request *model.UpdateInstanceTopicRequest) *UpdateInstanceTopicInvoker {
	requestDef := GenReqDefForUpdateInstanceTopic()
	return &UpdateInstanceTopicInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSinkTaskQuota 修改转储任务的配额
//
// 修改转储任务的配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) UpdateSinkTaskQuota(request *model.UpdateSinkTaskQuotaRequest) (*model.UpdateSinkTaskQuotaResponse, error) {
	requestDef := GenReqDefForUpdateSinkTaskQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSinkTaskQuotaResponse), nil
	}
}

// UpdateSinkTaskQuotaInvoker 修改转储任务的配额
func (c *KafkaClient) UpdateSinkTaskQuotaInvoker(request *model.UpdateSinkTaskQuotaRequest) *UpdateSinkTaskQuotaInvoker {
	requestDef := GenReqDefForUpdateSinkTaskQuota()
	return &UpdateSinkTaskQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTopicAccessPolicy 设置用户权限
//
// 设置用户权限。
//
// Kafka实例开启SASL功能时，才支持多用户管理的功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) UpdateTopicAccessPolicy(request *model.UpdateTopicAccessPolicyRequest) (*model.UpdateTopicAccessPolicyResponse, error) {
	requestDef := GenReqDefForUpdateTopicAccessPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTopicAccessPolicyResponse), nil
	}
}

// UpdateTopicAccessPolicyInvoker 设置用户权限
func (c *KafkaClient) UpdateTopicAccessPolicyInvoker(request *model.UpdateTopicAccessPolicyRequest) *UpdateTopicAccessPolicyInvoker {
	requestDef := GenReqDefForUpdateTopicAccessPolicy()
	return &UpdateTopicAccessPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTopicReplica 修改Kafka实例Topic分区的副本
//
// 修改Kafka实例Topic分区的副本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KafkaClient) UpdateTopicReplica(request *model.UpdateTopicReplicaRequest) (*model.UpdateTopicReplicaResponse, error) {
	requestDef := GenReqDefForUpdateTopicReplica()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTopicReplicaResponse), nil
	}
}

// UpdateTopicReplicaInvoker 修改Kafka实例Topic分区的副本
func (c *KafkaClient) UpdateTopicReplicaInvoker(request *model.UpdateTopicReplicaRequest) *UpdateTopicReplicaInvoker {
	requestDef := GenReqDefForUpdateTopicReplica()
	return &UpdateTopicReplicaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
