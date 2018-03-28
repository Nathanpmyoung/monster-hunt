package main

import "fmt"
import "math/rand"
import "math"
import "time"

type Game struct{
	Hand []string `json:"hand"`
	InPlay []string `json:"inplay"`
	Deck []string `json:"deck"`
	Discard []string `json:"discard"`
}

func sumCoins(cards[]string) int {
	totalx := 0
	for _, value := range cards {
		switch value {
		case "Copper":
			totalx += 1
		case "Silver":
			totalx += 2
		case "Gold":
			totalx += 3
		}
	}
    return totalx
}

func shuffle(a []string) []string {
    b := []string{}
    k := len(a)
    for i := 0; i < k; i++ {
    j := rand.Intn(len(a))
    b = append(b, a[j])
    a = append(a[:j], a[j+1:]...)
    }
    return b
}

func lookThrough(name string,cards []string) int {
	howMany := 0
	for _, card := range cards {
		if card == name {
			howMany += 1
		}
	}
	return howMany
}

func smithyCondition(game Game) bool {
	ratio := 0.2
	wholeDeck := append(game.Hand, game.InPlay, game.Deck, game.Discard...)
	if (lookThrough("Smithy", wholeDeck) < 10 && lookThrough("Smithy", wholeDeck)/ len(wholeDeck) < ratio {
		return true
	}
	return false
}

//draw function
func draw(number int,game Game) (Game){
    for i := 0; i < number; i++ {
        if len(game.Deck) > 0 {
            game.Hand = append(game.Hand, game.Deck[:1]...)
            game.Deck = game.Deck[1:]
        } else {
            if len(game.Discard) > 0 {
                game.Deck = shuffle(game.Discard)
                game.Hand = append(game.Hand, game.Deck[:1]...)
                game.Deck = game.Deck [1:]
                game.Discard = nil
            } else {
                return game
            }
        }
    }
    return game
}

func main() {

rand.Seed(time.Now().UTC().UnixNano())

    totalturns := 0
	totalstdev := 0.0
    n := 10
	//strat := {"","","","Silver","Smithy","","Gold","","Province"}

	for i := 0; i < n; i++ {
		turns := 0
		game := Game{Hand:[]string{}, InPlay:[]string{}, Deck:[]string{}, Discard:[]string{"Copper","Copper","Copper","Copper","Copper","Copper","Copper","Estate","Estate","Estate"}}
		fmt.Println(game.Hand)
		for provinces := 0; provinces < 5; {
			turns ++
	        game = draw(5, game)
			if lookThrough("Smithy", game.Hand) > 0 {
				fmt.Println("Play Smithy:", game)
				// put smithy in play game.Discard = append([]string{"Smithy"}, game.Discard...)
				// remove Smithy from game.Hand
				game = draw(3, game)
			}
	        if sumCoins(game.Hand) > 7 {
				provinces ++
				game.Discard = append([]string{"Province"}, game.Discard...)
			} else if sumCoins(game.Hand) > 5 {
				game.Discard = append([]string{"Gold"}, game.Discard...)
			} else if sumCoins(game.Hand) > 3 && smithyCondition(game) {
				game.Discard = append([]string{"Smithy"}, game.Discard...)
				//If the card density is higher
				//card totals?
			} else if sumCoins(game.Hand) > 2 {
				game.Discard = append([]string{"Silver"}, game.Discard...)
			}
	    	fmt.Println(game)
			game.Discard = append(game.Hand, game.Discard...)
			game.Hand = nil
		}
		totalturns += turns
		totalstdev += math.Pow((float64(turns)-21.773),2)
    }

	averageturns := float64(totalturns)/float64(n)
	fullstdev := math.Pow((totalstdev/(float64(n)-1)),.5)
    fmt.Println("Total turns:", totalturns, "n:", n, "Average turns:", averageturns, "Standard dev", fullstdev)
}
