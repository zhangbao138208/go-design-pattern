package abstractfactory

import "fmt"

type OrderMainDAO interface {
	SaveOrderMain()
}

type OrderMainDetail interface {
	SaveOrderDetail()
}

type DAOFactory interface {
	CreateOrderMainDAO() 	OrderMainDAO
	CreateOrderMainDetail() OrderMainDetail
}

type RDBMainDAO struct {

}

func (r RDBMainDAO)SaveOrderMain()  {
	fmt.Println("rdb main save")
}

type RDBDetailDAO struct {

}

func (d RDBDetailDAO) SaveOrderDetail()  {
	fmt.Println("rdb detail save")
}

type RDBDAOFactory struct {

}

func (f *RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
   return &RDBMainDAO{}
}

func (f *RDBDAOFactory) CreateOrderMainDetail() OrderMainDetail {
	return &RDBDetailDAO{}
}

type XMLMainDAO struct {

}

func (x XMLMainDAO) SaveOrderMain()  {
	fmt.Println("xml main save")
}

type XMLDetailDAO struct {

}

func (x XMLDetailDAO) SaveOrderDetail()  {
	fmt.Println("xml detail save")
}

type XMLDABFactory struct {

}

func (x XMLDABFactory) CreateOrderMainDAO() OrderMainDAO  {
	return &XMLMainDAO{}
}

func (x XMLDABFactory) CreateOrderMainDetail() OrderMainDetail  {
	return XMLDetailDAO{}
}
