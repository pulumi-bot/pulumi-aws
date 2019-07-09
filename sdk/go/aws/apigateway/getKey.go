// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the name and value of a pre-existing API Key, for
// example to supply credentials for a dependency microservice.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/api_gateway_api_key.html.markdown.
func LookupKey(ctx *pulumi.Context, args *GetKeyArgs) (*GetKeyResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["id"] = args.Id
	}
	outputs, err := ctx.Invoke("aws:apigateway/getKey:getKey", inputs)
	if err != nil {
		return nil, err
	}
	return &GetKeyResult{
		Id: outputs["id"],
		Name: outputs["name"],
		Value: outputs["value"],
	}, nil
}

// A collection of arguments for invoking getKey.
type GetKeyArgs struct {
	// The ID of the API Key to look up.
	Id interface{}
}

// A collection of values returned by getKey.
type GetKeyResult struct {
	// Set to the ID of the API Key.
	Id interface{}
	// Set to the name of the API Key.
	Name interface{}
	// Set to the value of the API Key.
	Value interface{}
}
