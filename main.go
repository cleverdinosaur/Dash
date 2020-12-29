package main

import (
	"database/sql"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/bitly/go-simplejson"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gorilla/mux"
)

func getDimension(imagePath string) map[string]int {
	file, err := os.Open(imagePath)
	defer file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	image, _, err := image.DecodeConfig(file) // Image Struct
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}
	var newMap map[string]int
	newMap = make(map[string]int)
	newMap["Width"] = image.Width
	newMap["Height"] = image.Height
	return newMap
}

func encode(filePath string, keyword string, outputFileName string) {
	dimension := getDimension(filePath)
	inputImage, _ := os.Open(filePath)
	defer inputImage.Close()
	decodeImage, _, _ := image.Decode(inputImage)

	myImage := image.NewRGBA(image.Rect(0, 0, dimension["Width"]+50, dimension["Height"]+50+1))
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		// Handle error
	}

	bound := myImage.Bounds()
	bound.Min.X = 0
	bound.Min.Y = 0
	background := color.RGBA{255, 255, 255, 255}
	draw.Draw(myImage, bound, &image.Uniform{background}, image.Point{0, 1}, draw.Src)
	newbound := myImage.Bounds()
	newbound.Min.X = 25
	newbound.Min.Y = 26
	draw.Draw(myImage, newbound, decodeImage, image.Point{0, 1}, draw.Src)

	i := 0
	for chIndex := 0; chIndex < len(keyword); chIndex++ {
		myImage.Pix[0+i] = keyword[chIndex]
		myImage.Pix[1+i] = 0
		myImage.Pix[2+i] = 0
		myImage.Pix[3+i] = 255
		i = i + 4
	}

	y := 13
	for x := 0; x <= myImage.Bounds().Max.X+50; x++ {
		if RandBool() {
			myImage.SetRGBA(x, y, color.RGBA{0, 0, 0, 0})
		}
	}

	y1 := 25 + dimension["Height"] + 13
	for x := 0; x <= myImage.Bounds().Max.X+50; x++ {
		if RandBool() {
			myImage.SetRGBA(x, y1, color.RGBA{0, 0, 0, 0})
		}
	}

	x := 13
	for y := 0; y <= myImage.Bounds().Max.Y+50; y++ {
		if RandBool() {
			myImage.SetRGBA(x, y, color.RGBA{0, 0, 0, 0})
		}
	}

	x1 := 25 + dimension["Width"] + 13
	for y := 0; y <= myImage.Bounds().Max.Y+50; y++ {
		if RandBool() {
			myImage.SetRGBA(x1, y, color.RGBA{0, 0, 0, 0})
		}
	}
	png.Encode(outputFile, myImage)
	outputFile.Close()
}

func generateKeyword(length int) string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	charset := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func insertDB(keyword string, url string) {
	sqliteDatabase, _ := sql.Open("sqlite3", "./data.db")
	defer sqliteDatabase.Close()

	statement, _ := sqliteDatabase.Prepare("CREATE TABLE IF NOT EXISTS encoding (keyword TEXT UNIQUE, url TEXT)")
	statement.Exec()

	insertStudentSQL := `INSERT INTO encoding (keyword, url) VALUES (?, ?)`
	statement, err := sqliteDatabase.Prepare(insertStudentSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(keyword, url)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func encodeHandler(url string, filePath string) string {
	keyword := generateKeyword(6)
	// url := generateKeyword(20)
	log.Println("keyword:" + keyword)
	log.Println("url:" + url)
	// filePath := "C:\\goproject\\1.jpg"
	now := time.Now()
	sec := now.Unix()
	outputFileName := "./static/" + strconv.FormatInt(int64(sec), 10) + ".png"
	log.Println(outputFileName)
	encode(filePath, keyword, outputFileName)
	insertDB(keyword, url)
	return outputFileName
}

func decode(filePath string) string {
	existingImageFile, err := os.Open(filePath)
	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		// Handle error
	}
	keyword := ""
	for x := 0; x < 6; x++ {
		rr, _, _, _ := loadedImage.At(x, 0).RGBA()
		rr = rr / 256
		keyword = keyword + string([]rune{rune(rr)})
	}

	sqliteDatabase, _ := sql.Open("sqlite3", "./data.db")
	defer sqliteDatabase.Close()

	statement, _ := sqliteDatabase.Prepare("CREATE TABLE IF NOT EXISTS encoding (keyword TEXT UNIQUE, url TEXT)")
	statement.Exec()

	query := "Select url from encoding where keyword = '" + keyword + "'"
	rows, _ := sqliteDatabase.Query(query)
	var url string
	for rows.Next() {
		rows.Scan(&url)
	}
	return url
}

func encodeUrl(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	file, handler, err := r.FormFile("file")
	url := r.FormValue("url")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, _ = io.Copy(f, file)
	time.Sleep(7 * time.Second)
	outputFileName := encodeHandler(url, handler.Filename)
	log.Println(handler.Filename)

	json := simplejson.New()
	json.Set("file", outputFileName)

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func decodeImage(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	file, handler, err := r.FormFile("file")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, _ = io.Copy(f, file)
	url := decode(handler.Filename)
	log.Println(url)

	json := simplejson.New()
	json.Set("url", url)

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	json := simplejson.New()
	json.Set("status", "welcome")

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func RandBool() bool {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/encode", encodeUrl).Methods("POST")
	router.HandleFunc("/decode", decodeImage).Methods("POST")

	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	router.PathPrefix("/static/").Handler(s)

	log.Println("SERVER STARTED at port: 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
