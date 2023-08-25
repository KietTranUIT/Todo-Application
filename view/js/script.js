function GetSignUpPage() {
    fetch('/signup', {
        method: 'GET'
    })
    .then(response => {
        if(response.ok) {
            window.location.href = response.url;
        } else {
            console.error('Yêu cầu không thành công. Mã trạng thái HTTP:', response.status);
        }
    })
}

function SignUp() {

  let form = {
    email: document.getElementById('email').value,
    code: document.getElementById('code').value,
    username: document.getElementById('username').value,
    password: document.getElementById('password').value
  }

  var formJson = JSON.stringify(form);
  console.log(formJson);

  fetch('/signup', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: formJson
  })
  .then(response => response.json)
  .then(data => {
    alert(data);
  })
  .then(data => {
    alert(data)
  })
}

function VerifyEmail(event) {
  event.preventDefault();
    let Email = {
        email: document.getElementById('email').value,
        username: document.getElementById('username').value,
        password: document.getElementById('password').value
    }

    var jsonData = JSON.stringify(Email);

    fetch('/verify', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: jsonData
      })
        .then(response => {
          console.log(response.text)
        })
        
        .catch(error => {
          // Xử lý lỗi
        });
}