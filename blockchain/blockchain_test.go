package blockchain

import (
	"reflect"
	"testing"
)

func TestBlockchain_createGenesisBlock(t *testing.T) {
	type fields struct {
		blockNumber int
		chain       chain
		difficulty  int
		powPrefix   string
	}
	tests := []struct {
		name   string
		fields fields
		want   block
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Blockchain{
				blockNumber: tt.fields.blockNumber,
				chain:       tt.fields.chain,
				difficulty:  tt.fields.difficulty,
				powPrefix:   tt.fields.powPrefix,
			}
			if got := b.createGenesisBlock(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Blockchain.createGenesisBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlockchain_lastBlock(t *testing.T) {
	type fields struct {
		blockNumber int
		chain       chain
		difficulty  int
		powPrefix   string
	}
	tests := []struct {
		name   string
		fields fields
		want   block
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Blockchain{
				blockNumber: tt.fields.blockNumber,
				chain:       tt.fields.chain,
				difficulty:  tt.fields.difficulty,
				powPrefix:   tt.fields.powPrefix,
			}
			if got := b.lastBlock(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Blockchain.lastBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlockchain_Chain(t *testing.T) {
	type fields struct {
		blockNumber int
		chain       chain
		difficulty  int
		powPrefix   string
	}
	tests := []struct {
		name   string
		fields fields
		want   []block
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Blockchain{
				blockNumber: tt.fields.blockNumber,
				chain:       tt.fields.chain,
				difficulty:  tt.fields.difficulty,
				powPrefix:   tt.fields.powPrefix,
			}
			if got := b.Chain(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Blockchain.Chain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlockchain_getPreviousBlockHash(t *testing.T) {
	type fields struct {
		blockNumber int
		chain       chain
		difficulty  int
		powPrefix   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Blockchain{
				blockNumber: tt.fields.blockNumber,
				chain:       tt.fields.chain,
				difficulty:  tt.fields.difficulty,
				powPrefix:   tt.fields.powPrefix,
			}
			if got := b.getPreviousBlockHash(); got != tt.want {
				t.Errorf("Blockchain.getPreviousBlockHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlockchain_CreatePayload(t *testing.T) {
	type fields struct {
		blockNumber int
		chain       chain
		difficulty  int
		powPrefix   string
	}
	type args struct {
		data string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   payload
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Blockchain{
				blockNumber: tt.fields.blockNumber,
				chain:       tt.fields.chain,
				difficulty:  tt.fields.difficulty,
				powPrefix:   tt.fields.powPrefix,
			}
			if got := b.CreatePayload(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Blockchain.CreatePayload() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlockchain_Mine(t *testing.T) {
	type fields struct {
		blockNumber int
		chain       chain
		difficulty  int
		powPrefix   string
	}
	type args struct {
		p payload
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   minedBlock
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Blockchain{
				blockNumber: tt.fields.blockNumber,
				chain:       tt.fields.chain,
				difficulty:  tt.fields.difficulty,
				powPrefix:   tt.fields.powPrefix,
			}
			if got := b.Mine(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Blockchain.Mine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlockchain_verifyBlock(t *testing.T) {
	type fields struct {
		blockNumber int
		chain       chain
		difficulty  int
		powPrefix   string
	}
	type args struct {
		bl block
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Blockchain{
				blockNumber: tt.fields.blockNumber,
				chain:       tt.fields.chain,
				difficulty:  tt.fields.difficulty,
				powPrefix:   tt.fields.powPrefix,
			}
			if got := b.verifyBlock(tt.args.bl); got != tt.want {
				t.Errorf("Blockchain.verifyBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlockchain_PushBlock(t *testing.T) {
	type fields struct {
		blockNumber int
		chain       chain
		difficulty  int
		powPrefix   string
	}
	type args struct {
		bl block
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   chain
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Blockchain{
				blockNumber: tt.fields.blockNumber,
				chain:       tt.fields.chain,
				difficulty:  tt.fields.difficulty,
				powPrefix:   tt.fields.powPrefix,
			}
			if got := b.PushBlock(tt.args.bl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Blockchain.PushBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		difficulty  int
		blockNumber int
	}
	tests := []struct {
		name string
		args args
		want Blockchain
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.difficulty, tt.args.blockNumber); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
