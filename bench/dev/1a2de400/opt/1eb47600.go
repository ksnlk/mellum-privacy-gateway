	return p
}

func (x SecurityScheme_In) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SecurityScheme_In) Descriptor() protoreflect.EnumDescriptor {
	return file_protoc_gen_openapiv2_options_openapiv2_proto_enumTypes[3].Descriptor()
}

func (SecurityScheme_In) Type() protoreflect.EnumType {
	return &file_protoc_gen_openapiv2_options_openapiv2_proto_enumTypes[3]
}

func (x SecurityScheme_In) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SecurityScheme_In.Descriptor instead.
func (SecurityScheme_In) EnumDescriptor() ([]byte, []int) {
	return file_protoc_gen_openapiv2_options_openapiv2_proto_rawDescGZIP(), []int{12, 1}
}

// The flow used by the OAuth2 security scheme. Valid values are
// "implicit", "password", "application" or "accessCode".
type SecurityScheme_Flow int32

const (
	SecurityScheme_FLOW_INVALID     SecurityScheme_Flow = 0
	SecurityScheme_FLOW_IMPLICIT    SecurityScheme_Flow = 1
	SecurityScheme_FLOW_PASSWORD    SecurityScheme_Flow = 2
	SecurityScheme_FLOW_APPLICATION SecurityScheme_Flow = 3
	SecurityScheme_FLOW_ACCESS_CODE SecurityScheme_Flow = 4
)

// Enum value maps for SecurityScheme_Flow.
var (
	SecurityScheme_Flow_name = map[int32]string{
		0: "FLOW_INVALID",
		1: "FLOW_IMPLICIT",
		2: "FLOW_PASSWORD",
		3: "FLOW_APPLICATION",
		4: "FLOW_ACCESS_CODE",
	}
	SecurityScheme_Flow_value = map[string]int32{
		"FLOW_INVALID":     0,
		"FLOW_IMPLICIT":    1,
		"FLOW_PASSWORD":    2,
		"FLOW_APPLICATION": 3,
		"FLOW_ACCESS_CODE": 4,
	Host string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	// The base path on which the API is served, which is relative to the host. If
	// it is not included, the API is served directly under the host. The value
	// MUST start with a leading slash (/). The basePath does not support path
	// templating.
	// Note that using `base_path` does not change the endpoint paths that are
	// generated in the resulting OpenAPI file. If you wish to use `base_path`
	// with relatively generated OpenAPI paths, the `base_path` prefix must be
	// manually removed from your `google.api.http` paths and your code changed to
	// serve the API from the `base_path`.
	BasePath string `protobuf:"bytes,4,opt,name=base_path,json=basePath,proto3" json:"base_path,omitempty"`
	// The transfer protocol of the API. Values MUST be from the list: "http",
	// "https", "ws", "wss". If the schemes is not included, the default scheme to
	// be used is the one used to access the OpenAPI definition itself.
	Schemes []Scheme `protobuf:"varint,5,rep,packed,name=schemes,proto3,enum=grpc.gateway.protoc_gen_openapiv2.options.Scheme" json:"schemes,omitempty"`
	// A list of MIME types the APIs can consume. This is global to all APIs but
	// can be overridden on specific API calls. Value MUST be as described under
	// Mime Types.
	Consumes []string `protobuf:"bytes,6,rep,name=consumes,proto3" json:"consumes,omitempty"`
	// A list of MIME types the APIs can produce. This is global to all APIs but
	// can be overridden on specific API calls. Value MUST be as described under
	// Mime Types.
	Produces []string `protobuf:"bytes,7,rep,name=produces,proto3" json:"produces,omitempty"`
	// An object to hold responses that can be used across operations. This
	// property does not define global responses for all operations.
	Responses map[string]*Response `protobuf:"bytes,10,rep,name=responses,proto3" json:"responses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Security scheme definitions that can be used across the specification.
	SecurityDefinitions *SecurityDefinitions `protobuf:"bytes,11,opt,name=security_definitions,json=securityDefinitions,proto3" json:"security_definitions,omitempty"`
	// A declaration of which security schemes are applied for the API as a whole.
	// The list of values describes alternative security schemes that can be used
	// (that is, there is a logical OR between the security requirements).
	// Individual operations can override this definition.
	Security []*SecurityRequirement `protobuf:"bytes,12,rep,name=security,proto3" json:"security,omitempty"`
	// Additional external documentation.
	ExternalDocs *ExternalDocumentation     `protobuf:"bytes,14,opt,name=external_docs,json=externalDocs,proto3" json:"external_docs,omitempty"`
	Extensions   map[string]*structpb.Value `protobuf:"bytes,15,rep,name=extensions,proto3" json:"extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Swagger) Reset() {
	*x = Swagger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_gen_openapiv2_options_openapiv2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Swagger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Swagger) ProtoMessage() {}

func (x *Swagger) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_gen_openapiv2_options_openapiv2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	// grouping of operations by resources or any other qualifier.
	Tags []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	// A short summary of what the operation does. For maximum readability in the
	// swagger-ui, this field SHOULD be less than 120 characters.
	Summary string `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	// A verbose explanation of the operation behavior. GFM syntax can be used for
	// rich text representation.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Additional external documentation for this operation.
	ExternalDocs *ExternalDocumentation `protobuf:"bytes,4,opt,name=external_docs,json=externalDocs,proto3" json:"external_docs,omitempty"`
	// Unique string used to identify the operation. The id MUST be unique among
	// all operations described in the API. Tools and libraries MAY use the
	// operationId to uniquely identify an operation, therefore, it is recommended
	// to follow common programming naming conventions.
	OperationId string `protobuf:"bytes,5,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	// A list of MIME types the operation can consume. This overrides the consumes
	// definition at the OpenAPI Object. An empty value MAY be used to clear the
	// global definition. Value MUST be as described under Mime Types.
	Consumes []string `protobuf:"bytes,6,rep,name=consumes,proto3" json:"consumes,omitempty"`
	// A list of MIME types the operation can produce. This overrides the produces
	// definition at the OpenAPI Object. An empty value MAY be used to clear the
	// global definition. Value MUST be as described under Mime Types.
	Produces []string `protobuf:"bytes,7,rep,name=produces,proto3" json:"produces,omitempty"`
	// The list of possible responses as they are returned from executing this
	// operation.
	Responses map[string]*Response `protobuf:"bytes,9,rep,name=responses,proto3" json:"responses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The transfer protocol for the operation. Values MUST be from the list:
	// "http", "https", "ws", "wss". The value overrides the OpenAPI Object
	// schemes definition.
	Schemes []Scheme `protobuf:"varint,10,rep,packed,name=schemes,proto3,enum=grpc.gateway.protoc_gen_openapiv2.options.Scheme" json:"schemes,omitempty"`
	// Declares this operation to be deprecated. Usage of the declared operation
	// should be refrained. Default value is false.
	Deprecated bool `protobuf:"varint,11,opt,name=deprecated,proto3" json:"deprecated,omitempty"`
	// A declaration of which security schemes are applied for this operation. The
	// list of values describes alternative security schemes that can be used
	// (that is, there is a logical OR between the security requirements). This
	// definition overrides any declared top-level security. To remove a top-level
	// security declaration, an empty array can be used.
	Security   []*SecurityRequirement     `protobuf:"bytes,12,rep,name=security,proto3" json:"security,omitempty"`
	Extensions map[string]*structpb.Value `protobuf:"bytes,13,rep,name=extensions,proto3" json:"extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_gen_openapiv2_options_openapiv2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_gen_openapiv2_options_openapiv2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
func (x *Header) GetPattern() string {
	if x != nil {
		return x.Pattern
	}
	return ""
}

// `Response` is a representation of OpenAPI v2 specification's Response object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#responseObject
//
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// `Description` is a short description of the response.
	// GFM syntax can be used for rich text representation.
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// `Schema` optionally defines the structure of the response.
	// If `Schema` is not provided, it means there is no content to the response.
	Schema *Schema `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
	// `Headers` A list of headers that are sent with the response.
	// `Header` name is expected to be a string in the canonical format of the MIME header key
	// See: https://golang.org/pkg/net/textproto/#CanonicalMIMEHeaderKey
	Headers map[string]*Header `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// `Examples` gives per-mimetype response examples.
	// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#example-object
	Examples   map[string]string          `protobuf:"bytes,4,rep,name=examples,proto3" json:"examples,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Extensions map[string]*structpb.Value `protobuf:"bytes,5,rep,name=extensions,proto3" json:"extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_gen_openapiv2_options_openapiv2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_gen_openapiv2_options_openapiv2_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
//        url: "https://github.com/grpc-ecosystem/grpc-gateway/blob/master/LICENSE.txt";
//      };
//    };
//    ...
//  };
//
type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The title of the application.
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// A short description of the application. GFM syntax can be used for rich
	// text representation.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The Terms of Service for the API.
	TermsOfService string `protobuf:"bytes,3,opt,name=terms_of_service,json=termsOfService,proto3" json:"terms_of_service,omitempty"`
	// The contact information for the exposed API.
	Contact *Contact `protobuf:"bytes,4,opt,name=contact,proto3" json:"contact,omitempty"`
	// The license information for the exposed API.
	License *License `protobuf:"bytes,5,opt,name=license,proto3" json:"license,omitempty"`
	// Provides the version of the application API (not to be confused
	// with the specification version).
	Version    string                     `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
	Extensions map[string]*structpb.Value `protobuf:"bytes,7,rep,name=extensions,proto3" json:"extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_gen_openapiv2_options_openapiv2_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_gen_openapiv2_options_openapiv2_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	return ""
}

func (x *Tag) GetExternalDocs() *ExternalDocumentation {
	if x != nil {
		return x.ExternalDocs
	}
	return nil
}

// `SecurityDefinitions` is a representation of OpenAPI v2 specification's
// Security Definitions object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#securityDefinitionsObject
//
// A declaration of the security schemes available to be used in the
// specification. This does not enforce the security schemes on the operations
// and only serves to provide the relevant details for each scheme.
type SecurityDefinitions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A single security scheme definition, mapping a "name" to the scheme it
	// defines.
	Security map[string]*SecurityScheme `protobuf:"bytes,1,rep,name=security,proto3" json:"security,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SecurityDefinitions) Reset() {
	*x = SecurityDefinitions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_gen_openapiv2_options_openapiv2_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityDefinitions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityDefinitions) ProtoMessage() {}

func (x *SecurityDefinitions) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_gen_openapiv2_options_openapiv2_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#securitySchemeObject
//
// Allows the definition of a security scheme that can be used by the
// operations. Supported schemes are basic authentication, an API key (either as
// a header or as a query parameter) and OAuth2's common flows (implicit,
// password, application and access code).
type SecurityScheme struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of the security scheme. Valid values are "basic",
	// "apiKey" or "oauth2".
	Type SecurityScheme_Type `protobuf:"varint,1,opt,name=type,proto3,enum=grpc.gateway.protoc_gen_openapiv2.options.SecurityScheme_Type" json:"type,omitempty"`
	// A short description for security scheme.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The name of the header or query parameter to be used.
	// Valid for apiKey.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The location of the API key. Valid values are "query" or
	// "header".
	// Valid for apiKey.
	In SecurityScheme_In `protobuf:"varint,4,opt,name=in,proto3,enum=grpc.gateway.protoc_gen_openapiv2.options.SecurityScheme_In" json:"in,omitempty"`
	// The flow used by the OAuth2 security scheme. Valid values are
	// "implicit", "password", "application" or "accessCode".
	// Valid for oauth2.
	Flow SecurityScheme_Flow `protobuf:"varint,5,opt,name=flow,proto3,enum=grpc.gateway.protoc_gen_openapiv2.options.SecurityScheme_Flow" json:"flow,omitempty"`
	// The authorization URL to be used for this flow. This SHOULD be in
	// the form of a URL.
	// Valid for oauth2/implicit and oauth2/accessCode.
	AuthorizationUrl string `protobuf:"bytes,6,opt,name=authorization_url,json=authorizationUrl,proto3" json:"authorization_url,omitempty"`
	// The token URL to be used for this flow. This SHOULD be in the
	// form of a URL.
	// Valid for oauth2/password, oauth2/application and oauth2/accessCode.
	TokenUrl string `protobuf:"bytes,7,opt,name=token_url,json=tokenUrl,proto3" json:"token_url,omitempty"`
	// The available scopes for the OAuth2 security scheme.
	// Valid for oauth2.
	Scopes     *Scopes                    `protobuf:"bytes,8,opt,name=scopes,proto3" json:"scopes,omitempty"`
	Extensions map[string]*structpb.Value `protobuf:"bytes,9,rep,name=extensions,proto3" json:"extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SecurityScheme) Reset() {
	*x = SecurityScheme{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_gen_openapiv2_options_openapiv2_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityScheme) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityScheme) ProtoMessage() {}

func (x *SecurityScheme) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_gen_openapiv2_options_openapiv2_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
		return x.Extensions
	}
	return nil
}

// `SecurityRequirement` is a representation of OpenAPI v2 specification's
// Security Requirement object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#securityRequirementObject
//
// Lists the required security schemes to execute this operation. The object can
// have multiple security schemes declared in it which are all required (that
// is, there is a logical AND between the schemes).
//
// The name used for each property MUST correspond to a security scheme
// declared in the Security Definitions.
type SecurityRequirement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Each name must correspond to a security scheme which is declared in
	// the Security Definitions. If the security scheme is of type "oauth2",
	// then the value is a list of scope names required for the execution.
	// For other security scheme types, the array MUST be empty.
	SecurityRequirement map[string]*SecurityRequirement_SecurityRequirementValue `protobuf:"bytes,1,rep,name=security_requirement,json=securityRequirement,proto3" json:"security_requirement,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SecurityRequirement) Reset() {
	*x = SecurityRequirement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_gen_openapiv2_options_openapiv2_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityRequirement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityRequirement) ProtoMessage() {}

func (x *SecurityRequirement) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_gen_openapiv2_options_openapiv2_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms

// Deprecated: Use SecurityRequirement.ProtoReflect.Descriptor instead.
func (*SecurityRequirement) Descriptor() ([]byte, []int) {
	return file_protoc_gen_openapiv2_options_openapiv2_proto_rawDescGZIP(), []int{13}
}

func (x *SecurityRequirement) GetSecurityRequirement() map[string]*SecurityRequirement_SecurityRequirementValue {
	if x != nil {
		return x.SecurityRequirement
	}
	return nil
}

// `Scopes` is a representation of OpenAPI v2 specification's Scopes object.
//
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#scopesObject
//
// Lists the available scopes for an OAuth2 security scheme.
type Scopes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Maps between a name of a scope to a short description of it (as the value
	// of the property).
	Scope map[string]string `protobuf:"bytes,1,rep,name=scope,proto3" json:"scope,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Scopes) Reset() {
	*x = Scopes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_gen_openapiv2_options_openapiv2_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scopes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scopes) ProtoMessage() {}

func (x *Scopes) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_gen_openapiv2_options_openapiv2_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms