package main
import (
	"fmt"
	"runtime"
	"sync"
)

func main(){

	contador := 0
	totalGoroutines1 := 10

	var wg sync.WaitGroup 
	wg.Add(totalGoroutines1)

	fmt.Println(" Sem Mutex ")
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())

	// Sem Mutex
	for i := 0; i < totalGoroutines1; i++ {
		go func(){
			v := contador
			runtime.Gosched()
			fmt.Println(i)
			v++
			contador = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("Contador: ", contador)
	fmt.Println(" ============================== ")
	
	// Com Mutex
	fmt.Println(" Com Mutex ")
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	totalGoroutines2 := 20
	wg.Add(totalGoroutines2)
	var mu sync.Mutex

	for i := 0; i < totalGoroutines2; i++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor final:", contador)
	fmt.Println(" ============================== ")
}

/*
- Agora vamos resolver a race condition do programa anterior utilizando mutex.
- Mutex é mutual exclusion, exclusão mútua.
- Utilizando mutex somente uma thread poderá utilizar a variável contador de cada vez, e as outras deve aguardar sua vez "na fila."
- Na prática:
    - type Mutex
        - func (m *Mutex) Lock()
        - func (m *Mutex) Unlock()
- RWMutex

- Como funciona?
    - Vá em github.com e crie um repo
    - Crie uma pasta com o mesmo nome no seu $GOPATH
        - $GOPATH/src/github.com/[username]/[repo]
    - Rode “git init” nesta pasta
    - Adicione arquivos, e.g. README.md e .gitignore
    - git add -A
    - git commit -m “here’s some commit message”
    - git remote add origin git@github.com:username/repo.git
    - git push -u origin master
- Comandos:
    - git status
    - git add --all
    - git commit -m "mensagem"
    - git push

*/