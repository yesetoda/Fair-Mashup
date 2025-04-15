# Fair-Mashup

[![Go Version](https://img.shields.io/github/go-mod/go-version/yesetoda/fair-mashup?style=flat-square)](https://golang.org/) [![Build Status](https://img.shields.io/github/actions/workflow/status/yesetoda/fair-mashup/go.yml?branch=main&style=flat-square)](https://github.com/yesetoda/fair-mashup/actions) [![GitHub issues](https://img.shields.io/github/issues/yesetoda/fair-mashup?style=flat-square)](https://github.com/yesetoda/fair-mashup/issues) [![GitHub stars](https://img.shields.io/github/stars/yesetoda/fair-mashup?style=flat-square)](https://github.com/yesetoda/fair-mashup/stargazers) [![Live Demo](https://img.shields.io/badge/Live-Demo-green?style=flat-square)](https://fair-mashup.onrender.com/)

Welcome to **Fair-Mashup** – the ultimate tool for creating fair and balanced Codeforces mashup contests!

---

## Overview

**Fair-Mashup** is a simple yet powerful web-based application built entirely in Go. It enables you to build Codeforces mashup contests where none of the participants will have solved—or even attempted—the selected problems before! This approach guarantees a level playing field, ensuring that your contests remain both challenging and impartial.

---

## How It Works

1. **Participants:**  
   Enter Codeforces handles (comma separated) for the contestants.  
   **Example:** `user1, user2, user3`

2. **Tags and Difficulty:**  
   - **Select Tags:** Filter problems by specific topics or themes.  
   - **Difficulty Range:** Specify minimum and maximum difficulty levels to curate your contest problems.

3. **Submit:**  
   After configuring the participants, tags, and difficulty range, click **Submit** to generate a contest where each problem is new to all contestants.

---

## Technologies Used

- **Go (Golang):**  
  Powers both the backend and the server-side rendered frontend using Go's built-in HTTP server and templating system.

- **HTML & CSS:**  
  Used to build and style static pages, ensuring a clean and responsive user interface.

---

## Installation and Setup

### Prerequisites

- **Go 1.18+** must be installed on your machine.

### Steps

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/yesetoda/fair-mashup.git
   cd fair-mashup
   ```
2. **Install Dependencies:**
   Update Go modules:
   ```bash
   go mod tidy
   ```
3. **Run the Application:**
   ```bash
   go run main.go
   ```
4. **Access the Application:**
   Open your browser and navigate to:
   ```
   http://localhost:8080
   ```
   Then, enter participant handles, select tags and difficulty ranges, and generate your contest.

---

## Contribution Guidelines

Contributions are welcome! To contribute:

1. **Fork the Repository.**
2. **Create a New Branch:**
   ```bash
   git checkout -b feature-branch-name
   ```
3. **Commit Your Changes:**
   ```bash
   git commit -m "Describe your feature or fix"
   ```
4. **Push Your Branch:**
   ```bash
   git push origin feature-branch-name
   ```
5. **Open a Pull Request** on GitHub with details about your changes.

---

## Testing

Fair-Mashup is a self-contained web application. To test manually:

- Run the application:
  ```bash
  go run main.go
  ```
- Verify participant input, tag/difficulty selection, and contest generation in your browser.

For automated testing, consider adding tests using Go’s `testing` package.

---

## Live Demo

Explore the live version of Fair-Mashup here:  
[Live Demo](https://fair-mashup.onrender.com/)

---

## Contact Information

- **Name:** Yeneineh Seiba  
- **GitHub:** [yesetoda](https://github.com/yesetoda)  
- **LinkedIn:** [yeneineh (yesetoda) seiba](https://www.linkedin.com/in/yeneineh-seiba-88110227b/)  
- **Email:** [yeneinehseiba@gmail.com](mailto:yeneinehseiba@gmail.com)

---

© 2024 Fair-Mashup. All rights reserved.
