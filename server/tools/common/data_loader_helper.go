package common

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// GetLocalFileNamePath 取得本地端檔案路徑
func GetLocalFileNamePath(fileName string) string {
	workingDirPath := GetCurrentDirectory()
	var jsonPath = fmt.Sprintf("%s/%s.json", workingDirPath, fileName)
	return jsonPath
}

// LoadJsonDataByFileName 透過傳入的檔名與結構取得參數檔
func LoadJsonDataByFileName(dataType interface{}, fileName string) interface{} {
	dataJson, err := os.Open(GetFilePath(fileName))
	if err != nil {
		panic(err)
	}
	defer dataJson.Close()

	byteValue, _ := ioutil.ReadAll(dataJson)
	err = json.Unmarshal(byteValue, &dataType)
	if err != nil {
		panic(err) //無法取得參數檔直接報錯結束程式
	}
	return dataType
}

// GetCurrentDirectory 取得目前執行檔的路徑
func GetCurrentDirectory() string {
	//返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	//log.Info(ctx,"dir :", dir)
	//将\替换成/
	outPath := strings.Replace(dir, "\\", "/", -1)
	//log.Info(ctx,"outPath :", outPath)
	return outPath
}

// GetFilePath 取得檔案路徑
func GetFilePath(fileName string) string {

	_, workingDirPath, _, ok := runtime.Caller(0)
	if !ok {
		panic("can't find file.")
	}
	//workingDirPath = filepath.Join(filepath.Dir(workingDirPath), "../../common/json")
	var jsonPath = fmt.Sprintf("/%s.json", fileName)
	outPath := fmt.Sprintf("%s%s", workingDirPath, jsonPath)

	return outPath
}

// GetDirectoryPath 取得目錄路徑
func GetDirectoryPath(dirName string) string {
	_, workingDirPath, _, ok := runtime.Caller(0)
	if !ok {
		panic("can't find directory.")
	}

	workingDirPath = filepath.Join(filepath.Dir(workingDirPath), "../../")
	var dirPath = fmt.Sprintf("/%s", dirName)
	outPath := fmt.Sprintf("%s%s", workingDirPath, dirPath)
	return outPath
}

// GetFilesInDirectory 取得目錄下的檔案
func GetFilesInDirectory(dir string) []fs.FileInfo {
	dir = GetDirectoryPath(dir)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err.Error())
	}
	return files
}

// GetFileShortName 取得檔案短名稱
func GetFileShortName(fileName string) string {
	names := strings.Split(fileName, ".")
	return names[0]
}

// LoadJsonData 透過傳入的檔名與結構取得參數檔資料
func LoadJsonData(fileName string, loadData interface{}) {
	path := GetJsonFilePathFromWorkingDir(fileName, true)
	jsonFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &loadData)
	if err != nil {
		panic(err)
	}
}

// LoadJsonFile 透過傳入的檔名取得參數檔資料
func LoadJsonFile(fileName string) ([]byte, error) {
	path := GetFilePath(fileName)
	jsonFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return ioutil.ReadAll(jsonFile)
}

func GetJsonFilePathFromWorkingDir(fileName string, isForceCreate bool) string {
	wd, err := os.Getwd()
	if err != nil {
		panic("can't find file.")
	}

	var jsonPath = fmt.Sprintf("/%s.json", fileName)
	outPath := fmt.Sprintf("%s%s", wd, jsonPath)

	//check file exist
	_, err = os.Stat(outPath)
	if os.IsNotExist(err) {
		if !isForceCreate {
			panic("file not exist:" + outPath)
		}
		//create directory
		dir := filepath.Dir(outPath)
		_, err := os.Stat(dir)
		if os.IsNotExist(err) {
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
		//create file
		file, err := os.Create(outPath)
		if err != nil {
			panic(err)
		} else {
			//write empty json object
			_, err := file.WriteString("{}")
			if err != nil {
				panic(err)
			}
			file.Close()
		}
	}

	return outPath
}

func WriteJsonData(fileName string, data interface{}) {
	path := GetJsonFilePathFromWorkingDir(fileName, true)
	jsonFile, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}

	_, err = jsonFile.Write(byteValue)
	if err != nil {
		panic(err)
	}
}
