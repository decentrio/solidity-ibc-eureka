//! Utility functions for IBC tests

use anchor_lang::InstructionData;
use solana_program::instruction::{AccountMeta, Instruction};
use solana_sdk::pubkey::Pubkey;
use ics26_router::state::CounterpartyInfo;

// PDAs
const ROUTER_STATE_SEED: &[u8] = b"router_state";
const CLIENT_SEED: &[u8] = b"client";
const CLIENT_SEQUENCE_SEED: &[u8] = b"client_sequence";
const IBC_APP_SEED: &[u8] = b"ibc_app";

/// Find PDA for router state
pub fn find_router_state_pda(program_id: &Pubkey) -> (Pubkey, u8) {
    Pubkey::find_program_address(
        &[ROUTER_STATE_SEED],
        program_id,
    )
}

/// Find PDA for client
pub fn find_client_pda(program_id: &Pubkey, client_id: &str) -> (Pubkey, u8) {
    Pubkey::find_program_address(
        &[CLIENT_SEED, client_id.as_bytes()],
        program_id,
    )
}

/// Find PDA for client sequence
pub fn find_client_sequence_pda(program_id: &Pubkey, client_id: &str) -> (Pubkey, u8) {
    Pubkey::find_program_address(
        &[CLIENT_SEQUENCE_SEED, client_id.as_bytes()],
        program_id,
    )
}

/// Find PDA for IBC app
pub fn find_ibc_app_pda(program_id: &Pubkey, port_id: &str) -> (Pubkey, u8) {
    Pubkey::find_program_address(
        &[IBC_APP_SEED, port_id.as_bytes()],
        program_id,
    )
}

/// Create initialize router instruction
pub fn create_initialize_router_instruction(
    program_id: &Pubkey,
    payer: &Pubkey,
    authority: &Pubkey,
) -> Instruction {
    let (router_state_pda, _) = find_router_state_pda(program_id);
    
    let accounts = vec![
        AccountMeta::new(router_state_pda, false),
        AccountMeta::new(*payer, true),
        AccountMeta::new_readonly(solana_program::system_program::ID, false),
    ];

    let data = ics26_router::instruction::Initialize { 
        authority: *authority 
    }.data();

    Instruction {
        program_id: *program_id,
        accounts,
        data,
    }
}

/// Create add client instruction
pub fn create_add_client_instruction(
    program_id: &Pubkey,
    authority: &Pubkey,
    client_id: &str,
    counterparty_info: CounterpartyInfo,
    light_client_program: &Pubkey,
) -> Instruction {
    let (router_state_pda, _) = find_router_state_pda(program_id);
    let (client_pda, _) = find_client_pda(program_id, client_id);
    let (client_sequence_pda, _) = find_client_sequence_pda(program_id, client_id);
    
    let accounts = vec![
        AccountMeta::new(*authority, true),
        AccountMeta::new_readonly(router_state_pda, false),
        AccountMeta::new(client_pda, false),
        AccountMeta::new(client_sequence_pda, false),
        AccountMeta::new_readonly(*authority, true), // relayer
        AccountMeta::new_readonly(*light_client_program, false),
        AccountMeta::new_readonly(solana_program::system_program::ID, false),
    ];

    let data = ics26_router::instruction::AddClient {
        client_id: client_id.to_string(),
        counterparty_info,
    }.data();

    Instruction {
        program_id: *program_id,
        accounts,
        data,
    }
}

/// Create add IBC app instruction
pub fn create_add_ibc_app_instruction(
    program_id: &Pubkey,
    authority: &Pubkey,
    payer: &Pubkey,
    port_id: &str,
    app_program_id: &Pubkey,
) -> Instruction {
    let (router_state_pda, _) = find_router_state_pda(program_id);
    let (ibc_app_pda, _) = find_ibc_app_pda(program_id, port_id);
    
    let accounts = vec![
        AccountMeta::new_readonly(router_state_pda, false),
        AccountMeta::new(ibc_app_pda, false),
        AccountMeta::new_readonly(*app_program_id, false),
        AccountMeta::new(*payer, true),
        AccountMeta::new_readonly(*authority, true),
        AccountMeta::new_readonly(solana_program::system_program::ID, false),
    ];

    let data = ics26_router::instruction::AddIbcApp {
        port_id: port_id.to_string(),
    }.data();

    Instruction {
        program_id: *program_id,
        accounts,
        data,
    }
}