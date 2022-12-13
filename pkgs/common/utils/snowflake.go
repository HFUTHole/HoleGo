package utils

import "github.com/bwmarrin/snowflake"

var node *snowflake.Node

func InitSnowflakeNode() {
	node, _ = snowflake.NewNode(1)
}

func NextSnowflake() int64 {
	return int64(node.Generate())
}
