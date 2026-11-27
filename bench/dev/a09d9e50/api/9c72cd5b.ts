export interface CompletionStats {
  size_in_bytes: long
  size?: ByteSize
  fields?: Record<Field, FieldSizeUsage>
}

export type Conflicts = 'abort' | 'proceed'

export type DataStreamName = string

export interface DateField {
  field: Field
  format?: string
  include_unmapped?: boolean
}

export type DateMath = string

export type DateMathTime = string

export type DateString = string

export type DefaultOperator = 'AND' | 'OR'

export interface DictionaryResponseBase<TKey = unknown, TValue = unknown> {
  [key: string]: TValue
}

export type Distance = string

export type DistanceUnit = 'in' | 'ft' | 'yd' | 'mi' | 'nmi' | 'km' | 'm' | 'cm' | 'mm'

export interface DocStats {
  count: long
  deleted: long
}

export interface ElasticsearchVersionInfo {
  build_date: DateString
  build_flavor: string
  build_hash: string
  build_snapshot: boolean
  build_type: string
  lucene_version: VersionString
  minimum_index_compatibility_version: VersionString
  minimum_wire_compatibility_version: VersionString
  number: string
}

export interface EmptyObject {
}

export type NodeRole = 'master' | 'data' | 'data_cold' | 'data_content' | 'data_frozen' | 'data_hot' | 'data_warm' | 'client' | 'ingest' | 'ml' | 'voting_only' | 'transform' | 'remote_cluster_client' | 'coordinating_only'

export type NodeRoles = NodeRole[]

export interface NodeShard {
  state: IndicesStatsShardRoutingState
  primary: boolean
  node?: NodeName
  shard: integer
  index: IndexName
  allocation_id?: Record<string, Id>
  recovery_source?: Record<string, Id>
  unassigned_info?: ClusterAllocationExplainUnassignedInformation
}

export interface NodeStatistics {
  failures?: ErrorCause[]
  total: integer
  successful: integer
  failed: integer
}

export type OpType = 'index' | 'create'

export type Password = string

export type Percentage = string | float

export type PipelineName = string

export interface PluginStats {
  classname: string
  description: string
  elasticsearch_version: VersionString
  extended_plugins: string[]
  has_native_controller: boolean
  java_version: VersionString
  name: Name
  version: VersionString
  licensed: boolean
  type: string
}

export type PropertyName = string

export interface QueryCacheStats {
  cache_count: integer
  cache_size: integer
  evictions: integer
  hit_count: integer
  script?: Script
  time_zone?: string
}

export interface AggregationsAverageAggregation extends AggregationsFormatMetricAggregationBase {
}

export interface AggregationsAverageBucketAggregation extends AggregationsPipelineAggregationBase {
}

export interface AggregationsBoxPlotAggregate extends AggregationsAggregateBase {
  min: double
  max: double
  q1: double
  q2: double
  q3: double
}

export interface AggregationsBoxplotAggregation extends AggregationsMetricAggregationBase {
  compression?: double
}

export type AggregationsBucket = AggregationsCompositeBucket | AggregationsDateHistogramBucket | AggregationsFiltersBucketItem | AggregationsIpRangeBucket | AggregationsRangeBucket | AggregationsRareTermsBucket<any> | AggregationsSignificantTermsBucket<any> | AggregationsKeyedBucket<any>

export interface AggregationsBucketAggregate extends AggregationsAggregateBase {
  after_key: Record<string, any>
  bg_count: long
  doc_count: long
  doc_count_error_upper_bound: long
  sum_other_doc_count: long
  interval: DateMathTime
  items: AggregationsBucket
}

export interface AggregationsBucketAggregationBase extends AggregationsAggregation {
  aggregations?: Record<string, AggregationsAggregationContainer>
}

export interface AggregationsBucketScriptAggregation extends AggregationsPipelineAggregationBase {
  script?: Script
}

export interface AggregationsBucketSelectorAggregation extends AggregationsPipelineAggregationBase {
  script?: Script
}

export interface AggregationsBucketSortAggregation extends AggregationsAggregation {
  from?: integer
  gap_policy?: AggregationsGapPolicy
  size?: integer
  sort?: SearchSort
  num_top_feature_importance_values?: integer
  prediction_field_type?: string
  results_field?: string
  top_classes_results_field?: string
}

export interface AggregationsCompositeAggregation extends AggregationsBucketAggregationBase {
  after?: Record<string, string | float | null>
  size?: integer
  sources?: Record<string, AggregationsCompositeAggregationSource>[]
}

export interface AggregationsCompositeAggregationSource {
  terms?: AggregationsTermsAggregation
  histogram?: AggregationsHistogramAggregation
  date_histogram?: AggregationsDateHistogramAggregation
  geotile_grid?: AggregationsGeoTileGridAggregation
}

export interface AggregationsCompositeBucketKeys {
}
export type AggregationsCompositeBucket = AggregationsCompositeBucketKeys |
    { [property: string]: AggregationsAggregate }

export interface AggregationsCompositeBucketAggregate extends AggregationsMultiBucketAggregate<Record<string, any>> {
  after_key: Record<string, any>
}

export interface AggregationsCumulativeCardinalityAggregation extends AggregationsPipelineAggregationBase {
}

export interface AggregationsCumulativeSumAggregation extends AggregationsPipelineAggregationBase {
}

export interface AggregationsDateHistogramAggregation extends AggregationsBucketAggregationBase {
  calendar_interval?: AggregationsDateInterval | Time
  extended_bounds?: AggregationsExtendedBounds<DateMath | long>
  hard_bounds?: AggregationsExtendedBounds<DateMath | long>
  field?: Field
  fixed_interval?: AggregationsDateInterval | Time
  format?: string
  interval?: AggregationsDateInterval | Time
  min_doc_count?: integer
  missing?: DateString
  offset?: Time
  order?: AggregationsHistogramOrder
  params?: Record<string, any>
  script?: Script
  time_zone?: string
}


export interface AggregationsGeoLineSort {
  field: Field
}

export interface AggregationsGeoTileGridAggregation extends AggregationsBucketAggregationBase {
  field?: Field
  precision?: GeoTilePrecision
  shard_size?: integer
  size?: integer
  bounds?: AggregationsGeoBounds
}

export interface AggregationsGlobalAggregation extends AggregationsBucketAggregationBase {
}

export interface AggregationsGoogleNormalizedDistanceHeuristic {
  background_is_superset: boolean
}

export interface AggregationsHdrMethod {
  number_of_significant_value_digits?: integer
}

export interface AggregationsHdrPercentileItem {
  key: double
  value: double
}

export interface AggregationsHdrPercentilesAggregate extends AggregationsAggregateBase {
  values: AggregationsHdrPercentileItem[]
}

export interface AggregationsHistogramAggregation extends AggregationsBucketAggregationBase {
  extended_bounds?: AggregationsExtendedBounds<double>
  hard_bounds?: AggregationsExtendedBounds<double>
  field?: Field
  interval?: double
  min_doc_count?: integer
  missing?: double
  offset?: double
  order?: AggregationsHistogramOrder
  script?: Script
  format?: string
}

export interface AggregationsHistogramOrder {
  _count?: SearchSortOrder
  _key?: SearchSortOrder
}


export interface AnalysisConditionTokenFilter extends AnalysisTokenFilterBase {
  filter: string[]
  script: Script
}

export type AnalysisDelimitedPayloadEncoding = 'int' | 'float' | 'identity'

export interface AnalysisDelimitedPayloadTokenFilter extends AnalysisTokenFilterBase {
  delimiter: string
  encoding: AnalysisDelimitedPayloadEncoding
}

export type AnalysisEdgeNGramSide = 'front' | 'back'

export interface AnalysisEdgeNGramTokenFilter extends AnalysisTokenFilterBase {
  max_gram: integer
  min_gram: integer
  side: AnalysisEdgeNGramSide
}

export interface AnalysisEdgeNGramTokenizer extends AnalysisTokenizerBase {
  custom_token_chars: string
  max_gram: integer
  min_gram: integer
  token_chars: AnalysisTokenChar[]
}

export interface AnalysisElisionTokenFilter extends AnalysisTokenFilterBase {
  articles: string[]
  articles_case: boolean
}

export interface AnalysisFingerprintTokenFilter extends AnalysisTokenFilterBase {
  max_output_size: integer
  separator: string
}

export interface AnalysisHtmlStripCharFilter extends AnalysisCharFilterBase {
}

export interface AnalysisHunspellTokenFilter extends AnalysisTokenFilterBase {
  dedup: boolean
  dictionary: string
  locale: string
  longest_only: boolean
}

export interface AnalysisHyphenationDecompounderTokenFilter extends AnalysisCompoundWordTokenFilterBase {
}

  keep_words_case: boolean
  keep_words_path: string
}

export interface AnalysisKeywordMarkerTokenFilter extends AnalysisTokenFilterBase {
  ignore_case: boolean
  keywords: string[]
  keywords_path: string
  keywords_pattern: string
}

export interface AnalysisKeywordTokenizer extends AnalysisTokenizerBase {
  buffer_size: integer
}

export interface AnalysisLengthTokenFilter extends AnalysisTokenFilterBase {
  max: integer
  min: integer
}

export interface AnalysisLetterTokenizer extends AnalysisTokenizerBase {
}

export interface AnalysisLimitTokenCountTokenFilter extends AnalysisTokenFilterBase {
  consume_all_tokens: boolean
  max_token_count: integer
}

export interface AnalysisLowercaseTokenFilter extends AnalysisTokenFilterBase {
  language: string
}

export interface AnalysisLowercaseTokenizer extends AnalysisTokenizerBase {
}

export interface AnalysisMappingCharFilter extends AnalysisCharFilterBase {
  mappings: string[]
  mappings_path: string
}

export interface AnalysisMultiplexerTokenFilter extends AnalysisTokenFilterBase {
  filters: string[]
  preserve_original: boolean
}

export interface AnalysisNGramTokenFilter extends AnalysisTokenFilterBase {
  max_gram: integer
  min_gram: integer
}

export interface AnalysisNGramTokenizer extends AnalysisTokenizerBase {
  custom_token_chars: string
  max_gram: integer
  min_gram: integer
  token_chars: AnalysisTokenChar[]
}

export type AnalysisNoriDecompoundMode = 'discard' | 'none' | 'mixed'

export interface AnalysisNoriPartOfSpeechTokenFilter extends AnalysisTokenFilterBase {
  stoptags: string[]
}

export interface AnalysisNoriTokenizer extends AnalysisTokenizerBase {
  decompound_mode: AnalysisNoriDecompoundMode
  discard_punctuation: boolean
  user_dictionary: string
  user_dictionary_rules: string[]
}

export interface AnalysisPathHierarchyTokenizer extends AnalysisTokenizerBase {
  buffer_size: integer
  delimiter: string
  replacement: string
  reverse: boolean
  skip: integer
}

export interface AnalysisPatternCaptureTokenFilter extends AnalysisTokenFilterBase {
  patterns: string[]
  preserve_original: boolean
}

export interface AnalysisPatternReplaceTokenFilter extends AnalysisTokenFilterBase {
  flags: string
  pattern: string
  replacement: string
}

export interface AnalysisPorterStemTokenFilter extends AnalysisTokenFilterBase {
}

export interface AnalysisPredicateTokenFilter extends AnalysisTokenFilterBase {
  script: Script
}

export interface AnalysisRemoveDuplicatesTokenFilter extends AnalysisTokenFilterBase {
}

export interface AnalysisReverseTokenFilter extends AnalysisTokenFilterBase {
}

export interface AnalysisShingleTokenFilter extends AnalysisTokenFilterBase {
  filler_token: string
  max_shingle_size: integer
  min_shingle_size: integer
  output_unigrams: boolean
  output_unigrams_if_no_shingles: boolean
  token_separator: string
}

export type AnalysisSnowballLanguage = 'Armenian' | 'Basque' | 'Catalan' | 'Danish' | 'Dutch' | 'English' | 'Finnish' | 'French' | 'German' | 'German2' | 'Hungarian' | 'Italian' | 'Kp' | 'Lovins' | 'Norwegian' | 'Porter' | 'Portuguese' | 'Romanian' | 'Russian' | 'Spanish' | 'Swedish' | 'Turkish'

export interface AnalysisSnowballTokenFilter extends AnalysisTokenFilterBase {
  language: AnalysisSnowballLanguage
}

export interface AnalysisStandardTokenizer extends AnalysisTokenizerBase {
  max_token_length: integer
}

export interface AnalysisStemmerOverrideTokenFilter extends AnalysisTokenFilterBase {
  rules: string[]
  rules_path: string
}

export interface AnalysisStemmerTokenFilter extends AnalysisTokenFilterBase {
  language: string
}

export interface AnalysisStopTokenFilter extends AnalysisTokenFilterBase {
  ignore_case?: boolean
  remove_trailing?: boolean
  stopwords: AnalysisStopWords
  stopwords_path?: string
}

export type AnalysisStopWords = string | string[]

export type AnalysisSynonymFormat = 'solr' | 'wordnet'

export interface AnalysisSynonymGraphTokenFilter extends AnalysisTokenFilterBase {
  expand: boolean
  format: AnalysisSynonymFormat

export type AnalysisTokenChar = 'letter' | 'digit' | 'whitespace' | 'punctuation' | 'symbol' | 'custom'

export type AnalysisTokenFilter = AnalysisAsciiFoldingTokenFilter | AnalysisCommonGramsTokenFilter | AnalysisConditionTokenFilter | AnalysisDelimitedPayloadTokenFilter | AnalysisEdgeNGramTokenFilter | AnalysisElisionTokenFilter | AnalysisFingerprintTokenFilter | AnalysisHunspellTokenFilter | AnalysisHyphenationDecompounderTokenFilter | AnalysisKeepTypesTokenFilter | AnalysisKeepWordsTokenFilter | AnalysisKeywordMarkerTokenFilter | AnalysisKStemTokenFilter | AnalysisLengthTokenFilter | AnalysisLimitTokenCountTokenFilter | AnalysisLowercaseTokenFilter | AnalysisMultiplexerTokenFilter | AnalysisNGramTokenFilter | AnalysisNoriPartOfSpeechTokenFilter | AnalysisPatternCaptureTokenFilter | AnalysisPatternReplaceTokenFilter | AnalysisPorterStemTokenFilter | AnalysisPredicateTokenFilter | AnalysisRemoveDuplicatesTokenFilter | AnalysisReverseTokenFilter | AnalysisShingleTokenFilter | AnalysisSnowballTokenFilter | AnalysisStemmerOverrideTokenFilter | AnalysisStemmerTokenFilter | AnalysisStopTokenFilter | AnalysisSynonymGraphTokenFilter | AnalysisSynonymTokenFilter | AnalysisTrimTokenFilter | AnalysisTruncateTokenFilter | AnalysisUniqueTokenFilter | AnalysisUppercaseTokenFilter | AnalysisWordDelimiterGraphTokenFilter | AnalysisWordDelimiterTokenFilter

export interface AnalysisTokenFilterBase {
  type: string
  version?: VersionString
}

export type AnalysisTokenizer = AnalysisCharGroupTokenizer | AnalysisEdgeNGramTokenizer | AnalysisKeywordTokenizer | AnalysisLetterTokenizer | AnalysisLowercaseTokenizer | AnalysisNGramTokenizer | AnalysisNoriTokenizer | AnalysisPathHierarchyTokenizer | AnalysisStandardTokenizer | AnalysisUaxEmailUrlTokenizer | AnalysisWhitespaceTokenizer

export interface AnalysisTokenizerBase {
  type: string
  version?: VersionString
}

export interface AnalysisTrimTokenFilter extends AnalysisTokenFilterBase {
}

export interface AnalysisTruncateTokenFilter extends AnalysisTokenFilterBase {
  length: integer
}

export interface AnalysisUaxEmailUrlTokenizer extends AnalysisTokenizerBase {
  max_token_length: integer
}

export interface AnalysisUniqueTokenFilter extends AnalysisTokenFilterBase {
  only_on_same_position: boolean
}

export interface AnalysisUppercaseTokenFilter extends AnalysisTokenFilterBase {
}

export interface AnalysisWhitespaceTokenizer extends AnalysisTokenizerBase {
  max_token_length: integer
}

export interface AnalysisWordDelimiterGraphTokenFilter extends AnalysisTokenFilterBase {
  adjust_offsets: boolean
  catenate_all: boolean
  catenate_numbers: boolean
  catenate_words: boolean
  generate_number_parts: boolean
  generate_word_parts: boolean
  preserve_original: boolean
  protected_words: string[]
  protected_words_path: string
  split_on_case_change: boolean
  split_on_numerics: boolean
  stem_english_possessive: boolean
  type_table: string[]
  type_table_path: string
}

export interface AnalysisWordDelimiterTokenFilter extends AnalysisTokenFilterBase {
  catenate_all: boolean
  catenate_numbers: boolean
  catenate_words: boolean
  generate_number_parts: boolean
  generate_word_parts: boolean
  allow_no_indices?: boolean
  expand_wildcards?: ExpandWildcards
  ignore_unavailable?: boolean
  master_timeout?: Time
  timeout?: Time
}

export interface IndicesAddBlockResponse extends AcknowledgedResponseBase {
  shards_acknowledged: boolean
  indices: IndicesAddBlockIndicesBlockStatus[]
}

export interface IndicesAnalyzeAnalyzeDetail {
  analyzer?: IndicesAnalyzeAnalyzerDetail
  charfilters?: IndicesAnalyzeCharFilterDetail[]
  custom_analyzer: boolean
  tokenfilters?: IndicesAnalyzeTokenDetail[]
  tokenizer?: IndicesAnalyzeTokenDetail
}

export interface IndicesAnalyzeAnalyzeToken {
  end_offset: long
  position: long
  position_length?: long
  start_offset: long
  token: string
  type: string
}

export interface IndicesAnalyzeAnalyzerDetail {
  name: string
  tokens: IndicesAnalyzeExplainAnalyzeToken[]
}

export interface IndicesAnalyzeCharFilterDetail {
  filtered_text: string[]
  name: string
}

export interface IndicesAnalyzeExplainAnalyzeToken {
  bytes: string
  end_offset: long
  keyword?: boolean
  position: long
  positionLength: long
  start_offset: long
  termFrequency: long
  token: string
  type: string
}

export interface IndicesAnalyzeRequest extends RequestBase {
  index?: IndexName
  body?: {
    analyzer?: string
    attributes?: string[]
    char_filter?: (string | AnalysisCharFilter)[]
    explain?: boolean
    field?: Field
    filter?: (string | AnalysisTokenFilter)[]
    normalizer?: string
    text?: IndicesAnalyzeTextToAnalyze
    tokenizer?: string | AnalysisTokenizer
  }
}

export interface IndicesAnalyzeResponse {
  detail?: IndicesAnalyzeAnalyzeDetail
  tokens?: IndicesAnalyzeAnalyzeToken[]
}

export type IndicesAnalyzeTextToAnalyze = string | string[]

}

export interface NodesInfoNodeInfoXpack {
  license?: NodesInfoNodeInfoXpackLicense
  security: NodesInfoNodeInfoXpackSecurity
  notification?: Record<string, any>
}

export interface NodesInfoNodeInfoXpackLicense {
  self_generated: NodesInfoNodeInfoXpackLicenseType
}

export interface NodesInfoNodeInfoXpackLicenseType {
  type: string
}

export interface NodesInfoNodeInfoXpackSecurity {
  http: NodesInfoNodeInfoXpackSecuritySsl
  enabled: string
  transport: NodesInfoNodeInfoXpackSecuritySsl
  authc?: NodesInfoNodeInfoXpackSecurityAuthc
}

export interface NodesInfoNodeInfoXpackSecurityAuthc {
  realms: NodesInfoNodeInfoXpackSecurityAuthcRealms
  token: NodesInfoNodeInfoXpackSecurityAuthcToken
}

export interface NodesInfoNodeInfoXpackSecurityAuthcRealms {
  file?: Record<string, NodesInfoNodeInfoXpackSecurityAuthcRealmsStatus>
  native?: Record<string, NodesInfoNodeInfoXpackSecurityAuthcRealmsStatus>
  pki?: Record<string, NodesInfoNodeInfoXpackSecurityAuthcRealmsStatus>
}

export interface NodesInfoNodeInfoXpackSecurityAuthcRealmsStatus {
  enabled?: string
  order: string
}

export interface NodesInfoNodeInfoXpackSecurityAuthcToken {
  enabled: string
}

export interface NodesInfoNodeInfoXpackSecuritySsl {
  ssl: Record<string, string>
}

export interface NodesInfoNodeJvmInfo {
  gc_collectors: string[]
  mem: NodesInfoNodeInfoJvmMemory
  memory_pools: string[]
  name: Names
}

export interface SecurityClearCachedServiceTokensResponse {
  _nodes: NodeStatistics
  cluster_name: Name
  nodes: Record<string, SecurityClusterNode>
}

export interface SecurityCreateApiKeyIndexPrivileges {
  names: Indices
  privileges: string[]
}

export interface SecurityCreateApiKeyRequest extends RequestBase {
  refresh?: Refresh
  body?: {
    expiration?: Time
    name?: Name
    role_descriptors?: Record<string, SecurityCreateApiKeyRoleDescriptor>
    metadata?: Metadata
  }
}

export interface SecurityCreateApiKeyResponse {
  api_key: string
  expiration?: long
  id: Id
  name: Name
}

export interface SecurityCreateApiKeyRoleDescriptor {
  cluster: string[]
  index: SecurityCreateApiKeyIndexPrivileges[]
  applications?: SecurityApplicationPrivileges[]
}

export interface SecurityCreateServiceTokenRequest extends RequestBase {
  namespace: Namespace
  service: Service
  name: Name
}

export interface SecurityCreateServiceTokenResponse {
  created: boolean
  token: SecurityCreateServiceTokenToken
}

export interface SecurityCreateServiceTokenToken {
  name: Name
  value: string
}

export interface SecurityDeletePrivilegesFoundStatus {
  found: boolean
}

export interface SecurityDeletePrivilegesRequest extends RequestBase {
  application: Name
  name: Name
  refresh?: Refresh
}

export interface SecurityDeletePrivilegesResponse extends DictionaryResponseBase<string, Record<string, SecurityDeletePrivilegesFoundStatus>> {
}

export interface SecurityDeleteRoleRequest extends RequestBase {
  name: Name
  refresh?: Refresh
}


export interface SecurityGetTokenAuthenticatedUser extends SecurityUser {
  authentication_realm: SecurityGetTokenUserRealm
  lookup_realm: SecurityGetTokenUserRealm
  authentication_provider?: SecurityGetTokenAuthenticationProvider
  authentication_type: string
}

export interface SecurityGetTokenAuthenticationProvider {
  type: string
  name: Name
}

export interface SecurityGetTokenRequest extends RequestBase {
  body?: {
    grant_type?: SecurityGetTokenAccessTokenGrantType
    scope?: string
    password?: Password
    kerberos_ticket?: string
    refresh_token?: string
    username?: Username
  }
}

export interface SecurityGetTokenResponse {
  access_token: string
  expires_in: long
  scope?: string
  type: string
  refresh_token: string
  kerberos_authentication_response_token?: string
  authentication: SecurityGetTokenAuthenticatedUser
}

export interface SecurityGetTokenUserRealm {
  name: Name
  type: string
}

export interface SecurityGetUserRequest extends RequestBase {
  username?: Username | Username[]
}

export interface SecurityGetUserResponse extends DictionaryResponseBase<string, SecurityUser> {
}

export interface SecurityGetUserPrivilegesRequest extends RequestBase {
  application?: Name
  priviledge?: Name
}

export interface SecurityGetUserPrivilegesResponse {
  applications: SecurityApplicationPrivileges[]
  cluster: string[]
  global: SecurityGlobalPrivileges[]
  indices: SecurityIndicesPrivileges[]
  run_as: string[]
}

export interface SecurityGrantApiKeyApiKey {
  name: Name
  expiration?: Time
  role_descriptors?: Record<string, any>[]
}

export type SecurityGrantApiKeyApiKeyGrantType = 'access_token' | 'password'

export interface SecurityGrantApiKeyRequest extends RequestBase {
  body?: {
    api_key: SecurityGrantApiKeyApiKey
    grant_type: SecurityGrantApiKeyApiKeyGrantType
    access_token?: string
    username?: Username
    password?: Password
  }
}

export interface SecurityGrantApiKeyResponse {
  api_key: string
  id: Id
  name: Name
  expiration?: EpochMillis
}

export interface SecurityHasPrivilegesApplicationPrivilegesCheck {
  application: string
  privileges: string[]
  resources: string[]
}

export type SecurityHasPrivilegesApplicationsPrivileges = Record<Name, SecurityHasPrivilegesResourcePrivileges>

export interface SecurityHasPrivilegesIndexPrivilegesCheck {
  names: string[]
  privileges: string[]
}

export type SecurityHasPrivilegesPrivileges = Record<string, boolean>

export interface SecurityHasPrivilegesRequest extends RequestBase {
  user?: Name
  body?: {
    application?: SecurityHasPrivilegesApplicationPrivilegesCheck[]
    cluster?: string[]
    owner?: boolean
    realm_name?: string
    username?: Username
  }
}

export interface SecurityInvalidateApiKeyResponse {
  error_count: integer
  error_details?: ErrorCause[]
  invalidated_api_keys: string[]
  previously_invalidated_api_keys: string[]
}

export interface SecurityInvalidateTokenRequest extends RequestBase {
  body?: {
    token?: string
    refresh_token?: string
    realm_name?: Name
    username?: Username
  }
}

export interface SecurityInvalidateTokenResponse {
  error_count: long
  error_details?: ErrorCause[]
  invalidated_tokens: long
  previously_invalidated_tokens: long
}

export interface SecurityPutPrivilegesActions {
  actions: string[]
  application?: string
  name?: Name
  metadata?: Metadata
}

export interface SecurityPutPrivilegesRequest extends RequestBase {
  refresh?: Refresh
  body?: Record<string, Record<string, SecurityPutPrivilegesActions>>
}

export interface SecurityPutPrivilegesResponse extends DictionaryResponseBase<string, Record<string, SecurityCreatedStatus>> {
}

export interface SecurityPutRoleRequest extends RequestBase {
  name: Name
  refresh?: Refresh
  body?: {
    applications?: SecurityApplicationPrivileges[]
    cluster?: string[]
    global?: Record<string, any>
    indices?: SecurityIndicesPrivileges[]
  rows: SqlQueryRow[]
}

export type SqlQueryRow = any[]

export interface SqlTranslateRequest extends RequestBase {
  body?: {
    fetch_size?: integer
    filter?: QueryDslQueryContainer
    query: string
    time_zone?: string
  }
}

export interface SqlTranslateResponse {
  size: long
  _source: boolean | Fields | SearchSourceFilter
  fields: Record<Field, string>[]
  sort: SearchSort
}

export interface SslGetCertificatesCertificateInformation {
  alias?: string
  expiry: DateString
  format: string
  has_private_key: boolean
  path: string
  serial_number: string
  subject_dn: string
}

export interface SslGetCertificatesRequest extends RequestBase {
}

export type SslGetCertificatesResponse = SslGetCertificatesCertificateInformation[]

export interface TaskInfo {
  action: string
  cancellable: boolean
  children?: TaskInfo[]
  description?: string
  headers: HttpHeaders
  id: long
  node: string
  running_time_in_nanos: long
  start_time_in_millis: long
  status?: TaskStatus
  type: string
  parent_task_id?: Id
}

}

export interface WatcherPagerDutyActionEventResult {
  event: WatcherPagerDutyEvent
  reason: string
  request: WatcherHttpInputRequestResult
  response: WatcherHttpInputResponseResult
}

export interface WatcherPagerDutyContext {
  href: string
  src: string
  type: WatcherPagerDutyContextType
}

export type WatcherPagerDutyContextType = 'link' | 'image'

export interface WatcherPagerDutyEvent {
  account: string
  attach_payload: boolean
  client: string
  client_url: string
  context: WatcherPagerDutyContext[]
  description: string
  event_type: WatcherPagerDutyEventType
  incident_key: string
}

export type WatcherPagerDutyEventType = 'trigger' | 'resolve' | 'acknowledge'

export interface WatcherPagerDutyResult {
  sent_event: WatcherPagerDutyActionEventResult
}

export type WatcherQuantifier = 'some' | 'all'

export type WatcherResponseContentType = 'json' | 'yaml' | 'text'

export interface WatcherScheduleBase {
}

export interface WatcherScheduleContainer {
  cron?: WatcherCronExpression
  daily?: WatcherDailySchedule
  hourly?: WatcherHourlySchedule
  interval?: Time
  monthly?: WatcherTimeOfMonth[]
  weekly?: WatcherTimeOfWeek[]
  yearly?: WatcherTimeOfYear[]
}
