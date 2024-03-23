package structs

type FileSaveData struct {
	Name    string `json:"name"`
	UrlPath string `json:"urlpath"`
}

type TempSaveData struct {
	Ip        string `json:"ip"`
	Name      string `json:"name"`
	UrlPath   string `json:"urlpath"`
	Connected bool   `json:"connected"`
}

type TempSaveDataMap map[string]TempSaveData
type FileSaveDataMap map[string]FileSaveData

func (t *TempSaveData) UpdateIP(ip string) TempSaveData {
	return TempSaveData{
		Ip:        ip,
		Name:      t.Name,
		UrlPath:   t.UrlPath,
		Connected: t.Connected,
	}
}

func (t *TempSaveData) UpdateName(name string) TempSaveData {
	return TempSaveData{
		Ip:        t.Ip,
		Name:      name,
		UrlPath:   t.UrlPath,
		Connected: t.Connected,
	}
}

func (t *TempSaveData) UpdateUrlPath(urlpath string) TempSaveData {
	return TempSaveData{
		Ip:        t.Ip,
		Name:      t.Name,
		UrlPath:   urlpath,
		Connected: t.Connected,
	}
}

func (t *TempSaveData) UpdateConnected(connected bool) TempSaveData {
	return TempSaveData{
		Ip:        t.Ip,
		Name:      t.Name,
		UrlPath:   t.UrlPath,
		Connected: connected,
	}
}

func GetTempSaveDataStructMap() map[string]TempSaveData {
	return make(map[string]TempSaveData)
}

func GetFileSaveDataStructMap() map[string]FileSaveData {
	return make(map[string]FileSaveData)
}

func (t TempSaveDataMap) UpdateTempSaveDataMap(ip, name, urlpath, connected interface{}, key string) {
	tempStruct := t[key]
	if ip != nil {
		tempStruct=tempStruct.UpdateIP(ip.(string))
	}
	if name != nil {
		tempStruct=tempStruct.UpdateName(name.(string))
	}
	if urlpath != nil {
		tempStruct=tempStruct.UpdateUrlPath(urlpath.(string))
	}
	if connected != nil {
		tempStruct=tempStruct.UpdateConnected(connected.(bool))
	}
	t[key] = tempStruct
}

func keyMap[K string, V interface{}](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func ConvTempSaveDataMaptoFileSaveDataMap(data TempSaveDataMap) FileSaveDataMap {
	tempKey := keyMap(data)
	tempReturnstruct := make(FileSaveDataMap)
	for _, j := range tempKey {
		tempReturnstruct[j] = FileSaveData{Name: data[j].Name,UrlPath: data[j].UrlPath}
	}
	return tempReturnstruct
}