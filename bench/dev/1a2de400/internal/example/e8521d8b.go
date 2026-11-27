	Uuid                string                                  `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Nested              []*ABitOfEverything_Nested              `protobuf:"bytes,2,rep,name=nested,proto3" json:"nested,omitempty"`
	FloatValue          float32                                 `protobuf:"fixed32,3,opt,name=float_value,json=floatValue,proto3" json:"float_value,omitempty"`
	DoubleValue         float64                                 `protobuf:"fixed64,4,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	Int64Value          int64                                   `protobuf:"varint,5,opt,name=int64_value,json=int64Value,proto3" json:"int64_value,omitempty"`
	Uint64Value         uint64                                  `protobuf:"varint,6,opt,name=uint64_value,json=uint64Value,proto3" json:"uint64_value,omitempty"`
	Int32Value          int32                                   `protobuf:"varint,7,opt,name=int32_value,json=int32Value,proto3" json:"int32_value,omitempty"`
	Fixed64Value        uint64                                  `protobuf:"fixed64,8,opt,name=fixed64_value,json=fixed64Value,proto3" json:"fixed64_value,omitempty"`
	Fixed32Value        uint32                                  `protobuf:"fixed32,9,opt,name=fixed32_value,json=fixed32Value,proto3" json:"fixed32_value,omitempty"`
	BoolValue           bool                                    `protobuf:"varint,10,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	StringValue         string                                  `protobuf:"bytes,11,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	BytesValue          []byte                                  `protobuf:"bytes,29,opt,name=bytes_value,json=bytesValue,proto3" json:"bytes_value,omitempty"`
	Uint32Value         uint32                                  `protobuf:"varint,13,opt,name=uint32_value,json=uint32Value,proto3" json:"uint32_value,omitempty"`
	EnumValue           NumericEnum                             `protobuf:"varint,14,opt,name=enum_value,json=enumValue,proto3,enum=grpc.gateway.examples.internal.proto.examplepb.NumericEnum" json:"enum_value,omitempty"`
	PathEnumValue       pathenum.PathEnum                       `protobuf:"varint,30,opt,name=path_enum_value,json=pathEnumValue,proto3,enum=grpc.gateway.examples.internal.pathenum.PathEnum" json:"path_enum_value,omitempty"`
	NestedPathEnumValue pathenum.MessagePathEnum_NestedPathEnum `protobuf:"varint,31,opt,name=nested_path_enum_value,json=nestedPathEnumValue,proto3,enum=grpc.gateway.examples.internal.pathenum.MessagePathEnum_NestedPathEnum" json:"nested_path_enum_value,omitempty"`
	Sfixed32Value       int32                                   `protobuf:"fixed32,15,opt,name=sfixed32_value,json=sfixed32Value,proto3" json:"sfixed32_value,omitempty"`
	Sfixed64Value       int64                                   `protobuf:"fixed64,16,opt,name=sfixed64_value,json=sfixed64Value,proto3" json:"sfixed64_value,omitempty"`
	Sint32Value         int32                                   `protobuf:"zigzag32,17,opt,name=sint32_value,json=sint32Value,proto3" json:"sint32_value,omitempty"`
	Sint64Value         int64                                   `protobuf:"zigzag64,18,opt,name=sint64_value,json=sint64Value,proto3" json:"sint64_value,omitempty"`
	RepeatedStringValue []string                                `protobuf:"bytes,19,rep,name=repeated_string_value,json=repeatedStringValue,proto3" json:"repeated_string_value,omitempty"`
	// Types that are assignable to OneofValue:
	//	*ABitOfEverything_OneofEmpty
	//	*ABitOfEverything_OneofString
	OneofValue               isABitOfEverything_OneofValue       `protobuf_oneof:"oneof_value"`
	MapValue                 map[string]NumericEnum              `protobuf:"bytes,22,rep,name=map_value,json=mapValue,proto3" json:"map_value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=grpc.gateway.examples.internal.proto.examplepb.NumericEnum"`
	MappedStringValue        map[string]string                   `protobuf:"bytes,23,rep,name=mapped_string_value,json=mappedStringValue,proto3" json:"mapped_string_value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MappedNestedValue        map[string]*ABitOfEverything_Nested `protobuf:"bytes,24,rep,name=mapped_nested_value,json=mappedNestedValue,proto3" json:"mapped_nested_value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	NonConventionalNameValue string                              `protobuf:"bytes,26,opt,name=nonConventionalNameValue,proto3" json:"nonConventionalNameValue,omitempty"`
	TimestampValue           *timestamppb.Timestamp              `protobuf:"bytes,27,opt,name=timestamp_value,json=timestampValue,proto3" json:"timestamp_value,omitempty"`
	// repeated enum value. it is comma-separated in query
	RepeatedEnumValue []NumericEnum `protobuf:"varint,28,rep,packed,name=repeated_enum_value,json=repeatedEnumValue,proto3,enum=grpc.gateway.examples.internal.proto.examplepb.NumericEnum" json:"repeated_enum_value,omitempty"`
	// repeated numeric enum comment (This comment is overridden by the field annotation)
	RepeatedEnumAnnotation []NumericEnum `protobuf:"varint,32,rep,packed,name=repeated_enum_annotation,json=repeatedEnumAnnotation,proto3,enum=grpc.gateway.examples.internal.proto.examplepb.NumericEnum" json:"repeated_enum_annotation,omitempty"`
	// numeric enum comment (This comment is overridden by the field annotation)
	EnumValueAnnotation NumericEnum `protobuf:"varint,33,opt,name=enum_value_annotation,json=enumValueAnnotation,proto3,enum=grpc.gateway.examples.internal.proto.examplepb.NumericEnum" json:"enum_value_annotation,omitempty"`
	// repeated string comment (This comment is overridden by the field annotation)
	RepeatedStringAnnotation []string `protobuf:"bytes,34,rep,name=repeated_string_annotation,json=repeatedStringAnnotation,proto3" json:"repeated_string_annotation,omitempty"`
	// repeated nested object comment (This comment is overridden by the field annotation)
	RepeatedNestedAnnotation []*ABitOfEverything_Nested `protobuf:"bytes,35,rep,name=repeated_nested_annotation,json=repeatedNestedAnnotation,proto3" json:"repeated_nested_annotation,omitempty"`
	// nested object comments (This comment is overridden by the field annotation)
	NestedAnnotation  *ABitOfEverything_Nested `protobuf:"bytes,36,opt,name=nested_annotation,json=nestedAnnotation,proto3" json:"nested_annotation,omitempty"`
	Int64OverrideType int64                    `protobuf:"varint,37,opt,name=int64_override_type,json=int64OverrideType,proto3" json:"int64_override_type,omitempty"`
	// mark a field as required in Open API definition
	RequiredStringViaFieldBehaviorAnnotation string `protobuf:"bytes,38,opt,name=required_string_via_field_behavior_annotation,json=requiredStringViaFieldBehaviorAnnotation,proto3" json:"required_string_via_field_behavior_annotation,omitempty"`
	// mark a field as readonly in Open API definition
	OutputOnlyStringViaFieldBehaviorAnnotation string  `protobuf:"bytes,39,opt,name=output_only_string_via_field_behavior_annotation,json=outputOnlyStringViaFieldBehaviorAnnotation,proto3" json:"output_only_string_via_field_behavior_annotation,omitempty"`
	OptionalStringValue                        *string `protobuf:"bytes,40,opt,name=optional_string_value,json=optionalStringValue,proto3,oneof" json:"optional_string_value,omitempty"`
}

func (x *ABitOfEverything) Reset() {
	*x = ABitOfEverything{}
	if protoimpl.UnsafeEnabled {