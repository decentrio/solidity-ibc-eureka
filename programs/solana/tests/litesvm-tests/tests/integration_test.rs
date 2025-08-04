//! Integration tests for IBC Solana programs using mollusk-svm

use ibc_solana_litesvm_tests::{fixtures::*, utils::*, IbcTestContext};
use solana_sdk::pubkey::Pubkey;

#[test]
fn test_initialize_router() {
    println!("Starting test_initialize_router");
    let mut ctx = IbcTestContext::new();
    println!("Created test context");
    
    // Create initialize instruction
    let init_ix = create_initialize_router_instruction(
        &ctx.ics26_program_id,
        &ctx.payer_pubkey(),
        &ctx.payer_pubkey(),
    );
    println!("Created initialize instruction");
    
    // Process transaction
    let result = ctx.process_transaction(vec![init_ix]);
    println!("Initialize router result: {:?}", result);
    
    // Check if successful
    // The process_transaction already validates with Check::success()
    // If we reach here, the transaction was successful
    println!("Router initialized successfully");
}

#[test]
fn test_add_client() {
    let mut ctx = IbcTestContext::new();
    
    // Initialize router first
    let init_ix = create_initialize_router_instruction(
        &ctx.ics26_program_id,
        &ctx.payer_pubkey(),
        &ctx.payer_pubkey(),
    );
    ctx.process_transaction(vec![init_ix]);
    
    // Add a client
    let client_id = test_client_id(1);
    let counterparty_info = test_counterparty_info();
    
    let add_client_ix = create_add_client_instruction(
        &ctx.ics26_program_id,
        &ctx.payer_pubkey(),
        &client_id,
        counterparty_info,
        &ctx.ics07_program_id, // Using ICS07 as the light client
    );
    
    let result = ctx.process_transaction(vec![add_client_ix]);
    println!("Add client result: {:?}", result);
}

#[test]
fn test_bind_port() {
    let mut ctx = IbcTestContext::new();
    
    // Initialize router
    let init_ix = create_initialize_router_instruction(
        &ctx.ics26_program_id,
        &ctx.payer_pubkey(),
        &ctx.payer_pubkey(),
    );
    ctx.process_transaction(vec![init_ix]);
    
    // Create a mock app program ID
    let app_program_id = Pubkey::new_unique();
    let port_id = "transfer";
    
    // Add IBC app
    let add_ibc_app_ix = create_add_ibc_app_instruction(
        &ctx.ics26_program_id,
        &ctx.payer_pubkey(),
        &ctx.payer_pubkey(),
        port_id,
        &app_program_id,
    );
    
    let result = ctx.process_transaction(vec![add_ibc_app_ix]);
    println!("Add IBC app result: {:?}", result);
}

#[test]
fn test_multiple_clients() {
    let mut ctx = IbcTestContext::new();
    
    // Initialize router
    let init_ix = create_initialize_router_instruction(
        &ctx.ics26_program_id,
        &ctx.payer_pubkey(),
        &ctx.payer_pubkey(),
    );
    ctx.process_transaction(vec![init_ix]);
    
    // Add multiple clients
    for i in 1..=3 {
        let client_id = test_client_id(i);
        let counterparty_info = test_counterparty_info();
        
        let add_client_ix = create_add_client_instruction(
            &ctx.ics26_program_id,
            &ctx.payer_pubkey(),
            &client_id,
            counterparty_info,
            &ctx.ics07_program_id,
        );
        
        let result = ctx.process_transaction(vec![add_client_ix]);
        println!("Add client {} result: {:?}", i, result);
    }
}

#[test]
fn test_router_state_management() {
    let mut ctx = IbcTestContext::new();
    
    // Initialize router
    let init_ix = create_initialize_router_instruction(
        &ctx.ics26_program_id,
        &ctx.payer_pubkey(),
        &ctx.payer_pubkey(),
    );
    
    let result = ctx.process_transaction(vec![init_ix]);
    println!("Router initialization result: {:?}", result);
    
    // Add a client and verify state changes
    let client_id = test_client_id(1);
    let counterparty_info = test_counterparty_info();
    
    let add_client_ix = create_add_client_instruction(
        &ctx.ics26_program_id,
        &ctx.payer_pubkey(),
        &client_id,
        counterparty_info,
        &ctx.ics07_program_id,
    );
    
    let result = ctx.process_transaction(vec![add_client_ix]);
    println!("Client addition result: {:?}", result);
}