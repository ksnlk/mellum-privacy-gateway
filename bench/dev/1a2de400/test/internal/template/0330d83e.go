			if err != nil {
				t.Fatalf("marshaler.Unmarshal failed: %v", err)
			}
			want := []string{"baz"}
			if diff := cmp.Diff(got, want); diff != "" {
				t.Errorf(diff)
			}
		}
		{
			var got string
			err = marshaler.Unmarshal(result.extensions[1].value, &got)
			if err != nil {
				t.Fatalf("marshaler.Unmarshal failed: %v", err)
			}
			want := "bar"
			if diff := cmp.Diff(got, want); diff != "" {
				t.Errorf(diff)
			}
		}

		var scheme openapiSecuritySchemeObject
		for _, v := range result.SecurityDefinitions {
			scheme = v
		}
		if want, is, name := []extension{
			{key: "x-security-baz", value: json.RawMessage("true")},
		}, scheme.extensions, "SecurityScheme.Extensions"; !reflect.DeepEqual(is, want) {
			t.Errorf("applyTemplate(%#v).%s = %s want to be %s", file, name, is, want)
		}

		if want, is, name := []extension{
			{key: "x-info-extension", value: json.RawMessage("\"bar\"")},
		}, result.Info.extensions, "Info.Extensions"; !reflect.DeepEqual(is, want) {
			t.Errorf("applyTemplate(%#v).%s = %s want to be %s", file, name, is, want)
		}

		var operation *openapiOperationObject
		var response openapiResponseObject
		for _, v := range result.Paths {
			operation = v.Get
			response = v.Get.Responses["200"]
		}
		if want, is, name := []extension{
			{key: "x-op-foo", value: json.RawMessage("\"baz\"")},
		}, operation.extensions, "operation.Extensions"; !reflect.DeepEqual(is, want) {
			t.Errorf("applyTemplate(%#v).%s = %s want to be %s", file, name, is, want)
		}
		if want, is, name := []extension{
			{key: "x-resp-id", value: json.RawMessage("\"resp1000\"")},
		}, response.extensions, "response.Extensions"; !reflect.DeepEqual(is, want) {
			t.Errorf("applyTemplate(%#v).%s = %s want to be %s", file, name, is, want)
		}
	}
	t.Run("verify template options set via proto options", func(t *testing.T) {
		file := newFile()
		proto.SetExtension(proto.Message(file.FileDescriptorProto.Options), openapi_options.E_Openapiv2Swagger, &swagger)
		proto.SetExtension(proto.Message(file.Services[0].Methods[0].Options), openapi_options.E_Openapiv2Operation, &openapiOperation)
		reg := descriptor.NewRegistry()
		verifyTemplateExtensions(t, reg, file, nil)
	})
	t.Run("verify template options set via annotations", func(t *testing.T) {
		file := newFile()
		opts := &openapiconfig.OpenAPIOptions{
			File: []*openapiconfig.OpenAPIFileOption{
				{
					File:   "example.proto",
					Option: &swagger,
				},
			},
			Method: []*openapiconfig.OpenAPIMethodOption{
				{
					Method: "example.ExampleService.Example",
					Option: &openapiOperation,
				},
							Number:  proto.Int32(1),
							Options: requiredField,
						},
					},
				},
			},
			schema: map[string]openapi_options.Schema{
				"Message": {
					JsonSchema: &openapi_options.JSONSchema{
						Title:       "title",
						Description: "desc",
						Required:    []string{"req"},
					},
				},
			},
			defs: map[string]openapiSchemaObject{
				"Message": {
					schemaCore: schemaCore{
						Type: "object",
					},
					Title:       "title",
					Description: "desc",
					Required:    []string{"req", "aRequiredField"},
					Properties: &openapiSchemaObjectProperties{
						{
							Key: "aRequiredField",
							Value: openapiSchemaObject{
								schemaCore: schemaCore{
									Type: "string",
								},
								Description: "field description",
								Title:       "field title",
								Required:    []string{"aRequiredField"},
							},
						},
					},
				},
			},
		},
		{
			descr: "JSONSchema with excluded fields",
			msgDescs: []*descriptorpb.DescriptorProto{
				{
					Name: proto.String("Message"),
					Field: []*descriptorpb.FieldDescriptorProto{
						{
							Name:    proto.String("aRequiredField"),
							Type:    descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
							Number:  proto.Int32(1),
							Options: requiredField,
						},
							Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
							Number: proto.Int32(2),
						},
					},
				},
			},
			schema: map[string]openapi_options.Schema{
				"Message": {
					JsonSchema: &openapi_options.JSONSchema{
						Title:       "title",
						Description: "desc",
						Required:    []string{"req"},
					},
				},
			},
			defs: map[string]openapiSchemaObject{
				"Message": {
					schemaCore: schemaCore{
						Type: "object",
					},
					Title:       "title",
					Description: "desc",
					Required:    []string{"req", "aRequiredField"},
					Properties: &openapiSchemaObjectProperties{
						{
							Key: "aRequiredField",
							Value: openapiSchemaObject{
								schemaCore: schemaCore{
									Type: "string",
								},
								Description: "field description",
								Title:       "field title",
								Required:    []string{"aRequiredField"},
							},
						},
					},
				},
			},
			excludedFields: []*descriptor.Field{
				{
					FieldDescriptorProto: &descriptorpb.FieldDescriptorProto{
						Name: strPtr("anExcludedField"),
					},
				},
			},
		},
		{
			descr: "JSONSchema with required properties via field_behavior",
			msgDescs: []*descriptorpb.DescriptorProto{
				{
					Name: proto.String("Message"),
							Number:  proto.Int32(2),
							Options: fieldBehaviorOutputOnlyOptions,
						},
					},
				},
			},
			schema: map[string]openapi_options.Schema{
				"Message": {
					JsonSchema: &openapi_options.JSONSchema{
						Title:       "title",
						Description: "desc",
						Required:    []string{"req"},
					},
				},
			},
			defs: map[string]openapiSchemaObject{
				"Message": {
					schemaCore: schemaCore{
						Type: "object",
					},
					Title:       "title",
					Description: "desc",
					Required:    []string{"req", "aRequiredField"},
					Properties: &openapiSchemaObjectProperties{
						{
							Key: "aRequiredField",
							Value: openapiSchemaObject{
								schemaCore: schemaCore{
									Type: "string",
								},
								Required: []string{"aRequiredField"},
							},
						},
						{
							Key: "aOutputOnlyField",
							Value: openapiSchemaObject{
								schemaCore: schemaCore{
									Type: "string",
								},
								ReadOnly: true,
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.descr, func(t *testing.T) {

			msgs := []*descriptor.Message{}
			for _, msgdesc := range test.msgDescs {
				msgdesc.Options = &descriptorpb.MessageOptions{}
				msgs = append(msgs, &descriptor.Message{DescriptorProto: msgdesc})
			}

			reg := descriptor.NewRegistry()
			file := descriptor.File{
				FileDescriptorProto: &descriptorpb.FileDescriptorProto{
			comments:      "First line.\n\nSecond line",
			expectedError: nil,
		},
		{
			descr:                "multi line comment with summary with dot",
			openapiSwaggerObject: &openapiOperationObject{},
			expectedOpenAPIObject: &openapiOperationObject{
				Summary:     "First line.",
				Description: "Second line",
			},
			comments:      "First line.\n\nSecond line",
			expectedError: nil,
		},
		{
			descr:                "multi line comment with summary no dot",
			openapiSwaggerObject: &openapiOperationObject{},
			expectedOpenAPIObject: &openapiOperationObject{
				Summary:     "First line",
				Description: "Second line",
			},
			comments:      "First line\n\nSecond line",
			expectedError: nil,
		},
		{
			descr:                 "multi line comment with summary no dot",
			openapiSwaggerObject:  &schemaCore{},
			expectedOpenAPIObject: &schemaCore{},
			comments:              "Any comment",
			expectedError:         errors.New("no description nor summary property"),
		},
		{
			descr:                "without use_go_template",
			openapiSwaggerObject: &openapiSchemaObject{},
			expectedOpenAPIObject: &openapiSchemaObject{
				Title:       "First line",
				Description: "{{import \"documentation.md\"}}",
			},
			comments:      "First line\n\n{{import \"documentation.md\"}}",
			expectedError: nil,
		},
		{
			descr:                "error with use_go_template",
			openapiSwaggerObject: &openapiSchemaObject{},
			expectedOpenAPIObject: &openapiSchemaObject{
				Title:       "First line",
				Description: "open noneexistingfile.txt: no such file or directory",
			},
			comments:      "First line\n\n{{import \"noneexistingfile.txt\"}}",
			expectedError: nil,
			useGoTemplate: true,
		},
		{
						"b_1": "b1JSONNAME",
					},
				},
				fieldNameToType: map[string]string{
					"abc": "pkg1.test.Msg",
					"bcd": "pkg1.test.Msg",
				},
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getReservedJSONName(tt.args.fieldName, tt.args.messageNameToFieldsToJSONName, tt.args.fieldNameToType); got != tt.want {
				t.Errorf("getReservedJSONName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseIncompleteSecurityRequirement(t *testing.T) {
	swagger := openapi_options.Swagger{
		Security: []*openapi_options.SecurityRequirement{
			{
				SecurityRequirement: map[string]*openapi_options.SecurityRequirement_SecurityRequirementValue{
					"key": nil,
				},
			},
		},
	}
	file := descriptor.File{
		FileDescriptorProto: &descriptorpb.FileDescriptorProto{
			SourceCodeInfo: &descriptorpb.SourceCodeInfo{},
			Name:           proto.String("example.proto"),
			Package:        proto.String("example"),
			Options: &descriptorpb.FileOptions{
				GoPackage: proto.String("github.com/grpc-ecosystem/grpc-gateway/runtime/internal/examplepb;example"),
			},
		},
	}
	proto.SetExtension(proto.Message(file.FileDescriptorProto.Options), openapi_options.E_Openapiv2Swagger, &swagger)
	reg := descriptor.NewRegistry()
	err := reg.Load(&pluginpb.CodeGeneratorRequest{ProtoFile: []*descriptorpb.FileDescriptorProto{file.FileDescriptorProto}})
	if err != nil {
		t.Errorf("failed to reg.Load(): %v", err)
		return
	}
	_, err = applyTemplate(param{File: crossLinkFixture(&file), reg: reg})
	if err == nil {
		t.Errorf("applyTemplate(%#v) did not error as expected", file)
		return