package utils

import (
	"go-jocy/config"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	lua "github.com/yuin/gopher-lua"
)

// NewLuaState 创建一个新的Lua状态
func NewLuaState() *lua.LState {
	return lua.NewState()
}

// RegisterLuaUtils 注册工具函数、设备信息、httpGet、json等到Lua环境
func RegisterLuaUtils(L *lua.LState, AuthIP string) {
	// 工具函数
	L.SetGlobal("utils", L.NewTable())
	L.SetField(L.GetGlobal("utils"), "md5", L.NewFunction(func(L *lua.LState) int {
		input := L.ToString(1)
		result := MD5Encryption(input)
		L.Push(lua.LString(result))
		return 1
	}))

	L.SetField(L.GetGlobal("utils"), "timestamp", L.NewFunction(func(L *lua.LState) int {
		ts := strconv.FormatInt(time.Now().Unix(), 10)
		L.Push(lua.LString(ts))
		return 1
	}))

	L.SetField(L.GetGlobal("utils"), "aes128cbc_decrypt", L.NewFunction(func(L *lua.LState) int {
		key := L.ToString(1)
		iv := L.ToString(2)
		encryptedText := L.ToString(3)
		decrypted, err := AesDecryption(encryptedText, key, iv)
		if err != nil {
			L.Push(lua.LString(""))
			return 1
		}
		L.Push(lua.LString(decrypted))
		return 1
	}))

	// 设备信息
	L.SetGlobal("device_info", L.NewTable())
	L.SetField(L.GetGlobal("device_info"), "platform", lua.LString("Android"))
	L.SetField(L.GetGlobal("device_info"), "app_version", lua.LString(config.GinConfig.App.AppVersion))

	// httpGet
	L.SetGlobal("httpGet", L.NewFunction(func(L *lua.LState) int {
		url := L.ToString(1)
		options := L.ToTable(2)
		client := resty.New()
		client.SetRetryCount(3)
		client.SetRetryWaitTime(time.Second / 2)
		req := client.R()
		if AuthIP != "" {
			req.SetHeader("X-Forwarded-For", AuthIP)
			req.SetHeader("X-Real-IP", AuthIP)
			req.SetHeader("True-Client-IP", AuthIP)
			req.SetHeader("Client-IP", AuthIP)
		}
		if options != nil {
			headerTable := options.RawGetString("header")
			if headerTable, ok := headerTable.(*lua.LTable); ok {
				headerTable.ForEach(func(k, v lua.LValue) {
					req.SetHeaderVerbatim(k.String(), v.String())
				})
			}
		}
		resp, err := req.Get(url)
		if err != nil {
			L.Push(lua.LString(""))
			return 1
		}
		L.Push(lua.LString(resp.String()))
		return 1
	}))

	// json模块
	jsonMod := L.NewTable()
	L.SetGlobal("json", jsonMod)
	L.SetField(jsonMod, "decode", L.NewFunction(func(L *lua.LState) int {
		jsonStr := L.ToString(1)
		var result interface{}
		err := json.Unmarshal([]byte(jsonStr), &result)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LNumber(0))
			L.Push(lua.LString(err.Error()))
			return 3
		}
		luaObj := valueToLua(L, result)
		L.Push(luaObj)
		L.Push(lua.LNumber(len(jsonStr) + 1))
		L.Push(lua.LNil)
		return 3
	}))
	L.SetField(jsonMod, "encode", L.NewFunction(func(L *lua.LState) int {
		value := L.CheckAny(1)
		goValue := luaToValue(L, value)
		jsonBytes, err := json.Marshal(goValue)
		if err != nil {
			L.Push(lua.LString(""))
			return 1
		}
		L.Push(lua.LString(jsonBytes))
		return 1
	}))
}
