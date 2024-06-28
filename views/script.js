let currentModel = 'Table QA';

function selectContact(model) {
    currentModel = model;

    document.getElementById('chat-header').innerHTML = `<h2>${model}</h2>`;
    document.getElementById('chat-box').innerHTML = '';

    const contactItems = document.querySelectorAll('.contact-item');
    contactItems.forEach(item => {
        if (item.getAttribute('data-model') === model) {
            item.classList.add('active');
        } else {
            item.classList.remove('active');
        }
    });


    const fileInput = document.getElementById('file-input');
    if (model === 'Table QA') {
        fileInput.style.display = 'block';
    } else {
        fileInput.style.display = 'none';
    }
}

function handleKeyPress(event) {
    if (event.key === 'Enter') {
        sendMessage();
    }
}

function sendMessage() {
    const input = document.getElementById('chat-input');
    const message = input.value.trim();
    if (!message) return;

    const chatBox = document.getElementById('chat-box');
    const userMessage = `<div class="chat-message user"><div class="message">${message}</div></div>`;
    chatBox.innerHTML += userMessage;

    input.value = '';

    if (currentModel === 'Table QA') {
        const fileInput = document.getElementById('file-input');
        const file = fileInput.files[0];
        if (!file) {
            alert('Please upload a CSV file.');
            return;
        }

        const formData = new FormData();
        formData.append('query', message);
        formData.append('file', file);

        fetch('/tableqa', {
            method: 'POST',
            body: formData,
        })
        .then(response => response.json())
        .then(data => {
            const botMessage = `<div class="chat-message bot"><div class="message">${data.answer}</div></div>`;
            chatBox.innerHTML += botMessage;
        });
    } else if (currentModel === 'Summarization') {
        fetch('/summarize', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: `input=${encodeURIComponent(message)}`,
        })
        .then(response => response.json())
        .then(data => {
            const botMessage = `<div class="chat-message bot"><div class="message">${data.summary}</div></div>`;
            chatBox.innerHTML += botMessage;
        });
    } else if (currentModel === 'Chat Me') {
        fetch('/chatme', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: `message=${encodeURIComponent(message)}`,
        })
        .then(response => response.json())
        .then(data => {
            const botMessage = `<div class="chat-message bot"><div class="message">${data.response}</div></div>`;
            chatBox.innerHTML += botMessage;
        });
    }
}
