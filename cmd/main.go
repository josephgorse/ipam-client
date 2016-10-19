package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/RackHD/ipam/resources"
	"github.com/ekaley/client"
)

var ipam string

func init() {
	flag.StringVar(&ipam, "ipam", "127.0.0.1:8000", "address:port of ipam")
}

func main() {
	//Bootstrap
	c := api.NewClient(ipam)

	//New Pools
	p := c.Pools()

	rPool := &resources.PoolV1{
		Name:     "Pool1",
		Metadata: "yodawg I heard you like interfaces",
	}

	pool, err := p.CreateShowPool(*rPool)
	if err == nil {
		fmt.Printf("%+v\n", pool)
	}

	time.Sleep(1 * time.Second)
	//New Subnets
	s := c.Subnets()

	rSubnet := &resources.SubnetV1{
		Name:  "Subnet1",
		Pool:  pool.ID,
		Start: "192.168.1.10",
		End:   "192.168.1.20",
	}

	subnet, err := s.CreateShowSubnet(pool.ID, *rSubnet)
	if err == nil {
		fmt.Printf("%+v\n", subnet)
	}

	time.Sleep(1 * time.Second)
	//New Reservations
	r := c.Reservations()

	rReservation := &resources.ReservationV1{
		Name:   "Reservation1",
		Subnet: subnet.ID,
	}

	reservation, err := r.CreateShowReservation(subnet.ID, *rReservation)
	if err == nil {
		fmt.Printf("%+v\n\n", reservation)
	}

	time.Sleep(1 * time.Second)
	//New Leases
	l := c.Leases()

	pI, err := p.Index()
	if err != nil {
	}
	fmt.Printf("%+v\n\n", pI)
	time.Sleep(1 * time.Second)

	sI, err := s.Index(pool.ID)
	if err != nil {
	}
	fmt.Printf("%+v\n\n", sI)
	time.Sleep(1 * time.Second)

	rI, err := r.Index(subnet.ID)
	if err != nil {
	}
	fmt.Printf("%+v\n\n", rI)
	time.Sleep(1 * time.Second)

	lI, err := l.Index(reservation.ID)
	if err != nil {
	}
	fmt.Printf("%+v\n\n", lI)

	fmt.Println("Finished")
	return
}
