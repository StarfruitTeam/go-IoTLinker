

type Tag struct {
	name string
	topic string
	tag_type string
	subscription integer
	data_Type string
	created datetime
}

type Device struct {
	id string
	name string
	protocol string
	tags []string
	location string
	timeout integer
	subscription integer
	data_type string
	created datetime
}

