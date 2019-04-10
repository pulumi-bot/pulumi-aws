// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The ECS container definition data source allows access to details of
// a specific container within an AWS ECS service.
func LookupContainerDefinition(ctx *pulumi.Context, args *GetContainerDefinitionArgs) (*GetContainerDefinitionResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["containerName"] = args.ContainerName
		inputs["taskDefinition"] = args.TaskDefinition
	}
	outputs, err := ctx.Invoke("aws:ecs/getContainerDefinition:getContainerDefinition", inputs)
	if err != nil {
		return nil, err
	}
	return &GetContainerDefinitionResult{
		ContainerName: outputs["containerName"],
		Cpu: outputs["cpu"],
		DisableNetworking: outputs["disableNetworking"],
		DockerLabels: outputs["dockerLabels"],
		Environment: outputs["environment"],
		Image: outputs["image"],
		ImageDigest: outputs["imageDigest"],
		Memory: outputs["memory"],
		MemoryReservation: outputs["memoryReservation"],
		TaskDefinition: outputs["taskDefinition"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getContainerDefinition.
type GetContainerDefinitionArgs struct {
	// The name of the container definition
	ContainerName interface{}
	// The ARN of the task definition which contains the container
	TaskDefinition interface{}
}

// A collection of values returned by getContainerDefinition.
type GetContainerDefinitionResult struct {
	ContainerName interface{}
	// The CPU limit for this container definition
	Cpu interface{}
	// Indicator if networking is disabled
	DisableNetworking interface{}
	// Set docker labels
	DockerLabels interface{}
	// The environment in use
	Environment interface{}
	// The docker image in use, including the digest
	Image interface{}
	// The digest of the docker image in use
	ImageDigest interface{}
	// The memory limit for this container definition
	Memory interface{}
	// The soft limit (in MiB) of memory to reserve for the container. When system memory is under contention, Docker attempts to keep the container memory to this soft limit
	MemoryReservation interface{}
	TaskDefinition interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
