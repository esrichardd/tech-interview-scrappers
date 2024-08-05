package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func InitializeDataHandler(w http.ResponseWriter, r *http.Request) {
	leagueResp, err := fetchFromHandler("/league")
	if err != nil {
		http.Error(w, "Error al obtener datos de la liga: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := postToEndpoint("http://tech-interview-core:3002/tournaments/with-associated-data", leagueResp); err != nil {
		http.Error(w, "Error al enviar datos de la liga: "+err.Error(), http.StatusInternalServerError)
		return
	}

	matchesResp, err := fetchFromHandler("/league/matches")
	if err != nil {
		http.Error(w, "Error al obtener datos de los partidos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := postToEndpoint("http://tech-interview-core:3002/tournaments/games", matchesResp); err != nil {
		http.Error(w, "Error al enviar datos de los partidos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Datos inicializados correctamente"))
}

func fetchFromHandler(path string) ([]byte, error) {
	resp, err := http.Get("http://localhost:3003" + path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error en la solicitud GET: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer el cuerpo de la respuesta: %v", err)
	}

	return data, nil
}

func postToEndpoint(url string, data []byte) error {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("error en la solicitud POST: %s", resp.Status)
	}

	return nil
}
