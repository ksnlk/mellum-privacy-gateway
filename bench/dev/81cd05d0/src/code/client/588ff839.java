    /** Provider for AWS credentials. */
    private final AWSCredentialsProvider awsCredentialsProvider;

    private static final Log log = LogFactory.getLog(AmazonCodeDeploy.class);

    /** Default signing name for the service. */
    private static final String DEFAULT_SIGNING_NAME = "codedeploy";

    private volatile AmazonCodeDeployWaiters waiters;

    /** Client configuration factory providing ClientConfigurations tailored to this client */
    protected static final ClientConfigurationFactory configFactory = new ClientConfigurationFactory();

    private final AdvancedConfig advancedConfig;

    private static final com.amazonaws.protocol.json.SdkJsonProtocolFactory protocolFactory = new com.amazonaws.protocol.json.SdkJsonProtocolFactory(
            new JsonClientMetadata()
                    .withProtocolVersion("1.1")
                    .withSupportsCbor(false)
                    .withSupportsIon(false)
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("ResourceValidationException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.ResourceValidationExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("GitHubAccountTokenDoesNotExistException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.GitHubAccountTokenDoesNotExistExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("InvalidOperationException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.InvalidOperationExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("InvalidTagsToAddException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.InvalidTagsToAddExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("TagSetListLimitExceededException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.TagSetListLimitExceededExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("DeploymentConfigDoesNotExistException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.DeploymentConfigDoesNotExistExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("ApplicationNameRequiredException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.ApplicationNameRequiredExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("InvalidLifecycleEventHookExecutionIdException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.InvalidLifecycleEventHookExecutionIdExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("InvalidAutoRollbackConfigException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.InvalidAutoRollbackConfigExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("DeploymentLimitExceededException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.DeploymentLimitExceededExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("TriggerTargetsLimitExceededException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.TriggerTargetsLimitExceededExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("InvalidLoadBalancerInfoException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.InvalidLoadBalancerInfoExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("MultipleIamArnsProvidedException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.MultipleIamArnsProvidedExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("InvalidDeploymentConfigNameException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.InvalidDeploymentConfigNameExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("LifecycleHookLimitExceededException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.LifecycleHookLimitExceededExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("IamUserArnAlreadyRegisteredException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.IamUserArnAlreadyRegisteredExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("RevisionRequiredException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.RevisionRequiredExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("DeploymentIsNotInReadyStateException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.DeploymentIsNotInReadyStateExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("InvalidGitHubAccountTokenNameException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.InvalidGitHubAccountTokenNameExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("InstanceNotRegisteredException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.InstanceNotRegisteredExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("InvalidBlueGreenDeploymentConfigurationException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.InvalidBlueGreenDeploymentConfigurationExceptionUnmarshaller
                                            .getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("ApplicationAlreadyExistsException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.ApplicationAlreadyExistsExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("InvalidDeployedStateFilterException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.InvalidDeployedStateFilterExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("DeploymentTargetIdRequiredException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.DeploymentTargetIdRequiredExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("InstanceDoesNotExistException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.InstanceDoesNotExistExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("InvalidECSServiceException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.InvalidECSServiceExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("InvalidAlarmConfigException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.InvalidAlarmConfigExceptionUnmarshaller.getInstance()))
                            new JsonErrorShapeMetadata().withErrorCode("InvalidTargetInstancesException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.InvalidTargetInstancesExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("TagRequiredException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.TagRequiredExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("BucketNameFilterRequiredException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.BucketNameFilterRequiredExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("InvalidIamUserArnException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.InvalidIamUserArnExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("InvalidInstanceTypeException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.InvalidInstanceTypeExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("InvalidTriggerConfigException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.InvalidTriggerConfigExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("IamArnRequiredException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.IamArnRequiredExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("InvalidKeyPrefixFilterException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.InvalidKeyPrefixFilterExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("GitHubAccountTokenNameRequiredException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.GitHubAccountTokenNameRequiredExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("ECSServiceMappingLimitExceededException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.ECSServiceMappingLimitExceededExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("InvalidBucketNameFilterException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.InvalidBucketNameFilterExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("InvalidComputePlatformException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.InvalidComputePlatformExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("RoleRequiredException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.RoleRequiredExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("ThrottlingException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.ThrottlingExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("DeploymentNotStartedException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.DeploymentNotStartedExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("LifecycleEventAlreadyCompletedException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.LifecycleEventAlreadyCompletedExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
                            new JsonErrorShapeMetadata().withErrorCode("InvalidIamSessionArnException").withExceptionUnmarshaller(
                                    com.amazonaws.services.codedeploy.model.transform.InvalidIamSessionArnExceptionUnmarshaller.getInstance()))
                    .addErrorMetadata(
     *         The specified resource could not be validated.
     * @throws OperationNotSupportedException
     *         The API used does not support the deployment.
     * @sample AmazonCodeDeploy.DeleteGitHubAccountToken
     * @see <a href="http://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/DeleteGitHubAccountToken"
     *      target="_top">AWS API Documentation</a>
     */
    @Override
    public DeleteGitHubAccountTokenResult deleteGitHubAccountToken(DeleteGitHubAccountTokenRequest request) {
        request = beforeClientExecution(request);
        return executeDeleteGitHubAccountToken(request);
    }

    @SdkInternalApi
    final DeleteGitHubAccountTokenResult executeDeleteGitHubAccountToken(DeleteGitHubAccountTokenRequest deleteGitHubAccountTokenRequest) {

        ExecutionContext executionContext = createExecutionContext(deleteGitHubAccountTokenRequest);
        AWSRequestMetrics awsRequestMetrics = executionContext.getAwsRequestMetrics();
        awsRequestMetrics.startEvent(Field.ClientExecuteTime);
        Request<DeleteGitHubAccountTokenRequest> request = null;
        Response<DeleteGitHubAccountTokenResult> response = null;

        try {
            awsRequestMetrics.startEvent(Field.RequestMarshallTime);
            try {
                request = new DeleteGitHubAccountTokenRequestProtocolMarshaller(protocolFactory).marshall(super
                        .beforeMarshalling(deleteGitHubAccountTokenRequest));
                // Binds the request metrics to the current request.
                request.setAWSRequestMetrics(awsRequestMetrics);
                request.addHandlerContext(HandlerContextKey.CLIENT_ENDPOINT, endpoint);
                request.addHandlerContext(HandlerContextKey.ENDPOINT_OVERRIDDEN, isEndpointOverridden());
                request.addHandlerContext(HandlerContextKey.SIGNING_REGION, getSigningRegion());
                request.addHandlerContext(HandlerContextKey.SERVICE_ID, "CodeDeploy");
                request.addHandlerContext(HandlerContextKey.OPERATION_NAME, "DeleteGitHubAccountToken");
                request.addHandlerContext(HandlerContextKey.ADVANCED_CONFIG, advancedConfig);

            } finally {
                awsRequestMetrics.endEvent(Field.RequestMarshallTime);
            }

            HttpResponseHandler<AmazonWebServiceResponse<DeleteGitHubAccountTokenResult>> responseHandler = protocolFactory.createResponseHandler(
                    new JsonOperationMetadata().withPayloadJson(true).withHasStreamingSuccessResponse(false),
                    new DeleteGitHubAccountTokenResultJsonUnmarshaller());
            response = invoke(request, responseHandler, executionContext);

            return response.getAwsResponse();

        } finally {

            endClientExecution(awsRequestMetrics, request, response);
        }
     *         The specified resource could not be validated.
     * @throws OperationNotSupportedException
     *         The API used does not support the deployment.
     * @sample AmazonCodeDeploy.ListGitHubAccountTokenNames
     * @see <a href="http://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/ListGitHubAccountTokenNames"
     *      target="_top">AWS API Documentation</a>
     */
    @Override
    public ListGitHubAccountTokenNamesResult listGitHubAccountTokenNames(ListGitHubAccountTokenNamesRequest request) {
        request = beforeClientExecution(request);
        return executeListGitHubAccountTokenNames(request);
    }

    @SdkInternalApi
    final ListGitHubAccountTokenNamesResult executeListGitHubAccountTokenNames(ListGitHubAccountTokenNamesRequest listGitHubAccountTokenNamesRequest) {

        ExecutionContext executionContext = createExecutionContext(listGitHubAccountTokenNamesRequest);
        AWSRequestMetrics awsRequestMetrics = executionContext.getAwsRequestMetrics();
        awsRequestMetrics.startEvent(Field.ClientExecuteTime);
        Request<ListGitHubAccountTokenNamesRequest> request = null;
        Response<ListGitHubAccountTokenNamesResult> response = null;

        try {
            awsRequestMetrics.startEvent(Field.RequestMarshallTime);
            try {
                request = new ListGitHubAccountTokenNamesRequestProtocolMarshaller(protocolFactory).marshall(super
                        .beforeMarshalling(listGitHubAccountTokenNamesRequest));
                // Binds the request metrics to the current request.
                request.setAWSRequestMetrics(awsRequestMetrics);
                request.addHandlerContext(HandlerContextKey.CLIENT_ENDPOINT, endpoint);
                request.addHandlerContext(HandlerContextKey.ENDPOINT_OVERRIDDEN, isEndpointOverridden());
                request.addHandlerContext(HandlerContextKey.SIGNING_REGION, getSigningRegion());
                request.addHandlerContext(HandlerContextKey.SERVICE_ID, "CodeDeploy");
                request.addHandlerContext(HandlerContextKey.OPERATION_NAME, "ListGitHubAccountTokenNames");
                request.addHandlerContext(HandlerContextKey.ADVANCED_CONFIG, advancedConfig);

            } finally {
                awsRequestMetrics.endEvent(Field.RequestMarshallTime);
            }

            HttpResponseHandler<AmazonWebServiceResponse<ListGitHubAccountTokenNamesResult>> responseHandler = protocolFactory.createResponseHandler(
                    new JsonOperationMetadata().withPayloadJson(true).withHasStreamingSuccessResponse(false),
                    new ListGitHubAccountTokenNamesResultJsonUnmarshaller());
            response = invoke(request, responseHandler, executionContext);

            return response.getAwsResponse();

        } finally {

            endClientExecution(awsRequestMetrics, request, response);
        }
    }

    /**
     * <p>
     * Gets a list of names for one or more on-premises instances.
     * </p>
     * <p>
     * Unless otherwise specified, both registered and deregistered on-premises instance names are listed. To list only
     * registered or deregistered on-premises instance names, use the registration status parameter.
     * </p>
     * 
     * @param listOnPremisesInstancesRequest
     *        Represents the input of a <code>ListOnPremisesInstances</code> operation.
     * @return Result of the ListOnPremisesInstances operation returned by the service.
     * @throws InvalidRegistrationStatusException
     *         The registration status was specified in an invalid format.
     * @throws InvalidTagFilterException