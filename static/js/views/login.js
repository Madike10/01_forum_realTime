import { postData } from "../../utils/helperFetch.js";

export default async function() {
  LoginPage();
  loginpage();
}

//LOGIN TRAITEMENT DE DONNEES
export function loginpage() {
  const loginForm = document.querySelector("#login-form");
  const errLoginMsg = document.querySelector(".errLoginMsg");
  errLoginMsg.innerText = "";
  loginForm.addEventListener("submit", (event) => {
    event.preventDefault();
    const usernameOremail = document.querySelector("#user_email").value;
    const password = document.querySelector("#password").value;
    console.log(usernameOremail, password);
    let credentials = {
      username_email: usernameOremail,
      password: password,
    };
    postData("/login-data", credentials)
      .then((value) => {
        if (value.error) {
          return value.error;
        }
        console.log(value.name_user);
        window.location.href = "/";
      })
      .catch((error) => {
        errLoginMsg.innerText = "wrong credentials";
        console.error("error", error);
        return;
      });
      errLoginMsg.innerText = "";
  });
}

export const LoginPage = () => {
  document.querySelector(".app").innerHTML = `
        <div class="containerLogin">
        <div class="login-right">
            <form method="POST" action="/home" class="form" id="login-form">
                <h1>Sign Up</h1>
                <span class="errLoginMsg" style="color : red; "></span>
                <input class="login-input" type="text" placeholder="Username" id="user_email">
                <input class="login-input" type="password" placeholder="Password" id="password">
                <input class="login-input" type="submit" class="submitBtn btnLogin" placeholder="Username" value="Login">
                </form>
                <a href="/register">
                  <button class="btn__register">Go to Register</button>
                </a>
        </div>
        <div class="login-left">
        </div>
        </div>
    `;
}

