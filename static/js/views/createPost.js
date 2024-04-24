import { getData, postData } from "../../utils/helperFetch.js";
import { routage } from "./route.js";

export default async function() {
  CreatePost();
  postpage();
}

// GET POST TRAITEMENT
export async function postpage() {
  const postForm = document.querySelector(".post__content");
  postForm.addEventListener("submit", async (event) => {
    event.preventDefault();
    const title = document.querySelector("#title_post").value;
    const content = document.querySelector("#content_post").value;
    const checkboxes = document.querySelectorAll(".category");
    const selectedValues = Array.from(checkboxes).map((checkbox) =>
    checkbox.checked ? checkbox.value : null
    );
    
    const date = new Date();
    const options = {
      weekday: "long",
      year: "numeric",
      month: "long",
      day: "numeric",
    };
    const dateFormatee = date.toLocaleDateString("fr-FR", options);
    const value = await getData("/posts-data");
    console.log("les valuer du user", value);
    let datapost = {
      title: title,
      content: content,
      categories: selectedValues,
      date: dateFormatee,
      id_user: value.user_data.user_id,
    };
    try {
      const value = await postData("/posts-data", datapost);
      console.log("Réponse du serveur :", value);
      if (value.error) {
        // Traitement en cas d'erreur
        return;
      }
      window.location.href = "/";
    } catch (error) {
      console.error("Erreur lors de l'envoi :", error);
    }
  });
}


export const  CreatePost = () =>{
    document.querySelector(".app").innerHTML = `
    <div class="feed">
    <!-- message sender starts -->
    <div class="messageSender">
        <div class="messageSender__top">
            <!-- <img class="user__avatar" src="../static/images/profile.png" alt="" /> -->
            <form  method="POST" class="post__content" enctype="multipart/form-data">
                <div class="title">
                    <label for="messageSender__input">Title</label>
                    <input class="messageSender__input" placeholder="Title" type="text" id="title_post"
                        required />
                </div>
                <div class="descrip">
                    <label for="messageSender__input "></label>
                    <textarea placeholder="What's on your mind?" type="text" id="content_post"
                        required rows="4" style="overflow: hidden; word-wrap: break-word; resize: none; height: 160px; "></textarea>
                </div>
                <div class="checkbox__wrapper">
                    <div>
                        <input type="checkbox" class="category" id="check" value="technologie" name="techno" />
                        <label for="check">Technologie</label>
                    </div>
                    <div>
                        <input type="checkbox" class="category" id="check" value="sport" name="sport" />
                        <label for="check">Sport</label>
                    </div>
                    <div>
                        <input type="checkbox" class="category" id="check" value="sante" name="sante" />
                        <label for="check">Sante</label>
                    </div>
                    <div>
                        <input type="checkbox" class="category" id="check" value="other" name="other" checked />
                        <label for="check">Other</label>
                    </div>
                </div>
                <div class="messageSender__bottom">
                <div class="messageSender__option">
                <!-- <span style="color: #FFB477" class="material-icons"> send  </span> -->
                <input style="color: #FFB477" class="material-icons" type="submit" value="send"
                name="send">
                <!-- <h3>Post</h3> -->
                </div>
                </div>
                <!-- <div class="error_form" style="color: red; text-align: center;">{{.Error}}</div> -->
                </form>
                </div>
                </div>
                </div>
                
                `;
            }