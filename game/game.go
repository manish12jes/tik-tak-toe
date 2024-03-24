package game

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	deque "github.com/edwingeng/deque/v2"

	board "tik-tak-toe/board"
	symb "tik-tak-toe/pieces"
	pl "tik-tak-toe/player"
)

func Start(p1, p2 pl.Player, b board.Board) {
	dq := deque.NewDeque[pl.Player]()
	dq.PushBack(p1)
	dq.PushBack(p2)

	boardSize := len(b)
	var gameCompleted = false
	maxMove := boardSize * boardSize
	count := 0

	for !gameCompleted {
		if count >= maxMove {
			fmt.Println("No new move possible, closing game.")
			gameCompleted = true
			continue
		}
		player := dq.PopFront()
		i, j, err := getUserMove(player)
		if err != nil {
			fmt.Println(err.Error())
			dq.PushFront(player)
			continue
		}

		if !b.ValidMove(i, j) {
			fmt.Println("Invalid move")
			dq.PushFront(player)
			continue
		}

		b.AddPieceToBoad(player.Piece, i, j)
		b.DisplayBoard()
		if player.IsWinner(b, player.Piece) {
			fmt.Println("Hey ", player.Name, " You won the game.")
			gameCompleted = true
			continue
		}
		dq.PushBack(player)
		count++
	}
}

func getUserMove(player pl.Player) (int, int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(player.Name, ", Your turn : ", "Enter i, j value: ")
	input, _ := reader.ReadString('\n')
	if len(strings.TrimSpace(input)) == 0 {
		return -1, -1, errors.New("invalid value entered")
	}
	inputSlice := strings.Split(input, ",")
	if len(inputSlice) != 2 {
		return -1, -1, errors.New("invalid value entered")
	}
	i, _ := strconv.Atoi(strings.TrimSpace(inputSlice[0]))
	j, _ := strconv.Atoi(strings.TrimSpace(inputSlice[1]))
	return i, j, nil
}

func Init() {
	p1, p2 := initPlayers()
	b := initBoard()
	fmt.Println("Starting game...")
	b.DisplayBoard()
	Start(*p1, *p2, b)
}

func initPlayers() (p1, p2 *pl.Player) {
	var (
		name, symbol string
		sign1, sign2 *symb.Pieces
	)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Player1 Name : ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Print(name, " pick your sumbol X or O : ")
	symbol, _ = reader.ReadString('\n')
	symbol = strings.TrimSpace(symbol)

	symbol = strings.ToLower(symbol)
	if symbol == "x" {
		sign1 = symb.NewPiece("X")
		sign2 = symb.NewPiece("O")
	} else {
		sign1 = symb.NewPiece("O")
		sign2 = symb.NewPiece("X")
	}

	p1 = pl.NewPlayer(name, *sign1)

	fmt.Print("Player2 Name : ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	p2 = pl.NewPlayer(name, *sign2)
	return p1, p2
}

func initBoard() board.Board {
	boardSize := 3
	b := board.NewBoard(boardSize)
	return b
}
