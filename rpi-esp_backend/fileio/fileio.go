package fileio

import (
	structs "backend/structs"
	"encoding/json"
	"errors"
	"os"
	"path"
)

type File struct {
	path string
}

func (f *File) Init() error {
	value1, state := os.LookupEnv("HOMEDRIVE")
	if !state {
		return errors.New("cannot find HOMEDRIVE")
	}
	value2, state := os.LookupEnv("HOMEPATH")
	if !state {
		return errors.New("cannot find HOMEPATH")
	}
	f.path = value1 + value2 + "\\.OpenCCTV\\test"
	return nil
}

func (f *File) CheckPath() (bool, error) {
	if f.path != "" {
		_, err := os.ReadFile(f.path)
		if err != nil {
			return false, err
		}
	} else {
		return false, errors.New("fill the path")
	}
	return true, nil
}

func (f *File) MakePath() error {
	dir := path.Dir(f.path)
	os.MkdirAll(dir, os.ModePerm)
	err := os.WriteFile(f.path, []byte(""), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (f *File) Write(data structs.TempSaveDataMap) error {
	tempfilesavestruct := structs.ConvTempSaveDataMaptoFileSaveDataMap(data)
	temp, err := json.Marshal(tempfilesavestruct)
	if err != nil {
		return err
	}
	err = os.WriteFile(f.path, temp, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (f *File) Read() (structs.TempSaveDataMap, error) {
	tempTempsavedatastruct := make(structs.TempSaveDataMap)
	data, err := os.ReadFile(f.path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tempTempsavedatastruct)
	if err != nil {
		return nil, err
	}
	return tempTempsavedatastruct, nil
}