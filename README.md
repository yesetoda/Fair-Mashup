# Fair-Mashup 🎯

Welcome to **Fair-Mashup** – the ultimate tool for creating fair and balanced Codeforces mashup contests! 🎉

---

## 🌟 What is Fair-Mashup?

**Fair-Mashup** is a simple yet powerful web-based tool designed to help you build Codeforces mashup contests. The twist? None of the participants in the contest will have solved the selected problems before! 🔥

Whether you're a competitive programmer, a trainer, or just a fan of mashups, this tool helps you select problems that match your difficulty preferences and tags, ensuring a level playing field for all participants. No more sneaky participants with unfair advantages! 😉

---

## 🚀 Features

- **Participant Filtering:** Add the participants’ usernames and make sure the selected problems haven’t been solved by any of them.
- **Problem Tags:** Choose problems based on specific tags like *dp*, *graphs*, *greedy*, and many more.
- **Difficulty Range:** Set a minimum and maximum difficulty level to target problems that suit your contest's skill range.
- **Problem Links:** Get direct links to unsolved problems from Codeforces – just one click away!
- **Mashup-ready:** Once the problems are generated, you can use them to create a Codeforces mashup directly.

---

## 🛠️ How It Works

1. Enter the participants' usernames.
2. Select the tags you're interested in.
3. Define a difficulty range (minimum and maximum problem ratings).
4. Hit that submit button! 🎯
5. Get a list of unsolved problems tailored to your mashup.
6. Click on the links and open the problems directly on Codeforces.

---

## 🧑‍💻 What's Under the Hood?

**Fair-Mashup** is built with:
- **Go** 🐹 – Our backend that handles the heavy lifting of fetching data from the Codeforces API and filtering it based on participants’ solved problems.
- **HTML/CSS/JavaScript** 🎨 – A clean and responsive frontend interface that makes it super easy to input data and view the results.
- **No external dependencies** 📦 – We like to keep things simple. No need for Node.js or any other package managers. Just pure Go, HTML, CSS, and JavaScript.

---

## 🛠️ Installation

Getting started is super easy:

1. **Clone the repo**:
   ```bash
   git clone https://github.com/yesetoda/Fair-Mashup.git
   ```
   
2. **Run the Go server**:
   ```bash
   go run main.go
   ```

3. **Open the HTML page**:  
   Open the `index.html` in your favorite browser, and you're ready to go!

4. **Enjoy crafting your contest!**

---

## 🎨 Interface

Once you load the page, you’ll find a simple form where you:
- Enter participant handles
- Select problem tags
- Set the difficulty range

After hitting submit, you’ll get a neatly formatted table with:
- Problem name
- Difficulty
- Tags (displayed as bullet points 🔵 for easy readability)
- Contest ID and index
- A link to solve the problem on Codeforces (because time is money, right? 💸)

---

## 💡 Use Cases

- **Trainers**: Ensure no participant has prior experience with the problems.
- **Friends**: Want to have a friendly contest? This tool ensures fair play!
- **Competitive Programming Groups**: Plan a contest that is perfectly balanced between all participants.

---

## 🔧 Future Improvements

Some cool features we’re thinking about:
- Adding more filtering options (like contest types or time constraints).
- Adding support for multiple programming contest platforms (Codeforces is just the beginning!).
- Dynamic difficulty adjustment based on participants’ previous performances.

---

## 🤝 Contributing

Found a bug? Have a cool feature idea? Pull requests are welcome! You can also open an issue if you want to discuss anything related to the project. Let’s build something awesome together!

---


Happy mashup-ing! 🎉