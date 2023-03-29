package order

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	inv "service/stores/case1/internal/inventory/models"
	"service/stores/case1/internal/models"
)

// Hlp me to apply th comments into the code to improve it reduce unsused code
const (
	WorkersLenght = 5
)

type Processor struct {
	jobChannel chan string
	done1      chan *Finance
	done2      chan *Accounting
	done3      chan *Admin
	done5      chan *Marketing
	workers1   []*Finance
	workers2   []*Accounting
	workers3   []*Admin
	workers5   []*Marketing
}

func (a *Finance) processJob1(data string, done chan *Finance) {
	// Use the data and process the job
	go func() {
		fmt.Println("Process Job 1 ")
		fmt.Printf("\n")
		done <- a
	}()
}

func (b *Accounting) processJob2(data string, done chan *Accounting) {
	// Use the data and process the job
	go func() {
		fmt.Println("Process Job 2 ")
		fmt.Printf("\n")
		done <- b
	}()
}

func (c *Admin) processJob3(data string, done chan *Admin) {
	// Use the data and process the job
	go func() {
		fmt.Println("Process Job 3 ")
		fmt.Printf("\n")
		done <- c
	}()
}

func (s *Marketing) processJob5(data string, done chan *Marketing) {
	go func() {
		fmt.Println("Process Job 5 ")
		fmt.Printf("\n")
		done <- s
	}()

}

func GetProcessor(a Finance, b Accounting, c Admin, s Marketing) *Processor {
	p := &Processor{
		jobChannel: make(chan string),
		workers1:   make([]*Finance, WorkersLenght),
		workers2:   make([]*Accounting, WorkersLenght),
		workers3:   make([]*Admin, WorkersLenght),
		workers5:   make([]*Marketing, WorkersLenght),
		done1:      make(chan *Finance),
		done2:      make(chan *Accounting),
		done3:      make(chan *Admin),
		done5:      make(chan *Marketing),
	}
	allProducts := make([]inv.Product, len(s.Inventory.InventoryProduct))

	for i := 0; i < WorkersLenght; i++ {
		a := &Finance{
			ID: s.ID,
			Inventory: inv.Inventory{
				ID:               s.Inventory.ID,
				InventoryProduct: allProducts,
				CreatedAt:        s.Inventory.CreatedAt,
			},
			StoreInfomation: models.StoreInfomation{},
			CreatedAt:       s.CreatedAt,
		}
		b := &Accounting{
			ID: s.ID,
			Inventory: inv.Inventory{
				ID:               s.Inventory.ID,
				InventoryProduct: allProducts,
				CreatedAt:        s.Inventory.CreatedAt,
			},
			StoreInfomation: models.StoreInfomation{},
			CreatedAt:       s.CreatedAt,
		}
		c := &Admin{
			ID: s.ID,
			Inventory: inv.Inventory{
				ID:               s.Inventory.ID,
				InventoryProduct: allProducts,
				CreatedAt:        s.Inventory.CreatedAt,
			},
			StoreInfomation: models.StoreInfomation{},
			CreatedAt:       s.CreatedAt,
		}
		e := &Marketing{
			ID: s.ID,
			Inventory: inv.Inventory{
				ID:               s.Inventory.ID,
				InventoryProduct: allProducts,
				CreatedAt:        s.Inventory.CreatedAt,
			},
			StoreInfomation: models.StoreInfomation{},
			Employee:        models.EmployeeInternalInformation{},
			CreatedAt:       s.CreatedAt,
		}

		writeJSONToFile(context.Background(), a, "a.json", 0777)
		writeJSONToFile(context.Background(), a, "b.json", 0777)
		writeJSONToFile(context.Background(), a, "c.json", 0777)
		writeJSONToFile(context.Background(), a, "e.json", 0777)

		p.workers1[i] = a
		p.workers2[i] = b
		p.workers3[i] = c
		p.workers5[i] = e
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
					e := p.workers5[4]

					a.processJob1(<-p.jobChannel, p.done1)
					b.processJob2(<-p.jobChannel, p.done2)
					c.processJob3(<-p.jobChannel, p.done3)
					e.processJob5(<-p.jobChannel, p.done5)
				}
			case a := <-p.done1:
				p.workers1 = append(p.workers1, a)

			case b := <-p.done2:
				p.workers2 = append(p.workers2, b)

			case c := <-p.done3:
				p.workers3 = append(p.workers3, c)

			case e := <-p.done5:
				p.workers5 = append(p.workers5, e)
			}

		}
	}()

}

func writeJSONToFile(ctx context.Context, data interface{}, filename string, perm os.FileMode) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	fileContents, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	_, err = file.Write(fileContents)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
