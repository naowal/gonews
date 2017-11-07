package app

import "net/http"

//Mount mounts handlers to mux
func Mount(mux *http.ServeMux) {
	mux.HandleFunc("/", index) //list all news
	mux.HandleFunc("/news/", newsView)

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/admin/login", adminLogin)
	adminMux.HandleFunc("/list", adminList)
	adminMux.HandleFunc("/create", adminCreate)
	adminMux.HandleFunc("/edit", adminEdit)

	mux.Handle("/admin/", http.StripPrefix("admin", onlyAdmin(adminMux)))

}

func onlyAdmin(h http.Handler) http.Handler {
	return h
}