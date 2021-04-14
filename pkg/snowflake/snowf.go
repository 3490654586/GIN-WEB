package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
	"time"
)


//雪花算法生成id
var node *sf.Node
func Init(startTime string,macineID int64)error{
	st ,err :=  time.Parse("2006-01-01",startTime)
	if err != nil {
		return err
	}

	sf.Epoch =st.UnixNano()/1000000

	node,_= sf.NewNode(macineID)
	return nil
}

func GetId()int64{
	return node.Generate().Int64()
}