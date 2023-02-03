// Code generated by user config generator. DO NOT EDIT.
// +kubebuilder:object:generate=true

package kafkauserconfig

import "encoding/json"

func (ip *IpFilter) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	var s string
	err := json.Unmarshal(data, &s)
	if err == nil {
		ip.Network = s
		return nil
	}

	type this struct {
		Network     string  `json:"network"`
		Description *string `json:"description,omitempty" `
	}

	var t *this
	err = json.Unmarshal(data, &t)
	if err != nil {
		return err
	}
	ip.Network = t.Network
	ip.Description = t.Description
	return nil
}

// IpFilter CIDR address block, either as a string, or in a dict with an optional description field
type IpFilter struct {
	// +kubebuilder:validation:MaxLength=1024
	// Description for IP filter list entry
	Description *string `groups:"create,update" json:"description,omitempty"`

	// +kubebuilder:validation:MaxLength=43
	// Network CIDR address block
	Network string `groups:"create,update" json:"network"`
}

// Kafka broker configuration values
type Kafka struct {
	// AutoCreateTopicsEnable Enable auto creation of topics
	AutoCreateTopicsEnable *bool `groups:"create,update" json:"auto_create_topics_enable,omitempty"`

	// +kubebuilder:validation:Enum=gzip;snappy;lz4;zstd;uncompressed;producer
	// CompressionType Specify the final compression type for a given topic. This configuration accepts the standard compression codecs ('gzip', 'snappy', 'lz4', 'zstd'). It additionally accepts 'uncompressed' which is equivalent to no compression; and 'producer' which means retain the original compression codec set by the producer.
	CompressionType *string `groups:"create,update" json:"compression_type,omitempty"`

	// +kubebuilder:validation:Minimum=1000
	// +kubebuilder:validation:Maximum=3600000
	// ConnectionsMaxIdleMs Idle connections timeout: the server socket processor threads close the connections that idle for longer than this.
	ConnectionsMaxIdleMs *int `groups:"create,update" json:"connections_max_idle_ms,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=10
	// DefaultReplicationFactor Replication factor for autocreated topics
	DefaultReplicationFactor *int `groups:"create,update" json:"default_replication_factor,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=300000
	// GroupInitialRebalanceDelayMs The amount of time, in milliseconds, the group coordinator will wait for more consumers to join a new group before performing the first rebalance. A longer delay means potentially fewer rebalances, but increases the time until processing begins. The default value for this is 3 seconds. During development and testing it might be desirable to set this to 0 in order to not delay test execution time.
	GroupInitialRebalanceDelayMs *int `groups:"create,update" json:"group_initial_rebalance_delay_ms,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1800000
	// GroupMaxSessionTimeoutMs The maximum allowed session timeout for registered consumers. Longer timeouts give consumers more time to process messages in between heartbeats at the cost of a longer time to detect failures.
	GroupMaxSessionTimeoutMs *int `groups:"create,update" json:"group_max_session_timeout_ms,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=60000
	// GroupMinSessionTimeoutMs The minimum allowed session timeout for registered consumers. Longer timeouts give consumers more time to process messages in between heartbeats at the cost of a longer time to detect failures.
	GroupMinSessionTimeoutMs *int `groups:"create,update" json:"group_min_session_timeout_ms,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=315569260000
	// LogCleanerDeleteRetentionMs How long are delete records retained?
	LogCleanerDeleteRetentionMs *int `groups:"create,update" json:"log_cleaner_delete_retention_ms,omitempty"`

	// +kubebuilder:validation:Minimum=30000
	// LogCleanerMaxCompactionLagMs The maximum amount of time message will remain uncompacted. Only applicable for logs that are being compacted
	LogCleanerMaxCompactionLagMs *int `groups:"create,update" json:"log_cleaner_max_compaction_lag_ms,omitempty"`

	// LogCleanerMinCleanableRatio Controls log compactor frequency. Larger value means more frequent compactions but also more space wasted for logs. Consider setting log.cleaner.max.compaction.lag.ms to enforce compactions sooner, instead of setting a very high value for this option.
	LogCleanerMinCleanableRatio *float64 `groups:"create,update" json:"log_cleaner_min_cleanable_ratio,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// LogCleanerMinCompactionLagMs The minimum time a message will remain uncompacted in the log. Only applicable for logs that are being compacted.
	LogCleanerMinCompactionLagMs *int `groups:"create,update" json:"log_cleaner_min_compaction_lag_ms,omitempty"`

	// +kubebuilder:validation:Enum=delete;compact;"compact,delete"
	// LogCleanupPolicy The default cleanup policy for segments beyond the retention window
	LogCleanupPolicy *string `groups:"create,update" json:"log_cleanup_policy,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// LogFlushIntervalMessages The number of messages accumulated on a log partition before messages are flushed to disk
	LogFlushIntervalMessages *int `groups:"create,update" json:"log_flush_interval_messages,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// LogFlushIntervalMs The maximum time in ms that a message in any topic is kept in memory before flushed to disk. If not set, the value in log.flush.scheduler.interval.ms is used
	LogFlushIntervalMs *int `groups:"create,update" json:"log_flush_interval_ms,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=104857600
	// LogIndexIntervalBytes The interval with which Kafka adds an entry to the offset index
	LogIndexIntervalBytes *int `groups:"create,update" json:"log_index_interval_bytes,omitempty"`

	// +kubebuilder:validation:Minimum=1048576
	// +kubebuilder:validation:Maximum=104857600
	// LogIndexSizeMaxBytes The maximum size in bytes of the offset index
	LogIndexSizeMaxBytes *int `groups:"create,update" json:"log_index_size_max_bytes,omitempty"`

	// LogMessageDownconversionEnable This configuration controls whether down-conversion of message formats is enabled to satisfy consume requests.
	LogMessageDownconversionEnable *bool `groups:"create,update" json:"log_message_downconversion_enable,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// LogMessageTimestampDifferenceMaxMs The maximum difference allowed between the timestamp when a broker receives a message and the timestamp specified in the message
	LogMessageTimestampDifferenceMaxMs *int `groups:"create,update" json:"log_message_timestamp_difference_max_ms,omitempty"`

	// +kubebuilder:validation:Enum=CreateTime;LogAppendTime
	// LogMessageTimestampType Define whether the timestamp in the message is message create time or log append time.
	LogMessageTimestampType *string `groups:"create,update" json:"log_message_timestamp_type,omitempty"`

	// LogPreallocate Should pre allocate file when create new segment?
	LogPreallocate *bool `groups:"create,update" json:"log_preallocate,omitempty"`

	// +kubebuilder:validation:Minimum=-1
	// LogRetentionBytes The maximum size of the log before deleting messages
	LogRetentionBytes *int `groups:"create,update" json:"log_retention_bytes,omitempty"`

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=2147483647
	// LogRetentionHours The number of hours to keep a log file before deleting it
	LogRetentionHours *int `groups:"create,update" json:"log_retention_hours,omitempty"`

	// +kubebuilder:validation:Minimum=-1
	// LogRetentionMs The number of milliseconds to keep a log file before deleting it (in milliseconds), If not set, the value in log.retention.minutes is used. If set to -1, no time limit is applied.
	LogRetentionMs *int `groups:"create,update" json:"log_retention_ms,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// LogRollJitterMs The maximum jitter to subtract from logRollTimeMillis (in milliseconds). If not set, the value in log.roll.jitter.hours is used
	LogRollJitterMs *int `groups:"create,update" json:"log_roll_jitter_ms,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// LogRollMs The maximum time before a new log segment is rolled out (in milliseconds).
	LogRollMs *int `groups:"create,update" json:"log_roll_ms,omitempty"`

	// +kubebuilder:validation:Minimum=10485760
	// +kubebuilder:validation:Maximum=1073741824
	// LogSegmentBytes The maximum size of a single log file
	LogSegmentBytes *int `groups:"create,update" json:"log_segment_bytes,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=3600000
	// LogSegmentDeleteDelayMs The amount of time to wait before deleting a file from the filesystem
	LogSegmentDeleteDelayMs *int `groups:"create,update" json:"log_segment_delete_delay_ms,omitempty"`

	// +kubebuilder:validation:Minimum=256
	// +kubebuilder:validation:Maximum=2147483647
	// MaxConnectionsPerIp The maximum number of connections allowed from each ip address (defaults to 2147483647).
	MaxConnectionsPerIp *int `groups:"create,update" json:"max_connections_per_ip,omitempty"`

	// +kubebuilder:validation:Minimum=1000
	// +kubebuilder:validation:Maximum=10000
	// MaxIncrementalFetchSessionCacheSlots The maximum number of incremental fetch sessions that the broker will maintain.
	MaxIncrementalFetchSessionCacheSlots *int `groups:"create,update" json:"max_incremental_fetch_session_cache_slots,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100001200
	// MessageMaxBytes The maximum size of message that the server can receive.
	MessageMaxBytes *int `groups:"create,update" json:"message_max_bytes,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=7
	// MinInsyncReplicas When a producer sets acks to 'all' (or '-1'), min.insync.replicas specifies the minimum number of replicas that must acknowledge a write for the write to be considered successful.
	MinInsyncReplicas *int `groups:"create,update" json:"min_insync_replicas,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=1000
	// NumPartitions Number of partitions for autocreated topics
	NumPartitions *int `groups:"create,update" json:"num_partitions,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2147483647
	// OffsetsRetentionMinutes Log retention window in minutes for offsets topic
	OffsetsRetentionMinutes *int `groups:"create,update" json:"offsets_retention_minutes,omitempty"`

	// +kubebuilder:validation:Minimum=10
	// +kubebuilder:validation:Maximum=10000
	// ProducerPurgatoryPurgeIntervalRequests The purge interval (in number of requests) of the producer request purgatory(defaults to 1000).
	ProducerPurgatoryPurgeIntervalRequests *int `groups:"create,update" json:"producer_purgatory_purge_interval_requests,omitempty"`

	// +kubebuilder:validation:Minimum=1048576
	// +kubebuilder:validation:Maximum=104857600
	// ReplicaFetchMaxBytes The number of bytes of messages to attempt to fetch for each partition (defaults to 1048576). This is not an absolute maximum, if the first record batch in the first non-empty partition of the fetch is larger than this value, the record batch will still be returned to ensure that progress can be made.
	ReplicaFetchMaxBytes *int `groups:"create,update" json:"replica_fetch_max_bytes,omitempty"`

	// +kubebuilder:validation:Minimum=10485760
	// +kubebuilder:validation:Maximum=1048576000
	// ReplicaFetchResponseMaxBytes Maximum bytes expected for the entire fetch response (defaults to 10485760). Records are fetched in batches, and if the first record batch in the first non-empty partition of the fetch is larger than this value, the record batch will still be returned to ensure that progress can be made. As such, this is not an absolute maximum.
	ReplicaFetchResponseMaxBytes *int `groups:"create,update" json:"replica_fetch_response_max_bytes,omitempty"`

	// +kubebuilder:validation:Minimum=10485760
	// +kubebuilder:validation:Maximum=209715200
	// SocketRequestMaxBytes The maximum number of bytes in a socket request (defaults to 104857600).
	SocketRequestMaxBytes *int `groups:"create,update" json:"socket_request_max_bytes,omitempty"`

	// +kubebuilder:validation:Minimum=600000
	// +kubebuilder:validation:Maximum=3600000
	// TransactionRemoveExpiredTransactionCleanupIntervalMs The interval at which to remove transactions that have expired due to transactional.id.expiration.ms passing (defaults to 3600000 (1 hour)).
	TransactionRemoveExpiredTransactionCleanupIntervalMs *int `groups:"create,update" json:"transaction_remove_expired_transaction_cleanup_interval_ms,omitempty"`

	// +kubebuilder:validation:Minimum=1048576
	// +kubebuilder:validation:Maximum=2147483647
	// TransactionStateLogSegmentBytes The transaction topic segment bytes should be kept relatively small in order to facilitate faster log compaction and cache loads (defaults to 104857600 (100 mebibytes)).
	TransactionStateLogSegmentBytes *int `groups:"create,update" json:"transaction_state_log_segment_bytes,omitempty"`
}

// KafkaAuthenticationMethods Kafka authentication methods
type KafkaAuthenticationMethods struct {
	// Certificate Enable certificate/SSL authentication
	Certificate *bool `groups:"create,update" json:"certificate,omitempty"`

	// Sasl Enable SASL authentication
	Sasl *bool `groups:"create,update" json:"sasl,omitempty"`
}

// KafkaConnectConfig Kafka Connect configuration values
type KafkaConnectConfig struct {
	// +kubebuilder:validation:Enum=None;All
	// ConnectorClientConfigOverridePolicy Defines what client configurations can be overridden by the connector. Default is None
	ConnectorClientConfigOverridePolicy *string `groups:"create,update" json:"connector_client_config_override_policy,omitempty"`

	// +kubebuilder:validation:Enum=earliest;latest
	// ConsumerAutoOffsetReset What to do when there is no initial offset in Kafka or if the current offset does not exist any more on the server. Default is earliest
	ConsumerAutoOffsetReset *string `groups:"create,update" json:"consumer_auto_offset_reset,omitempty"`

	// +kubebuilder:validation:Minimum=1048576
	// +kubebuilder:validation:Maximum=104857600
	// ConsumerFetchMaxBytes Records are fetched in batches by the consumer, and if the first record batch in the first non-empty partition of the fetch is larger than this value, the record batch will still be returned to ensure that the consumer can make progress. As such, this is not a absolute maximum.
	ConsumerFetchMaxBytes *int `groups:"create,update" json:"consumer_fetch_max_bytes,omitempty"`

	// +kubebuilder:validation:Enum=read_uncommitted;read_committed
	// ConsumerIsolationLevel Transaction read isolation level. read_uncommitted is the default, but read_committed can be used if consume-exactly-once behavior is desired.
	ConsumerIsolationLevel *string `groups:"create,update" json:"consumer_isolation_level,omitempty"`

	// +kubebuilder:validation:Minimum=1048576
	// +kubebuilder:validation:Maximum=104857600
	// ConsumerMaxPartitionFetchBytes Records are fetched in batches by the consumer.If the first record batch in the first non-empty partition of the fetch is larger than this limit, the batch will still be returned to ensure that the consumer can make progress.
	ConsumerMaxPartitionFetchBytes *int `groups:"create,update" json:"consumer_max_partition_fetch_bytes,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2147483647
	// ConsumerMaxPollIntervalMs The maximum delay in milliseconds between invocations of poll() when using consumer group management (defaults to 300000).
	ConsumerMaxPollIntervalMs *int `groups:"create,update" json:"consumer_max_poll_interval_ms,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=10000
	// ConsumerMaxPollRecords The maximum number of records returned in a single call to poll() (defaults to 500).
	ConsumerMaxPollRecords *int `groups:"create,update" json:"consumer_max_poll_records,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=100000000
	// OffsetFlushIntervalMs The interval at which to try committing offsets for tasks (defaults to 60000).
	OffsetFlushIntervalMs *int `groups:"create,update" json:"offset_flush_interval_ms,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2147483647
	// OffsetFlushTimeoutMs Maximum number of milliseconds to wait for records to flush and partition offset data to be committed to offset storage before cancelling the process and restoring the offset data to be committed in a future attempt (defaults to 5000).
	OffsetFlushTimeoutMs *int `groups:"create,update" json:"offset_flush_timeout_ms,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=5242880
	// ProducerBatchSize This setting gives the upper bound of the batch size to be sent. If there are fewer than this many bytes accumulated for this partition, the producer will 'linger' for the linger.ms time waiting for more records to show up. A batch size of zero will disable batching entirely (defaults to 16384).
	ProducerBatchSize *int `groups:"create,update" json:"producer_batch_size,omitempty"`

	// +kubebuilder:validation:Minimum=5242880
	// +kubebuilder:validation:Maximum=134217728
	// ProducerBufferMemory The total bytes of memory the producer can use to buffer records waiting to be sent to the broker (defaults to 33554432).
	ProducerBufferMemory *int `groups:"create,update" json:"producer_buffer_memory,omitempty"`

	// +kubebuilder:validation:Enum=gzip;snappy;lz4;zstd;none
	// ProducerCompressionType Specify the default compression type for producers. This configuration accepts the standard compression codecs ('gzip', 'snappy', 'lz4', 'zstd'). It additionally accepts 'none' which is the default and equivalent to no compression.
	ProducerCompressionType *string `groups:"create,update" json:"producer_compression_type,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=5000
	// ProducerLingerMs This setting gives the upper bound on the delay for batching: once there is batch.size worth of records for a partition it will be sent immediately regardless of this setting, however if there are fewer than this many bytes accumulated for this partition the producer will 'linger' for the specified time waiting for more records to show up. Defaults to 0.
	ProducerLingerMs *int `groups:"create,update" json:"producer_linger_ms,omitempty"`

	// +kubebuilder:validation:Minimum=131072
	// +kubebuilder:validation:Maximum=67108864
	// ProducerMaxRequestSize This setting will limit the number of record batches the producer will send in a single request to avoid sending huge requests.
	ProducerMaxRequestSize *int `groups:"create,update" json:"producer_max_request_size,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2147483647
	// SessionTimeoutMs The timeout in milliseconds used to detect failures when using Kafka’s group management facilities (defaults to 10000).
	SessionTimeoutMs *int `groups:"create,update" json:"session_timeout_ms,omitempty"`
}

// KafkaRestConfig Kafka REST configuration
type KafkaRestConfig struct {
	// ConsumerEnableAutoCommit If true the consumer's offset will be periodically committed to Kafka in the background
	ConsumerEnableAutoCommit *bool `groups:"create,update" json:"consumer_enable_auto_commit,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=671088640
	// ConsumerRequestMaxBytes Maximum number of bytes in unencoded message keys and values by a single request
	ConsumerRequestMaxBytes *int `groups:"create,update" json:"consumer_request_max_bytes,omitempty"`

	// +kubebuilder:validation:Minimum=1000
	// +kubebuilder:validation:Maximum=30000
	// +kubebuilder:validation:Enum=1000;15000;30000
	// ConsumerRequestTimeoutMs The maximum total time to wait for messages for a request if the maximum number of messages has not yet been reached
	ConsumerRequestTimeoutMs *int `groups:"create,update" json:"consumer_request_timeout_ms,omitempty"`

	// +kubebuilder:validation:Enum=all;-1;0;1
	// ProducerAcks The number of acknowledgments the producer requires the leader to have received before considering a request complete. If set to 'all' or '-1', the leader will wait for the full set of in-sync replicas to acknowledge the record.
	ProducerAcks *string `groups:"create,update" json:"producer_acks,omitempty"`

	// +kubebuilder:validation:Enum=gzip;snappy;lz4;zstd;none
	// ProducerCompressionType Specify the default compression type for producers. This configuration accepts the standard compression codecs ('gzip', 'snappy', 'lz4', 'zstd'). It additionally accepts 'none' which is the default and equivalent to no compression.
	ProducerCompressionType *string `groups:"create,update" json:"producer_compression_type,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=5000
	// ProducerLingerMs Wait for up to the given delay to allow batching records together
	ProducerLingerMs *int `groups:"create,update" json:"producer_linger_ms,omitempty"`

	// +kubebuilder:validation:Minimum=10
	// +kubebuilder:validation:Maximum=250
	// SimpleconsumerPoolSizeMax Maximum number of SimpleConsumers that can be instantiated per broker
	SimpleconsumerPoolSizeMax *int `groups:"create,update" json:"simpleconsumer_pool_size_max,omitempty"`
}

// PrivateAccess Allow access to selected service ports from private networks
type PrivateAccess struct {
	// Kafka Allow clients to connect to kafka with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations
	Kafka *bool `groups:"create,update" json:"kafka,omitempty"`

	// KafkaConnect Allow clients to connect to kafka_connect with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations
	KafkaConnect *bool `groups:"create,update" json:"kafka_connect,omitempty"`

	// KafkaRest Allow clients to connect to kafka_rest with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations
	KafkaRest *bool `groups:"create,update" json:"kafka_rest,omitempty"`

	// Prometheus Allow clients to connect to prometheus with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations
	Prometheus *bool `groups:"create,update" json:"prometheus,omitempty"`

	// SchemaRegistry Allow clients to connect to schema_registry with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations
	SchemaRegistry *bool `groups:"create,update" json:"schema_registry,omitempty"`
}

// PrivatelinkAccess Allow access to selected service components through Privatelink
type PrivatelinkAccess struct {
	// Jolokia Enable jolokia
	Jolokia *bool `groups:"create,update" json:"jolokia,omitempty"`

	// Kafka Enable kafka
	Kafka *bool `groups:"create,update" json:"kafka,omitempty"`

	// KafkaConnect Enable kafka_connect
	KafkaConnect *bool `groups:"create,update" json:"kafka_connect,omitempty"`

	// KafkaRest Enable kafka_rest
	KafkaRest *bool `groups:"create,update" json:"kafka_rest,omitempty"`

	// Prometheus Enable prometheus
	Prometheus *bool `groups:"create,update" json:"prometheus,omitempty"`

	// SchemaRegistry Enable schema_registry
	SchemaRegistry *bool `groups:"create,update" json:"schema_registry,omitempty"`
}

// PublicAccess Allow access to selected service ports from the public Internet
type PublicAccess struct {
	// Kafka Allow clients to connect to kafka from the public internet for service nodes that are in a project VPC or another type of private network
	Kafka *bool `groups:"create,update" json:"kafka,omitempty"`

	// KafkaConnect Allow clients to connect to kafka_connect from the public internet for service nodes that are in a project VPC or another type of private network
	KafkaConnect *bool `groups:"create,update" json:"kafka_connect,omitempty"`

	// KafkaRest Allow clients to connect to kafka_rest from the public internet for service nodes that are in a project VPC or another type of private network
	KafkaRest *bool `groups:"create,update" json:"kafka_rest,omitempty"`

	// Prometheus Allow clients to connect to prometheus from the public internet for service nodes that are in a project VPC or another type of private network
	Prometheus *bool `groups:"create,update" json:"prometheus,omitempty"`

	// SchemaRegistry Allow clients to connect to schema_registry from the public internet for service nodes that are in a project VPC or another type of private network
	SchemaRegistry *bool `groups:"create,update" json:"schema_registry,omitempty"`
}

// SchemaRegistryConfig Schema Registry configuration
type SchemaRegistryConfig struct {
	// LeaderEligibility If true, Karapace / Schema Registry on the service nodes can participate in leader election. It might be needed to disable this when the schemas topic is replicated to a secondary cluster and Karapace / Schema Registry there must not participate in leader election. Defaults to `true`.
	LeaderEligibility *bool `groups:"create,update" json:"leader_eligibility,omitempty"`

	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=249
	// TopicName The durable single partition topic that acts as the durable log for the data. This topic must be compacted to avoid losing data due to retention policy. Please note that changing this configuration in an existing Schema Registry / Karapace setup leads to previous schemas being inaccessible, data encoded with them potentially unreadable and schema ID sequence put out of order. It's only possible to do the switch while Schema Registry / Karapace is disabled. Defaults to `_schemas`.
	TopicName *string `groups:"create,update" json:"topic_name,omitempty"`
}
type KafkaUserConfig struct {
	// +kubebuilder:validation:MaxItems=1
	// AdditionalBackupRegions Additional Cloud Regions for Backup Replication
	AdditionalBackupRegions []string `groups:"create,update" json:"additional_backup_regions,omitempty"`

	// +kubebuilder:validation:MaxLength=255
	// CustomDomain Serve the web frontend using a custom CNAME pointing to the Aiven DNS name
	CustomDomain *string `groups:"create,update" json:"custom_domain,omitempty"`

	// +kubebuilder:validation:MaxItems=1024
	// IpFilter Allow incoming connections from CIDR address block, e.g. '10.20.0.0/16'
	IpFilter []*IpFilter `groups:"create,update" json:"ip_filter,omitempty"`

	// Kafka broker configuration values
	Kafka *Kafka `groups:"create,update" json:"kafka,omitempty"`

	// KafkaAuthenticationMethods Kafka authentication methods
	KafkaAuthenticationMethods *KafkaAuthenticationMethods `groups:"create,update" json:"kafka_authentication_methods,omitempty"`

	// KafkaConnect Enable Kafka Connect service
	KafkaConnect *bool `groups:"create,update" json:"kafka_connect,omitempty"`

	// KafkaConnectConfig Kafka Connect configuration values
	KafkaConnectConfig *KafkaConnectConfig `groups:"create,update" json:"kafka_connect_config,omitempty"`

	// KafkaRest Enable Kafka-REST service
	KafkaRest *bool `groups:"create,update" json:"kafka_rest,omitempty"`

	// KafkaRestAuthorization Enable authorization in Kafka-REST service
	KafkaRestAuthorization *bool `groups:"create,update" json:"kafka_rest_authorization,omitempty"`

	// KafkaRestConfig Kafka REST configuration
	KafkaRestConfig *KafkaRestConfig `groups:"create,update" json:"kafka_rest_config,omitempty"`

	// +kubebuilder:validation:Enum="2.8";"3.0";"3.1";"3.2";"3.3"
	// KafkaVersion Kafka major version
	KafkaVersion *string `groups:"create,update" json:"kafka_version,omitempty"`

	// PrivateAccess Allow access to selected service ports from private networks
	PrivateAccess *PrivateAccess `groups:"create,update" json:"private_access,omitempty"`

	// PrivatelinkAccess Allow access to selected service components through Privatelink
	PrivatelinkAccess *PrivatelinkAccess `groups:"create,update" json:"privatelink_access,omitempty"`

	// PublicAccess Allow access to selected service ports from the public Internet
	PublicAccess *PublicAccess `groups:"create,update" json:"public_access,omitempty"`

	// SchemaRegistry Enable Schema-Registry service
	SchemaRegistry *bool `groups:"create,update" json:"schema_registry,omitempty"`

	// SchemaRegistryConfig Schema Registry configuration
	SchemaRegistryConfig *SchemaRegistryConfig `groups:"create,update" json:"schema_registry_config,omitempty"`

	// StaticIps Use static public IP addresses
	StaticIps *bool `groups:"create,update" json:"static_ips,omitempty"`
}
