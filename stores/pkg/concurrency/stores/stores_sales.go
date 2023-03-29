package stores

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const (
	WorkersLenght = 4
)

type Processor struct {
	jobChannel chan string
	done1      chan *Worker1
	done2      chan *Worker2
	done3      chan *Worker3
	done4      chan *Worker4
	done5      chan *CustomerOrderCompleted
	workers1   []*Worker1
	workers2   []*Worker2
	workers3   []*Worker3
	workers4   []*Worker4
	workers5   []*CustomerOrderCompleted
}
type Worker1 struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	FloatNum    float64   `json:"float_num"`
	CreatedAt   string    `json:"created_at"`
}

type Worker2 struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	FloatNum    float64   `json:"float_num"`
	CreatedAt   string    `json:"created_at"`
}

type Worker3 struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	FloatNum    float64   `json:"float_num"`
	CreatedAt   string    `json:"created_at"`
}

type Worker4 struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	FloatNum    float64   `json:"float_num"`
	CreatedAt   string    `json:"created_at"`
}

func (a *Worker1) processJob1(data string, done chan *Worker1) {
	// Use the data and process the job
	go func() {
		fmt.Println("Working on data ", data, a.Name)
		fmt.Println("Working on data 2", data, a.Description)
		fmt.Println("Working on FloatNum 1", data, a.FloatNum)
		fmt.Printf("\n")
		done <- a
	}()
}

func (b *Worker2) processJob2(data string, done chan *Worker2) {
	// Use the data and process the job
	go func() {
		fmt.Println("Working on data ", data, b.Name)
		fmt.Println("Working on data 2", data, b.Description)
		fmt.Println("Working on FloatNum 1", data, b.FloatNum)
		fmt.Printf("\n")
		done <- b
	}()
}

func (c *Worker3) processJob3(data string, done chan *Worker3) {
	// Use the data and process the job
	go func() {
		fmt.Println("Working on data ", data, c.Name)
		fmt.Println("Working on data 2", data, c.Description)
		fmt.Println("Working on FloatNum 1", data, c.FloatNum)
		fmt.Printf("\n")
		done <- c
	}()
}

func (d *Worker4) processJob4(data string, done chan *Worker4) {
	// Use the data and process the job
	go func() {
		fmt.Println("Working on data ", data, d.Name)
		fmt.Println("Working on data 2", data, d.Description)
		fmt.Println("Working on FloatNum 1", data, d.FloatNum)
		fmt.Printf("\n")
		done <- d
	}()
}

func (s *CustomerOrderCompleted) processJob5(ctx *fiber.Ctx, data string, done chan *CustomerOrderCompleted) {
	go func() {
		fmt.Println("Process Job 5 ")
		fmt.Printf("\n")
		done <- s
	}()

}

func OperationA(a float64) float64 {
	return a * 10 / 2
}

func OperationB(b float64) float64 {
	return b * 10 / 2
}

func OperationC(c float64) float64 {
	return c * 10 / 2
}

func OperationD(d float64) float64 {
	return d * 10 / 2
}

func GetProcessor(a Worker1, b Worker2, c Worker3, d Worker4, s CustomerOrderCompleted) *Processor {
	p := &Processor{
		jobChannel: make(chan string),
		workers1:   make([]*Worker1, WorkersLenght),
		workers2:   make([]*Worker2, WorkersLenght),
		workers3:   make([]*Worker3, WorkersLenght),
		workers4:   make([]*Worker4, WorkersLenght),
		workers5:   make([]*CustomerOrderCompleted, WorkersLenght),
		done1:      make(chan *Worker1),
		done2:      make(chan *Worker2),
		done3:      make(chan *Worker3),
		done4:      make(chan *Worker4),
		done5:      make(chan *CustomerOrderCompleted),
	}
	for i := 0; i < WorkersLenght; i++ {
		a := &Worker1{ID: uuid.New(), Name: a.Name, Description: a.Description, FloatNum: OperationA(a.FloatNum), CreatedAt: time.Now().GoString()}
		b := &Worker2{ID: uuid.New(), Name: b.Name, Description: b.Description, FloatNum: OperationA(b.FloatNum), CreatedAt: time.Now().GoString()}
		c := &Worker3{ID: uuid.New(), Name: c.Name, Description: c.Description, FloatNum: OperationA(c.FloatNum), CreatedAt: time.Now().GoString()}
		d := &Worker4{ID: uuid.New(), Name: d.Name, Description: d.Description, FloatNum: OperationA(d.FloatNum), CreatedAt: time.Now().GoString()}

		filea, _ := json.MarshalIndent(a, "", " ")
		os.WriteFile("a.json", filea, 0777) //0644

		fileb, _ := json.MarshalIndent(b, "", " ")
		os.WriteFile("b.json", fileb, 0777) //0644

		filec, _ := json.MarshalIndent(c, "", " ")
		os.WriteFile("c.json", filec, 0777) //0644

		filed, _ := json.MarshalIndent(d, "", " ")
		os.WriteFile("d.json", filed, 0777) //0644

		p.workers1[i] = a
		p.workers2[i] = b
		p.workers3[i] = c
		p.workers4[i] = d
	}

	p.startProcess()

	return p
}

func (p *Processor) startProcess() {
	go func() {
		for {
			select {
			default:
				if len(p.workers1) > 0 {
					a := p.workers1[0]
					b := p.workers2[1]
					c := p.workers3[2]
					d := p.workers4[3]

					a.processJob1(<-p.jobChannel, p.done1)
					b.processJob2(<-p.jobChannel, p.done2)
					c.processJob3(<-p.jobChannel, p.done3)
					d.processJob4(<-p.jobChannel, p.done4)
				}
			case a := <-p.done1:
				p.workers1 = append(p.workers1, a)

			case b := <-p.done2:
				p.workers2 = append(p.workers2, b)

			case c := <-p.done3:
				p.workers3 = append(p.workers3, c)

			case d := <-p.done4:
				p.workers4 = append(p.workers4, d)
			}
		}
	}()

}
