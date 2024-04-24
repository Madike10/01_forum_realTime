import { postData } from "../../utils/helperFetch.js";
import { routage } from "./route.js"

export default async function() {
    RegisterPage();
    registerpage();
}

// REGISTER TRAITEMENT DE DONNEES
export function registerpage() {
  const registerForm = document.querySelector("#register-form");
  registerForm.addEventListener("submit", (event) => {
    event.preventDefault();
    const nickname = document.querySelector("#nickname").value;
    const age = document.querySelector("#age").value;
    const gender = document.querySelector("#gender").value;
    const firstname = document.querySelector("#firstname").value;
    const lastname = document.querySelector("#lastname").value;
    const email = document.querySelector("#email").value;
    const password = document.querySelector("#password").value;
    const errorMsg = document.querySelector(".errorMsg");
    errorMsg.innerText = "";
    let isValid = true;
    if (firstname === "") {
      errorMsg.innerText += "Enter a firstname\n";
      isValid = false;
    }
    if (lastname === "") {
      errorMsg.innerText += "Enter a lastname\n";
      isValid = false;
    }
    if (email === "") {
      errorMsg.innerText += "Enter a email\n";
      isValid = false;
    }
    if (password === "") {
      errorMsg.innerText += "Enter a password\n";
      isValid = false;
    }
    if (nickname === "") {
      errorMsg.innerText += "Enter a nickname\n";
      isValid = false;
    }
    if (age === "") {
      errorMsg.innerText += "Enter a age\n";
      isValid = false;
    }
    if (!isValid) {
      return;
    }
    // Check validity password
    if (password.length < 4) {
      errorMsg.innerText += "Password must be at least 4 characters\n";
      return;
    }
    //check validity email address
    if (!isValidEmail(email)) {
      errorMsg.innerText += "Email address must be valid\n";
      return;
    }
    let data = {
      name: nickname,
      email: email,
      password: password,
      age: age,
      first_name: firstname,
      last_name: lastname,
      gender: gender,
    };
    postData("/register-data", data)
      .then((value) => {
        if (value.error) {
          return;
        }
        routage("/login");
      })
      .catch((error) => {
        console.error("Erreur lors de la requête POST :", error);
      });
  });
}

// Function to check if email is valid
function isValidEmail(email) {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return emailRegex.test(email);
}

export const RegisterPage = () =>{
  document.querySelector(".app").innerHTML = `
      
    <div class="containerRegister">
          <div class="register-left">
          </div>

          <div class="register-right">
              <form class="form" id="register-form">
                  <h1>Sign In</h1>
                  <span class="errorMsg" style="color:red;"></span>
                  <input class="register-input" id="nickname" type="text" placeholder="Nickname">
                  <input class="register-input" id="age" type="age" placeholder="age">
                  <select name="" class="" id="gender">
                      <option id="gender" value="Male">Male</option>
                      <option id="gender" value="Female">Female</option>
                  </select>
                  <input class="register-input" id="firstname" type="text" placeholder="First Name">
                  <input class="register-input" id="lastname" type="text" placeholder="Last Name">
                  <input class="register-input" id="email" type="email" placeholder="Email">
                  <input class="register-input" id="password" type="password" placeholder="Password">
                  <input class="register-input" type="submit" class="submitBtn" value="Register">
                  </form>
                  <button class="btn__login">Go to Login</button>
          </div>
          </div>
  `;
}
