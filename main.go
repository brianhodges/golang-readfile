package main
import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
    "./team"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    var teams []team.Team
    var filename = "top3list.csv"
    
    file, err := os.Open("./"+filename)
    check(err)
    defer file.Close()
    
    f, err := os.Stat("./"+filename)
    check(err)
    
    fmt.Println()
    fmt.Println("Filename: " + filename)
    fmt.Printf("Modified At: " + f.ModTime().String())
    fmt.Println()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        s := strings.Split(scanner.Text(), ",")
        n, w := s[0], s[1]
        wins, err := strconv.Atoi(w)
        check(err)
        t := team.Team{
            Name: n,
            Super_Bowl_Wins: wins,
        }
        teams = append(teams, t)
    }
    
    fmt.Println()
	for _, team := range teams {
        line := len("| Team: " + team.Name + " |") - 2
        fmt.Println("+" + strings.Repeat("-", line) + "+")
		fmt.Println("| Team: " + team.Name + " |")
        ws := len("| Super Bowl Wins: " + strconv.Itoa(team.Super_Bowl_Wins))
        fmt.Println("| Super Bowl Wins: " + strconv.Itoa(team.Super_Bowl_Wins) + strings.Repeat(" ", line - ws) + " |")
        fmt.Println("+" + strings.Repeat("-", line) + "+")
        fmt.Println()
	}
}