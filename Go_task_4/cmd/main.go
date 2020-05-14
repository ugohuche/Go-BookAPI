func main()  {
  mux := http.NewServerMux()
  mux.HandleFunc("api/v1/books", logger())

  http.ListenandServe(":8000", mux)
}
