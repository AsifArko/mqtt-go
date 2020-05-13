package models

/*
	These are the Neo4j Nodes for a person to be used in routines for finding friends and so on ...

	This model will only be used for getting data from neo4j bolt database and
	will only be used to map with our primary database in couchbase
*/

type Node struct {
	NodeId     int64    `json:"NodeIdentity"`
	Labels     []string `json:"Labels"`
	Properties Person   `json:"properties"`
}

type Person struct {
	Id       string `json:"id"`
	Dob      string `json:"dob"`
	Name     string `json:"name"`
	HomeTown string `json:"hometown"`
}
