package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"path/filepath"
	"time"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

// generates qr code png files
// also generates qr relations data INSERT sql

const hashBytes = 16
const startID = 1
const numQR = 1600
const dst = "./qr"
const contentTemplate = "https://seikatsusoken.jp/Ny11r/q/"

func main() {
	fmt.Println("-- genqr main")

	start := time.Now()

	pairs := generateIDHashPairs(startID, numQR)

	for _, p := range pairs {
		fu := generateFileURLPair(p.Hash)
		generateQRPNG(fu.URL, fu.File)
	}

	dumpInsertStatement(pairs)

	fmt.Println("-- ", time.Since(start), "elapsed")
}

func generateQRPNG(content, filename string) {
	qrpath := filepath.Join(dst, filename)

	err := qrcode.WriteFile(content, qrcode.Medium, 320, qrpath)
	if err != nil {
		panic(err)
	}
}

func generateHash() string {
	b := make([]byte, hashBytes)
	if n, err := rand.Read(b); err != nil || n != hashBytes {
		fmt.Println("rand.Read fail. n:", n)
		panic(err)
	}
	return hex.EncodeToString(b)
}

type IDHash struct {
	ID   int
	Hash string
}

func generateIDHashPairs(start, num int) []IDHash {

	ret := make([]IDHash, 0, num)
	set := map[string]struct{}{}

	i := start

	for {
		h := generateHash()
		p := IDHash{
			ID:   i,
			Hash: h,
		}
		i++
		set[h] = struct{}{}
		ret = append(ret, p)
		if len(ret) == num {
			break
		}
	}

	if len(ret) != len(set) {
		panic("collision!")
	}

	return ret
}

type FileURLPair struct {
	File string
	URL  string
}

func generateFileURLPair(hash string) FileURLPair {
	return FileURLPair{
		File: hash + ".png",
		URL:  contentTemplate + hash,
	}
}

func dumpInsertStatement(pairs []IDHash) {
	const stmtSize = 10

	for lb := 0; lb < len(pairs); lb += stmtSize {
		ub := lb + stmtSize
		if ub > len(pairs) {
			ub = len(pairs)
		}

		chunk := pairs[lb:ub]
		query := "INSERT INTO `ny11r_qr_relations` (`entry_id`, `code`) VALUES"
		stmt := fmt.Sprintln(query)
		fmt.Println(query)
		for i, p := range chunk {
			stmt += fmt.Sprintf("(%d, \"%s\")", p.ID, p.Hash)
			fmt.Printf("(%d, \"%s\")", p.ID, p.Hash)
			if i+1 == len(chunk) {
				stmt += fmt.Sprintln(";")
				fmt.Println(";")
			} else {
				stmt += fmt.Sprintln(",")
				fmt.Println(",")
			}
		}
		writeInsertStatement(stmt)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeInsertStatement(statement string) {
	txtpath := filepath.Join(dst, "qr.txt")
	f, err := os.OpenFile(txtpath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	check(err)
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%s%s", statement, "\n"))
	check(err)
}