package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"01.kood.tech/git/obudarah/ascii-art-web-all/asciiart"
)

const (
	TEMPLATES_PATH = "./templates/"
	BANNER_PATH    = "./banners/"
)

type outputData struct {
	Input    string
	Output   string
	Color    string
	BkColor  string
	Err      string
	isResult bool
}

func main() {
	out := outputData{
		isResult: false,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/textfile/AsciiResult.txt", out.downloadFile)
	mux.HandleFunc("/", out.home)
	fileServer := http.FileServer(http.Dir(TEMPLATES_PATH + "static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	port, err := parseArgs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Starting server at port %s\n", *port)
	if err := http.ListenAndServe(":"+*port, mux); err != nil {
		log.Fatal(err)
	}
}

/*
a handler for the main page
*/
func (out *outputData) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		NotFound(w, r)
		return
	}

	*out = outputData{isResult: false}

	method := r.Method
	if method == "POST" {
		out.postHandler(w, r)
	}

	// assemble the page from templates
	site := []string{
		TEMPLATES_PATH + "base.layout.html",
		TEMPLATES_PATH + "home.page.tmpl",
	}

	tm, err := template.ParseFiles(site...)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tm.Execute(w, out)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

/*
handles data passed by POST method
*/
func (out *outputData) postHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("ParseForm() error: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	out.Input = r.FormValue("text_string")
	if ok, _ := asciiart.IsAsciiString(out.Input); !ok {
		log.Printf("not ascii symbols in the input: %s \n", out.Input)
		out.Err = "Error: only ASCII characters are accepted. Please try again."
		return
	}

	fontName := BANNER_PATH + r.FormValue("text_style") + ".txt"
	aText, err := asciiart.TextToArt(out.Input, fontName)
	if err != nil {
		log.Printf("error occures during making art ascii string\n text: %s banners: %s error: %s", out.Input, fontName, err)
		http.NotFound(w, r)
		return
	}
	// data for output
	out.Output = aText // ascii presentation of the string
	out.Color = r.FormValue("color")
	out.BkColor = r.FormValue("bk-color")
	out.isResult=true
	// os.WriteFile("./textfile/AsciiResult.txt", []byte(out.Output), 0o666)
}

/*
Downloading file
*/
func (out *outputData) downloadFile(w http.ResponseWriter, r *http.Request) {
	if !out.isResult{
		NotFound(w,r)
		return
	}
	
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.Header().Add("Content-Disposition", "attachment; filename=AsciiResult.txt")
	w.Header().Add("Content-Length", strconv.Itoa(len([]byte(out.Output))))
	// http.ServeFile(w, r, "./textfile/AsciiResult.txt")
	w.Write([]byte(out.Output))
	return
}

/*
replies to the request with HTTP 404 error using a pretty 404 page
*/
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound) // set status code at 404
	// send a pretty 404 page
	tm, _ := template.ParseFiles(TEMPLATES_PATH + "error404.html")
	err := tm.Execute(w, nil)
	if err != nil {
		http.NotFound(w, r)
		log.Println(err)
	}
	return
}

/*
parse program's arguments in the aim to obtain the server port. If there are no argumens, will return port 8080
*/
func parseArgs() (*string, error) {
	port := flag.String("port", "8080", "server port")
	flag.Parse()
	if flag.NArg() > 0 {
		return nil, fmt.Errorf("Wrong arguments\nUsage: go run .  --port=PORT_NUMBER\n")
	}
	_, err := strconv.ParseUint(*port, 10, 16)
	if err != nil {
		return nil, fmt.Errorf("error: port must be a 16-bit unsigned number ")
	}
	return port, nil
}
