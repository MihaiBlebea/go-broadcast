themeToggle = document.getElementById("theme-toggle")

if (localStorage.getItem("theme") !== null) {
    document.documentElement.setAttribute("data-theme", localStorage.getItem("theme"))
    if (localStorage.getItem("theme") === "light") {
        themeToggle.innerHTML = "Light"
    } else {
        themeToggle.innerHTML = "Dark"
    }
}

themeToggle.addEventListener("click", function() {
    trans()
    if (document.documentElement.getAttribute("data-theme") === "light") {
        document.documentElement.setAttribute("data-theme", "dark")
        localStorage.setItem("theme", "dark");
        this.innerHTML = "Dark"
    } else {
        document.documentElement.setAttribute("data-theme", "light")
        localStorage.setItem("theme", "light");
        this.innerHTML = "Light"
    }
})

let trans = ()=> {
    document.documentElement.classList.add("transition")
    window.setTimeout(()=> {
        document.documentElement.classList.remove("transition")
    }, 1000)
}