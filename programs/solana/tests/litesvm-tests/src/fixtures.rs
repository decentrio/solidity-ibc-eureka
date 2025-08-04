//! Test fixtures and data for IBC tests

use ics26_router::state::CounterpartyInfo;

/// Create a test client ID
pub fn test_client_id(index: u8) -> String {
    format!("test-client-{:02}", index)
}

/// Create test counterparty info
pub fn test_counterparty_info() -> CounterpartyInfo {
    CounterpartyInfo {
        client_id: "counterparty-client".to_string(),
        connection_id: "connection-0".to_string(),
        merkle_prefix: vec![0x01, 0x02, 0x03],
    }
}

/// Test light client update data
pub struct TestClientUpdate {
    pub header: Vec<u8>,
    pub trusted_height: u64,
    pub current_timestamp: i64,
}

impl TestClientUpdate {
    pub fn new() -> Self {
        Self {
            header: vec![0; 32], // Mock header
            trusted_height: 100,
            current_timestamp: 1700000000,
        }
    }
}

/// Test packet data
pub struct TestPacket {
    pub sequence: u64,
    pub source_port: String,
    pub source_channel: String,
    pub destination_port: String,
    pub destination_channel: String,
    pub data: Vec<u8>,
    pub timeout_height: u64,
    pub timeout_timestamp: u64,
}

impl TestPacket {
    pub fn new(sequence: u64) -> Self {
        Self {
            sequence,
            source_port: "transfer".to_string(),
            source_channel: "channel-0".to_string(),
            destination_port: "transfer".to_string(),
            destination_channel: "channel-1".to_string(),
            data: b"test packet data".to_vec(),
            timeout_height: 1000,
            timeout_timestamp: 0,
        }
    }
}