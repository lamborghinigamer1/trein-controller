package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

var (
	mu          sync.Mutex
	lastRequest time.Time
	minInterval = 2 * time.Second // Minimum interval between requests
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func handleFileUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("fileupload")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern

	if !strings.HasSuffix(handler.Filename, ".mp3") {
		return
	}
	
	tempFile, err := os.Create(handler.Filename)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Check if the minimum interval has passed since the last request
	if time.Since(lastRequest) < minInterval {
		http.Error(w, "Too many requests, please wait", http.StatusTooManyRequests)
		return
	}

	// Update the last request time
	lastRequest = time.Now()

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Access form data
	data := r.Form.Get("key") // Change "key" to the actual form field name

	// Process the data
	fmt.Println("Received data:", data)

	// Set CORS headers
	enableCors(&w)

	// Respond to the request
	fmt.Fprint(w, "Data received successfully!")

}

func handlePlaysound(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Set CORS headers
	enableCors(&w)

	// Update the last request time
	lastRequest = time.Now()

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Access form data
	data := r.Form.Get("key") // Change "key" to the actual form field name

	// Process the data
	fmt.Println("Received data:", data)

	go playSound(data)
}

type SoundInfo struct {
	Name string `json:"name"`
}

func getSounds(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	// Set CORS headers
	enableCors(&w)
	folderPath := "."

	fileInfo, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatal(err)
	}

	mp3Files := []SoundInfo{}

	for _, file := range fileInfo {
		if file.IsDir() {
			// Skip directories
			continue
		}

		if strings.HasSuffix(file.Name(), ".mp3") {
			soundInfo := SoundInfo{Name: file.Name()}
			mp3Files = append(mp3Files, soundInfo)
		}
	}

	// Convert the slice of SoundInfo to JSON
	jsonData, err := json.Marshal(mp3Files)
	if err != nil {
		http.Error(w, "Failed to convert to JSON", http.StatusInternalServerError)
		return
	}

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON data to the response writer
	w.Write(jsonData)
}

func playSound(mp3 string) {
	args := []string{mp3}

	audioplayer := "mpg123"
	cmd := exec.Command(audioplayer, args[0:]...)

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("playing %s\n", args[0])
}

func main() {
	http.HandleFunc("/postdata", handlePostRequest)
	http.HandleFunc("/playsound", handlePlaysound)
	http.HandleFunc("/upload", handleFileUpload)
	http.HandleFunc("/allsounds", getSounds)
	fmt.Println("Server listening on 8080...")
	http.ListenAndServe(":8080", nil)
}
