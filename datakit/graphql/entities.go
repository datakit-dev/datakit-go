// Code generated by go generate; DO NOT EDIT.
// This file was generated from GraphQL schema

package graphql

type EnumSchemaNamespace string

const EnumSchemaNamespacesync EnumSchemaNamespace = "sync"
const EnumSchemaNamespaceview EnumSchemaNamespace = "view"
const EnumSchemaNamespaceworkspace EnumSchemaNamespace = "workspace"

type EnumSyncStatus string

const EnumSyncStatusERROR EnumSyncStatus = "ERROR"
const EnumSyncStatusPENDING EnumSyncStatus = "PENDING"
const EnumSyncStatusRUNNING EnumSyncStatus = "RUNNING"
const EnumSyncStatusSCHEDULED EnumSyncStatus = "SCHEDULED"
const EnumSyncStatusSUCCESS EnumSyncStatus = "SUCCESS"

type Enum__TypeKind string

const Enum__TypeKindSCALAR Enum__TypeKind = "SCALAR"
const Enum__TypeKindOBJECT Enum__TypeKind = "OBJECT"
const Enum__TypeKindINTERFACE Enum__TypeKind = "INTERFACE"
const Enum__TypeKindUNION Enum__TypeKind = "UNION"
const Enum__TypeKindENUM Enum__TypeKind = "ENUM"
const Enum__TypeKindINPUT_OBJECT Enum__TypeKind = "INPUT_OBJECT"
const Enum__TypeKindLIST Enum__TypeKind = "LIST"
const Enum__TypeKindNON_NULL Enum__TypeKind = "NON_NULL"

type Enum__DirectiveLocation string

const Enum__DirectiveLocationQUERY Enum__DirectiveLocation = "QUERY"
const Enum__DirectiveLocationMUTATION Enum__DirectiveLocation = "MUTATION"
const Enum__DirectiveLocationSUBSCRIPTION Enum__DirectiveLocation = "SUBSCRIPTION"
const Enum__DirectiveLocationFIELD Enum__DirectiveLocation = "FIELD"
const Enum__DirectiveLocationFRAGMENT_DEFINITION Enum__DirectiveLocation = "FRAGMENT_DEFINITION"
const Enum__DirectiveLocationFRAGMENT_SPREAD Enum__DirectiveLocation = "FRAGMENT_SPREAD"
const Enum__DirectiveLocationINLINE_FRAGMENT Enum__DirectiveLocation = "INLINE_FRAGMENT"
const Enum__DirectiveLocationSCHEMA Enum__DirectiveLocation = "SCHEMA"
const Enum__DirectiveLocationSCALAR Enum__DirectiveLocation = "SCALAR"
const Enum__DirectiveLocationOBJECT Enum__DirectiveLocation = "OBJECT"
const Enum__DirectiveLocationFIELD_DEFINITION Enum__DirectiveLocation = "FIELD_DEFINITION"
const Enum__DirectiveLocationARGUMENT_DEFINITION Enum__DirectiveLocation = "ARGUMENT_DEFINITION"
const Enum__DirectiveLocationINTERFACE Enum__DirectiveLocation = "INTERFACE"
const Enum__DirectiveLocationUNION Enum__DirectiveLocation = "UNION"
const Enum__DirectiveLocationENUM Enum__DirectiveLocation = "ENUM"
const Enum__DirectiveLocationENUM_VALUE Enum__DirectiveLocation = "ENUM_VALUE"
const Enum__DirectiveLocationINPUT_OBJECT Enum__DirectiveLocation = "INPUT_OBJECT"
const Enum__DirectiveLocationINPUT_FIELD_DEFINITION Enum__DirectiveLocation = "INPUT_FIELD_DEFINITION"

type __Schema struct {
	Types            []__Type      `json:"types"`
	QueryType        __Type        `json:"queryType"`
	MutationType     *__Type       `json:"mutationType"`
	SubscriptionType *__Type       `json:"subscriptionType"`
	Directives       []__Directive `json:"directives"`
}

type __Type struct {
	Kind          Enum__TypeKind `json:"kind"`
	Name          *string        `json:"name"`
	Description   *string        `json:"description"`
	Fields        []__Field      `json:"fields"`
	Interfaces    []__Type       `json:"interfaces"`
	PossibleTypes []__Type       `json:"possibleTypes"`
	EnumValues    []__EnumValue  `json:"enumValues"`
	InputFields   []__InputValue `json:"inputFields"`
	OfType        *__Type        `json:"ofType"`
}

type __Field struct {
	Name              string         `json:"name"`
	Description       *string        `json:"description"`
	Args              []__InputValue `json:"args"`
	Type              __Type         `json:"type"`
	IsDeprecated      bool           `json:"isDeprecated"`
	DeprecationReason *string        `json:"deprecationReason"`
}

type __InputValue struct {
	Name         string  `json:"name"`
	Description  *string `json:"description"`
	Type         __Type  `json:"type"`
	DefaultValue *string `json:"defaultValue"`
}

type __EnumValue struct {
	Name              string  `json:"name"`
	Description       *string `json:"description"`
	IsDeprecated      bool    `json:"isDeprecated"`
	DeprecationReason *string `json:"deprecationReason"`
}

type __Directive struct {
	Name        string                    `json:"name"`
	Description *string                   `json:"description"`
	Locations   []Enum__DirectiveLocation `json:"locations"`
	Args        []__InputValue            `json:"args"`
}

type AirbyteCatalog struct {
	Streams []AirbyteStream `json:"streams"`
}

type AirbyteRecord struct {
	Data      JSON    `json:"data"`
	EmittedAt int64   `json:"emittedAt"`
	Namespace *string `json:"namespace"`
	Stream    string  `json:"stream"`
}

type AirbyteStream struct {
	Default_cursor_field       []*string  `json:"default_cursor_field"`
	Json_schema                JSONSchema `json:"json_schema"`
	Name                       string     `json:"name"`
	Namespace                  *string    `json:"namespace"`
	Source_defined_cursor      *bool      `json:"source_defined_cursor"`
	Source_defined_primary_key []*string  `json:"source_defined_primary_key"`
	Supported_sync_modes       []*string  `json:"supported_sync_modes"`
}

type Assistant struct {
	Chats       []*Chat             `json:"chats"`
	CreatedAt   DateTime            `json:"createdAt"`
	Datasets    []*AssistantDataset `json:"datasets"`
	Description string              `json:"description"`
	Id          string              `json:"id"`
	Name        string              `json:"name"`
	UpdatedAt   DateTime            `json:"updatedAt"`
	Workspace   Workspace           `json:"workspace"`
	WorkspaceId string              `json:"workspaceId"`
}

type AssistantDataset struct {
	Assistant      Assistant                `json:"assistant"`
	AssistantId    string                   `json:"assistantId"`
	Columns        []AssistantDatasetColumn `json:"columns"`
	Id             string                   `json:"id"`
	IdColumn       Column                   `json:"idColumn"`
	IdColumnId     string                   `json:"idColumnId"`
	Name           string                   `json:"name"`
	ResultFunction *string                  `json:"resultFunction"`
	Table          Table                    `json:"table"`
	TableId        string                   `json:"tableId"`
}

type AssistantDatasetColumn struct {
	Column     Column           `json:"column"`
	ColumnId   string           `json:"columnId"`
	Dataset    AssistantDataset `json:"dataset"`
	DatasetId  string           `json:"datasetId"`
	Filterable bool             `json:"filterable"`
	Searchable bool             `json:"searchable"`
}

type AssistantDatasetIndexStats struct {
	TotalRowsIndexed int64 `json:"totalRowsIndexed"`
	TotalRowsQueried int64 `json:"totalRowsQueried"`
}

type AssistantDatasetResults struct {
	Results      []JSON `json:"results"`
	TotalResults int64  `json:"totalResults"`
}

type Chat struct {
	ClientId    *string        `json:"clientId"`
	CreatedAt   DateTime       `json:"createdAt"`
	Id          string         `json:"id"`
	Messages    []*ChatMessage `json:"messages"`
	Parent      *ChatParent    `json:"parent"`
	ParentId    *string        `json:"parentId"`
	ParentType  *string        `json:"parentType"`
	UpdatedAt   DateTime       `json:"updatedAt"`
	User        *User          `json:"user"`
	UserId      *string        `json:"userId"`
	Workspace   Workspace      `json:"workspace"`
	WorkspaceId string         `json:"workspaceId"`
}

type ChatMessage struct {
	Chat        Chat           `json:"chat"`
	ChatId      string         `json:"chatId"`
	Children    []*ChatMessage `json:"children"`
	Content     string         `json:"content"`
	ContentId   *string        `json:"contentId"`
	ContentType *string        `json:"contentType"`
	CreatedAt   DateTime       `json:"createdAt"`
	Id          string         `json:"id"`
	Name        *string        `json:"name"`
	Parent      *ChatMessage   `json:"parent"`
	ParentId    *string        `json:"parentId"`
	Role        string         `json:"role"`
	TokenCount  int64          `json:"tokenCount"`
	UpdatedAt   DateTime       `json:"updatedAt"`
	User        *User          `json:"user"`
	UserId      *string        `json:"userId"`
}

type ChatMessageFilterInput struct {
	After       *DateTime `json:"after"`
	ContentId   *string   `json:"contentId"`
	ContentType *string   `json:"contentType"`
	Limit       *int64    `json:"limit"`
	Roles       []string  `json:"roles"`
}

type ChatParent struct {
	TypeName string `json:"__typeName"`
	Assistant
	TextQuery
}

type Column struct {
	CreatedAt  DateTime `json:"createdAt"`
	Id         string   `json:"id"`
	IsNullable bool     `json:"isNullable"`
	IsPrimary  bool     `json:"isPrimary"`
	Name       string   `json:"name"`
	OrdinalPos int64    `json:"ordinalPos"`
	Table      Table    `json:"table"`
	TableId    string   `json:"tableId"`
	Type       string   `json:"type"`
	UpdatedAt  DateTime `json:"updatedAt"`
}

type ColumnSearchResult struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Rank  float64 `json:"rank"`
	Table string  `json:"table"`
}

type ColumnSearchResults struct {
	Results []ColumnSearchResult `json:"results"`
}

type CopySchemaInput struct {
	WorkspaceId string `json:"workspaceId"`
}

type CreateAssistantDatasetColumnInput struct {
	ColumnId   string `json:"columnId"`
	Filterable bool   `json:"filterable"`
	Searchable bool   `json:"searchable"`
}

type CreateAssistantDatasetInput struct {
	AssistantId    string                              `json:"assistantId"`
	Columns        []CreateAssistantDatasetColumnInput `json:"columns"`
	IdColumnId     string                              `json:"idColumnId"`
	Name           string                              `json:"name"`
	ResultFunction *string                             `json:"resultFunction"`
	TableId        string                              `json:"tableId"`
}

type CreateAssistantInput struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	WorkspaceId string `json:"workspaceId"`
}

type CreateChatInput struct {
	ClientId      *string                   `json:"clientId"`
	IntroMessage  *OptionalChatMessageInput `json:"introMessage"`
	ParentId      *string                   `json:"parentId"`
	ParentType    *string                   `json:"parentType"`
	SystemMessage *OptionalChatMessageInput `json:"systemMessage"`
	WorkspaceId   string                    `json:"workspaceId"`
}

type CreateChatMessageInput struct {
	ChatId   string  `json:"chatId"`
	ClientId *string `json:"clientId"`
	Content  string  `json:"content"`
	Context  *string `json:"context"`
}

type CreateDestinationInput struct {
	Auth        *JSON  `json:"auth"`
	Config      *JSON  `json:"config"`
	TypeId      string `json:"typeId"`
	WorkspaceId string `json:"workspaceId"`
}

type CreateSchemaInput struct {
	Description string              `json:"description"`
	Name        string              `json:"name"`
	Namespace   EnumSchemaNamespace `json:"namespace"`
	NamespaceId string              `json:"namespaceId"`
	Title       string              `json:"title"`
	WorkspaceId string              `json:"workspaceId"`
}

type CreateSourceInput struct {
	Auth        *JSON  `json:"auth"`
	Config      *JSON  `json:"config"`
	TypeId      string `json:"typeId"`
	WorkspaceId string `json:"workspaceId"`
}

type CreateSqlQueryInput struct {
	ParentId    *string `json:"parentId"`
	Sql         *string `json:"sql"`
	Text        *string `json:"text"`
	WorkspaceId string  `json:"workspaceId"`
}

type CreateSyncInput struct {
	DestinationId     *string `json:"destinationId"`
	DestinationTypeId *string `json:"destinationTypeId"`
	SchedFreq         *int64  `json:"schedFreq"`
	SchedHour         *int64  `json:"schedHour"`
	SchedTz           *string `json:"schedTz"`
	SchedUtc          *string `json:"schedUtc"`
	SourceId          string  `json:"sourceId"`
	WorkspaceId       string  `json:"workspaceId"`
}

type CreateTextQueryInput struct {
	Columns     []TextQueryColumnInput `json:"columns"`
	SqlQueryId  *string                `json:"sqlQueryId"`
	Text        *string                `json:"text"`
	WorkspaceId string                 `json:"workspaceId"`
}

type CreateViewInput struct {
	Name        string  `json:"name"`
	SchemaId    *string `json:"schemaId"`
	SqlQueryId  string  `json:"sqlQueryId"`
	WorkspaceId string  `json:"workspaceId"`
}

type CreateVisualizationInput struct {
	Config      JSON    `json:"config"`
	SqlQueryId  string  `json:"sqlQueryId"`
	TextQueryId *string `json:"textQueryId"`
	Title       string  `json:"title"`
	Type        string  `json:"type"`
	WorkspaceId string  `json:"workspaceId"`
}

type CreateWorkspaceInput struct {
	Name string `json:"name"`
}

type Destination struct {
	CheckSuccess bool            `json:"checkSuccess"`
	Config       *JSON           `json:"config"`
	CreatedAt    *DateTime       `json:"createdAt"`
	Id           string          `json:"id"`
	Type         DestinationType `json:"type"`
	TypeId       string          `json:"typeId"`
	UpdatedAt    *DateTime       `json:"updatedAt"`
	Workspace    Workspace       `json:"workspace"`
	WorkspaceId  string          `json:"workspaceId"`
}

type DestinationType struct {
	Config       JSON           `json:"config"`
	Description  *string        `json:"description"`
	Destinations []*Destination `json:"destinations"`
	HasOauth     bool           `json:"hasOauth"`
	Id           string         `json:"id"`
	IsEnabled    bool           `json:"isEnabled"`
	Key          *string        `json:"key"`
	Logo         *string        `json:"logo"`
	Name         string         `json:"name"`
	Object       *string        `json:"object"`
	Provider     string         `json:"provider"`
	UiSchema     *JSON          `json:"uiSchema"`
}

type JSONSchema struct {
	Properties JSON   `json:"properties"`
	Type       string `json:"type"`
}

type OptionalChatMessageInput struct {
	Content *string `json:"content"`
	Context *string `json:"context"`
}

type PublicSchema struct {
	IsEnabled bool   `json:"isEnabled"`
	Schema    Schema `json:"schema"`
}

type Redwood struct {
	CurrentUser   *JSON   `json:"currentUser"`
	PrismaVersion *string `json:"prismaVersion"`
	Version       *string `json:"version"`
}

type Schema struct {
	Db          string              `json:"db"`
	Description string              `json:"description"`
	Id          string              `json:"id"`
	Name        string              `json:"name"`
	Namespace   EnumSchemaNamespace `json:"namespace"`
	NamespaceId string              `json:"namespaceId"`
	Tables      []*Table            `json:"tables"`
	Title       string              `json:"title"`
	Workspace   Workspace           `json:"workspace"`
	WorkspaceId string              `json:"workspaceId"`
}

type Source struct {
	Assets       []SourceAsset   `json:"assets"`
	Catalog      *AirbyteCatalog `json:"catalog"`
	CheckSuccess bool            `json:"checkSuccess"`
	Config       *JSON           `json:"config"`
	CreatedAt    *DateTime       `json:"createdAt"`
	DisplayName  *string         `json:"displayName"`
	Id           string          `json:"id"`
	Type         SourceType      `json:"type"`
	TypeId       string          `json:"typeId"`
	UpdatedAt    *DateTime       `json:"updatedAt"`
	Workspace    Workspace       `json:"workspace"`
	WorkspaceId  string          `json:"workspaceId"`
}

type SourceAsset struct {
	Data JSON   `json:"data"`
	Name string `json:"name"`
}

type SourceCatalog struct {
	Catalog AirbyteCatalog `json:"catalog"`
}

type SourceType struct {
	AuthProvider *string   `json:"authProvider"`
	Config       JSON      `json:"config"`
	Description  *string   `json:"description"`
	HasOauth     bool      `json:"hasOauth"`
	Id           string    `json:"id"`
	IsEnabled    bool      `json:"isEnabled"`
	Key          *string   `json:"key"`
	Logo         *string   `json:"logo"`
	Name         string    `json:"name"`
	Object       *string   `json:"object"`
	Provider     string    `json:"provider"`
	Sources      []*Source `json:"sources"`
	UiSchema     *JSON     `json:"uiSchema"`
}

type SqlQuery struct {
	Children       []SqlQuery       `json:"children"`
	Columns        []SqlQueryColumn `json:"columns"`
	CreatedAt      DateTime         `json:"createdAt"`
	Downvotes      int64            `json:"downvotes"`
	Fields         []SqlQueryField  `json:"fields"`
	Id             string           `json:"id"`
	Occurences     int64            `json:"occurences"`
	Sql            *string          `json:"sql"`
	Text           *string          `json:"text"`
	UpdatedAt      DateTime         `json:"updatedAt"`
	Upvotes        int64            `json:"upvotes"`
	View           *Table           `json:"view"`
	ViewId         *string          `json:"viewId"`
	Visualizations []Visualization  `json:"visualizations"`
}

type SqlQueryColumn struct {
	ColumnId   string `json:"columnId"`
	ColumnName string `json:"columnName"`
	Id         string `json:"id"`
	SchemaName string `json:"schemaName"`
	SqlQueryId string `json:"sqlQueryId"`
	TableName  string `json:"tableName"`
}

type SqlQueryField struct {
	Id         string `json:"id"`
	Mode       string `json:"mode"`
	Name       string `json:"name"`
	OrdinalPos int64  `json:"ordinalPos"`
	SqlQueryId string `json:"sqlQueryId"`
	Type       string `json:"type"`
}

type SqlQueryResult struct {
	Data    *SqlQueryResultData `json:"data"`
	Error   bool                `json:"error"`
	Message *string             `json:"message"`
	Meta    *SqlQueryResultMeta `json:"meta"`
}

type SqlQueryResultData struct {
	Rows []*JSON `json:"rows"`
}

type SqlQueryResultField struct {
	Mode *string `json:"mode"`
	Name *string `json:"name"`
	Type *string `json:"type"`
}

type SqlQueryResultMeta struct {
	Fields    []*SqlQueryResultField       `json:"fields"`
	NextQuery *SqlQueryResultMetaNextQuery `json:"nextQuery"`
	TotalRows *int64                       `json:"totalRows"`
}

type SqlQueryResultMetaNextQuery struct {
	JobId     *string `json:"jobId"`
	PageToken *string `json:"pageToken"`
}

type SqlQuerySearchResult struct {
	Rank     float64  `json:"rank"`
	SqlQuery SqlQuery `json:"sqlQuery"`
}

type SqlQuerySearchResults struct {
	Results []SqlQuerySearchResult `json:"results"`
}

type StripePlan struct {
	Description string        `json:"description"`
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	Prices      []StripePrice `json:"prices"`
}

type StripePrice struct {
	TrialPeriodDays *int64 `json:"trialPeriodDays"`
	UnitAmount      int64  `json:"unitAmount"`
	UnitInterval    string `json:"unitInterval"`
}

type Sync struct {
	CreatedAt     *DateTime      `json:"createdAt"`
	Destination   Destination    `json:"destination"`
	DestinationId string         `json:"destinationId"`
	Error         *string        `json:"error"`
	Id            string         `json:"id"`
	LastSyncStats *SyncStats     `json:"lastSyncStats"`
	SchedFreq     int64          `json:"schedFreq"`
	SchedHour     int64          `json:"schedHour"`
	SchedTz       string         `json:"schedTz"`
	SchedUtc      string         `json:"schedUtc"`
	Schema        Schema         `json:"schema"`
	SchemaId      string         `json:"schemaId"`
	Source        Source         `json:"source"`
	SourceId      string         `json:"sourceId"`
	StartedAt     *DateTime      `json:"startedAt"`
	State         *JSON          `json:"state"`
	Status        EnumSyncStatus `json:"status"`
	SyncedAt      *DateTime      `json:"syncedAt"`
	UpdatedAt     *DateTime      `json:"updatedAt"`
	Workspace     Workspace      `json:"workspace"`
	WorkspaceId   string         `json:"workspaceId"`
}

type SyncStats struct {
	Duration     float64  `json:"duration"`
	EndAt        DateTime `json:"endAt"`
	Id           string   `json:"id"`
	MaxCPU       float64  `json:"maxCPU"`
	MaxMem       float64  `json:"maxMem"`
	StartAt      DateTime `json:"startAt"`
	Sync         Sync     `json:"sync"`
	SyncId       string   `json:"syncId"`
	TotalBytes   BigInt   `json:"totalBytes"`
	TotalRecords BigInt   `json:"totalRecords"`
}

type Table struct {
	Columns   []*Column `json:"columns"`
	CreatedAt DateTime  `json:"createdAt"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Schema    Schema    `json:"schema"`
	SchemaId  string    `json:"schemaId"`
	SqlQuery  *SqlQuery `json:"sqlQuery"`
	UpdatedAt DateTime  `json:"updatedAt"`
}

type TableSearchResult struct {
	Rank  float64 `json:"rank"`
	Table Table   `json:"table"`
}

type TextQuery struct {
	Chats      []*Chat           `json:"chats"`
	Columns    []TextQueryColumn `json:"columns"`
	CreatedAt  *DateTime         `json:"createdAt"`
	Id         string            `json:"id"`
	Occurences int64             `json:"occurences"`
	SqlQueries []SqlQuery        `json:"sqlQueries"`
	Text       string            `json:"text"`
	UpdatedAt  *DateTime         `json:"updatedAt"`
}

type TextQueryColumn struct {
	ColumnId     string  `json:"columnId"`
	ColumnName   string  `json:"columnName"`
	Idx          *int64  `json:"idx"`
	SchemaPrefix *string `json:"schemaPrefix"`
	TableName    string  `json:"tableName"`
	TextQueryId  string  `json:"textQueryId"`
	Word         *string `json:"word"`
}

type TextQueryColumnInput struct {
	ColumnId string `json:"columnId"`
	Idx      int64  `json:"idx"`
	Word     string `json:"word"`
}

type TextQuerySearchResult struct {
	Id   string  `json:"id"`
	Rank float64 `json:"rank"`
	Text string  `json:"text"`
}

type TextQuerySearchResults struct {
	Results []TextQuerySearchResult `json:"results"`
}

type UpdateAssistantDatasetColumnInput struct {
	ColumnId   string `json:"columnId"`
	Filterable bool   `json:"filterable"`
	Searchable bool   `json:"searchable"`
}

type UpdateAssistantDatasetInput struct {
	Columns        []UpdateAssistantDatasetColumnInput `json:"columns"`
	IdColumnId     *string                             `json:"idColumnId"`
	Name           *string                             `json:"name"`
	ResultFunction *string                             `json:"resultFunction"`
	TableId        *string                             `json:"tableId"`
}

type UpdateAssistantInput struct {
	Description *string `json:"description"`
	Name        *string `json:"name"`
	WorkspaceId *string `json:"workspaceId"`
}

type UpdateDestinationInput struct {
	Auth        *JSON   `json:"auth"`
	Config      *JSON   `json:"config"`
	WorkspaceId *string `json:"workspaceId"`
}

type UpdateSourceInput struct {
	Auth        *JSON   `json:"auth"`
	Catalog     *JSON   `json:"catalog"`
	Config      *JSON   `json:"config"`
	WorkspaceId *string `json:"workspaceId"`
}

type UpdateSyncInput struct {
	SchedFreq *int64  `json:"schedFreq"`
	SchedHour *int64  `json:"schedHour"`
	SchedTz   *string `json:"schedTz"`
	SchedUtc  *string `json:"schedUtc"`
}

type UpdateUserInput struct {
	AcceptTerms *bool     `json:"acceptTerms"`
	Company     *string   `json:"company"`
	Country     *string   `json:"country"`
	Name        *string   `json:"name"`
	SignedInAt  *DateTime `json:"signedInAt"`
}

type UpdateViewInput struct {
	Name       *string `json:"name"`
	SqlQueryId *string `json:"sqlQueryId"`
}

type UpdateVisualizationInput struct {
	Config      JSON    `json:"config"`
	SqlQueryId  *string `json:"sqlQueryId"`
	TextQueryId *string `json:"textQueryId"`
	Title       string  `json:"title"`
	Type        *string `json:"type"`
	WorkspaceId *string `json:"workspaceId"`
}

type UpdateWorkspaceInput struct {
	Name *string `json:"name"`
}

type User struct {
	AcceptTerms bool      `json:"acceptTerms"`
	Company     *string   `json:"company"`
	Country     *string   `json:"country"`
	CreatedAt   DateTime  `json:"createdAt"`
	CustomerId  *string   `json:"customerId"`
	Email       string    `json:"email"`
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	SignedInAt  *DateTime `json:"signedInAt"`
	UpdatedAt   *DateTime `json:"updatedAt"`
}

type Visualization struct {
	Config      JSON       `json:"config"`
	CreatedAt   DateTime   `json:"createdAt"`
	Id          string     `json:"id"`
	SqlQuery    SqlQuery   `json:"sqlQuery"`
	SqlQueryId  string     `json:"sqlQueryId"`
	TextQuery   *TextQuery `json:"textQuery"`
	TextQueryId *string    `json:"textQueryId"`
	Title       string     `json:"title"`
	Type        string     `json:"type"`
	UpdatedAt   DateTime   `json:"updatedAt"`
	Workspace   Workspace  `json:"workspace"`
	WorkspaceId string     `json:"workspaceId"`
}

type Workspace struct {
	Assistants []Assistant     `json:"assistants"`
	Chats      []Chat          `json:"chats"`
	CreatedAt  *DateTime       `json:"createdAt"`
	Id         string          `json:"id"`
	Name       string          `json:"name"`
	Schemata   []Schema        `json:"schemata"`
	Syncs      []Sync          `json:"syncs"`
	UpdatedAt  *DateTime       `json:"updatedAt"`
	Users      []WorkspaceUser `json:"users"`
}

type WorkspaceUsage struct {
	Id            string  `json:"id"`
	QueryBytes    BigInt  `json:"queryBytes"`
	QueryDuration float64 `json:"queryDuration"`
	QueryRecords  BigInt  `json:"queryRecords"`
	SyncBytes     BigInt  `json:"syncBytes"`
	SyncDuration  float64 `json:"syncDuration"`
	SyncRecords   BigInt  `json:"syncRecords"`
}

type WorkspaceUser struct {
	Role string `json:"role"`
	User User   `json:"user"`
}
