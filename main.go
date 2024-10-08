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
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/problems", problemsHandler)
	http.HandleFunc("/api/tags", tagsHandler)

	fmt.Println("Server started at fair-mashup.onrender.com")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func problemsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	fmt.Printf("Received request: %s %s\n", r.Method, r.URL)

	var requestData struct {
		Participants  []string `json:"participants"`
		Tags          []string `json:"tags"`
		MinDifficulty int      `json:"minDifficulty"`
		MaxDifficulty int      `json:"maxDifficulty"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Request Data: %+v\n", requestData)

	solved, tried := getSolvedAndTriedProblems(requestData.Participants)

	allProblems := getProblemsByTags(requestData.Tags)

	validProblems := filterValidProblems(allProblems, solved, tried, requestData.MinDifficulty, requestData.MaxDifficulty)

	refinedProblems := validProblems[:min(len(validProblems), 10)]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"validProblems": refinedProblems,
	})
}

func tagsHandler(w http.ResponseWriter, r *http.Request) {
	tags := []string{
		"2-sat",
		"binary search",
		"bitmasks",
		"brute force",
		"combinatorics",
		"constructive algorithms",
		"data structures",
		"dfs and similar",
		"divide and conquer",
		"dp",
		"dsu",
		"expression parsing",
		"fft",
		"flows",
		"games",
		"geometry",
		"graphs",
		"greedy",
		"hashing",
		"implementation",
		"interactive",
		"math",
		"matrices",
		"meet-in-the-middle",
		"number theory",
		"probabilities",
		"schedules",
		"shortest paths",
		"sortings",
		"string suffix structures",
		"strings",
		"ternary search",
		"trees",
		"two pointers",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tags)
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
	ContestID int `json:"contestId"`
	Problem   struct {
		Index string `json:"index"`
	} `json:"problem"`
	Verdict string `json:"verdict"`
}

type StatusResponse struct {
	Status string      `json:"status"`
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
