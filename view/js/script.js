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

function GetSignInPage() {
  fetch('/signin', {
    method: 'GET'
  })
  .then(response => {
    if(response.ok) {
      window.location.href = response.url;
    } else {
      console.error('Yêu cầu không thành công. Mã trạng thái HTTP:', response.status)
    }
  })
}

function SignIn() {
  let form = {
    username: document.getElementById('username').value,
    password: document.getElementById('password').value
  }

  var formJson = JSON.stringify(form);

  fetch('/signin', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: formJson
  })
  .then(response => response.json())
  .then(data => {
    if(!data['status']) {
      if(data['err_msg'] == "not exist user") {
        failNotification("Username không tồn tại")
      } else if(data['err_msg'] == "wrong password") {
        failNotification("Password không chính xác")
      } else {
        failNotification("Lỗi server")
      }
    } else {
      alert("Login success")
    }
  })
}

function SignUp() {
  let email = document.getElementById('email').value;
  let code = document.getElementById('code').value;
  let username = document.getElementById('username').value;
  let password = document.getElementById('password').value;

  if(email=="") {
    failNotification("Email không được để trống")
    return
  } else if(code=="") {
    failNotification("Code không được để trống")
    return
  } else if (username=="") {
    failNotification("Username không được để trống")
    return
  } else if(password=="") {
    failNotification("Password không được để trống")
    return
  }

  let form = {
    email: email,
    code: code,
    username: username,
    password: password
  }

  var formJson = JSON.stringify(form);

  fetch('/signup', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: formJson
  })
  .then(response => response.json())
  .then(data => {
    if(!data['status']) {
      if(data['err_msg'] == "duplicate user") {
        failNotification("Username đã tồn tại")
      } else if(data['err_msg'] == 'code expired') {
        failNotification("Mã xác thực đã hết hạn")
      } else if(data['err_msg'] == 'wrong code') {
        failNotification("Mã xác thực không chính xác")
      } 
      else {
        failNotification("Lỗi server")
      }
    } else {
      alert("Đăng kí thành công")
      GetSignInPage()
    }
  })
}

function ValidateEmailInput(email) {
  var emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/; // Biểu thức chính quy kiểm tra định dạng email

  if (!emailPattern.test(email)) {
    // Người dùng nhập sai định dạng email
    alert('Email không hợp lệ');
    return false;
  }
  return true;
} 

function VerifyEmail() {
    let Email = {
        email: document.getElementById('email').value,
        type: 'verify sign up'
    }

    if(!ValidateEmailInput(Email.email)) {
      document.getElementById('email').value = ""
      return
    }


    var jsonData = JSON.stringify(Email);

    fetch('/sendmail', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: jsonData
      })
        .then(response => response.json())
        .then(resJson => {
          if(resJson['status']) {
            successNotification("Đã gửi mã xác thực đến email")
          } else {
            if(resJson['err_msg'] == "duplicate code") {
              failNotification("Mã đã gửi đến email! Vui lòng kiểm tra lại email!")
            } else if(resJson['err_msg'] == "duplicate user") {
              failNotification("Email đã được sử dụng")
            } else {
              failNotification("Lỗi server!")
            }
          }
        })
        
        .catch(error => {
          // Xử lý lỗi
        });
}


// create a notify success message
function successNotification(message) {
  text1 = document.getElementById('text1')
  text2 = document.getElementById('text2')
  toast = document.querySelector(".toast");
  signal = document.querySelector(".check");

  text1.innerHTML = "Success"
  text2.innerHTML = message
  signal.style.backgroundColor = "#4BB543";
  toast.style.visibility = "visible";

  (closeIcon = document.querySelector(".close")),
  (progress = document.querySelector(".progress"));

  let timer1, timer2;

  toast.classList.add("active");
  progress.classList.add("active");

  timer1 = setTimeout(() => {
    toast.classList.remove("active");
  }, 5000); //1s = 1000 milliseconds

  timer2 = setTimeout(() => {
    progress.classList.remove("active");
  }, 5300);

closeIcon.addEventListener("click", () => {
  toast.classList.remove("active");

  setTimeout(() => {
    progress.classList.remove("active");
  }, 300);
  clearTimeout(timer1);
  clearTimeout(timer2);
  toast.style.visibility = "hidden";

});

}

// Create a Notify Fail message
function failNotification(message) {
  text1 = document.getElementById('text1')
  text2 = document.getElementById('text2')
  toast = document.querySelector(".toast");
  signal = document.querySelector(".check")

  text1.innerHTML = "Failed"
  text2.innerHTML = message
  signal.style.backgroundColor = "red"
  toast.style.visibility = "visible";


  (closeIcon = document.querySelector(".close")),
  (progress = document.querySelector(".progress"));

let timer1, timer2;

  toast.classList.add("active");
  progress.classList.add("active");

  timer1 = setTimeout(() => {
    toast.classList.remove("active");
  }, 5000); //1s = 1000 milliseconds

  timer2 = setTimeout(() => {
    progress.classList.remove("active");
  }, 5300);

closeIcon.addEventListener("click", () => {
  toast.classList.remove("active");

  setTimeout(() => {
    progress.classList.remove("active");
  }, 300);
  clearTimeout(timer1);
  clearTimeout(timer2);
  toast.style.visibility = "hidden";
});

}
