import { routage, router } from "./views/route.js";

window.addEventListener("popstate", router);
// export let socket;
document.addEventListener("DOMContentLoaded", () => {
  let conn = socket();
  console.log(conn);
  document.body.addEventListener("click", (event) => {
    if (event.target.matches("[data-link]")) {
      event.preventDefault();
      routage(event.target.href);
    }
  });
  router();
});

export const socket = () => {
  console.log("Statar socket in js");
  // connect to websocket
  let conn = new WebSocket("ws://localhost:8080/ws");
 return conn
};
