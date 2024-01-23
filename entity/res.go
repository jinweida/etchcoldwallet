package entity

var Res *Response

func init() {
	Res = &Response{}
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r *Response) setCode(code int) *Response {
	r.Code = code
	return r
}

func (r *Response) Ok() *Response {
	r.Code = 1
	r.Msg = "success"
	return r
}
func (r *Response) Error() *Response {
	r.Code = 0
	return r
}
func (r *Response) SetMsg(msg string) *Response {
	r.Msg = msg
	return r
}

func (r *Response) Result(data interface{}) *Response {
	r.Data = data
	return r
}

//home区块运行状况
type StateEntity struct {
	LastTime      int64   `json:"time"`
	CurrentBlock  int64   `json:"block"`
	Tps           float64 `json:"tps"`
	Ctps          float64 `json:"ctps"`
	MaxTps        float64 `json:"maxtps"`
	MaxCtps       float64 `json:"maxctps"`
	MaxTxs        float64 `json:"maxtxs"`
	Txs           float64 `json:"txs"`
	BlockDay      int     `json:"day"`
	BlockDuration int64   `json:"duration"`
	PeakCtps      float64 `json:"peakctps"`
	PeakTps       float64 `json:"peaktps"`
}

//chart
type ChartEntity struct {
	//当前事务接收量
	RecTxs float64 `json:"rec_txs"`
	//平均块中事务量
	MeanRecTxs float64 `json:"mean_rec_txs"`
	//当前区块生成速度
	LastInterval int64 `json:"last_interval"`
	//当前区块中事务量
	LastTxs int64 `json:"last_txs"`
	//事务堆积
	LastQueueSize int64 `json:"last_queuesize"`
	//当前事务上链时间
	PreNum float64 `json:"prenum"`
	//当前事务上链完成率
	TxsRate float64 `json:"txs_rate"`
	//平均区块生成速度
	MeanInterval float64 `json:"mean_interval"`
	//平均块中事务量
	MeanTxs float64 `json:"mean_txs"`
	//平均上链时间
	MeanPreNum float64 `json:"mean_pernum"`
	MeanAci    float64 `json:"mean_aci"`
	//平均事务上链完成率
	MeanTxsRate float64 `json:"mean_txs_rate"`
	//当前上链事务量
	UpNum float64 `json:"up_num"`
	//平均上链量
	MeanUpNum float64 `json:"mean_up_num"`
	//平均事务大小
	MeanTxSize float64 `json:"mean_txsize"`
}
