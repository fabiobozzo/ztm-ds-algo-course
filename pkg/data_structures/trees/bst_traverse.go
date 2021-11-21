package trees

import (
	"strconv"
)

func bstToMap(rootNode *BSTNode) map[string]interface{} {
	output := map[string]interface{}{
		"value": strconv.Itoa(rootNode.Value),
	}

	output["left"] = nil
	if rootNode.Left != nil {
		output["left"] = bstToMap(rootNode.Left)
	}

	output["right"] = nil
	if rootNode.Right != nil {
		output["right"] = bstToMap(rootNode.Right)
	}

	return output
}

//func PrintBST(bst BST) (string, error) {
//	bstMap := bstToMap(bst.GetRoot())
//
//	bytes, err := json.Marshal(bstMap)
//	if err != nil {
//		return "", err
//	}
//
//	return string(bytes), nil
//}
