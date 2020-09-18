package application

type AppConfig struct {
	AmqpURL  string
	Consumer ConsumerConf
}

type ConsumerConf struct {
	Tag      string
	Exchange ExchangeConf
}

type ExchangeConf struct {
	Name       string
	Type       string
	BindingKey string
	QueName    string
}
