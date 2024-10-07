document.getElementById('tags-difficulty-form').addEventListener('submit', async function (e) {
    e.preventDefault();
    
    const participants = document.getElementById('participants').value.trim().split(',');
    const tags = document.getElementById('tags').value.trim().split(',');
    const minDifficulty = document.getElementById('min-difficulty').value;
    const maxDifficulty = document.getElementById('max-difficulty').value;

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
                tags: tags,
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
    } catch (err) {
        console.error('Error fetching problems:', err);
    } finally {
        document.getElementById('loading-container').style.display = 'none';
    }
});
