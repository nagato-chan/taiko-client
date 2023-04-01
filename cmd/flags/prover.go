package flags

import (
	"github.com/urfave/cli/v2"
)

// Required flags used by prover.
var (
	ZkEvmRpcdEndpoint = &cli.StringFlag{
		Name:     "zkevmRpcdEndpoint",
		Usage:    "RPC endpoint of a ZKEVM RPCD service",
		Required: true,
		Category: proverCategory,
	}
	ZkEvmRpcdParamsPath = &cli.StringFlag{
		Name:     "zkevmRpcdParamsPath",
		Usage:    "Path of ZKEVM parameters file to use",
		Required: true,
		Category: proverCategory,
	}
	L1ProverPrivKey = &cli.StringFlag{
		Name: "l1.proverPrivKey",
		Usage: "Private key of L1 prover, " +
			"who will send TaikoL1.proveBlock / TaikoL1.proveBlockInvalid transactions",
		Required: true,
		Category: proverCategory,
	}
)

// Optional flags used by prover.
var (
	StartingBlockID = &cli.Uint64Flag{
		Name:     "startingBlockID",
		Usage:    "If set, prover will start proving blocks from the block with this ID",
		Category: proverCategory,
	}
	MaxConcurrentProvingJobs = &cli.UintFlag{
		Name:     "maxConcurrentProvingJobs",
		Usage:    "Limits the number of concurrent proving blocks jobs",
		Value:    1,
		Category: proverCategory,
	}
	// Special flags for testing.
	Dummy = &cli.BoolFlag{
		Name:     "dummy",
		Usage:    "Produce dummy proofs, testing purposes only",
		Value:    false,
		Category: proverCategory,
	}
	RandomDummyProofDelay = &cli.StringFlag{
		Name: "randomDummyProofDelay",
		Usage: "Set the random dummy proof delay between the bounds using the format: " +
			"`lowerBound-upperBound` (e.g. `30m-1h`), testing purposes only",
		Category: proverCategory,
	}
	OnlyProveOddNumberBlocks = &cli.BoolFlag{
		Name:     "onlyProveOddNumberBlocks",
		Usage:    "Let the prover only proves blocks with odd number",
		Value:    false,
		Category: proverCategory,
	}
	OnlyProveEvenNumberBlocks = &cli.BoolFlag{
		Name:     "onlyProveEvenNumberBlocks",
		Usage:    "Let the prover only proves blocks with even number",
		Value:    false,
		Category: proverCategory,
	}
	ProofNumberRate = &cli.Uint64Flag{
		Name:     "proofNumberRate",
		Usage:    "The rate of proof number to verify",
		Value:  0,
		Category: proverCategory,
	}
	ProofSubmittorPrivKey = &cli.StringFlag{
		Name:     "proofSubmittorPrivKey",
		Usage:    "Private key of L1 proof submittor",
		Category: proverCategory,
	}
)

// All prover flags.
var ProverFlags = MergeFlags(CommonFlags, []cli.Flag{
	L1HTTPEndpoint,
	L2WSEndpoint,
	L2HTTPEndpoint,
	ZkEvmRpcdEndpoint,
	ZkEvmRpcdParamsPath,
	L1ProverPrivKey,
	ProofSubmittorPrivKey,
	StartingBlockID,
	MaxConcurrentProvingJobs,
	Dummy,
	RandomDummyProofDelay,
	OnlyProveOddNumberBlocks,
	OnlyProveEvenNumberBlocks,
})
