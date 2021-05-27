// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The ECS container definition data source allows access to details of
// a specific container within an AWS ECS service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ecs"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ecs.GetContainerDefinition(ctx, &ecs.GetContainerDefinitionArgs{
// 			TaskDefinition: aws_ecs_task_definition.Mongo.Id,
// 			ContainerName:  "mongodb",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetContainerDefinition(ctx *pulumi.Context, args *GetContainerDefinitionArgs, opts ...pulumi.InvokeOption) (*GetContainerDefinitionResult, error) {
	var rv GetContainerDefinitionResult
	err := ctx.Invoke("aws:ecs/getContainerDefinition:getContainerDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContainerDefinition.
type GetContainerDefinitionArgs struct {
	// The name of the container definition
	ContainerName string `pulumi:"containerName"`
	// The ARN of the task definition which contains the container
	TaskDefinition string `pulumi:"taskDefinition"`
}

// A collection of values returned by getContainerDefinition.
type GetContainerDefinitionResult struct {
	ContainerName string `pulumi:"containerName"`
	// The CPU limit for this container definition
	Cpu int `pulumi:"cpu"`
	// Indicator if networking is disabled
	DisableNetworking bool `pulumi:"disableNetworking"`
	// Set docker labels
	DockerLabels map[string]string `pulumi:"dockerLabels"`
	// The environment in use
	Environment map[string]string `pulumi:"environment"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The docker image in use, including the digest
	Image string `pulumi:"image"`
	// The digest of the docker image in use
	ImageDigest string `pulumi:"imageDigest"`
	// The memory limit for this container definition
	Memory int `pulumi:"memory"`
	// The soft limit (in MiB) of memory to reserve for the container. When system memory is under contention, Docker attempts to keep the container memory to this soft limit
	MemoryReservation int    `pulumi:"memoryReservation"`
	TaskDefinition    string `pulumi:"taskDefinition"`
}

func GetContainerDefinitionOutput(ctx *pulumi.Context, args GetContainerDefinitionOutputArgs, opts ...pulumi.InvokeOption) GetContainerDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetContainerDefinitionResult, error) {
			args := v.(GetContainerDefinitionArgs)
			r, err := GetContainerDefinition(ctx, &args, opts...)
			return *r, err
		}).(GetContainerDefinitionResultOutput)
}

// A collection of arguments for invoking getContainerDefinition.
type GetContainerDefinitionOutputArgs struct {
	// The name of the container definition
	ContainerName pulumi.StringInput `pulumi:"containerName"`
	// The ARN of the task definition which contains the container
	TaskDefinition pulumi.StringInput `pulumi:"taskDefinition"`
}

func (GetContainerDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContainerDefinitionArgs)(nil)).Elem()
}

// A collection of values returned by getContainerDefinition.
type GetContainerDefinitionResultOutput struct{ *pulumi.OutputState }

func (GetContainerDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContainerDefinitionResult)(nil)).Elem()
}

func (o GetContainerDefinitionResultOutput) ToGetContainerDefinitionResultOutput() GetContainerDefinitionResultOutput {
	return o
}

func (o GetContainerDefinitionResultOutput) ToGetContainerDefinitionResultOutputWithContext(ctx context.Context) GetContainerDefinitionResultOutput {
	return o
}

func (o GetContainerDefinitionResultOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerDefinitionResult) string { return v.ContainerName }).(pulumi.StringOutput)
}

// The CPU limit for this container definition
func (o GetContainerDefinitionResultOutput) Cpu() pulumi.IntOutput {
	return o.ApplyT(func(v GetContainerDefinitionResult) int { return v.Cpu }).(pulumi.IntOutput)
}

// Indicator if networking is disabled
func (o GetContainerDefinitionResultOutput) DisableNetworking() pulumi.BoolOutput {
	return o.ApplyT(func(v GetContainerDefinitionResult) bool { return v.DisableNetworking }).(pulumi.BoolOutput)
}

// Set docker labels
func (o GetContainerDefinitionResultOutput) DockerLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetContainerDefinitionResult) map[string]string { return v.DockerLabels }).(pulumi.StringMapOutput)
}

// The environment in use
func (o GetContainerDefinitionResultOutput) Environment() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetContainerDefinitionResult) map[string]string { return v.Environment }).(pulumi.StringMapOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetContainerDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The docker image in use, including the digest
func (o GetContainerDefinitionResultOutput) Image() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerDefinitionResult) string { return v.Image }).(pulumi.StringOutput)
}

// The digest of the docker image in use
func (o GetContainerDefinitionResultOutput) ImageDigest() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerDefinitionResult) string { return v.ImageDigest }).(pulumi.StringOutput)
}

// The memory limit for this container definition
func (o GetContainerDefinitionResultOutput) Memory() pulumi.IntOutput {
	return o.ApplyT(func(v GetContainerDefinitionResult) int { return v.Memory }).(pulumi.IntOutput)
}

// The soft limit (in MiB) of memory to reserve for the container. When system memory is under contention, Docker attempts to keep the container memory to this soft limit
func (o GetContainerDefinitionResultOutput) MemoryReservation() pulumi.IntOutput {
	return o.ApplyT(func(v GetContainerDefinitionResult) int { return v.MemoryReservation }).(pulumi.IntOutput)
}

func (o GetContainerDefinitionResultOutput) TaskDefinition() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerDefinitionResult) string { return v.TaskDefinition }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetContainerDefinitionResultOutput{})
}
