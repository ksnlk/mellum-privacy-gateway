//
// See the AWS API reference guide for Amazon Elastic Compute Cloud's
// API operation AllocateHosts for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AllocateHosts
func (c *EC2) AllocateHosts(input *AllocateHostsInput) (*AllocateHostsOutput, error) {
	req, out := c.AllocateHostsRequest(input)
	return out, req.Send()
}

// AllocateHostsWithContext is the same as AllocateHosts with the addition of
// the ability to pass a context and additional request options.
//
// See AllocateHosts for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EC2) AllocateHostsWithContext(ctx aws.Context, input *AllocateHostsInput, opts ...request.Option) (*AllocateHostsOutput, error) {
	req, out := c.AllocateHostsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAssignIpv6Addresses = "AssignIpv6Addresses"

// AssignIpv6AddressesRequest generates a "aws/request.Request" representing the
// client's request for the AssignIpv6Addresses operation. The "output" return
// value will be populated with the request's response once the request completes
// successfuly.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See AssignIpv6Addresses for more information on using the AssignIpv6Addresses
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the AssignIpv6AddressesRequest method.
//    req, resp := client.AssignIpv6AddressesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssignIpv6Addresses
//
// See the AWS API reference guide for Amazon Elastic Compute Cloud's
// API operation AssignIpv6Addresses for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssignIpv6Addresses
func (c *EC2) AssignIpv6Addresses(input *AssignIpv6AddressesInput) (*AssignIpv6AddressesOutput, error) {
	req, out := c.AssignIpv6AddressesRequest(input)
	return out, req.Send()
}

// AssignIpv6AddressesWithContext is the same as AssignIpv6Addresses with the addition of
// the ability to pass a context and additional request options.
//
// See AssignIpv6Addresses for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EC2) AssignIpv6AddressesWithContext(ctx aws.Context, input *AssignIpv6AddressesInput, opts ...request.Option) (*AssignIpv6AddressesOutput, error) {
	req, out := c.AssignIpv6AddressesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAssignPrivateIpAddresses = "AssignPrivateIpAddresses"

// AssignPrivateIpAddressesRequest generates a "aws/request.Request" representing the
// client's request for the AssignPrivateIpAddresses operation. The "output" return
// value will be populated with the request's response once the request completes
// successfuly.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See AssignPrivateIpAddresses for more information on using the AssignPrivateIpAddresses
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the AssignPrivateIpAddressesRequest method.
//    req, resp := client.AssignPrivateIpAddressesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssignPrivateIpAddresses
//
// See the AWS API reference guide for Amazon Elastic Compute Cloud's
// API operation AssociateSubnetCidrBlock for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssociateSubnetCidrBlock
func (c *EC2) AssociateSubnetCidrBlock(input *AssociateSubnetCidrBlockInput) (*AssociateSubnetCidrBlockOutput, error) {
	req, out := c.AssociateSubnetCidrBlockRequest(input)
	return out, req.Send()
}

// AssociateSubnetCidrBlockWithContext is the same as AssociateSubnetCidrBlock with the addition of
// the ability to pass a context and additional request options.
//
// See AssociateSubnetCidrBlock for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EC2) AssociateSubnetCidrBlockWithContext(ctx aws.Context, input *AssociateSubnetCidrBlockInput, opts ...request.Option) (*AssociateSubnetCidrBlockOutput, error) {
	req, out := c.AssociateSubnetCidrBlockRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAssociateVpcCidrBlock = "AssociateVpcCidrBlock"

// AssociateVpcCidrBlockRequest generates a "aws/request.Request" representing the
// client's request for the AssociateVpcCidrBlock operation. The "output" return
// value will be populated with the request's response once the request completes
// successfuly.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See AssociateVpcCidrBlock for more information on using the AssociateVpcCidrBlock
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the AssociateVpcCidrBlockRequest method.
//    req, resp := client.AssociateVpcCidrBlockRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssociateVpcCidrBlock
//
// See the AWS API reference guide for Amazon Elastic Compute Cloud's
// API operation GetLaunchTemplateData for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/GetLaunchTemplateData
func (c *EC2) GetLaunchTemplateData(input *GetLaunchTemplateDataInput) (*GetLaunchTemplateDataOutput, error) {
	req, out := c.GetLaunchTemplateDataRequest(input)
	return out, req.Send()
}

// GetLaunchTemplateDataWithContext is the same as GetLaunchTemplateData with the addition of
// the ability to pass a context and additional request options.
//
// See GetLaunchTemplateData for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EC2) GetLaunchTemplateDataWithContext(ctx aws.Context, input *GetLaunchTemplateDataInput, opts ...request.Option) (*GetLaunchTemplateDataOutput, error) {
	req, out := c.GetLaunchTemplateDataRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetPasswordData = "GetPasswordData"

// GetPasswordDataRequest generates a "aws/request.Request" representing the
// client's request for the GetPasswordData operation. The "output" return
// value will be populated with the request's response once the request completes
// successfuly.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetPasswordData for more information on using the GetPasswordData
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetPasswordDataRequest method.
//    req, resp := client.GetPasswordDataRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/GetPasswordData
func (c *EC2) GetPasswordDataRequest(input *GetPasswordDataInput) (req *request.Request, output *GetPasswordDataOutput) {
	op := &request.Operation{
		Name:       opGetPasswordData,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetPasswordDataInput{}
	}

	output = &GetPasswordDataOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetPasswordData API operation for Amazon Elastic Compute Cloud.
//
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// One or more filters. Filter names and values are case-sensitive.
	//
	//    * allocation-id - [EC2-VPC] The allocation ID for the address.
	//
	//    * association-id - [EC2-VPC] The association ID for the address.
	//
	//    * domain - Indicates whether the address is for use in EC2-Classic (standard)
	//    or in a VPC (vpc).
	//
	//    * instance-id - The ID of the instance the address is associated with,
	//    if any.
	//
	//    * network-interface-id - [EC2-VPC] The ID of the network interface that
	//    the address is associated with, if any.
	//
	//    * network-interface-owner-id - The AWS account ID of the owner.
	//
	//    * private-ip-address - [EC2-VPC] The private IP address associated with
	//    the Elastic IP address.
	//
	//    * public-ip - The Elastic IP address.
	//
	//    * tag:key=value - The key/value combination of a tag assigned to the resource.
	//    Specify the key of the tag in the filter name and the value of the tag
	//    in the filter value. For example, for the tag Purpose=X, specify tag:Purpose
	//    for the filter name and X for the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. This filter is
	//    independent of the tag-value filter. For example, if you use both the
	//    filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
	//    assigned both the tag key Purpose (regardless of what the tag's value
	//    is), and the tag value X (regardless of the tag's key). If you want to
	//    list only resources where Purpose is X, see the tag:key=value filter.
	Filters []*Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// [EC2-Classic] One or more Elastic IP addresses.
	//
	// Default: Describes all your Elastic IP addresses.
	PublicIps []*string `locationName:"PublicIp" locationNameList:"PublicIp" type:"list"`
}

// String returns the string representation
func (s DescribeAddressesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeAddressesInput) GoString() string {
	return s.String()
}

// SetAllocationIds sets the AllocationIds field's value.
func (s *DescribeAddressesInput) SetAllocationIds(v []*string) *DescribeAddressesInput {
	s.AllocationIds = v
	return s
}

// SetDryRun sets the DryRun field's value.
}

// SetBundleTasks sets the BundleTasks field's value.
func (s *DescribeBundleTasksOutput) SetBundleTasks(v []*BundleTask) *DescribeBundleTasksOutput {
	s.BundleTasks = v
	return s
}

// Contains the parameters for DescribeClassicLinkInstances.
type DescribeClassicLinkInstancesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// One or more filters.
	//
	//    * group-id - The ID of a VPC security group that's associated with the
	//    instance.
	//
	//    * instance-id - The ID of the instance.
	//
	//    * tag:key=value - The key/value combination of a tag assigned to the resource.
	//
	//    * tag-key - The key of a tag assigned to the resource. This filter is
	//    independent of the tag-value filter. For example, if you use both the
	//    filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
	//    assigned both the tag key Purpose (regardless of what the tag's value
	//    is), and the tag value X (regardless of what the tag's key is). If you
	//    want to list only resources where Purpose is X, see the tag:key=value
	//    filter.
	//
	//    * tag-value - The value of a tag assigned to the resource. This filter
	//    is independent of the tag-key filter.
	//
	//    * vpc-id - The ID of the VPC that the instance is linked to.
	Filters []*Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// One or more instance IDs. Must be instances linked to a VPC through ClassicLink.
	InstanceIds []*string `locationName:"InstanceId" locationNameList:"InstanceId" type:"list"`

	// The maximum number of results to return for the request in a single page.
	// The remaining results of the initial request can be seen by sending another
	// request with the returned NextToken value. This value can be between 5 and
	// 1000; if MaxResults is given a value larger than 1000, only 1000 results
	// are returned. You cannot specify this parameter and the instance IDs parameter
	// in the same request.
	//
	// Constraint: If the value is greater than 1000, we return only 1000 items.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The token to retrieve the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

	// Default: Describes all your customer gateways.
	CustomerGatewayIds []*string `locationName:"CustomerGatewayId" locationNameList:"CustomerGatewayId" type:"list"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// One or more filters.
	//
	//    * bgp-asn - The customer gateway's Border Gateway Protocol (BGP) Autonomous
	//    System Number (ASN).
	//
	//    * customer-gateway-id - The ID of the customer gateway.
	//
	//    * ip-address - The IP address of the customer gateway's Internet-routable
	//    external interface.
	//
	//    * state - The state of the customer gateway (pending | available | deleting
	//    | deleted).
	//
	//    * type - The type of customer gateway. Currently, the only supported type
	//    is ipsec.1.
	//
	//    * tag:key=value - The key/value combination of a tag assigned to the resource.
	//    Specify the key of the tag in the filter name and the value of the tag
	//    in the filter value. For example, for the tag Purpose=X, specify tag:Purpose
	//    for the filter name and X for the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. This filter is
	//    independent of the tag-value filter. For example, if you use both the
	//    filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
	//    assigned both the tag key Purpose (regardless of what the tag's value
	//    is), and the tag value X (regardless of what the tag's key is). If you
	//    want to list only resources where Purpose is X, see the tag:key=value
	//    filter.
	//
	//    * tag-value - The value of a tag assigned to the resource. This filter
	//    is independent of the tag-key filter.
	Filters []*Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`
}

// String returns the string representation
func (s DescribeCustomerGatewaysInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCustomerGatewaysInput) GoString() string {
	return s.String()
}

// SetCustomerGatewayIds sets the CustomerGatewayIds field's value.
func (s *DescribeCustomerGatewaysInput) SetCustomerGatewayIds(v []*string) *DescribeCustomerGatewaysInput {
	s.CustomerGatewayIds = v
	return s
}

// SetDryRun sets the DryRun field's value.
func (s *DescribeCustomerGatewaysInput) SetDryRun(v bool) *DescribeCustomerGatewaysInput {
}

// Contains the parameters for DescribeDhcpOptions.
type DescribeDhcpOptionsInput struct {
	_ struct{} `type:"structure"`

	// The IDs of one or more DHCP options sets.
	//
	// Default: Describes all your DHCP options sets.
	DhcpOptionsIds []*string `locationName:"DhcpOptionsId" locationNameList:"DhcpOptionsId" type:"list"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// One or more filters.
	//
	//    * dhcp-options-id - The ID of a set of DHCP options.
	//
	//    * key - The key for one of the options (for example, domain-name).
	//
	//    * value - The value for one of the options.
	//
	//    * tag:key=value - The key/value combination of a tag assigned to the resource.
	//    Specify the key of the tag in the filter name and the value of the tag
	//    in the filter value. For example, for the tag Purpose=X, specify tag:Purpose
	//    for the filter name and X for the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. This filter is
	//    independent of the tag-value filter. For example, if you use both the
	//    filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
	//    assigned both the tag key Purpose (regardless of what the tag's value
	//    is), and the tag value X (regardless of what the tag's key is). If you
	//    want to list only resources where Purpose is X, see the tag:key=value
	//    filter.
	//
	//    * tag-value - The value of a tag assigned to the resource. This filter
	//    is independent of the tag-key filter.
	Filters []*Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`
}

// String returns the string representation
func (s DescribeDhcpOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDhcpOptionsInput) GoString() string {
	return s.String()
}

// SetDhcpOptionsIds sets the DhcpOptionsIds field's value.
func (s *DescribeDhcpOptionsInput) SetDhcpOptionsIds(v []*string) *DescribeDhcpOptionsInput {
	s.DhcpOptionsIds = v
	return s
}

// SetDryRun sets the DryRun field's value.
func (s *DescribeDhcpOptionsInput) SetDryRun(v bool) *DescribeDhcpOptionsInput {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters.
	//
	//    * create-time - The creation time of the AFI.
	//
	//    * fpga-image-id - The FPGA image identifier (AFI ID).
	//
	//    * fpga-image-global-id - The global FPGA image identifier (AGFI ID).
	//
	//    * name - The name of the AFI.
	//
	//    * owner-id - The AWS account ID of the AFI owner.
	//
	//    * product-code - The product code.
	//
	//    * shell-version - The version of the AWS Shell that was used to create
	//    the bitstream.
	//
	//    * state - The state of the AFI (pending | failed | available | unavailable).
	//
	//    * tag:key=value - The key/value combination of a tag assigned to the resource.
	//    Specify the key of the tag in the filter name and the value of the tag
	//    in the filter value. For example, for the tag Purpose=X, specify tag:Purpose
	//    for the filter name and X for the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. This filter is
	//    independent of the tag-value filter. For example, if you use both the
	//    filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
	//    assigned both the tag key Purpose (regardless of what the tag's value
	//    is), and the tag value X (regardless of what the tag's key is). If you
	//    want to list only resources where Purpose is X, see the tag:key=value
	//    filter.
	//
	//    * tag-value - The value of a tag assigned to the resource. This filter
	//    is independent of the tag-key filter.
	//
	//    * update-time - The time of the most recent update.
	Filters []*Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// One or more AFI IDs.
	FpgaImageIds []*string `locationName:"FpgaImageId" locationNameList:"item" type:"list"`

	// The maximum number of results to return in a single call.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token to retrieve the next page of results.
	NextToken *string `min:"1" type:"string"`

	// Filters the AFI by owner. Specify an AWS account ID, self (owner is the sender
	// of the request), or an AWS owner alias (valid values are amazon | aws-marketplace).
	Owners []*string `locationName:"Owner" locationNameList:"Owner" type:"list"`
}

// String returns the string representation
func (s DescribeFpgaImagesInput) String() string {
	return awsutil.Prettify(s)
	//
	//    * owner-id - The AWS account ID of the image owner.
	//
	//    * platform - The platform. To only list Windows-based AMIs, use windows.
	//
	//    * product-code - The product code.
	//
	//    * product-code.type - The type of the product code (devpay | marketplace).
	//
	//    * ramdisk-id - The RAM disk ID.
	//
	//    * root-device-name - The device name of the root device volume (for example,
	//    /dev/sda1).
	//
	//    * root-device-type - The type of the root device volume (ebs | instance-store).
	//
	//    * state - The state of the image (available | pending | failed).
	//
	//    * state-reason-code - The reason code for the state change.
	//
	//    * state-reason-message - The message for the state change.
	//
	//    * sriov-net-support - A value of simple indicates that enhanced networking
	//    with the Intel 82599 VF interface is enabled.
	//
	//    * tag:key=value - The key/value combination of a tag assigned to the resource.
	//    Specify the key of the tag in the filter name and the value of the tag
	//    in the filter value. For example, for the tag Purpose=X, specify tag:Purpose
	//    for the filter name and X for the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. This filter is
	//    independent of the tag-value filter. For example, if you use both the
	//    filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
	//    assigned both the tag key Purpose (regardless of what the tag's value
	//    is), and the tag value X (regardless of what the tag's key is). If you
	//    want to list only resources where Purpose is X, see the tag:key=value
	//    filter.
	//
	//    * tag-value - The value of a tag assigned to the resource. This filter
	//    is independent of the tag-key filter.
	//
	//    * virtualization-type - The virtualization type (paravirtual | hvm).
	Filters []*Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// One or more image IDs.
	//
	// Default: Describes all images available to you.
	ImageIds []*string `locationName:"ImageId" locationNameList:"ImageId" type:"list"`

	// Filters the images by the owner. Specify an AWS account ID, self (owner is
	// the sender of the request), or an AWS owner alias (valid values are amazon
	// | aws-marketplace | microsoft). Omitting this option returns all images for
	// which you have launch permissions, regardless of ownership.
	Owners []*string `locationName:"Owner" locationNameList:"Owner" type:"list"`
}

// String returns the string representation
func (s DescribeImagesInput) String() string {
	return awsutil.Prettify(s)
}

	//    ID is created any time you launch an instance. A reservation ID has a
	//    one-to-one relationship with an instance launch request, but can be associated
	//    with more than one instance if you launch multiple instances using the
	//    same launch request. For example, if you launch one instance, you get
	//    one reservation ID. If you launch ten instances using the same launch
	//    request, you also get one reservation ID.
	//
	//    * root-device-name - The device name of the root device volume (for example,
	//    /dev/sda1).
	//
	//    * root-device-type - The type of the root device volume (ebs | instance-store).
	//
	//    * source-dest-check - Indicates whether the instance performs source/destination
	//    checking. A value of true means that checking is enabled, and false means
	//    that checking is disabled. The value must be false for the instance to
	//    perform network address translation (NAT) in your VPC.
	//
	//    * spot-instance-request-id - The ID of the Spot Instance request.
	//
	//    * state-reason-code - The reason code for the state change.
	//
	//    * state-reason-message - A message that describes the state change.
	//
	//    * subnet-id - The ID of the subnet for the instance.
	//
	//    * tag:key=value - The key/value combination of a tag assigned to the resource.
	//    Specify the key of the tag in the filter name and the value of the tag
	//    in the filter value. For example, for the tag Purpose=X, specify tag:Purpose
	//    for the filter name and X for the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. This filter is
	//    independent of the tag-value filter. For example, if you use both the
	//    filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
	//    assigned both the tag key Purpose (regardless of what the tag's value
	//    is), and the tag value X (regardless of the tag's key). If you want to
	//    list only resources where Purpose is X, see the tag:key=value filter.
	//
	//    * tag-value - The value of a tag assigned to the resource. This filter
	//    is independent of the tag-key filter.
	//
	//    * tenancy - The tenancy of an instance (dedicated | default | host).
	//
	//    * virtualization-type - The virtualization type of the instance (paravirtual
	//    | hvm).
	//
	//    * vpc-id - The ID of the VPC that the instance is running in.
	Filters []*Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// One or more instance IDs.
	//
	// Default: Describes all your instances.
	InstanceIds []*string `locationName:"InstanceId" locationNameList:"InstanceId" type:"list"`

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value. This
	// value can be between 5 and 1000. You cannot specify this parameter and the
	// instance IDs parameter or tag filters in the same call.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The token to request the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`
// SetReservations sets the Reservations field's value.
func (s *DescribeInstancesOutput) SetReservations(v []*Reservation) *DescribeInstancesOutput {
	s.Reservations = v
	return s
}

// Contains the parameters for DescribeInternetGateways.
type DescribeInternetGatewaysInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// One or more filters.
	//
	//    * attachment.state - The current state of the attachment between the gateway
	//    and the VPC (available). Present only if a VPC is attached.
	//
	//    * attachment.vpc-id - The ID of an attached VPC.
	//
	//    * internet-gateway-id - The ID of the Internet gateway.
	//
	//    * tag:key=value - The key/value combination of a tag assigned to the resource.
	//    Specify the key of the tag in the filter name and the value of the tag
	//    in the filter value. For example, for the tag Purpose=X, specify tag:Purpose
	//    for the filter name and X for the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. This filter is
	//    independent of the tag-value filter. For example, if you use both the
	//    filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
	//    assigned both the tag key Purpose (regardless of what the tag's value
	//    is), and the tag value X (regardless of what the tag's key is). If you
	//    want to list only resources where Purpose is X, see the tag:key=value
	//    filter.
	//
	//    * tag-value - The value of a tag assigned to the resource. This filter
	//    is independent of the tag-key filter.
	Filters []*Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// One or more Internet gateway IDs.
	//
	// Default: Describes all your Internet gateways.
	InternetGatewayIds []*string `locationName:"internetGatewayId" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeInternetGatewaysInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInternetGatewaysInput) GoString() string {
	return s.String()
}

// SetDryRun sets the DryRun field's value.
func (s *DescribeInternetGatewaysInput) SetDryRun(v bool) *DescribeInternetGatewaysInput {
	s.DryRun = &v
	s.LaunchTemplateVersions = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeLaunchTemplateVersionsOutput) SetNextToken(v string) *DescribeLaunchTemplateVersionsOutput {
	s.NextToken = &v
	return s
}

type DescribeLaunchTemplatesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters.
	//
	//    * create-time - The time the launch template was created.
	//
	//    * launch-template-name - The name of the launch template.
	//
	//    * tag:key=value - The key/value combination of a tag assigned to the resource.
	//    Specify the key of the tag in the filter name and the value of the tag
	//    in the filter value. For example, for the tag Purpose=X, specify tag:Purpose
	//    for the filter name and X for the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. This filter is
	//    independent of the tag-value filter. For example, if you use both the
	//    filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
	//    assigned both the tag key Purpose (regardless of what the tag's value
	//    is), and the tag value X (regardless of the tag's key). If you want to
	//    list only resources where Purpose is X, see the tag:key=value filter.
	Filters []*Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// One or more launch template IDs.
	LaunchTemplateIds []*string `locationName:"LaunchTemplateId" locationNameList:"item" type:"list"`

	// One or more launch template names.
	LaunchTemplateNames []*string `locationName:"LaunchTemplateName" locationNameList:"item" type:"list"`

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value. This
	// value can be between 5 and 1000.
	MaxResults *int64 `type:"integer"`

	// The token to request the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeLaunchTemplatesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeLaunchTemplatesInput) GoString() string {
	return s.String()
// SetMovingAddressStatuses sets the MovingAddressStatuses field's value.
func (s *DescribeMovingAddressesOutput) SetMovingAddressStatuses(v []*MovingAddressStatus) *DescribeMovingAddressesOutput {
	s.MovingAddressStatuses = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeMovingAddressesOutput) SetNextToken(v string) *DescribeMovingAddressesOutput {
	s.NextToken = &v
	return s
}

// Contains the parameters for DescribeNatGateways.
type DescribeNatGatewaysInput struct {
	_ struct{} `type:"structure"`

	// One or more filters.
	//
	//    * nat-gateway-id - The ID of the NAT gateway.
	//
	//    * state - The state of the NAT gateway (pending | failed | available |
	//    deleting | deleted).
	//
	//    * subnet-id - The ID of the subnet in which the NAT gateway resides.
	//
	//    * tag:key=value - The key/value combination of a tag assigned to the resource.
	//    Specify the key of the tag in the filter name and the value of the tag
	//    in the filter value. For example, for the tag Purpose=X, specify tag:Purpose
	//    for the filter name and X for the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. This filter is
	//    independent of the tag-value filter. For example, if you use both the
	//    filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
	//    assigned both the tag key Purpose (regardless of what the tag's value
	//    is), and the tag value X (regardless of what the tag's key is). If you
	//    want to list only resources where Purpose is X, see the tag:key=value
	//    filter.
	//
	//    * tag-value - The value of a tag assigned to the resource. This filter
	//    is independent of the tag-key filter.
	//
	//    * vpc-id - The ID of the VPC in which the NAT gateway resides.
	Filter []*Filter `locationNameList:"Filter" type:"list"`

	// The maximum number of items to return for this request. The request returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	//
	// Constraint: If the value specified is greater than 1000, we return only 1000
	// items.
	MaxResults *int64 `type:"integer"`

	// One or more NAT gateway IDs.
	NatGatewayIds []*string `locationName:"NatGatewayId" locationNameList:"item" type:"list"`

	// The token to retrieve the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeNatGatewaysInput) String() string {
	//
	//    * entry.egress - Indicates whether the entry applies to egress traffic.
	//
	//    * entry.icmp.code - The ICMP code specified in the entry, if any.
	//
	//    * entry.icmp.type - The ICMP type specified in the entry, if any.
	//
	//    * entry.ipv6-cidr - The IPv6 CIDR range specified in the entry.
	//
	//    * entry.port-range.from - The start of the port range specified in the
	//    entry.
	//
	//    * entry.port-range.to - The end of the port range specified in the entry.
	//
	//
	//    * entry.protocol - The protocol specified in the entry (tcp | udp | icmp
	//    or a protocol number).
	//
	//    * entry.rule-action - Allows or denies the matching traffic (allow | deny).
	//
	//    * entry.rule-number - The number of an entry (in other words, rule) in
	//    the ACL's set of entries.
	//
	//    * network-acl-id - The ID of the network ACL.
	//
	//    * tag:key=value - The key/value combination of a tag assigned to the resource.
	//    Specify the key of the tag in the filter name and the value of the tag
	//    in the filter value. For example, for the tag Purpose=X, specify tag:Purpose
	//    for the filter name and X for the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. This filter is
	//    independent of the tag-value filter. For example, if you use both the
	//    filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
	//    assigned both the tag key Purpose (regardless of what the tag's value
	//    is), and the tag value X (regardless of what the tag's key is). If you
	//    want to list only resources where Purpose is X, see the tag:key=value
	//    filter.
	//
	//    * tag-value - The value of a tag assigned to the resource. This filter
	//    is independent of the tag-key filter.
	//
	//    * vpc-id - The ID of the VPC for the network ACL.
	Filters []*Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// One or more network ACL IDs.
	//
	// Default: Describes all your network ACLs.
	NetworkAclIds []*string `locationName:"NetworkAclId" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeNetworkAclsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeNetworkAclsInput) GoString() string {
	return s.String()
}

// SetDryRun sets the DryRun field's value.
	//
	//    * private-ip-address - The private IPv4 address or addresses of the network
	//    interface.
	//
	//    * private-dns-name - The private DNS name of the network interface (IPv4).
	//
	//    * requester-id - The ID of the entity that launched the instance on your
	//    behalf (for example, AWS Management Console, Auto Scaling, and so on).
	//
	//    * requester-managed - Indicates whether the network interface is being
	//    managed by an AWS service (for example, AWS Management Console, Auto Scaling,
	//    and so on).
	//
	//    * source-desk-check - Indicates whether the network interface performs
	//    source/destination checking. A value of true means checking is enabled,
	//    and false means checking is disabled. The value must be false for the
	//    network interface to perform network address translation (NAT) in your
	//    VPC.
	//
	//    * status - The status of the network interface. If the network interface
	//    is not attached to an instance, the status is available; if a network
	//    interface is attached to an instance the status is in-use.
	//
	//    * subnet-id - The ID of the subnet for the network interface.
	//
	//    * tag:key=value - The key/value combination of a tag assigned to the resource.
	//    Specify the key of the tag in the filter name and the value of the tag
	//    in the filter value. For example, for the tag Purpose=X, specify tag:Purpose
	//    for the filter name and X for the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. This filter is
	//    independent of the tag-value filter. For example, if you use both the
	//    filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
	//    assigned both the tag key Purpose (regardless of what the tag's value
	//    is), and the tag value X (regardless of what the tag's key is). If you
	//    want to list only resources where Purpose is X, see the tag:key=value
	//    filter.
	//
	//    * tag-value - The value of a tag assigned to the resource. This filter
	//    is independent of the tag-key filter.
	//
	//    * vpc-id - The ID of the VPC for the network interface.
	Filters []*Filter `locationName:"filter" locationNameList:"Filter" type:"list"`

	// One or more network interface IDs.
	//
	// Default: Describes all your network interfaces.
	NetworkInterfaceIds []*string `locationName:"NetworkInterfaceId" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeNetworkInterfacesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeNetworkInterfacesInput) GoString() string {
	return s.String()
}

// SetDryRun sets the DryRun field's value.
	//    * fixed-price - The purchase price of the Reserved Instance (for example,
	//    9800.0).
	//
	//    * instance-type - The instance type that is covered by the reservation.
	//
	//    * scope - The scope of the Reserved Instance (Region or Availability Zone).
	//
	//    * product-description - The Reserved Instance product platform description.
	//    Instances that include (Amazon VPC) in the product platform description
	//    will only be displayed to EC2-Classic account holders and are for use
	//    with Amazon VPC (Linux/UNIX | Linux/UNIX (Amazon VPC) | SUSE Linux | SUSE
	//    Linux (Amazon VPC) | Red Hat Enterprise Linux | Red Hat Enterprise Linux
	//    (Amazon VPC) | Windows | Windows (Amazon VPC) | Windows with SQL Server
	//    Standard | Windows with SQL Server Standard (Amazon VPC) | Windows with
	//    SQL Server Web | Windows with SQL Server Web (Amazon VPC) | Windows with
	//    SQL Server Enterprise | Windows with SQL Server Enterprise (Amazon VPC)).
	//
	//    * reserved-instances-id - The ID of the Reserved Instance.
	//
	//    * start - The time at which the Reserved Instance purchase request was
	//    placed (for example, 2014-08-07T11:54:42.000Z).
	//
	//    * state - The state of the Reserved Instance (payment-pending | active
	//    | payment-failed | retired).
	//
	//    * tag:key=value - The key/value combination of a tag assigned to the resource.
	//    Specify the key of the tag in the filter name and the value of the tag
	//    in the filter value. For example, for the tag Purpose=X, specify tag:Purpose
	//    for the filter name and X for the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. This filter is
	//    independent of the tag-value filter. For example, if you use both the
	//    filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
	//    assigned both the tag key Purpose (regardless of what the tag's value
	//    is), and the tag value X (regardless of what the tag's key is). If you
	//    want to list only resources where Purpose is X, see the tag:key=value
	//    filter.
	//
	//    * tag-value - The value of a tag assigned to the resource. This filter
	//    is independent of the tag-key filter.
	//
	//    * usage-price - The usage price of the Reserved Instance, per hour (for
	//    example, 0.84).
	Filters []*Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// Describes whether the Reserved Instance is Standard or Convertible.
	OfferingClass *string `type:"string" enum:"OfferingClassType"`

	// The Reserved Instance offering type. If you are using tools that predate
	// the 2011-11-01 API version, you only have access to the Medium Utilization
	// Reserved Instance offering type.
	OfferingType *string `locationName:"offeringType" type:"string" enum:"OfferingTypeValues"`

	// One or more Reserved Instance IDs.
	//
	// Default: Describes all your Reserved Instances, or only those otherwise specified.
	ReservedInstancesIds []*string `locationName:"ReservedInstancesId" locationNameList:"ReservedInstancesId" type:"list"`
}

// String returns the string representation
func (s DescribeReservedInstancesInput) String() string {
	//
	//    * route.egress-only-internet-gateway-id - The ID of an egress-only Internet
	//    gateway specified in a route in the route table.
	//
	//    * route.gateway-id - The ID of a gateway specified in a route in the table.
	//
	//    * route.instance-id - The ID of an instance specified in a route in the
	//    table.
	//
	//    * route.nat-gateway-id - The ID of a NAT gateway.
	//
	//    * route.origin - Describes how the route was created. CreateRouteTable
	//    indicates that the route was automatically created when the route table
	//    was created; CreateRoute indicates that the route was manually added to
	//    the route table; EnableVgwRoutePropagation indicates that the route was
	//    propagated by route propagation.
	//
	//    * route.state - The state of a route in the route table (active | blackhole).
	//    The blackhole state indicates that the route's target isn't available
	//    (for example, the specified gateway isn't attached to the VPC, the specified
	//    NAT instance has been terminated, and so on).
	//
	//    * route.vpc-peering-connection-id - The ID of a VPC peering connection
	//    specified in a route in the table.
	//
	//    * tag:key=value - The key/value combination of a tag assigned to the resource.
	//    Specify the key of the tag in the filter name and the value of the tag
	//    in the filter value. For example, for the tag Purpose=X, specify tag:Purpose
	//    for the filter name and X for the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. This filter is
	//    independent of the tag-value filter. For example, if you use both the
	//    filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
	//    assigned both the tag key Purpose (regardless of what the tag's value
	//    is), and the tag value X (regardless of what the tag's key is). If you
	//    want to list only resources where Purpose is X, see the tag:key=value
	//    filter.
	//
	//    * tag-value - The value of a tag assigned to the resource. This filter
	//    is independent of the tag-key filter.
	//
	//    * vpc-id - The ID of the VPC for the route table.
	Filters []*Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// One or more route table IDs.
	//
	// Default: Describes all your route tables.
	RouteTableIds []*string `locationName:"RouteTableId" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeRouteTablesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeRouteTablesInput) GoString() string {
	return s.String()
}

// SetDryRun sets the DryRun field's value.
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// One or more filters.
	//
	//    * description - A description of the snapshot.
	//
	//    * owner-alias - Value from an Amazon-maintained list (amazon | aws-marketplace
	//    | microsoft) of snapshot owners. Not to be confused with the user-configured
	//    AWS account alias, which is set from the IAM console.
	//
	//    * owner-id - The ID of the AWS account that owns the snapshot.
	//
	//    * progress - The progress of the snapshot, as a percentage (for example,
	//    80%).
	//
	//    * snapshot-id - The snapshot ID.
	//
	//    * start-time - The time stamp when the snapshot was initiated.
	//
	//    * status - The status of the snapshot (pending | completed | error).
	//
	//    * tag:key=value - The key/value combination of a tag assigned to the resource.
	//    Specify the key of the tag in the filter name and the value of the tag
	//    in the filter value. For example, for the tag Purpose=X, specify tag:Purpose
	//    for the filter name and X for the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. This filter is
	//    independent of the tag-value filter. For example, if you use both the
	//    filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
	//    assigned both the tag key Purpose (regardless of what the tag's value
	//    is), and the tag value X (regardless of what the tag's key is). If you
	//    want to list only resources where Purpose is X, see the tag:key=value
	//    filter.
	//
	//    * tag-value - The value of a tag assigned to the resource. This filter
	//    is independent of the tag-key filter.
	//
	//    * volume-id - The ID of the volume the snapshot is for.
	//
	//    * volume-size - The size of the volume, in GiB.
	Filters []*Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of snapshot results returned by DescribeSnapshots in paginated
	// output. When this parameter is used, DescribeSnapshots only returns MaxResults
	// results in a single page along with a NextToken response element. The remaining
	// results of the initial request can be seen by sending another DescribeSnapshots
	// request with the returned NextToken value. This value can be between 5 and
	// 1000; if MaxResults is given a value larger than 1000, only 1000 results
	// are returned. If this parameter is not used, then DescribeSnapshots returns
	// all results. You cannot specify this parameter and the snapshot IDs parameter
	// in the same request.
	MaxResults *int64 `type:"integer"`

	// The NextToken value returned from a previous paginated DescribeSnapshots
	// request where MaxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the NextToken value. This value is null when there are no more results
	//    * network-interface.private-ip-address - The primary private IP address
	//    of the network interface.
	//
	//    * network-interface.subnet-id - The ID of the subnet for the instance.
	//
	//    * product-description - The product description associated with the instance
	//    (Linux/UNIX | Windows).
	//
	//    * spot-instance-request-id - The Spot Instance request ID.
	//
	//    * spot-price - The maximum hourly price for any Spot Instance launched
	//    to fulfill the request.
	//
	//    * state - The state of the Spot Instance request (open | active | closed
	//    | cancelled | failed). Spot request status information can help you track
	//    your Amazon EC2 Spot Instance requests. For more information, see Spot
	//    Request Status (http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-bid-status.html)
	//    in the Amazon EC2 User Guide for Linux Instances.
	//
	//    * status-code - The short code describing the most recent evaluation of
	//    your Spot Instance request.
	//
	//    * status-message - The message explaining the status of the Spot Instance
	//    request.
	//
	//    * tag:key=value - The key/value combination of a tag assigned to the resource.
	//    Specify the key of the tag in the filter name and the value of the tag
	//    in the filter value. For example, for the tag Purpose=X, specify tag:Purpose
	//    for the filter name and X for the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. This filter is
	//    independent of the tag-value filter. For example, if you use both the
	//    filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
	//    assigned both the tag key Purpose (regardless of what the tag's value
	//    is), and the tag value X (regardless of what the tag's key is). If you
	//    want to list only resources where Purpose is X, see the tag:key=value
	//    filter.
	//
	//    * tag-value - The value of a tag assigned to the resource. This filter
	//    is independent of the tag-key filter.
	//
	//    * type - The type of Spot Instance request (one-time | persistent).
	//
	//    * valid-from - The start date of the request.
	//
	//    * valid-until - The end date of the request.
	Filters []*Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// One or more Spot Instance request IDs.
	SpotInstanceRequestIds []*string `locationName:"SpotInstanceRequestId" locationNameList:"SpotInstanceRequestId" type:"list"`
}

// String returns the string representation
func (s DescribeSpotInstanceRequestsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSpotInstanceRequestsInput) GoString() string {
	return s.String()
}
	//    use availability-zone as the filter name.
	//
	//    * available-ip-address-count - The number of IPv4 addresses in the subnet
	//    that are available.
	//
	//    * cidrBlock - The IPv4 CIDR block of the subnet. The CIDR block you specify
	//    must exactly match the subnet's CIDR block for information to be returned
	//    for the subnet. You can also use cidr or cidr-block as the filter names.
	//
	//    * defaultForAz - Indicates whether this is the default subnet for the
	//    Availability Zone. You can also use default-for-az as the filter name.
	//
	//    * ipv6-cidr-block-association.ipv6-cidr-block - An IPv6 CIDR block associated
	//    with the subnet.
	//
	//    * ipv6-cidr-block-association.association-id - An association ID for an
	//    IPv6 CIDR block associated with the subnet.
	//
	//    * ipv6-cidr-block-association.state - The state of an IPv6 CIDR block
	//    associated with the subnet.
	//
	//    * state - The state of the subnet (pending | available).
	//
	//    * subnet-id - The ID of the subnet.
	//
	//    * tag:key=value - The key/value combination of a tag assigned to the resource.
	//    Specify the key of the tag in the filter name and the value of the tag
	//    in the filter value. For example, for the tag Purpose=X, specify tag:Purpose
	//    for the filter name and X for the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. This filter is
	//    independent of the tag-value filter. For example, if you use both the
	//    filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
	//    assigned both the tag key Purpose (regardless of what the tag's value
	//    is), and the tag value X (regardless of what the tag's key is). If you
	//    want to list only resources where Purpose is X, see the tag:key=value
	//    filter.
	//
	//    * tag-value - The value of a tag assigned to the resource. This filter
	//    is independent of the tag-key filter.
	//
	//    * vpc-id - The ID of the VPC for the subnet.
	Filters []*Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// One or more subnet IDs.
	//
	// Default: Describes all your subnets.
	SubnetIds []*string `locationName:"SubnetId" locationNameList:"SubnetId" type:"list"`
}

// String returns the string representation
func (s DescribeSubnetsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSubnetsInput) GoString() string {
	return s.String()
}

// SetDryRun sets the DryRun field's value.
	//
	//    * action.code - The action code for the event (for example, enable-volume-io).
	//
	//    * action.description - A description of the action.
	//
	//    * action.event-id - The event ID associated with the action.
	//
	//    * availability-zone - The Availability Zone of the instance.
	//
	//    * event.description - A description of the event.
	//
	//    * event.event-id - The event ID.
	//
	//    * event.event-type - The event type (for io-enabled: passed | failed;
	//    for io-performance: io-performance:degraded | io-performance:severely-degraded
	//    | io-performance:stalled).
	//
	//    * event.not-after - The latest end time for the event.
	//
	//    * event.not-before - The earliest start time for the event.
	//
	//    * volume-status.details-name - The cause for volume-status.status (io-enabled
	//    | io-performance).
	//
	//    * volume-status.details-status - The status of volume-status.details-name
	//    (for io-enabled: passed | failed; for io-performance: normal | degraded
	//    | severely-degraded | stalled).
	//
	//    * volume-status.status - The status of the volume (ok | impaired | warning
	//    | insufficient-data).
	Filters []*Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of volume results returned by DescribeVolumeStatus in
	// paginated output. When this parameter is used, the request only returns MaxResults
	// results in a single page along with a NextToken response element. The remaining
	// results of the initial request can be seen by sending another request with
	// the returned NextToken value. This value can be between 5 and 1000; if MaxResults
	// is given a value larger than 1000, only 1000 results are returned. If this
	// parameter is not used, then DescribeVolumeStatus returns all results. You
	// cannot specify this parameter and the volume IDs parameter in the same request.
	MaxResults *int64 `type:"integer"`

	// The NextToken value to include in a future DescribeVolumeStatus request.
	// When the results of the request exceed MaxResults, this value can be used
	// to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string `type:"string"`

	// One or more volume IDs.
	//
	// Default: Describes all your volumes.
	//
	//    * attachment.delete-on-termination - Whether the volume is deleted on
	//    instance termination.
	//
	//    * attachment.device - The device name specified in the block device mapping
	//    (for example, /dev/sda1).
	//
	//    * attachment.instance-id - The ID of the instance the volume is attached
	//    to.
	//
	//    * attachment.status - The attachment state (attaching | attached | detaching).
	//
	//    * availability-zone - The Availability Zone in which the volume was created.
	//
	//    * create-time - The time stamp when the volume was created.
	//
	//    * encrypted - The encryption status of the volume.
	//
	//    * size - The size of the volume, in GiB.
	//
	//    * snapshot-id - The snapshot from which the volume was created.
	//
	//    * status - The status of the volume (creating | available | in-use | deleting
	//    | deleted | error).
	//
	//    * tag:key=value - The key/value combination of a tag assigned to the resource.
	//    Specify the key of the tag in the filter name and the value of the tag
	//    in the filter value. For example, for the tag Purpose=X, specify tag:Purpose
	//    for the filter name and X for the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. This filter is
	//    independent of the tag-value filter. For example, if you use both the
	//    filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
	//    assigned both the tag key Purpose (regardless of what the tag's value
	//    is), and the tag value X (regardless of what the tag's key is). If you
	//    want to list only resources where Purpose is X, see the tag:key=value
	//    filter.
	//
	//    * tag-value - The value of a tag assigned to the resource. This filter
	//    is independent of the tag-key filter.
	//
	//    * volume-id - The volume ID.
	//
	//    * volume-type - The Amazon EBS volume type. This can be gp2 for General
	//    Purpose SSD, io1 for Provisioned IOPS SSD, st1 for Throughput Optimized
	//    HDD, sc1 for Cold HDD, or standard for Magnetic volumes.
	Filters []*Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of volume results returned by DescribeVolumes in paginated
	// output. When this parameter is used, DescribeVolumes only returns MaxResults
	// results in a single page along with a NextToken response element. The remaining
	// results of the initial request can be seen by sending another DescribeVolumes
	// request with the returned NextToken value. This value can be between 5 and
	// 500; if MaxResults is given a value larger than 500, only 500 results are
	// returned. If this parameter is not used, then DescribeVolumes returns all
	// results. You cannot specify this parameter and the volume IDs parameter in
	// the same request.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The NextToken value returned from a previous paginated DescribeVolumes request
	// where MaxResults was used and the results exceeded the value of that parameter.
	s.NextToken = &v
	return s
}

// SetVpcs sets the Vpcs field's value.
func (s *DescribeVpcClassicLinkDnsSupportOutput) SetVpcs(v []*ClassicLinkDnsSupport) *DescribeVpcClassicLinkDnsSupportOutput {
	s.Vpcs = v
	return s
}

// Contains the parameters for DescribeVpcClassicLink.
type DescribeVpcClassicLinkInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// One or more filters.
	//
	//    * is-classic-link-enabled - Whether the VPC is enabled for ClassicLink
	//    (true | false).
	//
	//    * tag:key=value - The key/value combination of a tag assigned to the resource.
	//    Specify the key of the tag in the filter name and the value of the tag
	//    in the filter value. For example, for the tag Purpose=X, specify tag:Purpose
	//    for the filter name and X for the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. This filter is
	//    independent of the tag-value filter. For example, if you use both the
	//    filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
	//    assigned both the tag key Purpose (regardless of what the tag's value
	//    is), and the tag value X (regardless of what the tag's key is). If you
	//    want to list only resources where Purpose is X, see the tag:key=value
	//    filter.
	//
	//    * tag-value - The value of a tag assigned to the resource. This filter
	//    is independent of the tag-key filter.
	Filters []*Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// One or more VPCs for which you want to describe the ClassicLink status.
	VpcIds []*string `locationName:"VpcId" locationNameList:"VpcId" type:"list"`
}

// String returns the string representation
func (s DescribeVpcClassicLinkInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVpcClassicLinkInput) GoString() string {
	return s.String()
}

// SetDryRun sets the DryRun field's value.
func (s *DescribeVpcClassicLinkInput) SetDryRun(v bool) *DescribeVpcClassicLinkInput {
	s.DryRun = &v
	return s
}
	// One or more filters.
	//
	//    * accepter-vpc-info.cidr-block - The IPv4 CIDR block of the accepter VPC.
	//
	//    * accepter-vpc-info.owner-id - The AWS account ID of the owner of the
	//    accepter VPC.
	//
	//    * accepter-vpc-info.vpc-id - The ID of the accepter VPC.
	//
	//    * expiration-time - The expiration date and time for the VPC peering connection.
	//
	//    * requester-vpc-info.cidr-block - The IPv4 CIDR block of the requester's
	//    VPC.
	//
	//    * requester-vpc-info.owner-id - The AWS account ID of the owner of the
	//    requester VPC.
	//
	//    * requester-vpc-info.vpc-id - The ID of the requester VPC.
	//
	//    * status-code - The status of the VPC peering connection (pending-acceptance
	//    | failed | expired | provisioning | active | deleting | deleted | rejected).
	//
	//    * status-message - A message that provides more information about the
	//    status of the VPC peering connection, if applicable.
	//
	//    * tag:key=value - The key/value combination of a tag assigned to the resource.
	//    Specify the key of the tag in the filter name and the value of the tag
	//    in the filter value. For example, for the tag Purpose=X, specify tag:Purpose
	//    for the filter name and X for the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. This filter is
	//    independent of the tag-value filter. For example, if you use both the
	//    filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
	//    assigned both the tag key Purpose (regardless of what the tag's value
	//    is), and the tag value X (regardless of what the tag's key is). If you
	//    want to list only resources where Purpose is X, see the tag:key=value
	//    filter.
	//
	//    * tag-value - The value of a tag assigned to the resource. This filter
	//    is independent of the tag-key filter.
	//
	//    * vpc-peering-connection-id - The ID of the VPC peering connection.
	Filters []*Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// One or more VPC peering connection IDs.
	//
	// Default: Describes all your VPC peering connections.
	VpcPeeringConnectionIds []*string `locationName:"VpcPeeringConnectionId" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeVpcPeeringConnectionsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVpcPeeringConnectionsInput) GoString() string {
	return s.String()
}

// SetDryRun sets the DryRun field's value.
	//
	//    * cidr-block-association.cidr-block - An IPv4 CIDR block associated with
	//    the VPC.
	//
	//    * cidr-block-association.association-id - The association ID for an IPv4
	//    CIDR block associated with the VPC.
	//
	//    * cidr-block-association.state - The state of an IPv4 CIDR block associated
	//    with the VPC.
	//
	//    * dhcp-options-id - The ID of a set of DHCP options.
	//
	//    * ipv6-cidr-block-association.ipv6-cidr-block - An IPv6 CIDR block associated
	//    with the VPC.
	//
	//    * ipv6-cidr-block-association.association-id - The association ID for
	//    an IPv6 CIDR block associated with the VPC.
	//
	//    * ipv6-cidr-block-association.state - The state of an IPv6 CIDR block
	//    associated with the VPC.
	//
	//    * isDefault - Indicates whether the VPC is the default VPC.
	//
	//    * state - The state of the VPC (pending | available).
	//
	//    * tag:key=value - The key/value combination of a tag assigned to the resource.
	//    Specify the key of the tag in the filter name and the value of the tag
	//    in the filter value. For example, for the tag Purpose=X, specify tag:Purpose
	//    for the filter name and X for the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. This filter is
	//    independent of the tag-value filter. For example, if you use both the
	//    filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
	//    assigned both the tag key Purpose (regardless of what the tag's value
	//    is), and the tag value X (regardless of what the tag's key is). If you
	//    want to list only resources where Purpose is X, see the tag:key=value
	//    filter.
	//
	//    * tag-value - The value of a tag assigned to the resource. This filter
	//    is independent of the tag-key filter.
	//
	//    * vpc-id - The ID of the VPC.
	Filters []*Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// One or more VPC IDs.
	//
	// Default: Describes all your VPCs.
	VpcIds []*string `locationName:"VpcId" locationNameList:"VpcId" type:"list"`
}

// String returns the string representation
func (s DescribeVpcsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVpcsInput) GoString() string {
	return s.String()
}

// SetDryRun sets the DryRun field's value.
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// One or more filters.
	//
	//    * customer-gateway-configuration - The configuration information for the
	//    customer gateway.
	//
	//    * customer-gateway-id - The ID of a customer gateway associated with the
	//    VPN connection.
	//
	//    * state - The state of the VPN connection (pending | available | deleting
	//    | deleted).
	//
	//    * option.static-routes-only - Indicates whether the connection has static
	//    routes only. Used for devices that do not support Border Gateway Protocol
	//    (BGP).
	//
	//    * route.destination-cidr-block - The destination CIDR block. This corresponds
	//    to the subnet used in a customer data center.
	//
	//    * bgp-asn - The BGP Autonomous System Number (ASN) associated with a BGP
	//    device.
	//
	//    * tag:key=value - The key/value combination of a tag assigned to the resource.
	//    Specify the key of the tag in the filter name and the value of the tag
	//    in the filter value. For example, for the tag Purpose=X, specify tag:Purpose
	//    for the filter name and X for the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. This filter is
	//    independent of the tag-value filter. For example, if you use both the
	//    filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
	//    assigned both the tag key Purpose (regardless of what the tag's value
	//    is), and the tag value X (regardless of what the tag's key is). If you
	//    want to list only resources where Purpose is X, see the tag:key=value
	//    filter.
	//
	//    * tag-value - The value of a tag assigned to the resource. This filter
	//    is independent of the tag-key filter.
	//
	//    * type - The type of VPN connection. Currently the only supported type
	//    is ipsec.1.
	//
	//    * vpn-connection-id - The ID of the VPN connection.
	//
	//    * vpn-gateway-id - The ID of a virtual private gateway associated with
	//    the VPN connection.
	Filters []*Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// One or more VPN connection IDs.
	//
	// Default: Describes your VPN connections.
	VpnConnectionIds []*string `locationName:"VpnConnectionId" locationNameList:"VpnConnectionId" type:"list"`
}

// String returns the string representation
func (s DescribeVpnConnectionsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeVpnGatewaysInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// One or more filters.
	//
	//    * amazon-side-asn - The Autonomous System Number (ASN) for the Amazon
	//    side of the gateway.
	//
	//    * attachment.state - The current state of the attachment between the gateway
	//    and the VPC (attaching | attached | detaching | detached).
	//
	//    * attachment.vpc-id - The ID of an attached VPC.
	//
	//    * availability-zone - The Availability Zone for the virtual private gateway
	//    (if applicable).
	//
	//    * state - The state of the virtual private gateway (pending | available
	//    | deleting | deleted).
	//
	//    * tag:key=value - The key/value combination of a tag assigned to the resource.
	//    Specify the key of the tag in the filter name and the value of the tag
	//    in the filter value. For example, for the tag Purpose=X, specify tag:Purpose
	//    for the filter name and X for the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. This filter is
	//    independent of the tag-value filter. For example, if you use both the
	//    filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
	//    assigned both the tag key Purpose (regardless of what the tag's value
	//    is), and the tag value X (regardless of what the tag's key is). If you
	//    want to list only resources where Purpose is X, see the tag:key=value
	//    filter.
	//
	//    * tag-value - The value of a tag assigned to the resource. This filter
	//    is independent of the tag-key filter.
	//
	//    * type - The type of virtual private gateway. Currently the only supported
	//    type is ipsec.1.
	//
	//    * vpn-gateway-id - The ID of the virtual private gateway.
	Filters []*Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// One or more virtual private gateway IDs.
	//
	// Default: Describes all your virtual private gateways.
	VpnGatewayIds []*string `locationName:"VpnGatewayId" locationNameList:"VpnGatewayId" type:"list"`
}

// String returns the string representation
func (s DescribeVpnGatewaysInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVpnGatewaysInput) GoString() string {
	return s.String()

	// InstanceTypeM512xlarge is a InstanceType enum value
	InstanceTypeM512xlarge = "m5.12xlarge"

	// InstanceTypeM524xlarge is a InstanceType enum value
	InstanceTypeM524xlarge = "m5.24xlarge"

	// InstanceTypeH12xlarge is a InstanceType enum value
	InstanceTypeH12xlarge = "h1.2xlarge"

	// InstanceTypeH14xlarge is a InstanceType enum value
	InstanceTypeH14xlarge = "h1.4xlarge"

	// InstanceTypeH18xlarge is a InstanceType enum value
	InstanceTypeH18xlarge = "h1.8xlarge"

	// InstanceTypeH116xlarge is a InstanceType enum value
	InstanceTypeH116xlarge = "h1.16xlarge"
)

const (
	// InterfacePermissionTypeInstanceAttach is a InterfacePermissionType enum value
	InterfacePermissionTypeInstanceAttach = "INSTANCE-ATTACH"

	// InterfacePermissionTypeEipAssociate is a InterfacePermissionType enum value
	InterfacePermissionTypeEipAssociate = "EIP-ASSOCIATE"
)

const (
	// LaunchTemplateErrorCodeLaunchTemplateIdDoesNotExist is a LaunchTemplateErrorCode enum value
	LaunchTemplateErrorCodeLaunchTemplateIdDoesNotExist = "launchTemplateIdDoesNotExist"

	// LaunchTemplateErrorCodeLaunchTemplateIdMalformed is a LaunchTemplateErrorCode enum value
	LaunchTemplateErrorCodeLaunchTemplateIdMalformed = "launchTemplateIdMalformed"

	// LaunchTemplateErrorCodeLaunchTemplateNameDoesNotExist is a LaunchTemplateErrorCode enum value
	LaunchTemplateErrorCodeLaunchTemplateNameDoesNotExist = "launchTemplateNameDoesNotExist"

	// LaunchTemplateErrorCodeLaunchTemplateNameMalformed is a LaunchTemplateErrorCode enum value
	LaunchTemplateErrorCodeLaunchTemplateNameMalformed = "launchTemplateNameMalformed"

	// LaunchTemplateErrorCodeLaunchTemplateVersionDoesNotExist is a LaunchTemplateErrorCode enum value
	LaunchTemplateErrorCodeLaunchTemplateVersionDoesNotExist = "launchTemplateVersionDoesNotExist"

	// LaunchTemplateErrorCodeUnexpectedError is a LaunchTemplateErrorCode enum value
	LaunchTemplateErrorCodeUnexpectedError = "unexpectedError"
)

const (
	// ListingStateAvailable is a ListingState enum value
	ListingStateAvailable = "available"
	RIProductDescriptionLinuxUnixamazonVpc = "Linux/UNIX (Amazon VPC)"

	// RIProductDescriptionWindows is a RIProductDescription enum value
	RIProductDescriptionWindows = "Windows"

	// RIProductDescriptionWindowsAmazonVpc is a RIProductDescription enum value
	RIProductDescriptionWindowsAmazonVpc = "Windows (Amazon VPC)"
)

const (
	// RecurringChargeFrequencyHourly is a RecurringChargeFrequency enum value
	RecurringChargeFrequencyHourly = "Hourly"
)

const (
	// ReportInstanceReasonCodesInstanceStuckInState is a ReportInstanceReasonCodes enum value
	ReportInstanceReasonCodesInstanceStuckInState = "instance-stuck-in-state"

	// ReportInstanceReasonCodesUnresponsive is a ReportInstanceReasonCodes enum value
	ReportInstanceReasonCodesUnresponsive = "unresponsive"

	// ReportInstanceReasonCodesNotAcceptingCredentials is a ReportInstanceReasonCodes enum value
	ReportInstanceReasonCodesNotAcceptingCredentials = "not-accepting-credentials"

	// ReportInstanceReasonCodesPasswordNotAvailable is a ReportInstanceReasonCodes enum value
	ReportInstanceReasonCodesPasswordNotAvailable = "password-not-available"

	// ReportInstanceReasonCodesPerformanceNetwork is a ReportInstanceReasonCodes enum value
	ReportInstanceReasonCodesPerformanceNetwork = "performance-network"

	// ReportInstanceReasonCodesPerformanceInstanceStore is a ReportInstanceReasonCodes enum value
	ReportInstanceReasonCodesPerformanceInstanceStore = "performance-instance-store"

	// ReportInstanceReasonCodesPerformanceEbsVolume is a ReportInstanceReasonCodes enum value
	ReportInstanceReasonCodesPerformanceEbsVolume = "performance-ebs-volume"

	// ReportInstanceReasonCodesPerformanceOther is a ReportInstanceReasonCodes enum value
	ReportInstanceReasonCodesPerformanceOther = "performance-other"

	// ReportInstanceReasonCodesOther is a ReportInstanceReasonCodes enum value
	ReportInstanceReasonCodesOther = "other"
)

const (
	// ReportStatusTypeOk is a ReportStatusType enum value
	ReportStatusTypeOk = "ok"

	// ReportStatusTypeImpaired is a ReportStatusType enum value
	ReportStatusTypeImpaired = "impaired"
)

	// StateFailed is a State enum value
	StateFailed = "Failed"

	// StateExpired is a State enum value
	StateExpired = "Expired"
)

const (
	// StatusMoveInProgress is a Status enum value
	StatusMoveInProgress = "MoveInProgress"

	// StatusInVpc is a Status enum value
	StatusInVpc = "InVpc"

	// StatusInClassic is a Status enum value
	StatusInClassic = "InClassic"
)

const (
	// StatusNameReachability is a StatusName enum value
	StatusNameReachability = "reachability"
)

const (
	// StatusTypePassed is a StatusType enum value
	StatusTypePassed = "passed"

	// StatusTypeFailed is a StatusType enum value
	StatusTypeFailed = "failed"

	// StatusTypeInsufficientData is a StatusType enum value
	StatusTypeInsufficientData = "insufficient-data"

	// StatusTypeInitializing is a StatusType enum value
	StatusTypeInitializing = "initializing"
)

const (
	// SubnetCidrBlockStateCodeAssociating is a SubnetCidrBlockStateCode enum value
	SubnetCidrBlockStateCodeAssociating = "associating"

	// SubnetCidrBlockStateCodeAssociated is a SubnetCidrBlockStateCode enum value
	SubnetCidrBlockStateCodeAssociated = "associated"

	// SubnetCidrBlockStateCodeDisassociating is a SubnetCidrBlockStateCode enum value
	SubnetCidrBlockStateCodeDisassociating = "disassociating"

	// SubnetCidrBlockStateCodeDisassociated is a SubnetCidrBlockStateCode enum value
	SubnetCidrBlockStateCodeDisassociated = "disassociated"

	// SubnetCidrBlockStateCodeFailing is a SubnetCidrBlockStateCode enum value