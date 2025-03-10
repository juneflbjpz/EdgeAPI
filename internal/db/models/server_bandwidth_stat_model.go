package models

// ServerBandwidthStat 服务峰值带宽统计
type ServerBandwidthStat struct {
	Id                  uint64 `field:"id"`                  // ID
	UserId              uint64 `field:"userId"`              // 用户ID
	ServerId            uint64 `field:"serverId"`            // 服务ID
	RegionId            uint32 `field:"regionId"`            // 区域ID
	Day                 string `field:"day"`                 // 日期YYYYMMDD
	TimeAt              string `field:"timeAt"`              // 时间点HHMM
	Bytes               uint64 `field:"bytes"`               // 带宽字节
	AvgBytes            uint64 `field:"avgBytes"`            // 平均流量
	CachedBytes         uint64 `field:"cachedBytes"`         // 缓存的流量
	AttackBytes         uint64 `field:"attackBytes"`         // 攻击流量
	CountRequests       uint64 `field:"countRequests"`       // 请求数
	CountCachedRequests uint64 `field:"countCachedRequests"` // 缓存的请求数
	CountAttackRequests uint64 `field:"countAttackRequests"` // 攻击请求数
	TotalBytes          uint64 `field:"totalBytes"`          // 总流量
}

type ServerBandwidthStatOperator struct {
	Id                  any // ID
	UserId              any // 用户ID
	ServerId            any // 服务ID
	RegionId            any // 区域ID
	Day                 any // 日期YYYYMMDD
	TimeAt              any // 时间点HHMM
	Bytes               any // 带宽字节
	AvgBytes            any // 平均流量
	CachedBytes         any // 缓存的流量
	AttackBytes         any // 攻击流量
	CountRequests       any // 请求数
	CountCachedRequests any // 缓存的请求数
	CountAttackRequests any // 攻击请求数
	TotalBytes          any // 总流量
}

func NewServerBandwidthStatOperator() *ServerBandwidthStatOperator {
	return &ServerBandwidthStatOperator{}
}
