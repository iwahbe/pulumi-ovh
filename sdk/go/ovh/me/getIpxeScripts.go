// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package me

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve a list of the names of the account's IPXE Scripts.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/Me"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Me.GetIpxeScripts(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetIpxeScripts(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetIpxeScriptsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIpxeScriptsResult
	err := ctx.Invoke("ovh:Me/getIpxeScripts:getIpxeScripts", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getIpxeScripts.
type GetIpxeScriptsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of the names of all the IPXE Scripts.
	Results []string `pulumi:"results"`
}

func GetIpxeScriptsOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetIpxeScriptsResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetIpxeScriptsResult, error) {
		r, err := GetIpxeScripts(ctx, opts...)
		var s GetIpxeScriptsResult
		if r != nil {
			s = *r
		}
		return s, err
	}).(GetIpxeScriptsResultOutput)
}

// A collection of values returned by getIpxeScripts.
type GetIpxeScriptsResultOutput struct{ *pulumi.OutputState }

func (GetIpxeScriptsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpxeScriptsResult)(nil)).Elem()
}

func (o GetIpxeScriptsResultOutput) ToGetIpxeScriptsResultOutput() GetIpxeScriptsResultOutput {
	return o
}

func (o GetIpxeScriptsResultOutput) ToGetIpxeScriptsResultOutputWithContext(ctx context.Context) GetIpxeScriptsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetIpxeScriptsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpxeScriptsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of the names of all the IPXE Scripts.
func (o GetIpxeScriptsResultOutput) Results() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpxeScriptsResult) []string { return v.Results }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIpxeScriptsResultOutput{})
}
