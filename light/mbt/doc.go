// Package mbt provides a test runner for model-based tests
//
// Model-based tests are generated by
// https://github.com/informalsystems/tendermint-rs/tree/master/testgen, which
// first turns TLA+ specifications into test scenarious. Those test scenarious
// are then in turn used to generate actual fixtures representing light blocks.
//
// The test runner initializes the light client with a trusted light block. For
// each next light block, it tries to verify the block and asserts the outcome
// ("verdict" field in .json files).
package mbt
