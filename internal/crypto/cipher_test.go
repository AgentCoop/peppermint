package crypto_test

import (
	"fmt"
	"github.com/AgentCoop/peppermint/internal/crypto"
	"github.com/AgentCoop/go-work"

	"encoding/hex"
	"testing"
	"bytes"
	//"fmt"
)

func cipherTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")

		data := []byte("send reinforcements, we're going to advance")
		original := make([]byte, len(data))
		copy(original, data)

		encryptor := crypto.NewSymCipher(key, nil, task)
		ciphertext := encryptor.Encrypt(data)
		fmt.Printf("%s %s\n", data, ciphertext)

		decryptor := crypto.NewSymCipher(key, encryptor.GetNonce(), task)
		plaintext := decryptor.Decrypt(ciphertext)

		T := j.GetValue().(*testing.T)
		if bytes.Compare(plaintext, original) != 0 {
			T.Fatalf("expected %s, got: %s", original, plaintext)
		}
		task.Done()
	}
	return nil, run, nil
}

func TestSymCipher(t *testing.T) {
	j := job.NewJob(t)
	j.AddTask(cipherTask)
	<-j.Run()
	_, err := j.GetInterruptedBy()
	if err != nil {
		t.Error(err)
	}
}
