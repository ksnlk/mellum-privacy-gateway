// limitations under the License.

package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
	"unicode"

	"github.com/nats-io/nuid"
)

// Request API subjects for JetStream.
const (
	// All API endpoints.
	jsAllAPI = "$JS.API.>"

	// For constructing JetStream domain prefixes.
	jsDomainAPI = "$JS.%s.API.>"

	JSApiPrefix = "$JS.API"

	// JSApiAccountInfo is for obtaining general information about JetStream for this account.
	// Will return JSON response.
	JSApiAccountInfo = "$JS.API.INFO"

	// JSApiTemplateCreate is the endpoint to create new stream templates.
	// Will return JSON response.
	JSApiTemplateCreate  = "$JS.API.STREAM.TEMPLATE.CREATE.*"
	JSApiTemplateCreateT = "$JS.API.STREAM.TEMPLATE.CREATE.%s"

	// JSApiTemplates is the endpoint to list all stream template names for this account.
	// Will return JSON response.
	JSApiTemplates = "$JS.API.STREAM.TEMPLATE.NAMES"

	// JSApiTemplateInfo is for obtaining general information about a named stream template.
	// Will return JSON response.
	JSApiTemplateInfo  = "$JS.API.STREAM.TEMPLATE.INFO.*"
	JSApiTemplateInfoT = "$JS.API.STREAM.TEMPLATE.INFO.%s"

	// JSApiTemplateDelete is the endpoint to delete stream templates.
	// Will return JSON response.
	JSApiTemplateDelete  = "$JS.API.STREAM.TEMPLATE.DELETE.*"
	JSApiTemplateDeleteT = "$JS.API.STREAM.TEMPLATE.DELETE.%s"

	// JSApiStreamCreate is the endpoint to create new streams.
	// JSApiRequestNextT is the prefix for the request next message(s) for a consumer in worker/pull mode.
	JSApiRequestNextT = "$JS.API.CONSUMER.MSG.NEXT.%s.%s"

	// For snapshots and restores. The ack will have additional tokens.
	jsSnapshotAckT    = "$JS.SNAPSHOT.ACK.%s.%s"
	jsRestoreDeliverT = "$JS.SNAPSHOT.RESTORE.%s.%s"

	// JSApiStreamRemovePeer is the endpoint to remove a peer from a clustered stream and its consumers.
	// Will return JSON response.
	JSApiStreamRemovePeer  = "$JS.API.STREAM.PEER.REMOVE.*"
	JSApiStreamRemovePeerT = "$JS.API.STREAM.PEER.REMOVE.%s"

	// JSApiStreamLeaderStepDown is the endpoint to have stream leader stepdown.
	// Will return JSON response.
	JSApiStreamLeaderStepDown  = "$JS.API.STREAM.LEADER.STEPDOWN.*"
	JSApiStreamLeaderStepDownT = "$JS.API.STREAM.LEADER.STEPDOWN.%s"

	// JSApiConsumerLeaderStepDown is the endpoint to have consumer leader stepdown.
	// Will return JSON response.
	JSApiConsumerLeaderStepDown  = "$JS.API.CONSUMER.LEADER.STEPDOWN.*.*"
	JSApiConsumerLeaderStepDownT = "$JS.API.CONSUMER.LEADER.STEPDOWN.%s.%s"

	// JSApiLeaderStepDown is the endpoint to have our metaleader stepdown.
	// Only works from system account.
	// Will return JSON response.
	JSApiLeaderStepDown = "$JS.API.META.LEADER.STEPDOWN"

	// JSApiRemoveServer is the endpoint to remove a peer server from the cluster.
	// Only works from system account.
	// Will return JSON response.
	JSApiRemoveServer = "$JS.API.SERVER.REMOVE"

	// jsAckT is the template for the ack message stream coming back from a consumer
	// when they ACK/NAK, etc a message.
	jsAckT   = "$JS.ACK.%s.%s"
	jsAckPre = "$JS.ACK."

	// jsFlowControl is for flow control subjects.
	jsFlowControlPre = "$JS.FC."
	// jsFlowControl is for FC responses.
	jsFlowControl = "$JS.FC.%s.%s.*"

	// JSAdvisoryPrefix is a prefix for all JetStream advisories.
	JSAdvisoryPrefix = "$JS.EVENT.ADVISORY"

	// JSMetricPrefix is a prefix for all JetStream metrics.
	JSMetricPrefix = "$JS.EVENT.METRIC"

	// JSMetricConsumerAckPre is a metric containing ack latency.
	JSMetricConsumerAckPre = "$JS.EVENT.METRIC.CONSUMER.ACK"

}

const JSApiStreamCreateResponseType = "io.nats.jetstream.api.v1.stream_create_response"

// JSApiStreamDeleteResponse stream removal.
type JSApiStreamDeleteResponse struct {
	ApiResponse
	Success bool `json:"success,omitempty"`
}

const JSApiStreamDeleteResponseType = "io.nats.jetstream.api.v1.stream_delete_response"

type JSApiStreamInfoRequest struct {
	DeletedDetails bool `json:"deleted_details,omitempty"`
}

type JSApiStreamInfoResponse struct {
	ApiResponse
	*StreamInfo
}

const JSApiStreamInfoResponseType = "io.nats.jetstream.api.v1.stream_info_response"

// JSApiNamesLimit is the maximum entries we will return for streams or consumers lists.
// TODO(dlc) - with header or request support could request chunked response.
const JSApiNamesLimit = 1024
const JSApiListLimit = 256

type JSApiStreamNamesRequest struct {
	ApiPagedRequest
	// These are filters that can be applied to the list.
	Subject string `json:"subject,omitempty"`
}

// JSApiStreamNamesResponse list of streams.
// A nil request is valid and means all streams.
type JSApiStreamNamesResponse struct {
	ApiResponse
	ApiPaged
	Streams []string `json:"streams"`
}

const JSApiStreamNamesResponseType = "io.nats.jetstream.api.v1.stream_names_response"

// JSApiStreamListResponse list of detailed stream information.
// A nil request is valid and means all streams.
type JSApiStreamListResponse struct {
	ApiResponse
	ApiPaged
	Streams []*StreamInfo `json:"streams"`
}