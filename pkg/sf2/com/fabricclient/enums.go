package fabricclient

// enum FABRIC_SERVICE_DESCRIPTION_KIND
var FABRIC_SERVICE_DESCRIPTION_KIND = struct {
	FABRIC_SERVICE_DESCRIPTION_KIND_INVALID   int32
	FABRIC_SERVICE_DESCRIPTION_KIND_STATELESS int32
	FABRIC_SERVICE_DESCRIPTION_KIND_STATEFUL  int32
}{
	FABRIC_SERVICE_DESCRIPTION_KIND_INVALID:   0,
	FABRIC_SERVICE_DESCRIPTION_KIND_STATELESS: 1,
	FABRIC_SERVICE_DESCRIPTION_KIND_STATEFUL:  2,
}

// enum FABRIC_PARTITION_KEY_TYPE
var FABRIC_PARTITION_KEY_TYPE = struct {
	FABRIC_PARTITION_KEY_TYPE_INVALID int32
	FABRIC_PARTITION_KEY_TYPE_NONE    int32
	FABRIC_PARTITION_KEY_TYPE_INT64   int32
	FABRIC_PARTITION_KEY_TYPE_STRING  int32
}{
	FABRIC_PARTITION_KEY_TYPE_INVALID: 0,
	FABRIC_PARTITION_KEY_TYPE_NONE:    1,
	FABRIC_PARTITION_KEY_TYPE_INT64:   2,
	FABRIC_PARTITION_KEY_TYPE_STRING:  3,
}

// enum FABRIC_SERVICE_PARTITION_KIND
var FABRIC_SERVICE_PARTITION_KIND = struct {
	FABRIC_SERVICE_PARTITION_KIND_INVALID     int32
	FABRIC_SERVICE_PARTITION_KIND_SINGLETON   int32
	FABRIC_SERVICE_PARTITION_KIND_INT64_RANGE int32
	FABRIC_SERVICE_PARTITION_KIND_NAMED       int32
}{
	FABRIC_SERVICE_PARTITION_KIND_INVALID:     0,
	FABRIC_SERVICE_PARTITION_KIND_SINGLETON:   1,
	FABRIC_SERVICE_PARTITION_KIND_INT64_RANGE: 2,
	FABRIC_SERVICE_PARTITION_KIND_NAMED:       3,
}

// enum FABRIC_SERVICE_ENDPOINT_ROLE
var FABRIC_SERVICE_ENDPOINT_ROLE = struct {
	FABRIC_SERVICE_ROLE_INVALID            int32
	FABRIC_SERVICE_ROLE_STATELESS          int32
	FABRIC_SERVICE_ROLE_STATEFUL_PRIMARY   int32
	FABRIC_SERVICE_ROLE_STATEFUL_SECONDARY int32
}{
	FABRIC_SERVICE_ROLE_INVALID:            0,
	FABRIC_SERVICE_ROLE_STATELESS:          1,
	FABRIC_SERVICE_ROLE_STATEFUL_PRIMARY:   2,
	FABRIC_SERVICE_ROLE_STATEFUL_SECONDARY: 3,
}

// enum FABRIC_CLAIMS_RETRIEVAL_METADATA_KIND
var FABRIC_CLAIMS_RETRIEVAL_METADATA_KIND = struct {
	FABRIC_CLAIMS_RETRIEVAL_METADATA_KIND_NONE int32
	FABRIC_CLAIMS_RETRIEVAL_METADATA_KIND_AAD  int32
}{
	FABRIC_CLAIMS_RETRIEVAL_METADATA_KIND_NONE: 0,
	FABRIC_CLAIMS_RETRIEVAL_METADATA_KIND_AAD:  1,
}

// enum FABRIC_SECURITY_CREDENTIAL_KIND
var FABRIC_SECURITY_CREDENTIAL_KIND = struct {
	FABRIC_SECURITY_CREDENTIAL_KIND_NONE    int32
	FABRIC_SECURITY_CREDENTIAL_KIND_X509    int32
	FABRIC_SECURITY_CREDENTIAL_KIND_WINDOWS int32
	FABRIC_SECURITY_CREDENTIAL_KIND_CLAIMS  int32
	FABRIC_SECURITY_CREDENTIAL_KIND_X509_2  int32
	FABRIC_SECURITY_CREDENTIAL_KIND_INVALID int32
}{
	FABRIC_SECURITY_CREDENTIAL_KIND_NONE:    0,
	FABRIC_SECURITY_CREDENTIAL_KIND_X509:    1,
	FABRIC_SECURITY_CREDENTIAL_KIND_WINDOWS: 2,
	FABRIC_SECURITY_CREDENTIAL_KIND_CLAIMS:  3,
	FABRIC_SECURITY_CREDENTIAL_KIND_X509_2:  4,
	FABRIC_SECURITY_CREDENTIAL_KIND_INVALID: 255,
}

// enum FABRIC_ENUMERATION_STATUS
var FABRIC_ENUMERATION_STATUS = struct {
	FABRIC_ENUMERATION_INVALID               int32
	FABRIC_ENUMERATION_BEST_EFFORT_MORE_DATA int32
	FABRIC_ENUMERATION_CONSISTENT_MORE_DATA  int32
	FABRIC_ENUMERATION_BEST_EFFORT_FINISHED  int32
	FABRIC_ENUMERATION_CONSISTENT_FINISHED   int32
	FABRIC_ENUMERATION_VALID_MASK            int32
	FABRIC_ENUMERATION_BEST_EFFORT_MASK      int32
	FABRIC_ENUMERATION_CONSISTENT_MASK       int32
	FABRIC_ENUMERATION_MORE_DATA_MASK        int32
	FABRIC_ENUMERATION_FINISHED_MASK         int32
}{
	FABRIC_ENUMERATION_INVALID:               0,
	FABRIC_ENUMERATION_BEST_EFFORT_MORE_DATA: 1,
	FABRIC_ENUMERATION_CONSISTENT_MORE_DATA:  2,
	FABRIC_ENUMERATION_BEST_EFFORT_FINISHED:  4,
	FABRIC_ENUMERATION_CONSISTENT_FINISHED:   8,
	FABRIC_ENUMERATION_VALID_MASK:            15,
	FABRIC_ENUMERATION_BEST_EFFORT_MASK:      5,
	FABRIC_ENUMERATION_CONSISTENT_MASK:       10,
	FABRIC_ENUMERATION_MORE_DATA_MASK:        3,
	FABRIC_ENUMERATION_FINISHED_MASK:         12,
}

// enum FABRIC_PROPERTY_TYPE_ID
var FABRIC_PROPERTY_TYPE_ID = struct {
	FABRIC_PROPERTY_TYPE_INVALID int32
	FABRIC_PROPERTY_TYPE_BINARY  int32
	FABRIC_PROPERTY_TYPE_INT64   int32
	FABRIC_PROPERTY_TYPE_DOUBLE  int32
	FABRIC_PROPERTY_TYPE_WSTRING int32
	FABRIC_PROPERTY_TYPE_GUID    int32
}{
	FABRIC_PROPERTY_TYPE_INVALID: 0,
	FABRIC_PROPERTY_TYPE_BINARY:  1,
	FABRIC_PROPERTY_TYPE_INT64:   2,
	FABRIC_PROPERTY_TYPE_DOUBLE:  3,
	FABRIC_PROPERTY_TYPE_WSTRING: 4,
	FABRIC_PROPERTY_TYPE_GUID:    5,
}

// enum FABRIC_PROPERTY_BATCH_OPERATION_KIND
var FABRIC_PROPERTY_BATCH_OPERATION_KIND = struct {
	FABRIC_PROPERTY_BATCH_OPERATION_KIND_INVALID        int32
	FABRIC_PROPERTY_BATCH_OPERATION_KIND_PUT            int32
	FABRIC_PROPERTY_BATCH_OPERATION_KIND_GET            int32
	FABRIC_PROPERTY_BATCH_OPERATION_KIND_CHECK_EXISTS   int32
	FABRIC_PROPERTY_BATCH_OPERATION_KIND_CHECK_SEQUENCE int32
	FABRIC_PROPERTY_BATCH_OPERATION_KIND_DELETE         int32
	FABRIC_PROPERTY_BATCH_OPERATION_KIND_PUT_CUSTOM     int32
	FABRIC_PROPERTY_BATCH_OPERATION_KIND_CHECK_VALUE    int32
}{
	FABRIC_PROPERTY_BATCH_OPERATION_KIND_INVALID:        0,
	FABRIC_PROPERTY_BATCH_OPERATION_KIND_PUT:            1,
	FABRIC_PROPERTY_BATCH_OPERATION_KIND_GET:            2,
	FABRIC_PROPERTY_BATCH_OPERATION_KIND_CHECK_EXISTS:   3,
	FABRIC_PROPERTY_BATCH_OPERATION_KIND_CHECK_SEQUENCE: 4,
	FABRIC_PROPERTY_BATCH_OPERATION_KIND_DELETE:         5,
	FABRIC_PROPERTY_BATCH_OPERATION_KIND_PUT_CUSTOM:     6,
	FABRIC_PROPERTY_BATCH_OPERATION_KIND_CHECK_VALUE:    7,
}

// enum FABRIC_SERVICE_NOTIFICATION_FILTER_FLAGS
var FABRIC_SERVICE_NOTIFICATION_FILTER_FLAGS = struct {
	FABRIC_SERVICE_NOTIFICATION_FILTER_FLAGS_NONE         int32
	FABRIC_SERVICE_NOTIFICATION_FILTER_FLAGS_NAME_PREFIX  int32
	FABRIC_SERVICE_NOTIFICATION_FILTER_FLAGS_PRIMARY_ONLY int32
}{
	FABRIC_SERVICE_NOTIFICATION_FILTER_FLAGS_NONE:         0,
	FABRIC_SERVICE_NOTIFICATION_FILTER_FLAGS_NAME_PREFIX:  1,
	FABRIC_SERVICE_NOTIFICATION_FILTER_FLAGS_PRIMARY_ONLY: 2,
}

// enum FABRIC_SERVICE_PACKAGE_ACTIVATION_MODE
var FABRIC_SERVICE_PACKAGE_ACTIVATION_MODE = struct {
	FABRIC_SERVICE_PACKAGE_ACTIVATION_MODE_SHARED_PROCESS    int32
	FABRIC_SERVICE_PACKAGE_ACTIVATION_MODE_EXCLUSIVE_PROCESS int32
}{
	FABRIC_SERVICE_PACKAGE_ACTIVATION_MODE_SHARED_PROCESS:    0,
	FABRIC_SERVICE_PACKAGE_ACTIVATION_MODE_EXCLUSIVE_PROCESS: 1,
}

// enum FABRIC_SERVICE_LOAD_METRIC_WEIGHT
var FABRIC_SERVICE_LOAD_METRIC_WEIGHT = struct {
	FABRIC_SERVICE_LOAD_METRIC_WEIGHT_ZERO   int32
	FABRIC_SERVICE_LOAD_METRIC_WEIGHT_LOW    int32
	FABRIC_SERVICE_LOAD_METRIC_WEIGHT_MEDIUM int32
	FABRIC_SERVICE_LOAD_METRIC_WEIGHT_HIGH   int32
}{
	FABRIC_SERVICE_LOAD_METRIC_WEIGHT_ZERO:   0,
	FABRIC_SERVICE_LOAD_METRIC_WEIGHT_LOW:    1,
	FABRIC_SERVICE_LOAD_METRIC_WEIGHT_MEDIUM: 2,
	FABRIC_SERVICE_LOAD_METRIC_WEIGHT_HIGH:   3,
}

// enum FABRIC_APPLICATION_UPGRADE_KIND
var FABRIC_APPLICATION_UPGRADE_KIND = struct {
	FABRIC_APPLICATION_UPGRADE_KIND_INVALID int32
	FABRIC_APPLICATION_UPGRADE_KIND_ROLLING int32
}{
	FABRIC_APPLICATION_UPGRADE_KIND_INVALID: 0,
	FABRIC_APPLICATION_UPGRADE_KIND_ROLLING: 1,
}

// enum FABRIC_APPLICATION_UPGRADE_STATE
var FABRIC_APPLICATION_UPGRADE_STATE = struct {
	FABRIC_APPLICATION_UPGRADE_STATE_INVALID                     int32
	FABRIC_APPLICATION_UPGRADE_STATE_ROLLING_BACK_IN_PROGRESS    int32
	FABRIC_APPLICATION_UPGRADE_STATE_ROLLING_BACK_COMPLETED      int32
	FABRIC_APPLICATION_UPGRADE_STATE_ROLLING_FORWARD_PENDING     int32
	FABRIC_APPLICATION_UPGRADE_STATE_ROLLING_FORWARD_IN_PROGRESS int32
	FABRIC_APPLICATION_UPGRADE_STATE_ROLLING_FORWARD_COMPLETED   int32
	FABRIC_APPLICATION_UPGRADE_STATE_FAILED                      int32
	FABRIC_APPLICATION_UPGRADE_STATE_ROLLING_BACK_PENDING        int32
}{
	FABRIC_APPLICATION_UPGRADE_STATE_INVALID:                     0,
	FABRIC_APPLICATION_UPGRADE_STATE_ROLLING_BACK_IN_PROGRESS:    1,
	FABRIC_APPLICATION_UPGRADE_STATE_ROLLING_BACK_COMPLETED:      2,
	FABRIC_APPLICATION_UPGRADE_STATE_ROLLING_FORWARD_PENDING:     3,
	FABRIC_APPLICATION_UPGRADE_STATE_ROLLING_FORWARD_IN_PROGRESS: 4,
	FABRIC_APPLICATION_UPGRADE_STATE_ROLLING_FORWARD_COMPLETED:   5,
	FABRIC_APPLICATION_UPGRADE_STATE_FAILED:                      6,
	FABRIC_APPLICATION_UPGRADE_STATE_ROLLING_BACK_PENDING:        7,
}

// enum FABRIC_UPGRADE_DOMAIN_STATE
var FABRIC_UPGRADE_DOMAIN_STATE = struct {
	FABRIC_UPGRADE_DOMAIN_STATE_INVALID     int32
	FABRIC_UPGRADE_DOMAIN_STATE_PENDING     int32
	FABRIC_UPGRADE_DOMAIN_STATE_IN_PROGRESS int32
	FABRIC_UPGRADE_DOMAIN_STATE_COMPLETED   int32
}{
	FABRIC_UPGRADE_DOMAIN_STATE_INVALID:     0,
	FABRIC_UPGRADE_DOMAIN_STATE_PENDING:     1,
	FABRIC_UPGRADE_DOMAIN_STATE_IN_PROGRESS: 2,
	FABRIC_UPGRADE_DOMAIN_STATE_COMPLETED:   3,
}

// enum FABRIC_ROLLING_UPGRADE_MODE
var FABRIC_ROLLING_UPGRADE_MODE = struct {
	FABRIC_ROLLING_UPGRADE_MODE_INVALID            int32
	FABRIC_ROLLING_UPGRADE_MODE_UNMONITORED_AUTO   int32
	FABRIC_ROLLING_UPGRADE_MODE_UNMONITORED_MANUAL int32
	FABRIC_ROLLING_UPGRADE_MODE_MONITORED          int32
}{
	FABRIC_ROLLING_UPGRADE_MODE_INVALID:            0,
	FABRIC_ROLLING_UPGRADE_MODE_UNMONITORED_AUTO:   1,
	FABRIC_ROLLING_UPGRADE_MODE_UNMONITORED_MANUAL: 2,
	FABRIC_ROLLING_UPGRADE_MODE_MONITORED:          3,
}

// enum FABRIC_PACKAGE_SHARING_POLICY_SCOPE
var FABRIC_PACKAGE_SHARING_POLICY_SCOPE = struct {
	FABRIC_PACKAGE_SHARING_POLICY_SCOPE_NONE   int32
	FABRIC_PACKAGE_SHARING_POLICY_SCOPE_ALL    int32
	FABRIC_PACKAGE_SHARING_POLICY_SCOPE_CODE   int32
	FABRIC_PACKAGE_SHARING_POLICY_SCOPE_CONFIG int32
	FABRIC_PACKAGE_SHARING_POLICY_SCOPE_DATA   int32
}{
	FABRIC_PACKAGE_SHARING_POLICY_SCOPE_NONE:   0,
	FABRIC_PACKAGE_SHARING_POLICY_SCOPE_ALL:    1,
	FABRIC_PACKAGE_SHARING_POLICY_SCOPE_CODE:   2,
	FABRIC_PACKAGE_SHARING_POLICY_SCOPE_CONFIG: 3,
	FABRIC_PACKAGE_SHARING_POLICY_SCOPE_DATA:   4,
}

// enum FABRIC_PROVISION_APPLICATION_TYPE_KIND
var FABRIC_PROVISION_APPLICATION_TYPE_KIND = struct {
	FABRIC_PROVISION_APPLICATION_TYPE_KIND_INVALID          int32
	FABRIC_PROVISION_APPLICATION_TYPE_KIND_IMAGE_STORE_PATH int32
	FABRIC_PROVISION_APPLICATION_TYPE_KIND_EXTERNAL_STORE   int32
}{
	FABRIC_PROVISION_APPLICATION_TYPE_KIND_INVALID:          0,
	FABRIC_PROVISION_APPLICATION_TYPE_KIND_IMAGE_STORE_PATH: 1,
	FABRIC_PROVISION_APPLICATION_TYPE_KIND_EXTERNAL_STORE:   2,
}

// enum FABRIC_NODE_DEACTIVATION_INTENT
var FABRIC_NODE_DEACTIVATION_INTENT = struct {
	FABRIC_NODE_DEACTIVATION_INTENT_INVALID     int32
	FABRIC_NODE_DEACTIVATION_INTENT_PAUSE       int32
	FABRIC_NODE_DEACTIVATION_INTENT_RESTART     int32
	FABRIC_NODE_DEACTIVATION_INTENT_REMOVE_DATA int32
	FABRIC_NODE_DEACTIVATION_INTENT_REMOVE_NODE int32
}{
	FABRIC_NODE_DEACTIVATION_INTENT_INVALID:     0,
	FABRIC_NODE_DEACTIVATION_INTENT_PAUSE:       1,
	FABRIC_NODE_DEACTIVATION_INTENT_RESTART:     2,
	FABRIC_NODE_DEACTIVATION_INTENT_REMOVE_DATA: 3,
	FABRIC_NODE_DEACTIVATION_INTENT_REMOVE_NODE: 4,
}

// enum FABRIC_UPGRADE_KIND
var FABRIC_UPGRADE_KIND = struct {
	FABRIC_UPGRADE_KIND_INVALID int32
	FABRIC_UPGRADE_KIND_ROLLING int32
}{
	FABRIC_UPGRADE_KIND_INVALID: 0,
	FABRIC_UPGRADE_KIND_ROLLING: 1,
}

// enum FABRIC_UPGRADE_STATE
var FABRIC_UPGRADE_STATE = struct {
	FABRIC_UPGRADE_STATE_INVALID                     int32
	FABRIC_UPGRADE_STATE_ROLLING_BACK_IN_PROGRESS    int32
	FABRIC_UPGRADE_STATE_ROLLING_BACK_COMPLETED      int32
	FABRIC_UPGRADE_STATE_ROLLING_FORWARD_PENDING     int32
	FABRIC_UPGRADE_STATE_ROLLING_FORWARD_IN_PROGRESS int32
	FABRIC_UPGRADE_STATE_ROLLING_FORWARD_COMPLETED   int32
	FABRIC_UPGRADE_STATE_FAILED                      int32
	FABRIC_UPGRADE_STATE_ROLLING_BACK_PENDING        int32
}{
	FABRIC_UPGRADE_STATE_INVALID:                     0,
	FABRIC_UPGRADE_STATE_ROLLING_BACK_IN_PROGRESS:    1,
	FABRIC_UPGRADE_STATE_ROLLING_BACK_COMPLETED:      2,
	FABRIC_UPGRADE_STATE_ROLLING_FORWARD_PENDING:     3,
	FABRIC_UPGRADE_STATE_ROLLING_FORWARD_IN_PROGRESS: 4,
	FABRIC_UPGRADE_STATE_ROLLING_FORWARD_COMPLETED:   5,
	FABRIC_UPGRADE_STATE_FAILED:                      6,
	FABRIC_UPGRADE_STATE_ROLLING_BACK_PENDING:        7,
}

// enum FABRIC_HEALTH_REPORT_KIND
var FABRIC_HEALTH_REPORT_KIND = struct {
	FABRIC_HEALTH_REPORT_KIND_INVALID                    int32
	FABRIC_HEALTH_REPORT_KIND_STATEFUL_SERVICE_REPLICA   int32
	FABRIC_HEALTH_REPORT_KIND_STATELESS_SERVICE_INSTANCE int32
	FABRIC_HEALTH_REPORT_KIND_PARTITION                  int32
	FABRIC_HEALTH_REPORT_KIND_NODE                       int32
	FABRIC_HEALTH_REPORT_KIND_SERVICE                    int32
	FABRIC_HEALTH_REPORT_KIND_APPLICATION                int32
	FABRIC_HEALTH_REPORT_KIND_DEPLOYED_APPLICATION       int32
	FABRIC_HEALTH_REPORT_KIND_DEPLOYED_SERVICE_PACKAGE   int32
	FABRIC_HEALTH_REPORT_KIND_CLUSTER                    int32
}{
	FABRIC_HEALTH_REPORT_KIND_INVALID:                    0,
	FABRIC_HEALTH_REPORT_KIND_STATEFUL_SERVICE_REPLICA:   1,
	FABRIC_HEALTH_REPORT_KIND_STATELESS_SERVICE_INSTANCE: 2,
	FABRIC_HEALTH_REPORT_KIND_PARTITION:                  3,
	FABRIC_HEALTH_REPORT_KIND_NODE:                       4,
	FABRIC_HEALTH_REPORT_KIND_SERVICE:                    5,
	FABRIC_HEALTH_REPORT_KIND_APPLICATION:                6,
	FABRIC_HEALTH_REPORT_KIND_DEPLOYED_APPLICATION:       7,
	FABRIC_HEALTH_REPORT_KIND_DEPLOYED_SERVICE_PACKAGE:   8,
	FABRIC_HEALTH_REPORT_KIND_CLUSTER:                    9,
}

// enum FABRIC_HEALTH_STATE
var FABRIC_HEALTH_STATE = struct {
	FABRIC_HEALTH_STATE_INVALID int32
	FABRIC_HEALTH_STATE_OK      int32
	FABRIC_HEALTH_STATE_WARNING int32
	FABRIC_HEALTH_STATE_ERROR   int32
	FABRIC_HEALTH_STATE_UNKNOWN int32
}{
	FABRIC_HEALTH_STATE_INVALID: 0,
	FABRIC_HEALTH_STATE_OK:      1,
	FABRIC_HEALTH_STATE_WARNING: 2,
	FABRIC_HEALTH_STATE_ERROR:   3,
	FABRIC_HEALTH_STATE_UNKNOWN: 65535,
}

// enum FABRIC_SERVICE_KIND
var FABRIC_SERVICE_KIND = struct {
	FABRIC_SERVICE_KIND_INVALID   int32
	FABRIC_SERVICE_KIND_STATELESS int32
	FABRIC_SERVICE_KIND_STATEFUL  int32
}{
	FABRIC_SERVICE_KIND_INVALID:   0,
	FABRIC_SERVICE_KIND_STATELESS: 1,
	FABRIC_SERVICE_KIND_STATEFUL:  2,
}

// enum FABRIC_QUERY_NODE_STATUS
var FABRIC_QUERY_NODE_STATUS = struct {
	FABRIC_QUERY_NODE_STATUS_INVALID   int32
	FABRIC_QUERY_NODE_STATUS_UP        int32
	FABRIC_QUERY_NODE_STATUS_DOWN      int32
	FABRIC_QUERY_NODE_STATUS_ENABLING  int32
	FABRIC_QUERY_NODE_STATUS_DISABLING int32
	FABRIC_QUERY_NODE_STATUS_DISABLED  int32
	FABRIC_QUERY_NODE_STATUS_UNKNOWN   int32
	FABRIC_QUERY_NODE_STATUS_REMOVED   int32
}{
	FABRIC_QUERY_NODE_STATUS_INVALID:   0,
	FABRIC_QUERY_NODE_STATUS_UP:        1,
	FABRIC_QUERY_NODE_STATUS_DOWN:      2,
	FABRIC_QUERY_NODE_STATUS_ENABLING:  3,
	FABRIC_QUERY_NODE_STATUS_DISABLING: 4,
	FABRIC_QUERY_NODE_STATUS_DISABLED:  5,
	FABRIC_QUERY_NODE_STATUS_UNKNOWN:   6,
	FABRIC_QUERY_NODE_STATUS_REMOVED:   7,
}

// enum FABRIC_APPLICATION_STATUS
var FABRIC_APPLICATION_STATUS = struct {
	FABRIC_APPLICATION_STATUS_INVALID   int32
	FABRIC_APPLICATION_STATUS_READY     int32
	FABRIC_APPLICATION_STATUS_UPGRADING int32
	FABRIC_APPLICATION_STATUS_CREATING  int32
	FABRIC_APPLICATION_STATUS_DELETING  int32
	FABRIC_APPLICATION_STATUS_FAILED    int32
}{
	FABRIC_APPLICATION_STATUS_INVALID:   0,
	FABRIC_APPLICATION_STATUS_READY:     1,
	FABRIC_APPLICATION_STATUS_UPGRADING: 2,
	FABRIC_APPLICATION_STATUS_CREATING:  3,
	FABRIC_APPLICATION_STATUS_DELETING:  4,
	FABRIC_APPLICATION_STATUS_FAILED:    5,
}

// enum FABRIC_DEPLOYMENT_STATUS
var FABRIC_DEPLOYMENT_STATUS = struct {
	FABRIC_DEPLOYMENT_STATUS_INVALID      int32
	FABRIC_DEPLOYMENT_STATUS_DOWNLOADING  int32
	FABRIC_DEPLOYMENT_STATUS_ACTIVATING   int32
	FABRIC_DEPLOYMENT_STATUS_ACTIVE       int32
	FABRIC_DEPLOYMENT_STATUS_UPGRADING    int32
	FABRIC_DEPLOYMENT_STATUS_DEACTIVATING int32
}{
	FABRIC_DEPLOYMENT_STATUS_INVALID:      0,
	FABRIC_DEPLOYMENT_STATUS_DOWNLOADING:  1,
	FABRIC_DEPLOYMENT_STATUS_ACTIVATING:   2,
	FABRIC_DEPLOYMENT_STATUS_ACTIVE:       3,
	FABRIC_DEPLOYMENT_STATUS_UPGRADING:    4,
	FABRIC_DEPLOYMENT_STATUS_DEACTIVATING: 5,
}

// enum FABRIC_SERVICE_TYPE_REGISTRATION_STATUS
var FABRIC_SERVICE_TYPE_REGISTRATION_STATUS = struct {
	FABRIC_SERVICE_TYPE_REGISTRATION_STATUS_INVALID        int32
	FABRIC_SERVICE_TYPE_REGISTRATION_STATUS_DISABLED       int32
	FABRIC_SERVICE_TYPE_REGISTRATION_STATUS_NOT_REGISTERED int32
	FABRIC_SERVICE_TYPE_REGISTRATION_STATUS_REGISTERED     int32
}{
	FABRIC_SERVICE_TYPE_REGISTRATION_STATUS_INVALID:        0,
	FABRIC_SERVICE_TYPE_REGISTRATION_STATUS_DISABLED:       1,
	FABRIC_SERVICE_TYPE_REGISTRATION_STATUS_NOT_REGISTERED: 2,
	FABRIC_SERVICE_TYPE_REGISTRATION_STATUS_REGISTERED:     3,
}

// enum FABRIC_ENTRY_POINT_STATUS
var FABRIC_ENTRY_POINT_STATUS = struct {
	FABRIC_ENTRY_POINT_STATUS_INVALID  int32
	FABRIC_ENTRY_POINT_STATUS_PENDING  int32
	FABRIC_ENTRY_POINT_STATUS_STARTING int32
	FABRIC_ENTRY_POINT_STATUS_STARTED  int32
	FABRIC_ENTRY_POINT_STATUS_STOPPING int32
	FABRIC_ENTRY_POINT_STATUS_STOPPED  int32
}{
	FABRIC_ENTRY_POINT_STATUS_INVALID:  0,
	FABRIC_ENTRY_POINT_STATUS_PENDING:  1,
	FABRIC_ENTRY_POINT_STATUS_STARTING: 2,
	FABRIC_ENTRY_POINT_STATUS_STARTED:  3,
	FABRIC_ENTRY_POINT_STATUS_STOPPING: 4,
	FABRIC_ENTRY_POINT_STATUS_STOPPED:  5,
}

// enum FABRIC_REPAIR_SCOPE_IDENTIFIER_KIND
var FABRIC_REPAIR_SCOPE_IDENTIFIER_KIND = struct {
	FABRIC_REPAIR_SCOPE_IDENTIFIER_KIND_INVALID int32
	FABRIC_REPAIR_SCOPE_IDENTIFIER_KIND_CLUSTER int32
}{
	FABRIC_REPAIR_SCOPE_IDENTIFIER_KIND_INVALID: 0,
	FABRIC_REPAIR_SCOPE_IDENTIFIER_KIND_CLUSTER: 1,
}

// enum FABRIC_REPAIR_TASK_STATE
var FABRIC_REPAIR_TASK_STATE = struct {
	FABRIC_REPAIR_TASK_STATE_INVALID   int32
	FABRIC_REPAIR_TASK_STATE_CREATED   int32
	FABRIC_REPAIR_TASK_STATE_CLAIMED   int32
	FABRIC_REPAIR_TASK_STATE_PREPARING int32
	FABRIC_REPAIR_TASK_STATE_APPROVED  int32
	FABRIC_REPAIR_TASK_STATE_EXECUTING int32
	FABRIC_REPAIR_TASK_STATE_RESTORING int32
	FABRIC_REPAIR_TASK_STATE_COMPLETED int32
}{
	FABRIC_REPAIR_TASK_STATE_INVALID:   0,
	FABRIC_REPAIR_TASK_STATE_CREATED:   1,
	FABRIC_REPAIR_TASK_STATE_CLAIMED:   2,
	FABRIC_REPAIR_TASK_STATE_PREPARING: 4,
	FABRIC_REPAIR_TASK_STATE_APPROVED:  8,
	FABRIC_REPAIR_TASK_STATE_EXECUTING: 16,
	FABRIC_REPAIR_TASK_STATE_RESTORING: 32,
	FABRIC_REPAIR_TASK_STATE_COMPLETED: 64,
}

// enum FABRIC_REPAIR_TARGET_KIND
var FABRIC_REPAIR_TARGET_KIND = struct {
	FABRIC_REPAIR_TARGET_KIND_INVALID int32
	FABRIC_REPAIR_TARGET_KIND_NODE    int32
}{
	FABRIC_REPAIR_TARGET_KIND_INVALID: 0,
	FABRIC_REPAIR_TARGET_KIND_NODE:    1,
}

// enum FABRIC_REPAIR_IMPACT_KIND
var FABRIC_REPAIR_IMPACT_KIND = struct {
	FABRIC_REPAIR_IMPACT_KIND_INVALID int32
	FABRIC_REPAIR_IMPACT_KIND_NODE    int32
}{
	FABRIC_REPAIR_IMPACT_KIND_INVALID: 0,
	FABRIC_REPAIR_IMPACT_KIND_NODE:    1,
}

// enum FABRIC_REPAIR_TASK_RESULT
var FABRIC_REPAIR_TASK_RESULT = struct {
	FABRIC_REPAIR_TASK_RESULT_INVALID     int32
	FABRIC_REPAIR_TASK_RESULT_SUCCEEDED   int32
	FABRIC_REPAIR_TASK_RESULT_CANCELLED   int32
	FABRIC_REPAIR_TASK_RESULT_INTERRUPTED int32
	FABRIC_REPAIR_TASK_RESULT_FAILED      int32
	FABRIC_REPAIR_TASK_RESULT_PENDING     int32
}{
	FABRIC_REPAIR_TASK_RESULT_INVALID:     0,
	FABRIC_REPAIR_TASK_RESULT_SUCCEEDED:   1,
	FABRIC_REPAIR_TASK_RESULT_CANCELLED:   2,
	FABRIC_REPAIR_TASK_RESULT_INTERRUPTED: 4,
	FABRIC_REPAIR_TASK_RESULT_FAILED:      8,
	FABRIC_REPAIR_TASK_RESULT_PENDING:     16,
}

// enum FABRIC_PARTITION_SELECTOR_TYPE
var FABRIC_PARTITION_SELECTOR_TYPE = struct {
	FABRIC_PARTITION_SELECTOR_TYPE_NONE          int32
	FABRIC_PARTITION_SELECTOR_TYPE_SINGLETON     int32
	FABRIC_PARTITION_SELECTOR_TYPE_NAMED         int32
	FABRIC_PARTITION_SELECTOR_TYPE_UNIFORM_INT64 int32
	FABRIC_PARTITION_SELECTOR_TYPE_PARTITION_ID  int32
	FABRIC_PARTITION_SELECTOR_TYPE_RANDOM        int32
}{
	FABRIC_PARTITION_SELECTOR_TYPE_NONE:          0,
	FABRIC_PARTITION_SELECTOR_TYPE_SINGLETON:     1,
	FABRIC_PARTITION_SELECTOR_TYPE_NAMED:         2,
	FABRIC_PARTITION_SELECTOR_TYPE_UNIFORM_INT64: 3,
	FABRIC_PARTITION_SELECTOR_TYPE_PARTITION_ID:  4,
	FABRIC_PARTITION_SELECTOR_TYPE_RANDOM:        5,
}

// enum FABRIC_DATA_LOSS_MODE
var FABRIC_DATA_LOSS_MODE = struct {
	FABRIC_DATA_LOSS_MODE_INVALID int32
	FABRIC_DATA_LOSS_MODE_PARTIAL int32
	FABRIC_DATA_LOSS_MODE_FULL    int32
}{
	FABRIC_DATA_LOSS_MODE_INVALID: 0,
	FABRIC_DATA_LOSS_MODE_PARTIAL: 1,
	FABRIC_DATA_LOSS_MODE_FULL:    2,
}

// enum FABRIC_TEST_COMMAND_PROGRESS_STATE
var FABRIC_TEST_COMMAND_PROGRESS_STATE = struct {
	FABRIC_TEST_COMMAND_PROGRESS_STATE_INVALID         int32
	FABRIC_TEST_COMMAND_PROGRESS_STATE_RUNNING         int32
	FABRIC_TEST_COMMAND_PROGRESS_STATE_ROLLING_BACK    int32
	FABRIC_TEST_COMMAND_PROGRESS_STATE_COMPLETED       int32
	FABRIC_TEST_COMMAND_PROGRESS_STATE_FAULTED         int32
	FABRIC_TEST_COMMAND_PROGRESS_STATE_CANCELLED       int32
	FABRIC_TEST_COMMAND_PROGRESS_STATE_FORCE_CANCELLED int32
}{
	FABRIC_TEST_COMMAND_PROGRESS_STATE_INVALID:         0,
	FABRIC_TEST_COMMAND_PROGRESS_STATE_RUNNING:         1,
	FABRIC_TEST_COMMAND_PROGRESS_STATE_ROLLING_BACK:    2,
	FABRIC_TEST_COMMAND_PROGRESS_STATE_COMPLETED:       3,
	FABRIC_TEST_COMMAND_PROGRESS_STATE_FAULTED:         4,
	FABRIC_TEST_COMMAND_PROGRESS_STATE_CANCELLED:       5,
	FABRIC_TEST_COMMAND_PROGRESS_STATE_FORCE_CANCELLED: 6,
}

// enum FABRIC_QUORUM_LOSS_MODE
var FABRIC_QUORUM_LOSS_MODE = struct {
	FABRIC_QUORUM_LOSS_MODE_INVALID         int32
	FABRIC_QUORUM_LOSS_MODE_QUORUM_REPLICAS int32
	FABRIC_QUORUM_LOSS_MODE_ALL_REPLICAS    int32
}{
	FABRIC_QUORUM_LOSS_MODE_INVALID:         0,
	FABRIC_QUORUM_LOSS_MODE_QUORUM_REPLICAS: 1,
	FABRIC_QUORUM_LOSS_MODE_ALL_REPLICAS:    2,
}

// enum FABRIC_RESTART_PARTITION_MODE
var FABRIC_RESTART_PARTITION_MODE = struct {
	FABRIC_RESTART_PARTITION_MODE_INVALID                   int32
	FABRIC_RESTART_PARTITION_MODE_ALL_REPLICAS_OR_INSTANCES int32
	FABRIC_RESTART_PARTITION_MODE_ONLY_ACTIVE_SECONDARIES   int32
}{
	FABRIC_RESTART_PARTITION_MODE_INVALID:                   0,
	FABRIC_RESTART_PARTITION_MODE_ALL_REPLICAS_OR_INSTANCES: 1,
	FABRIC_RESTART_PARTITION_MODE_ONLY_ACTIVE_SECONDARIES:   2,
}

// enum FABRIC_TEST_COMMAND_STATE_FILTER
var FABRIC_TEST_COMMAND_STATE_FILTER = struct {
	FABRIC_TEST_COMMAND_STATE_FILTER_DEFAULT                int32
	FABRIC_TEST_COMMAND_STATE_FILTER_ALL                    int32
	FABRIC_TEST_COMMAND_STATE_FILTER_RUNNING                int32
	FABRIC_TEST_COMMAND_STATE_FILTER_ROLLING_BACK           int32
	FABRIC_TEST_COMMAND_STATE_FILTER_COMPLETED_SUCCESSFULLY int32
	FABRIC_TEST_COMMAND_STATE_FILTER_FAILED                 int32
	FABRIC_TEST_COMMAND_STATE_FILTER_CANCELLED              int32
	FABRIC_TEST_COMMAND_STATE_FILTER_FORCE_CANCELLED        int32
}{
	FABRIC_TEST_COMMAND_STATE_FILTER_DEFAULT:                0,
	FABRIC_TEST_COMMAND_STATE_FILTER_ALL:                    65535,
	FABRIC_TEST_COMMAND_STATE_FILTER_RUNNING:                1,
	FABRIC_TEST_COMMAND_STATE_FILTER_ROLLING_BACK:           2,
	FABRIC_TEST_COMMAND_STATE_FILTER_COMPLETED_SUCCESSFULLY: 8,
	FABRIC_TEST_COMMAND_STATE_FILTER_FAILED:                 16,
	FABRIC_TEST_COMMAND_STATE_FILTER_CANCELLED:              32,
	FABRIC_TEST_COMMAND_STATE_FILTER_FORCE_CANCELLED:        64,
}

// enum FABRIC_TEST_COMMAND_TYPE_FILTER
var FABRIC_TEST_COMMAND_TYPE_FILTER = struct {
	FABRIC_TEST_COMMAND_TYPE_FILTER_DEFAULT               int32
	FABRIC_TEST_COMMAND_TYPE_FILTER_ALL                   int32
	FABRIC_TEST_COMMAND_TYPE_FILTER_PARTITION_DATA_LOSS   int32
	FABRIC_TEST_COMMAND_TYPE_FILTER_PARTITION_QUORUM_LOSS int32
	FABRIC_TEST_COMMAND_TYPE_FILTER_PARTITION_RESTART     int32
}{
	FABRIC_TEST_COMMAND_TYPE_FILTER_DEFAULT:               0,
	FABRIC_TEST_COMMAND_TYPE_FILTER_ALL:                   65535,
	FABRIC_TEST_COMMAND_TYPE_FILTER_PARTITION_DATA_LOSS:   1,
	FABRIC_TEST_COMMAND_TYPE_FILTER_PARTITION_QUORUM_LOSS: 2,
	FABRIC_TEST_COMMAND_TYPE_FILTER_PARTITION_RESTART:     4,
}

// enum FABRIC_CHAOS_STATUS
var FABRIC_CHAOS_STATUS = struct {
	FABRIC_CHAOS_STATUS_INVALID int32
	FABRIC_CHAOS_STATUS_RUNNING int32
	FABRIC_CHAOS_STATUS_STOPPED int32
}{
	FABRIC_CHAOS_STATUS_INVALID: 0,
	FABRIC_CHAOS_STATUS_RUNNING: 1,
	FABRIC_CHAOS_STATUS_STOPPED: 2,
}

// enum FABRIC_CHAOS_EVENT_KIND
var FABRIC_CHAOS_EVENT_KIND = struct {
	FABRIC_CHAOS_EVENT_KIND_INVALID           int32
	FABRIC_CHAOS_EVENT_KIND_STARTED           int32
	FABRIC_CHAOS_EVENT_KIND_EXECUTING_FAULTS  int32
	FABRIC_CHAOS_EVENT_KIND_WAITING           int32
	FABRIC_CHAOS_EVENT_KIND_VALIDATION_FAILED int32
	FABRIC_CHAOS_EVENT_KIND_TEST_ERROR        int32
	FABRIC_CHAOS_EVENT_KIND_STOPPED           int32
}{
	FABRIC_CHAOS_EVENT_KIND_INVALID:           0,
	FABRIC_CHAOS_EVENT_KIND_STARTED:           1,
	FABRIC_CHAOS_EVENT_KIND_EXECUTING_FAULTS:  2,
	FABRIC_CHAOS_EVENT_KIND_WAITING:           3,
	FABRIC_CHAOS_EVENT_KIND_VALIDATION_FAILED: 4,
	FABRIC_CHAOS_EVENT_KIND_TEST_ERROR:        5,
	FABRIC_CHAOS_EVENT_KIND_STOPPED:           6,
}

// enum FABRIC_NODE_TRANSITION_TYPE
var FABRIC_NODE_TRANSITION_TYPE = struct {
	FABRIC_NODE_TRANSITION_TYPE_INVALID int32
	FABRIC_NODE_TRANSITION_TYPE_START   int32
	FABRIC_NODE_TRANSITION_TYPE_STOP    int32
}{
	FABRIC_NODE_TRANSITION_TYPE_INVALID: 0,
	FABRIC_NODE_TRANSITION_TYPE_START:   1,
	FABRIC_NODE_TRANSITION_TYPE_STOP:    2,
}

// enum FABRIC_RESTART_NODE_DESCRIPTION_KIND
var FABRIC_RESTART_NODE_DESCRIPTION_KIND = struct {
	FABRIC_RESTART_NODE_DESCRIPTION_KIND_INVALID         int32
	FABRIC_RESTART_NODE_DESCRIPTION_KIND_USING_NODE_NAME int32
}{
	FABRIC_RESTART_NODE_DESCRIPTION_KIND_INVALID:         0,
	FABRIC_RESTART_NODE_DESCRIPTION_KIND_USING_NODE_NAME: 1,
}

// enum FABRIC_START_NODE_DESCRIPTION_KIND
var FABRIC_START_NODE_DESCRIPTION_KIND = struct {
	FABRIC_START_NODE_DESCRIPTION_KIND_INVALID         int32
	FABRIC_START_NODE_DESCRIPTION_KIND_USING_NODE_NAME int32
}{
	FABRIC_START_NODE_DESCRIPTION_KIND_INVALID:         0,
	FABRIC_START_NODE_DESCRIPTION_KIND_USING_NODE_NAME: 1,
}

// enum FABRIC_STOP_NODE_DESCRIPTION_KIND
var FABRIC_STOP_NODE_DESCRIPTION_KIND = struct {
	FABRIC_STOP_NODE_DESCRIPTION_KIND_INVALID         int32
	FABRIC_STOP_NODE_DESCRIPTION_KIND_USING_NODE_NAME int32
}{
	FABRIC_STOP_NODE_DESCRIPTION_KIND_INVALID:         0,
	FABRIC_STOP_NODE_DESCRIPTION_KIND_USING_NODE_NAME: 1,
}

// enum FABRIC_RESTART_DEPLOYED_CODE_PACKAGE_DESCRIPTION_KIND
var FABRIC_RESTART_DEPLOYED_CODE_PACKAGE_DESCRIPTION_KIND = struct {
	FABRIC_RESTART_DEPLOYED_CODE_PACKAGE_DESCRIPTION_KIND_INVALID         int32
	FABRIC_RESTART_DEPLOYED_CODE_PACKAGE_DESCRIPTION_KIND_USING_NODE_NAME int32
}{
	FABRIC_RESTART_DEPLOYED_CODE_PACKAGE_DESCRIPTION_KIND_INVALID:         0,
	FABRIC_RESTART_DEPLOYED_CODE_PACKAGE_DESCRIPTION_KIND_USING_NODE_NAME: 1,
}

// enum FABRIC_MOVE_PRIMARY_DESCRIPTION_KIND
var FABRIC_MOVE_PRIMARY_DESCRIPTION_KIND = struct {
	FABRIC_MOVE_PRIMARY_DESCRIPTION_KIND_INVALID                int32
	FABRIC_MOVE_PRIMARY_DESCRIPTION_KIND_USING_NODE_NAME        int32
	FABRIC_MOVE_PRIMARY_DESCRIPTION_KIND_USING_REPLICA_SELECTOR int32
}{
	FABRIC_MOVE_PRIMARY_DESCRIPTION_KIND_INVALID:                0,
	FABRIC_MOVE_PRIMARY_DESCRIPTION_KIND_USING_NODE_NAME:        1,
	FABRIC_MOVE_PRIMARY_DESCRIPTION_KIND_USING_REPLICA_SELECTOR: 2,
}

// enum FABRIC_MOVE_SECONDARY_DESCRIPTION_KIND
var FABRIC_MOVE_SECONDARY_DESCRIPTION_KIND = struct {
	FABRIC_MOVE_SECONDARY_DESCRIPTION_KIND_INVALID                int32
	FABRIC_MOVE_SECONDARY_DESCRIPTION_KIND_USING_NODE_NAME        int32
	FABRIC_MOVE_SECONDARY_DESCRIPTION_KIND_USING_REPLICA_SELECTOR int32
}{
	FABRIC_MOVE_SECONDARY_DESCRIPTION_KIND_INVALID:                0,
	FABRIC_MOVE_SECONDARY_DESCRIPTION_KIND_USING_NODE_NAME:        1,
	FABRIC_MOVE_SECONDARY_DESCRIPTION_KIND_USING_REPLICA_SELECTOR: 2,
}

// enum FABRIC_NETWORK_TYPE
var FABRIC_NETWORK_TYPE = struct {
	FABRIC_NETWORK_TYPE_INVALID   int32
	FABRIC_NETWORK_TYPE_LOCAL     int32
	FABRIC_NETWORK_TYPE_FEDERATED int32
}{
	FABRIC_NETWORK_TYPE_INVALID:   0,
	FABRIC_NETWORK_TYPE_LOCAL:     1,
	FABRIC_NETWORK_TYPE_FEDERATED: 2,
}

// enum FABRIC_HEALTH_EVALUATION_KIND
var FABRIC_HEALTH_EVALUATION_KIND = struct {
	FABRIC_HEALTH_EVALUATION_KIND_INVALID                              int32
	FABRIC_HEALTH_EVALUATION_KIND_EVENT                                int32
	FABRIC_HEALTH_EVALUATION_KIND_REPLICAS                             int32
	FABRIC_HEALTH_EVALUATION_KIND_PARTITIONS                           int32
	FABRIC_HEALTH_EVALUATION_KIND_DEPLOYED_SERVICE_PACKAGES            int32
	FABRIC_HEALTH_EVALUATION_KIND_DEPLOYED_APPLICATIONS                int32
	FABRIC_HEALTH_EVALUATION_KIND_SERVICES                             int32
	FABRIC_HEALTH_EVALUATION_KIND_NODES                                int32
	FABRIC_HEALTH_EVALUATION_KIND_APPLICATIONS                         int32
	FABRIC_HEALTH_EVALUATION_KIND_SYSTEM_APPLICATION                   int32
	FABRIC_HEALTH_EVALUATION_KIND_UPGRADE_DOMAIN_DEPLOYED_APPLICATIONS int32
	FABRIC_HEALTH_EVALUATION_KIND_UPGRADE_DOMAIN_NODES                 int32
	FABRIC_HEALTH_EVALUATION_KIND_NODE                                 int32
	FABRIC_HEALTH_EVALUATION_KIND_REPLICA                              int32
	FABRIC_HEALTH_EVALUATION_KIND_PARTITION                            int32
	FABRIC_HEALTH_EVALUATION_KIND_SERVICE                              int32
	FABRIC_HEALTH_EVALUATION_KIND_DEPLOYED_SERVICE_PACKAGE             int32
	FABRIC_HEALTH_EVALUATION_KIND_DEPLOYED_APPLICATION                 int32
	FABRIC_HEALTH_EVALUATION_KIND_APPLICATION                          int32
	FABRIC_HEALTH_EVALUATION_KIND_DELTA_NODES_CHECK                    int32
	FABRIC_HEALTH_EVALUATION_KIND_UPGRADE_DOMAIN_DELTA_NODES_CHECK     int32
	FABRIC_HEALTH_EVALUATION_KIND_APPLICATION_TYPE_APPLICATIONS        int32
}{
	FABRIC_HEALTH_EVALUATION_KIND_INVALID:                              0,
	FABRIC_HEALTH_EVALUATION_KIND_EVENT:                                1,
	FABRIC_HEALTH_EVALUATION_KIND_REPLICAS:                             2,
	FABRIC_HEALTH_EVALUATION_KIND_PARTITIONS:                           3,
	FABRIC_HEALTH_EVALUATION_KIND_DEPLOYED_SERVICE_PACKAGES:            4,
	FABRIC_HEALTH_EVALUATION_KIND_DEPLOYED_APPLICATIONS:                5,
	FABRIC_HEALTH_EVALUATION_KIND_SERVICES:                             6,
	FABRIC_HEALTH_EVALUATION_KIND_NODES:                                7,
	FABRIC_HEALTH_EVALUATION_KIND_APPLICATIONS:                         8,
	FABRIC_HEALTH_EVALUATION_KIND_SYSTEM_APPLICATION:                   9,
	FABRIC_HEALTH_EVALUATION_KIND_UPGRADE_DOMAIN_DEPLOYED_APPLICATIONS: 10,
	FABRIC_HEALTH_EVALUATION_KIND_UPGRADE_DOMAIN_NODES:                 11,
	FABRIC_HEALTH_EVALUATION_KIND_NODE:                                 12,
	FABRIC_HEALTH_EVALUATION_KIND_REPLICA:                              13,
	FABRIC_HEALTH_EVALUATION_KIND_PARTITION:                            14,
	FABRIC_HEALTH_EVALUATION_KIND_SERVICE:                              15,
	FABRIC_HEALTH_EVALUATION_KIND_DEPLOYED_SERVICE_PACKAGE:             16,
	FABRIC_HEALTH_EVALUATION_KIND_DEPLOYED_APPLICATION:                 17,
	FABRIC_HEALTH_EVALUATION_KIND_APPLICATION:                          18,
	FABRIC_HEALTH_EVALUATION_KIND_DELTA_NODES_CHECK:                    19,
	FABRIC_HEALTH_EVALUATION_KIND_UPGRADE_DOMAIN_DELTA_NODES_CHECK:     20,
	FABRIC_HEALTH_EVALUATION_KIND_APPLICATION_TYPE_APPLICATIONS:        21,
}

// enum FABRIC_NODE_UPGRADE_PHASE
var FABRIC_NODE_UPGRADE_PHASE = struct {
	FABRIC_NODE_UPGRADE_PHASE_INVALID                   int32
	FABRIC_NODE_UPGRADE_PHASE_PRE_UPGRADE_SAFETY_CHECK  int32
	FABRIC_NODE_UPGRADE_PHASE_UPGRADING                 int32
	FABRIC_NODE_UPGRADE_PHASE_POST_UPGRADE_SAFETY_CHECK int32
}{
	FABRIC_NODE_UPGRADE_PHASE_INVALID:                   0,
	FABRIC_NODE_UPGRADE_PHASE_PRE_UPGRADE_SAFETY_CHECK:  1,
	FABRIC_NODE_UPGRADE_PHASE_UPGRADING:                 2,
	FABRIC_NODE_UPGRADE_PHASE_POST_UPGRADE_SAFETY_CHECK: 3,
}

// enum FABRIC_UPGRADE_SAFETY_CHECK_KIND
var FABRIC_UPGRADE_SAFETY_CHECK_KIND = struct {
	FABRIC_UPGRADE_SAFETY_CHECK_KIND_INVALID                                  int32
	FABRIC_UPGRADE_SEED_NODE_SAFETY_CHECK_KIND_ENSURE_QUORUM                  int32
	FABRIC_UPGRADE_PARTITION_SAFETY_CHECK_KIND_ENSURE_QUORUM                  int32
	FABRIC_UPGRADE_PARTITION_SAFETY_CHECK_KIND_WAIT_FOR_PRIMARY_PLACEMENT     int32
	FABRIC_UPGRADE_PARTITION_SAFETY_CHECK_KIND_WAIT_FOR_PRIMARY_SWAP          int32
	FABRIC_UPGRADE_PARTITION_SAFETY_CHECK_KIND_WAIT_FOR_RECONFIGURATION       int32
	FABRIC_UPGRADE_PARTITION_SAFETY_CHECK_KIND_WAIT_FOR_INBUILD_REPLICA       int32
	FABRIC_UPGRADE_PARTITION_SAFETY_CHECK_KIND_ENSURE_AVAILABILITY            int32
	FABRIC_UPGRADE_PARTITION_SAFETY_CHECK_KIND_WAIT_FOR_RESOURCE_AVAILABILITY int32
}{
	FABRIC_UPGRADE_SAFETY_CHECK_KIND_INVALID:                                  0,
	FABRIC_UPGRADE_SEED_NODE_SAFETY_CHECK_KIND_ENSURE_QUORUM:                  1,
	FABRIC_UPGRADE_PARTITION_SAFETY_CHECK_KIND_ENSURE_QUORUM:                  2,
	FABRIC_UPGRADE_PARTITION_SAFETY_CHECK_KIND_WAIT_FOR_PRIMARY_PLACEMENT:     3,
	FABRIC_UPGRADE_PARTITION_SAFETY_CHECK_KIND_WAIT_FOR_PRIMARY_SWAP:          4,
	FABRIC_UPGRADE_PARTITION_SAFETY_CHECK_KIND_WAIT_FOR_RECONFIGURATION:       5,
	FABRIC_UPGRADE_PARTITION_SAFETY_CHECK_KIND_WAIT_FOR_INBUILD_REPLICA:       6,
	FABRIC_UPGRADE_PARTITION_SAFETY_CHECK_KIND_ENSURE_AVAILABILITY:            7,
	FABRIC_UPGRADE_PARTITION_SAFETY_CHECK_KIND_WAIT_FOR_RESOURCE_AVAILABILITY: 8,
}

// enum FABRIC_MONITORED_UPGRADE_FAILURE_ACTION
var FABRIC_MONITORED_UPGRADE_FAILURE_ACTION = struct {
	FABRIC_MONITORED_UPGRADE_FAILURE_ACTION_INVALID  int32
	FABRIC_MONITORED_UPGRADE_FAILURE_ACTION_ROLLBACK int32
	FABRIC_MONITORED_UPGRADE_FAILURE_ACTION_MANUAL   int32
}{
	FABRIC_MONITORED_UPGRADE_FAILURE_ACTION_INVALID:  0,
	FABRIC_MONITORED_UPGRADE_FAILURE_ACTION_ROLLBACK: 1,
	FABRIC_MONITORED_UPGRADE_FAILURE_ACTION_MANUAL:   2,
}

// enum FABRIC_CHAOS_SCHEDULE_STATUS
var FABRIC_CHAOS_SCHEDULE_STATUS = struct {
	FABRIC_CHAOS_SCHEDULE_STATUS_INVALID int32
	FABRIC_CHAOS_SCHEDULE_STATUS_ACTIVE  int32
	FABRIC_CHAOS_SCHEDULE_STATUS_EXPIRED int32
	FABRIC_CHAOS_SCHEDULE_STATUS_PENDING int32
	FABRIC_CHAOS_SCHEDULE_STATUS_STOPPED int32
}{
	FABRIC_CHAOS_SCHEDULE_STATUS_INVALID: 0,
	FABRIC_CHAOS_SCHEDULE_STATUS_ACTIVE:  1,
	FABRIC_CHAOS_SCHEDULE_STATUS_EXPIRED: 2,
	FABRIC_CHAOS_SCHEDULE_STATUS_PENDING: 3,
	FABRIC_CHAOS_SCHEDULE_STATUS_STOPPED: 4,
}

// enum FABRIC_CLIENT_ROLE
var FABRIC_CLIENT_ROLE = struct {
	FABRIC_CLIENT_ROLE_UNKNOWN int32
	FABRIC_CLIENT_ROLE_USER    int32
	FABRIC_CLIENT_ROLE_ADMIN   int32
}{
	FABRIC_CLIENT_ROLE_UNKNOWN: 0,
	FABRIC_CLIENT_ROLE_USER:    1,
	FABRIC_CLIENT_ROLE_ADMIN:   2,
}