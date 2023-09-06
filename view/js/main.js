
function CreateCategory() {
    let input = document.getElementById('NameCategory')
    let Username = localStorage.getItem('username')

    const category = {
        name: input.value,
        owner: Username
    }
    let dataJson = JSON.stringify(category)

    fetch('/createcategory', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body:dataJson
    })
    .then(response => response.json())
    .then(data => {
        console.log(data)
    })
}