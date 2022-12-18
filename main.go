package main

import (
  "fmt"
  "math/rand"
  "os"
  "strconv"
  "sync"
  "time"
)

var waitGroup sync.WaitGroup
var mutex sync.Mutex
var winner = 0

func main() {
  // quantidade padrão de raias
  quantityTrack := 6

  // verifica se foi passado algum argumento pela CLI
  // caso tenha sido passado e valor for válido ele é atribuído a variável `quantityTrack`
  if len(os.Args) > 1 {
    cliArg, err := strconv.Atoi(os.Args[1])
    if err != nil {
      os.Exit(3)
    }

    quantityTrack = cliArg
  }

  // adiciona a quantidade de raias para o `waitGroup`
  waitGroup.Add(quantityTrack)

  // cria a quantidade de raias definidas como gorotinas
  for i := 0; i < quantityTrack; i++ {
    go track(i + 1)
  }

  // manda o `waitGroup` esperar a finalização das gorotinas
  waitGroup.Wait()

  // imprimi o vencedor da corrida
  fmt.Printf("o vencedor foi: raia %d\n", winner)
}

func track(id int) {
  // esse `for` simula os corredores que vão revezar (são 4 corredores)
  for i := 0; i < 4; i++ {
    // caso seja o primeiro corredor mostramos a mensagem que começou a correr
    if i == 0 {
      fmt.Printf("(raia %d) começou a correr: 1 corredor\n", id)
    }
    // manda a gorotina dormir por um tempo aleatório que pode ir de 1 à 4 segundos
    time.Sleep(time.Duration(rand.Intn(4-1)+1) * time.Second)
    // caso seja o último corredor mostramos a mensagem que terminou de correr
    // senão mostramos que passou o bastão para o próximo corredor
    if i == 3 {
      fmt.Printf("(raia %d) terminou a corrida\n", id)
    } else {
      fmt.Printf("(raia %d) passou o bastão para: %d corredor\n", id, i+2)
    }
  }

  // bloqueamos o acesso e verificamos se o valor de winner ainda é o inicial (winner = 0),
  // se for significa que ninguem terminou ainda.
  // atribuimos o `id` da raia para a variável `winner` para indicar o vencedor,
  // quando liberarmos o acesso a winner e as outras gorotinas tentarem
  // acessar o valor de `winner` ele já não vai mais ser o inicial e não vai alterar o `winner`
  // e depois liberamos o acesso a `winner`
  // e por fim avisamos ao `waitGroup` que terminou a gorotina
  mutex.Lock()
  if winner == 0 {
    winner = id
  }
  mutex.Unlock()

  waitGroup.Done()
}
