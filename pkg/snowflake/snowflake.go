package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init(machineId int64) (err error) {
	//var st = startTime
	//st, err = time.Parse("2021-01-01", startTime)
	//if err != nil {
	//	return
	//}
	//sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineId)
	return
}

func GetId() int64 {
	return node.Generate().Int64()
}

//snowflake.Init(1)
//fmt.Println(snowflake.GetId())
