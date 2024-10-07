package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"strings"
)

func main() {
	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handle the main page and API
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/problems", problemsHandler)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// indexHandler serves the HTML file
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

// problemsHandler processes problem requests from the frontend
func problemsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		Participants  []string `json:"participants"`
		Tags          []string `json:"tags"`
		MinDifficulty int      `json:"minDifficulty"`
		MaxDifficulty int      `json:"maxDifficulty"`
	}

	// Parse JSON from the request body
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get participants' solved and tried problems
	solved, tried := getSolvedAndTriedProblems(requestData.Participants)

	// Get all problems matching the requested tags
	allProblems := getProblemsByTags(requestData.Tags)

	// Filter problems by difficulty and ensure they haven't been solved/tried
	validProblems := filterValidProblems(allProblems, solved, tried, requestData.MinDifficulty, requestData.MaxDifficulty)

	// Limit valid problems to 10
	refinedProblems := validProblems[:min(len(validProblems), 10)]

	// Send the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"validProblems": refinedProblems,
	})
}

func getSolvedAndTriedProblems(participants []string) (map[string]bool, map[string]bool) {
	solved := make(map[string]bool)
	tried := make(map[string]bool)

	for _, handle := range participants {
		handle = strings.TrimSpace(handle)
		submissions := getSubmissions(handle)
		for _, sub := range submissions {
			link := fmt.Sprintf("https://codeforces.com/contest/%d/problem/%s", sub.ContestID, sub.Problem.Index)
			if sub.Verdict == "OK" {
				solved[link] = true
			} else {
				tried[link] = true
			}
		}
	}
	return solved, tried
}

func getSubmissions(handle string) Submissions {
	resp, err := http.Get(fmt.Sprintf("https://codeforces.com/api/user.status?handle=%s", handle))
	if err != nil {
		fmt.Println("Failed to fetch submissions:", err)
		return nil
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result StatusResponse
	json.Unmarshal(body, &result)

	if result.Status == "FAILED" {
		fmt.Printf("Handle %s incorrect\n", handle)
	}
	return result.Result
}

func getProblemsByTags(tags []string) []ProblemInfo {
	url := "https://codeforces.com/api/problemset.problems?tags=" + strings.Join(tags, ";")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching problems:", err)
		return nil
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result ProblemsResponse
	json.Unmarshal(body, &result)

	if result.Status == "FAILED" {
		fmt.Println("Incorrect tags! Please try again.")
		return nil
	}
	return result.Result.Problems
}

func filterValidProblems(problems []ProblemInfo, solved, tried map[string]bool, minDifficulty, maxDifficulty int) []ProblemInfo {
	validProblems := []ProblemInfo{}
	for _, problem := range problems {
		if problem.Rating >= minDifficulty && problem.Rating <= maxDifficulty {
			link := fmt.Sprintf("https://codeforces.com/contest/%d/problem/%s", problem.ContestID, problem.Index)
			if !solved[link] && !tried[link] {
				validProblems = append(validProblems, problem)
			}
		}
	}
	return validProblems
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Submissions []struct {
	ContestID int    `json:"contestId"`
	Problem   struct {
		Index string `json:"index"`
	} `json:"problem"`
	Verdict string `json:"verdict"`
}

type StatusResponse struct {
	Status string     `json:"status"`
	Result Submissions `json:"result"`
}

type ProblemInfo struct {
	ContestID int      `json:"contestId"`
	Index     string   `json:"index"`
	Name      string   `json:"name"`
	Rating    int      `json:"rating"`
	Tags      []string `json:"tags"`
}

type ProblemsResponse struct {
	Status string `json:"status"`
	Result struct {
		Problems []ProblemInfo `json:"problems"`
	} `json:"result"`
}
