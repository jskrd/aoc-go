package aoc

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func SubmitAnswer(
	session string,
	year int,
	day int,
	level int,
	answer int,
) error {
	request, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day),
		strings.NewReader(
			url.Values{
				"level":  {fmt.Sprintf("%d", level)},
				"answer": {fmt.Sprintf("%d", answer)},
			}.Encode(),
		),
	)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Cookie", "session="+session)

	response, error := http.DefaultClient.Do(request)
	if error != nil {
		return fmt.Errorf("Failed to submit answer: %w", error)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to submit answer: %s", string(body))
	}

	if strings.Contains(string(body), "already complete") {
		return fmt.Errorf("Already complete")
	}

	if strings.Contains(string(body), "That's the right answer!") {
		return nil
	}

	return fmt.Errorf("Unknown error: %s", string(body))
}
