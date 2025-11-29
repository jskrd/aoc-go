package aoc

import (
	"fmt"
	"io"
	"net/http"
)

func FetchInput(session string, year int, day int, level int) (string, error) {
	request, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day),
		nil,
	)
	request.Header.Set("Cookie", "session="+session)

	response, error := http.DefaultClient.Do(request)
	if error != nil {
		return "", fmt.Errorf("Failed to fetch input: %w", error)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	if response.StatusCode != 200 {
		return "", fmt.Errorf("Failed to fetch input: %s", string(body))
	}

	return string(body), nil
}
