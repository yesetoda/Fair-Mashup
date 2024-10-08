document.addEventListener('DOMContentLoaded', () => {
    const tags = [
        "2-sat", "binary search", "bitmasks", "brute force", "combinatorics", 
        "constructive algorithms", "data structures", "dfs and similar", 
        "divide and conquer", "dp", "dsu", "expression parsing", "fft", 
        "flows", "games", "geometry", "graphs", "greedy", "hashing", 
        "implementation", "interactive", "math", "matrices", "meet-in-the-middle", 
        "number theory", "probabilities", "schedules", "shortest paths", 
        "sortings", "string suffix structures", "strings", "ternary search", 
        "trees", "two pointers"
    ];

    const tagsDropdown = document.getElementById('tags-dropdown');
    const selectedTagsDiv = document.getElementById('selected-tags');
    const selectTagsBtn = document.getElementById('select-tags-btn');
    let selectedTags = [];

    tags.forEach(tag => {
        const checkbox = document.createElement('input');
        checkbox.type = 'checkbox';
        checkbox.value = tag;
        checkbox.id = `tag-${tag}`;

        const label = document.createElement('label');
        label.htmlFor = `tag-${tag}`;
        label.textContent = tag;

        const div = document.createElement('div');
        div.classList.add('tag-option');
        div.appendChild(checkbox);
        div.appendChild(label);

        tagsDropdown.appendChild(div);
    });

    selectTagsBtn.addEventListener('click', () => {
        tagsDropdown.classList.toggle('hidden');
    });

    tagsDropdown.addEventListener('change', (event) => {
        const tag = event.target.value;

        if (event.target.checked) {
            if (!selectedTags.includes(tag)) {
                selectedTags.push(tag);
            }
        } else {
            selectedTags = selectedTags.filter(t => t !== tag);
        }

        updateSelectedTagsDisplay();
    });

    function updateSelectedTagsDisplay() {
        selectedTagsDiv.innerHTML = ''; 

        if (selectedTags.length > 0) {
            selectedTags.forEach(tag => {
                const tagDiv = document.createElement('div');
                tagDiv.classList.add('selected-tag');
                tagDiv.textContent = tag;
                selectedTagsDiv.appendChild(tagDiv);
            });
            selectTagsBtn.textContent = 'Edit Selected Tags';
        } else {
            selectTagsBtn.textContent = 'Select Tags';
        }
    }
});

document.getElementById('tags-difficulty-form').addEventListener('submit', async function (e) {
    e.preventDefault();
    
    const participants = document.getElementById('participants').value.trim().split(',');
    const minDifficulty = document.getElementById('min-difficulty').value;
    const maxDifficulty = document.getElementById('max-difficulty').value;
    
    const selectedTags = Array.from(document.querySelectorAll('#tags-dropdown input[type="checkbox"]:checked'))
                              .map(input => input.value);

    document.getElementById('loading-container').style.display = 'block';
    document.getElementById('output-container').style.display = 'none';
    document.getElementById('no-data-container').style.display = 'none';

    try {
        const response = await fetch('/api/problems', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                participants: participants,
                tags: selectedTags,
                minDifficulty: parseInt(minDifficulty),
                maxDifficulty: parseInt(maxDifficulty)
            })
        });

        const data = await response.json();
        
        if (data.validProblems.length > 0) {
            
            const tableBody = document.getElementById('valid-problems');
            tableBody.innerHTML = '';

            data.validProblems.forEach(problem => {
                const tagList = problem.tags.map(tag => `<span>${tag}</span>`).join('<br>'); 
                const row = document.createElement('tr');
                row.innerHTML = `
                    <td>${problem.name}</td>
                    <td>${problem.rating}</td>
                    <td>${tagList}</td>
                    <td>${problem.contestId}${problem.index}</td>
                    <td><a href="https://codeforces.com/contest/${problem.contestId}/problem/${problem.index}" target="_blank">Open Problem</a></td>
                `;
                tableBody.appendChild(row);
            });

            document.getElementById('output-container').style.display = 'block';
        } else {
            document.getElementById('no-data-container').style.display = 'block';
        }
    } catch (error) {
        console.error('Error fetching problems:', error);
    } finally {
        document.getElementById('loading-container').style.display = 'none';
    }
});


document.getElementById('tags-difficulty-form').addEventListener('submit', async function (e) {
    e.preventDefault();
    
    const participants = document.getElementById('participants').value.trim().split(',');
    const minDifficulty = document.getElementById('min-difficulty').value;
    const maxDifficulty = document.getElementById('max-difficulty').value;
    const selectedTags = Array.from(document.querySelectorAll('.tag.selected')).map(tag => tag.textContent);

    document.getElementById('loading-container').style.display = 'block';
    document.getElementById('output-container').style.display = 'none';
    document.getElementById('no-data-container').style.display = 'none';

    try {
        const response = await fetch('/api/problems', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                participants: participants,
                tags: selectedTags,
                minDifficulty: parseInt(minDifficulty),
                maxDifficulty: parseInt(maxDifficulty)
            })
        });

        const data = await response.json();
        
        if (data.validProblems.length > 0) {
            const tableBody = document.getElementById('valid-problems');
            tableBody.innerHTML = '';

            data.validProblems.forEach(problem => {
                const tagList = problem.tags.map(tag => `<span>${tag}</span>`).join('<br>'); 
                const row = document.createElement('tr');
                row.innerHTML = `
                    <td>${problem.name}</td>
                    <td>${problem.rating}</td>
                    <td>${tagList}</td>
                    <td>${problem.contestId}${problem.index}</td>
                    <td><a href="https://codeforces.com/contest/${problem.contestId}/problem/${problem.index}" target="_blank">Open Problem</a></td>
                `;
                tableBody.appendChild(row);
            });

            document.getElementById('output-container').style.display = 'block';
        } else {
            document.getElementById('no-data-container').style.display = 'block';
        }
    } catch (error) {
        console.error('Error fetching problems:', error);
    } finally {
        document.getElementById('loading-container').style.display = 'none';
    }
});

window.onload = async function() {
    const tagsContainer = document.getElementById('tags-container');
    const response = await fetch('/api/tags');
    const tags = await response.json();
    
    tags.forEach(tag => {
        const tagElement = document.createElement('span');
        tagElement.textContent = tag;
        tagElement.classList.add('tag');
        tagsContainer.appendChild(tagElement);

        tagElement.addEventListener('click', function () {
            tagElement.classList.toggle('selected');
        });
    });
};
