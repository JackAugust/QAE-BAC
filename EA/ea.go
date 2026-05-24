package EA

// type DC struct {
// 	ID      string `json:"id"`
// 	EmeList *Queue `json:"-"`
// 	succNum int
// 	faliNum int
// 	// salveChain *service.ServiceSetup
// }

// func NewDC(id string, num ...int) *DC {
// 	succNum := 0
// 	failNum := 0
// 	list := 2
// 	if len(num) == 3 {
// 		succNum = num[0]
// 		failNum = num[1]
// 		list = num[2]
// 	} else if len(num) == 1 {
// 		succNum = num[0]
// 	} else if len(num) == 2 {
// 		succNum = num[0]
// 		failNum = num[1]
// 	}
// 	return &DC{
// 		ID:      id,
// 		succNum: succNum,
// 		faliNum: failNum,
// 		EmeList: NewQueue(list),
// 		// salveChain: nil,
// 	}
// }
// func (dc *DC) GetID() string {
// 	return dc.ID
// }

// // func (dc *DC) CombineWithServ(serv *service.ServiceSetup) {
// // 	dc.salveChain = serv
// // }
// // func (dc *DC) GetServ() *service.ServiceSetup {
// // 	return dc.salveChain
// // }
// func (dc *DC) IncreaseSuc() {
// 	dc.succNum++
// }

// func (dc *DC) DecreaseFail() {
// 	dc.faliNum++
// }

// // 紧急委派
// // 当A的救活率不足以救治时，调用此函数
// func (dc *DC) EmerAss(token *abac.MyToken, req *abac.ABACRequest) (*abac.MyToken, error) {
// 	// 从急救列表中获得可急救对象
// 	rand.Seed(time.Now().UnixNano())

// 	// fmt.Println("emerDC is ", emerDC)
// 	// 注意，这里emerDC为id，这样设计是正确合理的，因为你不能直接获得人家DC的详细数据
// 	//  TODO: 这里的ID暂时直接用了DC，后期改成每个小dc上
// 	delToken, err := abac.DelegateToken(token, req.Sub.UID, req.Recipient)
// 	if err != nil {
// 		// fmt.Println("委派失败：", err.Error())
// 		return nil, err
// 	}
// 	return delToken, nil
// }

// // 救活率计算
// // 每个DC都有一个且结果只有DC自己可以查看
// // 这个合约部署在每个DC中
// func (dc *DC) CalSurRate() int64 {
// 	surRate := float64(dc.succNum) / float64((dc.succNum + dc.faliNum)) * 100
// 	return int64(surRate)
// }

// // 返回的是DCid，还需要返回给具体的数据
// func (dc *DC) GetEmerDC() (string, error) {
// 	ResponseData := dc.EmeList.Dequeue()
// 	if ResponseData == nil {
// 		return "", fmt.Errorf("no emerDC to get")
// 	}
// 	id := fmt.Sprintf("%s", ResponseData)
// 	return id, nil
// }

// func (dc *DC) AddEmerDC(emerDC *DC) {
// 	// n := dc.CalSurRate()
// 	m := emerDC.CalSurRate()
// 	goal := dc.CalSurRate()
// 	flag := zk.ZK(goal, m)
// 	if flag {
// 		// TODO: 这里是id吗, 暂时写成id。如果后期需要改再改-----改成dc了 20231114
// 		// fmt.Println("emerDC.ID is ", emerDC.ID)
// 		dc.EmeList.Enqueue(emerDC.ID)
// 	}
// }
