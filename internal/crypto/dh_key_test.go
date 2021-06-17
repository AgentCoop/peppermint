package crypto_test

import (
	"github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/crypto"

	"bytes"
	"testing"
)

func keyExchangeTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		// Alice's side
		dhKeyAlice, err := crypto.NewKeyExchange()
		task.Assert(err)

		alicePubKey := dhKeyAlice.GetPublicKey()
		// Bob's side
		dhKeyBob, err := crypto.NewKeyExchange()
		task.Assert(err)

		bobPubKey := dhKeyBob.GetPublicKey()
		aliceKey, err := dhKeyAlice.ComputeKey(bobPubKey)
		task.Assert(err)

		bobKey, err := dhKeyBob.ComputeKey(alicePubKey)
		task.Assert(err)

		T := task.GetJob().GetValue().(*testing.T)
		if bytes.Compare(aliceKey, bobKey) != 0 {
			T.Fatalf("alice's secret key %x, bob's secret key %x\n", aliceKey, bobKey)
		}
		task.Done()
	}
	return nil, run, nil
}

func TestKeyExchange(t *testing.T) {
	j := job.NewJob(t)
	j.AddTask(keyExchangeTask)
	<-j.Run()
	_, err := j.GetInterruptedBy()
	if err != nil {
		t.Error(err)
	}
}
