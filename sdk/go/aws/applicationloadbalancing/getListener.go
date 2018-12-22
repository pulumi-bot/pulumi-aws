// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package applicationloadbalancing

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > **Note:** `aws_alb_listener` is known as `aws_lb_listener`. The functionality is identical.
// 
// Provides information about a Load Balancer Listener.
// 
// This data source can prove useful when a module accepts an LB Listener as an
// input variable and needs to know the LB it is attached to, or other
// information specific to the listener in question.
func LookupListener(ctx *pulumi.Context, args *GetListenerArgs) (*GetListenerResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["arn"] = args.Arn
		inputs["loadBalancerArn"] = args.LoadBalancerArn
		inputs["port"] = args.Port
	}
	outputs, err := ctx.Invoke("aws:applicationloadbalancing/getListener:getListener", inputs)
	if err != nil {
		return nil, err
	}
	return &GetListenerResult{
		Arn: outputs["arn"],
		CertificateArn: outputs["certificateArn"],
		DefaultActions: outputs["defaultActions"],
		LoadBalancerArn: outputs["loadBalancerArn"],
		Port: outputs["port"],
		Protocol: outputs["protocol"],
		SslPolicy: outputs["sslPolicy"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getListener.
type GetListenerArgs struct {
	// The arn of the listener. Required if `load_balancer_arn` and `port` is not set.
	Arn interface{}
	// The arn of the load balander. Required if `arn` is not set.
	LoadBalancerArn interface{}
	// The port of the listener. Required if `arn` is not set.
	Port interface{}
}

// A collection of values returned by getListener.
type GetListenerResult struct {
	Arn interface{}
	CertificateArn interface{}
	DefaultActions interface{}
	LoadBalancerArn interface{}
	Port interface{}
	Protocol interface{}
	SslPolicy interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
