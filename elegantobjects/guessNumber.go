package elegantobjects

import (
	"fmt"
	"math/rand"
)

// PlayGuessingGame written object oriented in Elegant Objects style from
// Yegor Bugayenko. The top level object starts the run using .Say(). All
// other objects call each other and pass them self to each other.
func PlayGuessingGame() {
	secret := NewSecret()
	NewFarewell(
		NewAttempts(
			NewVerboseDiff(
				NewDiff(
					secret,
					NewGuess(),
				),
			),
			5,
		),
		secret,
	).Say()
}

type Secret struct {
	num int
}

func NewSecret() Secret {
	return Secret{rand.Intn(99)}
}

func (s Secret) Number() int {
	return s.num
}

type Guess struct{}

func NewGuess() Guess {
	return Guess{}
}

func (g Guess) Number() int {
	fmt.Printf("Guess a number in 0..99 range: ")

	var num int
	if _, err := fmt.Scanf("%d", &num); err != nil {
		panic(err)
	}
	fmt.Println()

	return num
}

type Diff struct {
	secret Secret
	guess  Guess
}

func NewDiff(n Secret, x Guess) Diff {
	return Diff{n, x}
}

func (d Diff) Number() int {
	return d.guess.Number() - d.secret.Number()
}

type VerboseDiff struct {
	diff Diff
}

func NewVerboseDiff(d Diff) VerboseDiff {
	return VerboseDiff{d}
}

func (v VerboseDiff) Number() int {
	x := v.diff.Number()
	if x < 0 {
		fmt.Println("Too small.")
	} else if x > 0 {
		fmt.Println("Too big.")
	} else {
		fmt.Println("Bingo!")
	}
	return x
}

type Attempts struct {
	diff VerboseDiff
	max  int
}

func NewAttempts(d VerboseDiff, t int) Attempts {
	return Attempts{d, t}
}

func (a Attempts) Matches() bool {
	t := 0
	for t <= a.max && a.diff.Number() != 0 {
		t++
	}
	return t <= a.max
}

type Farewell struct {
	attempts Attempts
	secret   Secret
}

func NewFarewell(a Attempts, s Secret) Farewell {
	return Farewell{a, s}
}

func (f Farewell) Say() {
	if !f.attempts.Matches() {
		fmt.Println("You lost, sorry. It was: ", f.secret.Number())
	}
	fmt.Println("Thanks for playing with me!")
}
