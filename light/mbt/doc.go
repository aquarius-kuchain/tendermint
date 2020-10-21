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
//
// In the first version (v1), JSON files are directly added to the repo. In
// the future (v2), they will be generated by the testgen binary right before
// testing on CI (the number of files will be around thousands).
//
// NOTE (v1): If a breaking change is introduced into the SignedHeader or
// ValidatorSet, you will need to regenerate the JSON files using testgen
// binary (may also require modifying tendermint-rs, e.g.
// https://github.com/informalsystems/tendermint-rs/pull/647)
package mbt
