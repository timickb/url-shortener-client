import "./style.css"

function validateURL(url) {
    if (!url.startsWith("http://") && !url.startsWith("https://") && !url.startsWith("ftp://")) {
        return false
    }

    var spaceParts = url.split(' ')
    var dotParts = url.split('.')

    if (spaceParts.length > 1) {
        return false
    }

    if (dotParts.length < 2) {
        return false
    }

    return true
}

const generateBtn = document.getElementById("generate_btn")
const urlInput = document.getElementById("url_input")
const resultText = document.getElementById("result_text")
const errorText = document.getElementById("error_text")

const domain = "https://shorturl.me/"
const server = "http://localhost/create"

generateBtn.addEventListener("click", () => {
    resultText.innerText = ""
    errorText.innerText = ""

    urlInput.value = urlInput.value.trim()

    if (!validateURL(urlInput.value)) {
        errorText.innerText = "Incorrect URL"
        return
    }

    const queryBody = {
        url: urlInput.value
    }

    fetch(server, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(queryBody)
    })
    .then(res => res.json())
    .then(data => resultText.innerText = domain + data.hash)
    .catch(err => {
        console.log(err)
        errorText.innerText = "Connection error"
    })
})