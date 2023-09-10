//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package module

import (
	contextx "context"
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
	"github.com/gin-gonic/gin"
	v9 "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &eventBus_{}
		},
	})
	eventBusStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &EventBus{}
		},
		ConstructFunc: func(i interface{}, _ interface{}) (interface{}, error) {
			impl := i.(*EventBus)
			var constructFunc EventBusConstructFunc = NewEventBus
			return constructFunc(impl)
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(eventBusStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &storage_{}
		},
	})
	storageStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &Storage{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	normal.RegisterStructDescriptor(storageStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &treasury_{}
		},
	})
	treasuryStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &Treasury{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	normal.RegisterStructDescriptor(treasuryStructDescriptor)
}

type EventBusConstructFunc func(impl *EventBus) (*EventBus, error)
type eventBus_ struct {
}

type storage_ struct {
	GetRedisClient_ func() (*v9.Client, error)
	GetMysqlClient_ func() (*gorm.DB, error)
}

func (s *storage_) GetRedisClient() (*v9.Client, error) {
	return s.GetRedisClient_()
}

func (s *storage_) GetMysqlClient() (*gorm.DB, error) {
	return s.GetMysqlClient_()
}

type treasury_ struct {
	CreateRechargeOrder_      func(ctx contextx.Context, request *CreateRechargeOrderRequest) (*CreateRechargeOrderResponse, error)
	CreateRechargeOrderApi_   func(ctx *gin.Context)
	CheckRechargeOrderStatus_ func() error
}

func (t *treasury_) CreateRechargeOrder(ctx contextx.Context, request *CreateRechargeOrderRequest) (*CreateRechargeOrderResponse, error) {
	return t.CreateRechargeOrder_(ctx, request)
}

func (t *treasury_) CreateRechargeOrderApi(ctx *gin.Context) {
	t.CreateRechargeOrderApi_(ctx)
}

func (t *treasury_) CheckRechargeOrderStatus() error {
	return t.CheckRechargeOrderStatus_()
}

type EventBusIOCInterface interface {
}

type StorageIOCInterface interface {
	GetRedisClient() (*v9.Client, error)
	GetMysqlClient() (*gorm.DB, error)
}

type TreasuryIOCInterface interface {
	CreateRechargeOrder(ctx contextx.Context, request *CreateRechargeOrderRequest) (*CreateRechargeOrderResponse, error)
	CreateRechargeOrderApi(ctx *gin.Context)
	CheckRechargeOrderStatus() error
}

var _eventBusSDID string

func GetEventBusSingleton() (*EventBus, error) {
	if _eventBusSDID == "" {
		_eventBusSDID = util.GetSDIDByStructPtr(new(EventBus))
	}
	i, err := singleton.GetImpl(_eventBusSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*EventBus)
	return impl, nil
}

func GetEventBusIOCInterfaceSingleton() (EventBusIOCInterface, error) {
	if _eventBusSDID == "" {
		_eventBusSDID = util.GetSDIDByStructPtr(new(EventBus))
	}
	i, err := singleton.GetImplWithProxy(_eventBusSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(EventBusIOCInterface)
	return impl, nil
}

type ThisEventBus struct {
}

func (t *ThisEventBus) This() EventBusIOCInterface {
	thisPtr, _ := GetEventBusIOCInterfaceSingleton()
	return thisPtr
}

var _storageSDID string

func GetStorage() (*Storage, error) {
	if _storageSDID == "" {
		_storageSDID = util.GetSDIDByStructPtr(new(Storage))
	}
	i, err := normal.GetImpl(_storageSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*Storage)
	return impl, nil
}

func GetStorageIOCInterface() (StorageIOCInterface, error) {
	if _storageSDID == "" {
		_storageSDID = util.GetSDIDByStructPtr(new(Storage))
	}
	i, err := normal.GetImplWithProxy(_storageSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(StorageIOCInterface)
	return impl, nil
}

var _treasurySDID string

func GetTreasury() (*Treasury, error) {
	if _treasurySDID == "" {
		_treasurySDID = util.GetSDIDByStructPtr(new(Treasury))
	}
	i, err := normal.GetImpl(_treasurySDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*Treasury)
	return impl, nil
}

func GetTreasuryIOCInterface() (TreasuryIOCInterface, error) {
	if _treasurySDID == "" {
		_treasurySDID = util.GetSDIDByStructPtr(new(Treasury))
	}
	i, err := normal.GetImplWithProxy(_treasurySDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(TreasuryIOCInterface)
	return impl, nil
}
