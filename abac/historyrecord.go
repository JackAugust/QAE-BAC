package abac

import (
	"math"
	"reflect"
	"sort"
	"strings"
)

// var PoolNum = 10
// var Attr_list = []string{"D", "C", "A", "X", "B", "Y"}
// // var Attr_list = []string{"D", "C", "A", "X", "B", "Y"}

// // var Attr_list = []string{"A", "B", "C", "D", "X", "Y", "O", "E"}

// 每个属性可以取的值的个数
var AttrinNum = 2
var PoolNum = 10
var Attr_list = []string{"A", "B", "E", "C", "D", "X", "Y", "O", "P", "Q", "R", "S"}
var Sub_Attr = []string{"A", "B", "C", "D"}
var Obj_Attr = []string{"X", "Y", "P", "Q", "R", "S"}
var MyHistoryPool = NewHistoryPool(PoolNum)

// 每种属性的比重，按v进行节点生成
// var attr_weights1 = map[string]int{"Owner": 10, "Env": 5, "SubRules": 3}
var attr_weights2 = map[string]float64{"X": 1.0, "Y": 1.0, "A": 1.0, "B": 1.0, "C": 1.0, "D": 1.0, "E": 1.0}

// 每条历史记录的数据格式
type HistoryRecord struct {
	FeatureNum int    // number
	Result     bool   //1 is true;0 is false； 作为标签
	Attr1      string `json:"attr1"`
	Attr2      string `json:"attr2"`
	Attr3      string `json:"attr3"`
	Attr4      string `json:"attr4"`
	Attr5      string `json:"attr5"`
	Attr6      string `json:"attr6"`
	Attr7      string `json:"attr7"`
	Attr8      string `json:"attr8"`
	Attr9      string `json:"attr9"`
	Attr10     string `json:"attr10"`
	Attr11     string `json:"attr11"`
	Attr12     string `json:"attr12"`
}

// 历史记录池，存放历史记录，是不断更新的
type HistoryPool struct {
	Pool           []*HistoryRecord
	MaxHistorySize int
	// Head           int
	// Tail           int
	Len int
	Cur int
}

func NewHistoryRecord(result bool) *HistoryRecord {
	return &HistoryRecord{
		Result:     result,
		FeatureNum: 0,
	}

}

func NewHistoryPool(maxHistorySize ...int) *HistoryPool {
	size := 100
	if len(maxHistorySize) != 0 {
		size = maxHistorySize[0]
	}
	return &HistoryPool{
		Pool:           make([]*HistoryRecord, size),
		MaxHistorySize: size,
		// Head:           0,
		// Tail:           -1,
		Len: 0,
		Cur: 0,
	}
}

func (pool *HistoryPool) IsEmpty() bool {
	return pool.Len == 0
}

func (pool *HistoryPool) IsFull() bool {
	return pool.Len == pool.MaxHistorySize
}

// TODO：缓存有三种淘汰策略：LRU(最近最少使用)、LFU(最少使用)、FIFO(先进先出)
// 考虑到这里实现的是历史记录池，与访问与否和命中率无关，所以FIFO，后期的缓存池应该为LRU
// 如果有必要再修改或增加？（比如效率提高不明显）

// Push 循环添加数据
func (pool *HistoryPool) Push(record *HistoryRecord) {
	if pool.IsFull() {
		// panic("队列已满")
		// 队列已满时，从头开始覆盖
		pool.Cur = pool.Cur % pool.MaxHistorySize
		pool.Len--

		// return
	}
	// pool.Pool = append(pool.Pool, record)
	// fmt.Println("Pushing ", pool.Cur)
	pool.Pool[pool.Cur] = record
	pool.Cur++
	pool.Len++
}
func (pool *HistoryPool) GetAll() []*HistoryRecord {
	return pool.Pool
}

// Pop 将第一个元素删除
// func (pool *HistoryPool) Pop() *HistoryRecord {
// 	if pool.IsEmpty() {
// 		panic("HistoryPool is empty")
// 	}
// 	val := pool.Pool[0]
// 	// pool.Head++
// 	pool.Len--
// 	pool.Pool = pool.Pool[1:]
// 	return val
// }

// func (pool *HistoryPool) Push(record *HistoryRecord) {
// 	if len(pool.Pool) >= pool.MaxHistorySize {

// 	}
// }

// 仅在这里对attr_weights进行排序，排序后写入attr_list其他地方不涉及到排序

// 按照value对k排序，然后把k存了
// var attr_list = []string{"org", "role", "researce", "action"}

type Item struct {
	Attr   string
	Weight float64
}
type ItemSorter []Item

// attr_weight排序,将排序结果放到s里
func AWSorter(m map[string]float64) []string {
	len := len(m)
	aw := make(ItemSorter, 0, len)
	for k, v := range m {
		aw = append(aw, Item{k, v})
		// fmt.Println("aw is ", aw)

	}

	// 排序
	sort.Slice(aw, func(i, j int) bool {
		// 权重越大说明信息增益越大，信息增益越大说明该特征对数据集的纯净度提升程度越大‌
		// 因此应该将权重最大的放在前面
		return aw[i].Weight > aw[j].Weight // 降序
		// return aw[i].Weight < aw[j].Weight // 升序
	})
	// fmt.Println("aw is ", aw)
	// 生成attr_list
	// var attrList []string
	attrList := make([]string, 0, len)
	// fmt.Println("len: ", len(attrList))
	for _, v := range aw {
		if v.Attr != "" {
			// fmt.Println("v is ", v)
			attrList = append(attrList, v.Attr)
			// fmt.Println(attrList)
		}
	}
	// s = attrList
	return attrList
}

// 信息熵
func calcEnt(dataSet []*HistoryRecord) float64 {
	// 数据行数
	num := len(dataSet)
	// fmt.Println("num is ", num)
	// 记录标签出现的次数,标签为bool，result
	labelMap := make(map[bool]int)
	for _, temp := range dataSet {
		// fmt.Println("temp is ", temp, i)
		// 定义结果
		curLabel := temp.Result
		if _, ok := labelMap[curLabel]; !ok {
			labelMap[curLabel] = 0
		}
		labelMap[curLabel]++
	}
	ent := 0.0

	// 计算经验熵
	for _, v := range labelMap {
		prob := float64(v) / float64(num)
		ent -= math.Log2(prob) * prob
	}
	// fmt.Println("label map: ", labelMap, "ent: ", ent)

	return ent
}

// 去重函数
func duplicate(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}
func distinct(val []string) []interface{} {
	sort.Strings(val)
	return duplicate(val)
}

// 划分子集
func splitDataSet(dataSet []*HistoryRecord, axis int, value string) []*HistoryRecord {
	var res []*HistoryRecord

	for _, temp := range dataSet {
		v := reflect.ValueOf(*temp)
		if strings.Compare(v.Field(axis+2).String(), value) == 0 {
			// 先复制一个切片，防止对数据集的修改
			// tar := make([]string, len(temp))
			// copy(tar, temp)
			// tar := temp
			// 这里决策树会去除掉tar[axis]，以便下次划分时不在使用该属性
			// 但是感觉我们这里没必要，先去掉此部分
			// reduceFeatVec := tar[:axis]
			// reduceFeatVec = append(reduceFeatVec, tar[axis+1:]...)
			//fmt.Println(reduceFeatVec)
			res = append(res, temp)
		}
	}

	return res
}

// 计算出每个属性的信息增益，之后调用AWSorter排列顺序
func ChooseBestFeature(pool *HistoryPool) []string {
	// 特征数量
	featureNum := pool.Pool[0].FeatureNum
	// fmt.Println("Choose")
	dataSet := pool.Pool
	// 计算数据集的熵
	baseEntropy := calcEnt(dataSet)
	// fmt.Println("-----", baseEntropy)

	// 信息增益
	// bestInfoGain := 0.0
	// 最优特征的索引值
	// bestFeatureIdx := -1
	// 遍历所有特征
	for i := 0; i < featureNum; i++ {
		// 获取当前特征的所有特征值
		var featList []string
		for _, temp := range dataSet {
			// Attr := "Attr1"
			value := reflect.ValueOf(*temp)
			// fmt.Println("value is ", value)
			featList = append(featList, value.Field(i+2).String())
			// fmt.Println("Attribute is ", value.Field(i+2))
		}
		// 获取不同的特征值
		uniqueFeatureValues := distinct(featList)
		// fmt.Println("uniqueFeatureValues is ", uniqueFeatureValues, Attr_list[i])

		// 经验条件熵
		newEntropy := 0.0
		// 计算信息增益
		for _, temp := range uniqueFeatureValues {
			// 划分子集
			subDataSet := splitDataSet(dataSet, i, temp.(string))
			// 计算子集的概率
			prob := float64(len(subDataSet)) / float64(len(dataSet))
			// 计算经验条件熵
			newEntropy += prob * calcEnt(subDataSet)
			// fmt.Println(temp.(string), prob*calcEnt(subDataSet))
		}
		// 信息增益
		infoGain := baseEntropy - newEntropy
		// 打印每个特征的信息增益
		// fmt.Printf("特征%s的增益为%.9f, uniqueFeatureValues is %s\n", Attr_list[i], infoGain, uniqueFeatureValues)
		attr_weights2[Attr_list[i]] = infoGain
		// fmt.Println("attr_weights2 is ", attr_weights2, len(attr_weights2))
		// // 计算信息增益
		// if infoGain > bestInfoGain {
		// 	// 更新信息增益，找到最大的信息增益
		// 	bestInfoGain = infoGain
		// 	// 记录信息增益最大的特征的索引
		// 	bestFeatureIdx = i
		// }
	}
	// fmt.Println(attr_weights2) // attr_list2 := Attr_list
	attr_list2 := AWSorter(attr_weights2)
	return attr_list2
}
