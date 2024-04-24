import { getData } from "../../utils/helperFetch.js";
import { socket } from "../index.js";
import { SendReponseToBack } from "../webscoket.js";
import { routage } from "./route.js";
// import { userConnected } from "./login.js";
// import { navigateTo, navigateToWithoutSavingHistory } from "./route.js";

let methode = "GET";
export default async function () {
  let conn = socket();
  HomePage();
  getPosts();
  SendReponseToBack(methode, "", conn);
  displayCreerMessage();
}
export async function getPosts() {
  let postsContainer = document.querySelector(".ibg");
  try {
    const value = await getData("/posts-data");
    let allPosts = value.post;
    if (allPosts == null) {
      const val = await getData("/login-data");
      console.log(val);
      let userC = document.querySelector("#user_session");
      userC.innerText = val.name_user;
      let emptyContainer = document.createElement("div");
      emptyContainer.classList.add("empty-container");
      emptyContainer.innerText = "Aucun post";
      postsContainer.appendChild(emptyContainer);
      return;   
    }
    console.log(value);
    // GET User from the session Data
    let userC = document.querySelector("#user_session");
    userC.innerText = value.user_data.username;
    // Supprimer tous les éléments enfants avant d'ajouter les nouveaux posts
    allPosts.forEach((post) => {
      let postElement = document.createElement("div");
      postElement.classList.add("Post");
      // Header du Post
      let headerPost = document.createElement("div");
      headerPost.classList.add("header-post");
      let imgDiv = document.createElement("div");
      imgDiv.innerHTML =
        '<img src="assets/profil.png" class="header-post-img" />';
      let infoDiv = document.createElement("div");
      infoDiv.classList.add("header-post-info");
      let nameUser = document.createElement("h4");
      nameUser.innerText = post.name_user_post;
      let dateElement = document.createElement("span");
      dateElement.innerText = post.date;
      infoDiv.appendChild(nameUser);
      infoDiv.appendChild(dateElement);
      headerPost.appendChild(imgDiv);
      headerPost.appendChild(infoDiv);
      // Contenu du Post
      let containPostInfo = document.createElement("div");
      containPostInfo.classList.add("contain-post-info");
      let titlepost = document.createElement("h4");
      titlepost.innerText = post.title;
      let contentElement = document.createElement("p");
      contentElement.innerText = post.content;
      containPostInfo.appendChild(titlepost);
      containPostInfo.appendChild(contentElement);
      // Options du Post (like, dislike, comment)

      let postOptions = document.createElement("div");
      postOptions.classList.add("post__options");
      // Bouton Like
      let likeOption = document.createElement("div");
      likeOption.classList.add("post__option");
      let likeForm = document.createElement("form");
      likeForm.classList.add("like-form");
      likeForm.action = "/like";
      likeForm.method = "post";
      let likeInputId = document.createElement("input");
      likeInputId.type = "hidden";
      likeInputId.name = "idpost";
      likeInputId.value = post.id; // Assurez-vous d'avoir l'id du post ici
      let likeInputSubmit = document.createElement("input");
      likeInputSubmit.type = "submit";
      likeInputSubmit.name = "like";
      likeInputSubmit.classList.add("material-icons");
      likeInputSubmit.value = "thumb_up";
      let likeCount = document.createElement("p");
      likeCount.innerText = "2"; // Remplacez par la valeur réelle du nombre de likes
      likeForm.appendChild(likeInputId);
      likeForm.appendChild(likeInputSubmit);
      likeOption.appendChild(likeForm);
      likeOption.appendChild(likeCount);
      // Bouton Dislike
      let dislikeOption = document.createElement("div");
      dislikeOption.classList.add("post__option");
      let dislikeForm = document.createElement("form");
      dislikeForm.classList.add("dislike-form");
      dislikeForm.action = "/like";
      dislikeForm.method = "post";
      let dislikeInputId = document.createElement("input");
      dislikeInputId.type = "hidden";
      dislikeInputId.name = "idpost";
      dislikeInputId.value = post.id; // Assurez-vous d'avoir l'id du post ici
      let dislikeInputSubmit = document.createElement("input");
      dislikeInputSubmit.type = "submit";
      dislikeInputSubmit.name = "like";
      dislikeInputSubmit.classList.add("material-icons");
      dislikeInputSubmit.value = "thumb_down";
      let dislikeCount = document.createElement("p");
      dislikeCount.innerText = "10";
      dislikeForm.appendChild(dislikeInputId);
      dislikeForm.appendChild(dislikeInputSubmit);
      dislikeOption.appendChild(dislikeForm);
      dislikeOption.appendChild(dislikeCount);
      // Bouton Comment
      let commentOption = document.createElement("div");
      commentOption.classList.add("post__option");
      let commentLabel = document.createElement("label");
      commentLabel.for = post.id; // Assurez-vous d'avoir l'id du post ici
      let commentInput = document.createElement("input");
      commentInput.type = "submit";
      commentInput.id = post.id; // Assurez-vous d'avoir l'id du post ici
      commentInput.classList.add("material-icons");
      commentInput.value = "chat_bubble_outline";
      let commentCount = document.createElement("p");
      commentCount.innerText = "20"; // Remplacez par la valeur réelle du nombre de commentaires
      commentLabel.appendChild(commentInput);
      commentLabel.appendChild(commentCount);
      commentOption.appendChild(commentLabel);
      postOptions.appendChild(likeOption);
      postOptions.appendChild(dislikeOption);
      postOptions.appendChild(commentOption);
      postElement.appendChild(headerPost);
      postElement.appendChild(containPostInfo);
      postElement.appendChild(postOptions);
      postsContainer.appendChild(postElement);
    });
  } catch (err) {
    routage("/login");
    console.log(err);
    console.error(err);
  }
}
export function HomePage() {
  document.querySelector(".app").innerHTML = `
  <div class="header-content">
  <header class="header">
  <div class="logo">
  <img class="logo" src="assets/accueil.png" />
  </div>
            <div class="header__input">
                <h2 class="titre">REAL TIME FORUM</h2>
            </div>
            <div class="tools">
              <span type="" id="user_session"></span>
              <button class="btn__logOut">LogOut</span></button>
            </div>
        </header>
        </div>
        <div class="containerHomepage">
        <div class="sidebar-left">
        <a href="/">
        <div class="sidebarRow sidebar__home">
        <span class="material-icons"> home </span>
        <h4>Home</h4>
        </div>
        </a>
        <a href="/profil">
        <div class="sidebarRow sidebar__profile">
        <span class="material-icons"> person </span>
        <h4>Profil</h4>
        </div>
        </a>
        <a href="/posts">
        <div class="sidebarRow sidebar__post">
        <span class="material-icons"> chat </span>
        <h4>POST</h4>
        </div>
        </a>
        </div>

        <div class="middle">
        <!-- elements du milieu -->
            <div class="ibg">
              
            </div>
       </div>

      <div class="sidebar-right">
        <div class="online-content">
          <span id="started"> User Online</span>
          <span> User Online</span>
          <span> User Online</span>
          <span> User Online</span>
        </div>
      </div>
</div>
`;
}

function displayCreerMessage() {
  let startmessage = document.querySelector("#started");
  startmessage.addEventListener("click", function () {
    chatContainer();
  });
}

export const chatContainer = () => {
  document.querySelector(".ibg").innerHTML = `
  <div id="chat-container">
  <div id="info-receiver"> Username receiver </div>
  <div id="message-zone">
  <div id="msg-sender"> Hello sender </div>
  <div id="msg-receiver">Message receiver</div>
  </div>
  <div id="make-msg">
  <input type="text" id = "message" placeholder="Entrez votre message...">
  <button id = "submitMessage">Envoyer</button>
  </div>
  </div>
  `;
};
