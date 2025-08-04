//! Integration test utilities for IBC Solana programs

use mollusk_svm::{Mollusk, result::Check};
use solana_sdk::{
    account::Account,
    pubkey::Pubkey,
    signature::{Keypair, Signer},
    instruction::Instruction,
    native_loader,
    system_program as system_program_mod,
};

pub mod fixtures;
pub mod utils;

/// Test context for IBC programs
pub struct IbcTestContext {
    pub mollusk: Mollusk,
    pub payer: Keypair,
    pub ics07_program_id: Pubkey,
    pub ics26_program_id: Pubkey,
}

impl IbcTestContext {
    /// Create a new test context with deployed programs
    pub fn new() -> Self {
        let payer = Keypair::new();
        
        // Initialize Mollusk with the ICS26 router program
        let ics26_program_id = ics26_router::ID;
        // Use relative path from test directory
        let program_path = "../../../../target/deploy/ics26_router";
        let mollusk = Mollusk::new(&ics26_program_id, program_path);
        
        // Use the actual program IDs
        let ics07_program_id = ics07_tendermint::ID;

        Self {
            mollusk,
            payer,
            ics07_program_id,
            ics26_program_id,
        }
    }
    
    /// Get the payer's pubkey
    pub fn payer_pubkey(&self) -> Pubkey {
        self.payer.pubkey()
    }

    /// Process a single instruction
    pub fn process_instruction(
        &mut self,
        instruction: &Instruction,
    ) -> mollusk_svm::result::InstructionResult {
        // Set up accounts based on instruction
        let mut accounts = vec![
            (
                self.payer_pubkey(),
                Account {
                    lamports: 1_000_000_000,
                    data: vec![],
                    owner: system_program_mod::ID,
                    executable: false,
                    rent_epoch: 0,
                },
            ),
        ];

        // Add accounts referenced in the instruction
        for account_meta in &instruction.accounts {
            // Skip if we already have this account
            if accounts.iter().any(|(pubkey, _)| pubkey == &account_meta.pubkey) {
                continue;
            }

            // Add account based on its type
            if account_meta.pubkey == system_program_mod::ID {
                accounts.push((
                    system_program_mod::ID,
                    Account {
                        lamports: 0,
                        data: vec![],
                        owner: native_loader::ID,
                        executable: true,
                        rent_epoch: 0,
                    },
                ));
            } else if account_meta.pubkey == self.ics07_program_id {
                accounts.push((
                    self.ics07_program_id,
                    Account {
                        lamports: 0,
                        data: vec![],
                        owner: native_loader::ID,
                        executable: true,
                        rent_epoch: 0,
                    },
                ));
            } else {
                // Default empty account for PDAs
                accounts.push((
                    account_meta.pubkey,
                    Account {
                        lamports: 0,
                        data: vec![],
                        owner: system_program_mod::ID,
                        executable: false,
                        rent_epoch: 0,
                    },
                ));
            }
        }

        // Process with basic success check
        let checks = vec![Check::success()];
        
        self.mollusk.process_and_validate_instruction(instruction, &accounts, &checks)
    }

    /// Process multiple instructions
    pub fn process_transaction(
        &mut self,
        instructions: Vec<Instruction>,
    ) -> mollusk_svm::result::InstructionResult {
        let mut last_result = mollusk_svm::result::InstructionResult::default();
        
        for instruction in instructions {
            last_result = self.process_instruction(&instruction);
            if last_result.program_result.is_err() {
                break;
            }
        }
        
        last_result
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_context_creation() {
        let ctx = IbcTestContext::new();
        assert_eq!(ctx.ics07_program_id, ics07_tendermint::ID);
        assert_eq!(ctx.ics26_program_id, ics26_router::ID);
    }
}