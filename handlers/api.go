package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

var words = []string{
	"AMBIDEXTROUS",
	"INCONCEIVABLE",
	"HYPERINTELLIGENT",
	"BIBLIOPHILE",
}

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetRandomWordAPI(w http.ResponseWriter, r *http.Request) {
	word := words[rng.Intn(len(words))]
	resp := map[string]string{"word": word}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
