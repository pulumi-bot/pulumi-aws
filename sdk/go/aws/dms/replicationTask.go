// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a DMS (Data Migration Service) replication task resource. DMS replication tasks can be created, updated, deleted, and imported.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/dms"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dms.NewReplicationTask(ctx, "test", &dms.ReplicationTaskArgs{
// 			CdcStartTime:            pulumi.String("1484346880"),
// 			MigrationType:           pulumi.String("full-load"),
// 			ReplicationInstanceArn:  pulumi.Any(aws_dms_replication_instance.Test - dms - replication - instance - tf.Replication_instance_arn),
// 			ReplicationTaskId:       pulumi.String("test-dms-replication-task-tf"),
// 			ReplicationTaskSettings: pulumi.String("..."),
// 			SourceEndpointArn:       pulumi.Any(aws_dms_endpoint.Test - dms - source - endpoint - tf.Endpoint_arn),
// 			TableMappings:           pulumi.String(fmt.Sprintf("%v%v%v%v%v", "{\"rules\":[{\"rule-type\":\"selection\",\"rule-id\":\"1\",\"rule-name\":\"1\",\"object-locator\":{\"schema-name\":\"", "%", "\",\"table-name\":\"", "%", "\"},\"rule-action\":\"include\"}]}")),
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("test"),
// 			},
// 			TargetEndpointArn: pulumi.Any(aws_dms_endpoint.Test - dms - target - endpoint - tf.Endpoint_arn),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Replication tasks can be imported using the `replication_task_id`, e.g.
//
// ```sh
//  $ pulumi import aws:dms/replicationTask:ReplicationTask test test-dms-replication-task-tf
// ```
type ReplicationTask struct {
	pulumi.CustomResourceState

	// The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
	CdcStartTime pulumi.StringPtrOutput `pulumi:"cdcStartTime"`
	// The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
	MigrationType pulumi.StringOutput `pulumi:"migrationType"`
	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn pulumi.StringOutput `pulumi:"replicationInstanceArn"`
	// The Amazon Resource Name (ARN) for the replication task.
	ReplicationTaskArn pulumi.StringOutput `pulumi:"replicationTaskArn"`
	// The replication task identifier.
	ReplicationTaskId pulumi.StringOutput `pulumi:"replicationTaskId"`
	// An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
	ReplicationTaskSettings pulumi.StringPtrOutput `pulumi:"replicationTaskSettings"`
	// The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
	SourceEndpointArn pulumi.StringOutput `pulumi:"sourceEndpointArn"`
	// An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
	TableMappings pulumi.StringOutput `pulumi:"tableMappings"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
	TargetEndpointArn pulumi.StringOutput `pulumi:"targetEndpointArn"`
}

// NewReplicationTask registers a new resource with the given unique name, arguments, and options.
func NewReplicationTask(ctx *pulumi.Context,
	name string, args *ReplicationTaskArgs, opts ...pulumi.ResourceOption) (*ReplicationTask, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MigrationType == nil {
		return nil, errors.New("invalid value for required argument 'MigrationType'")
	}
	if args.ReplicationInstanceArn == nil {
		return nil, errors.New("invalid value for required argument 'ReplicationInstanceArn'")
	}
	if args.ReplicationTaskId == nil {
		return nil, errors.New("invalid value for required argument 'ReplicationTaskId'")
	}
	if args.SourceEndpointArn == nil {
		return nil, errors.New("invalid value for required argument 'SourceEndpointArn'")
	}
	if args.TableMappings == nil {
		return nil, errors.New("invalid value for required argument 'TableMappings'")
	}
	if args.TargetEndpointArn == nil {
		return nil, errors.New("invalid value for required argument 'TargetEndpointArn'")
	}
	var resource ReplicationTask
	err := ctx.RegisterResource("aws:dms/replicationTask:ReplicationTask", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationTask gets an existing ReplicationTask resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationTaskState, opts ...pulumi.ResourceOption) (*ReplicationTask, error) {
	var resource ReplicationTask
	err := ctx.ReadResource("aws:dms/replicationTask:ReplicationTask", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationTask resources.
type replicationTaskState struct {
	// The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
	CdcStartTime *string `pulumi:"cdcStartTime"`
	// The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
	MigrationType *string `pulumi:"migrationType"`
	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn *string `pulumi:"replicationInstanceArn"`
	// The Amazon Resource Name (ARN) for the replication task.
	ReplicationTaskArn *string `pulumi:"replicationTaskArn"`
	// The replication task identifier.
	ReplicationTaskId *string `pulumi:"replicationTaskId"`
	// An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
	ReplicationTaskSettings *string `pulumi:"replicationTaskSettings"`
	// The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
	SourceEndpointArn *string `pulumi:"sourceEndpointArn"`
	// An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
	TableMappings *string `pulumi:"tableMappings"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
	TargetEndpointArn *string `pulumi:"targetEndpointArn"`
}

type ReplicationTaskState struct {
	// The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
	CdcStartTime pulumi.StringPtrInput
	// The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
	MigrationType pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) for the replication task.
	ReplicationTaskArn pulumi.StringPtrInput
	// The replication task identifier.
	ReplicationTaskId pulumi.StringPtrInput
	// An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
	ReplicationTaskSettings pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
	SourceEndpointArn pulumi.StringPtrInput
	// An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
	TableMappings pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
	TargetEndpointArn pulumi.StringPtrInput
}

func (ReplicationTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationTaskState)(nil)).Elem()
}

type replicationTaskArgs struct {
	// The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
	CdcStartTime *string `pulumi:"cdcStartTime"`
	// The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
	MigrationType string `pulumi:"migrationType"`
	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn string `pulumi:"replicationInstanceArn"`
	// The replication task identifier.
	ReplicationTaskId string `pulumi:"replicationTaskId"`
	// An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
	ReplicationTaskSettings *string `pulumi:"replicationTaskSettings"`
	// The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
	SourceEndpointArn string `pulumi:"sourceEndpointArn"`
	// An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
	TableMappings string `pulumi:"tableMappings"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
	TargetEndpointArn string `pulumi:"targetEndpointArn"`
}

// The set of arguments for constructing a ReplicationTask resource.
type ReplicationTaskArgs struct {
	// The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
	CdcStartTime pulumi.StringPtrInput
	// The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
	MigrationType pulumi.StringInput
	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn pulumi.StringInput
	// The replication task identifier.
	ReplicationTaskId pulumi.StringInput
	// An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
	ReplicationTaskSettings pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
	SourceEndpointArn pulumi.StringInput
	// An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
	TableMappings pulumi.StringInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
	TargetEndpointArn pulumi.StringInput
}

func (ReplicationTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationTaskArgs)(nil)).Elem()
}

type ReplicationTaskInput interface {
	pulumi.Input

	ToReplicationTaskOutput() ReplicationTaskOutput
	ToReplicationTaskOutputWithContext(ctx context.Context) ReplicationTaskOutput
}

func (ReplicationTask) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationTask)(nil)).Elem()
}

func (i ReplicationTask) ToReplicationTaskOutput() ReplicationTaskOutput {
	return i.ToReplicationTaskOutputWithContext(context.Background())
}

func (i ReplicationTask) ToReplicationTaskOutputWithContext(ctx context.Context) ReplicationTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationTaskOutput)
}

type ReplicationTaskOutput struct {
	*pulumi.OutputState
}

func (ReplicationTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationTaskOutput)(nil)).Elem()
}

func (o ReplicationTaskOutput) ToReplicationTaskOutput() ReplicationTaskOutput {
	return o
}

func (o ReplicationTaskOutput) ToReplicationTaskOutputWithContext(ctx context.Context) ReplicationTaskOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ReplicationTaskOutput{})
}
