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
		dhKeyAlice := crypto.NewKeyExchange(task)
		alicePubKey := dhKeyAlice.GetPublicKey()

		// Bob's side
		dhKeyBob := crypto.NewKeyExchange(task)
		bobPubKey := dhKeyBob.GetPublicKey()

		aliceKey := dhKeyAlice.ComputeKey(bobPubKey)
		bobKey := dhKeyBob.ComputeKey(alicePubKey)

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
