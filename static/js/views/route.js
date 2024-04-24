import  home  from "./home.js";
import register  from "./register.js";
import  login  from "./login.js";
import createPost  from "./createPost.js";
import error from "./error.js";

export const routage = (url) => {
    history.pushState(null, null, url)
    router()
}

export const navigateToWithoutSavingHistory = (url) => {
    history.replaceState(null, null, url)
    router()
}

export const router = async () => {
    const routes = [
        { path: "/", view: home },
        { path: "/home", view: home},
        { path: "/register", view: register },
        { path: "/login", view: login },
        { path: "/posts", view: createPost, home  },
        // { path: "/logout-data", view: logout },
    ]
    const potentialMatches = routes.map(route => {
        return {
            route: route,
            isMatch: location.pathname === route.path
        }
    })

    let match = potentialMatches.find(potentialMatch => potentialMatch.isMatch)
    if(match) {
        match.route.view()
    } else {
        error();
        console.log("Could not find");
    }
}

