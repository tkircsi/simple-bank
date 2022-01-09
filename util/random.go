package util

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var userStream <-chan []string

func init() {
	rand.Seed(time.Now().UnixNano())

	done := make(chan interface{})

	userStream = randomUserGenerator(done)

}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	curr := []string{"EUR", "USD", "HUF"}
	return curr[rand.Intn(len(curr))]
}

func GenerateRandomString(n int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num := rand.Int31n(int32(len(letters)))
		ret[i] = letters[num]
	}
	return string(ret)
}

func RandomUser() []string {
	rec := <-userStream
	fullName := rec[0]
	password := GenerateRandomString(10)
	userName := strings.ToLower(strings.ReplaceAll(fullName, " ", ""))
	email := fmt.Sprintf("%s@simplemail.com", userName)
	return []string{userName, fullName, email, password}

}

func randomUserGenerator(done <-chan interface{}) <-chan []string {
	userStream := make(chan []string)
	r := csv.NewReader(strings.NewReader(names))
	go func() {
		defer close(userStream)
		for {
			rec, err := r.Read()
			if err != nil {
				return
			}
			select {
			case <-done:
				return
			case userStream <- rec:
			}
		}
	}()
	return userStream
}

var names = `Ellen Bruce
Melvin Hoffman
Nicholas Smith
Zachary Mullins
Anthony Hall
Willie Lambert
Max Bonilla
Cindy Schneider
Michael Drake
Jacob Martin
Michael James
Lawrence Moore
Julie Steele
Jeffrey Edwards
Isaac Yang
Nicole Mendoza
Dana Williams
Jorge Coleman
Melanie Klein
James Bowers
Christopher Jordan
Sharon Scott
Benjamin Lewis
Kathleen Simpson
Cynthia Smith
Benjamin Hall
Matthew Smith
James Odonnell
Whitney Robinson
Lori Pitts
Bailey Adams
Dylan Morgan
Amy Andrade
Lisa Harris
Michael Bishop
Jeremiah David
Kristen Weaver
Kathy Shelton
Jack Powell
Kathleen Castaneda
Robert Peterson
Miguel Johnston
Gabriella Richardson
Marco Molina
Stephen Riggs
Samuel Ruiz
Peter Boone
David Lee
Rachel Schneider
James Thompson
Joe Copeland
April Campbell
Andrea Hernandez
Michael Cameron
Karen Terry
Ashley Howe
Sean Moore
Pamela Merritt
Mia Stewart
Peter Fisher
Chelsea Simpson
Tyler Norris
Wendy Brewer
Cheryl Kline
Ann Casey
Julie Rios
Hector Duran
Jake Schaefer
Jennifer Conley
Javier Farmer
Cindy Eaton
Alicia White
Brandon Moore
Christina Williams
Brian Lopez
Christine Copeland
Sara Shelton
Sandra Boyer
Ruben Lester
Lisa Perkins
Edwin Sheppard
Christopher Mccullough
Destiny Warren
Jackie Miles
Lauren Kline
Valerie Jenkins
Christine Gordon
Heidi Dillon MD
Anthony Cunningham
Dustin Harrison
Justin Johnson
Jordan Ford
Sarah Mckinney
George Goodwin
Monica Nichols
Christina Morris`
