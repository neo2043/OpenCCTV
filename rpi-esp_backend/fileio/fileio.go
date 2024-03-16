package fileio

import (
	"encoding/json"
	"errors"
	"path"
	"os"
)

type fileSaveData struct {
	Name string `json:"name"`
	// Id   string `json:"id"`
}

type tempSaveData struct {
	// Id        string `json:"id"`
	Ip        string `json:"ip"`
	Name      string `json:"name"`
	Connected bool   `json:"connected"`
}

type File struct{
	path string
}

func (f *File) Init() error {
	value1, state := os.LookupEnv("HOMEDRIVE")
	if !state{
		return errors.New("cannot find HOMEDRIVE")
	}
	value2, state := os.LookupEnv("HOMEPATH")
	if !state{
		return errors.New("cannot find HOMEPATH")
	}
	f.path=value1+value2+"\\.OpenCCTV\\test"
	return nil
}

func (f *File) CheckPath() (bool,error) {
	if(f.path!=""){
		_,err := os.ReadFile(f.path)
		if err!=nil{
			return false,err
		}
	}else{
		return false,errors.New("fill the path")
	}
	return true,nil
}

func (f *File) MakePath() (error) {
	dir:=path.Dir(f.path)
	os.MkdirAll(dir,os.ModePerm)
	err := os.WriteFile(f.path,[]byte(""),0644)
	if err!=nil{
		return err
	}
	return nil
}


func (f *File) Write(data map[string]tempSaveData) error {
	tempfilesavestruct:=tempTosave(data)	
	temp,err:=json.Marshal(tempfilesavestruct)
	if err!=nil{
		return err
	}
	err = os.WriteFile(f.path,temp,0644)
	if err!=nil{
		return err
	}
	return nil
}

func (f *File) Read() (map[string]tempSaveData,error) {
	temptempsavedatastruct:=make(map[string]tempSaveData)
	data,err:=os.ReadFile(f.path)
	if err!=nil{
		return nil,err
	}
	err = json.Unmarshal(data,&temptempsavedatastruct)
	if err!=nil{
		return nil,err
	}
	return temptempsavedatastruct,nil
}

func keyMap[K string, V interface{}](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func tempTosave(data map[string]tempSaveData) map[string]fileSaveData {
	tempKey:=keyMap(data)
	tempReturnstruct:=make(map[string]fileSaveData)
	for _,j:=range tempKey{
		tempReturnstruct[j]=fileSaveData{Name: data[j].Name}
	}
	return tempReturnstruct
}

