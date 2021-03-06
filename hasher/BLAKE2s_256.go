package hasher

import (
	"HashDehash/db"
	"fmt"
	"golang.org/x/crypto/blake2s"
)

func (h *Hasher) BLAKE2s_256() {
	h.wg.Add(1)

	go func() {
		defer h.wg.Done()

		var err error
		dbToFrom := db.NewDB("hash_db/BLAKE2b_256/to_from")
		dbFromTo := db.NewDB("hash_db/BLAKE2b_256/from_to")

		h.hashes.BLAKE2s_256, err = dbFromTo.Get(h.input)
		if err == nil {
			return
		}

		h.hashes.BLAKE2s_256 = fmt.Sprintf("%x",
			blake2s.Sum256([]byte(h.input)))

		dbToFrom.Set(h.hashes.BLAKE2s_256, h.input)
		dbFromTo.Set(h.input, h.hashes.BLAKE2s_256)
	}()
}
