package seed

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
)

type SeedConfig struct {
	API                string
	NumWorkers         int
	NumUsers           int
	NumRoles           int
	NumPermissions     int
	NumRolePermissions int
	NumUserRoles       int
}

var adjectives = []string{
	"quick", "silent", "bright", "brave", "fuzzy", "wild", "wise", "sneaky", "gentle", "bold",
	"zesty", "sleepy", "fiery", "curious", "grumpy", "cheeky", "stormy", "loyal", "mystic", "nimble",
	"rusty", "fierce", "happy", "icy", "jumpy", "crazy", "clever", "shy", "chilly", "noisy",
	"sassy", "dizzy", "dusty", "lucky", "spicy", "tiny", "chunky", "frosty", "groovy", "breezy",
	"crispy", "cranky", "snappy", "picky", "sunny", "quirky", "moody", "vivid", "jazzy", "cloudy",
	"rowdy", "proud", "snug", "snazzy", "frisky", "chubby", "kooky", "peppy", "bouncy", "zany",
	"flashy", "neat", "dopey", "nifty", "witty", "silly", "clumsy", "perky", "tidy", "glossy",
	"yappy", "punchy", "nervy", "twinkly", "toasty", "gritty", "wavy", "flaky", "feisty", "giddy",
	"skippy", "minty", "chilly", "dandy", "fizzy", "cheery", "plucky", "nutty", "spacey", "dreamy",
	"smug", "nosy", "zippy", "charming", "jolly", "sparkly", "classy", "quirky", "whimsical", "chic",
}

var animals = []string{
	"fox", "panda", "tiger", "koala", "eagle", "otter", "lion", "rabbit", "wolf", "moose",
	"cat", "dog", "falcon", "badger", "sloth", "kangaroo", "toucan", "squid", "chimp", "buffalo",
	"penguin", "lizard", "pony", "boar", "giraffe", "hyena", "lemur", "duck", "turtle", "bat",
	"crab", "swan", "bee", "antelope", "armadillo", "parrot", "beaver", "raccoon", "whale", "stingray",
	"frog", "newt", "yak", "reindeer", "hedgehog", "meerkat", "hamster", "macaw", "donkey", "platypus",
	"rooster", "rhino", "sheep", "gorilla", "iguana", "ferret", "seal", "crow", "horse", "ox",
	"vulture", "goose", "pigeon", "turkey", "wasp", "dragonfly", "grasshopper", "mole", "gecko", "lobster",
	"starfish", "eel", "clam", "snail", "octopus", "caribou", "coyote", "gazelle", "mallard", "panther",
	"manatee", "bear", "wolverine", "bluejay", "salamander", "elk", "dingo", "chinchilla", "skunk", "hummingbird",
	"quokka", "porcupine", "caterpillar", "mongoose", "narwhal", "jaguar", "alpaca", "stingray", "dragon", "phoenix",
}
var consonants = []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "r", "s", "t", "v", "z"}
var vowels = []string{"a", "e", "i", "o", "u", "y"}

func Seed(api string) {
	total := 100_000
	workers := 20

	var wg sync.WaitGroup
	jobs := make(chan int)

	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range jobs {
				user := map[string]string{
					"username": username(),
					"password": password(),
				}

				jsonData, err := json.Marshal(user)
				if err != nil {
					log.Fatal("Failed to marshal JSON:", err)
				}

				res, err := http.Post(api+"users", "application/json", bytes.NewBuffer(jsonData))
				if err != nil {
					fmt.Println(err)
					continue
				}
				res.Body.Close()
			}
		}()
	}

	for i := range total {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

func username() string {
	adjective := adjectives[rand.Intn(len(adjectives))]
	animal := animals[rand.Intn(len(animals))]
	number := rand.Intn(1000)

	return fmt.Sprintf("%s%s%d", adjective, animal, number)
}

func password() string {
	var b strings.Builder

	for range 10 {
		b.WriteString(consonants[rand.Intn(len(consonants))])
		b.WriteString(vowels[rand.Intn(len(vowels))])
		if rand.Intn(4) == 0 { // occasional extra consonant
			b.WriteString(consonants[rand.Intn(len(consonants))])
		}
	}

	// optional extras for flair
	if rand.Intn(2) == 1 {
		b.WriteString(fmt.Sprintf("%d", rand.Intn(100)))
	}

	return b.String()
}
