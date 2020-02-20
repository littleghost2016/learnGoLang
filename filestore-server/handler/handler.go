package handler

import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"net/http"
	"time"
	"encoding/json"
	// "bufio"

	"filestore-server/meta"
	"filestore-server/util"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 返回上传HTML页面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "Internal server error\n")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		// 接收文件六及存储到本地目录
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get data, err: %s\n", err.Error())
			return
		}
		defer file.Close()

		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			Location: ".\\tmp\\" + head.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		

		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			fmt.Printf("Failed to create file, err: %s\n", err.Error())
			return
		}
		defer newFile.Close()

		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Failed to save data into file. err: %s", err.Error())
			return
		}

		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		meta.UpdateFileMeta(fileMeta)


		// filePath := ".\\tmp\\text.txt"
		// tempFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
		// if err != nil {
		// 	fmt.Println("文件打开失败", err)
		// }
		// //及时关闭file句柄
		// defer tempFile.Close()
		// //写入文件时，使用带缓存的 *Writer
		// write := bufio.NewWriter(tempFile)
		// write.WriteString(fileMeta.FileSha1)
		// //Flush将缓存的文件真正写入到文件中
		// write.Flush()


		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

// 上传已完成
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload finished.")
}

// 获取文件元信息
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	filehash := r.Form["filehash"][0]
	fMeta := meta.GetFileMeta(filehash)
	data, err := json.Marshal(fMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fsha1 := r.Form.Get("filehash")
	fm := meta.GetFileMeta(fsha1)

	f, err := os.Open(fm.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("Content-disposition", "attachment; filename=\""+fm.FileName+"\"")
	w.Write(data)
}