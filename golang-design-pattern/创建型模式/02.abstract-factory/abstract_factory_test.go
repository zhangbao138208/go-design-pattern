package abstractfactory

func GetMainAndDetail(factory DAOFactory) {
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderMainDetail().SaveOrderDetail()
}

func ExampleRDBDAOFactory() {
	var factory DAOFactory
	factory = &RDBDAOFactory{}
	GetMainAndDetail(factory)
	// Output:
	// rdb main save
	// rdb detail save
}

func ExampleXMLDABFactory() {
	var factory DAOFactory
	factory = &XMLDABFactory{}
	GetMainAndDetail(factory)
	// Output:
	// xml main save
	// xml detail save
}
