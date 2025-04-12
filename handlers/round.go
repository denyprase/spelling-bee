package handlers

import (
	"fmt"
	"net/http"
	"spelling-bee/utils"
	"strconv"
)

func (h *AppHandler) RoundDetailHandler(w http.ResponseWriter, r *http.Request) {
	roundIDStr := r.URL.Query().Get("id")
	roundID, err := strconv.Atoi(roundIDStr)
	if err != nil {
		fmt.Println("Error converting round ID:", err.Error())
		http.Error(w, "Invalid round ID", http.StatusBadRequest)
		return
	}

	round, err := h.DB.GetRoundByID(roundID)
	if err != nil {
		fmt.Println("Error fetching round:", err.Error())
		http.NotFound(w, r)
		return
	}

	words, err := h.DB.GetWordsByRoundID(round.ID)
	if err != nil {
		fmt.Println("Error fetching words:", err.Error())
		http.NotFound(w, r)
		return
	}

	data := map[string]interface{}{
		"Round": round,
		"Words": words,
	}

	utils.RenderTemplate(w, r, "round-detail.html", data)

}
