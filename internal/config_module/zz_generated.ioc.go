//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package config_module

import (
	"github.com/DwGoing/funds-system/internal/shared"
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &configModule_{}
		},
	})
	configModuleStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &ConfigModule{}
		},
		ConstructFunc: func(i interface{}, _ interface{}) (interface{}, error) {
			impl := i.(*ConfigModule)
			var constructFunc ConfigModuleConstructFunc = NewConfigModule
			return constructFunc(impl)
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(configModuleStructDescriptor)
}

type ConfigModuleConstructFunc func(impl *ConfigModule) (*ConfigModule, error)
type configModule_ struct {
	Load_ func() (*shared.Configs, error)
	Set_  func(key string, value any) error
}

func (c *configModule_) Load() (*shared.Configs, error) {
	return c.Load_()
}

func (c *configModule_) Set(key string, value any) error {
	return c.Set_(key, value)
}

type ConfigModuleIOCInterface interface {
	Load() (*shared.Configs, error)
	Set(key string, value any) error
}

var _configModuleSDID string

func GetConfigModuleSingleton() (*ConfigModule, error) {
	if _configModuleSDID == "" {
		_configModuleSDID = util.GetSDIDByStructPtr(new(ConfigModule))
	}
	i, err := singleton.GetImpl(_configModuleSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*ConfigModule)
	return impl, nil
}

func GetConfigModuleIOCInterfaceSingleton() (ConfigModuleIOCInterface, error) {
	if _configModuleSDID == "" {
		_configModuleSDID = util.GetSDIDByStructPtr(new(ConfigModule))
	}
	i, err := singleton.GetImplWithProxy(_configModuleSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(ConfigModuleIOCInterface)
	return impl, nil
}

type ThisConfigModule struct {
}

func (t *ThisConfigModule) This() ConfigModuleIOCInterface {
	thisPtr, _ := GetConfigModuleIOCInterfaceSingleton()
	return thisPtr
}
